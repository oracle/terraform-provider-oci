// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "db_node_console_connection_public_key" {
  default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
}


resource "oci_database_db_node_console_connection" "test_db_node_console_connection" {
  #Required
  db_node_id = data.oci_database_db_nodes.test_db_nodes.db_nodes[0]["id"]
  public_key = var.db_node_console_connection_public_key

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

data "oci_database_db_node_console_connections" "test_db_node_console_connections" {
  #Required
  db_node_id = oci_database_db_node.test_db_node.id
}
