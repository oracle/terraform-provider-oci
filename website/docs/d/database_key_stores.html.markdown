---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_key_stores"
sidebar_current: "docs-oci-datasource-database-key_stores"
description: |-
  Provides the list of Key Stores in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_key_stores
This data source provides the list of Key Stores in Oracle Cloud Infrastructure Database service.

Gets a list of key stores in the specified compartment.


## Example Usage

```hcl
data "oci_database_key_stores" "test_key_stores" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `key_stores` - The list of key_stores.

### KeyStore Reference

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

