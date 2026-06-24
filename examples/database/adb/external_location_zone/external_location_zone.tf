// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

resource "random_string" "autonomous_database_name_suffix" {
  length  = 10
  upper   = false
  special = false
}

resource "oci_database_autonomous_database" "external_location_zone" {
  admin_password              = random_string.autonomous_database_admin_password.result
  compartment_id              = var.compartment_ocid
  compute_count               = "8.0"
  compute_model               = "ECPU"
  data_storage_size_in_tbs    = "1"
  db_name                     = "adb${random_string.autonomous_database_name_suffix.result}"
  display_name                = "adb${random_string.autonomous_database_name_suffix.result}"
  is_mtls_connection_required = true

  lifecycle {
    postcondition {
      condition     = self.external_location_zone != null && self.external_location_zone != ""
      error_message = "The created autonomous database did not return external_location_zone."
    }
  }
}

data "oci_database_autonomous_database" "external_location_zone" {
  autonomous_database_id = oci_database_autonomous_database.external_location_zone.id

  lifecycle {
    postcondition {
      condition     = self.external_location_zone == oci_database_autonomous_database.external_location_zone.external_location_zone
      error_message = "The autonomous database data source returned a different external_location_zone than the resource."
    }
  }
}

output "resource_external_location_zone" {
  value = oci_database_autonomous_database.external_location_zone.external_location_zone
}

output "datasource_external_location_zone" {
  value = data.oci_database_autonomous_database.external_location_zone.external_location_zone
}

output "external_location_zone_matches_datasource" {
  value = oci_database_autonomous_database.external_location_zone.external_location_zone == data.oci_database_autonomous_database.external_location_zone.external_location_zone
}
