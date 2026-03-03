---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_byol_allocation"
sidebar_current: "docs-oci-datasource-ocvp-byol_allocation"
description: |-
  Provides details about a specific Byol Allocation in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_byol_allocation
This data source provides details about a specific Byol Allocation resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Gets the specified BYOL Allocation's information.

## Example Usage

```hcl
data "oci_ocvp_byol_allocation" "test_byol_allocation" {
	#Required
	byol_allocation_id = oci_ocvp_byol_allocation.test_byol_allocation.id
}
```

## Argument Reference

The following arguments are supported:

* `byol_allocation_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation.


## Attributes Reference

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

