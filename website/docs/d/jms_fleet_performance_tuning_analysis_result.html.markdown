---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_performance_tuning_analysis_result"
sidebar_current: "docs-oci-datasource-jms-fleet_performance_tuning_analysis_result"
description: |-
  Provides details about a specific Fleet Performance Tuning Analysis Result in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_performance_tuning_analysis_result
This data source provides details about a specific Fleet Performance Tuning Analysis Result resource in Oracle Cloud Infrastructure Jms service.

Retrieve metadata of the Performance Tuning Analysis result.

## Example Usage

```hcl
data "oci_jms_fleet_performance_tuning_analysis_result" "test_fleet_performance_tuning_analysis_result" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	performance_tuning_analysis_result_id = var.fleet_performance_tuning_analysis_result_id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `performance_tuning_analysis_result_id` - (Required) The OCID of the performance tuning analysis result.


## Attributes Reference

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

