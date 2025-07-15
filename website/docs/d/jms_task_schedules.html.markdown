---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_task_schedules"
sidebar_current: "docs-oci-datasource-jms-task_schedules"
description: |-
  Provides the list of Task Schedules in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_task_schedules
This data source provides the list of Task Schedules in Oracle Cloud Infrastructure Jms service.

Returns a list of task schedules.


## Example Usage

```hcl
data "oci_jms_task_schedules" "test_task_schedules" {

	#Optional
	fleet_id = oci_jms_fleet.test_fleet.id
	id = var.task_schedule_id
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	name = var.task_schedule_name
	task_schedule_name_contains = var.task_schedule_task_schedule_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Optional) The ID of the Fleet.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to identify this task schedule.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `name` - (Optional) The task name.
* `task_schedule_name_contains` - (Optional) Filter the list with task schedule name contains the given value. 


## Attributes Reference

The following attributes are exported:

* `task_schedule_collection` - The list of task_schedule_collection.

### TaskSchedule Reference

The following attributes are exported:

* `created_by` - Name of the task creator.
* `execution_recurrences` - Recurrence specification for the task schedule execution (formatted according to [RFC-5545](https://icalendar.org/RFC-Specifications/iCalendar-RFC-5545/)). To run daily for 10 occurrences starts on September 2, 2024 09:00 EDT, it should be 'DTSTART;TZID=America/New_York:20240902T090000 RRULE:FREQ=DAILY;COUNT=10'. To run every 3 hours from 9:00 AM to 5:00 PM on August 5, 2024 EDT, it should be 'DTSTART;TZID=America/New_York:20240805T090000 RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=20240805T170000Z'. 
* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to identify this task schedule.
* `name` - The name of the task schedule.
* `state` - All possible status of task schedule.
* `task_details` - The minimum details of a task.
	* `add_installation_site_task_request` - The list of Java installation sites to add.
		* `installation_sites` - The list of installation sites to add.
			* `artifact_content_type` - Artifact content type for the Java version.
			* `force_install` - Forces the installation request even if a more recent release is already present in the host.
			* `headless_mode` - Flag to install headless or headful Java installation. Only valid for Oracle Linux in OCI.
			* `installation_path` - Custom path to install new Java installation site.
			* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
			* `release_version` - The release version of the Java Runtime.
		* `post_installation_actions` - Optional list of post java installation actions
	* `crypto_task_request` - Details of the request to start a JFR crypto event analysis. When the targets aren't specified, then all managed instances currently in the fleet are selected. 
		* `recording_duration_in_minutes` - Duration of the JFR recording in minutes.
		* `targets` - The attachment targets to start JFR.
			* `application_installation_key` - Unique key that identifies the application installation for JFR data collection.
			* `application_key` - Unique key that identifies the application for JFR data collection.
			* `container_key` - Unique key that identifies the container for JFR data collection.
			* `jre_key` - Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `deployed_application_migration_task_request` - Details of the request to start a Java migration analyses. The analyses requires the managed instance OCID, deployed application key, source JDK version, and target JDK version of each selected application. 
		* `targets` - An array of migration analyses requests.
			* `deployed_application_installation_key` - The unique key that identifies the deployed application's installation path that is to be used for the Java migration analyses.
			* `exclude_package_prefixes` - Excludes the packages that starts with the prefix from the migration analyses result. Either this or includePackagePrefixes can be specified.
			* `include_package_prefixes` - Includes the packages that starts with the prefix from the migration analyses result. Either this or excludePackagePrefixes can be specified.
			* `managed_instance_id` - The OCID of the managed instance that hosts the application for which the Java migration analyses was performed.
			* `source_jdk_version` - The JDK version the application is currently running on.
			* `target_jdk_version` - The JDK version against which the migration analyses was performed to identify effort required to move from source JDK.
	* `java_migration_task_request` - Details of the request to start a Java migration analysis. The analysis requires the managed instance OCID, application installation key, source JDK version, and target JDK version of each selected application. 
		* `targets` - An array of migration analysis requests.
			* `application_installation_key` - The unique key that identifies the application's installation path that is to be used for the Java migration analysis.
			* `exclude_package_prefixes` - Excludes the packages that starts with the prefix from the migration analysis result. Either this or includePackagePrefixes can be specified.
			* `include_package_prefixes` - includes the packages that starts with the prefix from the migration analysis result. Either this or excludePackagePrefixes can be specified.
			* `managed_instance_id` - The OCID of the managed instance that hosts the application for which the Java migration analysis was performed.
			* `source_jdk_version` - The JDK version the application is currently running on.
			* `target_jdk_version` - The JDK version against which the migration analysis was performed to identify effort required to move from source JDK.
	* `jfr_task_request` - Details of the request to start JFR recordings. When the targets aren't specified, then all managed instances currently in the Fleet are selected. 
		* `jfc_profile_name` - The profile used for JFR events selection. If the name isn't recognized, the settings from jfcV1 or jfcV2 will be used depending on the JVM version. Both jfcV2 and jfcV1 should be provided to ensure JFR collection on different JVM versions. 
		* `jfc_v1` - The BASE64 encoded string of JFR settings XML with schema used by JDK 8.
		* `jfc_v2` - The BASE64 encoded string of JFR settings XML with [schema used by JDK 9 and after](https://raw.githubusercontent.com/openjdk/jdk/master/src/jdk.jfr/share/classes/jdk/jfr/internal/jfc/jfc.xsd). 
		* `recording_duration_in_minutes` - Duration of the JFR recording in minutes.
		* `recording_size_in_mb` - The maximum size limit for the JFR file collected.
		* `targets` - The attachment targets to start JFR.
			* `application_installation_key` - Unique key that identifies the application installation for JFR data collection.
			* `application_key` - Unique key that identifies the application for JFR data collection.
			* `container_key` - Unique key that identifies the container for JFR data collection.
			* `jre_key` - Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `performance_tuning_task_request` - Details of the request to start a JFR performance tuning analysis. 
		* `recording_duration_in_minutes` - Duration of the JFR recording in minutes.
		* `targets` - The attachment targets to start JFR.
			* `application_installation_key` - Unique key that identifies the application installation for JFR data collection.
			* `application_key` - Unique key that identifies the application for JFR data collection.
			* `container_key` - Unique key that identifies the container for JFR data collection.
			* `jre_key` - Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `remove_installation_site_task_request` - The list of Java installation sites to remove.
		* `installation_sites` - The list of installation sites to remove.
			* `installation_key` - The unique identifier for the installation of a Java Runtime at a specific path on a specific operating system.
			* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
	* `scan_java_server_task_request` - The list of managed instances to scan.
		* `managed_instance_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of managed instances to scan.
	* `scan_library_task_request` - The list of managed instances to scan. 
		* `dynamic_scan_duration_in_minutes` - The duration of the dynamic scan in minutes. 
		* `is_dynamic_scan` - Indicates whether the scan is dynamic or static. 
		* `managed_instance_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of managed instances to scan. 
	* `task_type` - Type of task.
* `time_created` - The date and time the task schedule was created (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_last_run` - The date and time the task schedule ran last (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_last_updated` - The date and time the task schedule was last updated (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_next_run` - The date and time the task schedule will run next (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).

