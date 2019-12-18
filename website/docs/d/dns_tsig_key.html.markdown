---
subcategory: "Dns"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_tsig_key"
sidebar_current: "docs-oci-datasource-dns-tsig_key"
description: |-
  Provides details about a specific Tsig Key in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_tsig_key
This data source provides details about a specific Tsig Key resource in Oracle Cloud Infrastructure Dns service.

Gets information about the specified TSIG key.


## Example Usage

```hcl
data "oci_dns_tsig_key" "test_tsig_key" {
	#Required
	tsig_key_id = "${oci_dns_tsig_key.test_tsig_key.id}"
}
```

## Argument Reference

The following arguments are supported:

* `tsig_key_id` - (Required) The OCID of the target TSIG key.


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

