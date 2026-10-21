---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_child_tenancy"
sidebar_current: "docs-oci-resource-tenantmanagercontrolplane-child_tenancy"
description: |-
  Provides the Child Tenancy resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# oci_tenantmanagercontrolplane_child_tenancy
This resource provides the Child Tenancy resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/organizations/latest/ChildTenancy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/tenantmanagercontrolplane

Creates a child tenancy asynchronously.

## Example Usage

```hcl
resource "oci_tenantmanagercontrolplane_child_tenancy" "test_child_tenancy" {
	#Required
	admin_email = var.child_tenancy_admin_email
	compartment_id = var.compartment_id
	home_region = var.child_tenancy_home_region
	tenancy_name = oci_identity_tenancy.test_tenancy.name

	#Optional
	features = var.child_tenancy_features
	governance_status = var.child_tenancy_governance_status
	policy_name = oci_identity_policy.test_policy.name
}
```

## Argument Reference

The following arguments are supported:

* `admin_email` - (Required) Email address of the child tenancy administrator.
* `compartment_id` - (Required) The tenancy ID of the parent tenancy.
* `features` - (Optional) List of features that the child and parent tenancies will have links for. Link with feature CORE will always be created.
* `governance_status` - (Optional) The governance status of the child tenancy.
* `home_region` - (Required) The home region to use for the child tenancy. This must be a region where the parent tenancy is subscribed.
* `policy_name` - (Optional) The name to use for the administrator policy in the child tenancy. Must contain only letters and underscores.
* `tenancy_name` - (Required) The tenancy name to use for the child tenancy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Child Tenancy
	* `update` - (Defaults to 20 minutes), when updating the Child Tenancy
	* `delete` - (Defaults to 20 minutes), when destroying the Child Tenancy


## Import

ChildTenancies can be imported using the `id`, e.g.

```
$ terraform import oci_tenantmanagercontrolplane_child_tenancy.test_child_tenancy "id"
```

