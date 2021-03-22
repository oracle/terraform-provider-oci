---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_distribution_statement"
sidebar_current: "docs-oci-resource-core-drg_route_distribution_statement"
description: |-
  Provides the Drg Route Distribution Statement resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_route_distribution_statement
This resource provides the Drg Route Distribution Statement resource in Oracle Cloud Infrastructure Core service.

Adds one route distribution statement to the specified route distribution.


## Example Usage

```hcl
resource "oci_core_drg_route_distribution_statement" "test_drg_route_distribution_statement" {
	#Required
	drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
	action = var.drg_route_distribution_statement_statements_action
	#Optional
	match_criteria {
	#Required
	match_type = var.drg_route_distribution_statement_statements_match_criteria_match_type

	#Optional
	attachment_type = var.drg_route_distribution_statement_statements_match_criteria_attachment_type
	drg_attachment_id = oci_core_drg_attachment.test_drg_attachment.id
	}
	priority = var.drg_route_distribution_statement_statements_priority

}
```

## Argument Reference

The following arguments are supported:

* `drg_route_distribution_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route distribution.
* `action` - (Required) Accept: import/export the route "as is" 
* `match_criteria` - (Required) (Updatable) The action is applied only if all of the match criteria is met. If there are no match criteria in a statement, match ALL is implied. 
		* `attachment_type` - (Required when match_type=DRG_ATTACHMENT_TYPE) The type of the network resource to be included in this match. A match for a network type implies that all DRG attachments of that type insert routes into the table. 
		* `drg_attachment_id` - (Required when match_type=DRG_ATTACHMENT_ID) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG attachment. 
		* `match_type` - (Required) (Updatable) The type of the match criteria for a route distribution statement.
* `priority` - (Required) (Updatable) This field is used to specify the priority of each statement in a route distribution. The priority will be represented as a number between 0 and 65535 where a lower number indicates a higher priority. When a route is processed, statements are applied in the order defined by their priority. The first matching rule dictates the action that will be taken on the route. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - `ACCEPT` indicates the route should be imported or exported as-is. 
* `id` - The Oracle-assigned ID of the route distribution statement. 
* `match_criteria` - The action is applied only if all of the match criteria is met. If there are no match criteria in a statement, any input is considered a match and the action is applied. 
	* `attachment_type` - The type of the network resource to be included in this match. A match for a network type implies that all DRG attachments of that type insert routes into the table. 
	* `drg_attachment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG attachment. 
	* `match_type` - The type of the match criteria for a route distribution statement. 
* `priority` - This field specifies the priority of each statement in a route distribution. Priorities must be unique within a particular route distribution. The priority will be represented as a number between 0 and 65535 where a lower number indicates a higher priority. When a route is processed, statements are applied in the order defined by their priority. The first matching rule dictates the action that will be taken on the route. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Route Distribution Statement
	* `update` - (Defaults to 20 minutes), when updating the Drg Route Distribution Statement
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Route Distribution Statement


## Import

DrgRouteDistributionStatement can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_route_distribution_statement.test_drg_route_distribution_statement "drgRouteDistributions/{drgRouteDistributionId}/statements/{id}" 
```

