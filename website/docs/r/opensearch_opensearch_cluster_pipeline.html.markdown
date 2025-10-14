---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_cluster_pipeline"
sidebar_current: "docs-oci-resource-opensearch-opensearch_cluster_pipeline"
description: |-
  Provides the Opensearch Cluster Pipeline resource in Oracle Cloud Infrastructure Opensearch service
---

# oci_opensearch_opensearch_cluster_pipeline
This resource provides the Opensearch Cluster Pipeline resource in Oracle Cloud Infrastructure Opensearch service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/opensearch/latest/OpensearchClusterPipeline

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/opensearch

Creates a new OpensearchCluster Pipeline.


## Example Usage

```hcl
resource "oci_opensearch_opensearch_cluster_pipeline" "test_opensearch_cluster_pipeline" {
	#Required
	compartment_id = var.compartment_id
	data_prepper_configuration_body = var.opensearch_cluster_pipeline_data_prepper_configuration_body
	display_name = var.opensearch_cluster_pipeline_display_name
	memory_gb = var.opensearch_cluster_pipeline_memory_gb
	node_count = var.opensearch_cluster_pipeline_node_count
	ocpu_count = var.opensearch_cluster_pipeline_ocpu_count
	pipeline_configuration_body = var.opensearch_cluster_pipeline_pipeline_configuration_body

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	node_shape = var.opensearch_cluster_pipeline_node_shape
	nsg_id = oci_opensearch_nsg.test_nsg.id
	opc_dry_run = var.opensearch_cluster_pipeline_opc_dry_run
	reverse_connection_endpoints {
		#Required
		customer_fqdn = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_fqdn
		customer_ip = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_ip
	}
	subnet_compartment_id = oci_identity_compartment.test_compartment.id
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_compartment_id = oci_identity_compartment.test_compartment.id
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to create the pipeline in.
* `data_prepper_configuration_body` - (Required) (Updatable) The data prepper config in YAML format. The command accepts the data prepper config as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The name of the cluster pipeline. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `memory_gb` - (Required) (Updatable) The amount of memory in GB, for each pipeline node.
* `node_count` - (Required) (Updatable) The number of nodes configured for the pipeline.
* `node_shape` - (Optional) (Updatable) The pipeline node shape.
* `nsg_id` - (Optional) (Updatable) The OCID of the NSG where the pipeline private endpoint vnic will be attached.
* `ocpu_count` - (Required) (Updatable) The number of OCPUs configured for each pipeline node.
* `opc_dry_run` - (Optional) (Updatable) Indicates that the request is a dry run, if set to "true". A dry run request does not modify the configuration item details and is used only to perform validation on the submitted data. 
* `pipeline_configuration_body` - (Required) (Updatable) The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \. 
* `reverse_connection_endpoints` - (Optional) (Updatable) The customer IP and the corresponding fully qualified domain name that the pipeline will connect to.
	* `customer_fqdn` - (Required) (Updatable) The fully qualified domain name of the customerIp in the customer VCN
	* `customer_ip` - (Required) (Updatable) The IPv4 address in the customer VCN
* `subnet_compartment_id` - (Optional) (Updatable) The OCID for the compartment where the pipeline's subnet is located.
* `subnet_id` - (Optional) (Updatable) The OCID of the pipeline's subnet.
* `vcn_compartment_id` - (Optional) (Updatable) The OCID for the compartment where the pipeline's VCN is located.
* `vcn_id` - (Optional) (Updatable) The OCID of the pipeline's VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Opensearch Cluster Pipeline
	* `update` - (Defaults to 20 minutes), when updating the Opensearch Cluster Pipeline
	* `delete` - (Defaults to 20 minutes), when destroying the Opensearch Cluster Pipeline


## Import

OpensearchClusterPipelines can be imported using the `id`, e.g.

```
$ terraform import oci_opensearch_opensearch_cluster_pipeline.test_opensearch_cluster_pipeline "id"
```

