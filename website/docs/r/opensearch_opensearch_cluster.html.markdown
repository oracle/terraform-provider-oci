---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_cluster"
sidebar_current: "docs-oci-resource-opensearch-opensearch_cluster"
description: |-
Provides the Opensearch Cluster resource in Oracle Cloud Infrastructure Opensearch service
---

# oci_opensearch_opensearch_cluster
This resource provides the Opensearch Cluster resource in Oracle Cloud Infrastructure Opensearch service.

Creates a new OpensearchCluster.


## Example Usage

```hcl
resource "oci_opensearch_opensearch_cluster" "test_opensearch_cluster" {
	#Required
	compartment_id = var.compartment_id
	data_node_count = var.opensearch_cluster_data_node_count
	data_node_host_memory_gb = var.opensearch_cluster_data_node_host_memory_gb
	data_node_host_ocpu_count = var.opensearch_cluster_data_node_host_ocpu_count
	data_node_host_type = var.opensearch_cluster_data_node_host_type
	data_node_storage_gb = var.opensearch_cluster_data_node_storage_gb
	display_name = var.opensearch_cluster_display_name
	master_node_count = var.opensearch_cluster_master_node_count
	master_node_host_memory_gb = var.opensearch_cluster_master_node_host_memory_gb
	master_node_host_ocpu_count = var.opensearch_cluster_master_node_host_ocpu_count
	master_node_host_type = var.opensearch_cluster_master_node_host_type
	opendashboard_node_count = var.opensearch_cluster_opendashboard_node_count
	opendashboard_node_host_memory_gb = var.opensearch_cluster_opendashboard_node_host_memory_gb
	opendashboard_node_host_ocpu_count = var.opensearch_cluster_opendashboard_node_host_ocpu_count
	software_version = var.opensearch_cluster_software_version
	subnet_compartment_id = oci_identity_compartment.test_compartment.id
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_compartment_id = oci_identity_compartment.test_compartment.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	data_node_host_bare_metal_shape = var.opensearch_cluster_data_node_host_bare_metal_shape
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	master_node_host_bare_metal_shape = var.opensearch_cluster_master_node_host_bare_metal_shape
	system_tags = var.opensearch_cluster_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to create the cluster in.
* `data_node_count` - (Required) The number of data nodes to configure for the cluster.
* `data_node_host_bare_metal_shape` - (Optional) The bare metal shape for the cluster's data nodes.
* `data_node_host_memory_gb` - (Required) The amount of memory in GB, to configure per node for the cluster's data nodes.
* `data_node_host_ocpu_count` - (Required) The number of OCPUs to configure for the cluster's data nodes.
* `data_node_host_type` - (Required) TThe instance type for the cluster's data nodes.
* `data_node_storage_gb` - (Required) The amount of storage in GB, to configure per node for the cluster's data nodes.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - (Required) (Updatable) The name of the cluster. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `master_node_count` - (Required) The number of master nodes to configure for the cluster.
* `master_node_host_bare_metal_shape` - (Optional) The bare metal shape for the cluster's master nodes.
* `master_node_host_memory_gb` - (Required) The amount of memory in GB, to configure per node for the cluster's master nodes.
* `master_node_host_ocpu_count` - (Required) The number of OCPUs to configure for the cluser's master nodes.
* `master_node_host_type` - (Required) The instance type for the cluster's master nodes.
* `opendashboard_node_count` - (Required) The number of OpenSearch Dashboard nodes to configure for the cluster.
* `opendashboard_node_host_memory_gb` - (Required) The amount of memory in GB, to configure for the cluster's OpenSearch Dashboard nodes.
* `opendashboard_node_host_ocpu_count` - (Required) The number of OCPUs to configure for the cluster's OpenSearch Dashboard nodes.
* `software_version` - (Required) (Updatable) The version of the software the cluster is running.
* `subnet_compartment_id` - (Required) The OCID for the compartment where the cluster's subnet is located.
* `subnet_id` - (Required) The OCID of the cluster's subnet.
* `system_tags` - (Optional) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `vcn_compartment_id` - (Required) The OCID for the compartment where the cluster's VCN is located.
* `vcn_id` - (Required) The OCID of the cluster's VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domains` - The availability domains to distribute the cluser nodes across.
* `compartment_id` - The OCID of the compartment where the cluster is located.
* `data_node_count` - The number of data nodes configured for the cluster.
* `data_node_host_bare_metal_shape` - The bare metal shape for the cluster's data nodes.
* `data_node_host_memory_gb` - The amount of memory in GB, for the cluster's data nodes.
* `data_node_host_ocpu_count` - The number of OCPUs configured for the cluster's data nodes.
* `data_node_host_type` - The instance type for the cluster's data nodes.
* `data_node_storage_gb` - The amount of storage in GB, to configure per node for the cluster's data nodes.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - The name of the cluster. Avoid entering confidential information.
* `fqdn` - The fully qualified domain name (FQDN) for the cluster's API endpoint.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - The OCID of the cluster.
* `lifecycle_details` - Additional information about the current lifecycle state of the cluster.
* `master_node_count` - The number of master nodes configured for the cluster.
* `master_node_host_bare_metal_shape` - The bare metal shape for the cluster's master nodes.
* `master_node_host_memory_gb` - The amount of memory in GB, for the cluster's master nodes.
* `master_node_host_ocpu_count` - The number of OCPUs configured for cluster's master nodes.
* `master_node_host_type` - The instance type for the cluster's master nodes.
* `opendashboard_fqdn` - The fully qualified domain name (FQDN) for the cluster's OpenSearch Dashboard API endpoint.
* `opendashboard_node_count` - The number of OpenSearch Dashboard nodes configured for the cluster.
* `opendashboard_node_host_memory_gb` - The amount of memory in GB, for the cluster's OpenSearch Dashboard nodes.
* `opendashboard_node_host_ocpu_count` - The amount of memory in GB, for the cluster's OpenSearch Dashboard nodes.
* `opendashboard_private_ip` - The private IP address for the cluster's OpenSearch Dashboard.
* `opensearch_fqdn` - The fully qualified domain name (FQDN) for the cluster's API endpoint.
* `opensearch_private_ip` - The cluster's private IP address.
* `software_version` - The software version the cluster is running.
* `state` - The current state of the cluster.
* `subnet_compartment_id` - The OCID for the compartment where the cluster's subnet is located.
* `subnet_id` - The OCID of the cluster's subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The amount of time in milliseconds since the cluster was created.
* `time_deleted` - The amount of time in milliseconds since the cluster was updated.
* `time_updated` - The amount of time in milliseconds since the cluster was updated.
* `total_storage_gb` - The size in GB of the cluster's total storage.
* `vcn_compartment_id` - The OCID for the compartment where the cluster's VCN is located.
* `vcn_id` - The OCID of the cluster's VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Opensearch Cluster
* `update` - (Defaults to 20 minutes), when updating the Opensearch Cluster
* `delete` - (Defaults to 20 minutes), when destroying the Opensearch Cluster


## Import

OpensearchClusters can be imported using the `id`, e.g.

```
$ terraform import oci_opensearch_opensearch_cluster.test_opensearch_cluster "id"
```
