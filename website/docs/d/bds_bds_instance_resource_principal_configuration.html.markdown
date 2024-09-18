---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_resource_principal_configuration"
sidebar_current: "docs-oci-datasource-bds-bds_instance_resource_principal_configuration"
description: |-
  Provides details about a specific Bds Instance Resource Principal Configuration in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_resource_principal_configuration
This data source provides details about a specific Bds Instance Resource Principal Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Returns details of the resourcePrincipalConfiguration identified by the given ID.


## Example Usage

```hcl
data "oci_bds_bds_instance_resource_principal_configuration" "test_bds_instance_resource_principal_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	resource_principal_configuration_id = oci_audit_configuration.test_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `resource_principal_configuration_id` - (Required) Unique Oracle-assigned identifier of the ResourcePrincipalConfiguration.


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

