---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_applications"
sidebar_current: "docs-oci-datasource-generative_ai-hosted_applications"
description: |-
  Provides the list of Hosted Applications in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_hosted_applications
This data source provides the list of Hosted Applications in Oracle Cloud Infrastructure Generative AI service.

Lists the hosted applications in a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_hosted_applications" "test_hosted_applications" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.hosted_application_display_name
	id = var.hosted_application_id
	state = var.hosted_application_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
* `state` - (Optional) A filter to return only the hosted applications that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `hosted_application_collection` - The list of hosted_application_collection.

### HostedApplication Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID to create the hosted application in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - An optional description of the hosted application.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `environment_variables` - The list of environment variables for the Hosted Application.  Defines a list of environment variables injected at runtime.
	* `name` - Name of the environment variable.
	* `type` - Type of the environment variable (PLAINTEXT or HASHED, no default value).
	* `value` - Value of the environment variable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application.
* `inbound_auth_config` - The client-side inbound authentication configuration for the Hosted Application.  Defines the network access rules. When unspecified, the service applies the default inbound authentication configuration type.
	* `idcs_config` - Oracle Identity Cloud Service (IDCS) configuration used  when inboundAuthConfigType is set to IDCS_AUTH_CONFIG. This object must be specified when inboundAuthConfigType is IDCS_AUTH_CONFIG.
		* `audience` - Audience for IDCS.
		* `domain_url` - Domain URL for IDCS.
		* `scope` - Scope for IDCS.
	* `inbound_auth_config_type` - Inbound authentication configuration type of network access (IDCS_AUTH_CONFIG).
* `lifecycle_details` - A message describing the current state of the endpoint in more detail that can provide actionable information.
* `networking_config` - Networking configuration.
	* `inbound_networking_config` - Inbound Networking configuration.
		* `endpoint_mode` - inbounding from public or private endpoint.
		* `private_endpoint_id` - The [OCID] of Private Endpoint when endpointMode=Private
	* `outbound_networking_config` - Outbound Networking configuration.
		* `custom_subnet_id` - ocid of customer subnet when networkMode=Custom
		* `network_mode` - outbounding to managed internet or customer network.
		* `nsg_ids` - A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
* `scaling_config` - The auto scaling configuration for the Hosted Application.  Defines the minimum and maximum number of replicas. When unspecified, the service applies service-defined default scaling values.
	* `max_replica` - Maximum number of replicas allowed.
	* `min_replica` - Minimum number of replicas to keep running.
	* `scaling_type` - scaling type for application.
	* `target_concurrency_threshold` - number of simultaneous requests that can be processed by each replica.
	* `target_cpu_threshold` - Scale up if average CPU utilization exceeds this threshold.
	* `target_memory_threshold` - Scale up if average memory utilization exceeds this threshold.
	* `target_rps_threshold` - requests-per-second per replica of an application.
* `state` - The current state of the hosted application.
* `storage_configs` - A list of storageConfigs managed by the Oracle Cloud Infrastructure GenAI Platform and attached to the application.
	* `environment_variable_key` - The key of environment variable to store the database connection.
	* `storage_id` - The [OCID] of ApplicationStorage.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the hosted application was created, in the format defined by RFC 3339
* `time_updated` - The date and time the hosted application was updated, in the format defined by RFC 3339
