    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|  
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Writing Terraform configurations for the Oracle Bare Metal Cloud

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

## OBMC resource and datasource details
https://github.com/oracle/terraform-provider-baremetal/tree/master/docs

## Overview
In the simplest terms Terraform turns configurations into a set of API calls against OBMCS API endpoints. The configuration language closely follows but does not mimic the API. Once you understand how to abstract the API into the configuration language writing configuration files is easy.

## Configuration file requirements
Every configuration file defines the provider that will be used, the OBMCS 
provider is called 'baremetal'. You must also specify where to get the 
required authentication details. You should never directly specify these 
values in a configuration file. This syntax will source the values from 
environment variables as covered in the README.  
```
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "baremetal" {
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
 
See the [ad_multi_region](https://github.com/oracle/terraform-provider-baremetal/tree/master/docs/examples/iam/ad_multi_region/ad_multi_region.tf)
or [vcn_multi_region](https://github.com/oracle/terraform-provider-baremetal/tree/master/docs/examples/networking/vcn_multi_region)
examples for details on how to target multiple regions from one plan.

## CamelCase
The OBMCS API uses CamelCase in multiple places. Terraform doesn't support CamelCase in configuration files so we've replaced it with underscores. For example -

	OBMCS native API			Terraform configuration
	----------------			-----------------------
	availabilityDomain			availability_domain
	cidrBlock				cidr_block
	compartmentId			 	compartment_id
	routeTableId			  	route_table_id
	securityListIds		   	    security_list_ids
	vcnId						 vcn_id

## Mutable resources -
Terraform abstracts the create, modify, and delete functions of the API. For example if you use the 'baremetal_core_subnet' resource to create a subnet and later want to change the display name of the subnet all you need to do is change the name in the configuration file. Terraform will determine you are making a change to an existing resource and make the appropriate API calls. If you wanted to delete a subnet you just remove the subnet resource in the configuration file and Terraform will make the appropriate API calls.

This version of the provider supports these resources -
```
	baremetal_core_console_history
    baremetal_core_cpe
    baremetal_core_dhcp_options
    baremetal_core_drg_attachment
    baremetal_core_drg
    baremetal_core_image
    baremetal_core_instance
    baremetal_core_internet_gateway
    baremetal_core_ipsec
    baremetal_core_route_table
    baremetal_core_subnet
    baremetal_core_virtual_network
    baremetal_core_volume_attachment
    baremetal_core_volume_backup
    baremetal_core_volume
    baremetal_identity_api_key
    baremetal_identity_compartment
    baremetal_identity_group
    baremetal_identity_policy
    baremetal_identity_ui_password
    baremetal_identity_user
```

## Non-mutable data sources -
Terraform supports collecting data from multiple sources to influence the application of a configuration. One of those data sources can be an existing Terraform managed OBMCS deployment.  

This version of the provider supports these OBMCS data sources -
```
	baremetal_core_console_history_data
	baremetal_core_cpes
	baremetal_core_dhcp_options
	baremetal_core_drg_attachments
	baremetal_core_drgs
	baremetal_core_images
	baremetal_core_instances
	baremetal_core_internet_gateways
	baremetal_core_ipsec_config
	baremetal_core_ipsec_connections
	baremetal_core_ipsec_status
	baremetal_core_route_tables
	baremetal_core_shape
	baremetal_core_subnets
	baremetal_core_virtual_networks
	baremetal_core_vnic_attachments
	baremetal_core_vnic
	baremetal_core_volume_attachments
	baremetal_core_volume_backups
	baremetal_core_volumes
	baremetal_identity_api_keys
	baremetal_identity_availability_domains
```
