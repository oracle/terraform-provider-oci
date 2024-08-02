---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_java_migration_analysis_results"
sidebar_current: "docs-oci-datasource-jms-fleet_java_migration_analysis_results"
description: |-
  Provides the list of Fleet Java Migration Analysis Results in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_java_migration_analysis_results
This data source provides the list of Fleet Java Migration Analysis Results in Oracle Cloud Infrastructure Jms service.

Lists the results of a Java migration analysis.

## Example Usage

```hcl
data "oci_jms_fleet_java_migration_analysis_results" "test_fleet_java_migration_analysis_results" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	application_name = var.fleet_java_migration_analysis_result_application_name
	host_name = var.fleet_java_migration_analysis_result_host_name
	managed_instance_id = var.fleet_java_migration_analysis_result_managed_instance_id
	time_end = var.fleet_java_migration_analysis_result_time_end
	time_start = var.fleet_java_migration_analysis_result_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_name` - (Optional) The name of the application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `host_name` - (Optional) The host [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `java_migration_analysis_result_collection` - The list of java_migration_analysis_result_collection.

### FleetJavaMigrationAnalysisResult Reference

The following attributes are exported:

* `application_execution_type` - Execution type of the application for an application type, such as WAR and EAR, that is deployed or installed.
* `application_key` - The unique key that identifies the application.
* `application_name` - The name of the application for which the Java migration analysis was performed.
* `application_path` - The installation path of the application for which the Java migration analysis was performed.
* `bucket` - The name of the object storage bucket that contains the results of the migration analysis.
* `fleet_id` - The fleet OCID.
* `host_name` - The hostname of the managed instance that hosts the application for which the Java migration analysis was performed.
* `id` - The OCID of the migration analysis report.
* `managed_instance_id` - The managed instance OCID.
* `metadata` - Additional info reserved for future use.
* `namespace` - The object storage namespace that contains the results of the migration analysis.
* `object_list` - The names of the object storage objects that contain the results of the migration analysis.
* `object_storage_upload_dir_path` - The directory path of the object storage bucket that contains the results of the migration analysis.
* `source_jdk_version` - The source JDK version of the application that's currently running.
* `target_jdk_version` - The target JDK version of the application to be migrated.
* `time_created` - The time the result is compiled.
* `work_request_id` - The OCID of the work request of this analysis.

