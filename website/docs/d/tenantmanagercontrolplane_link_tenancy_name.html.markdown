---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_link_tenancy_name"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-link_tenancy_name"
description: |-
  Provides details about a specific Link Tenancy Name in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_link_tenancy_name
This data source provides details about a specific Link Tenancy Name resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the link along with the parent and child tenancy names.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_link_tenancy_name" "test_link_tenancy_name" {
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
* `child_tenancy_name` - Name of the child tenancy.
* `feature` - The feature associated with this link. Default value is CORE.
* `id` - OCID of the link.
* `parent_tenancy_id` - OCID of the parent tenancy.
* `parent_tenancy_name` - Name of the parent tenancy.
* `state` - Lifecycle state of the link.
* `time_created` - Date-time when this link was created.
* `time_terminated` - Date-time when this link was terminated.
* `time_updated` - Date-time when this link was last updated.
