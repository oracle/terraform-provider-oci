---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_byol"
sidebar_current: "docs-oci-resource-ocvp-byol"
description: |-
  Provides the Byol resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_byol
This resource provides the Byol resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/vmware/latest/Byol

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/ocvp

Creates an Oracle Cloud VMware Solution Bring-You-Own-License (BYOL).

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/WorkRequest/) operations to track the
creation of the BYOL.


## Example Usage

```hcl
resource "oci_ocvp_byol" "test_byol" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.byol_display_name
	entitlement_key = var.byol_entitlement_key
	software_type = var.byol_software_type
	time_term_end = var.byol_time_term_end
	time_term_start = var.byol_time_term_start
	total_units = var.byol_total_units

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.byol_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the BYOL. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A description of the BYOL.
* `display_name` - (Required) (Updatable) A descriptive name for the BYOL. 
* `entitlement_key` - (Required) (Updatable) The Broadcom-supplied identifier of a BYOL license. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `software_type` - (Required) (Updatable) The type of VMware software the BYOL applies to.  Supported values:
	* VCF (VMware Cloud Foundation)
	* VSAN (VMware vSAN)
	* VDEFEND (VMware vDefend Firewall)
	* AVI_LOAD_BALANCER (VMware Avi Load Balancer) 
* `time_term_end` - (Required) (Updatable) The date and time when the BYOL expires and becomes inactive. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_term_start` - (Required) (Updatable) The date and time when the BYOL becomes active. VMware software functionality cannot begin before this time. In the format defined by[RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `total_units` - (Required) (Updatable) Total quantity of licensed units for the specified `softwareType`:
	* VCF, VDEFEND: number of OCPUs
	* VSAN: storage capacity in TiB (tebibytes)
	* AVI_LOAD_BALANCER: number of instances 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Byol
	* `update` - (Defaults to 20 minutes), when updating the Byol
	* `delete` - (Defaults to 20 minutes), when destroying the Byol


## Import

Byols can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_byol.test_byol "id"
```

