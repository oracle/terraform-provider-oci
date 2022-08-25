---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_admin_users"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_admin_users"
description: |-
  Provides the list of Fusion Environment Admin Users in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_admin_users
This data source provides the list of Fusion Environment Admin Users in Oracle Cloud Infrastructure Fusion Apps service.

List all FusionEnvironment admin users

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_admin_users" "test_fusion_environment_admin_users" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `admin_user_collection` - The list of admin_user_collection.

### FusionEnvironmentAdminUser Reference

The following attributes are exported:

* `items` - A page of AdminUserSummary objects.
	* `email_address` - Admin users email address
	* `first_name` - Admin users first name
	* `last_name` - Admin users last name
	* `username` - Admin username

