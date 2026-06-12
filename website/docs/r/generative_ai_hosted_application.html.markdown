---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_application"
sidebar_current: "docs-oci-resource-generative_ai-hosted_application"
description: |-
  Provides the Hosted Application resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_hosted_application
This resource provides the Hosted Application resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/HostedApplication

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a hosted application.

## Example Usage

```hcl
resource "oci_generative_ai_hosted_application" "test_hosted_application" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.hosted_application_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.hosted_application_description
	environment_variables {
		#Required
		name = var.hosted_application_environment_variables_name
		type = var.hosted_application_environment_variables_type
		value = var.hosted_application_environment_variables_value
	}
	freeform_tags = {"Department"= "Finance"}
	inbound_auth_config {
		#Required
		inbound_auth_config_type = var.hosted_application_inbound_auth_config_inbound_auth_config_type

		#Optional
		idcs_config {
			#Required
			domain_url = var.hosted_application_inbound_auth_config_idcs_config_domain_url
			scope = var.hosted_application_inbound_auth_config_idcs_config_scope

			#Optional
			audience = var.hosted_application_inbound_auth_config_idcs_config_audience
		}
	}
	networking_config {
		#Required
		inbound_networking_config {
			#Required
			endpoint_mode = var.hosted_application_networking_config_inbound_networking_config_endpoint_mode

			#Optional
			private_endpoint_id = oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id
		}
		outbound_networking_config {
			#Required
			network_mode = var.hosted_application_networking_config_outbound_networking_config_network_mode

			#Optional
			custom_subnet_id = oci_core_subnet.test_subnet.id
			nsg_ids = var.hosted_application_networking_config_outbound_networking_config_nsg_ids
		}
	}
	scaling_config {
		#Required
		scaling_type = var.hosted_application_scaling_config_scaling_type

		#Optional
		max_replica = var.hosted_application_scaling_config_max_replica
		min_replica = var.hosted_application_scaling_config_min_replica
		target_concurrency_threshold = var.hosted_application_scaling_config_target_concurrency_threshold
		target_cpu_threshold = var.hosted_application_scaling_config_target_cpu_threshold
		target_memory_threshold = var.hosted_application_scaling_config_target_memory_threshold
		target_rps_threshold = var.hosted_application_scaling_config_target_rps_threshold
	}
	storage_configs {
		#Required
		environment_variable_key = var.hosted_application_storage_configs_environment_variable_key
		storage_id = oci_generative_ai_storage.test_storage.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment OCID for the Hosted Application.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) The description for the Hosted Application.
* `display_name` - (Required) (Updatable) The user-friendly display name for the Hosted Application.  Does not need to be unique and can be updated after creation.
* `environment_variables` - (Optional) (Updatable) The list of environment variables for the Hosted Application.  Defines a list of environment variables injected at runtime.
	* `name` - (Required) (Updatable) Name of the environment variable.
	* `type` - (Required) (Updatable) Type of the environment variable (PLAINTEXT or HASHED, no default value).
	* `value` - (Required) (Updatable) Value of the environment variable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `inbound_auth_config` - (Optional) (Updatable) The client-side inbound authentication configuration for the Hosted Application.  Defines the network access rules. When unspecified, the service applies the default inbound authentication configuration type.
	* `idcs_config` - (Optional) (Updatable) Oracle Identity Cloud Service (IDCS) configuration used  when inboundAuthConfigType is set to IDCS_AUTH_CONFIG. This object must be specified when inboundAuthConfigType is IDCS_AUTH_CONFIG.
		* `audience` - (Optional) (Updatable) Audience for IDCS.
		* `domain_url` - (Required) (Updatable) Domain URL for IDCS.
		* `scope` - (Required) (Updatable) Scope for IDCS.
	* `inbound_auth_config_type` - (Required) (Updatable) Inbound authentication configuration type of network access (IDCS_AUTH_CONFIG).
* `networking_config` - (Optional) Networking configuration.
	* `inbound_networking_config` - (Required) Inbound Networking configuration.
		* `endpoint_mode` - (Required) inbounding from public or private endpoint.
		* `private_endpoint_id` - (Optional) The [OCID] of Private Endpoint when endpointMode=Private
	* `outbound_networking_config` - (Required) Outbound Networking configuration.
		* `custom_subnet_id` - (Optional) ocid of customer subnet when networkMode=Custom
		* `network_mode` - (Required) outbounding to managed internet or customer network.
		* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
* `scaling_config` - (Optional) (Updatable) The auto scaling configuration for the Hosted Application.  Defines the minimum and maximum number of replicas. When unspecified, the service applies service-defined default scaling values.
	* `max_replica` - (Optional) (Updatable) Maximum number of replicas allowed.
	* `min_replica` - (Optional) (Updatable) Minimum number of replicas to keep running.
	* `scaling_type` - (Required) (Updatable) scaling type for application.
	* `target_concurrency_threshold` - (Optional) (Updatable) number of simultaneous requests that can be processed by each replica.
	* `target_cpu_threshold` - (Optional) (Updatable) Scale up if average CPU utilization exceeds this threshold.
	* `target_memory_threshold` - (Optional) (Updatable) Scale up if average memory utilization exceeds this threshold.
	* `target_rps_threshold` - (Optional) (Updatable) requests-per-second per replica of an application.
* `storage_configs` - (Optional) The list of storage configuration for the Hosted Application.  Defines a list of service-managed storage back-ends.
	* `environment_variable_key` - (Required) The key of environment variable to store the database connection.
	* `storage_id` - (Required) The [OCID] of ApplicationStorage.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Hosted Application
	* `update` - (Defaults to 20 minutes), when updating the Hosted Application
	* `delete` - (Defaults to 20 minutes), when destroying the Hosted Application


## Import

HostedApplications can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_hosted_application.test_hosted_application "id"
```
