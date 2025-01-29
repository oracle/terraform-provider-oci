---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_links"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-links"
description: |-
  Provides the list of Links in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_links
This data source provides the list of Links in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a (paginated) list of links.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_links" "test_links" {

	#Optional
	child_tenancy_id = oci_identity_tenancy.test_tenancy.id
	parent_tenancy_id = oci_identity_tenancy.test_tenancy.id
	state = var.link_state
}
```

## Argument Reference

The following arguments are supported:

* `child_tenancy_id` - (Optional) The ID of the child tenancy this link is associated with.
* `parent_tenancy_id` - (Optional) The ID of the parent tenancy this link is associated with.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `link_collection` - The list of link_collection.

### Link Reference

The following attributes are exported:

* `child_tenancy_id` - OCID of the child tenancy.
* `id` - OCID of the link.
* `parent_tenancy_id` - OCID of the parent tenancy.
* `state` - Lifecycle state of the link.
* `time_created` - Date-time when this link was created.
* `time_terminated` - Date-time when this link was terminated.
* `time_updated` - Date-time when this link was last updated.

