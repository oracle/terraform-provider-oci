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
	inbound_cluster_ids = var.opensearch_cluster_inbound_cluster_ids
	maintenance_details {

		#Optional
		notification_email_ids = var.opensearch_cluster_maintenance_details_notification_email_ids
	}
	master_node_host_bare_metal_shape = var.opensearch_cluster_master_node_host_bare_metal_shape
	outbound_cluster_config {
		#Required
		is_enabled = var.opensearch_cluster_outbound_cluster_config_is_enabled
		outbound_clusters {
			#Required
			display_name = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_display_name
			seed_cluster_id = oci_opensearch_opensearch_cluster.test_opensearch_cluster.id

			#Optional
			is_skip_unavailable = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_is_skip_unavailable
			mode = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_mode
			ping_schedule = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_ping_schedule
		}
	}
	reverse_connection_endpoint_customer_ips = var.opensearch_cluster_reverse_connection_endpoint_customer_ips
	security_master_user_name = oci_identity_user.test_user.name
	security_master_user_password_hash = var.opensearch_cluster_security_master_user_password_hash
	security_mode = var.opensearch_cluster_security_mode
	system_tags = var.opensearch_cluster_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to create the cluster in.
* `data_node_count` - (Required) (Updatable) The number of data nodes to configure for the cluster.
* `data_node_host_bare_metal_shape` - (Optional) The bare metal shape for the cluster's data nodes.
* `data_node_host_memory_gb` - (Required) (Updatable) The amount of memory in GB, to configure per node for the cluster's data nodes.
* `data_node_host_ocpu_count` - (Required) (Updatable) The number of OCPUs to configure for the cluster's data nodes.
* `data_node_host_type` - (Required) TThe instance type for the cluster's data nodes.
* `data_node_storage_gb` - (Required) (Updatable) The amount of storage in GB, to configure per node for the cluster's data nodes.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - (Required) (Updatable) The name of the cluster. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `inbound_cluster_ids` - (Optional) List of inbound clusters that will be queried using cross cluster search
* `maintenance_details` - (Optional) (Updatable) Details for creation of maintenance details
	* `notification_email_ids` - (Optional) (Updatable) The Email IDs given by the customer to get notified about maintenance activities
* `master_node_count` - (Required) (Updatable) The number of master nodes to configure for the cluster.
* `master_node_host_bare_metal_shape` - (Optional) The bare metal shape for the cluster's master nodes.
* `master_node_host_memory_gb` - (Required) (Updatable) The amount of memory in GB, to configure per node for the cluster's master nodes.
* `master_node_host_ocpu_count` - (Required) (Updatable) The number of OCPUs to configure for the cluser's master nodes.
* `master_node_host_type` - (Required) The instance type for the cluster's master nodes.
* `opendashboard_node_count` - (Required) (Updatable) The number of OpenSearch Dashboard nodes to configure for the cluster.
* `opendashboard_node_host_memory_gb` - (Required) (Updatable) The amount of memory in GB, to configure for the cluster's OpenSearch Dashboard nodes.
* `opendashboard_node_host_ocpu_count` - (Required) (Updatable) The number of OCPUs to configure for the cluster's OpenSearch Dashboard nodes.
* `outbound_cluster_config` - (Optional) (Updatable) This configuration is used for passing request details to connect outbound cluster(s) to the inbound cluster (coordinating cluster) 
	* `is_enabled` - (Required) (Updatable) Flag to indicate whether outbound cluster configuration is enabled
	* `outbound_clusters` - (Required) (Updatable) List of outbound clusters to be connected to the inbound cluster
		* `display_name` - (Required) (Updatable) Name of the Outbound cluster. Avoid entering confidential information.
		* `is_skip_unavailable` - (Optional) (Updatable) Flag to indicate whether to skip the Outbound cluster during cross cluster search, if it is unavailable
		* `mode` - (Optional) (Updatable) Mode for the cross cluster connection
		* `ping_schedule` - (Optional) (Updatable) Sets the time interval between regular application-level ping messages that are sent to try and keep outbound cluster connections alive. If set to -1, application-level ping messages to this outbound cluster are not sent. If unset, application-level ping messages are sent according to the global transport.ping_schedule setting, which defaults to -1 meaning that pings are not sent.
		* `seed_cluster_id` - (Required) (Updatable) OCID of the Outbound cluster
* `reverse_connection_endpoint_customer_ips` - (Optional) (Updatable) The customer IP addresses of the endpoint in customer VCN
* `security_master_user_name` - (Optional) (Updatable) The name of the master user that are used to manage security config
* `security_master_user_password_hash` - (Optional) (Updatable) The password hash of the master user that are used to manage security config
* `security_mode` - (Optional) (Updatable) The security mode of the cluster.
* `software_version` - (Required) (Updatable) The version of the software the cluster is running.
* `subnet_compartment_id` - (Required) The OCID for the compartment where the cluster's subnet is located.
* `subnet_id` - (Required) The OCID of the cluster's subnet.
* `system_tags` - (Optional) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `vcn_compartment_id` - (Required) The OCID for the compartment where the cluster's VCN is located.
* `vcn_id` - (Required) The OCID of the cluster's VCN.
* `configure_outbound_cluster_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure Outbound Cluster. Could be set to any integer value.


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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 45 minutes), when creating the Opensearch Cluster
* `update` - (Defaults to 45 minutes), when updating the Opensearch Cluster
* `delete` - (Defaults to 45 minutes), when destroying the Opensearch Cluster


## Import

OpensearchClusters can be imported using the `id`, e.g.

```
$ terraform import oci_opensearch_opensearch_cluster.test_opensearch_cluster "id"
```
