---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_asm_configuration"
sidebar_current: "docs-oci-datasource-database_management-external_asm_configuration"
description: |-
  Provides details about a specific External Asm Configuration in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_asm_configuration
This data source provides details about a specific External Asm Configuration resource in Oracle Cloud Infrastructure Database Management service.

Gets configuration details including disk groups for the external ASM specified by `externalAsmId`.


## Example Usage

```hcl
data "oci_database_management_external_asm_configuration" "test_external_asm_configuration" {
	#Required
	external_asm_id = oci_database_management_external_asm.test_external_asm.id

	#Optional
	opc_named_credential_id = var.external_asm_configuration_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `external_asm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `init_parameters` - An array of initialization parameters for the external ASM instances.
	* `asm_instance_display_name` - The user-friendly name for the ASM instance. The name does not have to be unique.
	* `asm_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM instance.
	* `auto_mount_disk_groups` - The list of disk group names that an ASM instance mounts at startup or when the `ALTER DISKGROUP ALL MOUNT` statement is issued. 
	* `disk_discovery_path` - An operating system-dependent value used to limit the set of disks considered for discovery. 
	* `preferred_read_failure_groups` - The list of failure groups that contain preferred read disks.
	* `rebalance_power` - The maximum power on an ASM instance for disk rebalancing. 

