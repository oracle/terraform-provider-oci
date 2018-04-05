    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
## Deploy the File Storage Service resources
This example creates file systems, mount targets, exports the file systems, and creates a snapshot.

One of the file systems is exported using two different paths on two different mount targets.

We see that a single mount target can export paths from two (or more) file systems.

We also see how we need to specify certain stateful ingress rules in a security list for the file system to be operational.

Finally, we set up a compute instance that mounts the NFS storage.

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`
* Update `variables.tf` with your instance options.

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

#### `file_system.tf`
Defines two file system resources

#### `mount_target.tf`
Defines two mount target resources

#### `export.tf`
Defines the exports - used to make the file systems accessible via the mount targets

#### `snapshot.tf`
Defines a snapshot for a file system

#### `network.tf`
Defines a virtual cloud network, internet gateway, route table and a subnet.

This basic setup is needed to enable SSH to our instance.

#### `security_list.tf`
Defines a security list setup to make our file system operational.

We keep this separate from our `network.tf` file since there are some interesting things to note here.

In particular, some specific ports are being opened to allow the NFS communication to happen.

### `instance.tf`
Defines our compute instance.

Note the remote action that we execute on our compute instance once it is launched.

We install the nfs-utils and then mount the NFS storage on to our compute instance.

#### `variables.tf`
Defines the variables used in the configuration

#### `datasources.tf`
Defines the data sources used in the configuration

#### `provider.tf`
Specifies and passes authentication details to the OCI TF provider
