---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_stages"
sidebar_current: "docs-oci-datasource-os_management_hub-lifecycle_stages"
description: |-
  Provides the list of Lifecycle Stages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_lifecycle_stages
This data source provides the list of Lifecycle Stages in Oracle Cloud Infrastructure Os Management Hub service.

Lists lifecycle stages that match the specified compartment or lifecycle stage OCID. Filter the list against 
a variety of criteria including but not limited to its name, status, architecture, and OS family.


## Example Usage

```hcl
data "oci_os_management_hub_lifecycle_stages" "test_lifecycle_stages" {

	#Optional
	arch_type = var.lifecycle_stage_arch_type
	compartment_id = var.compartment_id
	display_name = var.lifecycle_stage_display_name
	display_name_contains = var.lifecycle_stage_display_name_contains
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id
	os_family = var.lifecycle_stage_os_family
	software_source_id {
	}
	state = var.lifecycle_stage_state
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Optional) A filter to return only profiles that match the given archType.
* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `lifecycle_stage_id` - (Optional) The OCID of the lifecycle stage.
* `os_family` - (Optional) A filter to return only profiles that match the given osFamily.
* `software_source_id` - (Optional) The OCID for the software source.
* `state` - (Optional) A filter to return only lifecycle stage whose lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `lifecycle_stage_collection` - The list of lifecycle_stage_collection.

### LifecycleStage Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the target instances.
* `compartment_id` - The OCID of the tenancy containing the lifecycle stage.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The lifecycle stage OCID that is immutable on creation.
* `lifecycle_environment_id` - The OCID of the lifecycle environment for the lifecycle stage.
* `managed_instance_ids` - The list of managed instances specified lifecycle stage.
	* `display_name` - Managed instance name.
	* `id` - The OCID of the managed instance.
* `os_family` - The operating system type of the target instances.
* `rank` - User specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages for a given lifecycle environment. 
* `software_source_id` - Identifying information for the specified software source.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The OCID of the software source.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the lifecycle stage.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the lifecycle stage was created. An RFC3339 formatted datetime string.
* `time_modified` - The time the lifecycle stage was last modified. An RFC3339 formatted datetime string.
* `vendor_name` - The software source vendor name.

