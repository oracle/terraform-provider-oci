---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_byol_allocations"
sidebar_current: "docs-oci-datasource-ocvp-byol_allocations"
description: |-
  Provides the list of Byol Allocations in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_byol_allocations
This data source provides the list of Byol Allocations in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists the BYOL Allocations in the specified compartment. The list can be
filtered by display name or availability domain.


## Example Usage

```hcl
data "oci_ocvp_byol_allocations" "test_byol_allocations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	available_units_greater_than_or_equal_to = var.byol_allocation_available_units_greater_than_or_equal_to
	byol_allocation_id = oci_ocvp_byol_allocation.test_byol_allocation.id
	byol_id = oci_ocvp_byol.test_byol.id
	display_name = var.byol_allocation_display_name
	software_type = var.byol_allocation_software_type
	state = var.byol_allocation_state
}
```

## Argument Reference

The following arguments are supported:

* `available_units_greater_than_or_equal_to` - (Optional) A filter to return only resources whose availableUnits greater than or equal to the given value.
* `byol_allocation_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation.
* `byol_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `software_type` - (Optional) A filter to return only resources whose softwareType matches the given value.
* `state` - (Optional) A filter to return only resources whose lifecycle state matches the given value.


## Attributes Reference

The following attributes are exported:

* `byol_allocation_collection` - The list of byol_allocation_collection.

### ByolAllocation Reference

The following attributes are exported:

* `allocated_units` - The quantity of licensed units that allocated to this region. 
* `available_units` - The quantity of licensed units that not yet consumed by resources. 
* `byol_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL resource from which this BYOL Allocation is derived. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the BYOL Allocation. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the BYOL Allocation. 
* `entitlement_key` - The Broadcom-supplied identifier of a BYOL license. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation. 
* `software_type` - The type of VMware software the BYOL applies to.  Supported values:
	* VCF (VMware Cloud Foundation)
	* VSAN (VMware vSAN)
	* VDEFEND (VMware vDefend Firewall)
	* AVI_LOAD_BALANCER (VMware Avi Load Balancer) 
* `state` - The current state of the BYOL Allocation.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time the BYOL Allocation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_term_end` - The date and time when the BYOL Allocation expires and becomes inactive. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_term_start` - The date and time when the BYOL Allocation becomes active. VMware software functionality cannot begin before this time. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the BYOL Allocation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

