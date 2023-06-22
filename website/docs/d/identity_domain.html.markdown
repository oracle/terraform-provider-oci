---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domain"
sidebar_current: "docs-oci-datasource-identity-domain"
description: |-
  Provides details about a specific Domain in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_domain
This data source provides details about a specific Domain resource in Oracle Cloud Infrastructure Identity service.

Get the specified domain's information.

- If the domain doesn't exists, returns 404 NOT FOUND.
- If any internal error occurs, returns 500 INTERNAL SERVER ERROR.


## Example Usage

```hcl
data "oci_identity_domain" "test_domain" {
	#Required
	domain_id = oci_identity_domain.test_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required) The OCID of the domain


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

