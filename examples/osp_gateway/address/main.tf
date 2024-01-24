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

variable "address_id" {
  default = "MX|LP|A|608700|5065_SPA"
}

variable "country_code" {
  default = "hu"
}

variable "address" {
  default = {
    "line1" = "Blvd Puerta de Hierro 5065"
    "line2" = "Col Puerta de Hierro"
    "line3" = "45116  ZAPOPAN, JAL"
    "country" = "MX"
    "county" = "Zapopan"
    "state" = "JAL"
    "postal_code" = "45116"
    "city" = "Zapopan"
  }
  type = object({
    line1 = string
    line2 = string
    line3 = string
    country = string
    county = string
    state = string
    postal_code = string
    city = string
  })
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_osp_gateway_address" "test_address" {
  address_id      = var.address_id
  compartment_id  = var.tenancy_ocid
  osp_home_region = var.region
}

data "oci_osp_gateway_address_rule" "test_address_rule" {
  compartment_id = var.tenancy_ocid
  osp_home_region = var.region
  country_code = var.country_code
}

resource "oci_osp_gateway_address_action_verification" "test_address_action_verification" {
  compartment_id = var.tenancy_ocid
  osp_home_region = var.region
  city = var.address.city
  country = var.address.country
  county = var.address.county
  line1 = var.address.line1
  line2 = var.address.line2
  line3 = var.address.line3
  postal_code = var.address.postal_code
  state = var.address.state
}