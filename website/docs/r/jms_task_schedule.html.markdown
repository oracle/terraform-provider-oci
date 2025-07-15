---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_task_schedule"
sidebar_current: "docs-oci-resource-jms-task_schedule"
description: |-
  Provides the Task Schedule resource in Oracle Cloud Infrastructure Jms service
---

# oci_jms_task_schedule
This resource provides the Task Schedule resource in Oracle Cloud Infrastructure Jms service.

Create a task schedule using the information provided.

## Example Usage

```hcl
resource "oci_jms_task_schedule" "test_task_schedule" {
	#Required
	execution_recurrences = var.task_schedule_execution_recurrences
	fleet_id = oci_jms_fleet.test_fleet.id
	task_details {
		#Required
		task_type = var.task_schedule_task_details_task_type

		#Optional
		add_installation_site_task_request {

			#Optional
			installation_sites {

				#Optional
				artifact_content_type = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_artifact_content_type
				force_install = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_force_install
				headless_mode = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_headless_mode
				installation_path = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_installation_path
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
				release_version = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_release_version
			}
			post_installation_actions = var.task_schedule_task_details_add_installation_site_task_request_post_installation_actions
		}
		crypto_task_request {

			#Optional
			recording_duration_in_minutes = var.task_schedule_task_details_crypto_task_request_recording_duration_in_minutes
			targets {

				#Optional
				application_installation_key = var.task_schedule_task_details_crypto_task_request_targets_application_installation_key
				application_key = var.task_schedule_task_details_crypto_task_request_targets_application_key
				container_key = var.task_schedule_task_details_crypto_task_request_targets_container_key
				jre_key = var.task_schedule_task_details_crypto_task_request_targets_jre_key
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
			}
			waiting_period_in_minutes = var.task_schedule_task_details_crypto_task_request_waiting_period_in_minutes
		}
		deployed_application_migration_task_request {

			#Optional
			targets {

				#Optional
				deployed_application_installation_key = var.task_schedule_task_details_deployed_application_migration_task_request_targets_deployed_application_installation_key
				exclude_package_prefixes = var.task_schedule_task_details_deployed_application_migration_task_request_targets_exclude_package_prefixes
				include_package_prefixes = var.task_schedule_task_details_deployed_application_migration_task_request_targets_include_package_prefixes
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
				source_jdk_version = var.task_schedule_task_details_deployed_application_migration_task_request_targets_source_jdk_version
				target_jdk_version = var.task_schedule_task_details_deployed_application_migration_task_request_targets_target_jdk_version
			}
		}
		java_migration_task_request {

			#Optional
			targets {

				#Optional
				application_installation_key = var.task_schedule_task_details_java_migration_task_request_targets_application_installation_key
				exclude_package_prefixes = var.task_schedule_task_details_java_migration_task_request_targets_exclude_package_prefixes
				include_package_prefixes = var.task_schedule_task_details_java_migration_task_request_targets_include_package_prefixes
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
				source_jdk_version = var.task_schedule_task_details_java_migration_task_request_targets_source_jdk_version
				target_jdk_version = var.task_schedule_task_details_java_migration_task_request_targets_target_jdk_version
			}
		}
		jfr_task_request {

			#Optional
			jfc_profile_name = oci_optimizer_profile.test_profile.name
			jfc_v1 = var.task_schedule_task_details_jfr_task_request_jfc_v1
			jfc_v2 = var.task_schedule_task_details_jfr_task_request_jfc_v2
			recording_duration_in_minutes = var.task_schedule_task_details_jfr_task_request_recording_duration_in_minutes
			recording_size_in_mb = var.task_schedule_task_details_jfr_task_request_recording_size_in_mb
			targets {

				#Optional
				application_installation_key = var.task_schedule_task_details_jfr_task_request_targets_application_installation_key
				application_key = var.task_schedule_task_details_jfr_task_request_targets_application_key
				container_key = var.task_schedule_task_details_jfr_task_request_targets_container_key
				jre_key = var.task_schedule_task_details_jfr_task_request_targets_jre_key
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
			}
			waiting_period_in_minutes = var.task_schedule_task_details_jfr_task_request_waiting_period_in_minutes
		}
		performance_tuning_task_request {

			#Optional
			recording_duration_in_minutes = var.task_schedule_task_details_performance_tuning_task_request_recording_duration_in_minutes
			targets {

				#Optional
				application_installation_key = var.task_schedule_task_details_performance_tuning_task_request_targets_application_installation_key
				application_key = var.task_schedule_task_details_performance_tuning_task_request_targets_application_key
				container_key = var.task_schedule_task_details_performance_tuning_task_request_targets_container_key
				jre_key = var.task_schedule_task_details_performance_tuning_task_request_targets_jre_key
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
			}
			waiting_period_in_minutes = var.task_schedule_task_details_performance_tuning_task_request_waiting_period_in_minutes
		}
		remove_installation_site_task_request {

			#Optional
			installation_sites {

				#Optional
				installation_key = var.task_schedule_task_details_remove_installation_site_task_request_installation_sites_installation_key
				managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
			}
		}
		scan_java_server_task_request {

			#Optional
			managed_instance_ids = var.task_schedule_task_details_scan_java_server_task_request_managed_instance_ids
		}
		scan_library_task_request {

			#Optional
			dynamic_scan_duration_in_minutes = var.task_schedule_task_details_scan_library_task_request_dynamic_scan_duration_in_minutes
			is_dynamic_scan = var.task_schedule_task_details_scan_library_task_request_is_dynamic_scan
			managed_instance_ids = var.task_schedule_task_details_scan_library_task_request_managed_instance_ids
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `execution_recurrences` - (Required) (Updatable) Recurrence specification for the task schedule execution (formatted according to [RFC-5545](https://icalendar.org/RFC-Specifications/iCalendar-RFC-5545/)). To run daily for 10 occurrences starts on September 2, 2024 09:00 EDT, it should be 'DTSTART;TZID=America/New_York:20240902T090000 RRULE:FREQ=DAILY;COUNT=10'. To run every 3 hours from 9:00 AM to 5:00 PM on August 5, 2024 EDT, it should be 'DTSTART;TZID=America/New_York:20240805T090000 RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=20240805T170000Z'. 
* `fleet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `task_details` - (Required) (Updatable) The minimum details of a task.
	* `add_installation_site_task_request` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) The list of Java installation sites to add.
		* `installation_sites` - (Required when task_type=ADD_INSTALLATION_SITE) (Updatable) The list of installation sites to add.
			* `artifact_content_type` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) Artifact content type for the Java version.
			* `force_install` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) Forces the installation request even if a more recent release is already present in the host.
			* `headless_mode` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) Flag to install headless or headful Java installation. Only valid for Oracle Linux in OCI.
			* `installation_path` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) Custom path to install new Java installation site.
			* `managed_instance_id` - (Required when task_type=ADD_INSTALLATION_SITE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
			* `release_version` - (Required when task_type=ADD_INSTALLATION_SITE) (Updatable) The release version of the Java Runtime.
		* `post_installation_actions` - (Applicable when task_type=ADD_INSTALLATION_SITE) (Updatable) Optional list of post java installation actions
	* `crypto_task_request` - (Applicable when task_type=CRYPTO) (Updatable) Details of the request to start a JFR crypto event analysis. When the targets aren't specified, then all managed instances currently in the fleet are selected. 
		* `recording_duration_in_minutes` - (Applicable when task_type=CRYPTO) (Updatable) Duration of the JFR recording in minutes.
		* `targets` - (Applicable when task_type=CRYPTO) (Updatable) The attachment targets to start JFR.
			* `application_installation_key` - (Applicable when task_type=CRYPTO) (Updatable) Unique key that identifies the application installation for JFR data collection.
			* `application_key` - (Applicable when task_type=CRYPTO) (Updatable) Unique key that identifies the application for JFR data collection.
			* `container_key` - (Applicable when task_type=CRYPTO) (Updatable) Unique key that identifies the container for JFR data collection.
			* `jre_key` - (Applicable when task_type=CRYPTO) (Updatable) Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - (Required when task_type=CRYPTO) (Updatable) OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - (Applicable when task_type=CRYPTO) (Updatable) Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `deployed_application_migration_task_request` - (Applicable when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) Details of the request to start a Java migration analyses. The analyses requires the managed instance OCID, deployed application key, source JDK version, and target JDK version of each selected application. 
		* `targets` - (Required when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) An array of migration analyses requests.
			* `deployed_application_installation_key` - (Required when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) The unique key that identifies the deployed application's installation path that is to be used for the Java migration analyses.
			* `exclude_package_prefixes` - (Applicable when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) Excludes the packages that starts with the prefix from the migration analyses result. Either this or includePackagePrefixes can be specified.
			* `include_package_prefixes` - (Applicable when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) Includes the packages that starts with the prefix from the migration analyses result. Either this or excludePackagePrefixes can be specified.
			* `managed_instance_id` - (Required when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) The OCID of the managed instance that hosts the application for which the Java migration analyses was performed.
			* `source_jdk_version` - (Required when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) The JDK version the application is currently running on.
			* `target_jdk_version` - (Required when task_type=DEPLOYED_APPLICATION_MIGRATION) (Updatable) The JDK version against which the migration analyses was performed to identify effort required to move from source JDK.
	* `java_migration_task_request` - (Applicable when task_type=JAVA_MIGRATION) (Updatable) Details of the request to start a Java migration analysis. The analysis requires the managed instance OCID, application installation key, source JDK version, and target JDK version of each selected application. 
		* `targets` - (Required when task_type=JAVA_MIGRATION) (Updatable) An array of migration analysis requests.
			* `application_installation_key` - (Required when task_type=JAVA_MIGRATION) (Updatable) The unique key that identifies the application's installation path that is to be used for the Java migration analysis.
			* `exclude_package_prefixes` - (Applicable when task_type=JAVA_MIGRATION) (Updatable) Excludes the packages that starts with the prefix from the migration analysis result. Either this or includePackagePrefixes can be specified.
			* `include_package_prefixes` - (Applicable when task_type=JAVA_MIGRATION) (Updatable) includes the packages that starts with the prefix from the migration analysis result. Either this or excludePackagePrefixes can be specified.
			* `managed_instance_id` - (Required when task_type=JAVA_MIGRATION) (Updatable) The OCID of the managed instance that hosts the application for which the Java migration analysis was performed.
			* `source_jdk_version` - (Required when task_type=JAVA_MIGRATION) (Updatable) The JDK version the application is currently running on.
			* `target_jdk_version` - (Required when task_type=JAVA_MIGRATION) (Updatable) The JDK version against which the migration analysis was performed to identify effort required to move from source JDK.
	* `jfr_task_request` - (Applicable when task_type=JFR) (Updatable) Details of the request to start JFR recordings. When the targets aren't specified, then all managed instances currently in the Fleet are selected. 
		* `jfc_profile_name` - (Required when task_type=JFR) (Updatable) The profile used for JFR events selection. If the name isn't recognized, the settings from jfcV1 or jfcV2 will be used depending on the JVM version. Both jfcV2 and jfcV1 should be provided to ensure JFR collection on different JVM versions. 
		* `jfc_v1` - (Applicable when task_type=JFR) (Updatable) The BASE64 encoded string of JFR settings XML with schema used by JDK 8.
		* `jfc_v2` - (Applicable when task_type=JFR) (Updatable) The BASE64 encoded string of JFR settings XML with [schema used by JDK 9 and after](https://raw.githubusercontent.com/openjdk/jdk/master/src/jdk.jfr/share/classes/jdk/jfr/internal/jfc/jfc.xsd). 
		* `recording_duration_in_minutes` - (Applicable when task_type=JFR) (Updatable) Duration of the JFR recording in minutes.
		* `recording_size_in_mb` - (Applicable when task_type=JFR) (Updatable) The maximum size limit for the JFR file collected.
		* `targets` - (Applicable when task_type=JFR) (Updatable) The attachment targets to start JFR.
			* `application_installation_key` - (Applicable when task_type=JFR) (Updatable) Unique key that identifies the application installation for JFR data collection.
			* `application_key` - (Applicable when task_type=JFR) (Updatable) Unique key that identifies the application for JFR data collection.
			* `container_key` - (Applicable when task_type=JFR) (Updatable) Unique key that identifies the container for JFR data collection.
			* `jre_key` - (Applicable when task_type=JFR) (Updatable) Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - (Required when task_type=JFR) (Updatable) OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - (Applicable when task_type=JFR) (Updatable) Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `performance_tuning_task_request` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Details of the request to start a JFR performance tuning analysis. 
		* `recording_duration_in_minutes` - (Required when task_type=PERFORMANCE_TUNING) (Updatable) Duration of the JFR recording in minutes.
		* `targets` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) The attachment targets to start JFR.
			* `application_installation_key` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Unique key that identifies the application installation for JFR data collection.
			* `application_key` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Unique key that identifies the application for JFR data collection.
			* `container_key` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Unique key that identifies the container for JFR data collection.
			* `jre_key` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Unique key that identify the JVM for JFR data collection.
			* `managed_instance_id` - (Required when task_type=PERFORMANCE_TUNING) (Updatable) OCID of the Managed Instance to collect JFR data.
		* `waiting_period_in_minutes` - (Applicable when task_type=PERFORMANCE_TUNING) (Updatable) Period to looking for JVMs. In addition to attach to running JVMs when given the command, JVM started within the waiting period will also be attached for JFR. The value should be larger than the agent polling interval setting for the fleet to ensure agent can get the instructions. If not specified, the agent polling interval for the fleet is used. 
	* `remove_installation_site_task_request` - (Applicable when task_type=REMOVE_INSTALLATION_SITE) (Updatable) The list of Java installation sites to remove.
		* `installation_sites` - (Required when task_type=REMOVE_INSTALLATION_SITE) (Updatable) The list of installation sites to remove.
			* `installation_key` - (Required when task_type=REMOVE_INSTALLATION_SITE) (Updatable) The unique identifier for the installation of a Java Runtime at a specific path on a specific operating system.
			* `managed_instance_id` - (Required when task_type=REMOVE_INSTALLATION_SITE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
	* `scan_java_server_task_request` - (Applicable when task_type=SCAN_JAVA_SERVER) (Updatable) The list of managed instances to scan.
		* `managed_instance_ids` - (Applicable when task_type=SCAN_JAVA_SERVER) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of managed instances to scan.
	* `scan_library_task_request` - (Applicable when task_type=SCAN_LIBRARY) (Updatable) The list of managed instances to scan. 
		* `dynamic_scan_duration_in_minutes` - (Applicable when task_type=SCAN_LIBRARY) (Updatable) The duration of the dynamic scan in minutes. 
		* `is_dynamic_scan` - (Applicable when task_type=SCAN_LIBRARY) (Updatable) Indicates whether the scan is dynamic or static. 
		* `managed_instance_ids` - (Applicable when task_type=SCAN_LIBRARY) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of managed instances to scan. 
	* `task_type` - (Required) (Updatable) Type of task.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Task Schedule
	* `update` - (Defaults to 20 minutes), when updating the Task Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Task Schedule


## Import

TaskSchedules can be imported using the `id`, e.g.

```
$ terraform import oci_jms_task_schedule.test_task_schedule "id"
```

