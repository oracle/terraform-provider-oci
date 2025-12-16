

data "oci_core_images" "instance_image" {
  compartment_id           = var.compartment_ocid
  operating_system         = var.image_operating_system
  operating_system_version = var.image_operating_system_version
  shape                    = var.instance_shape
  sort_by                  = "TIMECREATED"
  sort_order               = "DESC"
}

resource "oci_core_compute_capacity_reservation" "cr" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  instance_reservation_configs {
    instance_shape = var.instance_shape
    reserved_count = var.instance_count
    instance_shape_config {
      ocpus = "2"
      memory_in_gbs = "18"
    }
  }

  instance_reservation_configs {
    instance_shape = var.instance_shape
    reserved_count = var.instance_count
    instance_shape_config {
      ocpus = "3"
      memory_in_gbs = "22"
    }
  }
}

resource "oci_core_instance" "test_instance" {
  count                   = var.instance_count
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id          = var.compartment_ocid
  capacity_reservation_id = oci_core_compute_capacity_reservation.cr.id
  display_name            = "${var.instance_name_prefix}${count.index}"
  shape                   = var.instance_shape
  create_vnic_details {
    assign_public_ip = false
    subnet_id        = oci_core_subnet.test_subnet.id
  }
  shape_config {
    ocpus = "2"
    memory_in_gbs = "18"
  }
  source_details {
    source_type = "image"
    source_id = data.oci_core_images.instance_image.images[0].id
  }
}
