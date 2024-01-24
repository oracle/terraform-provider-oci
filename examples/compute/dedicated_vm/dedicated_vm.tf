// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_dedicated_vm_host" "test_dedicated_vm_host" {
  #Required
  availability_domain     = data.oci_identity_availability_domain.ad.name
  compartment_id          = var.compartment_ocid
  dedicated_vm_host_shape = "DVH.Standard2.52"

  #Optional
  #  defined_tags = {
  #   "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  #  }
  #freeform_tags = var.dedicated_vm_host_freeform_tags
  display_name = "TestDedicatedVmHost"
}

# instance using dedicated vm host
resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "TestInstanceLabel"
  }

  dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
    # Apply this to set the size of the boot volume that's created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    #boot_volume_size_in_gbs = "60"
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
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

data "oci_core_dedicated_vm_hosts" "test_dedicated_vm_hosts" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
  display_name        = "TestDedicatedVmHost"
  instance_shape_name = "VM.Standard2.1"
  state               = "ACTIVE"
}

data "oci_core_dedicated_vm_host_instance_shapes" "test_dedicated_vm_host_instance_shapes" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain     = data.oci_identity_availability_domain.ad.name
  dedicated_vm_host_shape = "DVH.Standard2.52"
}

data "oci_core_dedicated_vm_host_shapes" "test_dedicated_vm_host_shapes" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_shape_name = "VM.Standard2.1"
}

data "oci_core_dedicated_vm_hosts_instances" "test_dedicated_vm_hosts_instances" {
  #Required
  compartment_id       = var.compartment_ocid
  dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
  depends_on          = [oci_core_instance.test_instance]
}

#output the dedidcated vm host ids
output "dedicated_hos_idst" {
  value = [data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts.id]
}

#output the dedidcated vm host ids
output "dedicated_host_shapes" {
  value = [data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes.dedicated_vm_host_shapes]
}

output "dedicated_vm_host_instances" {
  value = [data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances.dedicated_vm_host_instances]
}

output "dedicated_vm_host_instance_shapes" {
  value = [data.oci_core_dedicated_vm_host_instance_shapes.test_dedicated_vm_host_instance_shapes.dedicated_vm_host_instance_shapes]
}

