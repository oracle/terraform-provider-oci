---
subcategory: "Identity Data Plane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_data_plane_generate_scoped_access_token"
sidebar_current: "docs-oci-resource-identity_data_plane-generate_scoped_access_token"
description: |-
  Provides the Generate Scoped Access Token resource in Oracle Cloud Infrastructure Identity Data Plane service
---

# oci_identity_data_plane_generate_scoped_access_token
This resource provides the Generate Scoped Access Token resource in Oracle Cloud Infrastructure Identity Data Plane service.

Based on the calling principal and the input payload, derive the claims and create a security token.


## Example Usage

```hcl
resource "oci_identity_data_plane_generate_scoped_access_token" "test_generate_scoped_access_token" {
	#Required
	public_key = var.generate_scoped_access_token_public_key
	scope = var.generate_scoped_access_token_scope
}
```

## Argument Reference

The following arguments are supported:

* `public_key` - (Required) A temporary public key, owned by the service. The service also owns the corresponding private key. This public key will by put inside the security token by the auth service after successful validation of the certificate. 
* `scope` - (Required) Scope definition for the scoped access token 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `token` - The security token, signed by auth service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Generate Scoped Access Token
	* `update` - (Defaults to 20 minutes), when updating the Generate Scoped Access Token
	* `delete` - (Defaults to 20 minutes), when destroying the Generate Scoped Access Token


## Import

GenerateScopedAccessToken can be imported using the `id`, e.g.

```
$ terraform import oci_identity_data_plane_generate_scoped_access_token.test_generate_scoped_access_token "id"
```

