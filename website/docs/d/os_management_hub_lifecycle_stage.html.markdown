---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_lifecycle_stage"
sidebar_current: "docs-oci-datasource-os_management_hub-lifecycle_stage"
description: |-
  Provides details about a specific Lifecycle Stage in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_lifecycle_stage
This data source provides details about a specific Lifecycle Stage resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns information about the specified lifecycle stage.


## Example Usage

```hcl
data "oci_os_management_hub_lifecycle_stage" "test_lifecycle_stage" {
	#Required
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_stage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.


## Attributes Reference

The following attributes are exported:

* `arch_type` - The CPU architecture of the managed instances in the lifecycle stage.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the lifecycle stage.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the lifecycle stage.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage.
* `lifecycle_environment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment that contains the lifecycle stage.
* `location` - The location of managed instances associated with the lifecycle stage.
* `managed_instance_ids` - The list of managed instances associated with the lifecycle stage.
	* `display_name` - Managed instance name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `os_family` - The operating system of the managed instances in the lifecycle stage.
* `rank` - User-specified rank for the lifecycle stage. Rank determines the hierarchy of the lifecycle stages within the lifecycle environment. 
* `software_source_id` - Provides identifying information for the specified software source.
	* `description` - Software source description.
	* `display_name` - Software source name.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
	* `software_source_type` - Type of the software source.
* `state` - The current state of the lifecycle stage.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the lifecycle stage was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_modified` - The time the lifecycle stage was last modified (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `vendor_name` - The vendor of the operating system used by the managed instances in the lifecycle stage.

