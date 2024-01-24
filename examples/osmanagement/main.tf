// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "ssh_public_key" {}
variable "instance_image_ocid" {}

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "tag_namespace_name" {
  default = "testexamples-tag-namespace"
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

resource "oci_identity_tag" "tag2" {
  #Required
  description      = "tf example tag 2"
  name             = "tf-example-tag-2"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_id
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid
    # Apply this to set the size of the boot volume that's created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    #boot_volume_size_in_gbs = "60"
  }

  # Apply the following flag only if you wish to preserve the attached boot volume upon destroying this instance
  # Setting this and destroying the instance will result in a boot volume that should be managed outside of this config.
  # When changing this value, make sure to run 'terraform apply' so that it takes effect before the resource is destroyed.
  #preserve_boot_volume = true

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
    user_data           = base64encode(file("./userdata/bootstrap"))
  }
  timeouts {
    create = "60m"
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "TestVcn"
  dns_label      = "testvcn"

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_id
  display_name   = "TestInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_core_default_route_table" "default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id
  display_name               = "DefaultRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "time_sleep" "wait_3_minutes" {
  depends_on = [oci_core_instance.test_instance]

  create_duration = "3m"
}

resource "oci_osmanagement_managed_instance_group" "test_managed_instance_group" {
  depends_on = [time_sleep.wait_3_minutes]

  #Required
  compartment_id = var.compartment_id
  display_name   = "TF-managed-instance-group"

  #Optional
  managed_instance_ids = [oci_core_instance.test_instance.id]
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF Managed instance group"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_osmanagement_managed_instance_group" "test_managed_instance_group2" {
  depends_on = [time_sleep.wait_3_minutes]

  #Required
  compartment_id = var.compartment_id
  display_name   = "TF-managed-instance-group2"

  #Optional
  managed_instance_ids = [oci_core_instance.test_instance.id]
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF Managed instance group2"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_osmanagement_software_source" "test_software_source" {
  #Required
  arch_type      = "X86_64"
  compartment_id = var.compartment_id
  display_name   = "TF-software-source"

  #Optional
  checksum_type = "SHA1"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF software source"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

resource "oci_osmanagement_software_source" "test_software_source2" {
  #Required
  arch_type      = "X86_64"
  compartment_id = var.compartment_id
  display_name   = "TF-software-source2"

  #Optional
  checksum_type = "SHA1"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF software source2"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  lifecycle {
    ignore_changes = ["defined_tags"]
  }
}

# attach/detach parent and child software sources, managed instance groups from/to managed instance
# NOTE: It can take some time for the compute instance to become managed instance after installing the OSMS agent
# This resource on CREATE will detach any already attached parent software source, child software sources, managed instance
# groups to the managed instance.
resource "oci_osmanagement_managed_instance_management" "test_managed_instance_management" {
  managed_instance_id = oci_core_instance.test_instance.id

  parent_software_source {
    id   = oci_osmanagement_software_source.test_software_source.id
    name = oci_osmanagement_software_source.test_software_source.display_name
  }

  managed_instance_groups {
    id           = oci_osmanagement_managed_instance_group.test_managed_instance_group.id
    display_name = oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name
  }

  managed_instance_groups {
    id           = oci_osmanagement_managed_instance_group.test_managed_instance_group2.id
    display_name = oci_osmanagement_managed_instance_group.test_managed_instance_group2.display_name
  }
}

data "oci_osmanagement_managed_instance_groups" "test_managed_instance_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name
  state        = "ACTIVE"
}

data "oci_osmanagement_software_sources" "test_software_sources" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = oci_osmanagement_software_source.test_software_source.display_name
  state        = "ACTIVE"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_osmanagement_managed_instance" "test_managed_instance" {
  #Required
  managed_instance_id = oci_osmanagement_managed_instance_management.test_managed_instance_management.id
}

data "oci_osmanagement_managed_instance_event_report" "test_managed_instance_event_report" {
  #Required
  managed_instance_id = oci_osmanagement_managed_instance_management.test_managed_instance_management.id
  compartment_id = var.tenancy_ocid
}

output "managed_instance_output" {
  value = [data.oci_osmanagement_managed_instance.test_managed_instance.managed_instance_groups]
}

output "managed_instance_groups_output" {
  value = [data.oci_osmanagement_managed_instance_groups.test_managed_instance_groups.display_name]
}

