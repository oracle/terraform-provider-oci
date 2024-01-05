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

variable "ssh_public_key" {
}

variable "ssh_private_key" {
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

variable "instance_shape" {
  default = "VM.Standard2.1"
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

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "Test_IG"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_default_route_table" "test_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

# Gets a list of vNIC attachments on the instance
data "oci_core_vnic_attachments" "instance_vnics" {
  compartment_id      = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_id         = oci_core_instance.test_instance.id
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "instance_vnic" {
  vnic_id = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id      = oci_core_subnet.test_subnet.id
    hostname_label = "testinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  extended_metadata = {
    ssh_authorized_keys = var.ssh_public_key
    some_string         = "stringA"
    nested_object       = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
    array               = "[\"stringA\", \"stringB\"]"
  }
}

resource "null_resource" "remote-exec" {
  depends_on = [oci_core_instance.test_instance]

  provisioner "remote-exec" {
    connection {
      agent       = false
      timeout     = "30m"
      host        = data.oci_core_vnic.instance_vnic.public_ip_address
      user        = "opc"
      private_key = var.ssh_private_key
    }

    inline = [
      "export STR=$(curl --header \"Authorization: Bearer Oracle\" http://169.254.169.254/opc/v2/instance/metadata/nested_object/object/some_string)",
      "if [[ \"$STR\" != \"stringC\" ]]; then exit 1; fi",
    ]
  }
}

output "instance_public_ip" {
  value = [data.oci_core_vnic.instance_vnic.public_ip_address]
}

