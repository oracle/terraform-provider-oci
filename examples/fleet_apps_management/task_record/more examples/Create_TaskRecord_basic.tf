
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "catalog_id" { default = "ocid1.famscatalogitem.oc1." }
variable "credential" { default = "ocid1.famsplatformconfiguration." }
variable "credential_name" { default = "test-credential" }

resource "oci_fleet_apps_management_task_record" "test_task_record" {
  compartment_id = "${var.compartment_id}"
  description    = "description"
  details {
    execution_details {
      command        = "pwd"
      execution_type = "SCRIPT"
    }
    os_type = "LINUX"
    properties {
      num_retries        = "10"
      timeout_in_seconds = "10"
    }
    scope = "LOCAL"
  }
  display_name = "displayName"
}
