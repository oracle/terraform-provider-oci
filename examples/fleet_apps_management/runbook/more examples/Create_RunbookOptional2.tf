
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id_for_update" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_runbook" "test_runbook" {
  compartment_id = "${var.compartment_id_for_update}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "value")}"
  description    = "description"
  display_name   = "displayName"
  estimated_time = "PT1H"
  freeform_tags = {
    "bar-key" = "value"
  }
  is_default            = "false"
  is_sudo_access_needed = "false"
  operation             = "PATCH"
  os_type               = "LINUX"
  platform              = "Oracle Linux"
  runbook_version {
    execution_workflow_details {
      workflow {
        group_name = "Parallel_resource_group"
        steps {
          step_name = "StepName"
          type      = "TASK"
        }
        type = "PARALLEL_RESOURCE_GROUP"
      }
    }
    groups {
      name = "Parallel_resource_group"
      properties {
        action_on_failure = "ABORT"
      }
      type = "PARALLEL_RESOURCE_GROUP"
    }
    is_latest = "false"
    tasks {
      step_name = "StepName"
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
              description = "description"
              name        = "name"
              type        = "STRING"
            }
            output_variables = ["outputVariables"]
          }
        }
        is_apply_subject_task      = "false"
        is_copy_to_library_enabled = "false"
        is_discovery_output_task   = "false"
        name                       = "StepName"
        os_type                    = "LINUX"
        platform                   = "Oracle Linux"
        properties {
          num_retries        = "10"
          timeout_in_seconds = "1000"
        }
        scope = "LOCAL"
      }
    }
    version = "1"
  }
}
