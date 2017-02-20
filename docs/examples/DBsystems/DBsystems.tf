variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  private_key = "${var.private_key}"
}

resource "baremetal_database_db_system" "dev_db_1" {
  availability_domain = "Uocm:PHX-AD-1"
  compartment_id = "${var.compartment_ocid}"
  cpu_core_count = 2
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      "admin_password" = "BEstrO0ng_#11"
      "db_name" = "mytfdb"
    }
    db_version = "12.1.0.2"
    display_name = "my-dev-db"
  }
  disk_redundancy = "HIGH"
  shape = "BM.DenseIO1.36"
  subnet_id = "ocid1.subnet.oc1.phx.aaaaaaaay6exxocfkho64s56qk6q2xrhshajbwursc635v2b7oxnz5i7udjq"
  ssh_public_keys = ["somesshkey"]
  display_name = "display_name"
  domain = "mycompany.com"
  hostname = "my-dev-db"
}
