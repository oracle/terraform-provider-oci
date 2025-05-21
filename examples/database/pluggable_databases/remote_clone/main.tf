# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - main file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/remote_clone
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#      Clone a Pluggable Database online remotely between DB Systems in Oracle Base Database Service using database link
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


# DB SYSTEM
resource "oci_database_db_system" "test_db_system" {
    compartment_id = var.compartment_id
    subnet_id = oci_core_subnet.test_subnet.id
    database_edition = "ENTERPRISE_EDITION"
    availability_domain = data.oci_identity_availability_domains.test_availability_domain.availability_domains.0.name
    disk_redundancy = "NORMAL"
    shape = "VM.Standard2.1"
    cpu_core_count = var.cpu_core_count
    ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
    display_name = "tfDBSystem-PDB"
    domain = "${oci_core_subnet.test_subnet.dns_label}.${oci_core_virtual_network.test_vcn.dns_label}.oraclevcn.com"
    hostname = "OracleDB" // this will be lowercased server side
    data_storage_size_in_gb = "256"
    license_model = "LICENSE_INCLUDED"
    node_count = "1"
    fault_domains = ["FAULT-DOMAIN-1"]
    db_home {
        db_version = "23.4.1.24.06"
        display_name = "tfDBHome"
        database {
            admin_password = "BEstrO0ng_#11"
            db_name = "tfDB"
            character_set = "AL32UTF8"
            defined_tags = tomap({"example-tag-namespace-all.example-tag" = "originalValue"})
            freeform_tags = {"Department" = "Finance"}
            ncharacter_set = "AL16UTF16"
            db_workload = "OLTP"
            pdb_name = "tfPDB"
        }
    }
    db_system_options {
        storage_management = "LVM"
    }
    defined_tags = tomap({"example-tag-namespace-all.example-tag" = "originalValue"})
    freeform_tags = {"Department" = "Finance"}
    # nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
    lifecycle {
        ignore_changes = [db_home.0.db_version, defined_tags, db_home.0.database.0.defined_tags]
    }

    timeouts {
        create = "60m"
        delete = "2h"
    }
}

# REMOTE DB SYSTEM
resource "oci_database_db_system" "test_db_system_remote" {
    compartment_id = var.compartment_id
    subnet_id = oci_core_subnet.test_subnet.id
    database_edition = "ENTERPRISE_EDITION"
    availability_domain = data.oci_identity_availability_domains.test_availability_domain.availability_domains.0.name
    disk_redundancy = "NORMAL"
    shape = "VM.Standard2.1"
    cpu_core_count = var.cpu_core_count
    ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
    display_name = "tfDBSystem-RemotePDB"
    domain = "${oci_core_subnet.test_subnet.dns_label}.${oci_core_virtual_network.test_vcn.dns_label}.oraclevcn.com"
    hostname = "OracleDBRemote" // this will be lowercased server side
    data_storage_size_in_gb = "256"
    license_model = "LICENSE_INCLUDED"
    node_count = "1"
    fault_domains = ["FAULT-DOMAIN-1"]
    db_home {
        db_version = "23.4.1.24.06"
        display_name = "tfDBHome"
        database {
            admin_password = "BEstrO0ng_#11"
            db_name = "tfDB"
            character_set = "AL32UTF8"
            defined_tags = tomap({"example-tag-namespace-all.example-tag" = "originalValue"})
            freeform_tags = {"Department" = "Finance"}
            ncharacter_set = "AL16UTF16"
            db_workload = "OLTP"
            pdb_name = "tfPDB"
        }
    }
    db_system_options {
        storage_management = "LVM"
    }
    defined_tags = tomap({"example-tag-namespace-all.example-tag" = "originalValue"})
    freeform_tags = {"Department" = "Finance"}
    # nsg_ids = [oci_core_network_security_group.test_network_security_group_remote.id]
    lifecycle {
        ignore_changes = [db_home.0.db_version, defined_tags, db_home.0.database.0.defined_tags]
    }

    timeouts {
        create = "60m"
        delete = "2h"
    }
}

resource "oci_database_pluggable_database" "test_pluggable_database" {
    container_database_id = data.oci_database_database.test_database.id
    pdb_admin_password = "BEstrO0ng_#11"
    pdb_name = "tfSalesPDB"
    tde_wallet_password = "BEstrO0ng_#11"

    lifecycle {
        ignore_changes = [defined_tags]
    }

    timeouts {
        create = "60m"
        delete = "2h"
    }
}

resource "oci_database_pluggable_database" "test_pluggable_databases_remote_clone" {
    pdb_name = "tfPDBRemoteClone"
    pdb_admin_password = "BEstrO0ng_#11"
    container_database_id = data.oci_database_database.test_database_remote.id
    tde_wallet_password = "BEstrO0ng_#11"

    pdb_creation_type_details {
        creation_type = "REMOTE_CLONE_PDB"
        source_pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
        source_container_database_admin_password = "BEstrO0ng_#11"
    }

    timeouts {
        create = "60m"
        delete = "2h"
    }
}

