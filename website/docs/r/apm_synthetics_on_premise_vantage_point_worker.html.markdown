---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_on_premise_vantage_point_worker"
sidebar_current: "docs-oci-resource-apm_synthetics-on_premise_vantage_point_worker"
description: |-
  Provides the On Premise Vantage Point Worker resource in Oracle Cloud Infrastructure Apm Synthetics service
---

# oci_apm_synthetics_on_premise_vantage_point_worker
This resource provides the On Premise Vantage Point Worker resource in Oracle Cloud Infrastructure Apm Synthetics service.

Registers a new worker.


## Example Usage

```hcl
resource "oci_apm_synthetics_on_premise_vantage_point_worker" "test_on_premise_vantage_point_worker" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	name = var.on_premise_vantage_point_worker_name
	on_premise_vantage_point_id = oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id
	resource_principal_token_public_key = var.on_premise_vantage_point_worker_resource_principal_token_public_key
	version = var.on_premise_vantage_point_worker_version

	#Optional
	configuration_details = var.on_premise_vantage_point_worker_configuration_details
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	priority = var.on_premise_vantage_point_worker_priority
	status = var.on_premise_vantage_point_worker_status
	worker_type = var.on_premise_vantage_point_worker_worker_type
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM domain ID the request is intended for. 
* `configuration_details` - (Optional) (Updatable) Configuration details of the On-premise VP worker.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) Unique On-premise VP worker name that cannot be edited. The name should not contain any confidential information.
* `on_premise_vantage_point_id` - (Required) The OCID of the On-premise vantage point.
* `priority` - (Optional) (Updatable) Priority of the On-premise VP worker to schedule monitors.
* `resource_principal_token_public_key` - (Required) public key for resource Principal Token based validation to be used in further calls.
* `status` - (Optional) (Updatable) Enables or disables the On-premise VP worker.
* `version` - (Required) Image version of the On-premise VP worker.
* `worker_type` - (Optional) Type of the On-premise VP worker.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `configuration_details` - Configuration details of the On-premise VP worker.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique On-premise VP worker name that cannot be edited. The name should not contain any confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `geo_info` - Geographical information of the On-premise VP worker.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the On-premise VP worker.
* `identity_info` - Domain details of the On-premise VP worker.
	* `apm_short_id` - Domain short id of the On-premise VP worker.
	* `collector_end_point` - Collector endpoint of the On-premise VP worker.
	* `region_name` - Domain region of the On-premise VP worker.
* `monitor_list` - Monitors list assigned to the On-premise VP worker.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitor.
	* `is_run_now` - If isRunNow is enabled, then the monitor will run immediately.
	* `monitor_type` - Type of monitor.
	* `time_assigned` - The time the resource was last assigned to an On-premise vantage point worker, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `name` - Unique permanent name of the On-premise VP worker. This is the same as the displayName.
* `opvp_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the On-premise vantage point.
* `opvp_name` - On-premise vantage point name.
* `priority` - Priority of the On-premise VP worker to schedule monitors.
* `runtime_id` - The runtime assigned id of the On-premise VP worker.
* `status` - Enables or disables the On-premise VP worker.
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_last_sync_up` - The time the resource was last synced, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `version_details` - Image version details of the On-premise VP worker.
	* `latest_version` - Latest image version of the On-premise VP worker.
	* `min_supported_version` - Minimum supported image version of the On-premise VP worker.
	* `version` - Current image version of the On-premise VP worker.
* `worker_type` - Type of the On-premise VP worker.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the On Premise Vantage Point Worker
	* `update` - (Defaults to 20 minutes), when updating the On Premise Vantage Point Worker
	* `delete` - (Defaults to 20 minutes), when destroying the On Premise Vantage Point Worker


## Import

OnPremiseVantagePointWorkers can be imported using the `id`, e.g.

```
$ terraform import oci_apm_synthetics_on_premise_vantage_point_worker.test_on_premise_vantage_point_worker "onPremiseVantagePoints/{onPremiseVantagePointId}/workers/{workerId}/apmDomainId/{apmDomainId}" 
```

