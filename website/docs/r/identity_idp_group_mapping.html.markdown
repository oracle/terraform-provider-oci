---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_idp_group_mapping"
sidebar_current: "docs-oci-resource-identity-idp_group_mapping"
description: |-
  Provides the Idp Group Mapping resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_idp_group_mapping
This resource provides the Idp Group Mapping resource in Oracle Cloud Infrastructure Identity service.

**Deprecated.** For more information, see [Deprecated IAM Service APIs](https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/deprecatediamapis.htm).

Creates a single mapping between an IdP group and an IAM Service
[group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/).


## Example Usage

```hcl
resource "oci_identity_idp_group_mapping" "test_idp_group_mapping" {
	#Required
	group_id = oci_identity_group.test_group.id
	identity_provider_id = oci_identity_identity_provider.test_identity_provider.id
	idp_group_name = var.idp_group_mapping_idp_group_name
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group. 
* `identity_provider_id` - (Required) The OCID of the identity provider.
* `idp_group_name` - (Required) (Updatable) The name of the IdP group you want to map.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the `IdentityProvider`.
* `group_id` - The OCID of the IAM Service group that is mapped to the IdP group.
* `id` - The OCID of the `IdpGroupMapping`.
* `identity_provider_id` - The OCID of the `IdentityProvider` this mapping belongs to.
* `idp_group_name` - The name of the IdP group that is mapped to the IAM Service group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The mapping's current state.
* `time_created` - Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Idp Group Mapping
	* `update` - (Defaults to 20 minutes), when updating the Idp Group Mapping
	* `delete` - (Defaults to 20 minutes), when destroying the Idp Group Mapping


## Import

IdpGroupMappings can be imported using the `id`, e.g.

```
$ terraform import oci_identity_idp_group_mapping.test_idp_group_mapping "identityProviders/{identityProviderId}/groupMappings/{mappingId}" 
```

