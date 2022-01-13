---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_clusters"
sidebar_current: "docs-oci-datasource-containerengine-clusters"
description: |-
  Provides the list of Clusters in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_clusters
This data source provides the list of Clusters in Oracle Cloud Infrastructure Container Engine service.

List all the cluster objects in a compartment.

## Example Usage

```hcl
data "oci_containerengine_clusters" "test_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.cluster_name
	state = var.cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) The name to filter on.
* `state` - (Optional) A cluster lifecycle state to filter on. Can have multiple parameters of this name.


## Attributes Reference

The following attributes are exported:

* `clusters` - The list of clusters.

### Cluster Reference

The following attributes are exported:

* `available_kubernetes_upgrades` - Available Kubernetes versions to which the clusters masters may be upgraded.
* `compartment_id` - The OCID of the compartment in which the cluster exists.
* `endpoint_config` - The network configuration for access to the Cluster control plane. 
	* `is_public_ip_enabled` - Whether the cluster should be assigned a public IP address. Defaults to false. If set to true on a private subnet, the cluster provisioning will fail.
	* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to apply to the cluster endpoint. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `subnet_id` - The OCID of the regional subnet in which to place the Cluster endpoint.
* `endpoints` - Endpoints served up by the cluster masters.
	* `kubernetes` - The non-native networking Kubernetes API server endpoint.
	* `private_endpoint` - The private native networking Kubernetes API server endpoint.
	* `public_endpoint` - The public native networking Kubernetes API server endpoint, if one was requested.
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
	* `kubernetes_network_config` - Network configuration for Kubernetes.
		* `pods_cidr` - The CIDR block for Kubernetes pods. Optional, defaults to 10.244.0.0/16.
		* `services_cidr` - The CIDR block for Kubernetes services. Optional, defaults to 10.96.0.0/16.
	* `service_lb_subnet_ids` - The OCIDs of the subnets used for Kubernetes services load balancers.
* `state` - The state of the cluster masters.
* `vcn_id` - The OCID of the virtual cloud network (VCN) in which the cluster exists.

