---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_host_insight"
sidebar_current: "docs-oci-resource-opsi-host_insight"
description: |-
  Provides the Host Insight resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_host_insight
This resource provides the Host Insight resource in Oracle Cloud Infrastructure Opsi service.

Create a Host Insight resource for a host in Operations Insights. The host will be enabled in Operations Insights. Host metric collection and analysis will be started.


## Example Usage

```hcl
resource "oci_opsi_host_insight" "test_host_insight" {
	#Required
	compartment_id = var.compartment_id
	entity_source = var.host_insight_entity_source
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = 'DISABLED'
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of host
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `entity_source` - (Required) (Updatable) Source of the host entity.
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `management_agent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values. The resource destruction here is basically a soft delete. User cannot create resource using the same Management agent OCID. If resource is in enabled state during destruction, the resource will be disabled automatically before performing delete operation.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `entity_source` - Source of the host entity.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_display_name` - The user-friendly name for the host. The name does not have to be unique.
* `host_name` - The host name. The host name is unique amongst the hosts managed by the same management agent.
* `host_type` - Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
* `platform_name` - Platform name.
* `platform_type` - Platform type.
* `platform_version` - Platform version.
* `processor_count` - Processor count.
* `state` - The current state of the host.
* `status` - Indicates the status of a host insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the host insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the host insight was updated. An RFC3339 formatted datetime string

## Import

HostInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_host_insight.test_host_insight "id"
```

