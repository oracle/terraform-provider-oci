// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

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

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaarjbcsqt4pg2hmuspw7rhpvjvua32yfjiajakehcd2nxskdnxrcia"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaa6bftra47564ph2uowoooiexeyfmyxokcu7bxaenldni3t7frm3ia"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
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
  compartment_id      = var.compartment_ocid
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
    source_id   = var.instance_image_ocid[var.region]
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
}

resource "null_resource" "remote-exec" {
  depends_on = [oci_core_instance.test_instance]

  provisioner "remote-exec" {
    connection {
      agent       = false
      timeout     = "60m"
      host        = oci_core_instance.test_instance.public_ip
      user        = "opc"
      private_key = var.ssh_private_key
    }

    # https://blogs.oracle.com/linux/oracle-instant-client-rpms-now-available-on-oracle-linux-yum-servers-in-oci
    inline = [
      "sudo -E wget http://yum-${var.region}.oracle.com/yum-${var.region}-ol7.repo",
      "sudo yum-config-manager --enable ol7_oci_included",
      "sudo yum -y --enablerepo ol7_oci_included install osms-agent",
    ]
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_default_route_table" "default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id
  display_name               = "DefaultRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

resource "oci_osmanagement_managed_instance_group" "test_managed_instance_group" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "TF-managed-instance-group"

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF Managed instance group"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
}

resource "oci_osmanagement_managed_instance_group" "test_managed_instance_group2" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "TF-managed-instance-group2"

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "awesome-app-server"
  }
  description = "TF Managed instance group2"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
}

resource "oci_osmanagement_software_source" "test_software_source" {
  #Required
  arch_type      = "X86_64"
  compartment_id = var.compartment_ocid
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
}

resource "oci_osmanagement_software_source" "test_software_source2" {
  #Required
  arch_type      = "X86_64"
  compartment_id = var.compartment_ocid
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
}

# attach/detach parent and child software sources, managed instance groups from/to managed instance
# NOTE: It can take some time for the compute instance to become managed instance after installing the OSMS agent
# This resource on CREATE will detach any already attached parent software source, child software sources, managed instance
# groups to the managed instance.
resource "oci_osmanagement_managed_instance_management" "test_managed_instance_management" {
  depends_on          = [null_resource.remote-exec]
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
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_osmanagement_managed_instance_group.test_managed_instance_group.display_name
  state        = "ACTIVE"
}

data "oci_osmanagement_software_sources" "test_software_sources" {
  #Required
  compartment_id = var.compartment_ocid

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

