// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

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

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "testinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_image" "custom_image" {
  compartment_id = var.compartment_ocid
  instance_id    = oci_core_instance.test_instance.id

  launch_mode = "NATIVE"

  timeouts {
    create = "30m"
  }
}

resource "oci_core_shape_management" "compatible_shape" {
  compartment_id = var.compartment_ocid
  image_id       = oci_core_image.custom_image.id
  shape_name     = "VM.Standard2.1"
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

# Gets a list of Availability Domains
data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

# Gets the custom image that will be created by this Terraform config
data "oci_core_images" "custom_images" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_image.custom_image.id]
  }
}

# Gets a list of images within a tenancy
data "oci_core_images" "supported_shape_images" {
  compartment_id = var.tenancy_ocid
  # Uncomment below to filter images that support a specific instance shape
  #shape                    = "VM.Standard2.1"

  # Uncomment below to filter images that are a specific OS
  #operating_system         = "Oracle Linux"

  # Uncomment below to filter images that are a specific OS version
  #operating_system_version = "7.5"

  # Uncomment below to sort images by creation time
  #sort_by                 = "TIMECREATED"
  # Default sort order for TIMECREATED is descending (DESC)
  #sort_order              = "ASC"

  # Uncomment below to sort images by display name, display name sort order is case-sensitive
  #sort_by                 = "DISPLAYNAME"
  # Default sort order for DISPLAYNAME is ascending (ASC)
  #sort_order              = "DESC"
}

# Another way to get the custom image that will be created by this Terraform config
data "oci_core_image" "supported_image" {
  image_id = oci_core_image.custom_image.id
}

output "supported_shape_images" {
  value = data.oci_core_images.supported_shape_images.images
}

resource "oci_core_instance" "test_instance_from_image" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstanceImage"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "testimage"
  }

  source_details {
    source_type = "image"
    source_id   = oci_core_image.custom_image.id
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}
