// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}
variable "internal_admin_product_description" {
  default = "description"
}

variable "internal_admin_product_limits_public_limit_name" {
  default = "policy-count"
}

variable "internal_admin_product_meters_name" {
  default = "A100_GPU_V2"
}

variable "internal_admin_product_name" {
  default = "opc-limits-testing-admin-example"
}

variable "internal_admin_product_limits_service_name" {
  default = "limits"
}

variable "internal_admin_service_name" {
  default = "COMPUTE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  region           = var.region
}

resource "oci_oci_product_catalog_internal_admin_product" "test_internal_admin_product" {
  #Required
  compartment_id = var.compartment_id
  description    = var.internal_admin_product_description
  name           = var.internal_admin_product_name
  service_name   = var.internal_admin_service_name

  #Optional
  limits {

    #Optional
    public_limit_name   = var.internal_admin_product_limits_public_limit_name
    public_service_name = var.internal_admin_product_limits_service_name
  }
  meters {

    #Optional
    name = var.internal_admin_product_meters_name
  }
}


data "oci_oci_product_catalog_internal_admin_product" "test_internal_admin_product" {
  #Required
  product_id = data.oci_oci_product_catalog_internal_admin_products.test_internal_admin_products.products.0.id
}

data "oci_oci_product_catalog_internal_admin_products" "test_internal_admin_products" {
  #Required
  compartment_id = var.compartment_id
}

output "test_internal_admin_product_id" {
  value = data.oci_oci_product_catalog_internal_admin_product.test_internal_admin_product.id
}