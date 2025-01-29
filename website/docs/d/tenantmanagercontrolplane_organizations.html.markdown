---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_organizations"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-organizations"
description: |-
  Provides the list of Organizations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_organizations
This data source provides the list of Organizations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Lists organizations associated with the caller.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_organizations" "test_organizations" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `organization_collection` - The list of organization_collection.

### Organization Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment containing the organization. Always a tenancy OCID.
* `default_ucm_subscription_id` - OCID of the default Universal Credits Model subscription. Any tenancy joining the organization will automatically get assigned this subscription, if a subscription is not explictly assigned.
* `display_name` - A display name for the organization. Avoid entering confidential information.
* `id` - OCID of the organization.
* `parent_name` - The name of the tenancy that is the organization parent.
* `state` - Lifecycle state of the organization.
* `time_created` - Date and time when the organization was created.
* `time_updated` - Date and time when the organization was last updated.

