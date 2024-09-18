---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_resource_principal_configuration"
sidebar_current: "docs-oci-resource-bds-bds_instance_resource_principal_configuration"
description: |-
  Provides the Bds Instance Resource Principal Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_resource_principal_configuration
This resource provides the Bds Instance Resource Principal Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Create a resource principal session token configuration.


## Example Usage

```hcl
resource "oci_bds_bds_instance_resource_principal_configuration" "test_bds_instance_resource_principal_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_resource_principal_configuration_cluster_admin_password
	display_name = var.bds_instance_resource_principal_configuration_display_name

	#Optional
	session_token_life_span_duration_in_hours = var.bds_instance_resource_principal_configuration_session_token_life_span_duration_in_hours
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) Base-64 encoded Cluster Admin Password for cluster admin user.
* `display_name` - (Required) (Updatable) A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `session_token_life_span_duration_in_hours` - (Optional) (Updatable) Life span in hours for the resource principal session token.
* `force_refresh_resource_principal_trigger` - (Optional) (Updatable) An optional property when incremented triggers Force Refresh Resource Principal. Could be set to any integer value.
* `remove_trigger` - (Optional) (Updatable) An optional property when incremented triggers Remove. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `display_name` - A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `id` - The id of the ResourcePrincipalConfiguration.
* `session_token_life_span_duration_in_hours` - Life span in hours of each resource principal session token.
* `state` - The state of the ResourcePrincipalConfiguration.
* `time_created` - The time the ResourcePrincipalConfiguration was created, shown as an RFC 3339 formatted datetime string.
* `time_token_expiry` - the time the resource principal session token will expired, shown as an rfc 3339 formatted datetime string.
* `time_token_refreshed` - the time the resource principal session token was refreshed, shown as an rfc 3339 formatted datetime string.
* `time_updated` - The time the ResourcePrincipalConfiguration was updated, shown as an RFC 3339 formatted datetime string. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Resource Principal Configuration
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Resource Principal Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Resource Principal Configuration


## Import

BdsInstanceResourcePrincipalConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration "bdsInstances/{bdsInstanceId}/resourcePrincipalConfigurations/{resourcePrincipalConfigurationId}" 
```

