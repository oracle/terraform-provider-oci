---
subcategory: "Desktops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_desktops_desktop_pool_volumes"
sidebar_current: "docs-oci-datasource-desktops-desktop_pool_volumes"
description: |-
  Provides the list of Desktop Pool Volumes in Oracle Cloud Infrastructure Desktops service
---

# Data Source: oci_desktops_desktop_pool_volumes
This data source provides the list of Desktop Pool Volumes in Oracle Cloud Infrastructure Desktops service.

Returns a list of volumes within the given desktop pool. You can limit the results to an availability domain, volume name, or volume state. You can limit the number of results returned, sort the results by time or name, and sort in ascending or descending order.


## Example Usage

```hcl
data "oci_desktops_desktop_pool_volumes" "test_desktop_pool_volumes" {
	#Required
	compartment_id = var.compartment_id
	desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id

	#Optional
	availability_domain = var.desktop_pool_volume_availability_domain
	display_name = var.desktop_pool_volume_display_name
	id = var.desktop_pool_volume_id
	state = var.desktop_pool_volume_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.
* `compartment_id` - (Required) The OCID of the compartment of the desktop pool.
* `desktop_pool_id` - (Required) The OCID of the desktop pool.
* `display_name` - (Optional) A filter to return only results with the given displayName.
* `id` - (Optional) A filter to return only results with the given OCID.
* `state` - (Optional) A filter to return only results with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `desktop_pool_volume_collection` - The list of desktop_pool_volume_collection.

### DesktopPoolVolume Reference

The following attributes are exported:

* `items` - A list of desktop pool volumes.
	* `availability_domain` - The availability domain of the desktop pool.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the desktop pool volume.
	* `name` - The name of the desktop pool volume.
	* `pool_id` - The OCID of the desktop pool to which this volume belongs.
	* `state` - The state of the desktop pool volume.
	* `user_name` - The owner of the desktop pool volume.

