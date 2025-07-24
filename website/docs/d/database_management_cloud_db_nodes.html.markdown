---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_nodes"
sidebar_current: "docs-oci-datasource-database_management-cloud_db_nodes"
description: |-
  Provides the list of Cloud Db Nodes in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_db_nodes
This data source provides the list of Cloud Db Nodes in Oracle Cloud Infrastructure Database Management service.

Lists the cloud DB nodes in the specified cloud DB system.

## Example Usage

```hcl
data "oci_database_management_cloud_db_nodes" "test_cloud_db_nodes" {

	#Optional
	cloud_db_system_id = oci_database_management_cloud_db_system.test_cloud_db_system.id
	compartment_id = var.compartment_id
	display_name = var.cloud_db_node_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `cloud_db_node_collection` - The list of cloud_db_node_collection.

### CloudDbNode Reference

The following attributes are exported:

* `additional_details` - The additional details of the cloud DB node defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `cloud_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the DB node is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud DB node.
* `cpu_core_count` - The number of CPU cores available on the DB node.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud DB node. The name does not have to be unique.
* `domain_name` - Name of the domain.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The host name for the DB node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
* `state` - The current lifecycle state of the cloud DB node.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud DB node was created.
* `time_updated` - The date and time the cloud DB node was last updated.

