run:
	cp wazuh.rego example.rego
	opa build -t wasm -e 'wazuh/ignore' ./example.rego
	tar zxf bundle.tar.gz
	go run main.go
