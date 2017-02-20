/* This configuration creates and attaches a block volume to an instance. 
   Valid volumes sizes are "2097152" (2TB) and "262144" (256GB)

   Note that we're using the volume OCID of the newly created volume in 
   the volume attach resource via the ${baremetal_core_volume.block-1.volume_id} 
   variable.
*/

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key = "${var.private_key}"
}

resource "baremetal_core_volume" "block-1" {
  availability_domain = "Uocm:PHX-AD-1"
  compartment_id = "${var.compartment_ocid}"
  display_name = "block-1"
  size_in_mbs = "2097152"
}

/*resource "baremetal_core_volume_attachment" "block-1-attach" {
    attachment_type = "iscsi"
    compartment_id = "${var.compartment_ocid}"
    instance_id = "ocid1.instance.oc1.phx.abyhqljslboyszgpx2csnmm3wx2nyfpusu5l7hbunuacbifalc3a6v2xc22q"
    volume_id = "${baremetal_core_volume.block-1.volume_id}"
}*/

#output "block-1.id" { 
#  value = "${baremetal_core_volume.block-1.id}" }
