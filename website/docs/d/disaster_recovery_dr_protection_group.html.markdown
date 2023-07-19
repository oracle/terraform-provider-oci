---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_protection_group"
sidebar_current: "docs-oci-datasource-disaster_recovery-dr_protection_group"
description: |-
  Provides details about a specific Dr Protection Group in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_dr_protection_group
This data source provides details about a specific Dr Protection Group resource in Oracle Cloud Infrastructure Disaster Recovery service.

Get the DR Protection Group identified by *drProtectionGroupId*.

## Example Usage

```hcl
data "oci_disaster_recovery_dr_protection_group" "test_dr_protection_group" {
	#Required
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
}
```

## Argument Reference

The following arguments are supported:

* `dr_protection_group_id` - (Required) The OCID of the DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.phx.exampleocid` 


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

