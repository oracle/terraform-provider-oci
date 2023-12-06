---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_idp_group_mappings"
sidebar_current: "docs-oci-datasource-identity-idp_group_mappings"
description: |-
  Provides the list of Idp Group Mappings in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_idp_group_mappings
This data source provides the list of Idp Group Mappings in Oracle Cloud Infrastructure Identity service.

**Deprecated.** For more information, see [Deprecated IAM Service APIs](https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/deprecatediamapis.htm).

Lists the group mappings for the specified identity provider.


## Example Usage

```hcl
data "oci_identity_idp_group_mappings" "test_idp_group_mappings" {
	#Required
	identity_provider_id = oci_identity_identity_provider.test_identity_provider.id
}
```

## Argument Reference

The following arguments are supported:

* `identity_provider_id` - (Required) The OCID of the identity provider.


## Attributes Reference

The following attributes are exported:

* `idp_group_mappings` - The list of idp_group_mappings.

### IdpGroupMapping Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the `IdentityProvider`.
* `group_id` - The OCID of the IAM Service group that is mapped to the IdP group.
* `id` - The OCID of the `IdpGroupMapping`.
* `identity_provider_id` - The OCID of the `IdentityProvider` this mapping belongs to.
* `idp_group_name` - The name of the IdP group that is mapped to the IAM Service group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The mapping's current state.
* `time_created` - Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

