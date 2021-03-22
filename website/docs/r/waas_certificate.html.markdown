---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_certificate"
sidebar_current: "docs-oci-resource-waas-certificate"
description: |-
  Provides the Certificate resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# oci_waas_certificate
This resource provides the Certificate resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Allows an SSL certificate to be added to a WAAS policy. The Web Application Firewall terminates SSL connections to inspect requests in runtime, and then re-encrypts requests before sending them to the origin for fulfillment.

For more information, see [WAF Settings](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/wafsettings.htm).

## Example Usage

```hcl
resource "oci_waas_certificate" "test_certificate" {
	#Required
	certificate_data = var.certificate_certificate_data
	compartment_id = var.compartment_id
	private_key_data = var.certificate_private_key_data

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.certificate_display_name
	freeform_tags = {"Department"= "Finance"}
	is_trust_verification_disabled = var.certificate_is_trust_verification_disabled
}
```

## Argument Reference

The following arguments are supported:

* `certificate_data` - (Required) The data of the SSL certificate.

	 **Note:** Many SSL certificate providers require an intermediate certificate chain to ensure a trusted status. If your SSL certificate requires an intermediate certificate chain, please append the intermediate certificate key in the `certificateData` field after the leaf certificate issued by the SSL certificate provider. If you are unsure if your certificate requires an intermediate certificate chain, see your certificate provider's documentation.

	 The example below shows an intermediate certificate appended to a leaf certificate. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_trust_verification_disabled` - (Optional) Set to `true` if the SSL certificate is self-signed.
* `private_key_data` - (Required) The private key of the SSL certificate.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `certificate_data` - The data of the SSL certificate.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name of the SSL certificate.
* `extensions` - Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
	* `is_critical` - The critical flag of the extension. Critical extensions must be processed, non-critical extensions can be ignored.
	* `name` - The certificate extension name.
	* `value` - The certificate extension value.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate.
* `is_trust_verification_disabled` - This indicates whether trust verification was disabled during the creation of SSL certificate. If `true` SSL certificate trust verification was disabled and this SSL certificate is most likely self-signed. 
* `issued_by` - 
* `issuer_name` - The issuer of the certificate.
	* `common_name` - The Certificate Authority (CA) name.
	* `country` - ISO 3166-1 alpha-2 code of the country where the organization is located. For a list of codes, see [ISO's website](https://www.iso.org/obp/ui/#search/code/).
	* `email_address` - The email address of the server's administrator.
	* `locality` - The city in which the organization is located.
	* `organization` - The organization name.
	* `organizational_unit` - The field to differentiate between divisions within an organization.
	* `state_province` - The province where the organization is located.
* `public_key_info` - Information about the public key and the algorithm used by the public key.
	* `algorithm` - The algorithm identifier and parameters for the public key.
	* `exponent` - The private key exponent.
	* `key_size` - The number of bits in a key used by a cryptographic algorithm.
* `serial_number` - A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
* `signature_algorithm` - The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
* `state` - The current lifecycle state of the SSL certificate.
* `subject_name` - The entity to be secured by the certificate.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Certificate
	* `update` - (Defaults to 20 minutes), when updating the Certificate
	* `delete` - (Defaults to 20 minutes), when destroying the Certificate


## Import

Import is not supported for this resource.

