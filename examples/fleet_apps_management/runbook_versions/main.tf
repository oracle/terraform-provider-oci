#########################################################################################################
# Variables for OCI Apps Management Runbook Version
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Runbook Version will be created."
}

variable "runbook_id" {
  type        = string
  description = "The OCID of the Runbook to which this version belongs."
}

variable "version_name" {
  type        = string
  description = "Version identifier or label for the Runbook Version (e.g., 'v1', '2', 'TestRunbookV3')."
  default     = "1"
}

variable "execution_type" {
  type        = string
  description = <<EOT
Specifies the execution type of task execution details.
Possible values:
  - SCRIPT     : Executes shell or OS-level commands.
  - TERRAFORM  : Executes Terraform configurations or catalog references.
EOT
  default     = "SCRIPT"
}

variable "os_type" {
  type        = string
  description = <<EOT
Operating system type applicable for the Runbook task.
Possible values:
  - LINUX
  - WINDOWS
  - GENERIC
EOT
  default     = "LINUX"
}

variable "platform" {
  type        = string
  description = <<EOT
Target platform or product family for which the task applies.
Possible values:
  - Oracle Exadata OS
  - Oracle Clusterware
  - Oracle Exadata Database Service
  - Microsoft Windows
  - Oracle HTTP Server
  - Oracle Base Database Service
  - Oracle Linux
  - Oracle Java
  - Oracle Weblogic Server
  - Oracle Weblogic Server and Java
  - Oracle Fusion Middleware
  - Oracle Database
EOT
  default     = "Oracle Linux"
}

variable "group_type" {
  type        = string
  description = <<EOT
Defines how task groups are executed.
Possible values:
  - ROLLING_RESOURCE_GROUP  : Sequential rolling execution.
  - PARALLEL_RESOURCE_GROUP : Parallel execution across resources.
  - PARALLEL_TASK_GROUP     : Parallel execution across tasks.
EOT
  default     = "PARALLEL_RESOURCE_GROUP"
}

variable "action_on_failure" {
  type        = string
  description = <<EOT
Defines the behavior when a task or group fails.
Possible values:
  - ABORT    : Stop execution on failure.
  - CONTINUE : Proceed to next task or group.
EOT
  default     = "CONTINUE"
}

variable "scope" {
  type        = string
  description = <<EOT
Execution scope for the task.
Possible values:
  - LOCAL  : Executes locally on the resource.
  - TARGET : Executes remotely on a target system.
EOT
  default     = "LOCAL"
}

variable "pause_kind" {
  type        = string
  description = <<EOT
Type of pause inserted between steps.
Possible values:
  - USER_ACTION : Requires manual user confirmation to continue.
  - SYSTEM_PAUSE : Automated pause for system validation.
EOT
  default     = "USER_ACTION"
}

variable "run_on_kind" {
  type        = string
  description = <<EOT
Defines where and how the task should execute.
Possible values:
  - SCHEDULED_INSTANCES : Executes only on scheduled instances.
  - ALL_TARGETS         : Executes on all defined targets.
EOT
  default     = "SCHEDULED_INSTANCES"
}

variable "defined_tags" {
  type        = map(string)
  description = "Predefined namespaced key-value tags for the Runbook Version."
  default     = {}
}

