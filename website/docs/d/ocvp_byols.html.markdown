---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_byols"
sidebar_current: "docs-oci-datasource-ocvp-byols"
description: |-
  Provides the list of Byols in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_byols
This data source provides the list of Byols in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists the BYOLs in the specified compartment. The list can be
filtered by display name or availability domain.


## Example Usage

```hcl
data "oci_ocvp_byols" "test_byols" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	available_units_greater_than_or_equal_to = var.byol_available_units_greater_than_or_equal_to
	byol_id = oci_ocvp_byol.test_byol.id
	display_name = var.byol_display_name
	software_type = var.byol_software_type
	state = var.byol_state
}
```

## Argument Reference

The following arguments are supported:

* `available_units_greater_than_or_equal_to` - (Optional) A filter to return only resources whose availableUnits greater than or equal to the given value.
* `byol_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `software_type` - (Optional) A filter to return only resources whose softwareType matches the given value.
* `state` - (Optional) A filter to return only resources whose lifecycle state matches the given value.


## Attributes Reference

The following attributes are exported:

* `byol_collection` - The list of byol_collection.

### Byol Reference

The following attributes are exported:

* `available_units` - The quantity of licensed units that not yet allocated to specific region. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the BYOL. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the BYOL.
* `display_name` - A descriptive name for the BYOL. 
* `entitlement_key` - The Broadcom-supplied identifier of a BYOL license. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL. 
* `software_type` - The type of VMware software the BYOL applies to.  Supported values:
	* VCF (VMware Cloud Foundation)
	* VSAN (VMware vSAN)
	* VDEFEND (VMware vDefend Firewall)
	* AVI_LOAD_BALANCER (VMware Avi Load Balancer) 
* `state` - The current state of the BYOL.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time the BYOL was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_term_end` - The date and time when the BYOL expires and becomes inactive. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_term_start` - The date and time when the BYOL becomes active. VMware software functionality cannot begin before this time. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the BYOL was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `total_units` - Total quantity of licensed units for the specified `softwareType`:
	* VCF, VDEFEND: number of OCPUs
	* VSAN: storage capacity in TiB (tebibytes)
	* AVI_LOAD_BALANCER: number of instances 

