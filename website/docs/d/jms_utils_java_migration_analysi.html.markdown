---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_java_migration_analysi"
sidebar_current: "docs-oci-datasource-jms_utils-java_migration_analysi"
description: |-
  Provides details about a specific Java Migration Analysi in Oracle Cloud Infrastructure Jms Utils service
---

# Data Source: oci_jms_utils_java_migration_analysi
This data source provides details about a specific Java Migration Analysi resource in Oracle Cloud Infrastructure Jms Utils service.

Gets information about a Java Migration Analysis.

## Example Usage

```hcl
data "oci_jms_utils_java_migration_analysi" "test_java_migration_analysi" {
	#Required
	java_migration_analysis_id = oci_jms_utils_java_migration_analysi.test_java_migration_analysi.id
}
```

## Argument Reference

The following arguments are supported:

* `java_migration_analysis_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Java Migration Analysis.


## Attributes Reference

The following attributes are exported:

* `analysis_project_name` - Name of the analysis project.
* `analysis_result_files` - The analysis application file names result in the Object Storage.
* `analysis_result_object_storage_path` - Path to the Object Storage analysis application result.
* `bucket` - Object storage bucket name.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Java Migration Analysis.
* `input_applications_object_storage_paths` - Object storage paths to the input files applications to be analysed.
* `metadata` - Additional info reserved for future use.
* `namespace` - Object storage namespace.
* `target_jdk_version` - Jdk Version of the Java Migration Analysis target.
* `time_created` - The date and time the Java Migration Analysis was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_finished` - The date and time the Java Migration Analysis was finished, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_started` - The date and time the Java Migration Analysis was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `work_request_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Work Request.

