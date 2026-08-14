---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_organization_tenancy"
sidebar_current: "docs-oci-resource-tenantmanagercontrolplane-organization_tenancy"
description: |-
  Provides the Organization Tenancy resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# oci_tenantmanagercontrolplane_organization_tenancy
This resource provides the Organization Tenancy resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Creates a child tenancy asynchronously inside the organization. The newly created tenancy joins the organization
as a child (member) tenancy.

~> **NOTE on deletion:** Running `terraform destroy` only removes the child tenancy from Terraform state; it does
**not** reliably terminate the tenancy. Child tenancy termination is not a self-service operation through this
organization API in the general case (for example, the parent organization is not authorized to terminate a child
that is `OPTED_OUT` of governance). To actually terminate a child tenancy you must either:

* Terminate it **from within the child tenancy itself** (sign in to that tenancy as an administrator and use
  Governance & Administration → Tenancy Management → Terminate), or
* Open an **Oracle Support service request** to have the tenancy terminated (required when you no longer have
  administrative access to the child tenancy, e.g. it was created with an incorrect `admin_email`).

After the tenancy has been terminated through one of the above paths, remove it from Terraform state with
`terraform state rm oci_tenantmanagercontrolplane_organization_tenancy.<name>`.

Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/organizations/latest/OrganizationTenancy

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/tenantmanagercontrolplane


## Example Usage

```hcl
resource "oci_tenantmanagercontrolplane_organization_tenancy" "test_organization_tenancy" {
	#Required
	admin_email     = var.organization_tenancy_admin_email
	compartment_id  = var.tenancy_ocid
	home_region     = var.organization_tenancy_home_region
	organization_id = var.organization_id
	tenancy_name    = var.organization_tenancy_name

	#Optional
	governance_status = var.organization_tenancy_governance_status
	policy_name       = var.organization_tenancy_policy_name
	subscription_id   = oci_tenantmanagercontrolplane_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `admin_email` - (Required) Email address of the child tenancy administrator.
* `compartment_id` - (Required) The tenancy ID of the parent tenancy.
* `governance_status` - (Optional) The governance status of the child tenancy.
* `home_region` - (Required) The home region to use for the child tenancy. This must be a region where the parent tenancy is subscribed.
* `organization_id` - (Required) OCID of the organization the child tenancy belongs to. Used to read the tenancy after creation.
* `policy_name` - (Optional) The name to use for the administrator policy in the child tenancy. Must contain only letters and underscores.
* `subscription_id` - (Optional) OCID of the subscription that needs to be assigned to the child tenancy.
* `tenancy_name` - (Required) The tenancy name to use for the child tenancy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Organization Tenancy
	* `update` - (Defaults to 20 minutes), when updating the Organization Tenancy
	* `delete` - (Defaults to 20 minutes), when destroying the Organization Tenancy

Note on create: the operation is asynchronous and tracked via a work request. See the deletion note at the top of this
page — destroying the resource does not by itself terminate the child tenancy.


## Import

OrganizationTenancies can be imported using a composite id of the form
`organizations/{organizationId}/tenancies/{tenancyId}` (the organization OCID is required to read the tenancy and is
not encoded in the tenancy OCID alone), e.g.

```
$ terraform import oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy "organizations/{organizationId}/tenancies/{tenancyId}"
```

Note: `GetOrganizationTenancy` does not return the original create-only inputs `admin_email`, `home_region`, and
`compartment_id` (the parent tenancy). After import these will be empty in state while present in your configuration,
which would otherwise plan a destroy/recreate. Add a `lifecycle` block to keep the imported tenancy stable:

```hcl
lifecycle {
  ignore_changes = [admin_email, home_region, compartment_id]
}
```

(`tenancy_name` is repopulated from the returned tenancy name, so it does not need to be ignored.)

