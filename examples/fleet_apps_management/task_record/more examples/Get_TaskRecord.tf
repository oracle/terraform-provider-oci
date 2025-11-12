
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}



data "oci_fleet_apps_management_task_record" "test_task_record" {
  task_record_id = "${oci_fleet_apps_management_task_record.test_task_record.id}"
}
variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_task_record" "test_task_record" {
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  description    = "description2"
  details {
    execution_details {
      command = "ls"
      content {
        bucket      = "bucket"
        checksum    = "checksum"
        namespace   = "namespace"
        object      = "object"
        source_type = "OBJECT_STORAGE_BUCKET"
      }
      execution_type        = "SCRIPT"
      is_executable_content = "true"
      is_locked             = "true"
      system_variables      = ["systemVariables"]
      variables {
        input_variables {
          description = "description"
          name        = "name"
          type        = "OUTPUT_VARIABLE"
        }
        output_variables = ["outputVariables"]
      }
    }
    is_apply_subject_task    = "false"
    is_discovery_output_task = "true"
    operation                = "Discovery"
    os_type                  = "WINDOWS"
    platform                 = "test-compatible-product2"
    properties {
      num_retries        = "11"
      timeout_in_seconds = "11"
    }
    scope = "SHARED"
  }
  display_name = "displayName2"
  freeform_tags = {
    "bar-key" = "value"
  }
}
