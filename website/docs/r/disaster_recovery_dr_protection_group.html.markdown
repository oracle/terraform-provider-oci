---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_protection_group"
sidebar_current: "docs-oci-resource-disaster_recovery-dr_protection_group"
description: |-
  Provides the Dr Protection Group resource in Oracle Cloud Infrastructure Disaster Recovery service
---

# oci_disaster_recovery_dr_protection_group
This resource provides the Dr Protection Group resource in Oracle Cloud Infrastructure Disaster Recovery service.

Create a new DR Protection Group.

## Example Usage

```hcl
resource "oci_disaster_recovery_dr_protection_group" "test_dr_protection_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.dr_protection_group_display_name
	log_location {
		#Required
		bucket = var.dr_protection_group_log_location_bucket
		namespace = var.dr_protection_group_log_location_namespace
	}

	#Optional
	association {
		#Required
		role = var.dr_protection_group_association_role

		#Optional
		peer_id = oci_blockchain_peer.test_peer.id
		peer_region = var.dr_protection_group_association_peer_region
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	members {
		#Required
		member_id = oci_disaster_recovery_member.test_member.id
		member_type = var.dr_protection_group_members_member_type

		#Optional
		destination_capacity_reservation_id = oci_disaster_recovery_destination_capacity_reservation.test_destination_capacity_reservation.id
		destination_compartment_id = oci_identity_compartment.test_compartment.id
		destination_dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
		is_movable = var.dr_protection_group_members_is_movable
		is_retain_fault_domain = var.dr_protection_group_members_is_retain_fault_domain
		password_vault_secret_id = oci_vault_secret.test_secret.id
		vnic_mapping {

			#Optional
			destination_nsg_id_list = var.dr_protection_group_members_vnic_mapping_destination_nsg_id_list
			destination_primary_private_ip_address = var.dr_protection_group_members_vnic_mapping_destination_primary_private_ip_address
			destination_primary_private_ip_hostname_label = var.dr_protection_group_members_vnic_mapping_destination_primary_private_ip_hostname_label
			destination_subnet_id = oci_core_subnet.test_subnet.id
			source_vnic_id = oci_core_vnic.test_vnic.id
		}
		vnic_mappings {

			#Optional
			destination_nsg_id_list = var.dr_protection_group_members_vnic_mappings_destination_nsg_id_list
			destination_primary_private_ip_address = var.dr_protection_group_members_vnic_mappings_destination_primary_private_ip_address
			destination_primary_private_ip_hostname_label = var.dr_protection_group_members_vnic_mappings_destination_primary_private_ip_hostname_label
			destination_subnet_id = oci_core_subnet.test_subnet.id
			source_vnic_id = oci_core_vnic.test_vnic.id
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `association` - (Optional) The details for associating this DR Protection Group with a peer (remote) DR Protection Group.
	* `peer_id` - (Optional) The OCID of the peer (remote) DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.iad.&lt;unique_id&gt;` 
	* `peer_region` - (Optional) The region of the peer (remote) DR Protection Group.  Example: `us-ashburn-1` 
	* `role` - (Required) The role of this DR Protection Group. 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the DR Protection Group.  Example: `ocid1.compartment.oc1..&lt;unique_id&gt;` 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) The display name of the DR Protection Group.  Example: `EBS PHX DRPG` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `log_location` - (Required) (Updatable) Information about creating an Object Storage log location for a DR Protection Group.
	* `bucket` - (Required) (Updatable) The bucket name inside the Object Storage namespace.  Example: `operation_logs` 
	* `namespace` - (Required) (Updatable) The namespace in Object Storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
