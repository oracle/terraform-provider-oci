---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_status"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_status"
description: |-
  Provides details about a specific Fusion Environment Status in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_status
This data source provides details about a specific Fusion Environment Status resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets the status of a Fusion environment identified by its OCID.

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_status" "test_fusion_environment_status" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier


## Attributes Reference

The following attributes are exported:

* `status` - The data plane status of FusionEnvironment.

