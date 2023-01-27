---
subcategory: "Opensearch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opensearch_opensearch_clusters"
sidebar_current: "docs-oci-datasource-opensearch-opensearch_clusters"
description: |-
  Provides the list of Opensearch Clusters in Oracle Cloud Infrastructure Opensearch service
---

# Data Source: oci_opensearch_opensearch_clusters
This data source provides the list of Opensearch Clusters in Oracle Cloud Infrastructure Opensearch service.

Returns a list of OpensearchClusters.

## Prerequisites
The below policies must be created in compartment before creating OpensearchCluster

##### {Compartment-Name} - Name of  your compartment
```
Allow service opensearch to manage vnics in compartment {Compartment-Name}
Allow service opensearch to use subnets in compartment {Compartment-Name}
Allow service opensearch to use network-security-groups in compartment {Compartment-Name}
Allow service opensearch to manage vcns in compartment {Compartment-Name}
```

For latest documentation on OpenSearch use please refer to https://docs.oracle.com/en-us/iaas/Content/search-opensearch/home.htm  
Required permissions: https://docs.oracle.com/en-us/iaas/Content/search-opensearch/Concepts/ocisearchpermissions.htm

## Example Usage

```hcl
data "oci_opensearch_opensearch_clusters" "test_opensearch_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.opensearch_cluster_display_name
	id = var.opensearch_cluster_id
	state = var.opensearch_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique OpensearchCluster identifier
* `state` - (Optional) A filter to return only OpensearchClusters their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `opensearch_cluster_collection` - The list of opensearch_cluster_collection.

### OpensearchCluster Reference

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
* `security_master_user_name` - The name of the master user that are used to manage security config
* `security_master_user_password_hash` - The password hash of the master user that are used to manage security config
* `security_mode` - The security mode of the cluster.
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
