---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_ssl_cipher_suites"
sidebar_current: "docs-oci-datasource-load_balancer-ssl_cipher_suites"
description: |-
  Provides the list of Ssl Cipher Suites in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_ssl_cipher_suites
This data source provides the list of Ssl Cipher Suites in Oracle Cloud Infrastructure Load Balancer service.

Lists all SSL cipher suites associated with the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_ssl_cipher_suites" "test_ssl_cipher_suites" {

	#Optional
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer. 


## Attributes Reference

The following attributes are exported:

* `ssl_cipher_suites` - The list of ssl_cipher_suites.

### SslCipherSuite Reference

The following attributes are exported:

* `name` - A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

	**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
	* oci-default-ssl-cipher-suite-v1
	* oci-modern-ssl-cipher-suite-v1
	* oci-compatible-ssl-cipher-suite-v1
	* oci-wider-compatible-ssl-cipher-suite-v1
	* oci-customized-ssl-cipher-suite

	example: `example_cipher_suite` 

