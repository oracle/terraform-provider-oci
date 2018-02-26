    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
This example creates a VCN with two management subnets, in two different availability domains. It then launches an instance in each of these management subnets and configures them to perform DNS forwarding for DNS hostnames in the VCN, and the DNS hostnames in the on-premises network. See ![Hybrid DNS configuration using DNS VMs in VCN.md](Hybrid-DNS-configuration-using-DNS-VM-in-VCN.md) for more details on the setup. 

To enable resolution of DNS hostnames from on-premises, you will need to update the default DHCP options of the VCN to use the DNS VMs as the DNS resolvers.

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`

Once the environment is built, the DNS VMs will be able to query the DNS hostnames within the VCN. You can run 'nslookup <fqdn-of-an-instance-in-vcn> <DNS VM IP>' from any instance in the VCN to verify this. By specifying an IP address at the end of the 'nslookup' command, the DNS query is sent to the DNS service at that IP address.

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

#### `dns.tf`
Defines the resources. 

