---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_environment"
sidebar_current: "docs-oci-datasource-os_management_hub-lifecycle_environment"
description: |-
  Provides details about a specific Lifecycle Environment in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_lifecycle_environment
This data source provides details about a specific Lifecycle Environment resource in Oracle Cloud Infrastructure Os Management Hub service.

Gets information about the specified lifecycle environment.

## Example Usage

```hcl
data "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
	#Required
	lifecycle_environment_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_environment_id` - (Required) The OCID of the lifecycle environment.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the target instances.
* `compartment_id` - The OCID of the tenancy containing the lifecycle environment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User specified information about the lifecycle environment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource that is immutable on creation.
* `managed_instance_ids` - The list of managed instance OCIDs specified in the lifecycle stage.
	* `display_name` - Managed instance name.
	* `id` - The OCID of the managed instance.
* `os_family` - The operating system type of the target instances.
* `stages` - User specified list of lifecycle stages to be created for the lifecycle environment.
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
* `state` - The current state of the lifecycle environment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the lifecycle environment was created. An RFC3339 formatted datetime string.
* `time_modified` - The time the lifecycle environment was last modified. An RFC3339 formatted datetime string.
* `vendor_name` - The software source vendor name.

