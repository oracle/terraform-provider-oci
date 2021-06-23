---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_tsig_key"
sidebar_current: "docs-oci-resource-dns-tsig_key"
description: |-
  Provides the Tsig Key resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_tsig_key
This resource provides the Tsig Key resource in Oracle Cloud Infrastructure DNS service.

Creates a new TSIG key in the specified compartment. There is no
`opc-retry-token` header since TSIG key names must be globally unique.


## Example Usage

```hcl
resource "oci_dns_tsig_key" "test_tsig_key" {
	#Required
	algorithm = var.tsig_key_algorithm
	compartment_id = var.compartment_id
	name = var.tsig_key_name
	secret = var.tsig_key_secret

	#Optional
	defined_tags = var.tsig_key_defined_tags
	freeform_tags = var.tsig_key_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `algorithm` - (Required) TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2). 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the TSIG key.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `name` - (Required) A globally unique domain name identifying the key for a given pair of hosts.
* `secret` - (Required) A base64 string encoding the binary shared secret.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `algorithm` - TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2). 
* `compartment_id` - The OCID of the compartment containing the TSIG key.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `name` - A globally unique domain name identifying the key for a given pair of hosts.
* `secret` - A base64 string encoding the binary shared secret.
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `time_created` - The date and time the resource was created, expressed in RFC 3339 timestamp format.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated, expressed in RFC 3339 timestamp format.

	**Example:** `2016-07-22T17:23:59:60Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Tsig Key
	* `update` - (Defaults to 20 minutes), when updating the Tsig Key
	* `delete` - (Defaults to 20 minutes), when destroying the Tsig Key


## Import

TsigKeys can be imported using the `id`, e.g.

```
$ terraform import oci_dns_tsig_key.test_tsig_key "id"
```

