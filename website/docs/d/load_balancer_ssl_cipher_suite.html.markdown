---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_ssl_cipher_suite"
sidebar_current: "docs-oci-datasource-load_balancer-ssl_cipher_suite"
description: |-
  Provides details about a specific Ssl Cipher Suite in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_ssl_cipher_suite
This data source provides details about a specific Ssl Cipher Suite resource in Oracle Cloud Infrastructure Load Balancer service.

Gets the specified SSL cipher suite's configuration information.

## Example Usage

```hcl
data "oci_load_balancer_ssl_cipher_suite" "test_ssl_cipher_suite" {
	#Required
	name = var.ssl_cipher_suite_name

	#Optional
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer. 
* `name` - (Required) The name of the SSL cipher suite to retrieve.

	example: `example_cipher_suite` 


## Attributes Reference

The following attributes are exported:

* `name` - A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

	**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
	* oci-default-ssl-cipher-suite-v1
	* oci-modern-ssl-cipher-suite-v1
	* oci-compatible-ssl-cipher-suite-v1
	* oci-wider-compatible-ssl-cipher-suite-v1
	* oci-customized-ssl-cipher-suite

	example: `example_cipher_suite` 

