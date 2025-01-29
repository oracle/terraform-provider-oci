---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_link"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-link"
description: |-
  Provides details about a specific Link in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_link
This data source provides details about a specific Link resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the link.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_link" "test_link" {
	#Required
	link_id = oci_tenantmanagercontrolplane_link.test_link.id
}
```

## Argument Reference

The following arguments are supported:

* `link_id` - (Required) OCID of the link to retrieve.


## Attributes Reference

The following attributes are exported:

* `child_tenancy_id` - OCID of the child tenancy.
* `id` - OCID of the link.
* `parent_tenancy_id` - OCID of the parent tenancy.
* `state` - Lifecycle state of the link.
* `time_created` - Date-time when this link was created.
* `time_terminated` - Date-time when this link was terminated.
* `time_updated` - Date-time when this link was last updated.

