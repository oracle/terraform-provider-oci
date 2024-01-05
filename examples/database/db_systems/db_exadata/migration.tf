// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// This is 1 time action to migrate test_db_system into db exaCs
// and the test_db_system will become `Migrated`
resource "oci_database_migration" "test_migration" {
  #Required
  db_system_id = oci_database_db_system.test_db_system.id
}
