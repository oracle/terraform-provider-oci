---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_saas_admin_user"
sidebar_current: "docs-oci-resource-database-autonomous_database_saas_admin_user"
description: |-
  Provides the Autonomous Database Saas Admin User resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_saas_admin_user
This resource creates and enables the Autonomous Database administrative user account in Oracle Cloud Infrastructure Database service.

## Example Usage

```hcl
resource "oci_database_autonomous_database_saas_admin_user" "test_autonomous_database_saas_admin_user" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	password = var.autonomous_database_saas_admin_user_password

	#Optional
	access_type = var.autonomous_database_saas_admin_user_access_type
	duration = var.autonomous_database_saas_admin_user_duration
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Optional) The access type for the SaaS administrative user. If no access type is specified, the READ_ONLY access type is used.
* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `duration` - (Optional) How long, in hours, the SaaS administrative user will stay enabled. If no duration is specified, the default value 1 will be used.
* `password` - (Optional) A strong password for SaaS administrative user. The password must be a minimum of nine (9) characters and contain a minimum of two (2) uppercase, two (2) lowercase, two (2) numbers, and two (2) special characters from _ (underscore), \# (hashtag), or - (dash). The password is mandatory if "secret_id" is not present.
* `secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). The secret is mandatory if "password" is not present.
* `secret_version_number` - (Optional) The version of the vault secret. If no version is specified, the latest version will be used.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Database Saas Admin User
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Database Saas Admin User
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Database Saas Admin User


## Import

Import is not supported for this resource.

