---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_resource_principal_configurations"
sidebar_current: "docs-oci-datasource-bds-bds_instance_resource_principal_configurations"
description: |-
  Provides the list of Bds Instance Resource Principal Configurations in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_resource_principal_configurations
This data source provides the list of Bds Instance Resource Principal Configurations in Oracle Cloud Infrastructure Big Data Service service.

Returns information about the ResourcePrincipalConfiguration.


## Example Usage

```hcl
data "oci_bds_bds_instance_resource_principal_configurations" "test_bds_instance_resource_principal_configurations" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	display_name = var.bds_instance_resource_principal_configuration_display_name
	state = var.bds_instance_resource_principal_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The state of the ResourcePrincipalConfiguration.


## Attributes Reference

The following attributes are exported:

* `resource_principal_configurations` - The list of resource_principal_configurations.

### BdsInstanceResourcePrincipalConfiguration Reference

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

