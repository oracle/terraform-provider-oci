// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "subnet_ocid" {
}

variable "image_ocid" {
}

variable "config_file_profile" {
}

# provider "oci" {
#   region              = var.region
#   auth                = "SecurityToken"
#   config_file_profile = var.config_file_profile
#   version             = "7.18.0"
# }

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_instance_configuration" "test_instance_configuration_with_aie" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      compartment_id = var.compartment_ocid
      shape = "BM.GPU.A10.4"
      is_ai_enterprise_enabled  = true

      source_details {
        source_type = "image"
        image_id = var.image_ocid
      }

      instance_options {
        are_legacy_imds_endpoints_disabled = true
      }

      create_vnic_details {
        subnet_id        = var.subnet_ocid
        display_name     = "Primaryvnic"
      }
    }
  }
}

resource "oci_core_instance_pool" "test_instance_pool_with_aie" {
  compartment_id            = var.compartment_ocid
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration_with_aie.id
  size                      = 1
  state                     = "RUNNING"
  display_name              = "TestInstancePool"

  placement_configurations {
    availability_domain = data.oci_identity_availability_domain.ad.name
    primary_subnet_id   = var.subnet_ocid
  }
}

resource "oci_core_instance" "test_instance_with_aie" {
  availability_domain       = data.oci_identity_availability_domain.ad.name
  compartment_id            = var.compartment_ocid
  display_name              = "TestInstance"
  shape                     = "BM.GPU.A10.4"
  is_ai_enterprise_enabled  = true

  create_vnic_details {
    subnet_id        = var.subnet_ocid
    display_name     = "Primaryvnic"
  }

  source_details {
    source_type = "image"
    source_id = var.image_ocid
  }

  instance_options {
    are_legacy_imds_endpoints_disabled = true
  }

  timeouts {
    create = "60m"
  }
}

output "aie_instance_config_data" {
  value = oci_core_instance_configuration.test_instance_configuration_with_aie
}

data "oci_core_instance_pool_instances" "aie_instance_pool_instance" {
  compartment_id   = var.compartment_ocid
  instance_pool_id = oci_core_instance_pool.test_instance_pool_with_aie.id
}

data "oci_core_instance" "aie_instance_pool_instance_data" {
  instance_id = data.oci_core_instance_pool_instances.aie_instance_pool_instance.instances[0].id
}

data "oci_core_instances" "aie_instance_data" {
  compartment_id = var.compartment_ocid
  filter {
    name = "id"
    values = [oci_core_instance.test_instance_with_aie.id]
  }
}

output "aie_instance_pool_instance_data" {
  value = data.oci_core_instance.aie_instance_pool_instance_data
}

output "aie_instance_data" {
  value = data.oci_core_instances.aie_instance_data
}