---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_accepted_agreement"
sidebar_current: "docs-oci-datasource-marketplace-accepted_agreement"
description: |-
  Provides details about a specific Accepted Agreement in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_accepted_agreement
This data source provides details about a specific Accepted Agreement resource in Oracle Cloud Infrastructure Marketplace service.

Gets the details of a specific, previously accepted terms of use agreement.


## Example Usage

```hcl
data "oci_marketplace_accepted_agreement" "test_accepted_agreement" {
	#Required
	accepted_agreement_id = "${oci_marketplace_accepted_agreement.test_accepted_agreement.id}"
}
```

## Argument Reference

The following arguments are supported:

* `accepted_agreement_id` - (Required) The unique identifier for the accepted terms of use agreement.


## Attributes Reference

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

