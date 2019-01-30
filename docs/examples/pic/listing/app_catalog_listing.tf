// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_core_app_catalog_listings" "test_app_catalog_listings" {
  filter {
    name   = "publisher_name"
    values = ["Oracle CCE Image Management Pipeline"]
  }
}

output "listings" {
  value = ["${data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings}"]
}
