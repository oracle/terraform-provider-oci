
# Get Environment Variables
os_user = node['omc']['os_user']
os_installer_group = node['omc']['os_installer_group']
omc_entity  = File.join(node['omc']['stage'], 'omc_entity.json')
hostname = node['fqdn']

#Load linux monitoring entity template
template omc_entity do
  source "omc_entity.json.erb"
  variables({
                :hostname 	=> hostname
            })
end

#Call CLI to add linux monitoring entity
add_entity = "#{node['omc']['app']}/cloud_agent/agent_inst/bin/omcli update_entity agent  #{node['omc']['stage']}/omc_entity.json"
execute add_entity do
  action :run
  cwd node['omc']['stage']
  user os_user
  group os_installer_group
end
