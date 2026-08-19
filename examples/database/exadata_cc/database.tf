resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
  source       = "VM_CLUSTER_NEW"
  db_version   = "19.0.0.0"
  display_name = "createdDbHome"
}

resource "oci_database_database" "test_exacc_database"{
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbVMClus"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    pdb_name       = "pdbName"

    freeform_tags = {
      "Department" = "Finance"
    }

    db_backup_config {
      auto_backup_enabled = true
      auto_backup_window  = "SLOT_TWO"

      backup_destination_details {
        id   = oci_database_backup_destination.test_backup_destination_nfs.id
        type = "NFS"
      }
    }
    encryption_key_location_details {
        #Required
        hsm_password  = "hsmPassword"
        provider_type = "EXTERNAL"
    }
  }
  db_home_id = oci_database_db_home.test_db_home_vm_cluster.id
  source     = "NONE"
}

resource "oci_database_backup_destination" "test_backup_destination_nfs" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "testBackupDestination"
  type           = "NFS"

  #Optional

  freeform_tags = {
    "Department" = "Finance"
  }
  mount_type_details {
    local_mount_point_path = "localMountPointPath"
    mount_type             = "SELF_MOUNT"
  }
}

resource "oci_database_backup_destination" "test_backup_destination_zdlra" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "testBackupDestination"
  type           = "RECOVERY_APPLIANCE"

  # Required for RECOVERY_APPLIANCE destinations
  connection_string = "ra.host:1521/RA.SERVICE"
  vpc_users         = ["bkupUser1"]

  #Optional
  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_database" "test_exacc_database_zdlra" {
  db_home_id = oci_database_db_home.test_db_home_vm_cluster.id
  source     = "NONE"

  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbRA"

    db_backup_config {
      backup_destination_details {
        id   = oci_database_backup_destination.test_backup_destination_zdlra.id
        type = "RECOVERY_APPLIANCE"

        # Provide RA credentials for the selected VPC user
        vpc_user     = "bkupUser1"
        vpc_password = "secret"

        tde_wallet_backup_destination {
          backup_destination_type = "OSS"
        }
      }
    }
  }
}

data "oci_database_backup_destinations" "test_database_backup_destinations" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  type = "NFS"
}

data "oci_database_backup_destination" "test_database_backup_destination" {
  #Required
  backup_destination_id = oci_database_backup_destination.test_backup_destination_nfs.id
}