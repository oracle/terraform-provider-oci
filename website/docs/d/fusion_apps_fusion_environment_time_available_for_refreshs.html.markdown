---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_time_available_for_refreshs"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_time_available_for_refreshs"
description: |-
  Provides the list of Fusion Environment Time Available For Refreshs in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_time_available_for_refreshs
This data source provides the list of Fusion Environment Time Available For Refreshs in Oracle Cloud Infrastructure Fusion Apps service.

Gets available refresh time for this fusion environment

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_time_available_for_refreshs" "test_fusion_environment_time_available_for_refreshs" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `time_available_for_refresh_collection` - The list of time_available_for_refresh_collection.

### FusionEnvironmentTimeAvailableForRefresh Reference

The following attributes are exported:

* `items` - A list of available refresh time objects.
	* `time_available_for_refresh` - refresh time.

