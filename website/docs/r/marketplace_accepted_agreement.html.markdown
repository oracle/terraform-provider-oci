---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_accepted_agreement"
sidebar_current: "docs-oci-resource-marketplace-accepted_agreement"
description: |-
  Provides the Accepted Agreement resource in Oracle Cloud Infrastructure Marketplace service
---

# oci_marketplace_accepted_agreement
This resource provides the Accepted Agreement resource in Oracle Cloud Infrastructure Marketplace service.

Accepts a terms of use agreement for a specific package version of a listing. You must accept all
terms of use for a package before you can deploy the package.


## Example Usage

```hcl
resource "oci_marketplace_accepted_agreement" "test_accepted_agreement" {
	#Required
	agreement_id = oci_marketplace_agreement.test_agreement.id
	compartment_id = var.compartment_id
	listing_id = oci_marketplace_listing.test_listing.id
	package_version = var.accepted_agreement_package_version
	signature = var.accepted_agreement_signature

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.accepted_agreement_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `agreement_id` - (Required) The agreement to accept.
* `compartment_id` - (Required) The unique identifier for the compartment where the agreement will be accepted.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A display name for the accepted agreement.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `listing_id` - (Required) The unique identifier for the listing associated with the agreement.
* `package_version` - (Required) The package version associated with the agreement.
* `signature` - (Required) A signature generated for the listing package agreements that you can retrieve with [GetAgreement](https://docs.cloud.oracle.com/iaas/api/#/en/marketplace/20181001/Agreement/GetAgreement). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agreement_id` - The unique identifier for the terms of use agreement itself.
* `compartment_id` - The unique identifier for the compartment where the agreement was accepted.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A display name for the accepted agreement.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The unique identifier for the acceptance of the agreement within a specific compartment.
* `listing_id` - The unique identifier for the listing associated with the agreement.
* `package_version` - The package version associated with the agreement.
* `time_accepted` - The time the agreement was accepted.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Accepted Agreement
	* `update` - (Defaults to 20 minutes), when updating the Accepted Agreement
	* `delete` - (Defaults to 20 minutes), when destroying the Accepted Agreement


## Import

AcceptedAgreements can be imported using the `id`, e.g.

```
$ terraform import oci_marketplace_accepted_agreement.test_accepted_agreement "id"
```

