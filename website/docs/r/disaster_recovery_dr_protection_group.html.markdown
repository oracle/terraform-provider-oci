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

Create a DR protection group.

## Example Usage

```hcl
variable "disassociate_trigger" { default = 0 }

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
		peer_id = var.dr_protection_group_association_peer_id
		peer_region = var.dr_protection_group_association_peer_region
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	disassociate_trigger = var.disassociate_trigger

	members {
		#Required
		member_id = var.dr_protection_group_members_member_id
		member_type = var.dr_protection_group_members_member_type

		#Optional
		autonomous_database_standby_type_for_dr_drills = var.dr_protection_group_members_autonomous_database_standby_type_for_dr_drills
		backend_set_mappings {

			#Optional
			destination_backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
			is_backend_set_for_non_movable = var.dr_protection_group_members_backend_set_mappings_is_backend_set_for_non_movable
			source_backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
		}
		block_volume_operations {

			#Optional
			attachment_details {

				#Optional
				volume_attachment_reference_instance_id = oci_core_instance.test_instance.id
			}
			block_volume_id = oci_core_volume.test_volume.id
			mount_details {

				#Optional
				mount_point = var.dr_protection_group_members_block_volume_operations_mount_details_mount_point
			}
		}
		connection_string_type = var.dr_protection_group_members_connection_string_type
		bucket = var.dr_protection_group_members_bucket
		destination_availability_domain = var.dr_protection_group_members_destination_availability_domain
		destination_capacity_reservation_id = var.destination_capacity_reservation_id
		destination_compartment_id = oci_identity_compartment.test_compartment.id
		destination_dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
		destination_load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		destination_network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		export_mappings {

			#Optional
			destination_mount_target_id = oci_file_storage_mount_target.test_mount_target.id
			export_id = oci_file_storage_export.test_export.id
		}
		file_system_operations {

			#Optional
			export_path = var.dr_protection_group_members_file_system_operations_export_path
			mount_details {

				#Optional
				mount_target_id = oci_file_storage_mount_target.test_mount_target.id
			}
			mount_point = var.dr_protection_group_members_file_system_operations_mount_point
			mount_target_id = oci_file_storage_mount_target.test_mount_target.id
			unmount_details {

				#Optional
				mount_target_id = oci_file_storage_mount_target.test_mount_target.id
			}
		}
		is_movable = var.dr_protection_group_members_is_movable
		is_retain_fault_domain = var.dr_protection_group_members_is_retain_fault_domain
		is_start_stop_enabled = var.dr_protection_group_members_is_start_stop_enabled
		namespace = var.dr_protection_group_members_namespace
		password_vault_secret_id = var.password_vault_secret_id
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

* `association` - (Optional) The details for associating a DR protection group with a peer DR protection group.
	* `peer_id` - (Optional) The OCID of the peer DR protection group.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
	* `peer_region` - (Optional) The region of the peer DR protection group.  Example: `us-ashburn-1` 
	* `role` - (Required) The role of the DR protection group.  Example: `STANDBY` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the DR protection group.  Example: `ocid1.compartment.oc1..uniqueID` 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) The display name of the DR protection group.  Example: `EBS PHX Group` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `log_location` - (Required) (Updatable) The details for creating an object storage log location for a DR protection group.
	* `bucket` - (Required) (Updatable) The bucket name inside the object storage namespace.  Example: `operation_logs` 
	* `namespace` - (Required) (Updatable) The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
* `members` - (Optional) (Updatable) A list of DR protection group members. 
	* `autonomous_database_standby_type_for_dr_drills` - (Applicable when member_type=AUTONOMOUS_DATABASE) (Updatable) This specifies the mechanism used to create a temporary Autonomous Database instance for DR Drills. See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-clone-about.html for information about these clone types. See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-data-guard-snapshot-standby.html for information about snapshot standby. 
	* `backend_set_mappings` - (Applicable when member_type=LOAD_BALANCER | NETWORK_LOAD_BALANCER) (Updatable) A list of backend set mappings that are used to transfer or update backends during DR. 
		* `destination_backend_set_name` - (Required when member_type=LOAD_BALANCER | NETWORK_LOAD_BALANCER) (Updatable) The name of the destination backend set.  Example: `Destination-BackendSet-1` 
		* `is_backend_set_for_non_movable` - (Required when member_type=LOAD_BALANCER | NETWORK_LOAD_BALANCER) (Updatable) This flag specifies if this backend set is used for traffic for non-movable compute instances. Backend sets that point to non-movable instances are only enabled or disabled during DR, their contents are not altered. For non-movable instances this flag should be set to 'true'. Backend sets that point to movable instances are emptied and their contents are transferred to the  destination region load balancer.  For movable instances this flag should be set to 'false'.   Example: `true` 
		* `source_backend_set_name` - (Required when member_type=LOAD_BALANCER | NETWORK_LOAD_BALANCER) (Updatable) The name of the source backend set.  Example: `Source-BackendSet-1` 
	* `block_volume_operations` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A list of operations performed on block volumes used by the compute instance. 
		* `attachment_details` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The details for creating a block volume attachment. 
			* `volume_attachment_reference_instance_id` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the reference compute instance from which to obtain the attachment details for the volume. This reference compute instance is from the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `block_volume_id` - (Required when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
		* `mount_details` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The details for creating a mount for a file system on a block volume. 
			* `mount_point` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The physical mount point used for mounting the file system on the block volume.  Example: `/mnt/yourmountpoint` 
	* `connection_string_type` - (Applicable when member_type=AUTONOMOUS_CONTAINER_DATABASE) (Updatable) The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation. See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types. 
	* `bucket` - (Required when member_type=OBJECT_STORAGE_BUCKET) (Updatable) The bucket name inside the object storage namespace.  Example: `bucket_name` 
	* `destination_availability_domain` - (Applicable when member_type=FILE_SYSTEM) (Updatable) The availability domain of the destination mount target.  Example: `BBTh:region-AD` 
	* `destination_capacity_reservation_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of a capacity reservation in the destination region which will be used to launch the compute instance.  Example: `ocid1.capacityreservation.oc1..uniqueID` 
	* `destination_compartment_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of a compartment in the destination region in which the compute instance should be launched.  Example: `ocid1.compartment.oc1..uniqueID` 
	* `destination_dedicated_vm_host_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of a dedicated VM host in the destination region where the compute instance should be launched.  Example: `ocid1.dedicatedvmhost.oc1..uniqueID` 
	* `destination_load_balancer_id` - (Applicable when member_type=LOAD_BALANCER) (Updatable) The OCID of the destination load balancer.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `destination_network_load_balancer_id` - (Applicable when member_type=NETWORK_LOAD_BALANCER) (Updatable) The OCID of the destination network load balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `export_mappings` - (Applicable when member_type=FILE_SYSTEM) (Updatable) A list of mappings between file system exports in the primary region and mount targets in the standby region. 
		* `destination_mount_target_id` - (Required when member_type=FILE_SYSTEM) (Updatable) The OCID of the destination mount target in the destination region which is used to export the file system.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `export_id` - (Required when member_type=FILE_SYSTEM) (Updatable) The OCID of the export path in the primary region used to mount or unmount the file system.  Example: `ocid1.export.oc1..uniqueID` 
	* `file_system_operations` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE | COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A list of operations performed on file systems used by the compute instance. 
		* `export_path` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE | COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The export path of the file system.  Example: `/fs-export-path` 
		* `mount_details` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The details for creating a file system mount. 
			* `mount_target_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the mount target for this file system.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `mount_point` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE | COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The physical mount point of the file system on a host.  Example: `/mnt/yourmountpoint` 
		* `mount_target_id` - (Required when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the mount target.  Example: `ocid1.mounttarget.oc1..uniqueID` 
		* `unmount_details` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The details for creating a file system unmount. 
			* `mount_target_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the mount target.  Example: `ocid1.mounttarget.oc1..uniqueID` 
	* `is_movable` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A flag indicating if the compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A flag indicating if the compute instance should be moved to the same fault domain in the destination region.  The compute instance launch will fail if this flag is set to true and capacity is not available in the  specified fault domain in the destination region.  Example: `false` 
	* `is_start_stop_enabled` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A flag indicating whether the non-movable compute instance should be started and stopped during DR operations. *Prechecks cannot be executed on stopped instances that are configured to be started.* 
	* `member_id` - (Required) (Updatable) The OCID of the member.  Example: `ocid1.instance.oc1..uniqueID` 
	* `member_type` - (Required) (Updatable) The type of the member. 
	* `namespace` - (Required when member_type=OBJECT_STORAGE_BUCKET) (Updatable) The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `password_vault_secret_id` - (Applicable when member_type=AUTONOMOUS_DATABASE | DATABASE) (Updatable) The OCID of the vault secret where the database SYSDBA password is stored. This password is required and used for performing database DR Drill operations when using full clone.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
	* `vnic_mapping` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID, ocid1.networksecuritygroup.oc1..uniqueID ]` 
		* `destination_primary_private_ip_address` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) The primary private IP address to be assigned to the VNIC in the destination region.  This address must belong to the destination subnet.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) The hostname label to be assigned in the destination subnet for the primary private IP of the source VNIC. This label is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, 'myhost1' in the FQDN 'myhost1.subnet123.vcn1.oraclevcn.com').  Example: `myhost1` 
		* `destination_subnet_id` - (Required when member_type=COMPUTE_INSTANCE) (Updatable) The OCID of the destination subnet to which this source VNIC should connect.  Example: `ocid1.subnet.oc1..uniqueID` 
		* `source_vnic_id` - (Required when member_type=COMPUTE_INSTANCE) (Updatable) The OCID of the VNIC.  Example: `ocid1.vnic.oc1..uniqueID` 
	* `vnic_mappings` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID, ocid1.networksecuritygroup.oc1..uniqueID ]` 
		* `destination_primary_private_ip_address` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The primary private IP address to be assigned to the source VNIC in the destination subnet.  This IP address must belong to the destination subnet.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The hostname label to be assigned in the destination subnet for the primary private IP of the source VNIC. This label is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, 'myhost1' in the FQDN 'myhost1.subnet123.vcn1.oraclevcn.com').  Example: `myhost1` 
		* `destination_subnet_id` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the destination subnet to which the source VNIC should connect.          Example: `ocid1.subnet.oc1..uniqueID` 
		* `source_vnic_id` - (Required when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the source VNIC.  Example: `ocid1.vnic.oc1..uniqueID` 
* `disassociate_trigger` - (Optional) (Updatable) An optional property when incremented triggers Disassociate. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
	* `bucket` - The bucket name inside the object storage namespace.  Example: `bucket_name` 
	* `connection_string_type` - The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation. See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types. 
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dr Protection Group
	* `update` - (Defaults to 20 minutes), when updating the Dr Protection Group
	* `delete` - (Defaults to 20 minutes), when destroying the Dr Protection Group

## Create

Create DR Protection Group resource with a default value of `disassociate_trigger` property, e.g.

```
terraform apply -var "disassociate_trigger=0"
```

## Delete

Disassociate DR Protection Group (if associated) before deleting it. Increment value of `disassociate_trigger` property to trigger Disassociate, e.g.

```
terraform destroy -var "disassociate_trigger=1"
```


## Import

DrProtectionGroups can be imported using the `id`, e.g.

```
$ terraform import oci_disaster_recovery_dr_protection_group.test_dr_protection_group "id"
```

