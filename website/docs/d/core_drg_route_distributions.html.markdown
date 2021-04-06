---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_distributions"
sidebar_current: "docs-oci-datasource-core-drg_route_distributions"
description: |-
  Provides the list of Drg Route Distributions in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_route_distributions
This data source provides the list of Drg Route Distributions in Oracle Cloud Infrastructure Core service.

Lists the route distributions in the specified DRG.

To retrieve the statements in a distribution, use the
ListDrgRouteDistributionStatements operation.


## Example Usage

```hcl
data "oci_core_drg_route_distributions" "test_drg_route_distributions" {
	#Required
	drg_id = oci_core_drg.test_drg.id

	#Optional
	display_name = var.drg_route_distribution_display_name
	state = var.drg_route_distribution_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `state` - (Optional) A filter that only returns resources that match the specified lifecycle state. The value is case insensitive. 


## Attributes Reference

The following attributes are exported:

* `drg_route_distributions` - The list of drg_route_distributions.

### DrgRouteDistribution Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the route distribution.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `distribution_type` - Whether this distribution defines how routes get imported into route tables or exported through DRG attachments. 
* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG that contains this route distribution. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The route distribution's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `state` - The route distribution's current state.
* `time_created` - The date and time the route distribution was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

