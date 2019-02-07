---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_certificates"
sidebar_current: "docs-oci-datasource-waas-certificates"
description: |-
  Provides the list of Certificates in Oracle Cloud Infrastructure Waas service
---

# Data Source: oci_waas_certificates
This data source provides the list of Certificates in Oracle Cloud Infrastructure Waas service.

Gets a list of SSL certificates that can be used in a WAAS policy.

## Example Usage

```hcl
data "oci_waas_certificates" "test_certificates" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_names = "${var.certificate_display_names}"
	ids = "${var.certificate_ids}"
	states = "${var.certificate_states}"
	time_created_greater_than_or_equal_to = "${var.certificate_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.certificate_time_created_less_than}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
* `display_names` - (Optional) Filter certificates using a list of display names.
* `ids` - (Optional) Filter certificates using a list of certificates OCIDs.
* `states` - (Optional) Filter certificates using a list of lifecycle states.
* `time_created_greater_than_or_equal_to` - (Optional) A filter that matches certificates created on or after the specified date-time.
* `time_created_less_than` - (Optional) A filter that matches certificates created before the specified date-time.


## Attributes Reference

The following attributes are exported:

* `certificates` - The list of certificates.

### Certificate Reference

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

