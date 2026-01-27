---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_snaps"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_snaps"
description: |-
  Provides the list of Managed Instance Snaps in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_snaps
This data source provides the list of Managed Instance Snaps in Oracle Cloud Infrastructure Os Management Hub service.

Retrieves a list of snaps for a managed instance. Filters may be applied to select a subset of snaps based on the filter criteria.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_snaps" "test_managed_instance_snaps" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	compartment_id = var.compartment_id
	name = var.managed_instance_snap_name
	name_contains = var.managed_instance_snap_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `name` - (Optional) The resource name. 
* `name_contains` - (Optional) A filter to return resources that may partially match the name given.


## Attributes Reference

The following attributes are exported:

* `snap_collection` - The list of snap_collection.

### ManagedInstanceSnap Reference

The following attributes are exported:

* `items` - The list of snaps.
	* `description` - The description of of snap.
	* `name` - The name of the snap.
	* `publisher` - The publisher of the snap.
	* `revision` - The revision number of the snap channel.
	* `store_url` - The snap's store url.
	* `time_refreshed` - The date and time of the snap's last refresh in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format.
	* `tracking` - The track this snap is following.
	* `version` - The version of the snap.

