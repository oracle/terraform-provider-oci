---
layout: "oci"
page_title: "OCI: oci_identity_identity_provider"
sidebar_current: "docs-oci-resource-identity_provider"
description: |-
Creates and manages an OCI IdentityProvider
---

# oci_identity_identity_provider
The `oci_identity_identity_provider` resource creates and manages an OCI IdentityProvider

Creates a new identity provider in your tenancy. For more information, see
[Identity Providers and Federation](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/federation.htm).

You must specify your tenancy's OCID as the compartment ID in the request object.
Remember that the tenancy is simply the root compartment. For information about
OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the `IdentityProvider`, which must be unique
across all `IdentityProvider` objects in your tenancy and cannot be changed.

You must also specify a *description* for the `IdentityProvider` (although
it can be an empty string). It does not have to be unique, and you can change
it anytime with
[UpdateIdentityProvider](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/IdentityProvider/UpdateIdentityProvider).

After you send your request, the new object's `lifecycleState` will temporarily
be CREATING. Before using the object, first make sure its `lifecycleState` has
changed to ACTIVE.


## Example Usage

```hcl
resource "oci_identity_identity_provider" "test_identity_provider" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.identity_provider_description}"
	metadata = "${var.identity_provider_metadata}"
	metadata_url = "${var.identity_provider_metadata_url}"
	name = "${var.identity_provider_name}"
	product_type = "${var.identity_provider_product_type}"
	protocol = "${var.identity_provider_protocol}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of your tenancy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `metadata` - (Required) (Updatable) The XML that contains the information required for federating. 
* `metadata_url` - (Required) (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating. 
* `name` - (Required) The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed. 
* `product_type` - (Required) The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS` 
* `protocol` - (Required) (Updatable) The protocol used for federation.  Example: `SAML2` 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the `IdentityProvider`.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the `IdentityProvider`.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed. This is the name federated users see when choosing which identity provider to use when signing in to the Oracle Cloud Infrastructure Console. 
* `product_type` - The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Allowed values are: - `ADFS` - `IDCS`  Example: `IDCS` 
* `protocol` - The protocol used for federation. Allowed value: `SAML2`.  Example: `SAML2` 
* `state` - The current state. After creating an `IdentityProvider`, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

IdentityProviders can be imported using the `id`, e.g.

```
$ terraform import oci_identity_identity_provider.test_identity_provider "id"
```
