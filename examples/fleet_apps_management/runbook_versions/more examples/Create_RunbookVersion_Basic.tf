
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

resource "oci_fleet_apps_management_runbook_version" "test_runbook_version" {
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
  groups {
    name = "Parallel_resource_group"
    type = "PARALLEL_RESOURCE_GROUP"
  }
  runbook_id = "${oci_fleet_apps_management_runbook.test_runbook.id}"
  tasks {
    step_name = "stepName"
    task_record_details {
      execution_details {
        command        = "pwd"
        execution_type = "SCRIPT"
      }
      scope = "LOCAL"
    }
  }
}
