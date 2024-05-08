---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_wlp_agent"
sidebar_current: "docs-oci-datasource-cloud_guard-wlp_agent"
description: |-
  Provides details about a specific Wlp Agent in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_wlp_agent
This data source provides details about a specific Wlp Agent resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a WlpAgent resource for an on-premise resource identified by wlpAgentId.

## Example Usage

```hcl
data "oci_cloud_guard_wlp_agent" "test_wlp_agent" {
	#Required
	wlp_agent_id = oci_cloud_guard_wlp_agent.test_wlp_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `wlp_agent_id` - (Required) WLP agent OCID.


## Attributes Reference

The following attributes are exported:

* `agent_version` - The version of the agent
* `certificate_id` - The certificate ID returned by Oracle Cloud Infrastructure certificates service
* `certificate_signed_request` - The updated certificate signing request
* `compartment_id` - Compartment OCID of WlpAgent.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `host_id` - OCID for instance in which WlpAgent is installed
* `id` - OCID for WlpAgent
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - TenantId of the host
* `time_created` - The date and time the WlpAgent was created. Format defined by RFC3339.
* `time_updated` - The date and time the WlpAgent was updated. Format defined by RFC3339.

