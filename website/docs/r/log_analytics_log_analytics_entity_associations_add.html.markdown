---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_associations_add"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_entity_associations_add"
description: |-
  Provides the Log Analytics Entity Associations Add resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_entity_associations_add
This resource provides the Log Analytics Entity Associations Add resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/LogAnalyticsEntity/AddEntityAssociations

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

Adds association between input source log analytics entity and one or more existing destination entities.

## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_entity_associations_add" "test_log_analytics_entity_associations_add" {
	#Required
	association_entities = var.log_analytics_entity_associations_add_association_entities
	log_analytics_entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
	namespace = var.log_analytics_entity_associations_add_namespace
}
```

## Argument Reference

The following arguments are supported:

* `association_entities` - (Required) Destination entities OCIDs with which associations are to be added. 
* `log_analytics_entity_id` - (Required) The Log analytics entity OCID. 
* `namespace` - (Required) The Log Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Entity Associations Add
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Entity Associations Add
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Entity Associations Add


## Import

LogAnalyticsEntityAssociationsAdd can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_log_analytics_entity_associations_add.test_log_analytics_entity_associations_add "id"
```

