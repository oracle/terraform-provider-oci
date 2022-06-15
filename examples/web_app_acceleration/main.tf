variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "web_app_acceleration_policy_display_name" {
  default = "WAAPolicy1"
}

variable "web_app_acceleration_policy_id" {
  default = "id"
}

variable "web_app_acceleration_policy_state" {
  default = []
}

variable "web_app_acceleration_display_name" {
  default = "WAA1"
}

variable "web_app_acceleration_id" {
  default = "id"
}

variable "web_app_acceleration_state" {
  default = []
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_waa_web_app_acceleration_policy" "test_web_app_acceleration_policy" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name  = var.web_app_acceleration_policy_display_name

  response_caching_policy {

    #Optional
    is_response_header_based_caching_enabled = true
  }
  response_compression_policy {

    #Optional
    gzip_compression {

      #Optional
      is_enabled = true
    }
  }

  #Optional
  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  freeform_tags = {
    "bar-key" = "value"
  }
}

data "oci_waa_web_app_acceleration_policies" "test_web_app_acceleration_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.web_app_acceleration_policy_display_name
  id           = var.web_app_acceleration_policy_id
  state        = var.web_app_acceleration_policy_state
}

resource "oci_waa_web_app_acceleration" "test_web_app_acceleration" {
  #Optional
  display_name  = var.web_app_acceleration_display_name

  #Required
  backend_type                   = "LOAD_BALANCER"
  compartment_id                 = var.compartment_ocid
  load_balancer_id               = oci_load_balancer.test_load_balancer.id
  web_app_acceleration_policy_id = oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id

  #Optional
  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  freeform_tags = {
    "bar-key" = "value"
  }
}

data "oci_waa_web_app_accelerations" "test_web_app_accelerations" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name                   = var.web_app_acceleration_display_name
  id                             = var.web_app_acceleration_id
  state                          = var.web_app_acceleration_state
  web_app_acceleration_policy_id = oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id
}

/* VCN for Network Address List and Load Balancers */
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "vcn1"
  dns_label      = "vcn1"
}

/* Data needed to create Load Balancer, which is needed for Web App Acceleration */
resource "oci_core_subnet" "test_subnet" {
  cidr_block                 = "10.1.20.0/24"
  display_name               = "subnet1"
  dns_label                  = "subnet1"
  compartment_id             = var.compartment_ocid
  vcn_id                     = oci_core_vcn.test_vcn.id
  security_list_ids          = [oci_core_vcn.test_vcn.default_security_list_id]
  route_table_id             = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id            = oci_core_vcn.test_vcn.default_dhcp_options_id
  prohibit_public_ip_on_vnic = true

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_load_balancer" "test_load_balancer" {
  shape          = "100Mbps"
  compartment_id = var.compartment_ocid

  subnet_ids = [
    oci_core_subnet.test_subnet.id,
  ]

  display_name               = "lb1"
  is_private                 = true
  network_security_group_ids = [oci_core_network_security_group.test_network_security_group.id]
}