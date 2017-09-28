#
# Cookbook:: example_webserver
# Recipe:: default
#
# Copyright:: 2017, The Authors, All Rights Reserved.

package "httpd" do
  action :install
end

service "httpd" do
  action [:enable, :start]
end

firewall 'default'

# enable platform default firewall
firewall 'default' do
  action :install
end

firewall_rule 'http' do
  port 80
  command :allow
end

firewall_rule 'ssh' do
  port 22
  command :allow
end

# create an example index page on the web server
file '/var/www/html/index.html' do
  content 'Hello World!'
end
