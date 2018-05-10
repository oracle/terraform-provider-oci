resource "oci_core_instance" "MongoP" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Primary MongoDB Server"
  hostname_label      = "MongoDB-Pri"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"  
  shape               = "${var.InstanceShape1}"
  subnet_id           = "${var.SubnetOCID}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}

resource "oci_core_instance" "MongoR1" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Replica MongoDB Server 1"
  hostname_label      = "MongoDB-R1"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"
  shape               = "${var.InstanceShape1}"
  subnet_id           = "${var.SubnetOCID}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}

resource "oci_core_instance" "MongoR2" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Replica MongoDB Server 2"
  hostname_label      = "MongoDB-R2"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"
  shape               = "${var.InstanceShape1}"
  subnet_id           = "${var.SubnetOCID}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}

resource "oci_core_instance" "Pritunl1" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Pritunl Server 1"
  hostname_label      = "Pritunl1"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"
  shape               = "${var.InstanceShape}"
  subnet_id           = "${var.SubnetOCID}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}

resource "oci_core_instance" "Pritunl2" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Pritunl Server 2"
  hostname_label      = "Pritunl2"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"
  shape               = "${var.InstanceShape}"
  subnet_id           = "${var.SubnetOCID}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}

resource "oci_core_instance" "Pritunllink" {
  count = "1"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "Pritunl Link"
  hostname_label      = "PTlink"
  image               = "${lookup(data.oci_core_images.OLImageOCID.images[3], "id")}"
  shape               = "${var.InstanceShape}"
  subnet_id           = "${var.SubnetOCID}"

  create_vnic_details {
    subnet_id         = "${var.SubnetOCID}"
    skip_source_dest_check = "true"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "8m"
  }
}
