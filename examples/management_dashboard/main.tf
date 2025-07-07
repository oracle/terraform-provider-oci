// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_management_dashboard_management_dashboards_export" "test_export" {
  export_dashboard_id = "%7B%22dashboardIds%22%3A%5B%22dummy_ocid%22%5D%7D"
}

// example showing import_details usage
resource "oci_management_dashboard_management_dashboards_import" "test_import_via_tf_variable" {
  import_details = var.test_import_details
}

// example showing import_details_file usage, sample.json content is same as var.test_import_details
resource "oci_management_dashboard_management_dashboards_import" "test_import_via_file" {
  import_details_file = "sample.json"
}

resource "oci_management_dashboard_management_dashboards_import" "test_import_override_name" {
  override_same_name = true
  import_details_file = "sample.json"
}

resource "oci_management_dashboard_management_dashboards_import" "test_import_override_ss_compartment" {
  override_saved_search_compartment_ocid = "dummy_ocid"
  import_details_file = "sample.json"
}

resource "oci_management_dashboard_management_dashboards_import" "test_import_override_db_compartment" {
  override_dashboard_compartment_ocid = "dummy_ocid"
  import_details_file = "sample.json"
}

output "test_export_data_all" {
  value = data.oci_management_dashboard_management_dashboards_export.test_export
}

output "test_export_data_dashboards" {
  value = data.oci_management_dashboard_management_dashboards_export.test_export.export_details
}

