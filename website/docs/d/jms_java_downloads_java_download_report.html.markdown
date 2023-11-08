---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_report"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_download_report"
description: |-
  Provides details about a specific Java Download Report in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_download_report
This data source provides details about a specific Java Download Report resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Gets a JavaDownloadReport by the specified identifier.

## Example Usage

```hcl
data "oci_jms_java_downloads_java_download_report" "test_java_download_report" {
	#Required
	java_download_report_id = oci_jms_java_downloads_java_download_report.test_java_download_report.id
}
```

## Argument Reference

The following arguments are supported:

* `java_download_report_id` - (Required) Unique Java download report identifier.


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
* `time_created` - The time the Java download report was created. An RFC3339 formatted datetime string.

