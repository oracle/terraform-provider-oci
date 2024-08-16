---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_protection_groups"
sidebar_current: "docs-oci-datasource-disaster_recovery-dr_protection_groups"
description: |-
  Provides the list of Dr Protection Groups in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_dr_protection_groups
This data source provides the list of Dr Protection Groups in Oracle Cloud Infrastructure Disaster Recovery service.

Get a summary list of all DR protection groups in a compartment.

## Example Usage

```hcl
data "oci_disaster_recovery_dr_protection_groups" "test_dr_protection_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.dr_protection_group_display_name
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
	lifecycle_sub_state = var.dr_protection_group_lifecycle_sub_state
	role = var.dr_protection_group_role
	state = var.dr_protection_group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID (OCID) of the compartment in which to list resources.  Example: `ocid1.compartment.oc1..uniqueID` 
* `display_name` - (Optional) A filter to return only resources that match the given display name.  Example: `MyResourceDisplayName` 
* `dr_protection_group_id` - (Optional) The OCID of the DR protection group. Optional query param.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `lifecycle_sub_state` - (Optional) A filter to return only DR protection groups that match the given lifecycle sub-state. 
* `role` - (Optional) The DR protection group Role.
* `state` - (Optional) A filter to return only DR protection groups that match the given lifecycle state. 


## Attributes Reference

The following attributes are exported:

* `dr_protection_group_collection` - The list of dr_protection_group_collection.

### DrProtectionGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DR protection group.  Example: `ocid1.compartment.oc1..uniqueID` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the DR protection group.  Example: `EBS PHX Group` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the DR protection group.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `life_cycle_details` - A message describing the DR protection group's current state in more detail. 
* `lifecycle_sub_state` - The current sub-state of the DR protection group. 
* `log_location` - The details of an object storage log location for a DR protection group.
	* `bucket` - The bucket name inside the object storage namespace.  Example: `operation_logs` 
	* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `object` - The object name inside the object storage bucket.  Example: `switchover_plan_executions` 
* `members` - A list of DR protection group members. 
	* `autonomous_database_standby_type_for_dr_drills` - This specifies the mechanism used to create a temporary Autonomous Database instance for DR Drills. See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-clone-about.html for information about these clone types. See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-data-guard-snapshot-standby.html for information about snapshot standby. 
	* `backend_set_mappings` - A list of backend set mappings that are used to transfer or update backends during DR. 
		* `destination_backend_set_name` - The name of the destination backend set.  Example: `My_Destination_Backend_Set` 
		* `is_backend_set_for_non_movable` - This flag specifies if this backend set is used for traffic for non-movable compute instances. Backend sets that point to non-movable instances are only enabled or disabled during DR. For non-movable instances this flag should be set to 'true'. Backend sets that point to movable instances are emptied and their contents are transferred to the destination region network load balancer.  For movable instances this flag should be set to 'false'.   Example: `true` 
		* `source_backend_set_name` - The name of the source backend set.  Example: `My_Source_Backend_Set` 
	* `block_volume_operations` - Operations performed on a list of block volumes used on the non-movable compute instance. 
		* `attachment_details` - The details for attaching or detaching a block volume. 
			* `volume_attachment_reference_instance_id` - The OCID of the reference compute instance from which to obtain the attachment details for the volume. This reference compute instance is from the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `block_volume_id` - The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
		* `mount_details` - The details for mounting or unmounting the file system on a block volume. 
			* `mount_point` - The physical mount point used for mounting and unmounting the file system on a block volume.  Example: `/mnt/yourmountpoint` 
	* `connection_string_type` - The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation. See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types. 
	* `bucket` - The bucket name inside the object storage namespace.  Example: `bucket_name` 
	* `destination_availability_domain` - The availability domain of the destination mount target. Example: `BBTh:region-AD` 
	* `destination_capacity_reservation_id` - The OCID of a capacity reservation in the destination region which will be used to launch the compute instance.  Example: `ocid1.capacityreservation.oc1..uniqueID` 
	* `destination_compartment_id` - The OCID of a compartment in the destination region in which the compute instance should be launched.  Example: `ocid1.compartment.oc1..uniqueID` 
	* `destination_dedicated_vm_host_id` - The OCID of a dedicated VM host in the destination region where the compute instance should be launched.  Example: `ocid1.dedicatedvmhost.oc1..uniqueID` 
	* `destination_load_balancer_id` - The OCID of the destination load balancer. The backend sets in this destination load balancer are updated during DR.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `destination_network_load_balancer_id` - The OCID of the destination network load balancer. The backend sets in this destination network load balancer are updated during DR.                Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `export_mappings` - A list of mappings between the primary region file system export and destination region mount target. 
		* `destination_mount_target_id` - The OCID of the destination mount target on which this file system export should be created.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `export_id` - The OCID of the export path.  Example: `ocid1.export.oc1..uniqueID` 
	* `file_system_operations` - Operations performed on a list of file systems used on the non-movable compute instance. 
		* `export_path` - The export path of the file system.  Example: `/fs-export-path` 
		* `mount_details` - Mount details of a file system.
			* `mount_target_id` - The OCID of the mount target for this file system.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `mount_point` - The physical mount point of the file system on a host.  Example: `/mnt/yourmountpoint` 
		* `mount_target_id` - The OCID of mount target.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `unmount_details` - Unmount details for a file system.
			* `mount_target_id` - The OCID of the mount target for this file system.  Example: `ocid1.mounttarget.oc1..uniqueID` 
	* `is_movable` - A flag indicating if the compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - A flag indicating if the compute instance should be moved to the same fault domain in the destination region.  The compute instance launch will fail if this flag is set to true and capacity is not available in the  specified fault domain in the destination region.  Example: `false` 
	* `is_start_stop_enabled` - A flag indicating whether the non-movable compute instance needs to be started and stopped during DR operations. 
	* `member_id` - The OCID of the member.  Example: `ocid1.instance.oc1..uniqueID` 
	* `member_type` - The type of the member. 
	* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `password_vault_secret_id` - The OCID of the vault secret where the database SYSDBA password is stored. This password is required and used for performing database DR Drill operations when using full clone.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
	* `vnic_mapping` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID1, ocid1.networksecuritygroup.oc1..uniqueID2 ]` 
		* `destination_subnet_id` - The OCID of the destination subnet to which the source VNIC should connect.  Example: `ocid1.subnet.oc1..uniqueID` 
		* `source_vnic_id` - The OCID of the VNIC.  Example: `ocid1.vnic.oc1..uniqueID` 
	* `vnic_mappings` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID, ocid1.networksecuritygroup.oc1..uniqueID ]` 
		* `destination_primary_private_ip_address` - The private IP address to be assigned as the VNIC's primary IP address in the destination subnet. This must be a valid IP address in the destination subnet and the IP address must be available.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - The hostname label to be assigned in the destination subnet for the primary private IP of the source VNIC. This label is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, 'myhost1' in the FQDN 'myhost1.subnet123.vcn1.oraclevcn.com').  Example: `myhost1` 
		* `destination_subnet_id` - The OCID of the destination subnet to which the source VNIC should connect.  Example: `ocid1.subnet.oc1..uniqueID` 
		* `source_vnic_id` - The OCID of the source VNIC.  Example: `ocid1.vnic.oc1..uniqueID` 
* `peer_id` - The OCID of the peer DR protection group.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `peer_region` - The region of the peer DR protection group.  Example: `us-ashburn-1` 
* `role` - The role of the DR protection group. 
* `state` - The current state of the DR protection group. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DR protection group was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The date and time the DR protection group was updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 

