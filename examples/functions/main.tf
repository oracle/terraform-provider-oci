// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tf-vcn"
  dns_label      = "dnslabel"
}

resource "oci_core_internet_gateway" "test_network_entity" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
  display_name   = "-tf-internet-gateway"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = "${var.compartment_ocid}"

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
  }

  vcn_id = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain        = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
  cidr_block                 = "10.0.0.0/16"
  compartment_id             = "${var.compartment_ocid}"
  dhcp_options_id            = "${oci_core_vcn.test_vcn.default_dhcp_options_id}"
  display_name               = "tf-subnet"
  dns_label                  = "dnslabel"
  prohibit_public_ip_on_vnic = "false"
  route_table_id             = "${oci_core_route_table.test_route_table.id}"
  security_list_ids          = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
  vcn_id                     = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_functions_application" "test_application" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "example-application"
  subnet_ids     = ["${oci_core_subnet.test_subnet.id}"]

  #Optional
  config = "${var.config}"
}

data "oci_functions_applications" "test_applications" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "example-application"
  id           = "${oci_functions_application.test_application.id}"
  state        = "${var.application_state}"
}

resource "oci_functions_function" "test_function" {
  #Required
  application_id = "${oci_functions_application.test_application.id}"
  display_name   = "example-function"
  image          = "${var.function_image}"
  memory_in_mbs  = "128"

  #Optional
  config             = "${var.config}"
  image_digest       = "${var.function_image_digest}"
  timeout_in_seconds = "30"
}

data "oci_functions_functions" "test_functions" {
  #Required
  application_id = "${oci_functions_application.test_application.id}"

  #Optional
  display_name = "example-function"
  id           = "${oci_functions_function.test_function.id}"
  state        = "AVAILABLE"
}

resource "oci_functions_invoke_function" "test_invoke_function" {
  fn_intent            = "httprequest"
  fn_invoke_type       = "sync"
  function_id          = "${oci_functions_function.test_function.id}"
  invoke_function_body = "${var.invoke_function_body}"
}

resource "oci_functions_invoke_function" "test_invoke_function_source_path" {
  fn_intent              = "httprequest"
  fn_invoke_type         = "sync"
  function_id            = "${oci_functions_function.test_function.id}"
  input_body_source_path = "${var.invoke_function_body_source_path}"
}

resource "oci_functions_invoke_function" "test_invoke_function_detached" {
  fn_intent            = "httprequest"
  fn_invoke_type       = "detached"
  function_id          = "${oci_functions_function.test_function.id}"
  invoke_function_body = "${var.invoke_function_body}"
}

resource "oci_functions_invoke_function" "test_invoke_function_encoded_body" {
  fn_intent                           = "cloudevent"
  fn_invoke_type                      = "sync"
  function_id                         = "${oci_functions_function.test_function.id}"
  invoke_function_body_base64_encoded = "${base64encode(var.invoke_function_body)}"
}

resource "oci_functions_invoke_function" "test_invoke_function_encoded_body_detached" {
  fn_intent                           = "httprequest"
  fn_invoke_type                      = "detached"
  function_id                         = "${oci_functions_function.test_function.id}"
  invoke_function_body_base64_encoded = "${base64encode(var.invoke_function_body)}"
}

resource "oci_functions_invoke_function" "test_invoke_function_encoded_content" {
  fn_intent             = "httprequest"
  fn_invoke_type        = "sync"
  function_id           = "${oci_functions_function.test_function.id}"
  base64_encode_content = true
}

output "test_invoke_function_content" {
  value = "${oci_functions_invoke_function.test_invoke_function.content}"
}

output "test_invoke_function_source_path_content" {
  value = "${oci_functions_invoke_function.test_invoke_function_source_path.content}"
}

output "test_invoke_function_encoded_body" {
  value = "${oci_functions_invoke_function.test_invoke_function_encoded_body.content}"
}

output "test_invoke_function_encoded_content" {
  value = "${base64decode(oci_functions_invoke_function.test_invoke_function_encoded_content.content)}"
}
