---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_exadata_insight"
sidebar_current: "docs-oci-datasource-opsi-exadata_insight"
description: |-
  Provides details about a specific Exadata Insight in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_exadata_insight
This data source provides details about a specific Exadata Insight resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an Exadata insight.

## Example Usage

```hcl
data "oci_opsi_exadata_insight" "test_exadata_insight" {
	#Required
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
}
```

## Argument Reference

The following arguments are supported:

* `exadata_insight_id` - (Required) Unique Exadata insight identifier


## Attributes Reference

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

