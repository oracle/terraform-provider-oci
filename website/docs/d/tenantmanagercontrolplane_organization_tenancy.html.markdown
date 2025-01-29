---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_organization_tenancy"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-organization_tenancy"
description: |-
  Provides details about a specific Organization Tenancy in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_organization_tenancy
This data source provides details about a specific Organization Tenancy resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the organization's tenancy.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_organization_tenancy" "test_organization_tenancy" {
	#Required
	organization_id = oci_tenantmanagercontrolplane_organization.test_organization.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required) OCID of the organization.
* `tenancy_id` - (Required) OCID of the tenancy to retrieve.


## Attributes Reference

The following attributes are exported:

* `governance_status` - The governance status of the tenancy.
* `is_approved_for_transfer` - Parameter to indicate the tenancy is approved for transfer to another organization.
* `name` - Name of the tenancy.
* `role` - Role of the organization tenancy.
* `state` - Lifecycle state of the organization tenancy.
* `tenancy_id` - OCID of the tenancy.
* `time_joined` - Date and time when the tenancy joined the organization.
* `time_left` - Date and time when the tenancy left the organization.

