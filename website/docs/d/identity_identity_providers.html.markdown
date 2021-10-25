---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_identity_providers"
sidebar_current: "docs-oci-datasource-identity-identity_providers"
description: |-
  Provides the list of Identity Providers in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_identity_providers
This data source provides the list of Identity Providers in Oracle Cloud Infrastructure Identity service.

**Deprecated.** For more information, see [Deprecated IAM Service APIs](https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/deprecatediamapis.htm).

Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_identity_providers" "test_identity_providers" {
	#Required
	compartment_id = var.tenancy_ocid
	protocol = var.identity_provider_protocol

	#Optional
	name = var.identity_provider_name
	state = var.identity_provider_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `protocol` - (Required) The protocol used for federation.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `identity_providers` - The list of identity_providers.

### IdentityProvider Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the `IdentityProvider`.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable. 
* `freeform_attributes` - Extra name value pairs associated with this identity provider. Example: `{"clientId": "app_sf3kdjf3"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the `IdentityProvider`.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `metadata` - The XML that contains the information required for federating Identity with SAML2 Identity Provider. 
* `metadata_url` - The URL for retrieving the identity provider's metadata, which contains information required for federating. 
* `name` - The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed. This is the name federated users see when choosing which identity provider to use when signing in to the Oracle Cloud Infrastructure Console. 
* `product_type` - The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).

	Allowed values are:
	* `ADFS`
	* `IDCS`

	Example: `IDCS` 
* `protocol` - The protocol used for federation. Allowed value: `SAML2`.  Example: `SAML2` 
* `redirect_url` - The URL to redirect federated users to for authentication with the identity provider. 
* `signing_certificate` - The identity provider's signing certificate used by the IAM Service to validate the SAML2 token. 
* `state` - The current state.
* `time_created` - Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

