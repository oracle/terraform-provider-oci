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

