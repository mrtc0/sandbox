require 'webrick'

srv = WEBrick::HTTPServer.new({
  DocumentRoot: './',
  BindAddress: '0.0.0.0',
  Port: 3000
})

srv.mount_proc '/' do |req, res|
  res.body = req.query['q']
end

srv.start
