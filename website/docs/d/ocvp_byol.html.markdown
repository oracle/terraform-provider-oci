---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_byol"
sidebar_current: "docs-oci-datasource-ocvp-byol"
description: |-
  Provides details about a specific Byol in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_byol
This data source provides details about a specific Byol resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Gets the specified BYOL's information.

## Example Usage

```hcl
data "oci_ocvp_byol" "test_byol" {
	#Required
	byol_id = oci_ocvp_byol.test_byol.id
}
```

## Argument Reference

The following arguments are supported:

* `byol_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL.


## Attributes Reference

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

