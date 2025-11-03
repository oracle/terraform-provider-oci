#########################################################################################################
# Variables for OCI Fleet Apps Management Scheduler Execution
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment that owns the Scheduler Executions."
}

variable "fleet_id" {
  type        = string
  description = "The OCID of the Fleet associated with this Scheduler Execution."
}

variable "runbook_id" {
  type        = string
  description = "The OCID of the Runbook associated with this Scheduler Execution."
}

variable "scheduler_definition_id" {
  type        = string
  description = "The OCID of the Scheduler Definition under which this execution is created."
}

variable "scheduler_job_id" {
  type        = string
  description = "The OCID of the Scheduler Job triggering this execution."
}

variable "display_name" {
  type        = string
  description = "The display name of the Scheduler Execution to filter on."
}


variable "substate" {
  type        = string
  description = <<EOT
Defines the detailed state of a Scheduler Execution.

Possible values:
  - QUEUED          : Execution is queued and waiting to start.
  - IN_PROGRESS     : Execution is currently running.
  - SUCCEEDED       : Execution completed successfully.
  - FAILED          : Execution failed to complete successfully.
  - CANCELED        : Execution was canceled before completion.
  - TIMED_OUT       : Execution exceeded the allowed time.
  - SKIPPED         : Execution was skipped due to conditions not met.
EOT
  default = "FAILED"
}

#########################################################################################################
# Time Range Variables
#########################################################################################################

variable "time_scheduled_greater_than_or_equal_to" {
  type        = string
  description = "Filter to return executions scheduled after or on this RFC3339 timestamp. e.g 2025-10-20T00:00:00.000Z"
}

variable "time_scheduled_less_than" {
  type        = string
  description = "Filter to return executions scheduled before this RFC3339 timestamp. e.g. 2025-10-24T00:00:00.000Z"
}

#########################################################################################################
# Data Source Definition (List)
#########################################################################################################

data "oci_fleet_apps_management_scheduler_executions" "test_scheduler_executions" {
  compartment_id                          = var.compartment_id
  compartment_id_in_subtree               = false
  display_name                            = var.display_name
  resource_id                             = var.fleet_id
  runbook_id                              = var.runbook_id
  runbook_version_name                    = "1"
  scheduler_defintion_id                  = var.scheduler_definition_id
  scheduler_job_id                        = var.scheduler_job_id
  substate                                = var.substate
  time_scheduled_greater_than_or_equal_to = var.time_scheduled_greater_than_or_equal_to
  time_scheduled_less_than                = var.time_scheduled_less_than
}

#########################################################################################################
# Computed Attributes (Returned by Data Source)
#########################################################################################################

# id                          : Unique OCID of the Scheduler Execution.
# compartment_id               : Compartment OCID of the execution.
# display_name                 : Display name of the execution record.
# resource_id                  : OCID of the associated resource (e.g., Fleet).
# runbook_id                   : OCID of the Runbook executed.
# runbook_version_name          : Version name of the executed Runbook.
# scheduler_definition_id       : OCID of the Scheduler Definition.
# scheduler_job_id              : OCID of the Scheduler Job.
# lifecycle_state               : Overall lifecycle state (ACTIVE / INACTIVE / DELETED).
# substate                      : Detailed execution substate (QUEUED / IN_PROGRESS / SUCCEEDED / FAILED / CANCELED / TIMED_OUT / SKIPPED).
# time_created                  : Timestamp when the execution was created.
# time_updated                  : Timestamp when the execution was last updated.
# time_scheduled                : Timestamp when the execution was scheduled to start.
# time_started                  : Timestamp when the execution actually began.
# time_finished                 : Timestamp when the execution completed.
# lifecycle_details              : Detailed message or reason for current execution state.
# defined_tags                  : Namespaced Oracle or user-defined tags.
# freeform_tags                 : Simple user-defined key-value tags.
# system_tags                   : Oracle-assigned metadata tags.
#########################################################################################################
