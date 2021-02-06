require 'spec_helper'

describe StreamClient do
  let(:client) { described_class.new }
  let(:chunks) {
    [
      '{"id": 1, "data": "data-1"}',
      '{"id": 2, "data": "data-2"}',
      '{"id": 3, "data": "data-3"}',
    ]
  }

  describe "#join_chunk" do
    it "some_book is called 3 times" do
      allow(client).to receive(:some_hook).and_return(true)

      chunks.each do |c|
        client.join_chunk.call(c, c.size)
      end

      expect(client).to have_received(:some_hook).exactly(3).times
    end

    context "duty chunks" do
      let(:duty_chunks) {
        [
          '{"id": 1, "data": "data-1"',
          '}',
          '{"id": 2, "data":',
          ' "data-2"}',
        ]
      }

      it "some_book is called 2 times" do
        allow(client).to receive(:some_hook).and_return(true)

        duty_chunks.each do |c|
          client.join_chunk.call(c, c.size)
        end

        expect(client).to have_received(:some_hook).exactly(2).times
      end
    end
  end
end
