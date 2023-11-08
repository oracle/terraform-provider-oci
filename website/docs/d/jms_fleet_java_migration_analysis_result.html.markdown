---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_java_migration_analysis_result"
sidebar_current: "docs-oci-datasource-jms-fleet_java_migration_analysis_result"
description: |-
  Provides details about a specific Fleet Java Migration Analysis Result in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_java_migration_analysis_result
This data source provides details about a specific Fleet Java Migration Analysis Result resource in Oracle Cloud Infrastructure Jms service.

Retrieve Java Migration Analysis result.

## Example Usage

```hcl
data "oci_jms_fleet_java_migration_analysis_result" "test_fleet_java_migration_analysis_result" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	java_migration_analysis_result_id = oci_apm_synthetics_result.test_result.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `java_migration_analysis_result_id` - (Required) The OCID of the analysis result.


## Attributes Reference

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

