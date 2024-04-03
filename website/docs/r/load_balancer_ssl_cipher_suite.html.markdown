---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_ssl_cipher_suite"
sidebar_current: "docs-oci-resource-load_balancer-ssl_cipher_suite"
description: |-
  Provides the Ssl Cipher Suite resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_ssl_cipher_suite
This resource provides the Ssl Cipher Suite resource in Oracle Cloud Infrastructure Load Balancer service.

Creates a custom SSL cipher suite.

## Example Usage

```hcl
resource "oci_load_balancer_ssl_cipher_suite" "test_ssl_cipher_suite" {
	#Required
	ciphers = var.ssl_cipher_suite_ciphers
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.ssl_cipher_suite_name
}
```

## Argument Reference

The following arguments are supported:

* `ciphers` - (Required)
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
* `name` - (Required) A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

	**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
	* oci-default-ssl-cipher-suite-v1
	* oci-modern-ssl-cipher-suite-v1
	* oci-compatible-ssl-cipher-suite-v1
	* oci-wider-compatible-ssl-cipher-suite-v1
	* oci-customized-ssl-cipher-suite
	* oci-default-http2-ssl-cipher-suite-v1
	* oci-default-http2-tls-13-ssl-cipher-suite-v1
	* oci-default-http2-tls-12-13-ssl-cipher-suite-v1
	* oci-tls-13-recommended-ssl-cipher-suite-v1
	* oci-tls-12-13-wider-ssl-cipher-suite-v1
	* oci-tls-11-12-13-wider-ssl-cipher-suite-v1

	example: `example_cipher_suite` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ciphers` - A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.

	The following ciphers are valid values for this property:
	*  __TLSv1.2 ciphers__

	"AES128-GCM-SHA256" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA256" "DH-DSS-AES128-GCM-SHA256" "DH-DSS-AES128-SHA256" "DH-DSS-AES256-GCM-SHA384" "DH-DSS-AES256-SHA256" "DH-RSA-AES128-GCM-SHA256" "DH-RSA-AES128-SHA256" "DH-RSA-AES256-GCM-SHA384" "DH-RSA-AES256-SHA256" "DHE-DSS-AES128-GCM-SHA256" "DHE-DSS-AES128-SHA256" "DHE-DSS-AES256-GCM-SHA384" "DHE-DSS-AES256-SHA256" "DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA256" "ECDH-ECDSA-AES128-GCM-SHA256" "ECDH-ECDSA-AES128-SHA256" "ECDH-ECDSA-AES256-GCM-SHA384" "ECDH-ECDSA-AES256-SHA384" "ECDH-RSA-AES128-GCM-SHA256" "ECDH-RSA-AES128-SHA256" "ECDH-RSA-AES256-GCM-SHA384" "ECDH-RSA-AES256-SHA384" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA384"
	*  __TLSv1 ciphers also supported by TLSv1.2__

	"AES128-SHA" "AES256-SHA" "CAMELLIA128-SHA" "CAMELLIA256-SHA" "DES-CBC3-SHA" "DH-DSS-AES128-SHA" "DH-DSS-AES256-SHA" "DH-DSS-CAMELLIA128-SHA" "DH-DSS-CAMELLIA256-SHA" "DH-DSS-DES-CBC3-SHAv" "DH-DSS-SEED-SHA" "DH-RSA-AES128-SHA" "DH-RSA-AES256-SHA" "DH-RSA-CAMELLIA128-SHA" "DH-RSA-CAMELLIA256-SHA" "DH-RSA-DES-CBC3-SHA" "DH-RSA-SEED-SHA" "DHE-DSS-AES128-SHA" "DHE-DSS-AES256-SHA" "DHE-DSS-CAMELLIA128-SHA" "DHE-DSS-CAMELLIA256-SHA" "DHE-DSS-DES-CBC3-SHA" "DHE-DSS-SEED-SHA" "DHE-RSA-AES128-SHA" "DHE-RSA-AES256-SHA" "DHE-RSA-CAMELLIA128-SHA" "DHE-RSA-CAMELLIA256-SHA" "DHE-RSA-DES-CBC3-SHA" "DHE-RSA-SEED-SHA" "ECDH-ECDSA-AES128-SHA" "ECDH-ECDSA-AES256-SHA" "ECDH-ECDSA-DES-CBC3-SHA" "ECDH-ECDSA-RC4-SHA" "ECDH-RSA-AES128-SHA" "ECDH-RSA-AES256-SHA" "ECDH-RSA-DES-CBC3-SHA" "ECDH-RSA-RC4-SHA" "ECDHE-ECDSA-AES128-SHA" "ECDHE-ECDSA-AES256-SHA" "ECDHE-ECDSA-DES-CBC3-SHA" "ECDHE-ECDSA-RC4-SHA" "ECDHE-RSA-AES128-SHA" "ECDHE-RSA-AES256-SHA" "ECDHE-RSA-DES-CBC3-SHA" "ECDHE-RSA-RC4-SHA" "IDEA-CBC-SHA" "KRB5-DES-CBC3-MD5" "KRB5-DES-CBC3-SHA" "KRB5-IDEA-CBC-MD5" "KRB5-IDEA-CBC-SHA" "KRB5-RC4-MD5" "KRB5-RC4-SHA" "PSK-3DES-EDE-CBC-SHA" "PSK-AES128-CBC-SHA" "PSK-AES256-CBC-SHA" "PSK-RC4-SHA" "RC4-MD5" "RC4-SHA" "SEED-SHA"

	example: `["ECDHE-RSA-AES256-GCM-SHA384","ECDHE-ECDSA-AES256-GCM-SHA384","ECDHE-RSA-AES128-GCM-SHA256"]` 
* `name` - A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

	**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
	* oci-default-ssl-cipher-suite-v1
	* oci-modern-ssl-cipher-suite-v1
	* oci-compatible-ssl-cipher-suite-v1
	* oci-wider-compatible-ssl-cipher-suite-v1
	* oci-customized-ssl-cipher-suite
	* oci-default-http2-ssl-cipher-suite-v1
	* oci-default-http2-tls-13-ssl-cipher-suite-v1
	* oci-default-http2-tls-12-13-ssl-cipher-suite-v1
	* oci-tls-13-recommended-ssl-cipher-suite-v1
	* oci-tls-12-13-wider-ssl-cipher-suite-v1
	* oci-tls-11-12-13-wider-ssl-cipher-suite-v1

	example: `example_cipher_suite` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ssl Cipher Suite
	* `update` - (Defaults to 20 minutes), when updating the Ssl Cipher Suite
	* `delete` - (Defaults to 20 minutes), when destroying the Ssl Cipher Suite


## Import

SslCipherSuites can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_ssl_cipher_suite.test_ssl_cipher_suite "loadBalancers/{loadBalancerId}/sslCipherSuites/{name}" 
```

