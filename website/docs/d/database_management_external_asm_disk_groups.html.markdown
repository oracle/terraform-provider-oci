---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_asm_disk_groups"
sidebar_current: "docs-oci-datasource-database_management-external_asm_disk_groups"
description: |-
  Provides the list of External Asm Disk Groups in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_asm_disk_groups
This data source provides the list of External Asm Disk Groups in Oracle Cloud Infrastructure Database Management service.

Lists ASM disk groups for the external ASM specified by `externalAsmId`.


## Example Usage

```hcl
data "oci_database_management_external_asm_disk_groups" "test_external_asm_disk_groups" {
	#Required
	external_asm_id = oci_database_management_external_asm.test_external_asm.id

	#Optional
	opc_named_credential_id = var.external_asm_disk_group_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `external_asm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `external_asm_disk_group_collection` - The list of external_asm_disk_group_collection.

### ExternalAsmDiskGroup Reference

The following attributes are exported:

* `items` - An array of external ASM disk groups.
	* `databases` - The unique names of the databases using the disk group.
	* `dismounting_instance_count` - The number of ASM instances that have the disk group in dismounted state.
	* `is_sparse` - Indicates whether the disk group is a sparse disk group or not.
	* `mounting_instance_count` - The number of ASM instances that have the disk group in mounted state.
	* `name` - The name of the ASM disk group.
	* `redundancy_type` - The redundancy type of the disk group.
	* `total_size_in_mbs` - The total capacity of the disk group (in megabytes).
	* `used_percent` - The percentage of used space in the disk group.
	* `used_size_in_mbs` - The used capacity of the disk group (in megabytes).

