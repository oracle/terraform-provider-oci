---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_asset_sources"
sidebar_current: "docs-oci-datasource-cloud_bridge-asset_sources"
description: |-
  Provides the list of Asset Sources in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_asset_sources
This data source provides the list of Asset Sources in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of asset sources.


## Example Usage

```hcl
data "oci_cloud_bridge_asset_sources" "test_asset_sources" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	asset_source_id = oci_cloud_bridge_asset_source.test_asset_source.id
	display_name = var.asset_source_display_name
	state = var.asset_source_state
}
```

## Argument Reference

The following arguments are supported:

* `asset_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asset source.
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The current state of the asset source.


## Attributes Reference

The following attributes are exported:

* `asset_source_collection` - The list of asset_source_collection.

### AssetSource Reference

The following attributes are exported:

* `are_historical_metrics_collected` - Flag indicating whether historical metrics are collected for assets, originating from this asset source.
* `are_realtime_metrics_collected` - Flag indicating whether real-time metrics are collected for assets, originating from this asset source.
* `assets_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `discovery_credentials` - Credentials for an asset source.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret in a vault. If the the type of the credentials is BASIC`, the secret must contain the username and password in JSON format, which is in the form of `{ "username": "<VMwareUser>", "password": "<VMwarePassword>" }`. 
	* `type` - Authentication type
* `discovery_schedule_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an attached discovery schedule.
* `display_name` - A user-friendly name for the asset source. Does not have to be unique, and it's mutable. Avoid entering confidential information. 
* `environment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the environment.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `inventory_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
* `lifecycle_details` - The detailed state of the asset source.
* `replication_credentials` - Credentials for an asset source.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret in a vault. If the the type of the credentials is BASIC`, the secret must contain the username and password in JSON format, which is in the form of `{ "username": "<VMwareUser>", "password": "<VMwarePassword>" }`. 
	* `type` - Authentication type
* `state` - The current state of the asset source.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the asset source was created in the RFC3339 format.
* `time_updated` - The point in time that the asset source was last updated in the RFC3339 format.
* `type` - The type of asset source. Indicates external origin of the assets that are read by assigning this asset source.
* `vcenter_endpoint` - Endpoint for VMware asset discovery and replication in the form of ```https://<host>:<port>/sdk```

