---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_replace_configuration"
sidebar_current: "docs-oci-resource-bds-bds_instance_node_replace_configuration"
description: |-
  Provides the Bds Instance Node Replace Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_node_replace_configuration
This resource provides the Bds Instance Node Replace Configuration resource in Oracle Cloud Infrastructure Big Data Service service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/bigdata/latest/BdsInstanceNodeReplaceConfiguration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/big_data_service

Add a nodeReplaceConfigurations to the cluster.


## Example Usage

```hcl
resource "oci_bds_bds_instance_node_replace_configuration" "test_bds_instance_node_replace_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_node_replace_configuration_cluster_admin_password
	duration_in_minutes = var.bds_instance_node_replace_configuration_duration_in_minutes
	level_type_details {
		#Required
		level_type = var.bds_instance_node_replace_configuration_level_type_details_level_type

		#Optional
		node_host_name = var.bds_instance_node_replace_configuration_level_type_details_node_host_name
		node_type = var.bds_instance_node_replace_configuration_level_type_details_node_type
	}
	metric_type = var.bds_instance_node_replace_configuration_metric_type

	#Optional
	display_name = var.bds_instance_node_replace_configuration_display_name
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster admin user.
* `display_name` - (Optional) (Updatable) A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `duration_in_minutes` - (Required) (Updatable) This value is the minimum period of time to wait before triggering node replacement. The value is in minutes.
* `level_type_details` - (Required) (Updatable) Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - (Required) (Updatable) Type of level used to trigger the creation of a new node backup configuration or node replacement configuration. Accepted values are NODE_LEVEL and NODE_TYPE_LEVEL.
	* `node_host_name` - (Required when level_type=NODE_LEVEL) (Updatable) Host name of the node to create backup configuration.
	* `node_type` - (Required when level_type=NODE_TYPE_LEVEL) (Updatable) Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created.
* `metric_type` - (Required) (Updatable) Type of compute instance health metric to use for node replacement
* `remove_trigger` - (Optional) (Updatable) An optional property when incremented triggers Remove. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `display_name` - A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `duration_in_minutes` - This value is the minimum period of time to wait for metric emission before triggering node replacement. The value is in minutes.
* `id` - The unique identifier for the NodeReplaceConfiguration.
* `level_type_details` - Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - Type of level used to trigger the creation of a new node backup configuration or node replacement configuration. Accepted values are NODE_LEVEL and NODE_TYPE_LEVEL.
	* `node_host_name` - Host name of the node to create backup configuration.
	* `node_type` - Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created.
* `metric_type` - Type of compute instance health metric to use for node replacement
* `state` - The state of the NodeReplaceConfiguration.
* `time_created` - The time the NodeReplaceConfiguration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the NodeReplaceConfiguration was updated, shown as an RFC 3339 formatted datetime string. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Node Replace Configuration
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Node Replace Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Node Replace Configuration


## Import

BdsInstanceNodeReplaceConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration "bdsInstances/{bdsInstanceId}/nodeReplaceConfigurations/{nodeReplaceConfigurationId}" 
```

