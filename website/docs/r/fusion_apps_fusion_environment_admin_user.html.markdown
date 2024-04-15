---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_admin_user"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment_admin_user"
description: |-
  Provides the Fusion Environment Admin User resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment_admin_user
This resource provides the Fusion Environment Admin User resource in Oracle Cloud Infrastructure Fusion Apps service.

Create a FusionEnvironment admin user

## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment_admin_user" "test_fusion_environment_admin_user" {
	#Required
	email_address = var.fusion_environment_admin_user_email_address
	first_name = var.fusion_environment_admin_user_first_name
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	last_name = var.fusion_environment_admin_user_last_name
	username = var.fusion_environment_admin_user_username

	#Optional
	password = var.fusion_environment_admin_user_password
}
```

## Argument Reference

The following arguments are supported:

* `email_address` - (Required) The email address for the administrator.
* `first_name` - (Required) The administrator's first name.
* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `last_name` - (Required) The administrator's last name.
* `password` - (Optional) The password for the administrator.
* `username` - (Required) The username for the administrator.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - A page of AdminUserSummary objects.
	* `email_address` - Admin users email address
	* `first_name` - Admin users first name
	* `last_name` - Admin users last name
	* `username` - Admin username

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment Admin User
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment Admin User
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment Admin User


## Import

FusionEnvironmentAdminUsers can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment_admin_user.test_fusion_environment_admin_user "fusionEnvironments/{fusionEnvironmentId}/adminUsers/{adminUsername}" 
```

