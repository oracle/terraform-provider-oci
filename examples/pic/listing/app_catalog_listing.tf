// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_core_app_catalog_listings" "test_app_catalog_listings" {
  filter {
    name   = "publisher_name"
    values = ["Oracle CCE Image Management Pipeline"]
  }
}

output "listings" {
  value = [data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings]
}

