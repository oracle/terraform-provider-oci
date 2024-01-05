// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_pluggable_database" "test_pluggable_databases_local_clone" {
    pdb_name = "localClonePdb"
    container_database_id = "${data.oci_database_database.t.id}"
    lifecycle {
        ignore_changes = ["defined_tags"]
    }
    pdb_admin_password = "BEstrO0ng_#11"
    tde_wallet_password = "BEstrO0ng_#11"
    pdb_creation_type_details {
        creation_type = "LOCAL_CLONE_PDB"
        source_pluggable_database_id = "${data.oci_database_pluggable_database.test_pluggable_database.id}"
    }
}