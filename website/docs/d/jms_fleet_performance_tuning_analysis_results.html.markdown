---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_performance_tuning_analysis_results"
sidebar_current: "docs-oci-datasource-jms-fleet_performance_tuning_analysis_results"
description: |-
  Provides the list of Fleet Performance Tuning Analysis Results in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_performance_tuning_analysis_results
This data source provides the list of Fleet Performance Tuning Analysis Results in Oracle Cloud Infrastructure Jms service.

List Performance Tuning Analysis results.

## Example Usage

```hcl
data "oci_jms_fleet_performance_tuning_analysis_results" "test_fleet_performance_tuning_analysis_results" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	application_id = var.fleet_performance_tuning_analysis_result_application_id
	host_name = var.fleet_performance_tuning_analysis_result_host_name
	managed_instance_id = var.fleet_performance_tuning_analysis_result_managed_instance_id
	time_end = var.fleet_performance_tuning_analysis_result_time_end
	time_start = var.fleet_performance_tuning_analysis_result_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the related application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `host_name` - (Optional) The host [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `performance_tuning_analysis_result_collection` - The list of performance_tuning_analysis_result_collection.

### FleetPerformanceTuningAnalysisResult Reference

The following attributes are exported:

* `application_id` - The OCID of the application for which the report has been generated.
* `application_installation_id` - The internal identifier of the application installation for which the report has been generated.
* `application_installation_path` - The installation path of the application for which the report has been generated.
* `application_name` - The name of the application for which the report has been generated.
* `bucket` - The Object Storage bucket name of this analysis result.
* `fleet_id` - The fleet OCID.
* `host_name` - The hostname of the managed instance.
* `id` - The OCID to identify this analysis results.
* `managed_instance_id` - The managed instance OCID.
* `namespace` - The Object Storage namespace of this analysis result.
* `object` - The Object Storage object name of this analysis result.
* `result` - Result of the analysis based on whether warnings have been found or not.
* `time_created` - The time the result is compiled.
* `time_finished` - The time the JFR recording has finished.
* `time_started` - The time the JFR recording has started.
* `warning_count` - Total number of warnings reported by the analysis.
* `work_request_id` - The OCID of the work request to start the analysis.

