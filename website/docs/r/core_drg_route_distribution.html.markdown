---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_distribution"
sidebar_current: "docs-oci-resource-core-drg_route_distribution"
description: |-
  Provides the Drg Route Distribution resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_route_distribution
This resource provides the Drg Route Distribution resource in Oracle Cloud Infrastructure Core service.

Creates a new route distribution for the specified DRG.
Assign the route distribution as an import distribution to a DRG route table using the `UpdateDrgRouteTable` or `CreateDrgRouteTable` operations.
Assign the route distribution as an export distribution to a DRG attachment
using the `UpdateDrgAttachment` or `CreateDrgAttachment` operations.


## Example Usage

```hcl
resource "oci_core_drg_route_distribution" "test_drg_route_distribution" {
	#Required
	distribution_type = var.drg_route_distribution_distribution_type
	drg_id = oci_core_drg.test_drg.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.drg_route_distribution_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `distribution_type` - (Required) Whether this distribution defines how routes get imported into route tables or exported through DRG attachments. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Route Distribution
	* `update` - (Defaults to 20 minutes), when updating the Drg Route Distribution
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Route Distribution


## Import

DrgRouteDistributions can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_route_distribution.test_drg_route_distribution "id"
```

