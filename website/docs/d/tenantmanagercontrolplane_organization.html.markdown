---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_organization"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-organization"
description: |-
  Provides details about a specific Organization in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_organization
This data source provides details about a specific Organization resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the organization.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_organization" "test_organization" {
	#Required
	organization_id = oci_tenantmanagercontrolplane_organization.test_organization.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required) OCID of the organization to retrieve.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment containing the organization. Always a tenancy OCID.
* `default_ucm_subscription_id` - OCID of the default Universal Credits Model subscription. Any tenancy joining the organization will automatically get assigned this subscription, if a subscription is not explictly assigned.
* `display_name` - A display name for the organization. Avoid entering confidential information.
* `id` - OCID of the organization.
* `parent_name` - The name of the tenancy that is the organization parent.
* `state` - Lifecycle state of the organization.
* `time_created` - Date and time when the organization was created.
* `time_updated` - Date and time when the organization was last updated.

