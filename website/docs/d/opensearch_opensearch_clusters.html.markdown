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
* `inbound_cluster_ids` - List of inbound clusters for which this cluster is an outbound cluster
* `lifecycle_details` - Additional information about the current lifecycle state of the cluster.
* `maintenance_details` - Details for the maintenance activity.
	* `end_time` - End time of the maintenance activity
	* `notification_email_ids` - The Email Ids given the by customer to get notified about maintenance activities
	* `start_time` - Start time of the maintenance activity
	* `state` - State of the maintenance activity
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
* `outbound_cluster_config` - This configuration is used for passing request details to connect outbound cluster(s) to the inbound cluster (coordinating cluster) 
	* `is_enabled` - Flag to indicate whether outbound cluster configuration is enabled
	* `outbound_clusters` - List of outbound clusters to be connected to the inbound cluster
		* `display_name` - Name of the Outbound cluster. Avoid entering confidential information.
		* `is_skip_unavailable` - Flag to indicate whether to skip the Outbound cluster during cross cluster search, if it is unavailable
		* `mode` - Mode for the cross cluster connection
		* `ping_schedule` - Sets the time interval between regular application-level ping messages that are sent to try and keep outbound cluster connections alive. If set to -1, application-level ping messages to this outbound cluster are not sent. If unset, application-level ping messages are sent according to the global transport.ping_schedule setting, which defaults to -1 meaning that pings are not sent.
		* `seed_cluster_id` - OCID of the Outbound cluster
* `reverse_connection_endpoint_customer_ips` - The customer IP addresses of the endpoint in customer VCN
* `reverse_connection_endpoints` - The list of reverse connection endpoints.
	* `customer_ip` - The IP addresses of the endpoint in customer VCN
	* `nat_ip` - The NAT IP addresses of the endpoint in service VCN
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
