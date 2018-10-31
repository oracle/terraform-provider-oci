# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

// Gets the detail of the vault.
data "oci_kms_vault" "test_vault" {
  #Required
  vault_id = "${var.vault_id}"
}

/*
//if want to create a new vault
resource "oci_kms_vault" "test_vault" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.vault_display_name}"
	vault_type = "${var.vault_vault_type}"
}
*/

// Gets the list of keys in the compartment and vault.
data "oci_kms_keys" "test_keys" {
  #Required
  compartment_id      = "${var.compartment_id}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

  filter {
    name   = "display_name"
    values = ["${var.key_display_name}"]
  }
}

data "oci_core_volumes" "test_volumes" {
  compartment_id = "${var.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_core_volume.my_volume.id}"]
  }
}
