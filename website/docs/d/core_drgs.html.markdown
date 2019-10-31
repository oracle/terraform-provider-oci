---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drgs"
sidebar_current: "docs-oci-datasource-core-drgs"
description: |-
  Provides the list of Drgs in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drgs
This data source provides the list of Drgs in Oracle Cloud Infrastructure Core service.

Lists the DRGs in the specified compartment.


## Example Usage

```hcl
data "oci_core_drgs" "test_drgs" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `drgs` - The list of drgs.

### Drg Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state.
* `time_created` - The date and time the DRG was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

