---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_asm_users"
sidebar_current: "docs-oci-datasource-database_management-external_asm_users"
description: |-
  Provides the list of External Asm Users in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_asm_users
This data source provides the list of External Asm Users in Oracle Cloud Infrastructure Database Management service.

Lists ASM users for the external ASM specified by `externalAsmId`.


## Example Usage

```hcl
data "oci_database_management_external_asm_users" "test_external_asm_users" {
	#Required
	external_asm_id = oci_database_management_external_asm.test_external_asm.id

	#Optional
	opc_named_credential_id = var.external_asm_user_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `external_asm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `external_asm_user_collection` - The list of external_asm_user_collection.

### ExternalAsmUser Reference

The following attributes are exported:

* `items` - An array of external ASM users.
	* `asm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.
	* `name` - The name of the ASM user.
	* `privileges` - The list of privileges of the ASM user.

