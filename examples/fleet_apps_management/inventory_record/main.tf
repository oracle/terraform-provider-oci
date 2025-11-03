variable "region" {
  default = "us-ashburn-1"
}

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}



data "oci_fleet_apps_management_inventory_records" "test_inventory_records" {
  compartment_id            = "${var.compartment_id}"
  compartment_id_in_subtree = "false"
  fleet_id                  = "${var.fleet_id}"
  is_details_required       = "false"
  resource_id               = "${var.instance_id}"
}
variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Runbook will be created."
}

variable "fleet_id" {
  type        = string
  description = "The OCID of the Fleet associated with this Scheduler Execution."
}

variable "instance_id" {
  type        = string
  description = "OCID of a compute instance associated with fleet credentials."
}
