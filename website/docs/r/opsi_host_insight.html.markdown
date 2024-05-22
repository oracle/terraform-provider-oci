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

Create a Host Insight resource for a host in Ops Insights. The host will be enabled in Ops Insights. Host metric collection and analysis will be started.


## Example Usage

```hcl
resource "oci_opsi_host_insight" "test_host_insight" {
	#Required
	compartment_id = var.compartment_id
	entity_source = var.host_insight_entity_source
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id

	#Optional
	compute_id = oci_opsi_compute.test_compute.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	enterprise_manager_entity_identifier = var.host_insight_enterprise_manager_entity_identifier
	enterprise_manager_identifier = var.host_insight_enterprise_manager_identifier
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
	freeform_tags = {"bar-key"= "value"}
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	status = "DISABLED"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of host
* `compute_id` - (Required when entity_source=MACS_MANAGED_CLOUD_HOST) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `enterprise_manager_bridge_id` - (Required when entity_source=EM_MANAGED_EXTERNAL_HOST) OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_identifier` - (Required when entity_source=EM_MANAGED_EXTERNAL_HOST) Enterprise Manager Entity Unique Identifier
* `enterprise_manager_identifier` - (Required when entity_source=EM_MANAGED_EXTERNAL_HOST) Enterprise Manager Unique Identifier
* `entity_source` - (Required) (Updatable) Source of the host entity.
* `exadata_insight_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_HOST) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `management_agent_id` - (Required when entity_source=MACS_MANAGED_CLOUD_HOST | MACS_MANAGED_EXTERNAL_HOST) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values. The resource destruction here is basically a soft delete. User cannot create resource using the same Management agent OCID. If resource is in enabled state during destruction, the resource will be disabled automatically before performing delete operation.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compute Instance
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
* `host_type` - Ops Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
* `opsi_private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
* `parent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster or DB System ID, depending on which configuration the resource belongs to.
* `platform_name` - Platform name.
* `platform_type` - Platform type. Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS]. Supported platformType(s) for MACS-managed cloud host insight: [LINUX]. Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX]. 
* `platform_version` - Platform version.
* `root_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure. This is the OCPU count for Autonomous Database and CPU core count for other database types.
* `root_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
* `state` - The current state of the host.
* `status` - Indicates the status of a host insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the host insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the host insight was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Host Insight
	* `update` - (Defaults to 20 minutes), when updating the Host Insight
	* `delete` - (Defaults to 20 minutes), when destroying the Host Insight


## Import

HostInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_host_insight.test_host_insight "id"
```

