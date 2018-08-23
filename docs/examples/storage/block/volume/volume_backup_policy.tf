variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

variable "DBSize" {
  default = "50" // size in GBs, min: 50, max 16384
}

variable "availability_domain" {
  default = 3
}

data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_volume" "t" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "-tf-volume"
  size_in_gbs         = "${var.DBSize}"
}

data "oci_core_volume_backup_policies" "test_volume_backup_policies" {}

output "policies" {
  value = "${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies}"
}

data "oci_core_volume_backup_policies" "test_volume_backup_policies_silver" {
  filter {
    name   = "display_name"
    values = ["silver"]
  }
}

output "silver_policy_id" {
  value = "${data.oci_core_volume_backup_policies.test_volume_backup_policies_silver.volume_backup_policies.0.id}"
}

resource "oci_core_volume_backup_policy_assignment" "test_backup_policy_assignment" {
  asset_id  = "${oci_core_volume.t.id}"
  policy_id = "${data.oci_core_volume_backup_policies.test_volume_backup_policies_silver.volume_backup_policies.0.id}"
}

data "oci_core_volume_backup_policy_assignments" "test_backup_policy_assignments" {
  asset_id = "${oci_core_volume.t.id}"

  filter {
    name   = "id"
    values = ["${oci_core_volume_backup_policy_assignment.test_backup_policy_assignment.id}"]
  }
}

output "test_backup_policy_assignments" {
  value = "${data.oci_core_volume_backup_policy_assignments.test_backup_policy_assignments.volume_backup_policy_assignments}"
}
