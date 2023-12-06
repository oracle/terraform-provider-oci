---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_family_limits_and_usage"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_family_limits_and_usage"
description: |-
  Provides details about a specific Fusion Environment Family Limits And Usage in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_family_limits_and_usage
This data source provides details about a specific Fusion Environment Family Limits And Usage resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets the number of environments (usage) of each type in the fusion environment family, as well as the limit that's allowed to be created based on the group's associated subscriptions.

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_family_limits_and_usage" "test_fusion_environment_family_limits_and_usage" {
	#Required
	fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_family_id` - (Required) The unique identifier (OCID) of the FusionEnvironmentFamily.


## Attributes Reference

The following attributes are exported:

* `development_limit_and_usage` - The limit and usage for a specific environment type, for example, production, development, or test.
	* `usage` - The usage of current environment.
* `production_limit_and_usage` - The limit and usage for a specific environment type, for example, production, development, or test.
	* `usage` - The usage of current environment.
* `test_limit_and_usage` - The limit and usage for a specific environment type, for example, production, development, or test.
	* `usage` - The usage of current environment.

