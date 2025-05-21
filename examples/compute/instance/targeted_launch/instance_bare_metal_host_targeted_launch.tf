variable "tenancy_ocid" {}
variable "customer_bare_metal_host_id" {}
variable "compartment_ocid" {}
variable "shape" {
  default = "BM.Standard3.64"
}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "instance_image_ocid" {}
variable "ad" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = var.ad
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
}

resource "oci_core_instance" "test_customer_bare_metal_host_targeted_launch" {
  availability_domain = var.ad
  compartment_id      = var.compartment_ocid
  display_name        = "TestCustomerBareMetalHostInstance"
  shape               = var.shape

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "tlcustomerbaremetalhost"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid
  }

  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  timeouts {
    create = "60m"
  }

  placement_constraint_details {
    type = "COMPUTE_BARE_METAL_HOST"
    compute_bare_metal_host_id = var.customer_bare_metal_host_id
  }

}