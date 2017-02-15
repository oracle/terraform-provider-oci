variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}


provider baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key = "${var.private_key}"
}


resource "baremetal_core_instance" "a_TF_managed_instance" {
  availability_domain = "<some AD>"
  compartment_id = "${var.compartment}"
  display_name = "a_TF_managed_instance"
  image = "<some image ocid>"
  shape = "BM.DenseIO1.36"
  subnet_id = "<some subnet ocid"
  metadata {
    ssh_authorized_keys = "<some public key>"
  }
}
