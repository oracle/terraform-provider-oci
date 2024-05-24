---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_report"
sidebar_current: "docs-oci-resource-jms_java_downloads-java_download_report"
description: |-
  Provides the Java Download Report resource in Oracle Cloud Infrastructure Jms Java Downloads service
---

# oci_jms_java_downloads_java_download_report
This resource provides the Java Download Report resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Create a new report in the specified format containing the download details
for the tenancy.


## Example Usage

```hcl
resource "oci_jms_java_downloads_java_download_report" "test_java_download_report" {
	#Required
	compartment_id = var.tenancy_ocid
	format = var.java_download_report_format

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	time_end = var.java_download_report_time_end
	time_start = var.java_download_report_time_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) here should be the tenancy OCID. 
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `format` - (Required) The format of the report that is generated.
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `time_end` - (Optional) The end time until when the download records have to be included (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `time_start` - (Optional) The start time from when download records have to be included (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `checksum_type` - The algorithm used for calculating the checksum.
* `checksum_value` - The checksum value of the Java download report file.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the Java download report. 
* `created_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `email` - The email of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `display_name` - Display name for the Java download report.
* `file_size_in_bytes` - Approximate size of the Java download report file in bytes.
* `format` - The file format of the Java download report.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Java download report. 
* `state` - The current state of the Java download report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Java download report was created, displayed as an [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
* `time_end` - The end time until when the download records are included (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `time_start` - The start time from when the download records are included (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Java Download Report
	* `update` - (Defaults to 20 minutes), when updating the Java Download Report
	* `delete` - (Defaults to 20 minutes), when destroying the Java Download Report


## Import

Import is not supported for this resource.

