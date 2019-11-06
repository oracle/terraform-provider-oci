---
subcategory: "Waas"
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
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name of the SSL certificate.
* `extensions` - Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
	* `is_critical` - The critical flag of the extension. Critical extensions must be processed, non-critical extensions can be ignored.
	* `name` - The certificate extension name.
	* `value` - The certificate extension value.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate.
* `issued_by` - 
* `issuer_name` - 
	* `common_name` - The Certificate Authority (CA) name.
	* `country` - ISO 3166-1 alpha-2 code of the country where the organization is located. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
	* `email_address` - The email address of the server's administrator.
	* `locality` - The city in which the organization is located.
	* `organization` - The organization name.
	* `organizational_unit` - The field to differentiate between divisions within an organization.
	* `state_province` - The province where the organization is located.
* `public_key_info` - 
	* `algorithm` - The algorithm identifier and parameters for the public key.
	* `exponent` - The private key exponent.
	* `key_size` - The number of bits in a key used by a cryptographic algorithm.
* `serial_number` - A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
* `signature_algorithm` - The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
* `state` - The current lifecycle state of the SSL certificate.
* `subject_name` - 
	* `common_name` - The fully qualified domain name used for DNS lookups of the server.
	* `country` - ISO 3166-1 alpha-2 code of the country where the organization is located. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
	* `email_address` - The email address of the server's administrator.
	* `locality` - The city in which the organization is located.
	* `organization` - The organization name.
	* `organizational_unit` - The field to differentiate between divisions within an organization.
	* `state_province` - The province where the organization is located.
* `time_created` - The date and time the certificate was created, expressed in RFC 3339 timestamp format.
* `time_not_valid_after` - The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
* `time_not_valid_before` - The date and time the certificate will become valid, expressed in RFC 3339 timestamp format.
* `version` - The version of the encoded certificate.

