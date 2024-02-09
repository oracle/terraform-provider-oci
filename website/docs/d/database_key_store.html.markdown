---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_key_store"
sidebar_current: "docs-oci-datasource-database-key_store"
description: |-
  Provides details about a specific Key Store in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_key_store
This data source provides details about a specific Key Store resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified key store.


## Example Usage

```hcl
data "oci_database_key_store" "test_key_store" {
	#Required
	key_store_id = oci_database_key_store.test_key_store.id
}
```

## Argument Reference

The following arguments are supported:

* `key_store_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.


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

