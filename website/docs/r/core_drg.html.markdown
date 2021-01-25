---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg"
sidebar_current: "docs-oci-resource-core-drg"
description: |-
  Provides the Drg resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg
This resource provides the Drg resource in Oracle Cloud Infrastructure Core service.

Creates a new dynamic routing gateway (DRG) in the specified compartment. For more information,
see [Dynamic Routing Gateways (DRGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingDRGs.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the DRG to reside. Notice that the DRG doesn't have to be in the same compartment as the VCN,
the DRG attachment, or other Networking Service components. If you're not sure which compartment
to use, put the DRG in the same compartment as the VCN. For more information about compartments
and access control, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the DRG, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.drg_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to contain the DRG.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The DRG's Oracle ID (OCID).
* `redundancy_status` - The redundancy status of the DRG specified.
* `state` - The DRG's current state.
* `time_created` - The date and time the DRG was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Import

Drgs can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg.test_drg "id"
```

