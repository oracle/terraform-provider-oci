    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|  
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Writing Terraform configurations for the Oracle Cloud Infrastructure

### All about Terraform
The Terraform book, https://terraformbook.com/

Scott Lowe writes a great blog, including this introduction to Terraform http://blog.scottlowe.org/2015/11/25/intro-to-terraform/


#### Read the Terraform configuration documentation
Overview - https://www.terraform.io/docs/configuration/index.html  
Configuration file syntax - https://www.terraform.io/docs/configuration/syntax.html  
Resources - https://www.terraform.io/docs/configuration/resources.html  
Outputs - https://www.terraform.io/docs/configuration/outputs.html  
Logging - https://www.terraform.io/docs/configuration/environment-variables.html  
Interpolation - https://www.terraform.io/docs/configuration/interpolation.html  

## OCI resource and datasource details
https://github.com/oracle/terraform-provider-oci/tree/master/docs

## Overview
In the simplest terms Terraform turns configurations into a set of API calls against OCI API endpoints. The configuration language closely follows but does not mimic the API. Once you understand how to abstract the API into the configuration language writing configuration files is easy.

## Configuration file requirements
Every configuration file defines the provider that will be used, the OCI
provider is called 'oci'. You must also specify where to get the 
required authentication details. You should never directly specify these 
values in a configuration file. This syntax will source the values from 
environment variables as covered in the README.  
```
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}
```

Use the region parameter in your provider definition to specify which region 
your resources will be created in. Currently, not specifying a region value 
will result in use of the `us-phoenix-1` region, **but this is discouraged 
as region will soon become a required parameter.**
 
See the [ad_multi_region](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/iam/ad_multi_region/ad_multi_region.tf)
or [vcn_multi_region](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/networking/vcn_multi_region)
examples for details on how to target multiple regions from one plan.

## CamelCase
The OCI API uses CamelCase in multiple places. Terraform doesn't support CamelCase in configuration files so we've replaced it with underscores. For example -

	OCI native API			Terraform configuration
	----------------			-----------------------
	availabilityDomain			availability_domain
	cidrBlock				cidr_block
	compartmentId			 	compartment_id
	routeTableId			  	route_table_id
	securityListIds		   	    security_list_ids
	vcnId						 vcn_id

## Mutable resources -
Terraform abstracts the create, modify, and delete functions of the API. For example if you use the 'oci_core_subnet' resource to create a subnet and later want to change the display name of the subnet all you need to do is change the name in the configuration file. Terraform will determine you are making a change to an existing resource and make the appropriate API calls. If you wanted to delete a subnet you just remove the subnet resource in the configuration file and Terraform will make the appropriate API calls.

This version of the provider supports these resources -
```
	oci_core_console_history
    oci_core_cpe
    oci_core_dhcp_options
    oci_core_drg_attachment
    oci_core_drg
    oci_core_image
    oci_core_instance
    oci_core_internet_gateway
    oci_core_ipsec
    oci_core_route_table
    oci_core_subnet
    oci_core_virtual_network
    oci_core_volume_attachment
    oci_core_volume_backup
    oci_core_volume
    oci_identity_api_key
    oci_identity_compartment
    oci_identity_group
    oci_identity_policy
    oci_identity_ui_password
    oci_identity_user
```

## Non-mutable data sources -
Terraform supports collecting data from multiple sources to influence the application of a configuration. One of those data sources can be an existing Terraform managed OCI deployment.

This version of the provider supports these OCI data sources -
```
	oci_core_console_history_data
	oci_core_cpes
	oci_core_dhcp_options
	oci_core_drg_attachments
	oci_core_drgs
	oci_core_images
	oci_core_instances
	oci_core_internet_gateways
	oci_core_ipsec_config
	oci_core_ipsec_connections
	oci_core_ipsec_status
	oci_core_route_tables
	oci_core_shape
	oci_core_subnets
	oci_core_virtual_networks
	oci_core_vnic_attachments
	oci_core_vnic
	oci_core_volume_attachments
	oci_core_volume_backups
	oci_core_volumes
	oci_identity_api_keys
	oci_identity_availability_domains
```
