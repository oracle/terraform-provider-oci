---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_autonomous_vm_cluster_kms_management"
sidebar_current: "docs-oci-resource-database-cloud_autonomous_vm_cluster_kms_management"
description: |-
  Provides the Cloud Autonomous Vm Cluster K M S Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_autonomous_vm_cluster_kms_management
This resource provides the Cloud Autonomous Vm Cluster K M S Management resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/ChangeKeyManagementSystem

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

Update the Cloud Autonomous VM Cluster's key management system details where TDE keys of this cluster's ACDs will be stored.


## Example Usage

```hcl
resource "oci_database_cloud_autonomous_vm_cluster_kms_management" "test_cloud_autonomous_vm_cluster_kms_management" {
	#Required
	cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	key_management_system = var.cloud_autonomous_vm_cluster_kms_management_key_management_system

	#Optional
	is_exclusive = var.cloud_autonomous_vm_cluster_kms_management_is_exclusive
	key_store_id = oci_database_key_store.test_key_store.id
	thales_hsm_port = var.cloud_autonomous_vm_cluster_kms_management_thales_hsm_port
}
```

## Argument Reference

The following arguments are supported:

* `cloud_autonomous_vm_cluster_id` - (Required) The Cloud VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_exclusive` - (Optional) If true, ACDs within this Cluster cannot use a different key management system than what is configured in AVM Cluster.
* `key_management_system` - (Required) The type of key management system used.
* `key_store_id` - (Applicable when key_management_system=THALES) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `thales_hsm_port` - (Applicable when key_management_system=THALES) Thales HSM Port. Default is 9000. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Autonomous Vm Cluster K M S Management
	* `update` - (Defaults to 20 minutes), when updating the Cloud Autonomous Vm Cluster K M S Management
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Autonomous Vm Cluster K M S Management


## Import

Import is not supported for this resource.

