
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

data "oci_fleet_apps_management_runbook_versions" "test_runbook_versions" {
  compartment_id = "${var.compartment_id}"
  filter {
    name   = "id"
    values = ["${oci_fleet_apps_management_runbook_version.test_runbook_version.id}"]
  }
  id         = "${oci_fleet_apps_management_runbook_version.test_runbook_version.id}"
  name       = "TestRunbookTFP6"
  runbook_id = "${oci_fleet_apps_management_runbook.test_runbook.id}"
  state      = "ACTIVE"
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

resource "oci_fleet_apps_management_runbook_version" "test_runbook_version" {
  defined_tags = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  execution_workflow_details {
    workflow {
      group_name = "Parallel_resource_group"
      steps {
        step_name = "stepName2"
        type      = "TASK"
      }
      type = "PARALLEL_RESOURCE_GROUP"
    }
  }
  freeform_tags = {
    "Department" = "Accounting"
  }
  groups {
    name = "Parallel_resource_group"
    properties {
      action_on_failure = "CONTINUE"
      notification_preferences {
        should_notify_on_pause        = "true"
        should_notify_on_task_failure = "true"
        should_notify_on_task_success = "true"
      }
      pause_details {
        kind = "USER_ACTION"
      }
      pre_condition = "target.product.name == \"Oracle Linux Server\""
      run_on {
        condition = "target.product.name == \"Oracle Linux Server\""
        kind      = "SCHEDULED_INSTANCES"
      }
    }
    type = "PARALLEL_RESOURCE_GROUP"
  }
  rollback_workflow_details {
    scope = "TARGET"
    workflow {
      group_name = "Parallel_resource_group"
      steps {
        step_name = "stepNameRollback2"
        type      = "TASK"
      }
      type = "PARALLEL_RESOURCE_GROUP"
    }
  }
  runbook_id = "${oci_fleet_apps_management_runbook.test_runbook.id}"
  tasks {
    step_name = "stepName2"
    step_properties {
      action_on_failure = "CONTINUE"
      notification_preferences {
        should_notify_on_pause        = "true"
        should_notify_on_task_failure = "true"
        should_notify_on_task_success = "true"
      }
      pause_details {
        kind = "USER_ACTION"
      }
      pre_condition = "target.product.name == \"Oracle Weblogic Server\""
      run_on {
        condition = "target.product.name == \"Oracle Linux Server\""
        kind      = "SCHEDULED_INSTANCES"
      }
    }
    task_record_details {
      description = "description2"
      execution_details {
        command                         = "ls -la"
        execution_type                  = "SCRIPT"
        is_executable_content           = "true"
        is_locked                       = "true"
        is_read_output_variable_enabled = "false"
        variables {
          input_variables {
            description = "inputVarDescription2"
            name        = "inputVarName2"
            type        = "STRING"
          }
          output_variables = ["outputVariables2"]
        }
      }
      is_apply_subject_task      = "true"
      is_copy_to_library_enabled = "false"
      is_discovery_output_task   = "true"
      name                       = "name2"
      os_type                    = "LINUX"
      platform                   = "Oracle Database"
      properties {
        num_retries        = "11"
        timeout_in_seconds = "11"
      }
      scope = "LOCAL"
    }
  }
  tasks {
    step_name = "stepNameRollback2"
    step_properties {
      action_on_failure = "CONTINUE"
      notification_preferences {
        should_notify_on_pause        = "true"
        should_notify_on_task_failure = "true"
        should_notify_on_task_success = "true"
      }
      pause_details {
        kind = "USER_ACTION"
      }
      pre_condition = "target.product.name == \"Oracle Weblogic Server\""
      run_on {
        condition = "target.product.name == \"Oracle Linux Server\""
        kind      = "SCHEDULED_INSTANCES"
      }
    }
    task_record_details {
      description = "description2"
      execution_details {
        command                         = "ls -la"
        execution_type                  = "SCRIPT"
        is_executable_content           = "true"
        is_locked                       = "true"
        is_read_output_variable_enabled = "false"
        variables {
          input_variables {
            description = "inputVarDescription2"
            name        = "inputVarName2"
            type        = "STRING"
          }
          output_variables = ["outputVariables2"]
        }
      }
      is_apply_subject_task      = "true"
      is_copy_to_library_enabled = "false"
      is_discovery_output_task   = "true"
      name                       = "name2"
      os_type                    = "LINUX"
      platform                   = "Oracle Database"
      properties {
        num_retries        = "11"
        timeout_in_seconds = "11"
      }
      scope = "LOCAL"
    }
  }
}
