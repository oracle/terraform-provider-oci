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
	cluster_pod_network_options {
		#Required
		cni_type = var.cluster_cluster_pod_network_options_cni_type
	}
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
		open_id_connect_token_authentication_config {
			#Required
			is_open_id_connect_auth_enabled = var.cluster_options_open_id_connect_token_authentication_config_is_open_id_connect_auth_enabled

			#Optional
			ca_certificate = var.cluster_options_open_id_connect_token_authentication_config_ca_certificate
			client_id = oci_containerengine_client.test_client.id
			groups_claim = var.cluster_options_open_id_connect_token_authentication_config_groups_claim
			groups_prefix = var.cluster_options_open_id_connect_token_authentication_config_groups_prefix
			issuer_url = var.cluster_options_open_id_connect_token_authentication_config_issuer_url
			required_claims {

				#Optional
				key = var.cluster_options_open_id_connect_token_authentication_config_required_claims_key
				value = var.cluster_options_open_id_connect_token_authentication_config_required_claims_value
			}
			signing_algorithms = var.cluster_options_open_id_connect_token_authentication_config_signing_algorithms
			username_claim = var.cluster_options_open_id_connect_token_authentication_config_username_claim
			username_prefix = var.cluster_options_open_id_connect_token_authentication_config_username_prefix
		}					
		open_id_connect_discovery {

			#Optional
			is_open_id_connect_discovery_enabled = var.cluster_options_open_id_connect_discovery_is_open_id_connect_discovery_enabled
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
	type = var.cluster_type
}
```

## Argument Reference

The following arguments are supported:

* `cluster_pod_network_options` - (Optional) Available CNIs and network options for existing and new node pools of the cluster
	* `cni_type` - (Required) The CNI used by the node pools of this cluster
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
	* `open_id_connect_token_authentication_config` - (Optional) (Updatable) The properties that configure OIDC token authentication in kube-apiserver. For more information, see [Configuring the API Server](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#using-flags). 
		* `ca_certificate` - (Optional) (Updatable) A Base64 encoded public RSA or ECDSA certificates used to signed your identity provider's web certificate. 
		* `client_id` - (Optional) (Updatable) A client id that all tokens must be issued for. 
		* `groups_claim` - (Optional) (Updatable) JWT claim to use as the user's group. If the claim is present it must be an array of strings. 
		* `groups_prefix` - (Optional) (Updatable) Prefix prepended to group claims to prevent clashes with existing names (such as system:groups). 
		* `is_open_id_connect_auth_enabled` - (Required) (Updatable) Whether the cluster has OIDC Auth Config enabled. Defaults to false. 
		* `issuer_url` - (Optional) (Updatable) URL of the provider that allows the API server to discover public signing keys.  Only URLs that use the https:// scheme are accepted. This is typically the provider's discovery URL,  changed to have an empty path. 
		* `required_claims` - (Optional) (Updatable) A key=value pair that describes a required claim in the ID Token. If set, the claim is verified to be present  in the ID Token with a matching value. Repeat this flag to specify multiple claims. 
			* `key` - (Optional) (Updatable) The key of the pair.
			* `value` - (Optional) (Updatable) The value of the pair.
		* `signing_algorithms` - (Optional) (Updatable) The signing algorithms accepted. Default is ["RS256"]. 
		* `username_claim` - (Optional) (Updatable) JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end  user. Admins can choose other claims, such as email or name, depending on their provider. However, claims  other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins. 
		* `username_prefix` - (Optional) (Updatable) Prefix prepended to username claims to prevent clashes with existing names (such as system:users).  For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and  --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where  ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. 
	* `open_id_connect_discovery` - (Optional) (Updatable) The property that define the status of the OIDC Discovery feature for a cluster. 
		* `is_open_id_connect_discovery_enabled` - (Optional) (Updatable) Whether the cluster has OIDC Discovery enabled. Defaults to false. If set to true, the cluster will be assigned a public OIDC Discovery endpoint. 
	* `persistent_volume_config` - (Optional) (Updatable) Configuration to be applied to block volumes created by Kubernetes Persistent Volume Claims (PVC)
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_config` - (Optional) (Updatable) Configuration to be applied to load balancers created by Kubernetes services
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_subnet_ids` - (Optional) The OCIDs of the subnets used for Kubernetes services load balancers.
* `type` - (Optional) (Updatable) Type of cluster
* `vcn_id` - (Required) The OCID of the virtual cloud network (VCN) in which to create the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `open_id_connect_discovery_endpoint` - The cluster-specific OpenID Connect Discovery endpoint 
* `options` - Optional attributes for the cluster.
	* `add_ons` - Configurable cluster add-ons
		* `is_kubernetes_dashboard_enabled` - Whether or not to enable the Kubernetes Dashboard add-on.
		* `is_tiller_enabled` - Whether or not to enable the Tiller add-on.
	* `admission_controller_options` - Configurable cluster admission controllers
		* `is_pod_security_policy_enabled` - Whether or not to enable the Pod Security Policy admission controller.
	* `kubernetes_network_config` - Network configuration for Kubernetes.
		* `pods_cidr` - The CIDR block for Kubernetes pods. Optional, defaults to 10.244.0.0/16.
		* `services_cidr` - The CIDR block for Kubernetes services. Optional, defaults to 10.96.0.0/16.
	* `open_id_connect_token_authentication_config` - The properties that configure OIDC token authentication in kube-apiserver. For more information, see [Configuring the API Server](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#using-flags). 
		* `ca_certificate` - A Base64 encoded public RSA or ECDSA certificates used to signed your identity provider's web certificate. 
		* `client_id` - A client id that all tokens must be issued for. 
		* `groups_claim` - JWT claim to use as the user's group. If the claim is present it must be an array of strings. 
		* `groups_prefix` - Prefix prepended to group claims to prevent clashes with existing names (such as system:groups). 
		* `is_open_id_connect_auth_enabled` - Whether the cluster has OIDC Auth Config enabled. Defaults to false. 
		* `issuer_url` - URL of the provider that allows the API server to discover public signing keys.  Only URLs that use the https:// scheme are accepted. This is typically the provider's discovery URL,  changed to have an empty path. 
		* `required_claims` - A key=value pair that describes a required claim in the ID Token. If set, the claim is verified to be present  in the ID Token with a matching value. Repeat this flag to specify multiple claims. 
			* `key` - The key of the pair.
			* `value` - The value of the pair.
		* `signing_algorithms` - The signing algorithms accepted. Default is ["RS256"]. 
		* `username_claim` - JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end  user. Admins can choose other claims, such as email or name, depending on their provider. However, claims  other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins. 
		* `username_prefix` - Prefix prepended to username claims to prevent clashes with existing names (such as system:users).  For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and  --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where  ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. 
	* `open_id_connect_discovery` - The property that define the status of the OIDC Discovery feature for a cluster. 
		* `is_open_id_connect_discovery_enabled` - Whether the cluster has OIDC Discovery enabled. Defaults to false. If set to true, the cluster will be assigned a public OIDC Discovery endpoint. 
	* `persistent_volume_config` - Configuration to be applied to block volumes created by Kubernetes Persistent Volume Claims (PVC)
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_config` - Configuration to be applied to load balancers created by Kubernetes services
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `service_lb_subnet_ids` - The OCIDs of the subnets used for Kubernetes services load balancers.
* `state` - The state of the cluster masters.
* `type` - Type of cluster. Values can be BASIC_CLUSTER or ENHANCED_CLUSTER. For more information, see [Cluster Types](https://docs.cloud.oracle.com/iaas/Content/ContEng/Tasks/contengcomparingenhancedwithbasicclusters_topic.htm)
* `vcn_id` - The OCID of the virtual cloud network (VCN) in which the cluster exists.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Cluster
	* `update` - (Defaults to 1 hours), when updating the Cluster
	* `delete` - (Defaults to 1 hours), when destroying the Cluster


## Import

Clusters can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster.test_cluster "id"
```

