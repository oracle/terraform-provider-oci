---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_database_insight"
sidebar_current: "docs-oci-resource-opsi-database_insight"
description: |-
  Provides the Database Insight resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_database_insight
This resource provides the Database Insight resource in Oracle Cloud Infrastructure Opsi service.

Create a Database Insight resource for a Enterprise Manager(EM) managed database in Operations Insights. The database will be enabled in Operations Insights. Database metric collection and analysis will be started. The Database Insight resource for Autonomous Database and Management Agent managed external Database needs to be created by Database service terraform provider. 


## Example Usage

```hcl
resource "oci_opsi_database_insight" "test_database_insight" {
	#Required
	compartment_id = var.compartment_id
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	enterprise_manager_entity_identifier = var.database_insight_enterprise_manager_entity_identifier
	enterprise_manager_identifier = var.database_insight_enterprise_manager_identifier
	entity_source = var.database_insight_entity_source

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of database
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - (Required) OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_identifier` - (Required) Enterprise Manager Entity Unique Identifier
* `enterprise_manager_identifier` - (Required) Enterprise Manager Unqiue Identifier
* `entity_source` - (Required) (Updatable) Source of the database entity. The supported type is "EM_MANAGED_EXTERNAL_DATABASE"
* `exadata_insight_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values. The resource destruction here is basically a soft delete. User cannot create resource using the same EM managed bridge OCID. If resource is in enabled state during destruction, the resource will be disabled automatically before performing delete operation.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the database
* `database_display_name` - Display name of database
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_name` - Name of database
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Insight
	* `update` - (Defaults to 20 minutes), when updating the Database Insight
	* `delete` - (Defaults to 20 minutes), when destroying the Database Insight


## Import

DatabaseInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_database_insight.test_database_insight "id"
```

