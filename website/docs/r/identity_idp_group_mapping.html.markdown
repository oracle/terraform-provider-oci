---
layout: "oci"
page_title: "OCI: oci_identity_idp_group_mapping"
sidebar_current: "docs-oci-resource-identity-idp_group_mapping"
description: |-
  Creates and manages an OCI IdpGroupMapping
---

# oci_identity_idp_group_mapping
The `oci_identity_idp_group_mapping` resource creates and manages an OCI IdpGroupMapping

Creates a single mapping between an IdP group and an IAM Service
[group](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Group/).


## Example Usage

```hcl
resource "oci_identity_idp_group_mapping" "test_idp_group_mapping" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"
	idp_group_name = "${var.idp_group_mapping_idp_group_name}"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) (Updatable) The OCID of the IAM Service [group](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Group/) you want to map to the IdP group. 
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
* `state` - The mapping's current state.  After creating a mapping object, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
