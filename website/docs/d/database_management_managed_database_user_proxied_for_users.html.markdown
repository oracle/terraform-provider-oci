---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_proxied_for_users"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_proxied_for_users"
description: |-
  Provides the list of Managed Database User Proxied For Users in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_proxied_for_users
This data source provides the list of Managed Database User Proxied For Users in Oracle Cloud Infrastructure Database Management service.

Gets the list of users on whose behalf the current user acts as proxy.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_proxied_for_users" "test_managed_database_user_proxied_for_users" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_proxied_for_user_name
	opc_named_credential_id = var.managed_database_user_proxied_for_user_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `proxied_for_user_collection` - The list of proxied_for_user_collection.

### ManagedDatabaseUserProxiedForUser Reference

The following attributes are exported:

* `items` - An array of user resources.
	* `authentication` - Indicates whether the proxy is required to supply the client credentials (YES) or not (NO).
	* `flags` - The flags associated with the proxy/client pair.
	* `name` - The name of a proxy user or the name of the client user.

