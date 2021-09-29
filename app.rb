# frozen_string_literal: true

require 'net/http'
require 'uri'

class LineNotification
  TOKEN = ''.freeze
  URL = 'https://notify-api.line.me/api/notify'.freeze

  attr_reader :message

  def self.send(message)
    new(message).send
  end

  def initialize(message)
    @message = message
  end

  def send
    Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
      http.request(request)
    end
  end

  private def request
    req = Net::HTTP::Post.new(uri)
    req['Authorization'] = "Bearer #{TOKEN}"
    req.set_form_data(message: message)
    req
  end

  private def uri
    URI.parse(URL)
  end
end

LineNotification.send('Line message sent by okubo')
