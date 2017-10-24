
# User and group creation
os_user = node['omc']['os_user']
os_installer_group = node['omc']['os_installer_group']
hostname = node['fqdn']

install_agent = "./AgentInstall.sh AGENT_TYPE=cloud_agent AGENT_BASE_DIR='/omc/app/cloud_agent' -staged AGENT_PORT=#{node['omc']['port']} AGENT_PROPERTIES=$PWD/agent.properties AGENT_REGISTRATION_KEY=#{node['omc']['regkey']} ORACLE_HOSTNAME=#{node['fqdn']}"
omcli = "#{node['omc']['app']}/cloud_agent/agent_inst/bin/omcli"

execute install_agent do
  action :run
  cwd node['omc']['stage']
  user os_user
  group os_installer_group
  creates omcli
end

script 'configure_cli' do
  interpreter "bash"
  cwd node['omc']['stage']
  code <<-EOH
      export core=$(/omc/app/cloud_agent/agent_inst/bin/omcli status agent | grep '^Binaries Location' | awk -F: '{print $2}')
      sudo $core/root.sh
    EOH
  not_if { ::File.exist?("omc/app/cloud_agent/agent_inst/bin/omcli") }
end


apm_agent_zip = File.join(node['omc']['apm'], 'ApmAgent-1.16.zip')
download_apm = "./AgentInstall.sh AGENT_TYPE=apm_java_as_agent STAGE_LOCATION=#{node['omc']['apm']} AGENT_REGISTRATION_KEY=#{node['omc']['regkey']}"
execute download_apm do
  action :run
  cwd node['omc']['stage']
  user os_user
  group os_installer_group
  creates apm_agent_zip
end

apm_provisioner = File.join(node['omc']['apm'], 'ProvisionApmJavaAsAgent.sh')
unzip_apm = "unzip AgentInstall.zip -d #{node['omc']['apm']}"
execute unzip_apm do
  action :run
  cwd node['omc']['apm']
  user os_user
  group os_installer_group
  creates apm_provisioner
end

change_mode = "chmod +x ProvisionApmJavaAsAgent.sh"
execute change_mode do
  action :run
  cwd node['omc']['apm']
  user os_user
  group os_installer_group
end
