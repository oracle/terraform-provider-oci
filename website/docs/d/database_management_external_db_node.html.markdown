---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_node"
sidebar_current: "docs-oci-datasource-database_management-external_db_node"
description: |-
Provides details about a specific External Db Node in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_node
This data source provides details about a specific External Db Node resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the external DB node specified by `externalDbNodeId`.


## Example Usage

```hcl
data "oci_database_management_external_db_node" "test_external_db_node" {
  #Required
  external_db_node_id = oci_database_management_external_db_node.test_external_db_node.id
}
```

## Argument Reference

The following arguments are supported:

* `external_db_node_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database node.


## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the external DB node defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}`
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external DB node.
* `cpu_core_count` - The number of CPU cores available on the DB node.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the external DB node. The name does not have to be unique.
* `domain_name` - Name of the domain.
* `external_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external connector.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the DB node is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The host name for the DB node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
* `state` - The current lifecycle state of the external DB node.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the external DB node was created.
* `time_updated` - The date and time the external DB node was last updated.
