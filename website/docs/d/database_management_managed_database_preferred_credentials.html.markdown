---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_preferred_credentials"
sidebar_current: "docs-oci-datasource-database_management-managed_database_preferred_credentials"
description: |-
  Provides the list of Managed Database Preferred Credentials in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_preferred_credentials
This data source provides the list of Managed Database Preferred Credentials in Oracle Cloud Infrastructure Database Management service.

Gets the list of preferred credentials for a given Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_preferred_credentials" "test_managed_database_preferred_credentials" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.


## Attributes Reference

The following attributes are exported:

* `preferred_credential_collection` - The list of preferred_credential_collection.

### ManagedDatabasePreferredCredential Reference

The following attributes are exported:

* `credential_name` - The name of the preferred credential.
* `is_accessible` - Indicates whether the preferred credential is accessible.
* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential that contains the database user password metadata.
* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
* `role` - The role of the database user.
* `status` - The status of the preferred credential.
* `type` - The type of preferred credential. Only 'BASIC' is supported currently.
* `user_name` - The user name used to connect to the database.

