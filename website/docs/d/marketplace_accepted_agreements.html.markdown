---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_accepted_agreements"
sidebar_current: "docs-oci-datasource-marketplace-accepted_agreements"
description: |-
  Provides the list of Accepted Agreements in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_accepted_agreements
This data source provides the list of Accepted Agreements in Oracle Cloud Infrastructure Marketplace service.

Lists the terms of use agreements that have been accepted in the specified compartment.
You can filter results by specifying query parameters.


## Example Usage

```hcl
data "oci_marketplace_accepted_agreements" "test_accepted_agreements" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	accepted_agreement_id = "${oci_marketplace_accepted_agreement.test_accepted_agreement.id}"
	display_name = "${var.accepted_agreement_display_name}"
	listing_id = "${oci_marketplace_listing.test_listing.id}"
	package_version = "${var.accepted_agreement_package_version}"
}
```

## Argument Reference

The following arguments are supported:

* `accepted_agreement_id` - (Optional) The unique identifier for the accepted terms of use agreement.
* `compartment_id` - (Required) The unique identifier for the compartment.
* `display_name` - (Optional) The display name of the resource.
* `listing_id` - (Optional) The unique identifier for the listing.
* `package_version` - (Optional) The version of the package. Package versions are unique within a listing.


## Attributes Reference

The following attributes are exported:

* `accepted_agreements` - The list of accepted_agreements.

### AcceptedAgreement Reference

The following attributes are exported:

* `agreement_id` - The unique identifier for the terms of use agreement itself.
* `compartment_id` - The unique identifier for the compartment where the agreement was accepted.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A display name for the accepted agreement.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The unique identifier for the acceptance of the agreement within a specific compartment.
* `listing_id` - The unique identifier for the listing associated with the agreement.
* `package_version` - The package version associated with the agreement.
* `time_accepted` - The time the agreement was accepted.

