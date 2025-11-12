#########################################################################################################
# Variables for OCI Apps Management Runbook
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Runbook will be created."
}

variable "display_name" {
  type        = string
  description = "Human-friendly display name of the Runbook."
}

variable "description" {
  type        = string
  description = "Detailed description of the Runbook’s purpose and operations."
  default     = null
}

variable "estimated_time" {
  type        = string
  description = "Estimated duration of the Runbook in ISO-8601 format (e.g., PT30M, PT1H, PT2H, PT3H)."
  default     = null
}

variable "is_default" {
  type        = bool
  description = "Indicates if this Runbook is a default system-defined template."
  default     = false
}

variable "is_sudo_access_needed" {
  type        = bool
  description = "Specifies whether sudo privileges are required for task execution."
  default     = false
}

variable "operation" {
  type        = string
  description = <<EOT
Specifies the operation category this Runbook performs.
Possible values:
  - DISCOVERY   : Performs discovery or audit operations.
  - PROVISION   : Handles provisioning and setup tasks.
  - PATCH       : Executes patching or maintenance workflows.
EOT
  default     = "PATCH"
}

variable "os_type" {
  type        = string
  description = <<EOT
Operating system type for which this Runbook is applicable.
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
Platform or software family targeted by the Runbook.
Observed values include:
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
  - <any custom platforms/ products>
EOT
  default     = "Oracle Linux"
}

variable "group_type" {
  type        = string
  description = <<EOT
Type of Runbook group defining execution behavior for steps.
Possible values:
  - ROLLING_RESOURCE_GROUP : Executes tasks sequentially on rolling targets.
  - PARALLEL_RESOURCE_GROUP : Executes resource-level steps in parallel.
  - PARALLEL_TASK_GROUP : Executes individual tasks in parallel.
EOT
  default     = "ROLLING_RESOURCE_GROUP"
}

variable "execution_type" {
  type        = string
  description = <<EOT
Execution type for task execution details.
Possible values:
  - SCRIPT     : Executes shell or OS-level commands.
  - TERRAFORM  : Executes Terraform configurations or catalog references.
EOT
  default     = "SCRIPT"
}

variable "action_on_failure" {
  type        = string
  description = <<EOT
Determines what happens when a step or group fails.
Possible values:
  - ABORT : Stop further execution.
  - CONTINUE : Proceed with next step or group.
EOT
  default     = "ABORT"
}

variable "scope" {
  type        = string
  description = <<EOT
Scope of task execution.
Possible values:
  - LOCAL  : Runs the task in the same resource context.
  - TARGET : Runs the task on the target system.
EOT
  default     = "LOCAL"
}

variable "defined_tags" {
  type        = map(string)
  description = "Predefined tag namespaces and key-value pairs for the Runbook."
  default     = {}
}

variable "freeform_tags" {
  type        = map(string)
  description = "Custom key-value tags applied to the Runbook."
  default     = {}
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_runbook" "test_runbook" {
  compartment_id        = var.compartment_id
  display_name          = var.display_name
  description           = var.description
  estimated_time        = var.estimated_time
  is_default            = var.is_default
  is_sudo_access_needed = var.is_sudo_access_needed
  operation             = var.operation
  os_type               = var.os_type
  platform              = var.platform
  defined_tags          = var.defined_tags
  freeform_tags         = var.freeform_tags

  runbook_version {
    version   = "1"
    is_latest = false

    groups {
      name = "Rolling_resource_group"
      type = var.group_type
      properties {
        action_on_failure = var.action_on_failure
      }
    }

    execution_workflow_details {
      workflow {
        group_name = "Rolling_resource_group"
        type       = var.group_type
        steps {
          step_name = "StepName"
          type      = "TASK"
        }
      }
    }

    tasks {
      step_name = "StepName"
      task_record_details {
        name        = "StepName"
        description = "Sample task execution"
        os_type     = var.os_type
        platform    = var.platform
        scope       = var.scope

        execution_details {
          execution_type                  = var.execution_type
          command                         = "ls"
          is_locked                       = false
          is_executable_content           = false
          is_read_output_variable_enabled = false
        }

        properties {
          num_retries        = 2
          timeout_in_seconds = 1800
        }

        is_copy_to_library_enabled = true
        is_discovery_output_task   = false
        is_apply_subject_task      = false
      }
    }
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_runbook" "test_runbook" {
  runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_runbooks" "test_runbooks" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  operation      = var.operation
  platform       = var.platform
  state          = "ACTIVE"
  type           = "USER_DEFINED"

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_runbook.test_runbook.id]
  }
}

#########################################################################################################
# Computed Attributes (Available in Data Source)
#########################################################################################################
# - id                      : Unique OCID of the Runbook.
# - compartment_id           : OCID of the compartment that owns the Runbook.
# - resource_region          : OCI region where the Runbook is deployed (e.g., us-ashburn-1).
# - display_name             : Display name of the Runbook.
# - description              : Description of the Runbook.
# - type                     : Runbook category.
#       Possible values:
#         - SYSTEM_DEFINED
#         - USER_DEFINED
# - operation                : Type of operation this Runbook supports (DISCOVERY / PROVISION / PATCH / DELETE / UPDATE).
# - os_type                  : Operating system type (LINUX / WINDOWS / GENERIC).
# - platform                 : Target platform (Oracle Linux / Oracle Java / Oracle Exadata OS / Oracle Clusterware).
# - is_default               : Whether the Runbook is a default system template.
# - is_sudo_access_needed    : Whether sudo privileges are required.
# - estimated_time           : Estimated execution duration (e.g., PT1H, PT2H, PT3H).
# - latest_version           : Version number of the latest published Runbook (if available).
# - has_draft_version        : Indicates whether a draft version exists.
# - lifecycle_state          : Lifecycle state of the Runbook.
#       Possible values:
#         - ACTIVE
#         - INACTIVE
#         - DELETED
# - lifecycle_details        : JSON-encoded details of substate or draft information.
# - runbook_version          : Object describing workflow details, tasks, and groups.
# - rollback_workflow_details: Workflow executed during rollback operations.
# - freeform_tags            : Map of free-form tags.
# - defined_tags             : Map of defined tags with namespaces.
# - system_tags              : Map of system-assigned tags.
#########################################################################################################
