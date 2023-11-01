---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_discovery_jobs"
sidebar_current: "docs-oci-datasource-stack_monitoring-discovery_jobs"
description: |-
  Provides the list of Discovery Jobs in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_discovery_jobs
This data source provides the list of Discovery Jobs in Oracle Cloud Infrastructure Stack Monitoring service.

API to get the details of all Discovery Jobs.


## Example Usage

```hcl
data "oci_stack_monitoring_discovery_jobs" "test_discovery_jobs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.discovery_job_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which data is listed.
* `name` - (Optional) A filter to return only discovery jobs that match the entire resource name given.


## Attributes Reference

The following attributes are exported:

* `discovery_job_collection` - The list of discovery_job_collection.

### DiscoveryJob Reference

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

