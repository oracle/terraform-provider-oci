---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_host_insights"
sidebar_current: "docs-oci-datasource-opsi-host_insights"
description: |-
  Provides the list of Host Insights in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_host_insights
This data source provides the list of Host Insights in Oracle Cloud Infrastructure Opsi service.

Gets a list of host insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
When both compartmentId and compartmentIdInSubtree are specified, a list of host insights in that compartment and in all sub-compartments will be returned.


## Example Usage

```hcl
data "oci_opsi_host_insights" "test_host_insights" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.host_insight_compartment_id_in_subtree
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
	host_type = var.host_insight_host_type
	id = var.host_insight_id
	state = var.host_insight_state
	status = var.host_insight_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments. 
* `enterprise_manager_bridge_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_HOST) Unique Enterprise Manager bridge identifier
* `exadata_insight_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_HOST) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource. 
* `host_type` - (Optional) Filter by one or more host types. Possible value is EXTERNAL-HOST.
* `id` - (Optional) Optional list of host insight resource [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `platform_type` - (Optional) Filter by one or more platform types. Supported platformType(s) for MACS-managed external host insight: [LINUX]. Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS].
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `host_insight_summary_collection` - The list of host_insight_summary_collection.

### HostInsight Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_display_name` - Enterprise Manager Entity Display Name
* `enterprise_manager_entity_identifier` - Enterprise Manager Entity Unique Identifier
* `enterprise_manager_entity_name` - Enterprise Manager Entity Name
* `enterprise_manager_entity_type` - Enterprise Manager Entity Type
* `enterprise_manager_identifier` - Enterprise Manager Unique Identifier
* `entity_source` - Source of the host entity.
* `exadata_insight_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_display_name` - The user-friendly name for the host. The name does not have to be unique.
* `host_name` - The host name. The host name is unique amongst the hosts managed by the same management agent.
* `host_type` - Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
* `platform_name` - Platform name.
* `platform_type` - Platform type. Supported platformType(s) for MACS-managed external host insight: [LINUX]. Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS]. 
* `platform_version` - Platform version.
* `processor_count` - Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
* `state` - The current state of the host.
* `status` - Indicates the status of a host insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the host insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the host insight was updated. An RFC3339 formatted datetime string

