---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_organization_tenancies"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-organization_tenancies"
description: |-
  Provides the list of Organization Tenancies in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_organization_tenancies
This data source provides the list of Organization Tenancies in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets a list of tenancies in the organization.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_organization_tenancies" "test_organization_tenancies" {
	#Required
	organization_id = oci_tenantmanagercontrolplane_organization.test_organization.id
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required) OCID of the organization.


## Attributes Reference

The following attributes are exported:

* `organization_tenancy_collection` - The list of organization_tenancy_collection.

### OrganizationTenancy Reference

The following attributes are exported:

* `governance_status` - The governance status of the tenancy.
* `is_approved_for_transfer` - Parameter to indicate the tenancy is approved for transfer to another organization.
* `name` - Name of the tenancy.
* `role` - Role of the organization tenancy.
* `state` - Lifecycle state of the organization tenancy.
* `tenancy_id` - OCID of the tenancy.
* `time_joined` - Date and time when the tenancy joined the organization.
* `time_left` - Date and time when the tenancy left the organization.