* `members` - (Optional) (Updatable) A list of DR Protection Group members. 
	* `destination_capacity_reservation_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the capacity reservation in the destination region using which this compute instance should be launched.  Example: `ocid1.capacityreservation.oc1..&lt;unique_id&gt;` 
	* `destination_compartment_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the compartment for this compute instance in the destination region.  Example: `ocid1.compartment.oc1..&lt;unique_id&gt;` 
	* `destination_dedicated_vm_host_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the dedicated VM Host in the destination region where this compute instance should be launched  Example: `ocid1.dedicatedvmhost.oc1.iad.&lt;unique_id&gt;` 
	* `is_movable` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A flag indicating if this compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A flag indicating if this compute instance should be moved to the same fault domain.  Compute instance launch will fail if this flag is set to true and capacity is not available in that specific fault domain in the destination region.  Example: `false` 
	* `member_id` - (Required) (Updatable) The OCID of the member.  Example: `ocid1.instance.oc1.phx.&lt;unique_id&gt;` 
	* `member_type` - (Required) (Updatable) The type of the member. 
	* `password_vault_secret_id` - (Applicable when member_type=DATABASE) (Updatable) The OCID of the vault secret where the database password is stored.  Example: `ocid1.vaultsecret.oc1.phx.&lt;unique_id&gt;` 
	* `vnic_mapping` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A list of Compute Instance VNIC mappings. 
		* `destination_nsg_id_list` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A list of network security group (NSG) IDs in the destination region which this VNIC should use.  Example: `[ ocid1.networksecuritygroup.oc1.iad.&lt;unique_id&gt;, ocid1.networksecuritygroup.oc1.iad.&lt;unique_id&gt; ]` 
		* `destination_primary_private_ip_address` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) The primary private IP address to assign. This address must belong to the destination subnet.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) The hostname to assign for this primary private IP. The value is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com).  Example: `bminstance1` 
		* `destination_subnet_id` - (Required when member_type=COMPUTE_INSTANCE) (Updatable) The OCID of the destination (remote) subnet to which this VNIC should connect.  Example: `ocid1.subnet.oc1.iad.&lt;unique_id&gt;` 
		* `source_vnic_id` - (Required when member_type=COMPUTE_INSTANCE) (Updatable) The OCID of the VNIC.  Example: `ocid1.vnic.oc1.phx.&lt;unique_id&gt;` 
	* `vnic_mappings` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A list of Compute Instance VNIC mappings. 
		* `destination_nsg_id_list` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A list of network security group (NSG) IDs in the destination region which this VNIC should use.  Example: `[ ocid1.networksecuritygroup.oc1..&lt;unique_id&gt;, ocid1.networksecuritygroup.oc1..&lt;unique_id&gt; ]` 
		* `destination_primary_private_ip_address` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The primary private IP address to assign. This address must belong to the destination subnet.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The hostname to assign for this primary private IP. The value is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com).  Example: `bminstance1` 
		* `destination_subnet_id` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the destination (remote) subnet to which this VNIC should connect.  Example: `ocid1.subnet.oc1..&lt;unique_id&gt;` 
		* `source_vnic_id` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the VNIC.  Example: `ocid1.vnic.oc1..&lt;unique_id&gt;` 
* `disassociate_trigger` - (Optional) (Updatable) An optional property when incremented triggers Disassociate. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DR Protection Group.  Example: `ocid1.compartment.oc1..&lt;unique_id&gt;` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the DR Protection Group.  Example: `EBS PHX DRPG` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The OCID of the DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.phx.&lt;unique_id&gt;` 
* `life_cycle_details` - A message describing the DR Protection Group's current state in more detail. 
* `log_location` - Information about an Object Storage log location for a DR Protection Group.
	* `bucket` - The bucket name inside the Object Storage namespace.  Example: `operation_logs` 
	* `namespace` - The namespace in Object Storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `object` - The object name inside the Object Storage bucket.  Example: `switchover_plan_executions` 
* `members` - A list of DR Protection Group members. 
	* `destination_capacity_reservation_id` - The OCID of the capacity reservation in the destination region using which this compute instance should be launched.  Example: `ocid1.capacityreservation.oc1..&lt;unique_id&gt;` 
	* `destination_compartment_id` - The OCID of the compartment for this compute instance in the destination region.  Example: `ocid1.compartment.oc1..&lt;unique_id&gt;` 
	* `destination_dedicated_vm_host_id` - The OCID of the dedicated VM Host for this compute instance in the destination region.  Example: `ocid1.dedicatedvmhost.oc1..&lt;unique_id&gt;` 
	* `is_movable` - A flag indicating if this compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - A flag indicating if this compute instance should be moved to the same fault domain.  Compute instance launch will fail if this flag is set to true and capacity is not available in that specific fault domain in the destination region.  Example: `false` 
	* `member_id` - The OCID of the member.  Example: `ocid1.instance.oc1.phx.&lt;unique_id&gt;` 
	* `member_type` - The type of the member. 
	* `password_vault_secret_id` - The ID of the vault secret where the database password is stored.  Example: `ocid1.vaultsecret.oc1.phx.exampleocid1` 
	* `vnic_mapping` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of destination region's network security group (NSG) OCIDs which this VNIC should use.  Example: `[ ocid1.networksecuritygroup.oc1.iad.exampleocid1, ocid1.networksecuritygroup.oc1.iad.exampleocid2 ]` 
		* `destination_subnet_id` - The OCID of the destination (remote) subnet to which this VNIC should connect.  Example: `ocid1.subnet.oc1.iad.exampleocid` 
		* `source_vnic_id` - The OCID of the VNIC.  Example: `ocid1.vnic.oc1.phx.exampleocid` 
	* `vnic_mappings` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of destination region's network security group (NSG) OCIDs which this VNIC should use.  Example: `[ ocid1.networksecuritygroup.oc1..&lt;unique_id&gt;, ocid1.networksecuritygroup.oc1..&lt;unique_id&gt; ]` 
		* `destination_primary_private_ip_address` - The primary private IP address to assign. This address must belong to the destination subnet.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - The hostname to assign for this primary private IP. The value is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com).  Example: `bminstance1` 
		* `destination_subnet_id` - The OCID of the destination (remote) subnet to which this VNIC should connect.  Example: `ocid1.subnet.oc1..&lt;unique_id&gt;` 
		* `source_vnic_id` - The OCID of the VNIC.  Example: `ocid1.vnic.oc1..&lt;unique_id&gt;` 
* `peer_id` - The OCID of the peer (remote) DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.iad.&lt;unique_id&gt;` 
* `peer_region` - The region of the peer (remote) DR Protection Group.  Example: `us-ashburn-1` 
* `role` - The role of the DR Protection Group. 
* `state` - The current state of the DR Protection Group. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DR Protection Group was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The date and time the DR Protection Group was updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dr Protection Group
	* `update` - (Defaults to 20 minutes), when updating the Dr Protection Group
	* `delete` - (Defaults to 20 minutes), when destroying the Dr Protection Group


## Import

DrProtectionGroups can be imported using the `id`, e.g.

```
$ terraform import oci_disaster_recovery_dr_protection_group.test_dr_protection_group "id"
```

