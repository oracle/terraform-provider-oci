---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_asm_users"
sidebar_current: "docs-oci-datasource-database_management-cloud_asm_users"
description: |-
  Provides the list of Cloud Asm Users in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_asm_users
This data source provides the list of Cloud Asm Users in Oracle Cloud Infrastructure Database Management service.

Lists ASM users for the cloud ASM specified by `cloudAsmId`.


## Example Usage

```hcl
data "oci_database_management_cloud_asm_users" "test_cloud_asm_users" {
	#Required
	cloud_asm_id = oci_database_management_cloud_asm.test_cloud_asm.id

	#Optional
	opc_named_credential_id = var.cloud_asm_user_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_asm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.


## Attributes Reference

The following attributes are exported:

* `cloud_asm_user_collection` - The list of cloud_asm_user_collection.

### CloudAsmUser Reference

The following attributes are exported:

* `items` - An array of cloud ASM users.
	* `asm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
	* `name` - The name of the ASM user.
	* `privileges` - The list of privileges of the ASM user.

