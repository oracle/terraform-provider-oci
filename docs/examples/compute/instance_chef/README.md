    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
## Bootstrap an instance as a Chef node
This example shows how to use Terraform to bootstrap an OCI instance as a Chef node, register it with an existing Chef server, and run a recipe. When the plan completes, you will have a running web server with content and the correct firewall configuration.

### Prerequisites

* `knife` must be installed and configured in your PATH. This allows the Chef node deletion to work properly when `terraform destroy` is called.
* Access to a Chef server to upload the example_webserver recipe to.

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`
* Update `variables.tf` with your instance options.
* Upload the example Chef recipe and its dependencies to your Chef server.
  * `$ cd cookbooks/example_webserver`
  * `$ berks install`
  * `$ berks upload`
* Run `terraform apply` from the instance_chef directory.
* Navigate to the public IP address (reported by Terraform, above) in your browser.

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -
`$ . env-vars`

#### `compute.tf`
Defines the compute resource. This demo connects to the running instance
so you will need to supply public/private keys to create an ssh connection.
**NOTE**: do not try to use your api keys, see [this doc](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm)
for more info on configuring keys.

#### `./userdata/bootstrap`
The user-data script that gets injected into the instance on launch. More information on user-data scripts can be [found at the cloud-init project.](https://cloudinit.readthedocs.io/en/latest/topics/format.html)

#### `variables.tf`
Defines the variables used in the configuration

#### `datasources.tf`
Defines the datasources used in the configuration

#### `outputs.tf`
Defines the outputs of the configuration

#### `provider.tf`
Specifies and passes authentication details to the OCI TF provider

#### `./cookbooks/example_webserver/recipes/default.rb`
Installs a web server and configures the firewall to allow inbound http and ssh connections.