variable "freeform_tags" {
  type        = map(string)
  description = "Simple key-value tags applied without predefined namespace."
  default     = {}
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_runbook_version" "test_runbook_version" {
  runbook_id    = var.runbook_id
  defined_tags  = var.defined_tags
  freeform_tags = var.freeform_tags

  # Workflow definition
  execution_workflow_details {
    workflow {
      group_name = "Parallel_resource_group"
      type       = var.group_type
      steps {
        step_name = "StepName"
        type      = "TASK"
      }
    }
  }

  # Group definition
  groups {
    name = "Parallel_resource_group"
    type = var.group_type
    properties {
      action_on_failure = var.action_on_failure
      pre_condition     = "target.product.name == \"Oracle Linux Server\""

      pause_details {
        kind = var.pause_kind
      }

      run_on {
        condition = "target.product.name == \"Oracle Linux Server\""
        kind      = var.run_on_kind
      }

      notification_preferences {
        should_notify_on_pause        = true
        should_notify_on_task_failure = true
        should_notify_on_task_success = true
      }
    }
  }

  # Rollback workflow
  rollback_workflow_details {
    scope = "TARGET"
    workflow {
      group_name = "Parallel_resource_group"
      type       = var.group_type
      steps {
        step_name = "stepNameRollback"
        type      = "TASK"
      }
    }
  }

  # Tasks definition
  tasks {
    step_name = "StepName"
    step_properties {
      action_on_failure = var.action_on_failure
      pre_condition     = "target.product.name == \"Oracle Weblogic Server\""

      pause_details {
        kind = var.pause_kind
      }

      run_on {
        condition = "target.product.name == \"Oracle Linux Server\""
        kind      = var.run_on_kind
      }

      notification_preferences {
        should_notify_on_pause        = true
        should_notify_on_task_failure = true
        should_notify_on_task_success = true
      }
    }

    task_record_details {
      name        = "TaskExample"
      description = "Sample Runbook Version task"
      os_type     = var.os_type
      platform    = var.platform
      scope       = var.scope

      execution_details {
        command                         = "ls -la"
        execution_type                  = var.execution_type
        is_executable_content           = true
        is_locked                       = true
        is_read_output_variable_enabled = false

        variables {
          input_variables {
            name        = "inputVarName"
            description = "Input variable description"
            type        = "STRING"
          }
          output_variables = ["outputVar"]
        }
      }

      properties {
        num_retries        = 3
        timeout_in_seconds = 1800
      }

      is_copy_to_library_enabled = false
      is_discovery_output_task   = false
      is_apply_subject_task      = true
    }
  }
  tasks {
    step_name = "stepNameRollback"
    step_properties {
      action_on_failure = "ABORT"
      notification_preferences {
        should_notify_on_pause        = "false"
        should_notify_on_task_failure = "false"
        should_notify_on_task_success = "false"
      }
      pause_details {
        kind = "USER_ACTION"
      }
      pre_condition = "target.product.name == \"Oracle Weblogic Server\""
      run_on {
        condition = "target.product.name == \"Oracle Weblogic Server\""
        kind      = "SCHEDULED_INSTANCES"
      }
    }
    task_record_details {
      description = "description"
      execution_details {
        command                         = "pwd"
        execution_type                  = "SCRIPT"
        is_executable_content           = "false"
        is_locked                       = "false"
        is_read_output_variable_enabled = "false"
        variables {
          input_variables {
            description = "inputVarDescription"
            name        = "inputVarName"
            type        = "STRING"
          }
          output_variables = ["outputVariables"]
        }
      }
      is_apply_subject_task      = "false"
      is_copy_to_library_enabled = "false"
      is_discovery_output_task   = "false"
      name                       = "name"
      os_type                    = "WINDOWS"
      platform                   = "Oracle Fusion Middleware"
      properties {
        num_retries        = "10"
        timeout_in_seconds = "10"
      }
      scope = "LOCAL"
    }
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_runbook_version" "test_runbook_version" {
  runbook_version_id = oci_fleet_apps_management_runbook_version.test_runbook_version.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_runbook_versions" "test_runbook_versions" {
  compartment_id = var.compartment_id
  runbook_id     = var.runbook_id
  name           = var.version_name
  state          = "ACTIVE"

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_runbook_version.test_runbook_version.id]
  }
}

#########################################################################################################
# Computed Attributes (Available in Data Source)
#########################################################################################################
# - id                     : Unique OCID of the Runbook Version.
# - runbook_id              : OCID of the parent Runbook.
# - compartment_id          : Compartment OCID of the Runbook Version.
# - name                    : Version name or label.
# - is_latest               : Whether this version is the latest published one.
# - lifecycle_state          : Lifecycle state of the version.
#       Possible values:
#         - ACTIVE
#         - INACTIVE
#         - DELETED
# - lifecycle_details        : JSON-encoded lifecycle details.
# - time_created             : RFC3339 timestamp of creation.
# - time_updated             : RFC3339 timestamp of last modification.
# - execution_workflow_details : Object representing workflow execution structure.
# - rollback_workflow_details  : Object representing rollback workflows.
# - groups                  : List of execution groups with their properties.
# - tasks                   : Task list containing step details, properties, and execution info.
# - defined_tags            : Predefined namespace tags.
# - freeform_tags           : Custom key-value tags.
# - system_tags             : System-assigned metadata tags.
#########################################################################################################
