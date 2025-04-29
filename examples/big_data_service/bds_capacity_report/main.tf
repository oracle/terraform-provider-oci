variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}
variable "shape" {
  default = "VM.Standard2.2"
}

resource "oci_bds_bds_capacity_report" "test_bds_capacity_report" {
    compartment_id = "${var.compartment_id}"
    shape_availabilities {
        shape = "${var.shape}"
    }
}

output "capacity_report" {
    value = oci_bds_bds_capacity_report.test_bds_capacity_report
}
