// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "bm_shielded_instance_shape" {
  default = "BM.Standard.E4.128"
}

variable "bm_shielded_instance_platform_config_type" {
  default = "AMD_MILAN_BM"
}

resource "oci_core_instance" "test_bm_shielded_instance_with_platform_config" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestShieldedBMInstance"
  shape               = var.bm_shielded_instance_shape


  platform_config {
    type = var.bm_shielded_instance_platform_config_type
    is_measured_boot_enabled = false
    is_secure_boot_enabled = true
    is_trusted_platform_module_enabled = true
  }

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleshbminstance"
  }

  source_details {
    source_type = "image"
    source_id   = data.oci_core_images.supported_bm_shielded_instances_shape_images.images[0]["id"]
  }

  # Apply the following flag only if you wish to preserve the attached boot volume upon destroying this instance
  # Setting this and destroying the instance will result in a boot volume that should be managed outside of this config.
  # When changing this value, make sure to run 'terraform apply' so that it takes effect before the resource is destroyed.
  #preserve_boot_volume = true

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

# Gets a list of all images that support a given VM Instance shape
data "oci_core_images" "supported_bm_shielded_instances_shape_images" {
  compartment_id   = var.tenancy_ocid
  shape            = var.bm_shielded_instance_shape
  operating_system = "Oracle Linux"

}

