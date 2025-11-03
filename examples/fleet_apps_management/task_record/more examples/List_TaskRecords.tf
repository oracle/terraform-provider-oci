
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }

data "oci_fleet_apps_management_task_records" "test_task_records" {
  compartment_id = "${var.compartment_id}"
  state          = "ACTIVE"
  type           = "USER_DEFINED"
}

