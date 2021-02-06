require 'webrick'

class EventStream
  def initialize(queue)
    @queue = queue
  end

  def next_chunk
    chunk = @queue.pop 
    raise EOFError if chunk.nil?

    chunk
  end

  def close
  end

  def readpartial(size, buf='')
    sleep 2
    buf.clear
    buf << next_chunk
    buf
  end
end

srv = WEBrick::HTTPServer.new({ Port: 8080 })

queue = IO.readlines("events.json")

srv.mount_proc '/' do |req, res|
  res.chunked = true
  res.body = EventStream.new(queue)
end

srv.start
