---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_marketplace_metadata_public_keys"
sidebar_current: "docs-oci-datasource-marketplace-marketplace_metadata_public_keys"
description: |-
  Provides the list of Marketplace Metadata Public Keys in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_marketplace_metadata_public_keys
This data source provides the list of Marketplace Metadata Public Keys in Oracle Cloud Infrastructure Marketplace service.

Get public certificates used in JWT signing, in JSON Web Key Sets format

## Example Usage

```hcl
data "oci_marketplace_marketplace_metadata_public_keys" "test_marketplace_metadata_public_keys" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.


## Attributes Reference

The following attributes are exported:

* `marketplace_metadata_public_keys` - The list of marketplace_metadata_public_keys.

### MarketplaceMetadataPublicKey Reference

The following attributes are exported:

* `certificate_chain` - chain of certificates used to sign JWT
* `certificate_thumbprint` - unique identifier of associated X509 certificate
* `exponent` - base64 encoded exponent for public key
* `key_algorithm` - algorithm for public key (i.e. RS256)
* `key_id` - unique id that maps to public certificate, directs user which certificate to use to verfiy
* `key_type` - key type (i.e. RSA)
* `key_use` - how key is to be used
* `modulus` - RSA public modulus

