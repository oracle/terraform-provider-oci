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

Creates a Big Data Service cluster.


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
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
	}
	util_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
	}
	worker_node {
		#Required
		shape = var.bds_instance_nodes_shape
		subnet_id = oci_core_subnet.test_subnet.id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes = var.bds_instance_number_of_nodes
	}

	#Optional
	defined_tags = var.bds_instance_defined_tags
	freeform_tags = var.bds_instance_freeform_tags
	network_config {

		#Optional
		cidr_block = var.bds_instance_network_config_cidr_block
		is_nat_gateway_required = var.bds_instance_network_config_is_nat_gateway_required
	}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
* `cluster_public_key` - (Required) The SSH public key used to authenticate the cluster connection.
* `cluster_version` - (Required) Version of the Hadoop distribution.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - (Required) (Updatable) Name of the Big Data Service cluster.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 
* `is_high_availability` - (Required) Boolean flag specifying whether or not the cluster is highly available (HA).
* `is_secure` - (Required) Boolean flag specifying whether or not the cluster should be set up as secure.
* `network_config` - (Optional) Additional configuration of the user's network.
	* `cidr_block` - (Optional) The CIDR IP address block of the VCN.
	* `is_nat_gateway_required` - (Optional) A boolean flag whether to configure a NAT gateway.
* `nodes` - (Required) The list of nodes in the Big Data Service cluster.
	* `block_volume_size_in_gbs` - (Required) The size of block volume in GB to be attached to a given node. All the details needed for attaching the block volume are managed by service itself. 
	* `node_type` - (Required) The Big Data Service cluster node type.
	* `shape` - (Required) (Updatable) Shape of the node.
	* `subnet_id` - (Required) The OCID of the subnet in which the node will be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cloud_sql_details` - The information about added Cloud SQL capability
	* `block_volume_size_in_gbs` - The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
	* `ip_address` - IP address of the Cloud SQL node.
	* `is_kerberos_mapped_to_database_users` - Boolean flag specifying whether or not Kerberos principals are mapped to database users. 
	* `kerberos_details` - Details about the Kerberos principals.
		* `keytab_file` - Location of the keytab file
		* `principal_name` - Name of the Kerberos principal.
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
	* `os_version` - Oracle Linux version installed in the cluster.
	* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	* `time_refreshed` - The time the cluster was automatically or manually refreshed, shown as an RFC 3339 formatted datetime string. 
* `cluster_version` - Version of the Hadoop distribution.
* `compartment_id` - The OCID of the compartment.
* `created_by` - The user who created the cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - The name of the cluster.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}` 
* `id` - The OCID of the Big Data Service resource.
* `is_cloud_sql_configured` - Boolean flag specifying whether or not Cloud SQL should be configured.
* `is_high_availability` - Boolean flag specifying whether or not the cluster is highly available (HA)
* `is_secure` - Boolean flag specifying whether or not the cluster should be set up as secure.
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
	* `node_type` - Cluster node type.
	* `shape` - Shape of the node.
	* `ssh_fingerprint` - The fingerprint of the SSH key used for node access.
	* `state` - The state of the node.
	* `subnet_id` - The OCID of the subnet in which the node is to be created.
	* `time_created` - The time the node was created, shown as an RFC 3339 formatted datetime string.
	* `time_updated` - The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
* `number_of_nodes` - The number of nodes that form the cluster.
* `state` - The state of the cluster.
* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the cluster was updated, shown as an RFC 3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance


## Import

BdsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance.test_bds_instance "id"
```

