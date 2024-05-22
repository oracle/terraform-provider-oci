---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_host_insight"
sidebar_current: "docs-oci-datasource-opsi-host_insight"
description: |-
  Provides details about a specific Host Insight in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_host_insight
This data source provides details about a specific Host Insight resource in Oracle Cloud Infrastructure Opsi service.

Gets details of a host insight.

## Example Usage

```hcl
data "oci_opsi_host_insight" "test_host_insight" {
	#Required
	host_insight_id = oci_opsi_host_insight.test_host_insight.id
}
```

## Argument Reference

The following arguments are supported:

* `host_insight_id` - (Required) Unique host insight identifier


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
* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
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
* `root_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
* `root_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
* `state` - The current state of the host.
* `status` - Indicates the status of a host insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the host insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the host insight was updated. An RFC3339 formatted datetime string

