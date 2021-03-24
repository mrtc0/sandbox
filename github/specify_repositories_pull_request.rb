require 'octokit'

Octokit.configure do |c|
  c.api_endpoint = ENV['GITHUB_API_ENDPOINT']
end

client = Octokit::Client.new(:access_token => ENV['GITHUB_TOKEN'])
client.create_pull_request(repo, base, head, title, body)
