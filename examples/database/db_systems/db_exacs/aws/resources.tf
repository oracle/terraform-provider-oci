resource "oci_database_cloud_vm_cluster" "aws_cluster" {
    # import this resource because we need identity connector set up which is a multi cloud dependency

}

resource "oci_database_db_home" "test_db_home_vm_cluster_no_db" {
  vm_cluster_id = oci_database_cloud_vm_cluster.aws_cluster.id

  # VM_CLUSTER_BACKUP can also be specified as a source for cloud VM clusters.
  source       = "VM_CLUSTER_NEW"
  db_version   = "23.0.0.0.0"
  display_name = "createdDbHomeNoDb"
}

resource "oci_database_database" "test_database_create" {
  #Required
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "TFdb1"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"

    db_backup_config {
      auto_backup_enabled = false
    }

    encryption_key_location_details {
        #Required
        azure_encryption_key_id  =  var.aws_encryption_key
        provider_type = "AWS"
    }
  }

  db_home_id = oci_database_db_home.test_db_home_vm_cluster_no_db.id
  source     = "NONE"
}


resource "oci_database_database" "test_database_migrate" {
  #Required
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "TFdb2"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"

    db_backup_config {
      auto_backup_enabled = false
    }
/*
    encryption_key_location_details {
        #Required
        aws_encryption_key_id  = var.aws_encryption_key
        provider_type = "AWS"
    }
*/
  }

  db_home_id = oci_database_db_home.test_db_home_vm_cluster_no_db.id
  source     = "NONE"
}