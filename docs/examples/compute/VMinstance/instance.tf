variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

# Note that the difference between launching a VM and BM instance is just a different shape name.

resource "baremetal_core_instance" "a_TF_managed_BM_instance" {
  availability_domain = "Uocm:PHX-AD-1"
  compartment_id = "${var.compartment_ocid}"
  display_name = "a_TF_managed_BM_instance"
  image = "ocid1.image.oc1.phx.aaaaaaaaclseho77fcdfgejstt2bflkugcx5waa6bhconbokvhdp3qw7txlq"
  shape = "VM.Standard1.4"
  subnet_id = "ocid1.subnet.oc1.phx.aaaaaaaa2sn6ilxcsyflkc6ievfqvmie6fs75uypzgpmmykkgs3oq52jknla"
  metadata {
    ssh_authorized_keys = "ssh-rsa AAAAB3NaC1yc2EAAAADAQABAAABAQCYx1...gQ4RjNLLwD09i4zHW4cON"
  }
}
