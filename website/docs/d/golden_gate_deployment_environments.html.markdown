---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_environments"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_environments"
description: |-
  Provides the list of Deployment Environments in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_environments
This data source provides the list of Deployment Environments in Oracle Cloud Infrastructure Golden Gate service.

Returns an array of DeploymentEnvironmentDescriptor


## Example Usage

```hcl
data "oci_golden_gate_deployment_environments" "test_deployment_environments" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 


## Attributes Reference

The following attributes are exported:

* `deployment_environment_collection` - The list of deployment_environment_collection.

### DeploymentEnvironment Reference

The following attributes are exported:

* `items` - Array of DeploymentEnvironmentSummary objects. 
	* `category` - The deployment category defines the broad separation of the deployment type into three categories. Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS' and 'DATA_TRANSFORMS'. 
	* `default_cpu_core_count` - The default CPU core count. 
	* `display_name` - An object's Display Name. 
	* `environment_type` - Specifies whether the deployment is used in a production or development/testing environment. 
	* `is_auto_scaling_enabled_by_default` - Specifies whether the "Auto scaling" option should be enabled by default or not. 
	* `max_cpu_core_count` - The maximum CPU core count. 
	* `memory_per_ocpu_in_gbs` - The multiplier value between CPU core count and memory size. 
	* `min_cpu_core_count` - The minimum CPU core count. 
	* `network_bandwidth_per_ocpu_in_gbps` - The multiplier value between CPU core count and network bandwidth. 
	* `storage_usage_limit_per_ocpu_in_gbs` - The multiplier value between CPU core count and storage usage limit size. 

