---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_key_store"
sidebar_current: "docs-oci-resource-database-key_store"
description: |-
  Provides the Key Store resource in Oracle Cloud Infrastructure Database service
---

# oci_database_key_store
This resource provides the Key Store resource in Oracle Cloud Infrastructure Database service.

Creates a Key Store.


## Example Usage

```hcl
resource "oci_database_key_store" "test_key_store" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.key_store_display_name
	type_details {
		#Required
		admin_username = var.key_store_type_details_admin_username
		connection_ips = var.key_store_type_details_connection_ips
		secret_id = oci_vault_secret.test_secret.id
		type = var.key_store_type_details_type
		vault_id = oci_kms_vault.test_vault.id
	}

	#Optional
	defined_tags = var.key_store_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the key store. The name does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type_details` - (Required) (Updatable) Key store type details.
	* `admin_username` - (Required) (Updatable) The administrator username to connect to Oracle Key Vault
	* `connection_ips` - (Required) (Updatable) The list of Oracle Key Vault connection IP addresses.
	* `secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	* `type` - (Required) (Updatable) The type of key store.
	* `vault_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `associated_databases` - List of databases associated with the key store.
	* `db_name` - The name of the database that is associated with the key store.
	* `id` - The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the key store. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the key store.
* `time_created` - The date and time that the key store was created.
* `type_details` - Key store type details.
	* `admin_username` - The administrator username to connect to Oracle Key Vault
	* `connection_ips` - The list of Oracle Key Vault connection IP addresses.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	* `type` - The type of key store.
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Key Store
	* `update` - (Defaults to 20 minutes), when updating the Key Store
	* `delete` - (Defaults to 20 minutes), when destroying the Key Store


## Import

KeyStores can be imported using the `id`, e.g.

```
$ terraform import oci_database_key_store.test_key_store "id"
```

