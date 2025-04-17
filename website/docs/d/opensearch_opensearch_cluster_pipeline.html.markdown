---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_cluster_pipeline"
sidebar_current: "docs-oci-datasource-opensearch-opensearch_cluster_pipeline"
description: |-
  Provides details about a specific Opensearch Cluster Pipeline in Oracle Cloud Infrastructure Opensearch service
---

# Data Source: oci_opensearch_opensearch_cluster_pipeline
This data source provides details about a specific Opensearch Cluster Pipeline resource in Oracle Cloud Infrastructure Opensearch service.

Gets a OpensearchCluster Pipeline by identifier

## Example Usage

```hcl
data "oci_opensearch_opensearch_cluster_pipeline" "test_opensearch_cluster_pipeline" {
	#Required
	opensearch_cluster_pipeline_id = oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline.id
}
```

## Argument Reference

The following arguments are supported:

* `opensearch_cluster_pipeline_id` - (Required) unique OpensearchClusterPipeline identifier


## Attributes Reference

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

