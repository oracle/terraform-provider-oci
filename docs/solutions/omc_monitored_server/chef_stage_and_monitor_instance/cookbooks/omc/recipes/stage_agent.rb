
#
# Cookbook Name:: oci-server
# Recipe:: bootstrap-omc
#
# Copyright (c) 2016 The Authors, All Rights Reserved.

include_recipe 'sudo'

# User and group creation
omc_installer_group = node['omc']['installer_group']

omc = data_bag_item("users", "omc")  # The id specified in the json
omc_user = omc['id']
group omc_installer_group do
  action :create
  append true
end

users_manage omc_installer_group  do
  group_id 1139
  action [ :create ]
  data_bag 'users'
end

sudo 'omc' do
  user omc_user
  nopasswd true
end


user_ulimit omc_user do
  filehandle_limit 8192 # optional
  filehandle_soft_limit 8192 # optional; not used if filehandle_limit is set)
  filehandle_hard_limit 8192 # optional; not used if filehandle_limit is set)
  process_limit 61504 # optional
  process_soft_limit 61504 # optional; not used if process_limit is set)
  process_hard_limit 61504 # optional; not used if process_limit is set)
  memory_limit 1024 # optional
  core_limit 2048 # optional
  core_soft_limit 1024 # optional
  core_hard_limit 'unlimited' # optional
  stack_soft_limit 2048 # optional
  stack_hard_limit 2048 # optional
end

reset_command = "ulimit -a"
run_reset = "su -l #{omc_user} -c '#{reset_command}'"
#update_profile = "echo umask 022 >> ~#{os_user}/.bash_profile"

execute run_reset do
 action :run
end

# Install packages
node['omc']['packages'].each do |package|
  package_title = "#{package['name']}.#{package['arch']}"
  yum_package package_title do
    action :install
  end
end

# Create omc Directories
directory node['omc']['install'] do
  owner omc_user
  group omc_installer_group
  recursive true
  action :create
end

installer = File.join(node['omc']['install'], 'AgentInstall.zip')

cookbook_file installer do
  source 'AgentInstall.zip'
  owner omc_user
  group omc_installer_group
  mode '0755'
  action :create
end

# Create omc Directories
directory node['omc']['stage'] do
  owner omc_user
  group omc_installer_group
  recursive true
  action :create
end

# Create omc Directories
directory node['omc']['app'] do
  owner omc_user
  group omc_installer_group
  recursive true
  action :create
end

# Create omc Directories
directory node['omc']['apm'] do
  owner omc_user
  group omc_installer_group
  recursive true
  action :create
end

unzip_installer = "unzip AgentInstall.zip -d #{node['omc']['install']}"
execute unzip_installer do
  action :run
  cwd node['omc']['install']
  user omc_user
  group omc_installer_group
  creates 'AgentInstall.sh'
end

change_mode = "sudo chmod +x AgentInstall.sh"
execute change_mode do
  action :run
  cwd node['omc']['install']
  user omc_user
  group omc_installer_group
end

agent  = File.join(node['omc']['stage'], 'AgentInstall.sh')
download_installer = "./AgentInstall.sh AGENT_TYPE=cloud_agent STAGE_LOCATION=#{node['omc']['stage']} -download_only AGENT_REGISTRATION_KEY=#{node['omc']['regkey']}"
execute download_installer do
  action :run
  cwd node['omc']['install']
  user omc_user
  group omc_installer_group
  creates agent
end

