// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# You need actual OCID value of fleet and application to create this resource.
# variable "task_schedule_execution_recurrences" {
#   default = "executionRecurrences"
# }

variable "task_schedule_id" {
  default = "id"
}

variable "task_schedule_name" {
  default = "name"
}

# variable "task_schedule_task_details_add_installation_site_task_request_installation_sites_artifact_content_type" {
#   default = "JDK"
# }

# variable "task_schedule_task_details_add_installation_site_task_request_installation_sites_force_install" {
#   default = false
# }

# variable "task_schedule_task_details_add_installation_site_task_request_installation_sites_headless_mode" {
#   default = false
# }

# variable "task_schedule_task_details_add_installation_site_task_request_installation_sites_installation_path" {
#   default = "installationPath"
# }

# variable "task_schedule_task_details_add_installation_site_task_request_installation_sites_release_version" {
#   default = "releaseVersion"
# }

# variable "task_schedule_task_details_add_installation_site_task_request_post_installation_actions" {
#   default = []
# }

# variable "task_schedule_task_details_crypto_task_request_recording_duration_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_crypto_task_request_targets_application_installation_key" {
#   default = "applicationInstallationKey"
# }

# variable "task_schedule_task_details_crypto_task_request_targets_application_key" {
#   default = "applicationKey"
# }

# variable "task_schedule_task_details_crypto_task_request_targets_container_key" {
#   default = "containerKey"
# }

# variable "task_schedule_task_details_crypto_task_request_targets_jre_key" {
#   default = "jreKey"
# }

# variable "task_schedule_task_details_crypto_task_request_waiting_period_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_deployed_application_migration_task_request_targets_deployed_application_installation_key" {
#   default = "deployedApplicationInstallationKey"
# }

# variable "task_schedule_task_details_deployed_application_migration_task_request_targets_exclude_package_prefixes" {
#   default = []
# }

# variable "task_schedule_task_details_deployed_application_migration_task_request_targets_include_package_prefixes" {
#   default = []
# }

# variable "task_schedule_task_details_deployed_application_migration_task_request_targets_source_jdk_version" {
#   default = "sourceJdkVersion"
# }

# variable "task_schedule_task_details_deployed_application_migration_task_request_targets_target_jdk_version" {
#   default = "targetJdkVersion"
# }

# variable "task_schedule_task_details_java_migration_task_request_targets_application_installation_key" {
#   default = "applicationInstallationKey"
# }

# variable "task_schedule_task_details_java_migration_task_request_targets_exclude_package_prefixes" {
#   default = []
# }

# variable "task_schedule_task_details_java_migration_task_request_targets_include_package_prefixes" {
#   default = []
# }

# variable "task_schedule_task_details_java_migration_task_request_targets_source_jdk_version" {
#   default = "sourceJdkVersion"
# }

# variable "task_schedule_task_details_java_migration_task_request_targets_target_jdk_version" {
#   default = "targetJdkVersion"
# }

# variable "task_schedule_task_details_jfr_task_request_jfc_profile_name" {
#   default = "jfcV1"
# }

# variable "task_schedule_task_details_jfr_task_request_jfc_v1" {
#   default = "jfcV1"
# }

# variable "task_schedule_task_details_jfr_task_request_jfc_v2" {
#   default = "jfcV2"
# }

# variable "task_schedule_task_details_jfr_task_request_recording_duration_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_jfr_task_request_recording_size_in_mb" {
#   default = 10
# }

# variable "task_schedule_task_details_jfr_task_request_targets_application_installation_key" {
#   default = "applicationInstallationKey"
# }

# variable "task_schedule_task_details_jfr_task_request_targets_application_key" {
#   default = "applicationKey"
# }

# variable "task_schedule_task_details_jfr_task_request_targets_container_key" {
#   default = "containerKey"
# }

# variable "task_schedule_task_details_jfr_task_request_targets_jre_key" {
#   default = "jreKey"
# }

# variable "task_schedule_task_details_jfr_task_request_waiting_period_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_performance_tuning_task_request_recording_duration_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_performance_tuning_task_request_targets_application_installation_key" {
#   default = "applicationInstallationKey"
# }

# variable "task_schedule_task_details_performance_tuning_task_request_targets_application_key" {
#   default = "applicationKey"
# }

# variable "task_schedule_task_details_performance_tuning_task_request_targets_container_key" {
#   default = "containerKey"
# }

# variable "task_schedule_task_details_performance_tuning_task_request_targets_jre_key" {
#   default = "jreKey"
# }

# variable "task_schedule_task_details_performance_tuning_task_request_waiting_period_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_remove_installation_site_task_request_installation_sites_installation_key" {
#   default = "installationKey"
# }

# variable "task_schedule_task_details_scan_java_server_task_request_managed_instance_ids" {
#   default = []
# }

# variable "task_schedule_task_details_scan_library_task_request_dynamic_scan_duration_in_minutes" {
#   default = 10
# }

# variable "task_schedule_task_details_scan_library_task_request_is_dynamic_scan" {
#   default = false
# }

# variable "task_schedule_task_details_scan_library_task_request_managed_instance_ids" {
#   default = []
# }

# variable "task_schedule_task_details_task_type" {
#   default = "CRYPTO"
# }

variable "task_schedule_task_schedule_name_contains" {
  default = "taskScheduleNameContains"
}

# resource "oci_jms_task_schedule" "test_task_schedule" {
#   #Required
#   execution_recurrences = var.task_schedule_execution_recurrences
#   fleet_id              = var.fleet_ocid
#   task_details {
#     #Required
#     task_type = var.task_schedule_task_details_task_type

#     #Optional
#     add_installation_site_task_request {

