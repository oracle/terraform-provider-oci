// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to import custom content
 */

# Create a custom content
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentNew" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  import_custom_content_file = "./files/TFSource1.zip"

}

# Create a custom content with overwrite false
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentOverwriteFalse" {
  depends_on = [oci_log_analytics_log_analytics_import_custom_content.importCustomContentNew]
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  import_custom_content_file = "./files/TFSource1.zip"
  is_overwrite               = "false"
}

# Create a custom content with overwrite true
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentOverwriteTrue" {
  depends_on = [oci_log_analytics_log_analytics_import_custom_content.importCustomContentOverwriteFalse]
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  import_custom_content_file = "./files/TFSource1.zip"
  is_overwrite               = "true"
}