# Setting up a compute instance hosted in the public subnet in availability domain 1.
# We might run applications like a web-service on this instance.
resource "oci_core_instance" "MyInstanceInPublicSubnetAD1" {
  availability_domain = "${oci_core_subnet.MyPublicSubnetAD1.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyInstance"
  image = "${lookup(data.oci_core_images.ImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${oci_core_subnet.MyPublicSubnetAD1.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}

# Setting up a compute instance hosted in the private subnet in availability domain 1.
# We might run a database on this instances.
resource "oci_core_instance" "MyInstanceInPrivateSubnetAD1" {
  availability_domain = "${oci_core_subnet.MyPrivateSubnetAD1.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyInstance"
  image = "${lookup(data.oci_core_images.ImageOCID.images[0], "id")}"
  shape = "${var.InstanceShape}"
  subnet_id = "${oci_core_subnet.MyPrivateSubnetAD1.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }
}