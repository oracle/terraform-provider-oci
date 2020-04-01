---
subcategory: "Bds"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instances"
sidebar_current: "docs-oci-datasource-bds-bds_instances"
description: |-
  Provides the list of Bds Instances in Oracle Cloud Infrastructure Bds service
---

# Data Source: oci_bds_bds_instances
This data source provides the list of Bds Instances in Oracle Cloud Infrastructure Bds service.

Returns a list of BDS instances.


## Example Usage

```hcl
data "oci_bds_bds_instances" "test_bds_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.bds_instance_display_name}"
	state = "${var.bds_instance_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The state of the BDS instance.


## Attributes Reference

The following attributes are exported:

* `bds_instances` - The list of bds_instances.

### BdsInstance Reference

The following attributes are exported:

* `cloud_sql_details` - The information about added Cloud SQL capability
	* `block_volume_size_in_gbs` - The size of block volume in GB that needs to be attached to a given node. All the necessary details needed for attachment are managed by service itself. 
	* `ip_address` - IP address of the Cloud SQL node
	* `is_kerberos_mapped_to_database_users` - Boolean flag specifying whether or not are Kerberos principals mapped to database users. 
	* `kerberos_details` - Details about Kerberos principals
		* `keytab_file` - Location of the keytab file
		* `principal_name` - Name of the Kerberos principal
	* `shape` - Shape of the node
* `cluster_details` - Specific info about a Hadoop cluster
	* `bda_version` - BDA version installed in the cluster
	* `bdm_version` - BDM version installed in the cluster
	* `big_data_manager_url` - The URL of a Big Data Manager
	* `cloudera_manager_url` - The URL of a Cloudera Manager
	* `hue_server_url` - The URL of a Hue Server
	* `time_created` - The time the cluster was created. An RFC3339 formatted datetime string
	* `time_refreshed` - The time the BDS instance was automatically, or manually refreshed. An RFC3339 formatted datetime string 
* `cluster_version` - Version of the Hadoop distribution
* `compartment_id` - The OCID of the compartment
* `created_by` - The user who created the BDS instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the BDS instance
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the BDS resource
* `is_cloud_sql_configured` - Boolean flag specifying whether we configure Cloud SQL or not
* `is_high_availability` - Boolean flag specifying whether or not the cluster is HA
* `is_secure` - Boolean flag specifying whether or not the cluster should be setup as secure.
* `network_config` - Additional configuration of customer's network.
	* `cidr_block` - The CIDR IP address block of the VCN.
	* `is_nat_gateway_required` - A boolean flag whether to configure a NAT gateway.
* `nodes` - The list of nodes in the BDS instance
	* `attached_block_volumes` - The list of block volumes attached to a given node.
		* `volume_attachment_id` - The OCID of the volume attachment.
		* `volume_size_in_gbs` - The size of the volume in GBs.
	* `availability_domain` - The name of the availability domain the node is running in
	* `display_name` - The name of the node
	* `fault_domain` - The name of the fault domain the node is running in
	* `image_id` - The OCID of the image from which the node was created
	* `instance_id` - The OCID of the underlying compute instance
	* `ip_address` - IP address of the node
	* `node_type` - BDS instance node type
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

