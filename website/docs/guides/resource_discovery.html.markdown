---
layout: "oci"
page_title: "Resource Discovery"
sidebar_current: "docs-oci-guide-resource_discovery"
description: |-
  The Oracle Cloud Infrastructure provider. Discovering resources in an existing compartment
---

## Discovering Terraform resources in an Oracle Cloud Infrastructure compartment

### Overview
Beginning with version 3.50, the terraform-oci-provider can be run as a command line tool to discover resources that have been created within Oracle Cloud Infrastructure compartments and generate Terraform configuration files for the discovered resources.

The latest version of the terraform-oci-provider can be downloaded using `terraform init` or by going to https://releases.hashicorp.com/terraform-provider-oci/

### Authentication
To discover resources in your compartment, the terraform-oci-provider will need authentication information about the user, tenancy, and region with which to discover
the resources. It is recommended to specify a user that has access to inspect and read the resources to discover. 

Resource discovery supports API Key based authentication and Instance Principal based authentication.

The authentication information can be specified using the following environment variables:

```
export TF_VAR_tenancy_ocid=<value>
export TF_VAR_user_ocid=<value>
export TF_VAR_fingerprint=<value>
export TF_VAR_private_key_path=<path to your private key>
export TF_VAR_region=<region of the resources, e.g. "us-phoenix-1">
```


If your private key is password-encrypted, you may also need to specify a password with this variable:

```
export TF_VAR_private_key_password=<password for private key>
```

The authentication information can also be specified using a configuration file. For details on setting this up, see [SDK and CLI configuration file](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/sdkconfig.htm)
A non-default profile can be set using environment variable: 

```
export TF_VAR_config_file_profile=<value>
``` 


If the parameters have multiple sources, the priority will be in the following order: 

    Environment variables
    Non-default profile
    DEFAULT profile


### Usage

Once you have specified the prerequisite authentication settings, the command can be used as follows with a compartment being specified by name or OCID:

```
terraform-provider-oci -command=export -compartment_name=<name of compartment to export> -output_path=<directory under which to generate Terraform files>
```


```
terraform-provider-oci -command=export -compartment_id=<OCID of compartment to export> -output_path=<directory under which to generate Terraform files>
```

This command will discover resources within your compartment and generates Terraform configuration files in the given `output_path`. 
The generated `.tf` files contain the Terraform configuration with the resources that the command has discovered.

**Parameter Description**

* `command` - Command to run. Supported commands include: 
    * `export` - Discovers Oracle Cloud Infrastructure resources within your compartment and generates Terraform configuration files for them
    * `list_export_resources` - Lists the Terraform Oracle Cloud Infrastructure resources types that can be discovered by the `export` command
* `compartment_id` - OCID of a compartment to export. If `compartment_id`  or `compartment_name` is not specified, the root compartment will be used.
* `compartment_name` - The name of a compartment to export. Use this instead of `compartment_id` to provide a compartment name.
* `ids` - Comma-separated list of resource IDs to export. The ID could either be an OCID or a Terraform import ID. By default, all resources are exported.
* `output_path` - Path to output generated configurations and state files of the exported compartment
* `services` - Comma-separated list of service resources to export. If not specified, all resources within the given compartment (which excludes identity resources) are exported. The following values can be specified:
    * `availability_domain` - Discovers availability domains used by your compartment-level resources. It is recommended to always specify this value.
    * `bds` - Discovers big data service resources within the specified compartment
    * `core` - Discovers compute, block storage, and networking resources within the specified compartment
    * `database` - Discovers database and autonomous database resources within the specified compartment
    * `identity` - Discovers identity resources across the entire tenancy
    * `load_balancer` - Discovers load balancer resources within the specified compartment
    * `object_storage` - Discovers object storage resources within the specified compartment
    * `tagging` - Discovers tag-related resources within the specified compartment
* `generate_state` - Provide this flag to import the discovered resources into a state file along with the Terraform configuration

> **Note**: The compartment export functionality currently supports discovery of the target compartment. The ability to discover resources in child compartments is not yet supported.  

### Generated Terraform Configuration Contents

The command will discover resources that are in an active or usable state. Resources that have been terminated or otherwise made inactive are generally excluded from the generated configuration.

By default, the Terraform names of the discovered resources will share the same name as the display name for that resource, if one exists.

The attributes of the resources will be populated with the values that are returned by the Oracle Cloud Infrastructure services.

