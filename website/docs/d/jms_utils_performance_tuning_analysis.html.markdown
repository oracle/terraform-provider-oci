---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_performance_tuning_analysis"
sidebar_current: "docs-oci-datasource-jms_utils-performance_tuning_analysis"
description: |-
  Provides the list of Performance Tuning Analysis in Oracle Cloud Infrastructure Jms Utils service
---

# Data Source: oci_jms_utils_performance_tuning_analysis
This data source provides the list of Performance Tuning Analysis in Oracle Cloud Infrastructure Jms Utils service.

Gets a list of Performance tuning Analysis.


## Example Usage

```hcl
data "oci_jms_utils_performance_tuning_analysis" "test_performance_tuning_analysis" {

	#Optional
	analysis_project_name = oci_ai_anomaly_detection_project.test_project.name
	compartment_id = var.compartment_id
	id = var.performance_tuning_analysi_id
	performance_tuning_analysis_result = var.performance_tuning_analysi_performance_tuning_analysis_result
}
```

## Argument Reference

The following arguments are supported:

* `analysis_project_name` - (Optional) The project name of the Performance Tuning Analysis to query for.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Performance Tuning Analysis.
* `performance_tuning_analysis_result` - (Optional) The result of the Performance Tuning Analysis to query for.


## Attributes Reference

The following attributes are exported:

* `performance_tuning_analysis_collection` - The list of performance_tuning_analysis_collection.

### PerformanceTuningAnalysi Reference

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

