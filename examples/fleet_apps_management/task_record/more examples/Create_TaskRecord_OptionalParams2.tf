
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}


variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "compartment_id_for_update" { default = "ocid1.compartment.oc1.." }
variable "catalog_id" { default = "ocid1.famscatalogitem.oc1." }
variable "credential" { default = "ocid1.famsplatformconfiguration.oc1." }
variable "credential_updated" { default = "ocid1.famsplatformconfiguration.oc1." }
variable "credential_name" { default = "test-credential" }
variable "credential_name_for_update" { default = "test-credential2" }

resource "oci_fleet_apps_management_task_record" "test_task_record" {
  compartment_id = "${var.compartment_id_for_update}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "value")}"
  description    = "description"
  details {
    execution_details {
      command = "pwd"
      content {
        bucket      = "bucket"
        checksum    = "checksum"
        namespace   = "namespace"
        object      = "object"
        source_type = "OBJECT_STORAGE_BUCKET"
      }
      execution_type        = "SCRIPT"
      is_executable_content = "false"
      is_locked             = "false"
      system_variables      = ["systemVariables"]
      variables {
        input_variables {
          description = "description1"
          name        = "name1"
          type        = "STRING"
        }
        output_variables = ["outputVariables"]
      }
    }
    is_apply_subject_task    = "false"
    is_discovery_output_task = "false"
    operation                = "Patch"
    os_type                  = "LINUX"
    platform                 = "Oracle Java"
    properties {
      num_retries        = "10"
      timeout_in_seconds = "10"
    }
    scope = "LOCAL"
  }
  display_name = "displayName"
  freeform_tags = {
    "bar-key" = "value"
  }
}
