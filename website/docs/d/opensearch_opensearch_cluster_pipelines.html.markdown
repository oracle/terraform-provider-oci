---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_cluster_pipelines"
sidebar_current: "docs-oci-datasource-opensearch-opensearch_cluster_pipelines"
description: |-
  Provides the list of Opensearch Cluster Pipelines in Oracle Cloud Infrastructure Opensearch service
---

# Data Source: oci_opensearch_opensearch_cluster_pipelines
This data source provides the list of Opensearch Cluster Pipelines in Oracle Cloud Infrastructure Opensearch service.

Returns a list of OpensearchClusterPipelines.


## Example Usage

```hcl
data "oci_opensearch_opensearch_cluster_pipelines" "test_opensearch_cluster_pipelines" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.opensearch_cluster_pipeline_display_name
	id = var.opensearch_cluster_pipeline_id
	pipeline_component_id = oci_opensearch_pipeline_component.test_pipeline_component.id
	state = var.opensearch_cluster_pipeline_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique OpensearchClusterPipeline identifier
* `pipeline_component_id` - (Optional) A filter to return pipelines whose any component has the given pipelineComponentId.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `opensearch_cluster_pipeline_collection` - The list of opensearch_cluster_pipeline_collection.

### OpensearchClusterPipeline Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where the pipeline is located.
* `data_prepper_configuration_body` - The data prepper config in YAML format. The command accepts the data prepper config as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the pipeline. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the cluster pipeline.
* `memory_gb` - The amount of memory in GB, for each pipeline node.
* `node_count` - The number of nodes configured for the pipeline.
* `node_shape` - The pipeline node shape.
* `nsg_id` - The OCID of the NSG where the pipeline private endpoint vnic will be attached.
* `ocpu_count` - The number of OCPUs configured for each pipeline node.
* `opensearch_pipeline_fqdn` - The fully qualified domain name (FQDN) for the cluster's API endpoint.
* `opensearch_pipeline_private_ip` - The pipeline's private IP address.
* `pipeline_configuration_body` - The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \. 
* `pipeline_mode` - The current state of the pipeline.
* `reverse_connection_endpoints` - The customer IP and the corresponding fully qualified domain name that the pipeline will connect to.
	* `customer_fqdn` - The fully qualified domain name of the customerIp in the customer VCN
	* `customer_ip` - The IPv4 address in the customer VCN
* `state` - The current state of the cluster backup.
* `subnet_compartment_id` - The OCID for the compartment where the pipeline's subnet is located.
* `subnet_id` - The OCID of the pipeline's subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cluster pipeline was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The amount of time in milliseconds since the pipeline was updated.
* `vcn_compartment_id` - The OCID for the compartment where the pipeline's VCN is located.
* `vcn_id` - The OCID of the pipeline's VCN.

