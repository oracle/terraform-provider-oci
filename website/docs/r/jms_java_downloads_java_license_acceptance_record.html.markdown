---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_license_acceptance_record"
sidebar_current: "docs-oci-resource-jms_java_downloads-java_license_acceptance_record"
description: |-
  Provides the Java License Acceptance Record resource in Oracle Cloud Infrastructure Jms Java Downloads service
---

# oci_jms_java_downloads_java_license_acceptance_record
This resource provides the Java License Acceptance Record resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Creates a Java license acceptance record for the specified license type in a tenancy.


## Example Usage

```hcl
resource "oci_jms_java_downloads_java_license_acceptance_record" "test_java_license_acceptance_record" {
	#Required
	compartment_id = var.tenancy_ocid
	license_acceptance_status = var.java_license_acceptance_record_license_acceptance_status
	license_type = var.java_license_acceptance_record_license_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user accepting the license.
* `license_acceptance_status` - (Required) (Updatable) Status of license acceptance.
* `license_type` - (Required) License type for the Java version.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user accepting the license.
* `created_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `email` - The email of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `id` - The unique identifier for the acceptance record.
* `last_updated_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `email` - The email of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `license_acceptance_status` - Status of license acceptance.
* `license_type` - License type associated with the acceptance.
* `state` - The current state of the JavaLicenseAcceptanceRecord.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_accepted` - The date and time of license acceptance(formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `time_last_updated` - The date and time of last update(formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Java License Acceptance Record
	* `update` - (Defaults to 20 minutes), when updating the Java License Acceptance Record
	* `delete` - (Defaults to 20 minutes), when destroying the Java License Acceptance Record


## Import

Import is not supported for this resource.

