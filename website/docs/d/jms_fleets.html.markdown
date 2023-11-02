---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleets"
sidebar_current: "docs-oci-datasource-jms-fleets"
description: |-
  Provides the list of Fleets in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleets
This data source provides the list of Fleets in Oracle Cloud Infrastructure Jms service.

Returns a list of all the Fleets contained by a compartment. The query parameter `compartmentId`
is required unless the query parameter `id` is specified.


## Example Usage

```hcl
data "oci_jms_fleets" "test_fleets" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.fleet_display_name
	display_name_contains = var.fleet_display_name_contains
	id = var.fleet_id
	state = var.fleet_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `display_name` - (Optional) The display name.
* `display_name_contains` - (Optional) Filter the list with displayName contains the given value. 
* `id` - (Optional) The ID of the Fleet.
* `state` - (Optional) The state of the lifecycle.


## Attributes Reference

The following attributes are exported:

* `fleet_collection` - The list of fleet_collection.

### Fleet Reference

The following attributes are exported:

* `approximate_application_count` - The approximate count of all unique applications in the Fleet in the past seven days. This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag. 
* `approximate_installation_count` - The approximate count of all unique Java installations in the Fleet in the past seven days. This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag. 
* `approximate_java_server_count` - The approximate count of all unique Java servers in the Fleet in the past seven days. This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag. 
* `approximate_jre_count` - The approximate count of all unique Java Runtimes in the Fleet in the past seven days. This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag. 
* `approximate_managed_instance_count` - The approximate count of all unique managed instances in the Fleet in the past seven days. This metric is provided on a best-effort manner, and isn't taken into account when computing the resource ETag. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `description` - The Fleet's description.
* `display_name` - The name of the Fleet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `inventory_log` - Custom Log for inventory or operation log. 
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `is_advanced_features_enabled` - Whether or not advanced features are enabled in this Fleet. Deprecated, use `/fleets/{fleetId}/advanceFeatureConfiguration` API instead. 
* `is_export_setting_enabled` - Whether or not export setting is enabled in this Fleet. 
* `operation_log` - Custom Log for inventory or operation log. 
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `state` - The lifecycle state of the Fleet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the Fleet (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

