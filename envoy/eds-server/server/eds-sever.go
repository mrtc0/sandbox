package main

import (
	"log"
	"net"
	"context"
	"strconv"
	"time"

	v3core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3server "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	v3service "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	v3endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	v3cache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	xdslog "github.com/envoyproxy/go-control-plane/pkg/log"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"

	"google.golang.org/grpc"
)

var clusterUpstreams = map[string][]struct {
	Address string
	Port    uint32
}{
	"nginx_cluster": {{"nginx1", 80}, {"nginx2", 80}},
	"httpd_cluster": {{"httpd1", 80}, {"httpd2", 80}},
}

const serverAddr = ":20000"
const node = "node0"

func newSnapshotVersion() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// Envoy が xDS API を呼び出したときの callback 関数
// ストリームのリクエスト/レスポンスのタイミングでログを出すだけ
// https://github.com/newrelic/sidecar/blob/ff3e175b802490b07af0514d43dd11cc6f9d1e9e/envoy/server.go#L29-L41
type callbacks struct{}

func (cb *callbacks) OnFetchRequest(ctx context.Context, req *discovery.DiscoveryRequest) error {
    log.Printf("OnFetchRequest")
    return nil
}

func (*callbacks) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
	log.Printf("OnFetchResponse")
}

func (cb *callbacks) OnStreamOpen(ctx context.Context, streamID int64, typeURL string) error {
    log.Printf("OnStreamOpen: StreamID [%d], Type URL [%s]", streamID, typeURL)
    return nil
}

func (cb *callbacks) OnStreamClosed(streamID int64) {
    log.Printf("OnStreamClosed: StreamID [%d]", streamID)
}

func (cb *callbacks) OnStreamRequest(streamID int64, req *discovery.DiscoveryRequest) error {
    log.Printf("OnStreamRequest: StreamID [%d]", streamID)
    return nil
}

func (cb *callbacks) OnStreamResponse(streamID int64, req *discovery.DiscoveryRequest, resp *discovery.DiscoveryResponse) {
    log.Printf("OnStreamResponse: StreamID [%d]", streamID)
}

// clutser の EDS 情報を返す
func getEdsResourceForCluster(cluster string) *v3endpoint.ClusterLoadAssignment {
	upstreams, ok := clusterUpstreams[cluster]	
	if !ok { return nil }

	lbEndpoints := []*v3endpoint.LbEndpoint{}
	for _, u := range upstreams {
		addr := &v3core.Address{
			Address: &v3core.Address_SocketAddress{
				SocketAddress: &v3core.SocketAddress{
					Protocol: v3core.SocketAddress_TCP,
					Address:  u.Address,
					PortSpecifier: &v3core.SocketAddress_PortValue{
						PortValue: u.Port,
					},
				},
			},
		}

		lbEndpoint := &v3endpoint.LbEndpoint{
			HostIdentifier: &v3endpoint.LbEndpoint_Endpoint{
				Endpoint: &v3endpoint.Endpoint{
					Address: addr,
				},
			},
		}

		lbEndpoints = append(lbEndpoints, lbEndpoint)
	}

	assignment := &v3endpoint.ClusterLoadAssignment{
		ClusterName: cluster,
		Endpoints: []*v3endpoint.LocalityLbEndpoints{
			{LbEndpoints: lbEndpoints},
		},
	}

	return assignment
}

func getEdsResource() []types.Resource {
	resources := []types.Resource{}
	for cluster := range clusterUpstreams {
		r := getEdsResourceForCluster(cluster)
		if r != nil {
			resources = append(resources, r)
		}
	}
	return resources
}

// スナップショットを作成する
func getSnapshot() v3cache.Snapshot {
	eds := getEdsResource()
	log.Printf("EDS : %#v", eds)

	return v3cache.NewSnapshot(
		newSnapshotVersion(), eds, nil, nil, nil, nil, nil,
	)
}

func main() {
	var xdsLogger = xdslog.LoggerFuncs{
		DebugFunc: log.Printf,
		InfoFunc:  log.Printf,
		WarnFunc:  log.Printf,
		ErrorFunc: log.Printf,
	}

	// Enovy へのレスポンスを保存するキャッシュを作成
	cache := v3cache.NewSnapshotCache(true, v3cache.IDHash{}, xdsLogger)

	// node 名をキーにしてスナップショットを作成してキャッシュに保存する
	ss := getSnapshot()
	if err := cache.SetSnapshot(node, ss); err != nil {
		log.Fatalf("Failed SetSnapshot: %s", err)
	}

	// xdsServer の生成
	xdsServer := v3server.NewServer(context.Background(), cache, &callbacks{})
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen at: %s", err)
	}

	// gRPC サーバーの起動
	grpcServer := grpc.NewServer()

	v3service.RegisterEndpointDiscoveryServiceServer(grpcServer, xdsServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve GRPC server : %v", err)
	}
}
