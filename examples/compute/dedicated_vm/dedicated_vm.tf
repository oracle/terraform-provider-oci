// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}

variable "user_ocid" {}

variable "fingerprint" {}

variable "private_key_path" {}

variable "region" {}

variable "compartment_ocid" {}

variable subnet_id {}

variable "instance_image_ocid" {}

locals {
  # For resource "oci_core_dedicated_vm_host" "test_dedicated_vm_host"
  dvh_shape        = "DVH.Standard.E4.128"
  dvh_display_name = "TestDedicatedVmHost"
  capacity_config  = "standard_e4_flex"
  is_memory_encryption_enabled = "true"

  # For resource "oci_core_instance" "test_instance"
  vmi_display_name = "TestInstance"
  vmi_shape                           = "VM.Standard.E4.Flex"
  instance_shape_config_memory_in_gbs = "16"
  instance_shape_config_ocpus = "1"

  # For data "oci_core_dedicated_vm_hosts" "test_dedicated_vm_hosts"
  dvh_lifecycle_state                              = "ACTIVE"
  remaining_memory_in_gbs_greater_than_or_equal_to = "1.0"
  remaining_ocpus_greater_than_or_equal_to = "1.0"

  # For data "oci_core_subnet" "test_subnet"
  subnet_id = var.subnet_id
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = ""
  tenancy_ocid        = var.tenancy_ocid
  user_ocid           = var.user_ocid
  fingerprint         = var.fingerprint
  private_key_path    = var.private_key_path
  region              = var.region
  # version             = "7.22.0"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number = 1
}

resource "oci_core_dedicated_vm_host" "test_dedicated_vm_host" {
  #Required
  availability_domain     = data.oci_identity_availability_domain.ad.name
  compartment_id          = var.compartment_ocid
  dedicated_vm_host_shape = local.dvh_shape

  #Optional
  display_name = local.dvh_display_name
  capacity_config              = local.capacity_config
  is_memory_encryption_enabled = local.is_memory_encryption_enabled

  timeouts {
    create = "60m"
  }
}

# instance using dedicated vm host
resource "oci_core_instance" "test_instance" {
  # Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  shape               = local.vmi_shape

  # Optional but required for DVH Testing
  dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id

  # Optional
  display_name = local.vmi_display_name
  shape_config {
    memory_in_gbs = local.instance_shape_config_memory_in_gbs
    ocpus         = local.instance_shape_config_ocpus
  }
  instance_options {
    are_legacy_imds_endpoints_disabled = true
  }
  availability_config {
    recovery_action = "RESTORE_INSTANCE"
  }

  create_vnic_details {
    subnet_id        = data.oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "TestInstanceLabel"
  }

  source_details {
    source_type = "image"
    source_id = var.instance_image_ocid
    # Apply this to set the size of the boot volume that's created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    # boot_volume_size_in_gbs = "60"
  }

  platform_config {
    type                         = "AMD_VM"
    is_memory_encryption_enabled = true
  }

  timeouts {
    create = "60m"
  }
}

data "oci_core_subnet" "test_subnet" {
  subnet_id = local.subnet_id
}

data "oci_core_dedicated_vm_hosts" "test_dedicated_vm_hosts" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain                              = data.oci_identity_availability_domain.ad.name
  display_name                                     = local.dvh_display_name
  is_memory_encryption_enabled = local.is_memory_encryption_enabled

  instance_shape_name                              = local.vmi_shape
  remaining_memory_in_gbs_greater_than_or_equal_to = local.remaining_memory_in_gbs_greater_than_or_equal_to
  remaining_ocpus_greater_than_or_equal_to         = local.remaining_ocpus_greater_than_or_equal_to
  state                                            = local.dvh_lifecycle_state
}

data "oci_core_dedicated_vm_host_instance_shapes" "test_dedicated_vm_host_instance_shapes" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain     = data.oci_identity_availability_domain.ad.name
  dedicated_vm_host_shape = local.dvh_shape
}

data "oci_core_dedicated_vm_host_shapes" "test_dedicated_vm_host_shapes" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_shape_name = local.vmi_shape
}

data "oci_core_dedicated_vm_host" "test_oci_core_dedicated_vm_host" {
  dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
}

data "oci_core_dedicated_vm_hosts_instances" "test_dedicated_vm_hosts_instances" {
  #Required
  compartment_id = var.compartment_ocid
  dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
}

# output of List DVHs
output "dedicated_host_ids" {
  value = data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts.id
}

# output of List DVH Shapes
output "dedicated_host_shapes" {
  value = data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes.dedicated_vm_host_shapes
}

# output of List DVH Instance Shapes
output "dedicated_vm_host_instance_shapes" {
  value = data.oci_core_dedicated_vm_host_instance_shapes.test_dedicated_vm_host_instance_shapes.dedicated_vm_host_instance_shapes
}

# output of List DVH Instances
output "dedicated_vm_host_instances_data" {
  value = data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances.dedicated_vm_host_instances
}

# output of Get DVH
output "dedicated_vm_host_data" {
  value = data.oci_core_dedicated_vm_host.test_oci_core_dedicated_vm_host.*
}
