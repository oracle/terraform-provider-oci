---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_license_acceptance_records"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_license_acceptance_records"
description: |-
  Provides the list of Java License Acceptance Records in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_license_acceptance_records
This data source provides the list of Java License Acceptance Records in Oracle Cloud Infrastructure Jms Java Downloads service.

Returns a list of all the Java license acceptance records in a tenancy.


## Example Usage

```hcl
data "oci_jms_java_downloads_java_license_acceptance_records" "test_java_license_acceptance_records" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	id = var.java_license_acceptance_record_id
	license_type = var.java_license_acceptance_record_license_type
	search_by_user = var.java_license_acceptance_record_search_by_user
	status = var.java_license_acceptance_record_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `id` - (Optional) Unique Java license acceptance record identifier.
* `license_type` - (Optional) Unique Java license type.
* `search_by_user` - (Optional) A filter to return only resources that match the user principal detail.  The search string can be any of the property values from the [Principal](https://docs.cloud.oracle.com/iaas/api/#/en/jms/latest/datatypes/Principal) object. This object is used as response datatype for the `createdBy` and `lastUpdatedBy` fields in applicable resource. 
* `status` - (Optional) The status of license acceptance.


## Attributes Reference

The following attributes are exported:

* `java_license_acceptance_record_collection` - The list of java_license_acceptance_record_collection.

### JavaLicenseAcceptanceRecord Reference

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
* `time_accepted` - The date and time of license acceptance (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `time_last_updated` - The date and time of last update (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

