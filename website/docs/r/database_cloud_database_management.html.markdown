---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_database_management"
sidebar_current: "docs-oci-resource-database-cloud_database_management"
description: |-
  Provides the Database Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_database_management
This resource provides the Database Management resource in Oracle Cloud Infrastructure Database service.

Enable / Update / Disable database management for the specified Oracle Database instance.

Database Management requires `USER_NAME`, `PASSWORD_SECRET_ID` and `PRIVATE_END_POINT_ID`.
`database.0.database_management_config` is updated to appropriate managementType and managementStatus for the specified Oracle Database instance.

## Example Usage

```hcl
resource "oci_database_cloud_database_management" "test" {
  database_id           = oci_database_database.test_database.id
  management_type       = var.database_cloud_database_management_details_management_type
  private_end_point_id  = var.database_cloud_database_management_details_private_end_point_id
  service_name          = var.database_cloud_database_management_details_service_name
  credentialdetails {
    user_name           = var.database_cloud_database_management_details_user_name
    password_secret_id  = var.database_cloud_database_management_details_password_secret_id
  }
  enable_management     = var.database_cloud_database_management_details_enable_management
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `private_end_point_id` - (Required) (Updatable) The private end point [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `service_name` - (Required) (Updatable) Database service name
* `management_type` - (Required) (Updatable) Specifies database management type
  enum:
  - `BASIC`
  - `ADVANCED`
* `credentaildetails` - (Required) (Updatable) Credential details to connect to the database
    * `user_name` - Database username
    * `password_secret_id` - Specific database username's password [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `enable_management` - (Required) (Updatable) Use this flag to enable/disable database management

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `database_management_config` - Management config for specific database.
    * `management_status` - Database management status.
    * `management_type` - Management type (BASIC / ADVANCED)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Database Management
* `update` - (Defaults to 20 minutes), when updating the Database Management
* `delete` - (Defaults to 20 minutes), when destroying the Database Management


## Import

Import is not supported for this resource.

