---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_exadata_insight"
sidebar_current: "docs-oci-resource-opsi-exadata_insight"
description: |-
  Provides the Exadata Insight resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_exadata_insight
This resource provides the Exadata Insight resource in Oracle Cloud Infrastructure Opsi service.

Create an Exadata insight resource for an Exadata system in Operations Insights. The Exadata system will be enabled in Operations Insights. Exadata-related metric collection and analysis will be started.


## Example Usage

```hcl
resource "oci_opsi_exadata_insight" "test_exadata_insight" {
	#Required
	compartment_id = var.compartment_id
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	enterprise_manager_entity_identifier = var.exadata_insight_enterprise_manager_entity_identifier
	enterprise_manager_identifier = var.exadata_insight_enterprise_manager_identifier
	entity_source = var.exadata_insight_entity_source

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_auto_sync_enabled = var.exadata_insight_is_auto_sync_enabled
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of Exadata insight
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - (Required) OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_identifier` - (Required) Enterprise Manager Entity Unique Identifier
* `enterprise_manager_identifier` - (Required) Enterprise Manager Unique Identifier
* `entity_source` - (Required) (Updatable) Source of the Exadata system.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_auto_sync_enabled` - (Optional) (Updatable) Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `is_auto_sync_enabled` - Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
* `is_virtualized_exadata` - true if virtualization is used in the Exadata system
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the Exadata insight.
* `status` - Indicates the status of an Exadata insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the Exadata insight was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Exadata Insight
	* `update` - (Defaults to 20 minutes), when updating the Exadata Insight
	* `delete` - (Defaults to 20 minutes), when destroying the Exadata Insight


## Import

ExadataInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_exadata_insight.test_exadata_insight "id"
```

