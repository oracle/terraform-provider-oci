// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


resource "oci_core_instance_configuration" "test_instance_configuration" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "instance_options"

    options {
      launch_details {
        compartment_id = var.compartment_ocid
        shape = var.instance_shape

        shape_config {
          vcpus = var.instance_vcpus
          memory_in_gbs = var.instance_shape_config_memory_in_gbs
        }

        source_details {
          source_type = "image"

          instance_source_image_filter_details {
            compartment_id   = var.compartment_ocid
            operating_system = "Oracle Linux"
          }
        }

        create_vnic_details {
          subnet_id        = oci_core_subnet.test_subnet.id
          display_name     = "TFExampleInstanceConfigurationVNIC"
          assign_public_ip = true
          skip_source_dest_check = false
        }
      }
    }
  }
}

resource "oci_core_instance" "test_instance_configuration_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstanceConfigurationInstance"
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
  }

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
