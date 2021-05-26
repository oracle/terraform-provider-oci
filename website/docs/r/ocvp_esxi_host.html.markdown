---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_esxi_host"
sidebar_current: "docs-oci-resource-ocvp-esxi_host"
description: |-
  Provides the Esxi Host resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_esxi_host
This resource provides the Esxi Host resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Adds another ESXi host to an existing SDDC. The attributes of the specified
`Sddc` determine the VMware software and other configuration settings used
by the ESXi host.

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/WorkRequest/) operations to track the
creation of the ESXi host.


## Example Usage

```hcl
resource "oci_ocvp_esxi_host" "test_esxi_host" {
	#Required
	sddc_id = oci_ocvp_sddc.test_sddc.id

	#Optional
	current_sku = var.esxi_host_current_sku
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.esxi_host_display_name
	freeform_tags = {"Department"= "Finance"}
	next_sku = var.esxi_host_next_sku
}
```

## Argument Reference

The following arguments are supported:

* `current_sku` - (Optional) Billing option selected during SDDC creation. Oracle Cloud Infrastructure VMware Solution supports the following billing interval SKUs: HOUR, MONTH, ONE_YEAR, and THREE_YEARS. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name for the ESXi host. It's changeable. Esxi Host name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the SDDC.

	If this attribute is not specified, the SDDC's `instanceDisplayNamePrefix` attribute is used to name and incrementally number the ESXi host. For example, if you're creating the fourth ESXi host in the SDDC, and `instanceDisplayNamePrefix` is `MySDDC`, the host's display name is `MySDDC-4`.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `next_sku` - (Optional) (Updatable) Billing option to switch to once existing billing cycle ends. If nextSku is null or empty, currentSku will be used to continue with next billing term. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `sddc_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC to add the ESXi host to. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `billing_contract_end_date` - Current billing cycle end date. If nextSku is different from existing SKU, then we switch to newSKu after this contractEndDate Example: `2016-08-25T21:10:29.600Z` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC. 
* `compute_instance_id` - In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance. 
* `current_sku` - Billing option selected during SDDC creation. Oracle Cloud Infrastructure VMware Solution supports the following billing interval SKUs: HOUR, MONTH, ONE_YEAR, and THREE_YEARS. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the ESXi host. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host. 
* `next_sku` - Billing option to switch to once existing billing cycle ends. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to. 
* `state` - The current state of the ESXi host.
* `time_created` - The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

## Import

EsxiHosts can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_esxi_host.test_esxi_host "id"
```

