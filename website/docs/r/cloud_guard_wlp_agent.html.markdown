---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_wlp_agent"
sidebar_current: "docs-oci-resource-cloud_guard-wlp_agent"
description: |-
  Provides the Wlp Agent resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_wlp_agent
This resource provides the Wlp Agent resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates and registers a WLP agent for an
on-premise resource.


## Example Usage

```hcl
resource "oci_cloud_guard_wlp_agent" "test_wlp_agent" {
	#Required
	agent_version = var.wlp_agent_agent_version
	certificate_signed_request = var.wlp_agent_certificate_signed_request
	compartment_id = var.compartment_id
	os_info = var.wlp_agent_os_info

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `agent_version` - (Required) The version of the agent making the request
* `certificate_signed_request` - (Required) (Updatable) The certificate signed request containing domain, organization names, organization units, city, state, country, email and public key, among other certificate details, signed by private key
* `compartment_id` - (Required) Compartment OCID of the host
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `os_info` - (Required) Concatenated OS name, OS version and agent architecture; for example, ubuntu_22.0_amd64.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Wlp Agent
	* `update` - (Defaults to 20 minutes), when updating the Wlp Agent
	* `delete` - (Defaults to 20 minutes), when destroying the Wlp Agent


## Import

WlpAgents can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_wlp_agent.test_wlp_agent "id"
```

