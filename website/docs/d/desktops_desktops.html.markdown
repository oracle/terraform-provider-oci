---
subcategory: "Desktops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_desktops_desktops"
sidebar_current: "docs-oci-datasource-desktops-desktops"
description: |-
  Provides the list of Desktops in Oracle Cloud Infrastructure Desktops service
---
# Data Source: oci_desktops_desktops

This data source provides the list of Desktops in Oracle Cloud Infrastructure Desktops service.

Returns a list of desktops filtered by the specified parameters. You can limit the results to an availability domain, desktop name, desktop OCID, desktop state, pool OCID, or compartment OCID. You can limit the number of results returned, sort the results by time or name, and sort in ascending or descending order.

## Example Usage

```hcl
data "oci_desktops_desktops" "test_desktops" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.desktop_availability_domain
	desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
	display_name = var.desktop_display_name
	id = var.desktop_id
	state = var.desktop_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.
* `compartment_id` - (Required) The OCID of the compartment of the desktop pool.
* `desktop_pool_id` - (Optional) The OCID of the desktop pool.
* `display_name` - (Optional) A filter to return only results with the given displayName.
* `id` - (Optional) A filter to return only results with the given OCID.
* `state` - (Optional) A filter to return only results with the given lifecycleState.

## Attributes Reference

The following attributes are exported:

* `desktop_collection` - The list of desktop_collection.

### Desktop Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user friendly display name. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - The OCID of the desktop.
* `pool_id` - The OCID of the desktop pool the desktop is a member of.
* `state` - The state of the desktop.
* `time_created` - The date and time the resource was created.
* `user_name` - The owner of the desktop.
