
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
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  description    = "description2"
  details {
    execution_details {
      command = "ls"
      content {
        bucket      = "bucket2"
        checksum    = "checksum2"
        namespace   = "namespace2"
        object      = "object2"
        source_type = "OBJECT_STORAGE_BUCKET"
      }
      execution_type        = "SCRIPT"
      is_executable_content = "true"
      is_locked             = "true"
      system_variables      = ["systemVariables2"]
      variables {
        input_variables {
          description = "description2"
          name        = "name2"
          type        = "OUTPUT_VARIABLE"
        }
        output_variables = ["outputVariables2"]
      }
    }
    is_apply_subject_task    = "false"
    is_discovery_output_task = "true"
    operation                = "Discovery"
    os_type                  = "WINDOWS"
    platform                 = "tersi-test-compatible-product2"
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
