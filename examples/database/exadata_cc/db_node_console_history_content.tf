// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_database_db_node_console_history_content" "test_db_node_console_history_content" {
  #Required
  console_history_id = oci_database_db_node_console_history.test_db_node_console_history.id
  db_node_id = oci_database_db_node.test_db_node.id
}


