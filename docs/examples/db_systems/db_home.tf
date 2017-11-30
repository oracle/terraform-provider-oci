variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_id" {}
variable "database_admin_password" {}
variable "database_character_set" {}
variable "database_db_name" {}
variable "database_db_workload" {}
variable "database_ncharacter_set" {}
variable "database_pdb_name" {}
variable "db_system_id" {}
variable "db_version" {}
variable "display_name" {}


provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}

resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"

		#Optional
		character_set = "${var.database_character_set}"
		db_workload = "${var.database_db_workload}"
		ncharacter_set = "${var.database_ncharacter_set}"
		pdb_name = "${var.database_pdb_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"

	#Optional
	display_name = "${var.display_name}"
}


data "oci_database_db_homes" "testDBHomes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${var.db_system_id}"
}
