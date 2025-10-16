// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# You need actual OCIDs and/or keys of 
# fleet, application, application installation, jre, container
# to create this resource.
#
# variable "task_schedule_task_details_task_type" {
#   default = "CRYPTO"
# }

# resource "oci_jms_task_schedule" "test_task_schedule" {
#   #Required
#   fleet_id              = var.fleet_ocid
#   execution_recurrences = "DTSTART=20240805T090000Z;RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=20240805T170000Z"
#
#   task_details {
#     #Required
#     task_type = var.task_schedule_task_details_task_type
#
#     # Required for task_type="ADD_INSTALLATION_SITE"
#     add_installation_site_task_request {
#       installation_sites {
#         #Optional
#         artifact_content_type = "JDK"
#         force_install         = false
#         headless_mode         = false
#         installation_path     = "/example/install/path"
#         managed_instance_id   = var.managed_instance_ocid
#         release_version       = "17.0.0"
#       }
#       post_installation_actions = []
#     }
#
#     # Required for task_type="CRYPTO"
#     crypto_task_request {
#       #Optional
#       recording_duration_in_minutes = 10
#       targets {
#         #Optional
#         application_installation_key = "example-application-installation-key"
#         application_key              = "example-application-key"
#         container_key                = "example-container-key"
#         jre_key                      = "example-jre-key"
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = 10
#     }
#
#     # Required for task_type="DEPLOYED_APPLICATION_MIGRATION"
#     deployed_application_migration_task_request {
#       #Optional
#       targets {
#         #Optional
#         deployed_application_installation_key = "example-application-installation-key"
#         exclude_package_prefixes              = []
#         include_package_prefixes              = []
#         managed_instance_id                   = var.managed_instance_ocid
#         source_jdk_version                    = "11"
#         target_jdk_version                    = "21"
#       }
#     }
#
#     # Required for task_type="JAVA_MIGRATION"
#     java_migration_task_request {
#       #Optional
#       targets {
#         #Optional
#         application_installation_key = "example-application-installation-key"
#         exclude_package_prefixes     = []
#         include_package_prefixes     = []
#         managed_instance_id          = var.managed_instance_ocid
#         source_jdk_version           = "11"
#         target_jdk_version           = "21"
#       }
#     }
#
#     # Required for task_type="JFR"
#     jfr_task_request {
#       #Optional
#       jfc_profile_name              = "jfcV1"
#       jfc_v1                        = "jfcV1"
#       jfc_v2                        = "jfvV2"
#       recording_duration_in_minutes = 10
#       recording_size_in_mb          = 10
#       targets {
#         #Optional
#         application_installation_key = "example-application-installation-key"
#         application_key              = "example-application-key"
#         container_key                = "example-container-key"
#         jre_key                      = "example-jre-key"
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = 10
#     }
#
#     # Required for task_type="PERFORMANCE_TUNING"
#     performance_tuning_task_request {
#       #Optional
#       recording_duration_in_minutes = 10
#       targets {
#         #Optional
#         application_installation_key = "example-application-installation-key"
#         application_key              = "example-application-key"
#         container_key                = "example-container-key"
#         jre_key                      = "example-jre-key"
#         managed_instance_id          = var.managed_instance_ocid
#       }
#       waiting_period_in_minutes = 10
#     }
#
#     # Required for task_type="REMOVE_INSTALLATION_SITE"
#     remove_installation_site_task_request {
#       #Optional
#       installation_sites {
#         #Optional
#         installation_key    = "example-jvm-installation-key"
#         managed_instance_id = var.managed_instance_ocid
#       }
#     }
#
#     # Required for task_type="SCAN_JAVA_SERVER"
#     scan_java_server_task_request {
#       #Optional
#       managed_instance_ids = []
#     }
#
#     # Required for task_type="SCAN_LIBRARY"
#     scan_library_task_request {
#       #Optional
#       dynamic_scan_duration_in_minutes = 10
#       is_dynamic_scan                  = false
#       managed_instance_ids             = []
#     }
#
#   }
# }

data "oci_jms_task_schedules" "test_task_schedules" {

  fleet_id = var.fleet_ocid
  #Optional
  id                          = "example-task-schedule-id"
  managed_instance_id         = var.managed_instance_ocid
  name                        = "JFR"
  task_schedule_name_contains = "JFR"
}
