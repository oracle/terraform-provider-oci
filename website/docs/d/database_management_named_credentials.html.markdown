---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_named_credentials"
sidebar_current: "docs-oci-datasource-database_management-named_credentials"
description: |-
  Provides the list of Named Credentials in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_named_credentials
This data source provides the list of Named Credentials in Oracle Cloud Infrastructure Database Management service.

Gets a single named credential specified by the name or all the named credentials in a specific compartment.


## Example Usage

```hcl
data "oci_database_management_named_credentials" "test_named_credentials" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	associated_resource = var.named_credential_associated_resource
	name = var.named_credential_name
	scope = var.named_credential_scope
	type = var.named_credential_type
}
```

## Argument Reference

The following arguments are supported:

* `associated_resource` - (Optional) The resource associated to the named credential.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) The name of the named credential.
* `scope` - (Optional) The scope of named credential.
* `type` - (Optional) The type of database that is associated to the named credential.


## Attributes Reference

The following attributes are exported:

* `named_credential_collection` - The list of named_credential_collection.

### NamedCredential Reference

The following attributes are exported:

* `associated_resource` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource that  is associated to the named credential. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `content` - The details of the named credential.
	* `credential_type` - The type of named credential. Only 'BASIC' is supported currently.
	* `password_secret_access_mode` - The mechanism used to access the password plain text value.
	* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	* `role` - The role of the database user.
	* `user_name` - The user name used to connect to the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The information specified by the user about the named credential.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential.
* `lifecycle_details` - The details of the lifecycle state.
* `name` - The name of the named credential.
* `scope` - The scope of the named credential.
* `state` - The current lifecycle state of the named credential.
* `time_created` - The date and time the named credential was created.
* `time_updated` - The date and time the named credential was last updated.
* `type` - The type of resource associated with the named credential.

