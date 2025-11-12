
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "fleet_id" { default = "ocid1.famsfleet.oc1." }
variable "runbook_id" { default = "ocid1.famsrunbook.oc1.i" }
variable "schedular_definition_id" { default = "ocid1.famsschedulerdefinition.oc1." }
variable "schedular_job_id" { default = "ocid1.famsschedulerjob.oc1." }

data "oci_fleet_apps_management_scheduler_executions" "test_scheduler_executions" {
  compartment_id                          = "${var.compartment_id}"
  compartment_id_in_subtree               = "false"
  display_name                            = "execution-tersi-schedule-1"
  resource_id                             = "${var.fleet_id}"
  runbook_id                              = "${var.runbook_id}"
  runbook_version_name                    = "1"
  scheduler_defintion_id                  = "${var.schedular_definition_id}"
  scheduler_job_id                        = "${var.schedular_job_id}"
  substate                                = "FAILED"
  time_scheduled_greater_than_or_equal_to = "2025-10-20T00:00:00.000Z"
  time_scheduled_less_than                = "2025-10-24T00:00:00.000Z"
}
