require 'faraday'
require 'faraday_middleware'

class StreamClient
  attr_reader :chunks, :last_id

  def initialize
    @chunks = []
    @last_id = nil
  end

  def client
    Faraday.new(url: 'http://localhost:8080/') do |faraday|
      faraday.request :json
      faraday.adapter :net_http
      faraday.response :chunked
    end
  end

  def some_hook(event)
    p event
  end

  def join_chunk
    Proc.new { |chunk, _|
      @chunks << chunk.strip

      begin
        event = JSON.parse @chunks.join
        @last_id = event['id']
        @chunks.clear

        some_hook(event)
      rescue JSON::ParserError
        p "parse error", @chunks.join
      end
    }
  end

  def fetch
    client.get do |req|
      req.url '/'

      req.options.on_data = join_chunk
    end
  end
end
