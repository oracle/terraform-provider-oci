[terraform]: https://terraform.io
[oci]: https://cloud.oracle.com/cloud-infrastructure
[oci provider]: https://github.com/oracle/terraform-provider-oci/releases
[API signing]: https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm

# Terraform DC/OS Installer for Oracle Cloud Infrastructure
Version: 0.1
Created by: Changbin Gong
Date: 03/12/2018

## About
The DC/OS Installer for Oracle Cloud Infrastructure provides a Terraform-based DC/OS installation for Oracle Cloud Infrastructure. It consists of a set of [Terraform][terraform] modules and an example base configuration that is used to provision and configure the resources needed to run a highly available and configurable DC/OS cluster on [Oracle Cloud Infrastructure][oci](OCI).

This installer is based on [packet-terraform](https://github.com/mesosphere/packet-terraform). 

### DC/OS Cluster Overview

Terraform is used to _provision_ the cloud infrastructure and any required local resources for the DC/OS cluster including:

#### OCI Infrastructure

- Virtual Cloud Network (VCN) with dedicated public subnets for bootstrap, DC/OS masters, public and private agents in each availability domain
- Dedicated compute instances for bootstrap, DC/OS master and agent nodes in corresponding availability domain based upon configuration

#### Cluster Configuration

Terraform uses remote provisioner to handle instance-level _configuration_ including:
- Bootstrap node
- Master node(s)
- Public Agent node(s)
- Agent node(s)
- GPU Agent node(s)

The Terraform scripts accept a number of other input variables to choose instance shapes (including GPU) and how they are placed across the availability domain (ADs), etc. If your requirements extend beyond the base configuration, the modules can be used to form your own customized configuration.

## Prerequisites

1. Download and install [Terraform][terraform] (v0.10.3 or later)
2. Download and install the [OCI Terraform Provider][oci provider] (v2.0.0 or later)
3. Create an Terraform configuration file at  `~/.terraformrc` that specifies the path to the OCI provider:
```
providers {
  oci = "<path_to_provider_binary>/terraform-provider-oci"
}
```

## Quick start

### Customize the configuration

Create a _env-vars_ file in the project root that specifies your configuration. Please look at "vars.tf" file for more information.

* Set mandatory OCI input variables related to your tenancy, user and compartment including:
   - Tenancy ocid
   - User ocid
   - API fingeriprint  
   - API private key
   - Compartment ocid
   - Public and Private Key pairs for SSH access

* Source env-vars
  * `$ . env-vars`

### Deploy the cluster

Initialize Terraform:

```
$ terraform init
```

View what Terraform plans do before actually doing it:

```
$ terraform plan
```

Use Terraform to Provision resources and stand-up DC/OS cluster on OCI:

```
$ terraform apply
```

### Access the cluster

The DC/OS cluster will be running after the configuration is applied successfully. Typically, this takes around 5 minutes after `terraform apply` and will vary depending on the overall configuration, instance counts, and shapes.

When provision completes, the Terraform scripts show following outputs:

Outputs:

master_private_ips = [
    x.x.x.x,
    x.x.x.x,
    x.x.x.x,
]
master_public_ips = [
    x.x.x.x,
    x.x.x.x,
    x.x.x.x
]

Then you can access DC/OS dashboard via "http://<master_public_ip>/"

### Scale, upgrade, or delete the cluster

After DC/OS cluster is provisioned, you can scale up or down the number of Agent or Public Agent nodes by using input variables defined in "vars.tf".  For instance, you can change tne number of "private agent" node as following:

```
$ terraform apply -var "dcos_agent_id1_count=xx"
```

If you do not need the DC/OS cluster,  you can use "terraform destroy" to delete the whole DC/OS deployment.  

```
$ terraform destroy
```

## Known issues and limitations

* DC/OS cluster is only deployed on public subnets
* After DC/OS deployment, you can not change the number of master nodes.
* No OCI Load Balancers support in this version. 

