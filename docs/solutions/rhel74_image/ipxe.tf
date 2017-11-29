# DO NOT ALTER THIS FILE

resource "oci_core_instance" "ipxe_node" {
  availability_domain = "${data.oci_identity_availability_domains.ad.availability_domains.0.name}"
  compartment_id      = "${data.oci_identity_compartments.compartment.compartments.0.id}"
  display_name        = "${var.ipxe_instance["name"]}"
  image               = "${data.oci_core_images.image.images.0.id}" 
  shape               = "${var.ipxe_instance["shape"]}"
  subnet_id           = "${data.oci_core_subnets.subnet.subnets.0.id}"
  hostname_label      = "${var.ipxe_instance["hostname"]}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(data.external.ipxe_gen.result["shell"]))}"
  }
}

resource "null_resource" "delete_ipxe" {
  triggers {
    ipxe_node_id = "${oci_core_instance.ipxe_node.id}"
  }
#  depends_on = [ "oci_core_instance.ipxe_node" ]
  provisioner "local-exec" {
    command = "rm -rf ./ipxe.sh"
  }
}

resource "null_resource" "delete_ipxe_destroy" {
  provisioner "local-exec" {
    when = "destroy"
    command = "rm -rf ./ipxe.sh"
  }
}

