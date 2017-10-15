    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
This example creates a VCN with a public subnet and a private subnet. Each subnet is created with a separate security list and route table. The template then launches a private instance in the private subnet, and a public instance in the public subnet. 
The public instance is configured as a NAT instance (by enabling forwarding and configuring firewall to do forwarding/masquerading).
The private subnet's route table is configured to use the NAT instance's private IP address as the default route target. See [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip) for more details on this feature.

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`
* Update variables in vcn.tf as applicable to your target environment.

Once the environment is built, the private instance has Internet connectivity even when it doesn't have a public IP address and it's subnet's route table doesn't contain Internet gateway. You can login into the private instance (from the nat instance) and then run a command like 'ping oracle.com' to verify connectivity.

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

#### `nat.tf`
Defines the resources. 

