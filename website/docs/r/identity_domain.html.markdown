---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domain"
sidebar_current: "docs-oci-resource-identity-domain"
description: |-
  Provides the Domain resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_domain
This resource provides the Domain resource in Oracle Cloud Infrastructure Identity service.

Creates a new domain in the tenancy with domain home in {@code homeRegion}. This is an asynchronous call - where, at start,
{@code lifecycleState} of this domain is set to CREATING and {@code lifecycleDetails} to UPDATING. On domain creation completion
this Domain's {@code lifecycleState} will be set to ACTIVE and {@code lifecycleDetails} to null.

To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
the async operation's status.

After creating a `Domain`, make sure its `lifecycleState` changes from CREATING to ACTIVE
before using it.
If the domain's {@code displayName} already exists, returns 400 BAD REQUEST.
If any one of admin related fields are provided and one of the following 3 fields
- {@code adminEmail}, {@code adminLastName} and {@code adminUserName} - is not provided,
returns 400 BAD REQUEST.
- If {@code isNotificationBypassed} is NOT provided when admin information is provided,
returns 400 BAD REQUEST.
- If any internal error occurs, return 500 INTERNAL SERVER ERROR.


## Example Usage

```hcl
resource "oci_identity_domain" "test_domain" {
	#Required
	compartment_id = var.compartment_id
	description = var.domain_description
	display_name = var.domain_display_name
	home_region = var.domain_home_region
	license_type = var.domain_license_type

	#Optional
	admin_email = var.domain_admin_email
	admin_first_name = var.domain_admin_first_name
	admin_last_name = var.domain_admin_last_name
	admin_user_name = oci_identity_user.test_user.name
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_hidden_on_login = var.domain_is_hidden_on_login
	is_notification_bypassed = var.domain_is_notification_bypassed
	is_primary_email_required = var.domain_is_primary_email_required
}
```

## Argument Reference

The following arguments are supported:

* `admin_email` - (Optional) The admin email address
* `admin_first_name` - (Optional) The admin first name
* `admin_last_name` - (Optional) The admin last name
* `admin_user_name` - (Optional) The admin user name
* `compartment_id` - (Required) (Updatable) The OCID of the Compartment where domain is created
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) Domain entity description
* `display_name` - (Required) (Updatable) The mutable display name of the domain.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_region` - (Required) The region's name. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 
* `is_hidden_on_login` - (Optional) (Updatable) Indicates whether domain is hidden on login screen or not. 
* `is_notification_bypassed` - (Optional) Indicates if admin user created in IDCS stripe would like to receive notification like welcome email or not. Required field only if admin information is provided, otherwise optional. 
* `is_primary_email_required` - (Optional) Optional field to indicate whether users in the domain are required to have a primary email address or not Defaults to true 
* `license_type` - (Required) The License type of Domain


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the domain.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The domain descripition
* `display_name` - The mutable display name of the domain
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_region` - The home region for the domain. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 
* `home_region_url` - Region specific domain URL.
* `id` - The OCID of the domain
* `is_hidden_on_login` - Indicates whether domain is hidden on login screen or not. 
* `license_type` - The License type of Domain
* `lifecycle_details` - Any additional details about the current state of the Domain. 
* `replica_regions` - The regions domain is replication to.
	* `region` - A REPLICATION_ENABLED region, e.g. us-ashburn-1. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names. 
	* `state` - The IDCS replicated region state 
	* `url` - Region agnostic domain URL.
* `state` - The current state. 
* `time_created` - Date and time the domain was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `type` - The type of the domain. 
* `url` - Region agnostic domain URL.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Domain
	* `update` - (Defaults to 20 minutes), when updating the Domain
	* `delete` - (Defaults to 20 minutes), when destroying the Domain


## Import

Domains can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domain.test_domain "id"
```

