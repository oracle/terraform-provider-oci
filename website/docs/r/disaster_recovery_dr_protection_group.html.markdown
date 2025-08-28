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
		backup_config {

			#Optional
			backup_schedule = var.dr_protection_group_members_backup_config_backup_schedule
			exclude_namespaces = var.dr_protection_group_members_backup_config_exclude_namespaces
			image_replication_vault_secret_id = oci_vault_secret.test_secret.id
			max_number_of_backups_retained = var.dr_protection_group_members_backup_config_max_number_of_backups_retained
			namespaces = var.dr_protection_group_members_backup_config_namespaces
			replicate_images = var.dr_protection_group_members_backup_config_replicate_images
		}
		backup_location {

			#Optional
			bucket = var.dr_protection_group_members_backup_location_bucket
			namespace = var.dr_protection_group_members_backup_location_namespace
		}
		block_volume_attach_and_mount_operations {

			#Optional
			attachments {

				#Optional
				block_volume_id = oci_core_volume.test_volume.id
				volume_attachment_reference_instance_id = oci_core_instance.test_instance.id
			}
			mounts {

				#Optional
				mount_point = var.dr_protection_group_members_block_volume_attach_and_mount_operations_mounts_mount_point
			}
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
		bucket = var.dr_protection_group_members_bucket
		common_destination_key {

			#Optional
			encryption_key_id = oci_kms_key.test_key.id
			vault_id = oci_kms_vault.test_vault.id
		}
		connection_string_type = var.dr_protection_group_members_connection_string_type
		db_system_admin_user_details {

			#Optional
			password_vault_secret_id = oci_vault_secret.test_secret.id
			username = var.dr_protection_group_members_db_system_admin_user_details_username
		}
		db_system_replication_user_details {

			#Optional
			password_vault_secret_id = oci_vault_secret.test_secret.id
			username = var.dr_protection_group_members_db_system_replication_user_details_username
		}
		destination_availability_domain = var.dr_protection_group_members_destination_availability_domain
		destination_backup_policy_id = oci_identity_policy.test_policy.id
		destination_capacity_reservation_id = var.destination_capacity_reservation_id
		destination_compartment_id = oci_identity_compartment.test_compartment.id
		destination_dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
		destination_encryption_key {

			#Optional
			encryption_key_id = oci_kms_key.test_key.id
			vault_id = oci_kms_vault.test_vault.id
		}
		destination_load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		destination_network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		destination_snapshot_policy_id = oci_identity_policy.test_policy.id
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
		gtid_reconciliation_timeout = var.dr_protection_group_members_gtid_reconciliation_timeout
		is_continue_on_gtid_reconciliation_timeout = var.dr_protection_group_members_is_continue_on_gtid_reconciliation_timeout
		is_movable = var.dr_protection_group_members_is_movable
		is_retain_fault_domain = var.dr_protection_group_members_is_retain_fault_domain
		is_start_stop_enabled = var.dr_protection_group_members_is_start_stop_enabled
		jump_host_id = oci_disaster_recovery_jump_host.test_jump_host.id
		load_balancer_mappings {

			#Optional
			destination_load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
			source_load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		}
		managed_node_pool_configs {

			#Optional
			id = var.dr_protection_group_members_managed_node_pool_configs_id
			maximum = var.dr_protection_group_members_managed_node_pool_configs_maximum
			minimum = var.dr_protection_group_members_managed_node_pool_configs_minimum
		}
		namespace = var.dr_protection_group_members_namespace
		password_vault_secret_id = var.password_vault_secret_id
		network_load_balancer_mappings {

			#Optional
			destination_network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
			source_network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
		}
		peer_cluster_id = oci_containerengine_cluster.test_cluster.id
		peer_db_system_id = oci_database_db_system.test_db_system.id
		source_volume_to_destination_encryption_key_mappings {

			#Optional
			destination_encryption_key {

				#Optional
				encryption_key_id = oci_kms_key.test_key.id
				vault_id = oci_kms_vault.test_vault.id
			}
			source_volume_id = oci_core_volume.test_volume.id
		}
		vault_mappings {

			#Optional
			destination_vault_id = oci_kms_vault.test_vault.id
			source_vault_id = oci_kms_vault.test_vault.id
		}
		virtual_node_pool_configs {

			#Optional
			id = var.dr_protection_group_members_virtual_node_pool_configs_id
			maximum = var.dr_protection_group_members_virtual_node_pool_configs_maximum
			minimum = var.dr_protection_group_members_virtual_node_pool_configs_minimum
		}
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
			destination_reserved_public_ip_id = oci_core_public_ip.test_public_ip.id
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
	* `backup_config` - (Applicable when member_type=OKE_CLUSTER) (Updatable) Create backup configuration properties for an OKE member.
		* `backup_schedule` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The schedule for backing up namespaces to the destination region. If a backup schedule is not specified, only a single backup will be created.  This format of the string specifying the backup schedule must conform with RFC-5545 (see examples below). This schedule will use the UTC timezone. This property applies to the OKE cluster member in primary region.

			The backup frequency can be HOURLY, DAILY, WEEKLY or MONTHLY, and the upper and lower interval bounds are as follows HOURLY
			* Minimum = 1
			* Maximum = 24 DAILY
			* Minimum = 1
			* Maximum = 30 WEEKLY
			* Minimum = 1
			* Maximum = 1 MONTHLY
			* Minimum = 1
			* Maximum = 12

			Examples:  FREQ=WEEKLY;BYDAY=MO,WE;BYHOUR=10;INTERVAL=1 -> Run a backup every Monday and Wednesday at 10:00 AM. FREQ=WEEKLY;BYDAY=MO,WE;BYHOUR=10;INTERVAL=2 -> Invalid configuration (Cannot specify an interval of 2).

			FREQ=HOURLY;INTERVAL=25 -> Invalid configuration (Cannot specify an interval of 25). FREQ=HOURLY;INTERVAL=0 -> Invalid configuration (Cannot specify an interval of 0). FREQ=HOURLY;INTERVAL=24 -> Run a backup every 24 hours. FREQ=HOURLY;INTERVAL=1 -> Run a backup every hour. FREQ=HOURLY;BYMINUTE=30;INTERVAL=15 -> Run a backup every 15 hours at the 30th minute. FREQ=DAILY;INTERVAL=31 -> Invalid configuration (Cannot specify an interval of 31). FREQ=DAILY;INTERVAL=0 -> Invalid configuration (Cannot specify an interval of 0). FREQ=DAILY;INTERVAL=30 -> Run a backup every 30 days at 12:00 midnight.  FREQ=DAILY;BYHOUR=17;BYMINUTE=10;INTERVAL=1 -> Run a backup daily at 05:10 PM. 
		* `exclude_namespaces` - (Applicable when member_type=OKE_CLUSTER) (Updatable) A list of namespaces to be excluded from the backup.  The default value is null. If a list of namespaces to exclude is not provided, all namespaces will be backed up. Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both. This property applies to OKE cluster members in the primary region.  Example: ["namespace_string_3", "namespace_string_4"] 
		* `image_replication_vault_secret_id` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The OCID of the vault secret that stores the image credential. This property applies to the OKE cluster member in both the primary and standby region. 
		* `max_number_of_backups_retained` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The maximum number of backups that should be retained. This property applies to the OKE cluster member in primary region. 
		* `namespaces` - (Applicable when member_type=OKE_CLUSTER) (Updatable) A list of namespaces to be included in the backup.  The default value is null. If a list of namespaces to include is not provided, all namespaces will be backed up. Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both. This property applies to the OKE cluster member in primary region.  Example: ["default", "pv-nginx"] 
		* `replicate_images` - (Applicable when member_type=OKE_CLUSTER) (Updatable) Controls the behaviour of image replication across regions. Image replication is enabled by default for DR Protection Groups with a primary role. This property applies to the OKE cluster member in primary region. 
	* `backup_location` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The details for creating the backup location of an OKE Cluster.
		* `bucket` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The bucket name inside the object storage namespace.  Example: `operation_logs` 
		* `namespace` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The namespace in the object storage bucket location (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `block_volume_attach_and_mount_operations` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The details for creating the operations performed on a block volume. 
		* `attachments` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A list of details of attach or detach operations performed on block volumes. 
			* `block_volume_id` - (Required when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
			* `volume_attachment_reference_instance_id` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the reference compute instance needed to obtain the volume attachment details. This reference compute instance belongs to the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `mounts` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A list of details of mount operations performed on block volumes. 
			* `mount_point` - (Required when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The physical mount point where the file system is mounted on the block volume.  Example: `/mnt/yourmountpoint` 
	* `block_volume_operations` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) Deprecated. Use the 'blockVolumeAttachAndMountOperations' attribute instead of this. A list of operations performed on block volumes used by the compute instance. 
		* `attachment_details` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) Deprecated. Use the 'CreateComputeInstanceNonMovableBlockVolumeAttachOperationDetails' definition instead of this. The details for creating a block volume attachment. 
			* `volume_attachment_reference_instance_id` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the reference compute instance needed to obtain the volume attachment details. This reference compute instance belongs to the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `block_volume_id` - (Required when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
		* `mount_details` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The details for creating a mount for a file system on a block volume. 
			* `mount_point` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) The physical mount point used for mounting the file system on the block volume.  Example: `/mnt/yourmountpoint` 
	* `bucket` - (Required when member_type=OBJECT_STORAGE_BUCKET) (Updatable) The bucket name inside the object storage namespace.  Example: `bucket_name` 
	* `common_destination_key` - (Applicable when member_type=VOLUME_GROUP) (Updatable) Create properties for a customer-managed vault and encryption key in the destination region.  The customer-managed encryption key in this will be used to encrypt the resource or containing resources after they  move to the destination region. 
		* `encryption_key_id` - (Required when member_type=VOLUME_GROUP) (Updatable) The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
		* `vault_id` - (Required when member_type=VOLUME_GROUP) (Updatable) The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
	* `connection_string_type` - (Applicable when member_type=AUTONOMOUS_CONTAINER_DATABASE) (Updatable) The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation. See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types. 
	* `db_system_admin_user_details` - (Applicable when member_type=MYSQL_DB_SYSTEM) (Updatable) The credentials for the HeatWave MySQL DB System administrator user, containing the username and the OCID of the Vault secret that stores the password.
		* `password_vault_secret_id` - (Required when member_type=MYSQL_DB_SYSTEM) (Updatable) The OCID of the vault secret where the HeatWave MySQL DB System password is stored.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
		* `username` - (Required when member_type=MYSQL_DB_SYSTEM) (Updatable) The user name for connecting to the HeatWave MySQL DB System node.  Example: `user` 
	* `db_system_replication_user_details` - (Applicable when member_type=MYSQL_DB_SYSTEM) (Updatable) The credentials for the HeatWave MySQL DB System replication user, containing the username and the OCID of the Vault secret that stores the password.
		* `password_vault_secret_id` - (Required when member_type=MYSQL_DB_SYSTEM) (Updatable) The OCID of the vault secret where the HeatWave MySQL DB System password is stored.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
		* `username` - (Required when member_type=MYSQL_DB_SYSTEM) (Updatable) The user name for connecting to the HeatWave MySQL DB System node.  Example: `user` 
	* `destination_availability_domain` - (Applicable when member_type=FILE_SYSTEM) (Updatable) The availability domain of the destination mount target.  Example: `BBTh:region-AD` 
	* `destination_backup_policy_id` - (Applicable when member_type=VOLUME_GROUP) (Updatable) The OCID of the backup policy to use in the destination region. This policy will be used to create backups  for this volume group after it moves the destination region.  Example: `ocid1.volumebackuppolicy.oc1..uniqueID` 
	* `destination_capacity_reservation_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of a capacity reservation in the destination region which will be used to launch the compute instance.  Example: `ocid1.capacityreservation.oc1..uniqueID` 
	* `destination_compartment_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE | FILE_SYSTEM | VOLUME_GROUP) (Updatable) The OCID of a compartment in the destination region in which the compute instance should be launched.  Example: `ocid1.compartment.oc1..uniqueID` 
	* `destination_dedicated_vm_host_id` - (Applicable when member_type=COMPUTE_INSTANCE | COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of a dedicated VM host in the destination region where the compute instance should be launched.  Example: `ocid1.dedicatedvmhost.oc1..uniqueID` 
	* `destination_encryption_key` - (Applicable when member_type=AUTONOMOUS_DATABASE | FILE_SYSTEM) (Updatable) Create properties for a customer-managed vault and encryption key in the destination region.  The customer-managed encryption key in this will be used to encrypt the resource or containing resources after they  move to the destination region. 
		* `encryption_key_id` - (Required when member_type=AUTONOMOUS_DATABASE | FILE_SYSTEM) (Updatable) The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
		* `vault_id` - (Required when member_type=AUTONOMOUS_DATABASE | FILE_SYSTEM) (Updatable) The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
	* `destination_load_balancer_id` - (Applicable when member_type=LOAD_BALANCER) (Updatable) The OCID of the destination load balancer.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `destination_network_load_balancer_id` - (Applicable when member_type=NETWORK_LOAD_BALANCER) (Updatable) The OCID of the destination network load balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `destination_snapshot_policy_id` - (Applicable when member_type=FILE_SYSTEM) (Updatable) The OCID of the snapshot policy to use in the destination region. This policy will be attached to the file system after it moves to the destination region.  Example: `ocid1.filesystemsnapshotpolicy.oc1..uniqueID` 
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
	* `gtid_reconciliation_timeout` - (Applicable when member_type=MYSQL_DB_SYSTEM) (Updatable) The maximum time (in seconds) to wait for the Global Transaction Identifier (GTID) synchronization process to complete before timing out.  Example: `600` 
	* `is_continue_on_gtid_reconciliation_timeout` - (Applicable when member_type=MYSQL_DB_SYSTEM) (Updatable) A flag indicating whether to continue with DR operation if the Global Transaction Identifier (GTID) reconciliation operation times out.  Example: `false` 
	* `is_movable` - (Applicable when member_type=COMPUTE_INSTANCE) (Updatable) A flag indicating if the compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) A flag indicating if the compute instance should be moved to the same fault domain in the destination region.  The compute instance launch will fail if this flag is set to true and capacity is not available in the  specified fault domain in the destination region.  Example: `false` 
	* `is_start_stop_enabled` - (Applicable when member_type=COMPUTE_INSTANCE_NON_MOVABLE) (Updatable) A flag indicating whether the non-movable compute instance should be started and stopped during DR operations. *Prechecks cannot be executed on stopped instances that are configured to be started.* 
	* `jump_host_id` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The OCID of the compute instance member that is designated as a jump host. This compute instance will be used to perform DR operations on the cluster using Oracle Cloud Agent's Run Command feature.  Example: `ocid1.instance.oc1..uniqueID` 
	* `load_balancer_mappings` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The list of source-to-destination load balancer mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_load_balancer_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the destination Load Balancer.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
		* `source_load_balancer_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the source Load Balancer. Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `managed_node_pool_configs` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The list of managed node pools with configurations for minimum and maximum node counts. This property applies to the OKE cluster member in both the primary and standby region. 
		* `id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the managed node pool in OKE cluster. 
		* `maximum` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The maximum number to which nodes in the managed node pool could be scaled up. 
		* `minimum` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The minimum number to which nodes in the managed node pool could be scaled down. 
	* `member_id` - (Required) (Updatable) The OCID of the member.  Example: `ocid1.instance.oc1..uniqueID` 
	* `member_type` - (Required) (Updatable) The type of the member. 
	* `namespace` - (Required when member_type=OBJECT_STORAGE_BUCKET) (Updatable) The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `network_load_balancer_mappings` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The list of source-to-destination network load balancer mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_network_load_balancer_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the Network Load Balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
		* `source_network_load_balancer_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the source Network Load Balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `password_vault_secret_id` - (Applicable when member_type=AUTONOMOUS_DATABASE | DATABASE) (Updatable) The OCID of the vault secret where the database SYSDBA password is stored. This password is required and used for performing database DR Drill operations when using full clone.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
	* `peer_cluster_id` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The OCID of the peer OKE cluster. This property applies to the OKE cluster member in both the primary and standby region.   Example: `ocid1.cluster.oc1..uniqueID` 
	* `peer_db_system_id` - (Applicable when member_type=MYSQL_DB_SYSTEM) (Updatable) The OCID of the peer HeatWave MySQL DB System from the peer region.  Example: `ocid1.mysqldbsystem.oc1..uniqueID` 
	* `source_volume_to_destination_encryption_key_mappings` - (Applicable when member_type=VOLUME_GROUP) (Updatable) A list of mappings between source volume IDs in the volume group and customer-managed encryption keys in the  destination region which will be used to encrypt the volume after it moves to the destination region.

		If you add the entry for source volumes and its corresponding vault and encryption keys here, you can not use  'commonDestinationKey' for encrypting all volumes with common encryption key. Similarly, if you specify common vault and encryption key using 'commonDestinationKey', you cannot specify vaults and encryption keys individually  for each volume using 'sourceVolumeToDestinationEncryptionKeyMappings'.

		An entry for each volume in volume group should be added in this list. The encryption key will not be updated  for the volumes that are part of volume group but missing in this list. 
		* `destination_encryption_key` - (Required when member_type=VOLUME_GROUP) (Updatable) Create properties for a customer-managed vault and encryption key in the destination region.  The customer-managed encryption key in this will be used to encrypt the resource or containing resources after they  move to the destination region. 
			* `encryption_key_id` - (Required when member_type=VOLUME_GROUP) (Updatable) The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
			* `vault_id` - (Required when member_type=VOLUME_GROUP) (Updatable) The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
		* `source_volume_id` - (Required when member_type=VOLUME_GROUP) (Updatable) The OCID of the source boot volume or block volume.  Example: `ocid1.volume.oc1..uniqueID` 
	* `vault_mappings` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The list of source-to-destination vault mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_vault_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the destination Vault.  Example: `ocid1.vault.oc1..uniqueID` 
		* `source_vault_id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the source Vault.  Example: `ocid1.vault.oc1..uniqueID` 
	* `virtual_node_pool_configs` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The list of virtual node pools with configurations for minimum and maximum node counts. This property applies to the OKE cluster member in both the primary and standby region. 
		* `id` - (Required when member_type=OKE_CLUSTER) (Updatable) The OCID of the virtual node pool in OKE cluster. 
		* `maximum` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The maximum number to which nodes in the virtual node pool could be scaled up. 
		* `minimum` - (Applicable when member_type=OKE_CLUSTER) (Updatable) The minimum number to which nodes in the virtual node pool could be scaled down. 
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
		* `destination_reserved_public_ip_id` - (Applicable when member_type=COMPUTE_INSTANCE_MOVABLE) (Updatable) The OCID of the reserved public IP address to be assigned to the compute instance in the destination region.  Example: `ocid1.publicip.oc1..uniqueID` 
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
	* `backup_config` - The details of backup performed on OKE Cluster. 
		* `backup_schedule` - The schedule for backing up namespaces to the destination region. If a backup schedule is not specified, only a single backup will be created. This format of the string specifying the backup schedule must conform with RFC-5545. This schedule will use the UTC timezone. This property applies to the OKE cluster member in primary region.  Example: FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;INTERVAL=1 
		* `exclude_namespaces` - A list of namespaces to be excluded from the backup.  The default value is null. If a list of namespaces to exclude is not provided, all namespaces will be backed up. Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both. This property applies to OKE cluster members in the primary region.  Example: ["namespace_string_3", "namespace_string_4"] 
		* `image_replication_vault_secret_id` - The OCID of the vault secret that stores the image credential. This property applies to the OKE cluster member in both the primary and standby region. 
		* `max_number_of_backups_retained` - The maximum number of backups that should be retained. This property applies to the OKE cluster member in primary region. 
		* `namespaces` - A list of namespaces to be included in the backup.  The default value is null. If a list of namespaces to include is not provided, all namespaces will be backed up. Specify either the `namespaces` or the `excludeNamespaces` parameter, but not both. This property applies to the OKE cluster member in primary region.  Example: ["default", "pv-nginx"] 
		* `replicate_images` - Controls the behaviour of image replication across regions. This property applies to the OKE cluster member in primary region. 
	* `backup_location` - The details for object storage backup location of an OKE Cluster
		* `bucket` - The bucket name inside the object storage namespace.  Example: `operation_logs` 
		* `namespace` - The namespace in object storage backup location(Note - this is usually the tenancy name).  Example: `myocitenancy` 
		* `object` - The object name inside the object storage bucket.  Example: `switchover_plan_executions` 
	* `block_volume_attach_and_mount_operations` - The details of the block volume operations performed on the non-movable compute instance. 
		* `attachments` - A list of details of attach or detach operations performed on block volumes. 
			* `block_volume_id` - The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
			* `volume_attachment_reference_instance_id` - The OCID of the reference compute instance needed to obtain the volume attachment details. This reference compute instance belongs to the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `mounts` - A list of details of mount operations performed on block volumes. 
			* `mount_point` - The physical mount point where the file system is mounted on the block volume.  Example: `/mnt/yourmountpoint` 
	* `block_volume_operations` - Deprecated. Use the 'blockVolumeAttachAndMountOperations' attribute instead of this. Operations performed on a list of block volumes used on the non-movable compute instance. 
		* `attachment_details` - Deprecated. Use the 'ComputeInstanceNonMovableBlockVolumeAttachOperationDetails' definition instead of this. The details for attaching or detaching a block volume. 
			* `volume_attachment_reference_instance_id` - The OCID of the reference compute instance needed to obtain the volume attachment details. This reference compute instance belongs to the peer DR protection group.  Example: `ocid1.instance.oc1..uniqueID` 
		* `block_volume_id` - The OCID of the block volume.  Example: `ocid1.volume.oc1..uniqueID` 
		* `mount_details` - Deprecated. Use the 'ComputeInstanceNonMovableBlockVolumeMountOperationDetails' definition instead of this. The details for mounting or unmounting the file system on a block volume. 
			* `mount_point` - The physical mount point used for mounting and unmounting the file system on a block volume.  Example: `/mnt/yourmountpoint` 
	* `bucket` - The bucket name inside the object storage namespace.  Example: `bucket_name` 
	* `common_destination_key` - The OCID of a vault and customer-managed encryption key in the destination region. 

		The customer-managed encryption key in this will be used to encrypt all the volumes of the volume group after they move to the destination region.  If you specify this common vault and encryption key, you cannot specify vaults and encryption keys individually for each volume  using 'sourceVolumeToDestinationEncryptionKeyMappings'.

		The customer-managed encryption key in this will be used to encrypt the file system when it move to the destination region. 
		* `encryption_key_id` - The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
		* `vault_id` - The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
	* `connection_string_type` - The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation. See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types. 
	* `db_system_admin_user_details` - The credentials for the HeatWave MySQL DB System administrator user, containing the username and the OCID of the vault secret that stores the password.
		* `password_vault_secret_id` - The OCID of the vault secret where the HeatWave MySQL DB System password is stored.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
		* `username` - The user name for connecting to the HeatWave MySQL DB System node.  Example: `user` 
	* `db_system_replication_user_details` - The credentials for the HeatWave MySQL DB System replication user, containing the username and the OCID of the vault secret that stores the password.
		* `password_vault_secret_id` - The OCID of the vault secret where the HeatWave MySQL DB System password is stored.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
		* `username` - The user name for connecting to the HeatWave MySQL DB System node.  Example: `user` 
	* `destination_availability_domain` - The availability domain of the destination mount target. Example: `BBTh:region-AD` 
	* `destination_backup_policy_id` - The OCID of the backup policy to use in the destination region. This policy will be used to create backups for this volume group after it moves the destination region.  Example: `ocid1.volumebackuppolicy.oc1..uniqueID` 
	* `destination_capacity_reservation_id` - The OCID of a capacity reservation in the destination region which will be used to launch the compute instance.  Example: `ocid1.capacityreservation.oc1..uniqueID` 
	* `destination_compartment_id` - The OCID of a compartment in the destination region in which the compute instance should be launched.  Example: `ocid1.compartment.oc1..uniqueID` 
	* `destination_dedicated_vm_host_id` - The OCID of a dedicated VM host in the destination region where the compute instance should be launched.  Example: `ocid1.dedicatedvmhost.oc1..uniqueID` 
	* `destination_encryption_key` - The OCID of a vault and customer-managed encryption key in the destination region. 

		The customer-managed encryption key in this will be used to encrypt all the volumes of the volume group after they move to the destination region.  If you specify this common vault and encryption key, you cannot specify vaults and encryption keys individually for each volume  using 'sourceVolumeToDestinationEncryptionKeyMappings'.

		The customer-managed encryption key in this will be used to encrypt the file system when it move to the destination region. 
		* `encryption_key_id` - The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
		* `vault_id` - The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
	* `destination_load_balancer_id` - The OCID of the destination load balancer. The backend sets in this destination load balancer are updated during DR.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `destination_network_load_balancer_id` - The OCID of the destination network load balancer. The backend sets in this destination network load balancer are updated during DR.                Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `destination_snapshot_policy_id` - The OCID of the snapshot policy to use in the destination region. This policy will be attached to the file system after it moves to the destination region.  Example: `ocid1.filesystemsnapshotpolicy.oc1..uniqueID` 
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
	* `gtid_reconciliation_timeout` - The maximum time (in seconds) to wait for the Global Transaction Identifier (GTID) synchronization process to complete before timing out.  Example: `600` 
	* `is_continue_on_gtid_reconciliation_timeout` - A flag indicating whether to continue with DR operation if the Global Transaction Identifier (GTID) reconciliation operation times out.  Example: `false` 
	* `is_movable` - A flag indicating if the compute instance should be moved during DR operations.  Example: `false` 
	* `is_retain_fault_domain` - A flag indicating if the compute instance should be moved to the same fault domain in the destination region.  The compute instance launch will fail if this flag is set to true and capacity is not available in the  specified fault domain in the destination region.  Example: `false` 
	* `is_start_stop_enabled` - A flag indicating whether the non-movable compute instance needs to be started and stopped during DR operations. 
	* `jump_host_id` - The OCID of the compute instance member that is designated as a jump host. This compute instance will be used to perform DR operations on the cluster using Oracle Cloud Agent's Run Command feature.  Example: `ocid1.instance.oc1..uniqueID` 
	* `load_balancer_mappings` - The list of source-to-destination load balancer mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_load_balancer_id` - The OCID of the destination Load Balancer.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
		* `source_load_balancer_id` - The OCID of the source Load Balancer.  Example: `ocid1.loadbalancer.oc1..uniqueID` 
	* `managed_node_pool_configs` - The list of node pools with configurations for minimum and maximum node counts. This property applies to the OKE cluster member in both the primary and standby region. 
		* `id` - The OCID of the managed node pool in OKE cluster. 
		* `maximum` - The maximum number to which nodes in the managed node pool could be scaled up. 
		* `minimum` - The minimum number to which nodes in the managed node pool could be scaled down. 
	* `member_id` - The OCID of the member.  Example: `ocid1.instance.oc1..uniqueID` 
	* `member_type` - The type of the member. 
	* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `network_load_balancer_mappings` - The list of source-to-destination network load balancer mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_network_load_balancer_id` - The OCID of the destination Network Load Balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
		* `source_network_load_balancer_id` - The OCID of the source Network Load Balancer.  Example: `ocid1.networkloadbalancer.oc1..uniqueID` 
	* `password_vault_secret_id` - The OCID of the vault secret where the database SYSDBA password is stored. This password is required and used for performing database DR Drill operations when using full clone.  Example: `ocid1.vaultsecret.oc1..uniqueID` 
	* `peer_cluster_id` - The OCID of the peer OKE cluster. This property applies to the OKE cluster member in both the primary and standby region.  Example: `ocid1.cluster.oc1.uniqueID` 
	* `peer_db_system_id` - The OCID of the peer HeatWave MySQL DB System from the peer region.  Example: `ocid1.mysqldbsystem.oc1..uniqueID` 
	* `source_volume_to_destination_encryption_key_mappings` - A list of mappings between source volume IDs in the volume group and customer-managed encryption keys in the  destination region which will be used to encrypt the volume after it moves to the destination region.

		If you add the entry for source volumes and its corresponding vault and encryption keys here, you can not use  'commonDestinationKey' for encrypting all volumes with common encryption key. Similarly, if you specify common vault and encryption key using 'commonDestinationKey', you cannot specify vaults and encryption keys individually  for each volume using 'sourceVolumeToDestinationEncryptionKeyMappings'.

		An entry for each volume in volume group should be added in this list. The encryption key will not be updated  for the volumes that are part of volume group but missing in this list. 
		* `destination_encryption_key` - The OCID of a vault and customer-managed encryption key in the destination region. 

			The customer-managed encryption key in this will be used to encrypt all the volumes of the volume group after they move to the destination region.  If you specify this common vault and encryption key, you cannot specify vaults and encryption keys individually for each volume  using 'sourceVolumeToDestinationEncryptionKeyMappings'.

			The customer-managed encryption key in this will be used to encrypt the file system when it move to the destination region. 
			* `encryption_key_id` - The OCID of the customer-managed encryption key in the destination region vault.  Example: `ocid1.key.oc1..uniqueID` 
			* `vault_id` - The OCID of the destination region vault for the customer-managed encryption key.  Example: `ocid1.vault.oc1..uniqueID` 
		* `source_volume_id` - The OCID of the source boot volume or block volume.  Example: `ocid1.volume.oc1..uniqueID` 
	* `vault_mappings` - The list of source-to-destination vault mappings required for DR operations. This property applies to the OKE cluster member in primary region. 
		* `destination_vault_id` - The OCID of the destination Vault.  Example: `ocid1.vault.oc1..uniqueID` 
		* `source_vault_id` - The OCID of the source Vault.  Example: `ocid1.vault.oc1..uniqueID` 
	* `virtual_node_pool_configs` - The list of node pools with configurations for minimum and maximum node counts. This property applies to the OKE cluster member in both the primary and standby region. 
		* `id` - The OCID of the virtual node pool in OKE cluster. 
		* `maximum` - The maximum number to which nodes in the virtual node pool could be scaled up. 
		* `minimum` - The minimum number to which nodes in the virtual node pool could be scaled down. 
	* `vnic_mapping` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID1, ocid1.networksecuritygroup.oc1..uniqueID2 ]` 
		* `destination_subnet_id` - The OCID of the destination subnet to which the source VNIC should connect.  Example: `ocid1.subnet.oc1..uniqueID` 
		* `source_vnic_id` - The OCID of the VNIC.  Example: `ocid1.vnic.oc1..uniqueID` 
	* `vnic_mappings` - A list of compute instance VNIC mappings. 
		* `destination_nsg_id_list` - A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to the source VNIC.  Example: `[ ocid1.networksecuritygroup.oc1..uniqueID, ocid1.networksecuritygroup.oc1..uniqueID ]` 
		* `destination_primary_private_ip_address` - The private IP address to be assigned as the VNIC's primary IP address in the destination subnet. This must be a valid IP address in the destination subnet and the IP address must be available.  Example: `10.0.3.3` 
		* `destination_primary_private_ip_hostname_label` - The hostname label to be assigned in the destination subnet for the primary private IP of the source VNIC. This label is the hostname portion of the private IP's fully qualified domain name (FQDN)  (for example, 'myhost1' in the FQDN 'myhost1.subnet123.vcn1.oraclevcn.com').  Example: `myhost1` 
		* `destination_reserved_public_ip_id` - The OCID of the reserved public IP address to be assigned to the compute instance in the destination region.  Example: `ocid1.publicip.oc1..uniqueID` 
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

