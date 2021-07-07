---
subcategory: "Content and Experience"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oce_oce_instances"
sidebar_current: "docs-oci-datasource-oce-oce_instances"
description: |-
  Provides the list of Oce Instances in Oracle Cloud Infrastructure Content and Experience service
---

# Data Source: oci_oce_oce_instances
This data source provides the list of Oce Instances in Oracle Cloud Infrastructure Content and Experience service.

Returns a list of OceInstances.


## Example Usage

```hcl
data "oci_oce_oce_instances" "test_oce_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oce_instance_display_name
	state = var.oce_instance_state
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) Filter results on lifecycleState.
* `tenancy_id` - (Optional) The ID of the tenancy in which to list resources.


## Attributes Reference

The following attributes are exported:

* `oce_instances` - The list of oce_instances.

### OceInstance Reference

The following attributes are exported:

* `admin_email` - Admin Email for Notification
* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - OceInstance description, can be updated
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `guid` - Unique GUID identifier that is immutable on creation
* `id` - Unique identifier that is immutable on creation
* `idcs_tenancy` - IDCS Tenancy Identifier
* `instance_access_type` - Flag indicating whether the instance access is private or public
* `instance_license_type` - Flag indicating whether the instance license is new cloud or bring your own license
* `instance_usage_type` - Instance type based on its usage
* `name` - OceInstance Name
* `object_storage_namespace` - Object Storage Namespace of tenancy
* `service` - SERVICE data. Example: `{"service": {"IDCS": "value"}}` 
* `state` - The current state of the file system.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - Tenancy Identifier
* `tenancy_name` - Tenancy Name
* `time_created` - The time the the OceInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the OceInstance was updated. An RFC3339 formatted datetime string
* `upgrade_schedule` - Upgrade schedule type representing service to be upgraded immediately whenever latest version is released or delay upgrade of the service to previous released version 
* `waf_primary_domain` - Web Application Firewall(WAF) primary domain

