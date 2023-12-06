---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_esxi_host"
sidebar_current: "docs-oci-datasource-ocvp-esxi_host"
description: |-
  Provides details about a specific Esxi Host in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_esxi_host
This data source provides details about a specific Esxi Host resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Gets the specified ESXi host's information.

## Example Usage

```hcl
data "oci_ocvp_esxi_host" "test_esxi_host" {
	#Required
	esxi_host_id = oci_ocvp_esxi_host.test_esxi_host.id
}
```

## Argument Reference

The following arguments are supported:

* `esxi_host_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host.


## Attributes Reference

The following attributes are exported:

* `billing_contract_end_date` - Current billing cycle end date. If the value in `currentCommitment` and `nextCommitment` are different, the value specified in `nextCommitment` becomes the new `currentCommitment` when the `contractEndDate` is reached. Example: `2016-08-25T21:10:29.600Z` 
* `billing_donor_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deleted ESXi Host with LeftOver billing cycle.
* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster that the ESXi host belongs to. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Cluster. 
* `compute_availability_domain` - The availability domain of the ESXi host. 
* `compute_instance_id` - In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance. 
* `current_sku` - (**Deprecated**) The billing option currently used by the ESXi host. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).  **Deprecated**. Please use `current_commitment` instead.
* `current_commitment` - The billing option currently used by the ESXi host. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedCommitmentSummary/ListSupportedCommitments).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the ESXi host. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `esxi_software_version` - The version of ESXi software that Oracle Cloud VMware Solution installed on the ESXi hosts. 
* `failed_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that failed. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `grace_period_end_date` - The date and time when the new esxi host should start billing cycle. [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2021-07-25T21:10:29.600Z` 
* `host_ocpu_count` - The OCPU count of the ESXi host. 
* `host_shape_name` - The compute shape name of the ESXi host. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedHostShapes/ListSupportedHostShapes). 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host. 
* `is_billing_continuation_in_progress` - Indicates whether this host is in the progress of billing continuation. 
* `is_billing_swapping_in_progress` - Indicates whether this host is in the progress of swapping billing.
* `next_commitment` - The billing option to switch to after the current billing cycle ends. If `nextCommitment` is null or empty, `currentCommitment` continues to the next billing cycle. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedCommitmentSummary/ListSupportedCommitments).
* `next_sku` - (**Deprecated**) The billing option to switch to after the current billing cycle ends. If `nextSku` is null or empty, `currentSku` continues to the next billing cycle. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).  **Deprecated**. Please use `next_commitment` instead.
* `non_upgraded_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that will be upgraded. 
* `replacement_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the esxi host that is newly created to replace the failed node.
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to.
* `state` - The current state of the ESXi host.
* `swap_billing_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the active ESXi Host to swap billing with current host. 
* `time_created` - The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `upgraded_replacement_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that is newly created to upgrade the original host. 
* `vmware_software_version` - The version of VMware software that Oracle Cloud VMware Solution installed on the ESXi hosts. 

