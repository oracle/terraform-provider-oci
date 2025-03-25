---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance"
sidebar_current: "docs-oci-resource-bds-bds_instance"
description: |-
  Provides the Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance
This resource provides the Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service.

Creates a new BDS instance.


## Example Usage

```hcl
resource "oci_bds_bds_instance" "test_bds_instance" {
	#Required
	cluster_admin_password = var.bds_instance_cluster_admin_password
	cluster_public_key = var.bds_instance_cluster_public_key
	cluster_version = var.bds_instance_cluster_version
	compartment_id = var.compartment_id
	display_name = var.bds_instance_display_name
	is_high_availability = var.bds_instance_is_high_availability
	is_secure = var.bds_instance_is_secure
	master_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
        block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
            nvmes = var.bds_instance_nodes_shape_config_nvmes
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	util_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
        block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
            nvmes = var.bds_instance_nodes_shape_config_nvmes
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	worker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
        block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
            nvmes = var.bds_instance_nodes_shape_config_nvmes
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
	compute_only_worker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
        block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			nvmes = var.bds_instance_nodes_shape_config_nvmes
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}
    edge_node {
    	#Required
    	shape = var.bds_instance_nodes_shape
    	subnet_id = oci_core_subnet.test_subnet.id
    	number_of_nodes = var.bds_instance_number_of_nodes
    	#Optional
        block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    	shape_config {
    
    		#Optional
    		memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
    		nvmes = var.bds_instance_nodes_shape_config_nvmes
    		ocpus = var.bds_instance_nodes_shape_config_ocpus
    	}
    }
	kafka_broker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		number_of_nodes = var.bds_instance_number_of_nodes
		#Optional
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		shape_config {

			#Optional
			memory_in_gbs = var.bds_instance_nodes_shape_config_memory_in_gbs
			nvmes = var.bds_instance_nodes_shape_config_nvmes
			ocpus = var.bds_instance_nodes_shape_config_ocpus
		}
	}

	#Optional
	bds_cluster_version_summary {
		#Required
		bds_version = var.bds_instance_bds_cluster_version_summary_bds_version

		#Optional
		odh_version = var.bds_instance_bds_cluster_version_summary_odh_version
	}
	bootstrap_script_url = var.bds_instance_bootstrap_script_url
	cluster_profile = var.bds_instance_cluster_profile
	defined_tags = var.bds_instance_defined_tags
	freeform_tags = var.bds_instance_freeform_tags
	kerberos_realm_name = var.bds_instance_kerberos_realm_name
	kms_key_id = var.bds_instance_kms_key_id
	ignore_existing_nodes_shape = var.ignore_existing_nodes_shape
	network_config {

		#Optional
		cidr_block = var.bds_instance_network_config_cidr_block
		is_nat_gateway_required = var.bds_instance_network_config_is_nat_gateway_required
	}
}
```

## Argument Reference

The following arguments are supported:

* `bds_cluster_version_summary` - (Optional) Cluster version details including bds and odh version information.
	* `bds_version` - (Required) BDS version to be used for cluster creation
	* `odh_version` - (Optional) ODH version to be used for cluster creation
