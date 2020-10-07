---
layout: "oci"
page_title: "Database DB System Migration"
sidebar_current: "docs-oci-guide-database_db_system_migration"
description: |-
  The Oracle Cloud Infrastructure provider. Database DB system migration
---

## Database DB System Migration

The X8M generation of Exadata hardware introduces a new resource model that replaces the Exadata DB system. The new resource model uses new APIs to provision and manage its resources. The existing DB system APIs for Exadata will be deprecated by Oracle Cloud Infrastructure for all users following written notification and a transition period allowing you to switch to the new API and Console interfaces.

If you have existing Exadata DB systems in Oracle Cloud Infrastructure, you can use Terraform to switch them to the new resource model and APIs.

### Warning 
Switching an Exadata DB system to the new resource model and APIs cannot be reversed. If you have automation for your system that utilizes the DB system APIs, you may need to update your applications prior to switching.

## Switching to the new resource model:
* Does not impact the DB system's existing Exadata databases or client connections. 
* Does not change the underlying hardware or shape family of your Exadata Cloud Service instance.
* Will not affect bare metal and virtual DB systems

After converting your DB system, you will have two new resources in place of the DB system resource: a cloud Exadata infrastructure resource, and a cloud VM cluster resource.

## What to Expect After Switching

* Your new cloud Exadata infrastructure resource and cloud VM cluster are created in the same compartment as the DB system they replace
* Your new cloud Exadata infrastructure resource and cloud VM cluster use the same networking configuration as the DB system they replace
* After the switch, you cannot perform operations on the old Exadata DB system resource
* Switching is permanent, and the change cannot be undone
* X6, X7, X8 and Exadata base systems retain their fixed shapes after the switch, and cannot be expanded

## To switch an Exadata DB system to the new Exadata resource model

These migration steps use the following example, which shows an existing Exadata Cloud Service instance using the old DB system resource model:

```hcl
resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  cpu_core_count      = var.cpu_core_count
  database_edition    = var.db_edition
  time_zone           = var.time_zone

  db_home {
    database {
      admin_password = var.db_admin_password
      db_name        = "TFdb1Exa"
      character_set  = var.character_set
      ncharacter_set = var.n_character_set
      db_workload    = var.db_workload
      pdb_name       = var.pdb_name

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = var.db_version
    display_name = "MyTFDBHome1Exa"
  }

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day       = ["4"]
    lead_time_in_weeks = 2

    months {
      name = "APRIL"
    }

    weeks_of_month = ["2"]
  }

  disk_redundancy  = var.db_disk_redundancy
  shape            = var.db_system_shape
  subnet_id        = oci_core_subnet.subnet.id
  backup_subnet_id = oci_core_subnet.subnet_backup.id
  ssh_public_keys  = [var.ssh_public_key]
  display_name     = var.db_system_display_name
  sparse_diskgroup = var.sparse_diskgroup

  hostname                = var.hostname
  data_storage_percentage = var.data_storage_percentage

  #data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model          = var.license_model
  node_count             = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
  backup_network_nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
  nsg_ids                = [oci_core_network_security_group.test_network_security_group_backup.id, oci_core_network_security_group.test_network_security_group.id]
}
```

To migrate the system to the new resource model, first create the `oci_database_migration` resource:

```hcl
// This is 1 time action to migrate test_db_system into db ExaCS
// and the test_db_system will become `Migrated`
resource "oci_database_migration" "test_migration" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

Provisioning the `oci_database_migration` resource creates two new resources: `oci_database_cloud_exadata_infrastructure` and `oci_database_cloud_vm_cluster`.

You can get OCIDs of these two resources from the `oci_database_migration` resource:
```hcl
output "cloud_exadata_infrastructure_id" {
  value = oci_database_migration.test_migration.cloud_exadata_infrastructure_id
}

output "cloud_vm_cluster_id" {
  value = oci_database_migration.test_migration.cloud_vm_cluster_id
}
```
For the two new resources, you need to create Terraform config:
```hcl
resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure"{}

resource "oci_database_cloud_vm_cluster" "test_cloud_vm_cluster" {}
```
Then run the Terraform import command:
```text
terraform import oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure <cloud_exadata_infrastructure_id>
terraform import oci_database_cloud_vm_cluster.test_cloud_vm_cluster <cloud_vm_cluster_id>
```

After switching to the new Exadata resource model, remove the old `oci_database_db_system` config.
Terraform now manages the two new resources.

### Tip
After the migration, you can use [Resource Discovery](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/resource_discovery) to create a full configuration and state file for importing these two new resources.