    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
This example creates a VCN with two management subnets, in two different availability domains. It then launches an instance in each of these management subnets and configures them to perform DNS forwarding for DNS hostnames in the VCN, and the DNS hostnames in the on-premises network. See document for more details on the setup. 

Once you apply the configuration, you will need to update the default DHCP options of the VCN to use the DNS VMs as the DNS resolvers. The script configure_default_dhcp.sh retrieves the information from 'terraform output' and constructs the required files to make this change. 

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`


#### `dns.tf`
Defines the resources. 

### configure_default_dhcp.sh
Once you apply, you need to run this script to update the default DHCP options to use the DNS servers launched with this template. 

