// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_id" {}

variable "db_node_console_history_defined_tags_value" {
  default = "definedTags"
}

variable "db_node_console_history_display_name" {
  default = "console-history-20221202-1943"
}

variable "db_node_console_history_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "db_node_console_history_state" {
  default = "SUCCEEDED"
}

resource "oci_database_db_node_console_history" "test_db_node_console_history" {
  #Required
  db_node_id   = oci_database_db_node.test_db_node.id
  display_name = var.db_node_console_history_display_name

  #Optional
  defined_tags  = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.db_node_console_history_defined_tags_value
  }
  freeform_tags = var.db_node_console_history_freeform_tags
}

data "oci_database_db_node_console_histories" "test_db_node_console_histories" {
  #Required
  db_node_id = data.oci_database_db_nodes.test_db_nodes.db_nodes[0]["id"]

  #Optional
  display_name = var.db_node_console_history_display_name
  state        = var.db_node_console_history_state
}