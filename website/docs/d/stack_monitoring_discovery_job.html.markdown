---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_discovery_job"
sidebar_current: "docs-oci-datasource-stack_monitoring-discovery_job"
description: |-
  Provides details about a specific Discovery Job in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_discovery_job
This data source provides details about a specific Discovery Job resource in Oracle Cloud Infrastructure Stack Monitoring service.

API to get the details of discovery Job by identifier.


## Example Usage

```hcl
data "oci_stack_monitoring_discovery_job" "test_discovery_job" {
	#Required
	discovery_job_id = oci_stack_monitoring_discovery_job.test_discovery_job.id
}
```

## Argument Reference

The following arguments are supported:

* `discovery_job_id` - (Required) The Discovery Job ID


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

