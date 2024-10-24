variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #required
  admin_password = "Ka2P7jb1_3?*##DZ"
  compartment_id = "ocid1.tenancy.oc1..aaaaaaaayxnndmkzcpajuuptcojciksbdrlwofhbe3hve6lypcqfgx56ssva"
  cpu_core_count = "6"
  data_storage_size_in_tbs = "1"
  db_name  = "testdbScheduledOperations"
  customer_contacts {
    email = "test1@oracle.com"
  }
  customer_contacts {
    email = "test2@oracle.com"
  }
  scheduled_operations {
    day_of_week {
      name = "MONDAY"
    }
    scheduled_start_time = "10:00"
    scheduled_stop_time = "18:00"
  }
  scheduled_operations {
      day_of_week {
        name = "TUESDAY"
      }
      scheduled_start_time = "10:00"
      scheduled_stop_time = "18:00"
    }
  scheduled_operations {
    day_of_week {
      name = "FRIDAY"
    }
    scheduled_start_time = "10:00"
    scheduled_stop_time = "18:00"
  }
}