---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_node"
sidebar_current: "docs-oci-resource-database_management-external_db_node"
description: |-
  Provides the External Db Node resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_db_node
This resource provides the External Db Node resource in Oracle Cloud Infrastructure Database Management service.

Updates the external DB node specified by `externalDbNodeId`.


## Example Usage

```hcl
resource "oci_database_management_external_db_node" "test_external_db_node" {
	#Required
	external_db_node_id = oci_database_management_external_db_node.test_external_db_node.id

	#Optional
	external_connector_id = oci_database_management_external_connector.test_external_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `external_connector_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_node_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database node.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the external DB node defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external DB node.
* `cpu_core_count` - The number of CPU cores available on the DB node.
* `display_name` - The user-friendly name for the external DB node. The name does not have to be unique.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the DB node is a part of.
* `host_name` - The host name for the DB node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
* `state` - The current lifecycle state of the external DB node.
* `time_created` - The date and time the external DB node was created.
* `time_updated` - The date and time the external DB node was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Db Node
	* `update` - (Defaults to 20 minutes), when updating the External Db Node
	* `delete` - (Defaults to 20 minutes), when destroying the External Db Node


## Import

ExternalDbNodes can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_db_node.test_external_db_node "id"
```

