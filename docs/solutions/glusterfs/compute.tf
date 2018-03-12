resource "oci_core_instance" "GlusterServerInstance" {
  availability_domain = "${oci_core_subnet.SubnetAD1.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-server1"
  hostname_label = "glusterfs-server1"
  image = "${var.ServerInstanceImageOCID[var.region]}"
  shape = "${var.ServerInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD1.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ServerBootStrapFile))}"
  }
}

resource "oci_core_instance" "GlusterServerInstance2" {
  availability_domain = "${oci_core_subnet.SubnetAD2.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-server2"
  hostname_label = "glusterfs-server2"
  image = "${var.ServerInstanceImageOCID[var.region]}"
  shape = "${var.ServerInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD2.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ServerBootStrapFile))}"
  }
}

resource "oci_core_instance" "GlusterServerInstance3" {
  availability_domain = "${oci_core_subnet.SubnetAD3.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-server3"
  hostname_label = "glusterfs-server3"
  image = "${var.ServerInstanceImageOCID[var.region]}"
  shape = "${var.ServerInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD3.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ServerBootStrapFile))}"
  }
}


resource "oci_core_instance" "GlusterClientInstance" {
  availability_domain = "${oci_core_subnet.SubnetAD1.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-client1"
  hostname_label = "glusterfs-client1"
  image = "${var.ClientInstanceImageOCID[var.region]}"
  shape = "${var.ClientInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD1.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ClientBootStrapFile))}"
  }
}

resource "oci_core_instance" "GlusterClientInstance2" {
  availability_domain = "${oci_core_subnet.SubnetAD2.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-client2"
  hostname_label = "glusterfs-client2"
  image = "${var.ClientInstanceImageOCID[var.region]}"
  shape = "${var.ClientInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD2.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ClientBootStrapFile))}"
  }
}

resource "oci_core_instance" "GlusterClientInstance3" {
  availability_domain = "${oci_core_subnet.SubnetAD3.availability_domain}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "glusterfs-client3"
  hostname_label = "glusterfs-client3"
  image = "${var.ClientInstanceImageOCID[var.region]}"
  shape = "${var.ClientInstanceShape}"
  subnet_id = "${oci_core_subnet.SubnetAD3.id}"
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data = "${base64encode(file(var.ClientBootStrapFile))}"
  }
}

