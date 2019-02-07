---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_certificate"
sidebar_current: "docs-oci-datasource-waas-certificate"
description: |-
  Provides details about a specific Certificate in Oracle Cloud Infrastructure Waas service
---

# Data Source: oci_waas_certificate
This data source provides details about a specific Certificate resource in Oracle Cloud Infrastructure Waas service.

Gets the details of an SSL certificate.

## Example Usage

```hcl
data "oci_waas_certificate" "test_certificate" {
	#Required
	certificate_id = "${oci_waas_certificate.test_certificate.id}"
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate used in the WAAS policy. This number is generated when the certificate is added to the policy.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate's compartment.
* `defined_tags` - A key-value pair with a defined schema that restricts the values of tags. These predefined keys are scoped to namespaces.
* `display_name` - The user-friendly name of the SSL certificate.
* `extensions` - 
	* `is_critical` - 
	* `name` - 
	* `value` - 
* `freeform_tags` - A simple key-value pair without any defined schema.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate.
* `issued_by` - 
* `issuer_name` - 
	* `common_name` - 
	* `country` - 
	* `email_address` - 
	* `locality` - 
	* `organization` - 
	* `organizational_unit` - 
	* `state_province` - 
* `public_key_info` - 
	* `algorithm` - 
	* `exponent` - 
	* `key_size` - 
* `serial_number` - 
* `signature_algorithm` - 
* `state` - The current lifecycle state of the SSL certificate.
* `subject_name` - 
	* `common_name` - 
	* `country` - 
	* `email_address` - 
	* `locality` - 
	* `organization` - 
	* `organizational_unit` - 
	* `state_province` - 
* `time_created` - The date and time the certificate was created, expressed in RFC 3339 timestamp format.
* `time_not_valid_after` - The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
* `time_not_valid_before` - 
* `version` - 

