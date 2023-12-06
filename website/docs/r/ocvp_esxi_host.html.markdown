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

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/WorkRequest/) operations to track the
creation of the ESXi host.


## Example Usage

```hcl
resource "oci_ocvp_esxi_host" "test_esxi_host" {
	#Required
	cluster_id = oci_ocvp_cluster.test_cluster.id

	#Optional
	billing_donor_host_id = oci_ocvp_billing_donor_host.test_billing_donor_host.id
	capacity_reservation_id = oci_ocvp_capacity_reservation.test_capacity_reservation.id
	compute_availability_domain = var.esxi_host_compute_availability_domain
	current_commitment = var.esxi_host_current_commitment
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.esxi_host_display_name
	esxi_software_version = var.esxi_host_esxi_software_version
	freeform_tags = {"Department"= "Finance"}
	host_ocpu_count = var.esxi_host_host_ocpu_count
	host_shape_name = oci_core_shape.test_shape.name
	next_commitment = var.esxi_host_next_commitment
}
```

## How to migrate from deprecated fields to new fields
1. Before starting migration, back up your terraform state file.
2. Get the ESXi host OCID for the `oci_ocvp_esxi_host` resource you plan to migrate.
3. Remove the `oci_ocvp_esxi_host` resource from Terraform state via [command: terraform state rm \<resource address\>](https://developer.hashicorp.com/terraform/cli/commands/state/rm)
4. Update `oci_ocvp_esxi_host` resource config. Remove all deprecated fields and add corresponding new fields to the resource config.
5. Import `oci_ocvp_esxi_host` resource to Terraform state using [command: terraform import \<resource address\> \<ESXi host OCID\>](https://developer.hashicorp.com/terraform/cli/commands/import).
6. Run `terraform plan` to check if there is any planned change for `oci_ocvp_esxi_host` resource. If there is any planned change, update `oci_ocvp_esxi_host` resource config until there is no planned change.

## Argument Reference

The following arguments are supported:

* `billing_donor_host_id` - (**Deprecated**) (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deleted ESXi Host with LeftOver billing cycle. 
* `capacity_reservation_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
* `cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster to add the ESXi host to. 
* `compute_availability_domain` - (Optional) The availability domain to create the ESXi host in. If keep empty, for AD-specific Cluster, new ESXi host will be created in the same availability domain; for multi-AD Cluster, new ESXi host will be auto assigned to the next availability domain following evenly distribution strategy. 
* `current_sku` - (**Deprecated**) (Optional) The billing option currently used by the ESXi host. It is only effective during resource creation. Changes to its value after creation will be ignored. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). **Deprecated**. Please use `current_commitment` instead.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name for the ESXi host. It's changeable. Esxi Host name requirements are 1-16 character length limit, Must start with a letter,  Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the Cluster.

	If this attribute is not specified, the Cluster's `instanceDisplayNamePrefix` attribute is used to name and incrementally number the ESXi host. For example, if you're creating the fourth ESXi host in the Cluster, and `instanceDisplayNamePrefix` is `MyCluster`, the host's display name is `MyCluster-4`.

	Avoid entering confidential information. 
* `esxi_software_version` - (Optional) (Updatable) The ESXi software bundle to install on the ESXi host.  Only versions under the same vmwareSoftwareVersion and have been validate by Oracle Cloud VMware Solution will be accepted. To get a list of the available versions, use [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions). 
* `failed_esxi_host_id` - (**Deprecated**) (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that is failed. This is an optional parameter. If this parameter is specified, a new ESXi host will be created to replace the failed one, and the `failedEsxiHostId` field will be updated in the newly created Esxi host.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `host_ocpu_count` - (Optional) The OCPU count of the ESXi host. 
* `host_shape_name` - (Optional) The compute shape name of the ESXi host. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedHostShapes/ListSupportedHostShapes). 
* `next_sku` - (**Deprecated**) (Optional) (Updatable) The billing option to switch to after the existing billing cycle ends. If `nextSku` is null or empty, `currentSku` continues to the next billing cycle. In case of [SwapBilling](https://docs.oracle.com/en-us/iaas/api/#/en/vmware/20200501/EsxiHost/SwapBilling) which is not supported by Terraform, its value may be swapped with the other ESXi host. In this case, `next_sku` needs to be updated manually for both ESXi hosts in Terraform config to match the updated values. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).  **Deprecated**. Please use `next_commitment` instead.
* `non_upgraded_esxi_host_id` - (**Deprecated**) (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that will be upgraded. This is an optional parameter. If this parameter is specified, an ESXi host with the new software version is created to replace the original one, and the `nonUpgradedEsxiHostId` field is updated in the newly created Esxi host. See [Upgrading VMware Software](https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/upgrade.htm) for more information.
* `sddc_id` - (**Deprecated**) (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC to add the ESXi host to. This field has been deprecated. Please use `cluster_id` instead. Either `sddc_id` or `cluster_id` must be configured for `oci_ocvp_esxi_host` resource.  **Deprecated**. Please use `cluster_id` instead.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `billing_contract_end_date` - Current billing cycle end date. If the value in `currentCommitment` and `nextCommitment` are different, the value specified in `nextCommitment` becomes the new `currentCommitment` when the `contractEndDate` is reached. Example: `2016-08-25T21:10:29.600Z` 
* `billing_donor_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deleted ESXi Host with LeftOver billing cycle.
* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster that the ESXi host belongs to. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Cluster. 
* `compute_availability_domain` - The availability domain of the ESXi host. 
* `compute_instance_id` - In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance. 
* `current_commitment` - The billing option currently used by the ESXi host. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedCommitmentSummary/ListSupportedCommitments).
* `current_sku` - (**Deprecated**) The billing option currently used by the ESXi host. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).  **Deprecated**. Please use `current_commitment` instead.
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
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to. This field is deprecated, please use `cluster_id` instead.
* `state` - The current state of the ESXi host.
* `swap_billing_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the active ESXi Host to swap billing with current host. 
* `time_created` - The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `upgraded_replacement_esxi_host_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that is newly created to upgrade the original host. 
* `vmware_software_version` - The version of VMware software that Oracle Cloud VMware Solution installed on the ESXi hosts. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Esxi Host
	* `update` - (Defaults to 20 minutes), when updating the Esxi Host
	* `delete` - (Defaults to 20 minutes), when destroying the Esxi Host


## Import

EsxiHosts can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_esxi_host.test_esxi_host "id"
```
