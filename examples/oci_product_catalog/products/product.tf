// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  region           = var.region
}

data "oci_oci_product_catalog_products" "test_products" {
}

data "oci_oci_product_catalog_product" "test_product" {
  #Required
  product_id = data.oci_oci_product_catalog_products.test_products.products.0.id
}

output "test_product_id" {
  value = data.oci_oci_product_catalog_product.test_product.id
}
