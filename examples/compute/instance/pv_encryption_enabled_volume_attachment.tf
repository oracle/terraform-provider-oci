// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Instance with encryption in transit enabled

resource "oci_core_instance" "test_instance_with_pv_encryption_in_transit" {
  count               = 2
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance${count.index}"
  shape               = var.instance_shape

  shape_config {
    ocpus = "${var.instance_ocpus}"
    memory_in_gbs = "${var.instance_shape_config_memory_in_gbs}"
  }

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "testinstance${count.index}"
  }

  source_details {
    source_type = "image"
    source_id   = data.oci_core_images.supported_shape_images.images[0]["id"]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }

  is_pv_encryption_in_transit_enabled = "true"
}

resource "oci_core_volume" "test_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestVolume"
}

resource "oci_core_volume_attachment" "test_volume_attachment" {
  count                               = 2
  attachment_type                     = "paravirtualized"
  instance_id                         = oci_core_instance.test_instance_with_pv_encryption_in_transit[count.index].id
  volume_id                           = oci_core_volume.test_volume.id
  display_name                        = "TestVolumeAttachment"
  is_read_only                        = true
  is_pv_encryption_in_transit_enabled = true
  is_shareable                        = "true"
}

# Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
data "oci_core_images" "supported_shape_images" {
  compartment_id   = var.tenancy_ocid
  shape            = var.instance_shape
  operating_system = "Oracle Linux"

  filter {
    name   = "launch_options.is_pv_encryption_in_transit_enabled"
    values = ["true"]
  }
}

