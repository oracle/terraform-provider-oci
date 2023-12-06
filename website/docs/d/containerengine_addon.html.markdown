---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_addon"
sidebar_current: "docs-oci-datasource-containerengine-addon"
description: |-
  Provides details about a specific Addon in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_addon
This data source provides details about a specific Addon resource in Oracle Cloud Infrastructure Container Engine service.

Get the specified addon for a cluster.

## Example Usage

```hcl
data "oci_containerengine_addon" "test_addon" {
	#Required
	addon_name = oci_containerengine_addon.test_addon.name
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Required) The name of the addon.
* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `addon_error` - The error info of the addon.
	* `code` - A short error code that defines the upstream error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
	* `message` - A human-readable error string of the upstream error.
	* `status` - The status of the HTTP response encountered in the upstream error.
* `configurations` - Addon configuration details.
	* `key` - configuration key name
	* `value` - configuration value name
* `current_installed_version` - current installed version of the addon
* `addon_name` - The name of the addon.
* `state` - The state of the addon.
* `time_created` - The time the cluster was created.
* `version` - selected addon version, or null indicates autoUpdate