---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_certificate"
sidebar_current: "docs-oci-resource-waas-certificate"
description: |-
  Provides the Certificate resource in Oracle Cloud Infrastructure Waas service
---

# oci_waas_certificate
This resource provides the Certificate resource in Oracle Cloud Infrastructure Waas service.

Allows an SSL certificate to be added to a WAAS policy. The Web Application Firewall terminates SSL connections to inspect requests in runtime, and then re-encrypts requests before sending them to the origin for fulfillment.

For more information, see [WAF Settings](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/wafsettings.htm).

## Example Usage

```hcl
resource "oci_waas_certificate" "test_certificate" {
	#Required
	certificate_data = "${var.certificate_certificate_data}"
	compartment_id = "${var.compartment_id}"
	private_key_data = "${var.certificate_private_key_data}"

	#Optional
	defined_tags = "${var.certificate_defined_tags}"
	display_name = "${var.certificate_display_name}"
	freeform_tags = "${var.certificate_freeform_tags}"
	is_trust_verification_disabled = "${var.certificate_is_trust_verification_disabled}"
}
```

## Argument Reference

The following arguments are supported:

* `certificate_data` - (Required) The data of the SSL certificate.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example:
* `display_name` - (Optional) (Updatable) A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `is_trust_verification_disabled` - (Optional) Set to true if the SSL certificate is self-signed.
* `private_key_data` - (Required) The private key of the SSL certificate.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

Import is not supported for this resource.

