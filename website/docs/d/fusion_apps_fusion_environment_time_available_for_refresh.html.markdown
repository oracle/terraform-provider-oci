---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_time_available_for_refresh"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_time_available_for_refresh"
description: |-
  Provides details about a specific Fusion Environment Time Available For Refresh in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_time_available_for_refresh
This data source provides details about a specific Fusion Environment Time Available For Refresh resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets available refresh time for this fusion environment

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_time_available_for_refresh" "test_fusion_environment_time_available_for_refresh" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `items` - A list of available refresh time objects.
	* `time_available_for_refresh` - refresh time.

