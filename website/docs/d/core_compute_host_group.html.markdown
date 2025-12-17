---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_host_group"
sidebar_current: "docs-oci-datasource-core-compute_host_group"
description: |-
  Provides details about a specific Compute Host Group in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_host_group
This data source provides details about a specific Compute Host Group resource in Oracle Cloud Infrastructure Core service.

Gets information about the specified compute host group

## Example Usage

```hcl
data "oci_core_compute_host_group" "test_compute_host_group" {
	#Required
	compute_host_group_id = oci_core_compute_host_group.test_compute_host_group.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_host_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute host group. 


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of a host group.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains host group. 
* `configurations` - A list of HostGroupConfiguration objects
	* `firmware_bundle_id` - The OCID for firmware bundle
	* `recycle_level` - Preferred recycle level for hosts associated with the reservation config.
		* `SKIP_RECYCLE` - Skips host wipe.
		* `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior. 
	* `state` - The state of the host group configuration.
	* `target` - Either the platform name or compute shape that the configuration is targeting
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique host group 
* `is_targeted_placement_required` - A flag that allows customers to restrict placement for hosts attached to the group. If true, the only way to place on hosts is to target the specific host group.
* `state` - The lifecycle state of the host group 
* `system_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `time_created` - The date and time the host group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the host group was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

