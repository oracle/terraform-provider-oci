---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_admin_user"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_admin_user"
description: |-
  Provides details about a specific Fusion Environment Admin User in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_admin_user
This data source provides details about a specific Fusion Environment Admin User resource in Oracle Cloud Infrastructure Fusion Apps service.

List all FusionEnvironment admin users

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_admin_user" "test_fusion_environment_admin_user" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `items` - A page of AdminUserSummary objects.
	* `email_address` - Admin users email address
	* `first_name` - Admin users first name
	* `last_name` - Admin users last name
	* `username` - Admin username

