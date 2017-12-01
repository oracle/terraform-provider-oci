    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# The getting started tutorial

## Goal

Using the Oracle Cloud Infrastructure (OCI) Terraform provider, we would like to set up a basic infrastructure, but with all the typical components.

We would build a network from scratch, spin up compute instances, set up some storage etc.

## Intended audience

We hope this tutorial can be used by people who are not familiar with Terraform or the Oracle Cloud Infrastructure (OCI).

This tutorial would provide some hands-on experience in setting up your infrastructure on the OCI platform.

This tutorial may also be used as sample code to start out with when building your own infrastructure on OCI.

## Install Terraform

In order to install Terraform on your machine, follow the straightforward steps as described [here](https://www.terraform.io/intro/getting-started/install.html).

Verify that your installation works as suggested in the linked page above.

Run the following command -

```bash
terraform version
```

Ensure that the Terraform version is `v0.10.x` or higher.

## Install the OCI Terraform provider

The OCI terraform provider is available for download [here](https://github.com/oracle/terraform-provider-oci/releases).

Pick the latest release to download the files.

#### On a Mac
Download the package `darwin.tar.gz`

Unpack the downloaded package using the following commands.
Ensure that you unpack the package in the directory `~/.terraform.d/plugins`

To unpack the tarball, run the following commands on the terminal.

```bash
$ gunzip darwin.tar.gz
$ tar xopf darwin.tar
```

#### On Windows
Download the package - `windows.zip`

Unzip the downloaded file to the directory `%APPDATA%/terraform.d/plugins/`

## Tenancy, Compartment, User setup
If you don't have a tenancy set up with OCI, go [here](https://docs.us-phoenix-1.oraclecloud.com/Content/GSG/Tasks/signingup.htm)

To know more about `compartments`, see [here](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingcompartments.htm)

```markdown
Tip: The tenancy OCID is also the OCID of the root compartment.
```

You would need some [keys](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/credentials.htm) and OCIDs to proceed.

Follow the instructions available [here](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm) to generate your API Signing Keys.

Once the steps described above are complete, you should have a directory named `~/.oci`.

You should have a RSA private-public key pair in the `PEM` format in `~/.oci` directory.

You should have uploaded the public key to your user profile using the console.

Make sure your user account has permissions to create the resources in the chosen `compartment`.

In addition to the API Signing Keys that we obtained, we also need to create an SSH private-public key pair.
This will be useful to SSH into the compute instances that we will spin up as part of this tutorial.

Generate the SSH key using the instructions [here](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm).

## Set up environment variables
Since we will run Terraform on our local machines, we will set up some environment variables that provide the tenancy, compartment, user details that can be used in the Terraform configuration files.

We also need to make the keys, that we generated above, available to Terraform so that OCI Terraform Provider can respond to authentication challenges.
 
```markdown
Note: We recommend that you do not commit files to source control that may have your OCIDs, keys etc.
```

#### On a Mac or other Unix-based OS
Create a file (say) `env_vars`. Save the file in the `~/.oci/` directory that you created earlier.

Add the following statements to your file.

```bash
### Authentication details
export TF_VAR_tenancy_ocid="<tenancy OCID>"
export TF_VAR_user_ocid="<user OCID>"
export TF_VAR_fingerprint="<PEM key fingerprint>"
export TF_VAR_private_key_path="<path to the private key that matches the fingerprint above>"

### Region
export TF_VAR_region="<region in which to operate, example: us-ashburn-1, us-phoenix-1>"

### Compartment
export TF_VAR_compartment_ocid="<compartment OCID>"

### Public/private keys used on the instance
export TF_VAR_ssh_public_key=$(cat <path to public key>)
export TF_VAR_ssh_private_key=$(cat <path to private key>)
```

#### On Windows
Use the following commands on the Windows command prompt to set your environment variables.

```bash
### Authentication details
setx TF_VAR_tenancy_ocid="<tenancy OCID>"
setx TF_VAR_user_ocid="<user OCID>"
setx TF_VAR_fingerprint="<PEM key fingerprint>"
setx TF_VAR_private_key_path="<path to the private key that matches the fingerprint above>"

### Region
setx TF_VAR_region="<region in which to operate, example: us-ashburn-1, us-phoenix-1>"

### Compartment
setx TF_VAR_compartment_ocid="<compartment OCID>"

### Public/private keys used on the instance
setx TF_VAR_ssh_public_key=$(cat <path to public key>)
setx TF_VAR_ssh_private_key=$(cat <path to private key>)
```
#### Make environment variables available
At this point, you should have the `~/.oci/env_vars` file set up.

Now, make the environment variables available to the shell.

On a Mac/Unix, you would typically do something like -

```bash
$ source ~/.oci/env_vars

```

This step makes the Tenancy OCID, keys etc. available to our Terraform Provider.

## Run Terraform - Build Infra
We are now basically done with all the setup work.

It is now time to run Terraform and build some infrastructure elements (network, instances, storage etc.) using it.

#### General instructions on executing the Terraform configurations
With the basic setup out of the way, the following is what we will do.

With Terraform, the infrastructure is defined using a set of configuration files. These configuration files end with a `*.tf` extension.

When the Terraform application is run from the command prompt, every `*.tf` available in the current working directory is picked up, and executed upon.

For this tutorial, all the Terraform configuration files are available under the directory `./examples/get_started/`.

If you are curious, see the next section to see what pieces of the infrastructure we would be deploying.
The `*.tf` also have comments in them that would provide more information on what we are deploying.

At this point we need to run the following command, in the sequence shown below. These commands would deploy our infrastructure elements.

```markdown
-- change to a directory which has the *.tf files
$ cd ./examples/getting_started

-- General Initialization Step. Does some basic syntactic, semantic checking of configuration files.
$ terraform init

-- Generates a plan, as in, lists what will happen if we executed upon our configuration files
$ terraform plan

-- Actually builds the infrastructure that we want.
$ terraform apply
```

The following sections provide information on what infrastructure pieces we will deploy (or, just deployed, if you ran the commands listed above).

#### Network
We will set up a `Virtual Cloud Network` (VCN) in our chosen `region` and `compartment` (defined in the `env_vars` file).

Then, we will set up a few `subnets` that we will be deployed to different `Availability Domains` (ADs) for redundancy.

We will have public `subnets` that will accept traffic from the open internet.

We will set up private `subnets` which will accept from traffic only from (compute) `instances` from within the VCN.

We will set up `Route Tables`, `Security Lists` and an `Internet Gateway` in our network.

The networking related components (`route tables`, `security lists` etc.) are described in detail [here](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/overview.htm)

You will find the Terraform configurations related to the network setup in the file `./examples/getting_started/network.tf`.

#### Storage
`Block storage` basically works like a Hard Disk Drive (HDD) that you can mount on to your compute `instance` to add storage capacity to it.
You can read more about `block storage` [here](https://docs.us-phoenix-1.oraclecloud.com/Content/Block/Concepts/overview.htm)

You will find the Terraform configurations related to the setup of block storage in the `./examples/getting_started/block_storage.tf` file.

Other storage solutions are available with the OCI platform.

`Object Storage` is a high-scale service provided by the OCI platform which is described in detail [here](https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Concepts/overview.htm).
Terraform configuration examples are available in the `./examples/object_storage/` directory.

`Database As A Service` (DBaaS) is also available with the OCI platform.
Terraform configuration examples are available in the `./examples/db_systems/` directory.

#### Compute instances

You will find the Terraform configurations related to the setup of a compute `instance` in the `./examples/getting_started/compute_instances.tf` file.

We set up instances in both public and private subnets that we create.

Instances running in the public subnet would often host internet-facing webservices.

Instances running in the private subnet would often host databases and other backend services.

## Validation
Now, with our infrastructure elements deployed to the Oracle Cloud Infrastructure, the first thing to do would be to appreciate our work.

At this time, open the [Console](https://console.us-phoenix-1.oraclecloud.com/). Login with your credentials.

Navigate around the console to see the various elements we deployed, and if they look good to you.

## Clean up
Once we are done playing with this tutorial, if would be good to clean up what we created.

If we created all the infrastructure elements (network/storage etc.) using Terraform, then cleaning up is easy.

Run the following command.

```markdown
-- change to a directory which has the *.tf files
$ cd ./examples/getting_started

-- Deletes the infrastructure. Good hygiene for a test/tutorial run
$ terraform destroy
```

Verify using the [Console](https://console.us-phoenix-1.oraclecloud.com/) that all the test infrastucture was indeed cleaned up.

```Tip
Compartments, once created, cannot be deleted. Be mindful of that when testing.
```

## Summary and Next Steps
In this tutorial, we learnt the following -

```
    1. What is OCI and Terraform
    2. We got ourselves acquainted with the basics of OCI - tenancy, OCIDs, keys etc.
    3. We installed Terraform
    4. We built a basic infrastructure from scratch
```

For more code examples on how to build the various infrastructure elements, check out the `./examples/` directory.

The `./solutions` has more examples which show how to build solutions - like deploy a webserver using Chef etc.