---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub_source_awrhubsources_management"
sidebar_current: "docs-oci-resource-opsi-awr_hub_source_awrhubsources_management"
description: |-
  Provides the Awr Hub Source Awrhubsources Management resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_awr_hub_source_awrhubsources_management
This resource provides the Awr Hub Source Awrhubsources Management resource in Oracle Cloud Infrastructure Opsi service.

Enables a Awr Hub source database in Operations Insights. This will resume the Awr data flow for the given Awr Hub source if it was stopped earlier.

## Example Usage

```hcl
resource "oci_opsi_awr_hub_source_awrhubsources_management" "test_awr_hub_source_awrhubsources_management" {
	#Required
	awr_hub_source_id = oci_opsi_awr_hub_source.test_awr_hub_source.id
	enable_awrhubsource = var.enable_awrhubsource
}
```

## Argument Reference

The following arguments are supported:

* `awr_hub_source_id` - (Required) Unique Awr Hub Source identifier
* `enable_awrhubsource` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Awr Hub Source Awrhubsources Management
	* `update` - (Defaults to 20 minutes), when updating the Awr Hub Source Awrhubsources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Awr Hub Source Awrhubsources Management
