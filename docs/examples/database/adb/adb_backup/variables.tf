variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

# The autonomous Database should be pre-configured for backups. Refer:
# https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/atpbackingup.htm#creatingbucket
variable "autonomous_database_id" { default = "" }
variable "autonomous_database_backup_display_name" { default = "Monthly Backup" }

provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}
