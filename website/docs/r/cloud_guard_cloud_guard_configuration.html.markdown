---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_cloud_guard_configuration"
sidebar_current: "docs-oci-resource-cloud_guard-cloud_guard_configuration"
description: |-
  Provides the Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_cloud_guard_configuration
This resource provides the Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service.

Updates configuration details for a Cloud Guard tenancy, identified by root compartment OCID.
The reporting region cannot be updated once created.


## Example Usage

```hcl
resource "oci_cloud_guard_cloud_guard_configuration" "test_cloud_guard_configuration" {
	#Required
	compartment_id = var.compartment_id
	reporting_region = var.cloud_guard_configuration_reporting_region
	status = var.cloud_guard_configuration_status

	#Optional
	self_manage_resources = var.cloud_guard_configuration_self_manage_resources
	service_configurations {
		#Required
		service_configuration_type = var.cloud_guard_configuration_service_configurations_service_configuration_type

		#Optional
		status = var.cloud_guard_configuration_service_configurations_status
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to list resources.
* `reporting_region` - (Required) (Updatable) The reporting region
* `self_manage_resources` - (Optional) (Updatable) Identifies if Oracle managed resources will be created by customers. If no value is specified false is the default. 
* `service_configurations` - (Optional) (Updatable) List of service configurations for tenant
	* `service_configuration_type` - (Required) (Updatable) Type of service configuration
	* `status` - (Optional) (Updatable) Partner service status
* `status` - (Required) (Updatable) Status of Cloud Guard tenant


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `reporting_region` - The reporting region
* `self_manage_resources` - Were Oracle-managed resources created by customer? 
* `service_configurations` - List of service configurations for this tenant
	* `service_configuration_type` - Type of service configuration
	* `status` - Partner service status
* `status` - Status of the Cloud Guard tenant

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Guard Configuration
	* `update` - (Defaults to 20 minutes), when updating the Cloud Guard Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Guard Configuration


## Import

Import is not supported for this resource.