* `bootstrap_script_url` - (Optional) (Updatable) Pre-authenticated URL of the script in Object Store that is downloaded and executed.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
* `cluster_profile` - (Optional) Profile of the Big Data Service cluster.
* `cluster_public_key` - (Required) The SSH public key used to authenticate the cluster connection.
* `cluster_version` - (Required) Version of the Hadoop distribution
* `compartment_id` - (Required) (Updatable) The OCID of the compartment
* `display_name` - (Required) (Updatable) Name of the BDS instance
* `is_cloud_sql_configured` -(Optional) (Updatable) Boolean flag specifying whether we configure Cloud SQL or not
* `cloud_sql_details` -(Optional) The information about added Cloud SQL capability
    * `block_volume_size_in_gbs` - (Required) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
    * `shape` - (Required) Shape of the node
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_high_availability` - (Required) Boolean flag specifying whether or not the cluster is HA
* `is_secure` - (Required) Boolean flag specifying whether or not the cluster should be setup as secure.
* `kerberos_realm_name` - (Optional) The user-defined kerberos realm name.
* `kms_key_id` - (Optional) (Updatable) The OCID of the Key Management master encryption key.
* `network_config` - (Optional) (Updatable) Additional configuration of the user's network.
	* `cidr_block` - (Optional) (Updatable) The CIDR IP address block of the VCN.
	* `is_nat_gateway_required` - (Optional) (Updatable) A boolean flag whether to configure a NAT gateway.
* `nodes` - (Required) The list of nodes in the Big Data Service cluster.
	* `block_volume_size_in_gbs` - (Required) The size of block volume in GB to be attached to a given node. All the details needed for attaching the block volume are managed by service itself. 
	* `node_type` - (Required) The Big Data Service cluster node type.
	* `shape` - (Required) (Updatable) Shape of the node.
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes.
		* `nvmes` - (Optional) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
	* `subnet_id` - (Required) The OCID of the subnet in which the node will be created.
* `state` - (Optional) (Updatable) The target state for the Bds Instance. Could be set to `ACTIVE` or `INACTIVE`. 
* `execute_bootstrap_script_trigger` - (Optional) (Updatable) An optional property when incremented triggers Execute Bootstrap Script. Could be set to any integer value.
* `install_os_patch_trigger` - (Optional) (Updatable) An optional property when incremented triggers Install Os Patch. Could be set to any integer value.
* `remove_kafka_trigger` - (Optional) (Updatable) An optional property when incremented triggers Remove Kafka. Could be set to any integer value.
* `remove_node` - (Optional) (Updatable) An optional property when used triggers Remove Node. Takes the node ocid as input.
* `install_os_patch_trigger` - (Optional) (Updatable) An optional property when incremented triggers Install Os Patch. Could be set to any integer value.
* `state` - (Optional) (Updatable) The target state for the Bds Instance. Could be set to `ACTIVE` or `INACTIVE`.
* `remove_node` - (Optional) (Updatable) An optional property when used triggers Remove Node from an Active Cluster. Takes the node ocid as input
* `is_force_stop_jobs` - (Optional) (Updatable) When setting state as `INACTIVE` for stopping a cluster, setting this flag to true forcefully stops the bds instance.
* `is_kafka_configured` - (Optional) Boolean flag specifying whether or not Kafka should be configured.
* `os_patch_version`  - (Optional) (Updatable) The version of the patch to be upated.
* `state` - (Optional) (Updatable) The target state for the Bds Instance. Could be set to `ACTIVE` or `INACTIVE` to start/stop the bds instance.
* `is_force_stop_jobs` - (Optional) (Updatable) When setting state as `INACTIVE` for stopping a cluster, setting this flag to true forcefully stops the bds instance.
* `ignore_existing_nodes_shape` - Tag to ignore changing the shape of existing worker, master, utility, compute_only_worker, edge, kafka_broker nodes, in a list format, when new nodes are added with a different shape.
* `master_node` - (Required) The master node in the BDS instance
	* `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
	* `number_of_nodes` - (Required) The amount of master nodes should be created.
	* `shape` - (Required) Shape of the node
	* `subnet_id` - (Required) The OCID of the subnet in which the node should be created
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
* `util_node` - (Required) The utility node in the BDS instance
	* `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
	* `number_of_nodes` - (Required) The amount of utility nodes should be created.
	* `shape` - (Required) Shape of the node
	* `subnet_id` - (Required) The OCID of the subnet in which the node should be created
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
* `woker_node` - (Required) The worker node in the BDS instance
	* `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
	* `number_of_nodes` - (Required) The amount of worker nodes should be created, at least be 3.
	* `shape` - (Required) Shape of the node
	* `subnet_id` - (Required) The OCID of the subnet in which the node should be created
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
* `compute_only_woker_node` - (Optional) The worker node in the BDS instance
	* `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
	* `number_of_nodes` - (Required) The amount of worker nodes should be created
	* `shape` - (Required) Shape of the node
	* `subnet_id` - (Required) The OCID of the subnet in which the node should be created
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
* `kafka_broker_node` - (Optional) The kafka broker node in the BDS instance
	* `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
	* `number_of_nodes` - (Required) The amount of worker nodes should be created
	* `shape` - (Required) Shape of the node
	* `subnet_id` - (Required) The OCID of the subnet in which the node should be created
	* `shape_config` - (Optional) The shape configuration requested for the node.
		* `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
		* `ocpus` - (Optional) The total number of OCPUs available to the node.
    * `block_volume_size_in_gbs` - (Optional) The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself.
    * `number_of_nodes` - (Required) The amount of worker nodes should be created
    * `shape` - (Required) Shape of the node
    * `subnet_id` - (Required) The OCID of the subnet in which the node should be created
    * `shape_config` - (Optional) The shape configuration requested for the node.
        * `memory_in_gbs` - (Optional) The total amount of memory available to the node, in gigabytes
        * `ocpus` - (Optional) The total number of OCPUs available to the node.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_cluster_version_summary` - Cluster version details including bds and odh version information.
	* `bds_version` - BDS version to be used for cluster creation
	* `odh_version` - ODH version to be used for cluster creation
* `bootstrap_script_url` - pre-authenticated URL of the bootstrap script in Object Store that can be downloaded and executed.
* `cloud_sql_details` - The information about added Cloud SQL capability
	* `block_volume_size_in_gbs` - The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
	* `ip_address` - IP address of the Cloud SQL node
	* `is_kerberos_mapped_to_database_users` - Boolean flag specifying whether or not are Kerberos principals mapped to database users. 
	* `kerberos_details` - Details about Kerberos principals
		* `keytab_file` - Location of the keytab file
		* `principal_name` - Name of the Kerberos principal
	* `kms_key_id` - The OCID of the Key Management master encryption key
	* `shape` - Shape of the node
