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

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

# Creates a VCN
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "testVcn"
  dns_label      = "tfVcn"
}

# Creates a subnet
resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "testSubnet"
  dns_label           = "tfsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

# Creates an instance (without assigning a public IP to the primary private IP on the VNIC)
resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "testInstance"
  shape               = var.instance_shape

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  create_vnic_details {
    assign_public_ip = false
    display_name     = "primaryVnic"
    subnet_id        = oci_core_subnet.test_subnet.id
    hostname_label   = "instance"
  }
}

# Creates a secondary VNIC on the instance using a VNIC attachment
resource "oci_core_vnic_attachment" "secondary_vnic_attachment" {
  instance_id  = oci_core_instance.test_instance.id
  display_name = "secondaryVnicAttachment"

  create_vnic_details {
    assign_public_ip       = false
    display_name           = "TFSecondaryVnic"
    skip_source_dest_check = true
    subnet_id              = oci_core_subnet.test_subnet.id
  }
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "instance_vnics" {
  compartment_id      = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  instance_id         = oci_core_instance.test_instance.id
}

# Gets the OCID of the first VNIC
data "oci_core_vnic" "instance_vnic1" {
  vnic_id = data.oci_core_vnic_attachments.instance_vnics.vnic_attachments[0]["vnic_id"]
}

# Gets the OCID of the second VNIC
data "oci_core_vnic" "instance_vnic2" {
  vnic_id = oci_core_vnic_attachment.secondary_vnic_attachment.vnic_id
}

# Gets a list of private IPs on the first VNIC
data "oci_core_private_ips" "private_ips1" {
  vnic_id = data.oci_core_vnic.instance_vnic1.id
}

# Gets a list of private IPs on the second VNIC
data "oci_core_private_ips" "private_ips2" {
  vnic_id = data.oci_core_vnic.instance_vnic2.id
}

# Creates 3 public IPs: 
#  - Assigned ephemeral public IP (assigned to the first private IP on the first VNIC)
#  - Assigned reserved public IP (assigned to the first private IP on the second VNIC)
#  - Unssigned reserved public IP (available for assignment)

resource "oci_core_public_ip" "ephemeral_public_ip_assigned" {
  compartment_id = var.compartment_ocid
  display_name   = "ephemeralPublicIPAssigned"
  lifetime       = "EPHEMERAL"
  private_ip_id  = data.oci_core_private_ips.private_ips1.private_ips[0]["id"]
}

resource "oci_core_public_ip" "reserved_public_ip_assigned" {
  compartment_id = var.compartment_ocid
  display_name   = "reservedPublicIPAssigned"
  lifetime       = "RESERVED"
  private_ip_id  = data.oci_core_private_ips.private_ips2.private_ips[0]["id"]
}

resource "oci_core_public_ip" "reserved_public_ip_unassigned" {
  compartment_id = var.compartment_ocid
  display_name   = "reservedPublicIPUnassigned"
  lifetime       = "RESERVED"
}

# List public IPs: 
#  - Public IP with availability domain scope (ephemeral).
#  - Public IP with regional scope (reserved).

data "oci_core_public_ips" "availability_domain_public_ips_list" {
  compartment_id      = var.compartment_ocid
  scope               = "AVAILABILITY_DOMAIN"
  availability_domain = data.oci_identity_availability_domain.ad.name

  filter {
    name   = "id"
    values = [oci_core_public_ip.ephemeral_public_ip_assigned.id]
  }
}

data "oci_core_public_ips" "region_public_ips_list" {
  compartment_id = var.compartment_ocid
  scope          = "REGION"

  filter {
    name   = "id"
    values = [oci_core_public_ip.reserved_public_ip_assigned.id, oci_core_public_ip.reserved_public_ip_unassigned.id]
  }
}

output "public_ips" {
  value = [
    data.oci_core_public_ips.availability_domain_public_ips_list.public_ips,
    data.oci_core_public_ips.region_public_ips_list.public_ips,
  ]
}

