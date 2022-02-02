---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster"
sidebar_current: "docs-oci-resource-containerengine-cluster"
description: |-
  Provides the Cluster resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster
This resource provides the Cluster resource in Oracle Cloud Infrastructure Container Engine service.

Create a new cluster.

## Example Usage

```hcl
resource "oci_containerengine_cluster" "test_cluster" {
	#Required
	compartment_id = var.compartment_id
	kubernetes_version = var.cluster_kubernetes_version
	name = var.cluster_name
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	endpoint_config {

		#Optional
		is_public_ip_enabled = var.cluster_endpoint_config_is_public_ip_enabled
		nsg_ids = var.cluster_endpoint_config_nsg_ids
		subnet_id = oci_core_subnet.test_subnet.id
	}
	freeform_tags = {"Department"= "Finance"}
	image_policy_config {

		#Optional
		is_policy_enabled = var.cluster_image_policy_config_is_policy_enabled
		key_details {

			#Optional
			kms_key_id = oci_kms_key.test_key.id
		}
	}
	kms_key_id = oci_kms_key.test_key.id
	options {

		#Optional
		add_ons {

			#Optional
			is_kubernetes_dashboard_enabled = var.cluster_options_add_ons_is_kubernetes_dashboard_enabled
			is_tiller_enabled = var.cluster_options_add_ons_is_tiller_enabled
		}
		admission_controller_options {

			#Optional
			is_pod_security_policy_enabled = var.cluster_options_admission_controller_options_is_pod_security_policy_enabled
		}
		kubernetes_network_config {

			#Optional
			pods_cidr = var.cluster_options_kubernetes_network_config_pods_cidr
			services_cidr = var.cluster_options_kubernetes_network_config_services_cidr
		}
		persistent_volume_config {

			#Optional
			defined_tags = {"Operations.CostCenter"= "42"}
			freeform_tags = {"Department"= "Finance"}
		}
		service_lb_config {

			#Optional
			defined_tags = {"Operations.CostCenter"= "42"}
			freeform_tags = {"Department"= "Finance"}
		}
		service_lb_subnet_ids = var.cluster_options_service_lb_subnet_ids
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to create the cluster.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `endpoint_config` - (Optional) The network configuration for access to the Cluster control plane. 
	* `is_public_ip_enabled` - (Optional) Whether the cluster should be assigned a public IP address. Defaults to false. If set to true on a private subnet, the cluster provisioning will fail.
	* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to apply to the cluster endpoint. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `subnet_id` - (Optional) The OCID of the regional subnet in which to place the Cluster endpoint.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `image_policy_config` - (Optional) (Updatable) The image verification policy for signature validation. Once a policy is created and enabled with one or more kms keys, the policy will ensure all images deployed has been signed with the key(s) attached to the policy. 
	* `is_policy_enabled` - (Optional) (Updatable) Whether the image verification policy is enabled. Defaults to false. If set to true, the images will be verified against the policy at runtime.
	* `key_details` - (Optional) (Updatable) A list of KMS key details.
		* `kms_key_id` - (Optional) (Updatable) The OCIDs of the KMS key that will be used to verify whether the images are signed by an approved source. 
* `kms_key_id` - (Optional) The OCID of the KMS key to be used as the master encryption key for Kubernetes secret encryption. When used, `kubernetesVersion` must be at least `v1.13.0`. 
* `kubernetes_version` - (Required) (Updatable) The version of Kubernetes to install into the cluster masters.
* `name` - (Required) (Updatable) The name of the cluster. Avoid entering confidential information.
* `options` - (Optional) (Updatable) Optional attributes for the cluster.
	* `add_ons` - (Optional) Configurable cluster add-ons
		* `is_kubernetes_dashboard_enabled` - (Optional) Whether or not to enable the Kubernetes Dashboard add-on.
		* `is_tiller_enabled` - (Optional) Whether or not to enable the Tiller add-on.
	* `admission_controller_options` - (Optional) (Updatable) Configurable cluster admission controllers
		* `is_pod_security_policy_enabled` - (Optional) (Updatable) Whether or not to enable the Pod Security Policy admission controller.
	* `kubernetes_network_config` - (Optional) Network configuration for Kubernetes.
		* `pods_cidr` - (Optional) The CIDR block for Kubernetes pods. Optional, defaults to 10.244.0.0/16.
		* `services_cidr` - (Optional) The CIDR block for Kubernetes services. Optional, defaults to 10.96.0.0/16.
	* `persistent_volume_config` - (Optional) (Updatable) Configuration to be applied to block volumes created by Kubernetes Persistent Volume Claims (PVC)
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_config` - (Optional) (Updatable) Configuration to be applied to load balancers created by Kubernetes services
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_subnet_ids` - (Optional) The OCIDs of the subnets used for Kubernetes services load balancers.
* `vcn_id` - (Required) The OCID of the virtual cloud network (VCN) in which to create the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `available_kubernetes_upgrades` - Available Kubernetes versions to which the clusters masters may be upgraded.
* `compartment_id` - The OCID of the compartment in which the cluster exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `endpoint_config` - The network configuration for access to the Cluster control plane. 
	* `is_public_ip_enabled` - Whether the cluster should be assigned a public IP address. Defaults to false. If set to true on a private subnet, the cluster provisioning will fail.
	* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to apply to the cluster endpoint. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `subnet_id` - The OCID of the regional subnet in which to place the Cluster endpoint.
* `endpoints` - Endpoints served up by the cluster masters.
	* `kubernetes` - The non-native networking Kubernetes API server endpoint.
	* `private_endpoint` - The private native networking Kubernetes API server endpoint.
	* `public_endpoint` - The public native networking Kubernetes API server endpoint, if one was requested.
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
	* `persistent_volume_config` - Configuration to be applied to block volumes created by Kubernetes Persistent Volume Claims (PVC)
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_config` - Configuration to be applied to load balancers created by Kubernetes services
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_subnet_ids` - The OCIDs of the subnets used for Kubernetes services load balancers.
* `state` - The state of the cluster masters.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `vcn_id` - The OCID of the virtual cloud network (VCN) in which the cluster exists.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Cluster
	* `update` - (Defaults to 1 hours), when updating the Cluster
	* `delete` - (Defaults to 1 hours), when destroying the Cluster


## Import

Clusters can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster.test_cluster "id"
```

