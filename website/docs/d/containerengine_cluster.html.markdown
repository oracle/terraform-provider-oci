---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster"
sidebar_current: "docs-oci-datasource-containerengine-cluster"
description: |-
  Provides details about a specific Cluster in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster
This data source provides details about a specific Cluster resource in Oracle Cloud Infrastructure Container Engine service.

Get the details of a cluster.

## Example Usage

```hcl
data "oci_containerengine_cluster" "test_cluster" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `available_kubernetes_upgrades` - Available Kubernetes versions to which the clusters masters may be upgraded.
* `cluster_pod_network_options` - Available CNIs and network options for existing and new node pools of the cluster
	* `cni_type` - The CNI used by the node pools of this cluster
* `compartment_id` - The OCID of the compartment in which the cluster exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `endpoint_config` - The network configuration for access to the Cluster control plane. 
	* `is_public_ip_enabled` - Whether the cluster should be assigned a public IP address. Defaults to false. If set to true on a private subnet, the cluster provisioning will fail.
	* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to apply to the cluster endpoint. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `subnet_id` - The OCID of the regional subnet in which to place the Cluster endpoint.
* `endpoints` - Endpoints served up by the cluster masters.
	* `ipv6endpoint` - The IPv6 networking Kubernetes API server endpoint.
	* `kubernetes` - The non-native networking Kubernetes API server endpoint.
	* `private_endpoint` - The private native networking Kubernetes API server endpoint.
	* `public_endpoint` - The public native networking Kubernetes API server endpoint, if one was requested.
	* `vcn_hostname_endpoint` - The FQDN assigned to the Kubernetes API private endpoint. Example: 'https://yourVcnHostnameEndpoint' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the cluster.
* `image_policy_config` - The image verification policy for signature validation. 
	* `is_policy_enabled` - Whether the image verification policy is enabled. Defaults to false. If set to true, the images will be verified against the policy at runtime.
	* `key_details` - A list of KMS key details.
		* `kms_key_id` - The OCIDs of the KMS key that will be used to verify whether the images are signed by an approved source. 
* `kms_key_id` - The OCID of the KMS key to be used as the master encryption key for Kubernetes secret encryption.
* `kubernetes_version` - The version of Kubernetes running on the cluster masters.
* `lifecycle_details` - Details about the state of the cluster masters.
* `metadata` - Metadata about the cluster.
	* `created_by_user_id` - The user who created the cluster.
	* `created_by_work_request_id` - The OCID of the work request which created the cluster.
	* `deleted_by_user_id` - The user who deleted the cluster.
	* `deleted_by_work_request_id` - The OCID of the work request which deleted the cluster.
	* `time_created` - The time the cluster was created.
	* `time_credential_expiration` - The time until which the cluster credential is valid.
	* `time_deleted` - The time the cluster was deleted.
	* `time_updated` - The time the cluster was updated.
	* `updated_by_user_id` - The user who updated the cluster.
	* `updated_by_work_request_id` - The OCID of the work request which updated the cluster.
* `name` - The name of the cluster.
* `options` - Optional attributes for the cluster.
	* `add_ons` - Configurable cluster add-ons
		* `is_kubernetes_dashboard_enabled` - Whether or not to enable the Kubernetes Dashboard add-on.
		* `is_tiller_enabled` - Whether or not to enable the Tiller add-on.
	* `admission_controller_options` - Configurable cluster admission controllers
		* `is_pod_security_policy_enabled` - Whether or not to enable the Pod Security Policy admission controller.
	* `ip_families` - IP family to use for single stack or define the order of IP families for dual-stack
	* `kubernetes_network_config` - Network configuration for Kubernetes.
		* `pods_cidr` - The CIDR block for Kubernetes pods. Optional, defaults to 10.244.0.0/16.
		* `services_cidr` - The CIDR block for Kubernetes services. Optional, defaults to 10.96.0.0/16.
	* `persistent_volume_config` - Configuration to be applied to block volumes created by Kubernetes Persistent Volume Claims (PVC)
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_config` - Configuration to be applied to load balancers created by Kubernetes services
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_subnet_ids` - The OCIDs of the subnets used for Kubernetes services load balancers.
* `state` - The state of the cluster masters.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `type` - Type of cluster. Values can be BASIC_CLUSTER or ENHANCED_CLUSTER. For more information, see [Cluster Types](https://docs.cloud.oracle.com/iaas/Content/ContEng/Tasks/contengcomparingenhancedwithbasicclusters_topic.htm)
* `vcn_id` - The OCID of the virtual cloud network (VCN) in which the cluster exists.

