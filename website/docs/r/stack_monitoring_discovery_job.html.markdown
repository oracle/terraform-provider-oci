---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_discovery_job"
sidebar_current: "docs-oci-resource-stack_monitoring-discovery_job"
description: |-
  Provides the Discovery Job resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_discovery_job
This resource provides the Discovery Job resource in Oracle Cloud Infrastructure Stack Monitoring service.

API to create discovery Job and submit discovery Details to agent.


## Example Usage

```hcl
resource "oci_stack_monitoring_discovery_job" "test_discovery_job" {
	#Required
	compartment_id = var.compartment_id
	discovery_details {
		#Required
		agent_id = var.management_agent_id
		properties {

			#Optional
			properties_map = var.discovery_job_discovery_details_properties_properties_map
		}
		resource_name = var.discovery_job_discovery_details_resource_name
		resource_type = var.discovery_job_discovery_details_resource_type

		#Optional
		credentials {
			#Required
			items {
				#Required
				credential_name = var.discovery_job_discovery_details_credentials_items_credential_name
				credential_type = var.discovery_job_discovery_details_credentials_items_credential_type
				properties {

					#Optional
					properties_map = var.discovery_job_discovery_details_credentials_items_properties_properties_map
				}
			}
		}
		license = var.discovery_job_discovery_details_license
		tags {

			#Optional
			properties_map = var.discovery_job_discovery_details_tags_properties_map
		}
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	discovery_client = var.discovery_job_discovery_client
	discovery_type = var.discovery_job_discovery_type
	freeform_tags = {"bar-key"= "value"}
	should_propagate_tags_to_discovered_resources = var.discovery_job_should_propagate_tags_to_discovered_resources
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of Compartment
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `discovery_client` - (Optional) Client who submits discovery job.
* `discovery_details` - (Required) The request of DiscoveryJob Resource details.
	* `agent_id` - (Required) The OCID of Management Agent
	* `credentials` - (Optional) List of DiscoveryJob Credential Details.
		* `items` - (Required) List of DiscoveryJob credentials.
			* `credential_name` - (Required) Name of Credential
			* `credential_type` - (Required) Name of Credential Type
			* `properties` - (Required) Property Details
				* `properties_map` - (Optional) Key/Value pair of Property
	* `license` - (Optional) License edition of the monitored resource.
	* `properties` - (Required) Property Details
		* `properties_map` - (Optional) Key/Value pair of Property
	* `resource_name` - (Required) The Name of resource type
	* `resource_type` - (Required) Resource Type.
	* `tags` - (Optional) Property Details
		* `properties_map` - (Optional) Key/Value pair of Property
* `discovery_type` - (Optional) Add option submits new discovery Job. Add with retry option to re-submit failed discovery job. Refresh option refreshes the existing discovered resources. 
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `should_propagate_tags_to_discovered_resources` - (Optional) If this parameter set to true, the specified tags will be applied  to all resources discovered in the current request.  Default is true. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the Compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `discovery_client` - Client who submits discovery job.
* `discovery_details` - The request of DiscoveryJob Resource details.
	* `agent_id` - The OCID of Management Agent
	* `credentials` - List of DiscoveryJOb Credential Details.
		* `items` - List of DiscoveryJob credentials.
			* `credential_name` - Name of Credential
			* `credential_type` - Name of Credential Type
			* `properties` - Property Details
				* `properties_map` - Key/Value pair of Property
	* `license` - License edition of the monitored resource.
	* `properties` - Property Details
		* `properties_map` - Key/Value pair of Property
	* `resource_name` - The Name of resource type
	* `resource_type` - Resource Type.
	* `tags` - Property Details
		* `properties_map` - Key/Value pair of Property
* `discovery_type` - Add option submits new discovery Job. Add with retry option to re-submit failed discovery job. Refresh option refreshes the existing discovered resources. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of Discovery job
* `state` - The current state of the DiscoveryJob Resource.
* `status` - Specifies the status of the discovery job
* `status_message` - The short summary of the status of the discovery job
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - The OCID of Tenant
* `time_updated` - The time the discovery Job was updated.
* `user_id` - The OCID of user in which the job is submitted

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Discovery Job
	* `update` - (Defaults to 20 minutes), when updating the Discovery Job
	* `delete` - (Defaults to 20 minutes), when destroying the Discovery Job


## Import

DiscoveryJobs can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_discovery_job.test_discovery_job "id"
```

