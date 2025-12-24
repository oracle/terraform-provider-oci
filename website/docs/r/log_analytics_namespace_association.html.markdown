---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_association"
sidebar_current: "docs-oci-resource-log_analytics-namespace_association"
description: |-
  Provides the Namespace Association resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_association
This resource provides the Namespace Association resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/Association/Upsert

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

Creates or updates associations between sources and entities. The entity should belong to the specified input compartment.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_association" "test_namespace_association" {
  #Required
  compartment_id = var.compartment_id
  entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
  log_group_id = oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group.id
  namespace = var.namespace_association_namespace
  source_name = var.namespace_association_source_name

  #Optional
  is_from_republish = var.namespace_association_is_from_republish
  association_properties {
    #Required
    name = var.namespace_association_association_properties_name
    
    #Optional
    patterns {
      #Required
      id = var.namespace_association_association_properties_patterns_id
      value = var.namespace_association_association_properties_patterns_value
    }
    value = var.namespace_association_association_properties_value
  }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment ID
* `is_from_republish` - (Optional) isFromRepublish
* `association_properties` - (Optional) A list of association properties.
    * `name` - (Required) The name of the association property.
    * `patterns` - (Optional) A list of pattern level overrides for this property.
        * `id` - (Required) The pattern id.
        * `value` - (Required) (Updatable) The value of the property.
    * `value` - (Optional) The value of the association property.
* `entity_id` - (Required) The entity unique identifier.
* `log_group_id` - (Required) (Updatable) The log group unique identifier.
* `source_name` - (Required) The source name.
* `namespace` - (Required) The Logging Analytics namespace used for the request.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:
* `agent_id` - The agent unique identifier.
* `association_properties` - A list of association properties.
    * `name` - The name of the association property.
    * `patterns` - A list of pattern level overrides for this property.
        * `id` - The pattern id.
        * `value` - The value of the property.
    * `value` - The value of the association property.
* `agent_entity_name` - The agent unique identifier.
* `entity_id` - The entity unique identifier.
* `entity_name` - The entity name.
* `entity_type_display_name` - The entity type display name.
* `entity_type_name` - The entity type internal name.
* `failure_message` - The failure message.
* `host` - The host name.
* `log_group_id` - The log group unique identifier.
* `log_group_name` - The log group name.
* `log_group_compartment` - The log group compartment.
* `time_last_attempted` - The last attempt date.
* `retry_count` - The number of times the association will be attempted before failing.
* `source_name` - The source name.
* `source_display_name` - The source display name.
* `source_type_name` - The source type internal name.
* `life_cycle_state` - The lifecycle status. Valid values are ACCEPTED, IN_PROGRESS, SUCCEEDED or FAILED.


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Namespace Association
* `update` - (Defaults to 20 minutes), when updating the Namespace Association
* `delete` - (Defaults to 20 minutes), when destroying the Namespace Association


## Import

NamespaceAssociations can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_namespace_association.test_namespace_association "namespaces/{namespaceName}/associations/{compartmentId}/{entityId}/{sourceName}"
```

