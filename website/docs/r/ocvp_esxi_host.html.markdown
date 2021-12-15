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
	compute_availability_domain = var.esxi_host_compute_availability_domain
	current_sku = var.esxi_host_current_sku
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.esxi_host_display_name
	failed_esxi_host_id = oci_ocvp_esxi_host.test_esxi_host.id
	freeform_tags = {"Department"= "Finance"}
	next_sku = var.esxi_host_next_sku
}
```

## Argument Reference

The following arguments are supported:

* `compute_availability_domain` - (Optional) The availability domain to create the ESXi host in. If keep empty, for AD-specific SDDC, new ESXi host will be created in the same availability domain; for multi-AD SDDC, new ESXi host will be auto assigned to the next availability domain following evenly distribution strategy.
* `current_sku` - (Optional) The billing option currently used by the ESXi host. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) (Updatable) A descriptive name for the ESXi host. It's changeable. Esxi Host name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the SDDC.

	If this attribute is not specified, the SDDC's `instanceDisplayNamePrefix` attribute is used to name and incrementally number the ESXi host. For example, if you're creating the fourth ESXi host in the SDDC, and `instanceDisplayNamePrefix` is `MySDDC`, the host's display name is `MySDDC-4`.

	Avoid entering confidential information.
* `failed_esxi_host_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is failed. It is an optional param, when user supplies this param, new Esxi Host will be created to replace the failed one, and failedEsxiHostId field will be udpated in the newly created EsxiHost.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `next_sku` - (Optional) (Updatable) The billing option to switch to after the existing billing cycle ends. If `nextSku` is null or empty, `currentSku` continues to the next billing cycle. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
* `sddc_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC to add the ESXi host to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `billing_contract_end_date` - Current billing cycle end date. If the value in `currentSku` and `nextSku` are different, the value specified in `nextSku` becomes the new `currentSKU` when the `contractEndDate` is reached. Example: `2016-08-25T21:10:29.600Z`
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC.
* `compute_availability_domain` - The availability domain of the ESXi host.
* `compute_instance_id` - In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance.
* `current_sku` - The billing option currently used by the ESXi host. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A descriptive name for the ESXi host. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `failed_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is failed.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `grace_period_end_date` - The date and time when the new esxi host should start billing cycle. [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2021-07-25T21:10:29.600Z` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host.
* `next_sku` - The billing option to switch to after the current billing cycle ends. If `nextSku` is null or empty, `currentSku` continues to the next billing cycle. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
* `replacement_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is newly created to replace the failed node.
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to.
* `state` - The current state of the ESXi host.
* `time_created` - The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `time_updated` - The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Esxi Host
	* `update` - (Defaults to 20 minutes), when updating the Esxi Host
	* `delete` - (Defaults to 20 minutes), when destroying the Esxi Host


## Import

EsxiHosts can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_esxi_host.test_esxi_host "id"
```
