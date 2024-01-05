// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

/*
 * This example shows how to manage a retention rule within a bucket
 */

resource "oci_objectstorage_bucket" "retention_rule_bucket" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = "tf-example-bucket-with-retention-rule"
  access_type    = "NoPublicAccess"

  retention_rules {
    display_name = "tf-example-retention-rule"

    duration {
      time_amount = "10"
      time_unit   = "DAYS"
    }

    time_rule_locked = "2120-05-04T17:23:46Z"
  }
}

data "oci_objectstorage_bucket" "retention_rule_bucket" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  name      = oci_objectstorage_bucket.retention_rule_bucket.name
}

output "retention_rules_on_bucket" {
  value = data.oci_objectstorage_bucket.retention_rule_bucket.retention_rules
}

