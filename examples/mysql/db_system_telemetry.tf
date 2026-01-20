variable "mysql_version" {
  default = "9.6.1"
}

variable "log_group_ocid" {
  # The OCID of an existing OCI Log Analytics log group where MySQL logs will be sent to.
  # This log group must be created before running this example.
}


resource "oci_mysql_mysql_db_system" "test_mysql_db_system_telemetry" {
  display_name        = "db-system-with-telemetry"
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  shape_name          = "MySQL.2"
  subnet_id           = oci_core_subnet.test_subnet.id
  mysql_version       = var.mysql_version

  telemetry_configuration {
    logs {
      destination = "LOG_ANALYTICS"
      destination_configurations {
        key   = "log-group-id"
        value = var.log_group_ocid
      }
      destination_configurations {
        key   = "log-set"
        value = "logset1"
      }
      log_types = ["ERROR_LOG"]
    }
  }
}