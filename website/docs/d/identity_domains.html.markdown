---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains"
sidebar_current: "docs-oci-datasource-identity-domains"
description: |-
  Provides the list of Domains in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_domains
This data source provides the list of Domains in Oracle Cloud Infrastructure Identity service.

(For tenancies that support identity domains) Lists all identity domains within a tenancy.


## Example Usage

```hcl
data "oci_identity_domains" "test_domains" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.domain_display_name
	home_region_url = var.domain_home_region_url
	is_hidden_on_login = var.domain_is_hidden_on_login
	license_type = var.domain_license_type
	name = var.domain_name
	state = var.domain_state
	type = var.domain_type
	url = var.domain_url
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `display_name` - (Optional) The mutable display name of the identity domain.
* `home_region_url` - (Optional) The region-specific identity domain URL.
* `is_hidden_on_login` - (Optional) Indicates whether or not the identity domain is visible at the sign-in screen.
* `license_type` - (Optional) The license type of the identity domain.
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `type` - (Optional) The identity domain type.
* `url` - (Optional) The region-agnostic identity domain URL.


## Attributes Reference

The following attributes are exported:

* `domains` - The list of domains.

### Domain Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the identity domain.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The identity domain description. You can have an empty description.
* `display_name` - The mutable display name of the identity domain.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_region` - The home region for the identity domain. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 
* `home_region_url` - Region-specific identity domain URL.
* `id` - The OCID of the identity domain.
* `is_hidden_on_login` - Indicates whether the identity domain is hidden on the sign-in screen or not. 
* `license_type` - The license type of the identity domain.
* `lifecycle_details` - Any additional details about the current state of the identity domain. 
* `replica_regions` - The regions where replicas of the identity domain exist.
	* `region` - A REPLICATION_ENABLED region, e.g. us-ashburn-1. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names. 
	* `regional_url` - Region-specific identity domain URL.
	* `state` - The IDCS-replicated region state. 
	* `url` - Region-agnostic identity domain URL.
* `state` - The current state. 
* `time_created` - Date and time the identity domain was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `type` - The type of the identity domain. 
* `url` - Region-agnostic identity domain URL.

