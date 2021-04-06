---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_distribution_statements"
sidebar_current: "docs-oci-datasource-core-drg_route_distribution_statements"
description: |-
  Provides the list of Drg Route Distribution Statements in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_route_distribution_statements
This data source provides the list of Drg Route Distribution Statements in Oracle Cloud Infrastructure Core service.

Lists the statements for the specified route distribution.

## Example Usage

```hcl
data "oci_core_drg_route_distribution_statements" "test_drg_route_distribution_statements" {
	#Required
	drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
}
```

## Argument Reference

The following arguments are supported:

* `drg_route_distribution_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route distribution.


## Attributes Reference

The following attributes are exported:

* `drg_route_distribution_statements` - The list of drg_route_distribution_statements.

### DrgRouteDistributionStatement Reference

The following attributes are exported:

* `action` - `ACCEPT` indicates the route should be imported or exported as-is. 
* `id` - The Oracle-assigned ID of the route distribution statement. 
* `match_criteria` - The action is applied only if all of the match criteria is met. If there are no match criteria in a statement, any input is considered a match and the action is applied. 
	* `attachment_type` - The type of the network resource to be included in this match. A match for a network type implies that all DRG attachments of that type insert routes into the table. 
	* `drg_attachment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG attachment. 
	* `match_type` - The type of the match criteria for a route distribution statement. 
* `priority` - This field specifies the priority of each statement in a route distribution. Priorities must be unique within a particular route distribution. The priority will be represented as a number between 0 and 65535 where a lower number indicates a higher priority. When a route is processed, statements are applied in the order defined by their priority. The first matching rule dictates the action that will be taken on the route. 

