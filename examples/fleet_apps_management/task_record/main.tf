#########################################################################################################
# Variables for OCI Apps Management Task Record
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Task Record will be created."
}

variable "catalog_id" {
  type        = string
  description = "The OCID of the catalog item associated with the task."
}

variable "credential_id" {
  type        = string
  description = "The OCID of the credential or platform configuration used for task execution."
}

variable "display_name" {
  type        = string
  description = "Display name of the Task Record."
  default     = "displayName"
}

#########################################################################################################
# Enumerations
#########################################################################################################

variable "execution_type" {
  type        = string
  description = <<EOT
Defines how the task will be executed.

Possible values:
  - SCRIPT
  - TERRAFORM    
  - API       
EOT
  default = "SCRIPT"
}

variable "source_type" {
  type        = string
  description = <<EOT
Defines where the task content is sourced from.

Possible values:
  - OBJECT_STORAGE_BUCKET : Content resides in an Object Storage bucket.
  - PAR_URL               : Content is referenced via a pre-authenticated request URL.
  - INLINE                : Content provided inline within the configuration.
EOT
  default = "OBJECT_STORAGE_BUCKET"
}

variable "operation" {
  type        = string
  description = <<EOT
Specifies the functional purpose of this task.

Possible values:
  - PATCH       : Task performs patching or updates.
  - PROVISION   : Task provisions new resources.
  - DISCOVERY   : Task discovers inventory or environment details.
  - VALIDATION  : Task validates configuration or state.
EOT
  default = "PATCH"
}

variable "os_type" {
  type        = string
  description = <<EOT
Defines the operating system type the task applies to.

Possible values:
  - LINUX    : Applicable to Linux-based systems.
  - WINDOWS  : Applicable to Microsoft Windows systems.
  - GENERIC  : Applicable across OS types.
EOT
  default = "LINUX"
}

variable "platform" {
  type        = string
  description = <<EOT
Defines the target product platform for which this task is designed.

Possible values:
  - Oracle Linux
  - Oracle Java
  - Oracle WebLogic Server
  - Oracle Database
  - Oracle Fusion Middleware
  - Oracle HTTP Server
  - Oracle Exadata OS
  - Oracle Exadata Database Service
  - Microsoft Windows
EOT
  default = "Oracle Java"
}

variable "task_scope" {
  type        = string
  description = <<EOT
Scope of task execution.
Possible values:
  - LOCAL  
  - SHARED 
EOT
  default     = "LOCAL"
}

variable "state" {
  type        = string
  description = <<EOT
Defines the lifecycle state of the Task Record.

Possible values:
  - ACTIVE
  - INACTIVE
  - DELETED
  - DELETING
  - FAILED
  - UPDATING
EOT
  default = "ACTIVE"
}

variable "type" {
  type        = string
  description = <<EOT
Defines whether the Task Record is user-defined or Oracle-defined.

Possible values:
  - STRING
  - OUTPUT_VARIABLE
  - FILE
EOT
  default = "USER_DEFINED"
}

variable "input_variable_type" {
  type        = string
  description = "input variable type"
  default = "STRING"
}


#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_task_record" "test_task_record" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  description    = "description"

  details {
    execution_details {
      command              = "pwd"
      execution_type        = var.execution_type
      is_executable_content = false
      is_locked             = false
      system_variables      = ["systemVariables"]

      content {
        namespace   = "namespace"
        bucket      = "bucket"
        object      = "object"
        checksum    = "checksum"
        source_type = var.source_type
      }

      variables {
        input_variables {
          description = "Input variable description"
          name        = "name1"
          type        = var.input_variable_type
        }
        output_variables = ["outputVariable1"]
      }
    }

    is_apply_subject_task    = false
    is_discovery_output_task = false
    operation                = var.operation
    os_type                  = var.os_type
    platform                 = var.platform
    scope                    = var.task_scope

    properties {
      num_retries        = 10
      timeout_in_seconds = 600
    }
  }

  defined_tags = {
    "Oracle-Tags.CreatedBy" = "value"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_task_record" "test_task_record" {
  task_record_id = oci_fleet_apps_management_task_record.test_task_record.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_task_records" "test_task_records" {
  compartment_id = var.compartment_id
  state          = var.state
  type           = var.type

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_task_record.test_task_record.id]
  }
}

#########################################################################################################
# Computed Attributes (Returned by Data Source)
#########################################################################################################

# id                    : Unique OCID of the Task Record.
# compartment_id         : The compartment OCID where the task record is created.
# display_name           : Display name of the task.
# description            : Description of the task record.
# details                : Full nested structure containing execution details and metadata.
# execution_details      : Command, content source, execution type, variables, and locking flags.
# operation              : PATCH / PROVISION / DISCOVERY / VALIDATION.
# os_type                : LINUX / WINDOWS / GENERIC.
# platform               : Oracle product or target platform.
# scope                  : LOCAL / GLOBAL / TARGET.
# lifecycle_state        : ACTIVE / INACTIVE / DELETED.
# time_created           : RFC3339 timestamp when the task was created.
# time_updated           : RFC3339 timestamp when the task was last updated.
# defined_tags           : Oracle or user-defined namespace tags.
# freeform_tags          : Simple key-value tags defined by the user.
# system_tags            : Oracle-assigned metadata tags.
#########################################################################################################