In some cases, a required or optional attribute may not be discoverable from the Oracle Cloud Infrastructure services and may be omitted from the generated Terraform configuration.
This may be expected behavior from the service, which may prevent discovery of certain sensitive attributes or secrets. In such cases, the generated Terraform configuration will contain a commented line like this:

```
#admin_password = <<Required attribute not found in discovery>>
```

Run 'terraform plan' against the generated configuration files to get more information about the missing values.

### Exporting Identity Resources

Some resources, such as identity resources, may exist only at the tenancy level and cannot be discovered within a specific compartment. To discover such resources, specify
the following command.

```
terraform-provider-oci -command=export -output_path=<directory under which to generate Terraform files> -services=identity
``` 

> **Note**: When exporting identity resources, a `compartment_id` is not required. If a `compartment_id` is specified, the value will be ignored for discovering identity resources.


### Exporting Resources to Another Compartment
Once the user has reviewed the generated configuration and made the necessary changes to reflect the desired settings, the configuration can be used with Terraform.
One such use case is the re-deploying of those resources in a new compartment or tenancy, using Terraform.
 
To do so, specify the following environment variables:

```
export TF_VAR_tenancy_ocid=<new tenancy OCID>
export TF_VAR_compartment_ocid=<new compartment OCID>
```

And run 

```
terraform apply
```

### Generating a Terraform State File

Using this command it is also possible to generate a Terraform state file to manage the discovered resources. To do so, run the following command:

```
terraform-provider-oci -command=export -compartment_id=<compartment to export> -output_path=<directory under which to generate Terraform files> -generate_state
```

The results of this command are both the `.tf` files representing the Terraform configuration and a `terraform.tfstate` file representing the state.

> **Note** The Terraform state file generated by this command is currently compatible with Terraform v0.12.4 and above


### Supported Resources
As of this writing, the list of Terraform services and resources that can be discovered by the command is as follows.
The list of supported resources can also be retrieved by running this command:

```
terraform-provider-oci -command=list_export_resources
```

identity (tenancy-scope resources)

* oci\_identity\_api\_key
* oci\_identity\_auth\_token
* oci\_identity\_authentication\_policy
* oci\_identity\_compartment
* oci\_identity\_customer\_secret\_key
* oci\_identity\_dynamic\_group
* oci\_identity\_group
* oci\_identity\_identity\_provider
* oci\_identity\_idp\_group\_mapping
* oci\_identity\_policy
* oci\_identity\_smtp\_credential
* oci\_identity\_ui\_password
* oci\_identity\_user
* oci\_identity\_user\_group\_membership

core (compartment-scope resources)

* oci\_core\_boot\_volume
* oci\_core\_cpe
* oci\_core\_cross\_connect
* oci\_core\_cross\_connect\_group
* oci\_core\_dhcp\_options
* oci\_core\_drg
* oci\_core\_drg\_attachment
* oci\_core\_image
* oci\_core\_instance
* oci\_core\_instance\_configuration
* oci\_core\_instance\_pool
* oci\_core\_internet\_gateway
* oci\_core\_ipsec
* oci\_core\_local\_peering\_gateway
* oci\_core\_nat\_gateway
* oci\_core\_network\_security\_group
* oci\_core\_network\_security\_group\_security\_rule
* oci\_core\_remote\_peering\_connection
* oci\_core\_route\_table
* oci\_core\_security\_list
* oci\_core\_service\_gateway
* oci\_core\_subnet
* oci\_core\_vcn
* oci\_core\_virtual\_circuit
* oci\_core\_vnic\_attachment
* oci\_core\_volume
* oci\_core\_volume\_attachment
* oci\_core\_volume\_backup\_policy\_assignment
* oci\_core\_volume\_group

database (compartment-scope resources)

* oci\_database\_autonomous\_container\_database
* oci\_database\_autonomous\_database
* oci\_database\_autonomous\_exadata\_infrastructure
* oci\_database\_db\_home
* oci\_database\_db\_system

load\_balancer (compartment-scope resources)

* oci\_load\_balancer\_backend
* oci\_load\_balancer\_backend\_set
* oci\_load\_balancer\_certificate
* oci\_load\_balancer\_hostname
* oci\_load\_balancer\_listener
* oci\_load\_balancer\_load\_balancer
* oci\_load\_balancer\_path\_route\_set
* oci\_load\_balancer\_rule\_set

object_storage (compartment-scope resources)

* oci\_objectstorage\_bucket

tagging (compartment-scope resources)

* oci\_identity\_tag
* oci\_identity\_tag\_default
* oci\_identity\_tag\_namespace
