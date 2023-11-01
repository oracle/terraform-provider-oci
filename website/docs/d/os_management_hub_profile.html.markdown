---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile"
sidebar_current: "docs-oci-datasource-os_management_hub-profile"
description: |-
  Provides details about a specific Profile in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_profile
This data source provides details about a specific Profile resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified registration profile.

## Example Usage

```hcl
data "oci_os_management_hub_profile" "test_profile" {
	#Required
	profile_id = oci_os_management_hub_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `profile_id` - (Required) The OCID of the registration profile.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type.
* `compartment_id` - The OCID of the tenancy containing the registration profile.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the registration profile.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the profile that is immutable on creation.
* `lifecycle_environment` - Identifying information for the specified lifecycle environment.
	* `display_name` - Lifecycle environment name.
	* `id` - The OCID of the lifecycle environment.
* `lifecycle_stage` - Identifying information for the specified lifecycle stage.
	* `display_name` - Lifecycle stage name.
	* `id` - The OCID of the lifecycle stage.
* `managed_instance_group` - Identifying information for the specified managed instance group.
	* `display_name` - Managed instance group displayName.
	* `id` - The OCID of the managed instance group.
* `management_station_id` - The OCID of the management station.
* `os_family` - The operating system family.
* `profile_type` - The type of Profile. One of SOFTWARESOURCE, GROUP or LIFECYCLE.
* `software_sources` - The list of software sources that the registration profile will use.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The OCID of the software source.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the registration profile.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the registration profile was created. An RFC3339 formatted datetime string.
* `vendor_name` - The software source vendor name.

