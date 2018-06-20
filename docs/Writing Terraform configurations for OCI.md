    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|  
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Writing Terraform configurations for Oracle Cloud Infrastructure

### All about Terraform
- The Terraform book: https://terraformbook.com/
- Scott Lowe writes a great blog, including [this introduction to Terraform](http://blog.scottlowe.org/2015/11/25/intro-to-terraform/).

#### Read the Terraform configuration documentation
- Overview - https://www.terraform.io/docs/configuration/index.html
- Configuration file syntax - https://www.terraform.io/docs/configuration/syntax.html
- Resources - https://www.terraform.io/docs/configuration/resources.html
- Outputs - https://www.terraform.io/docs/configuration/outputs.html
- Logging - https://www.terraform.io/docs/configuration/environment-variables.html
- Interpolation - https://www.terraform.io/docs/configuration/interpolation.html

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

Use the 'region' attribute in your provider definition to specify which region 
your resources will be created in. See the [ad_multi_region](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/iam/ad_multi_region/ad_multi_region.tf)
or [vcn_multi_region](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/networking/vcn_multi_region)
examples for details on how to target multiple regions from one plan.

### Enabling Instance Principal Authorization
To enable instance principal authorization, you can set 'auth' attribute to "InstancePrincipal"
in the provider definition as follows ('tenancy_ocid', 'user_ocid', 'fingerprint'
and 'private_key_path' are not necessary):
```
variable "region" {}

provider "oci" {
  auth = "InstancePrincipal"
  region = "${var.region}"
}
```

See [Calling Services from an instance](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/callingservicesfrominstances.htm)
for setting up and using instances as principals.

## OCI resource and data source details
A list of all supported OCI resources and data sources can be found in the [Table of Contents](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Table%20of%20Contents.md).

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

## Filtering
Most OCI data sources support filtering - see [docs/Filters.md](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Filters.md) for details.

## Tagging
See [docs/Tagging Resources.md](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Tagging%20Resources.md) for how to manage tags and list of OCI resources that support tagging.