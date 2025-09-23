---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_performance_tuning_analysi"
sidebar_current: "docs-oci-datasource-jms_utils-performance_tuning_analysi"
description: |-
  Provides details about a specific Performance Tuning Analysi in Oracle Cloud Infrastructure Jms Utils service
---

# Data Source: oci_jms_utils_performance_tuning_analysi
This data source provides details about a specific Performance Tuning Analysi resource in Oracle Cloud Infrastructure Jms Utils service.

Gets information about a Performance Tuning Analysis.

## Example Usage

```hcl
data "oci_jms_utils_performance_tuning_analysi" "test_performance_tuning_analysi" {
	#Required
	performance_tuning_analysis_id = oci_jms_utils_performance_tuning_analysi.test_performance_tuning_analysi.id
}
```

## Argument Reference

The following arguments are supported:

* `performance_tuning_analysis_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Performance Tuning Analysis.


## Attributes Reference

The following attributes are exported:

* `analysis_project_name` - Name of the analysis project.
* `artifact_object_storage_path` - Object storage path to the artifact.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Performance Tuning Analysis.
* `result` - Possible Performance Tuning Result statuses.
* `result_object_storage_path` - Object storage path to the analysis.
* `time_created` - The date and time the Performance Tuning Analysis was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_finished` - The date and time the Performance Tuning Analysis was finished, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_started` - The date and time the Performance Tuning Analysis was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `warning_count` - Number of warnings in the Performance Tuning Analysis.
* `work_request_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Work Request.

