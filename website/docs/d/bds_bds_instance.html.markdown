---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance"
sidebar_current: "docs-oci-datasource-bds-bds_instance"
description: |-
  Provides details about a specific Bds Instance in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance
This data source provides details about a specific Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service.

Returns information about the Big Data Service cluster identified by the given ID.

## Example Usage

```hcl
data "oci_bds_bds_instance" "test_bds_instance" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `bootstrap_script_url` - pre-authenticated URL of the bootstrap script in Object Store that can be downloaded and executed.
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