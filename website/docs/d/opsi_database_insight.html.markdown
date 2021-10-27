---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_database_insight"
sidebar_current: "docs-oci-datasource-opsi-database_insight"
description: |-
  Provides details about a specific Database Insight in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_database_insight
This data source provides details about a specific Database Insight resource in Oracle Cloud Infrastructure Opsi service.

Gets details of a database insight.

## Example Usage

```hcl
data "oci_opsi_database_insight" "test_database_insight" {
	#Required
	database_insight_id = oci_opsi_database_insight.test_database_insight.id
}
```

## Argument Reference

The following arguments are supported:

* `database_insight_id` - (Required) Unique database insight identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the database
* `database_type` - Operations Insights internal representation of the database type.
* `database_version` - The version of the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_display_name` - Enterprise Manager Entity Display Name
* `enterprise_manager_entity_identifier` - Enterprise Manager Entity Unique Identifier
* `enterprise_manager_entity_name` - Enterprise Manager Entity Name
* `enterprise_manager_entity_type` - Enterprise Manager Entity Type
* `enterprise_manager_identifier` - Enterprise Manager Unqiue Identifier
* `entity_source` - Source of the database entity.
* `exadata_insight_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Database insight identifier
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `processor_count` - Processor count.
* `state` - The current state of the database.
* `status` - Indicates the status of a database insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the database insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the database insight was updated. An RFC3339 formatted datetime string

