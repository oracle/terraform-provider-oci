---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_recovery_service_subnet"
sidebar_current: "docs-oci-resource-recovery-recovery_service_subnet"
description: |-
  Provides the Recovery Service Subnet resource in Oracle Cloud Infrastructure Recovery service
---

# oci_recovery_recovery_service_subnet
This resource provides the Recovery Service Subnet resource in Oracle Cloud Infrastructure Recovery service.

Creates a new Recovery Service Subnet.


## Example Usage

```hcl
resource "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.recovery_service_subnet_display_name
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	nsg_ids = var.recovery_service_subnet_nsg_ids
	subnet_id = oci_core_subnet.test_subnet.id
	subnets = var.recovery_service_subnet_subnets
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment OCID.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - (Required) (Updatable) A user-provided name for the recovery service subnet. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `nsg_ids` - (Optional) (Updatable) A list of network security group (NSG) OCIDs that are associated with the Recovery Service subnet. You can specify a maximum of 5 unique OCIDs, which implies that you can associate a maximum of 5 NSGs to each Recovery Service subnet. Specify an empty array if you want to remove all the associated NSGs from a Recovery Service subnet. See [Network Security Groups](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/) for more information. 
* `subnet_id` - (Optional) Deprecated. One of the subnets associated with the Recovery Service subnet. 
* `subnets` - (Optional) (Updatable) A list of OCIDs of the subnets associated with the Recovery Service subnet.
* `vcn_id` - (Required) The OCID of the virtual cloud network (VCN) that contains the recovery service subnet. You can create a single recovery service subnet per VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - A user-provided name for the recovery service subnet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The recovery service subnet OCID.
* `lifecycle_details` - Detailed description about the current lifecycle state of the recovery service subnet. For example, it can be used to provide actionable information for a resource in a Failed state
* `nsg_ids` - A list of network security group (NSG) OCIDs that are associated with the Recovery Service subnet. You can specify a maximum of 5 unique OCIDs, which implies that you can associate a maximum of 5 NSGs to each Recovery Service subnet. Specify an empty array if you want to remove all the associated NSGs from a Recovery Service subnet. See [Network Security Groups](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/) for more information. 
* `state` - The current state of the recovery service subnet. 
* `subnet_id` - Deprecated. One of the subnets associated with the Recovery Service subnet. 
* `subnets` - A list of OCIDs of all the subnets associated with the Recovery Service subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_created` - An RFC3339 formatted datetime string that indicates the last created time for a recovery service subnet. For example: '2020-05-22T21:10:29.600Z'. 
* `time_updated` - An RFC3339 formatted datetime string that indicates the last updated time for a recovery service subnet. For example: '2020-05-22T21:10:29.600Z'. 
* `vcn_id` - VCN Identifier.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Recovery Service Subnet
	* `update` - (Defaults to 20 minutes), when updating the Recovery Service Subnet
	* `delete` - (Defaults to 20 minutes), when destroying the Recovery Service Subnet


## Import

RecoveryServiceSubnets can be imported using the `id`, e.g.

```
$ terraform import oci_recovery_recovery_service_subnet.test_recovery_service_subnet "id"
```

