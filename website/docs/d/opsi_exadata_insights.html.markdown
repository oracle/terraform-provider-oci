---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_exadata_insights"
sidebar_current: "docs-oci-datasource-opsi-exadata_insights"
description: |-
  Provides the list of Exadata Insights in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_exadata_insights
This data source provides the list of Exadata Insights in Oracle Cloud Infrastructure Opsi service.

Gets a list of Exadata insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
When both compartmentId and compartmentIdInSubtree are specified, a list of Exadata insights in that compartment and in all sub-compartments will be returned.


## Example Usage

```hcl
data "oci_opsi_exadata_insights" "test_exadata_insights" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.exadata_insight_compartment_id_in_subtree
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	exadata_type = var.exadata_insight_exadata_type
	id = var.exadata_insight_id
	state = var.exadata_insight_state
	status = var.exadata_insight_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments. 
* `enterprise_manager_bridge_id` - (Optional) Unique Enterprise Manager bridge identifier
* `exadata_type` - (Optional) Filter by one or more Exadata types. Possible value are DBMACHINE, EXACS, and EXACC. 
* `id` - (Optional) Optional list of Exadata insight resource [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `exadata_insight_summary_collection` - The list of exadata_insight_summary_collection.

### ExadataInsight Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the Exadata insight resource
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_display_name` - Enterprise Manager Entity Display Name
* `enterprise_manager_entity_identifier` - Enterprise Manager Entity Unique Identifier
* `enterprise_manager_entity_name` - Enterprise Manager Entity Name
* `enterprise_manager_entity_type` - Enterprise Manager Entity Type
* `enterprise_manager_identifier` - Enterprise Manager Unique Identifier
* `entity_source` - Source of the Exadata system.
* `exadata_display_name` - The user-friendly name for the Exadata system. The name does not have to be unique.
* `exadata_name` - The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
* `exadata_rack_type` - Exadata rack type.
* `exadata_type` - Operations Insights internal representation of the the Exadata system type.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Exadata insight identifier
* `is_auto_sync_enabled` - Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight. This should be always set true when using terraform, when it is set to false member associations may be missing from the application.
* `is_virtualized_exadata` - true if virtualization is used in the Exadata system
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the Exadata insight.
* `status` - Indicates the status of an Exadata insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the Exadata insight was updated. An RFC3339 formatted datetime string

