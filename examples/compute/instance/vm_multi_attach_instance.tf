// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "vm_multi_attach_instance_shape" {
  default = "VM.Standard2.1"
}

resource "oci_core_volume" "test_block_volume_multi_attach_iscsi" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "test_attach_existing_volume_on_instance_launch_iscsi"
  size_in_gbs         = var.db_size
}

resource "oci_core_volume" "test_block_volume_multi_attach_pv" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "test_attach_existing_volume_on_instance_launch_pv"
  size_in_gbs         = var.db_size
}

resource "oci_core_instance" "test_vm_multi_attach_instance_launch" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "test_vm_multi_attach_instance"
  shape               = var.vm_multi_attach_instance_shape

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "primary_vnic"
    assign_public_ip = true
    hostname_label   = "tfexampleshinstance"
  }

  source_details {
    source_type = "image"
    source_id = var.instance_image_ocid[var.region]
    # Apply this to set the size of the boot volume that is created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    #boot_volume_size_in_gbs = "60"
    kms_key_id = var.kms_key_ocid
  }

  // Create and attach a volume - iscsi
  launch_volume_attachments {
    type = "iscsi"
    display_name = "test_create_and_attach_volume_on_launch_1"
    launch_create_volume_details {
      volume_creation_type = "ATTRIBUTES"
      compartment_id = var.compartment_ocid
      display_name = "test_create_and_attach_volume_on_launch_1"
      size_in_gbs = var.db_size
    }
  }

  // Create and attach a volume - pv
  launch_volume_attachments {
    type = "paravirtualized"
    display_name = "test_create_and_attach_volume_on_launch_2"
    launch_create_volume_details {
      volume_creation_type = "ATTRIBUTES"
      compartment_id = var.compartment_ocid
      display_name = "test_create_and_attach_volume_on_launch_2"
      size_in_gbs = var.db_size
    }
  }

  // Attach an existing volume - iscsi
  launch_volume_attachments {
    type = "iscsi"
    display_name = "test_attach_existing_volume_on_launch"
    volume_id = oci_core_volume.test_block_volume_multi_attach_iscsi.id
  }

  // Attach an existing volume - pv
  launch_volume_attachments {
    type = "paravirtualized"
    display_name = "test_attach_existing_volume_on_launch"
    volume_id = oci_core_volume.test_block_volume_multi_attach_pv.id
  }

  # Apply the following flag only if you wish to preserve the attached boot volume upon destroying this instance
  # Setting this and destroying the instance will result in a boot volume that should be managed outside of this config.
  # When changing this value, make sure to run 'terraform apply' so that it takes effect before the resource is destroyed.
  #preserve_boot_volume = true

  // Since preserve_data_volumes_created_at_launch is a required parameter for instances launched with volumes,
  // defaulting it to false.
  preserve_data_volumes_created_at_launch = false

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
    user_data           = base64encode(file("./userdata/bootstrap"))
  }
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}" = "awesome-app-server"
  }

  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
  timeouts {
    create = "60m"
  }
}