#       #Optional
#       installation_sites {

#         #Optional
#         artifact_content_type = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_artifact_content_type
#         force_install         = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_force_install
#         headless_mode         = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_headless_mode
#         installation_path     = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_installation_path
#         managed_instance_id   = var.managed_instance_ocid
#         release_version       = var.task_schedule_task_details_add_installation_site_task_request_installation_sites_release_version
#       }
#       post_installation_actions = var.task_schedule_task_details_add_installation_site_task_request_post_installation_actions
#     }
#     crypto_task_request {

#       #Optional
#       recording_duration_in_minutes = var.task_schedule_task_details_crypto_task_request_recording_duration_in_minutes
#       targets {

#         #Optional
#         application_installation_key = var.task_schedule_task_details_crypto_task_request_targets_application_installation_key
#         application_key              = var.task_schedule_task_details_crypto_task_request_targets_application_key
#         container_key                = var.task_schedule_task_details_crypto_task_request_targets_container_key
#         jre_key                      = var.task_schedule_task_details_crypto_task_request_targets_jre_key
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = var.task_schedule_task_details_crypto_task_request_waiting_period_in_minutes
#     }
#     deployed_application_migration_task_request {

#       #Optional
#       targets {

#         #Optional
#         deployed_application_installation_key = var.task_schedule_task_details_deployed_application_migration_task_request_targets_deployed_application_installation_key
#         exclude_package_prefixes              = var.task_schedule_task_details_deployed_application_migration_task_request_targets_exclude_package_prefixes
#         include_package_prefixes              = var.task_schedule_task_details_deployed_application_migration_task_request_targets_include_package_prefixes
#         managed_instance_id                   = var.managed_instance_ocid
#         source_jdk_version                    = var.task_schedule_task_details_deployed_application_migration_task_request_targets_source_jdk_version
#         target_jdk_version                    = var.task_schedule_task_details_deployed_application_migration_task_request_targets_target_jdk_version
#       }
#     }
#     java_migration_task_request {

#       #Optional
#       targets {

#         #Optional
#         application_installation_key = var.task_schedule_task_details_java_migration_task_request_targets_application_installation_key
#         exclude_package_prefixes     = var.task_schedule_task_details_java_migration_task_request_targets_exclude_package_prefixes
#         include_package_prefixes     = var.task_schedule_task_details_java_migration_task_request_targets_include_package_prefixes
#         managed_instance_id          = var.managed_instance_ocid
#         source_jdk_version           = var.task_schedule_task_details_java_migration_task_request_targets_source_jdk_version
#         target_jdk_version           = var.task_schedule_task_details_java_migration_task_request_targets_target_jdk_version
#       }
#     }
#     jfr_task_request {

#       #Optional
#       jfc_profile_name              = var.task_schedule_task_details_jfr_task_request_jfc_profile_name
#       jfc_v1                        = var.task_schedule_task_details_jfr_task_request_jfc_v1
#       jfc_v2                        = var.task_schedule_task_details_jfr_task_request_jfc_v2
#       recording_duration_in_minutes = var.task_schedule_task_details_jfr_task_request_recording_duration_in_minutes
#       recording_size_in_mb          = var.task_schedule_task_details_jfr_task_request_recording_size_in_mb
#       targets {

#         #Optional
#         application_installation_key = var.task_schedule_task_details_jfr_task_request_targets_application_installation_key
#         application_key              = var.task_schedule_task_details_jfr_task_request_targets_application_key
#         container_key                = var.task_schedule_task_details_jfr_task_request_targets_container_key
#         jre_key                      = var.task_schedule_task_details_jfr_task_request_targets_jre_key
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = var.task_schedule_task_details_jfr_task_request_waiting_period_in_minutes
#     }
#     performance_tuning_task_request {

#       #Optional
#       recording_duration_in_minutes = var.task_schedule_task_details_performance_tuning_task_request_recording_duration_in_minutes
#       targets {

#         #Optional
#         application_installation_key = var.task_schedule_task_details_performance_tuning_task_request_targets_application_installation_key
#         application_key              = var.task_schedule_task_details_performance_tuning_task_request_targets_application_key
#         container_key                = var.task_schedule_task_details_performance_tuning_task_request_targets_container_key
#         jre_key                      = var.task_schedule_task_details_performance_tuning_task_request_targets_jre_key
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = var.task_schedule_task_details_performance_tuning_task_request_waiting_period_in_minutes
#     }
#     remove_installation_site_task_request {

#       #Optional
#       installation_sites {

#         #Optional
#         installation_key    = var.task_schedule_task_details_remove_installation_site_task_request_installation_sites_installation_key
#         managed_instance_id = var.managed_instance_ocid
#       }
#     }
#     scan_java_server_task_request {

#       #Optional
#       managed_instance_ids = var.task_schedule_task_details_scan_java_server_task_request_managed_instance_ids
#     }
#     scan_library_task_request {

#       #Optional
#       dynamic_scan_duration_in_minutes = var.task_schedule_task_details_scan_library_task_request_dynamic_scan_duration_in_minutes
#       is_dynamic_scan                  = var.task_schedule_task_details_scan_library_task_request_is_dynamic_scan
#       managed_instance_ids             = var.task_schedule_task_details_scan_library_task_request_managed_instance_ids
#     }
#   }
# }

data "oci_jms_task_schedules" "test_task_schedules" {

  #Optional
  fleet_id                    = var.fleet_ocid
  id                          = var.task_schedule_id
  managed_instance_id         = var.managed_instance_ocid
  name                        = var.task_schedule_name
  task_schedule_name_contains = var.task_schedule_task_schedule_name_contains
}