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

variable "oda_instance_state" {
  default = "INACTIVE"
}

variable "compartment_ocid" {
}

variable "oda_private_endpoint_display_name" {
  default = "oda_pe_name"
}

variable "oda_private_endpoint_description" {
  default = "oda_pe_description"
}

variable "oda_private_endpoint_freeform_tags" {
  default = { "test-key" = "value" }
}

variable "oda_private_endpoint_nsg_ids" {
  default = []
}

variable "oda_private_endpoint_defined_tags_value" {
  default = "value"
}

variable "oda_private_endpoint_is_used_for_rac_dbs" {
  default = true
}

variable "oda_private_endpoint_scan_proxy_protocol" {
  default = "TCP"
}

variable "oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_fqdn" {
  default = "myprefix-test-scan"
}

variable "oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_port" {
  default = 1521
}

variable "oda_private_endpoint_scan_proxy_scan_listener_type" {
  default = "FQDN"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_oda_oda_instance" "TFOdaInstance" {
  compartment_id = var.compartment_ocid
  shape_name     = "DEVELOPMENT"
  description    = "test instance"
  display_name   = "TestInstance"

  #Optional
  state = var.oda_instance_state
}

data "oci_oda_oda_instances" "TFOdaInstances" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_oda_oda_instance.TFOdaInstance.display_name
}

data "oci_oda_oda_instance" "TFOdaInstance" {
  #Required
  oda_instance_id = oci_oda_oda_instance.TFOdaInstance.id
}

resource "oci_oda_oda_private_endpoint" "TFOdaPrivateEndpoint" {
  compartment_id = var.compartment_ocid
  subnet_id      = var.oda_private_endpoint_description
  vcn_id         = oci_core_vcn.pe_vcn.id
  display_name   = var.oda_private_endpoint_display_name
  description    = var.oda_private_endpoint_description
  freeform_tags  = var.oda_private_endpoint_freeform_tags
  nsg_ids        = var.oda_private_endpoint_nsg_ids
  defined_tags   = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.oda_private_endpoint_defined_tags_value}")}"
  is_used_for_rac_dbs = var.oda_private_endpoint_is_used_for_rac_dbs
}

data "oci_oda_oda_private_endpoint" "test_oda_private_endpoint" {
  #Required
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
}

data "oci_oda_oda_private_endpoints" "test_oda_private_endpoints" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.oda_private_endpoint_display_name
}

resource "oci_oda_oda_private_endpoint_attachment" "TFOdaPrivateEndpointAttachment" {
  oda_instance_id = oci_oda_oda_instance.TFOdaInstance.id
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
}

data "oci_oda_oda_private_endpoint_attachment" "test_oda_private_endpoint_attachment" {
  #Required
  oda_private_endpoint_attachment_id = oci_oda_oda_private_endpoint_attachment.TFOdaPrivateEndpointAttachment.id
}

data "oci_oda_oda_private_endpoint_attachments" "test_oda_private_endpoint_attachments" {
  #Required
  compartment_id = var.compartment_ocid
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
}

resource "oci_oda_oda_private_endpoint_scan_proxy" "TFOdaPrivateEndpointScanProxy" {
  #Required
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
  protocol = var.oda_private_endpoint_scan_proxy_protocol
  scan_listener_infos {
    #Optional
    scan_listener_fqdn = var.oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_fqdn
    scan_listener_port = var.oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_port
  }
  scan_listener_type = var.oda_private_endpoint_scan_proxy_scan_listener_type
}

data "oci_oda_oda_private_endpoint_scan_proxy" "test_oda_private_endpoint_scan_proxy" {
  #Required
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
  oda_private_endpoint_scan_proxy_id = oci_oda_oda_private_endpoint_scan_proxy.TFOdaPrivateEndpointScanProxy.id
}

data "oci_oda_oda_private_endpoint_scan_proxies" "test_oda_private_endpoint_scan_proxies" {
  #Required
  oda_private_endpoint_id = oci_oda_oda_private_endpoint.TFOdaPrivateEndpoint.id
}