# OMC Cookbook

This cookbook automates the installation and configuration of the Oracle Management Cloud Service monitoring agent on Oracle Cloud Infrastructure Infrastucture.

## Requirements

- Subscription to the Oracle Management Cloud Service  -https://cloud.oracle.com/management 
- Download the AgentInstall.zip - https://docs.oracle.com/en/cloud/paas/management-cloud/emcad/deploying-oracle-management-cloud-agents.html
- Generate the registration key for your agent - https://docs.oracle.com/en/cloud/paas/management-cloud/emcad/managing-registration-keys.html 

### Platforms

- Oracle Linux 7.x
- Centos 7.x

### Chef

- Chef 12.0 or later

### Cookbooks

- `install_omc_agent` - Creates OMC install user, and loads & stages the AgentInstall.zip on your instance.  
- `monitor_linux` - Uses the staged OMC agent, installs the agent and automatically configures the agent 
                                  to monitor Linux.

## Attributes

### omc::omc

<table>
  <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Description</th>
    <th>Default</th>
  </tr>
  <tr>
    <td><tt>['omc']['reg_key']</tt></td>
    <td>String</td>
    <td>OMC agent registration key.  This key is retrieved from your OMC cloud account - https://docs.oracle.com/en/cloud/paas/management-cloud/emcad/managing-registration-keys.html</td>
    <td><tt>omc</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['os_user']</tt></td>
    <td>String</td>
    <td>OMC Install User Name</td>
    <td><tt>omc</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['installer_group']</tt></td>
    <td>String</td>
    <td>OMC User Group</td>
    <td><tt>omcinstall</tt></td>
  </tr>   
  <tr>
    <td><tt>['omc']['install']</tt></td>
    <td>String</td>
    <td>Install Path for OMC Agent on server</td>
    <td><tt>/omc/install</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['stage']</tt></td>
    <td>String</td>
    <td>Stage path for OMC Agent downloaded on server</td>
    <td><tt>/omc/install</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['app']</tt></td>
    <td>String</td>
    <td>Application install path for OMC Agent on server</td>
    <td><tt>/omc/install</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['apm']</tt></td>
    <td>String</td>
    <td>APM agent install path for on server</td>
    <td><tt>/omc/install</tt></td>
  </tr>
  <tr>
    <td><tt>['omc']['port']</tt></td>
    <td>String</td>
    <td>Custom port to connect to OMC to the agent on</td>
    <td><tt>1830</tt></td>
  </tr>
</table>

## Usage

To create an instance that you can use as the base for a custom image that will have the users
and agentInstall loaded.

### omc::install_omc_agent

```json
{
  "name":"my_node",
  "run_list": [
    "recipe[omc::stage_agent]"
  ]
}
```
To create an instance that is monitored for linux and based on custom image which is pre-loaded with the agent.

### omc::monitor_linux

```json
{
  "name":"my_node",
  "run_list": [
    "recipe[omc::install_agent]",
    "recipe[omc::monitor_linux]"
  ]
}
```
To create an instance Oracle Linux 7 or Centos 7 instance that monitors linux.  Full install and configuration of the 
monitoring agent.

### omc::monitor_linux

```json
{
  "name":"my_node",
  "run_list": [
    "recipe[omc::stage_agent]",
    "recipe[omc::install_agent]",
    "recipe[omc::monitor_linux]"
  ]
}

## Contributing

- Work in progress, email james.calise@oracle.com

## License and Authors

Authors: James Calise