* `cluster_details` - Specific info about a Hadoop cluster
	* `ambari_url` - The URL of Ambari
	* `bd_cell_version` - Cloud SQL cell version.
	* `bda_version` - BDA version installed in the cluster
	* `bdm_version` - Big Data Manager version installed in the cluster.
	* `bds_version` - Big Data Service version installed in the cluster.
	* `big_data_manager_url` - The URL of Big Data Manager.
	* `cloudera_manager_url` - The URL of Cloudera Manager
	* `csql_cell_version` - Big Data SQL version.
	* `db_version` - Cloud SQL query server database version.
	* `hue_server_url` - The URL of the Hue server.
	* `jupyter_hub_url` - The URL of the Jupyterhub.
	* `odh_version` - Version of the ODH (Oracle Distribution including Apache Hadoop) installed on the cluster.
	* `os_version` - Oracle Linux version installed in the cluster.
	* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	* `time_refreshed` - The time the cluster was automatically or manually refreshed, shown as an RFC 3339 formatted datetime string. 
* `cluster_profile` - Profile of the Big Data Service cluster.
* `cluster_version` - Version of the Hadoop distribution.
* `compartment_id` - The OCID of the compartment.
* `created_by` - The user who created the cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - The name of the cluster.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 
* `id` - The OCID of the Big Data Service resource.
* `is_cloud_sql_configured` - Boolean flag specifying whether or not Cloud SQL should be configured.
* `is_high_availability` - Boolean flag specifying whether or not the cluster is highly available (HA)
* `is_kafka_configured` - Boolean flag specifying whether or not Kafka should be configured.
* `is_secure` - Boolean flag specifying whether or not the cluster should be set up as secure.
* `kms_key_id` - The OCID of the Key Management master encryption key.
* `network_config` - Additional configuration of the user's network.
	* `cidr_block` - The CIDR IP address block of the VCN.
	* `is_nat_gateway_required` - A boolean flag whether to configure a NAT gateway.
* `nodes` - The list of nodes in the cluster.
	* `attached_block_volumes` - The list of block volumes attached to a given node.
		* `volume_attachment_id` - The OCID of the volume attachment.
		* `volume_size_in_gbs` - The size of the volume in GBs.
	* `availability_domain` - The name of the availability domain in which the node is running.
	* `display_name` - The name of the node.
	* `fault_domain` - The name of the fault domain in which the node is running.
	* `hostname` - The fully-qualified hostname (FQDN) of the node.
	* `image_id` - The OCID of the image from which the node was created.
	* `instance_id` - The OCID of the underlying Oracle Cloud Infrastructure Compute instance.
	* `ip_address` - IP address of the node.
	* `is_reboot_required` - Indicates if the node requires a reboot to either reflect the latest os kernel or take actions for maintenance reboot.
	* `local_disks_total_size_in_gbs` - The aggregate size of all local disks, in gigabytes. If the instance does not have any local disks, this field is null.
	* `memory_in_gbs` - The total amount of memory available to the node, in gigabytes.
	* `node_type` - Cluster node type.
	* `nvmes` - The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
	* `ocpus` - The total number of OCPUs available to the node.
	* `odh_version` - Version of the ODH (Oracle Distribution including Apache Hadoop) for the node.
	* `os_version` - BDS-assigned Operating System version for the node.
	* `shape` - Shape of the node.
	* `ssh_fingerprint` - The fingerprint of the SSH key used for node access.
	* `state` - The state of the node.
	* `subnet_id` - The OCID of the subnet in which the node is to be created.
	* `time_created` - The time the node was created, shown as an RFC 3339 formatted datetime string.
	* `time_maintenance_reboot_due` - The date and time the instance is expected to be stopped / started, in the format defined by RFC3339.
	* `time_updated` - The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
* `number_of_nodes` - The number of nodes that form the cluster.
* `number_of_nodes_requiring_maintenance_reboot` - Number of nodes that require a maintenance reboot
* `state` - The state of the cluster.
* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
* `nodes` - The list of nodes in the BDS instance
    * `attached_block_volumes` - The list of block volumes attached to a given node.
        * `volume_attachment_id` - The OCID of the volume attachment.
        * `volume_size_in_gbs` - The size of the volume in GBs.
    * `availability_domain` - The name of the availability domain the node is running in
    * `display_name` - The name of the node
    * `fault_domain` - The name of the fault domain the node is running in
    * `hostname` - The fully-qualified hostname (FQDN) of the node
    * `image_id` - The OCID of the image from which the node was created
    * `instance_id` - The OCID of the underlying compute instance
    * `ip_address` - IP address of the node
    * `memory_in_gbs` - The total amount of memory available to the node, in gigabytes.
    * `node_type` - BDS instance node type
    * `ocpus` - The total number of OCPUs available to the node.
    * `shape` - Shape of the node
    * `ssh_fingerprint` - The fingerprint of the SSH key used for node access
    * `state` - The state of the node
    * `subnet_id` - The OCID of the subnet in which the node should be created
    * `time_created` - The time the node was created. An RFC3339 formatted datetime string
    * `time_updated` - The time the BDS instance was updated. An RFC3339 formatted datetime string
* `number_of_nodes` - Number of nodes that forming the cluster
* `state` - The state of the BDS instance
* `time_created` - The time the BDS instance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the BDS instance was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance


## Import

BdsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance.test_bds_instance "id"
```

