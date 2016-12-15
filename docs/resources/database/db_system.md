# baremetal\_database\_db\_system

Provides an DBSystem resource.

## Example Usage

```
resource "baremetal_database_db_system" "t" {
  availability_domain = "availability_domain"
  compartment_id = "compartment_id"
  cpu_core_count = 2
  database_edition = "db_edition"
  db_home {
    database {
      "admin_password" = "apassword"
      "db_name" = "db_name"
    }
    db_version = "db_version"
    display_name = "display_name"
  }
  disk_redundancy = "disk_redundancy"
  shape = "shape"
  subnet_id = "subnet_id"
  ssh_public_keys = ["somesshkey"]
  display_name = "display_name"
  domain = "domain.com"
  hostname = "hostname"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain that the DB System is located in.
* `compartment_id` - (Required) The OCID of the compartment.
* `cpu_core_count` - (Required) The number of CPU cores enabled on the DB System.
* `database_edition` - (Optional) The Oracle Database Edition that applies to all the databases on the DB System.
* `db_home` - (Optional) Create DBHome details. See [Create DBHome Details](#create-dbhome-details) below for detials.
* `disk_redundancy` - (Optional) The type of redundancy configured for the DB System.
* `display_name` - (Optional) The user-friendly name for the DB System. It does not have to be unique.
* `domain` - (Optional) A domain name to assign to the DB System.
* `hostname` - (Optional) The host name to assign to the DB Node.
* `shape` - (Required) The shape of the DB System.
* `ssh_public_keys` - (Required) The public key portion of the key pair to use for SSH access to the DB System.
* `subnet_id` - (Required) The OCID of the subnet the DB System is associated with.

## Create DBHome Details

The following arguments are supported:

* `database` - (Required) Create Database details. See [Create Database Details](#create-database-details) below for details.
* `db_version` - (Required) A valid Oracle database version.
* `display_name` - (Optional) The user-provided name of the database home.

## Create Database Details

The following arguments are supported:

* `admin_password` - (Required) A strong password for SYS, SYSTEM, and PDB Admin.
* `db_name` - (Required) The database name.

## Attributes Reference

The following attributes are exported:

* `id` - The OCID of the DB System.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB System.
* `state` - The current state of the DB System.
* `time_created` - The date and time the DB System was created.
