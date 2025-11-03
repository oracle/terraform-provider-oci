
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_runbook" "test_runbook" {
  compartment_id = "${var.compartment_id}"
  display_name   = "displayName"
  operation      = "PATCH"
  platform       = "Oracle Linux"
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
      type = "PARALLEL_RESOURCE_GROUP"
    }
    tasks {
      step_name = "StepName"
      task_record_details {
        execution_details {
          command        = "pwd"
          execution_type = "SCRIPT"
        }
        scope = "LOCAL"
      }
    }
    version = "1"
  }
}
variable "creds_display_name" { default = "tersi-test-credential" }
variable "creds_ocid" { default = "ocid1.famsplatformconfiguration.oc1.." }

resource "oci_fleet_apps_management_runbook_version" "test_runbook_version" {
  defined_tags = "${map("Oracle-Tags.CreatedBy", "value")}"
  execution_workflow_details {
    workflow {
      group_name = "Parallel_resource_group"
      steps {
        step_name = "stepName"
        type      = "TASK"
      }
      type = "PARALLEL_RESOURCE_GROUP"
    }
  }
  freeform_tags = {
    "bar-key" = "value"
  }
  groups {
    name = "Parallel_resource_group"
    properties {
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
    type = "PARALLEL_RESOURCE_GROUP"
  }
  rollback_workflow_details {
    scope = "ACTION_GROUP"
    workflow {
      group_name = "Parallel_resource_group"
      steps {
        step_name = "stepNameRollback"
        type      = "TASK"
      }
      type = "PARALLEL_RESOURCE_GROUP"
    }
  }
  runbook_id = "${oci_fleet_apps_management_runbook.test_runbook.id}"
  tasks {
    step_name = "stepName"
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
