// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_pluggable_databases_local_clone" "test_pluggable_databases_local_clone" {
    cloned_pdb_name = "NewSalesPdb"
    lifecycle {
        ignore_changes = ["defined_tags"]
    }
    pdb_admin_password = "BEstrO0ng_#11"
    pluggable_database_id = "${oci_database_pluggable_database.test_pluggable_database.id}"
    target_tde_wallet_password = "BEstrO0ng_#11"
}