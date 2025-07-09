---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_replace_configuration"
sidebar_current: "docs-oci-datasource-bds-bds_instance_node_replace_configuration"
description: |-
  Provides details about a specific Bds Instance Node Replace Configuration in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_node_replace_configuration
This data source provides details about a specific Bds Instance Node Replace Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Returns details of the nodeReplaceConfiguration identified by the given ID.


## Example Usage

```hcl
data "oci_bds_bds_instance_node_replace_configuration" "test_bds_instance_node_replace_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	node_replace_configuration_id = oci_audit_configuration.test_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `node_replace_configuration_id` - (Required) Unique Oracle-assigned identifier of the  NodeReplaceConfiguration.


## Attributes Reference

The following attributes are exported:

* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `display_name` - A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `duration_in_minutes` - This value is the minimum period of time to wait for metric emission before triggering node replacement. The value is in minutes.
* `id` - The unique identifier for the NodeReplaceConfiguration.
* `level_type_details` - Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - Type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `node_host_name` - Host name of the node to create backup configuration.
	* `node_type` - Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created.
* `metric_type` - Type of compute instance health metric to use for node replacement
* `state` - The state of the NodeReplaceConfiguration.
* `time_created` - The time the NodeReplaceConfiguration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the NodeReplaceConfiguration was updated, shown as an RFC 3339 formatted datetime string. 

