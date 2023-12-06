---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_identity_provider_groups"
sidebar_current: "docs-oci-datasource-identity-identity_provider_groups"
description: |-
  Provides the list of Identity Provider Groups in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_identity_provider_groups
This data source provides the list of Identity Provider Groups in Oracle Cloud Infrastructure Identity service.

**Deprecated.** For more information, see [Deprecated IAM Service APIs](https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/deprecatediamapis.htm).

Lists the identity provider groups.


## Example Usage

```hcl
data "oci_identity_identity_provider_groups" "test_identity_provider_groups" {
	#Required
	identity_provider_id = oci_identity_identity_provider.test_identity_provider.id

	#Optional
	name = var.identity_provider_group_name
	state = var.identity_provider_group_state
}
```

## Argument Reference

The following arguments are supported:

* `identity_provider_id` - (Required) The OCID of the identity provider.
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `identity_provider_groups` - The list of identity_provider_groups.

### IdentityProviderGroup Reference

The following attributes are exported:

* `display_name` - Display name of the group
* `external_identifier` - Identifier of the group in the identity provider
* `id` - The OCID of the `IdentityProviderGroup`.
* `identity_provider_id` - The OCID of the `IdentityProvider` this group belongs to.
* `name` - Display name of the group
* `time_created` - Date and time the `IdentityProviderGroup` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_modified` - Date and time the `IdentityProviderGroup` was last modified, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

