variable "generative_ai_private_endpoint_display_name" {
  default = "PrivateEndpointTF"
}
variable "generative_ai_private_endpoint_description" {
  default = "PrivateEndpointTF"
}
variable "generative_ai_private_endpoint_dns_prefix" {
  default = "tersi"
}
variable "generative_ai_private_endpoint_state" {
  default = "ACTIVE"
}
variable "tenancy_ocid" {
}

resource "oci_generative_ai_generative_ai_private_endpoint" "test_generative_ai_private_endpoint" {
  #Required
  compartment_id = var.compartment_ocid
  dns_prefix     = var.generative_ai_private_endpoint_dns_prefix
  subnet_id      = oci_core_subnet.vcn_subnet.id

  #Optional
  description   = var.generative_ai_private_endpoint_description
  display_name  = var.generative_ai_private_endpoint_display_name
  freeform_tags = { "Department" = "Marketing" }
  nsg_ids       = []
}

resource "oci_core_subnet" "vcn_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "PE"
  dns_label           = "PE"
  security_list_ids   = [oci_core_vcn.pe_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.pe_vcn.id
  dhcp_options_id     = oci_core_vcn.pe_vcn.default_dhcp_options_id
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "pe_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "pe"
  dns_label      = "pe"
}

data "oci_generative_ai_generative_ai_private_endpoint" "test_generative_ai_private_endpoint" {
  #Required
  generative_ai_private_endpoint_id = oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id
}

data "oci_generative_ai_generative_ai_private_endpoints" "test_generative_ai_private_endpoints" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  state = var.generative_ai_private_endpoint_state
}