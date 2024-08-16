---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_addon"
sidebar_current: "docs-oci-resource-containerengine-addon"
description: |-
  Provides the Addon resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_addon
This resource provides the Addon resource in Oracle Cloud Infrastructure Container Engine service.

Install the specified addon for a cluster.

## Example Usage

```hcl
resource "oci_containerengine_addon" "test_addon" {
	#Required
	addon_name = oci_containerengine_addon.test_addon.name
	cluster_id = oci_containerengine_cluster.test_cluster.id
	remove_addon_resources_on_delete = true

	#Optional
	configurations {

		#Optional
		key = var.addon_configurations_key
		value = var.addon_configurations_value
	}
	override_existing = false
	version = var.addon_version
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Required) The name of the addon.
* `cluster_id` - (Required) The OCID of the cluster.
* `remove_addon_resources_on_delete` - (Required) Whether to remove addon resource in deletion.
* `configurations` - (Optional) (Updatable) Addon configuration details
	* `key` - (Optional) (Updatable) configuration key name
	* `value` - (Optional) (Updatable) configuration value name
* `override_existing` - (Optional) Whether or not to override an existing addon installation. Defaults to false. If set to true, any existing addon installation would be overridden as per new installation details.
* `version` - (Optional) (Updatable) The version of addon to be installed.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Addon
* `update` - (Defaults to 20 minutes), when updating the Addon
* `delete` - (Defaults to 20 minutes), when destroying the Addon


## Import

Addons can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_addon.test_addon "clusters/{clusterId}/addons/{addonName}" 
```
