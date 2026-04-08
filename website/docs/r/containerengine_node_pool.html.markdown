---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_node_pool"
sidebar_current: "docs-oci-resource-containerengine-node_pool"
description: |-
  Provides the Node Pool resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_node_pool
This resource provides the Node Pool resource in Oracle Cloud Infrastructure Container Engine service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/containerengine/latest/NodePool

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/container_engine

Create a new node pool.

## Example Usage

```hcl
resource "oci_containerengine_node_pool" "test_node_pool" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compartment_id = var.compartment_id
	name = var.node_pool_name
	node_shape = var.node_pool_node_shape

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	initial_node_labels {

		#Optional
		key = var.node_pool_initial_node_labels_key
		value = var.node_pool_initial_node_labels_value
	}
	kubernetes_version = var.node_pool_kubernetes_version
	network_launch_type = var.node_pool_network_launch_type
	node_config_details {
		#Required
		placement_configs {
			#Required
			availability_domain = var.node_pool_node_config_details_placement_configs_availability_domain
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			capacity_reservation_id = oci_containerengine_capacity_reservation.test_capacity_reservation.id
			fault_domains = var.node_pool_node_config_details_placement_configs_fault_domains
			preemptible_node_config {
				#Required
				preemption_action {
					#Required
					type = var.node_pool_node_config_details_placement_configs_preemptible_node_config_preemption_action_type

					#Optional
					is_preserve_boot_volume = var.node_pool_node_config_details_placement_configs_preemptible_node_config_preemption_action_is_preserve_boot_volume
				}
			}
		}
		size = var.node_pool_node_config_details_size

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		freeform_tags = {"Department"= "Finance"}
		is_pv_encryption_in_transit_enabled = var.node_pool_node_config_details_is_pv_encryption_in_transit_enabled
		kms_key_id = oci_kms_key.test_key.id
		node_pool_pod_network_option_details {
			#Required
			cni_type = var.node_pool_node_config_details_node_pool_pod_network_option_details_cni_type

			#Optional
			max_pods_per_node = var.node_pool_node_config_details_node_pool_pod_network_option_details_max_pods_per_node
			pod_nsg_ids = var.node_pool_node_config_details_node_pool_pod_network_option_details_pod_nsg_ids
			pod_subnet_ids = var.node_pool_node_config_details_node_pool_pod_network_option_details_pod_subnet_ids
		}
		nsg_ids = var.node_pool_node_config_details_nsg_ids
	}
	node_eviction_node_pool_settings {

		#Optional
		eviction_grace_duration = var.node_pool_node_eviction_node_pool_settings_eviction_grace_duration
		is_force_action_after_grace_duration = var.node_pool_node_eviction_node_pool_settings_is_force_action_after_grace_duration
		is_force_delete_after_grace_duration = var.node_pool_node_eviction_node_pool_settings_is_force_delete_after_grace_duration
	}
	node_image_name = oci_core_image.test_image.name
	node_metadata = var.node_pool_node_metadata
	node_pool_cycling_details {

		#Optional
		cycle_modes = var.node_pool_node_pool_cycling_details_cycle_modes
		is_node_cycling_enabled = var.node_pool_node_pool_cycling_details_is_node_cycling_enabled
		maximum_surge = var.node_pool_node_pool_cycling_details_maximum_surge
		maximum_unavailable = var.node_pool_node_pool_cycling_details_maximum_unavailable
	}
	node_shape_config {

		#Optional
		memory_in_gbs = var.node_pool_node_shape_config_memory_in_gbs
		ocpus = var.node_pool_node_shape_config_ocpus
	}
	node_source_details {
		#Required
		image_id = oci_core_image.test_image.id
		source_type = var.node_pool_node_source_details_source_type

		#Optional
		boot_volume_size_in_gbs = var.node_pool_node_source_details_boot_volume_size_in_gbs
	}
	quantity_per_subnet = var.node_pool_quantity_per_subnet
	secondary_vnics {
		#Required
		create_vnic_details {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			application_resources = var.node_pool_secondary_vnics_create_vnic_details_application_resources
			assign_ipv6ip = var.node_pool_secondary_vnics_create_vnic_details_assign_ipv6ip
			assign_public_ip = var.node_pool_secondary_vnics_create_vnic_details_assign_public_ip
			defined_tags = {"Operations.CostCenter"= "42"}
			display_name = var.node_pool_secondary_vnics_create_vnic_details_display_name
			freeform_tags = {"Department"= "Finance"}
			ip_count = var.node_pool_secondary_vnics_create_vnic_details_ip_count
			ipv6address_ipv6subnet_cidr_pair_details {

				#Optional
				ipv6address = var.node_pool_secondary_vnics_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6address
				ipv6subnet_cidr = var.node_pool_secondary_vnics_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
			}
			nsg_ids = var.node_pool_secondary_vnics_create_vnic_details_nsg_ids
			skip_source_dest_check = var.node_pool_secondary_vnics_create_vnic_details_skip_source_dest_check
		}

		#Optional
		display_name = var.node_pool_secondary_vnics_display_name
		nic_index = var.node_pool_secondary_vnics_nic_index
	}
	ssh_public_key = var.node_pool_ssh_public_key
	subnet_ids = var.node_pool_subnet_ids
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster to which this node pool is attached.
* `compartment_id` - (Required) The OCID of the compartment in which the node pool exists.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `initial_node_labels` - (Optional) (Updatable) A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	* `key` - (Optional) (Updatable) The key of the pair.
	* `value` - (Optional) (Updatable) The value of the pair.
* `kubernetes_version` - (Optional) (Updatable) The version of Kubernetes to install on the nodes in the node pool.
* `name` - (Required) (Updatable) The name of the node pool. Avoid entering confidential information.
* `network_launch_type` - (Optional) (Updatable) Emulation type for the physical network interface card (NIC) for nodes
* `node_config_details` - (Optional) (Updatable) The configuration of nodes in the node pool. Exactly one of the subnetIds or nodeConfigDetails properties must be specified. 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `is_pv_encryption_in_transit_enabled` - (Optional) (Updatable) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	* `kms_key_id` - (Optional) (Updatable) The OCID of the Key Management Service key assigned to the boot volume.
	* `node_pool_pod_network_option_details` - (Optional) (Updatable) The CNI related configuration of pods in the node pool. 
		* `cni_type` - (Required) (Updatable) The CNI plugin used by this node pool
		* `max_pods_per_node` - (Applicable when cni_type=OCI_VCN_IP_NATIVE) (Updatable) The max number of pods per node in the node pool. This value will be limited by the number of VNICs attachable to the node pool shape 
		* `pod_nsg_ids` - (Applicable when cni_type=OCI_VCN_IP_NATIVE) (Updatable) The OCIDs of the Network Security Group(s) to associate pods for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
		* `pod_subnet_ids` - (Applicable when cni_type=OCI_VCN_IP_NATIVE) (Updatable) The OCIDs of the subnets in which to place pods for this node pool. This can be one of the node pool subnet IDs 
	* `nsg_ids` - (Optional) (Updatable) The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `placement_configs` - (Required) (Updatable) The placement configurations for the node pool. Provide one placement configuration for each availability domain in which you intend to launch a node.

		To use the node pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
		* `availability_domain` - (Required) (Updatable) The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - (Optional) (Updatable) The OCID of the compute capacity reservation in which to place the compute instance.
		* `fault_domains` - (Optional) (Updatable) A list of fault domains in which to place nodes. 
		* `preemptible_node_config` - (Optional) (Updatable) Configuration options for preemptible nodes.
			* `preemption_action` - (Required) (Updatable) The action to run when the preemptible node is interrupted for eviction.
				* `is_preserve_boot_volume` - (Optional) (Updatable) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - (Required) (Updatable) The type of action to run when the instance is interrupted for eviction.
		* `subnet_id` - (Required) (Updatable) The OCID of the subnet in which to place nodes.
	* `size` - (Required) (Updatable) The number of nodes that should be in the node pool. 
* `node_eviction_node_pool_settings` - (Optional) (Updatable) Node Eviction Details configuration
	* `eviction_grace_duration` - (Optional) (Updatable) Duration after which OKE will give up eviction of the pods on the node. PT0M will indicate you want to delete the node without cordon and drain. Default PT60M, Min PT0M, Max: PT60M. Format ISO 8601 e.g PT30M 
	* `is_force_action_after_grace_duration` - (Optional) (Updatable) If the node action should be performed if not all the pods can be evicted in the grace period
	* `is_force_delete_after_grace_duration` - (Optional) (Updatable) If the underlying compute instance should be deleted if you cannot evict all the pods in grace period
* `node_image_name` - (Optional) Deprecated. Use `nodeSourceDetails` instead. If you specify values for both, this value is ignored. The name of the image running on the nodes in the node pool. Cannot be used when `node_image_id` is specified.
* `node_metadata` - (Optional) (Updatable) A list of key/value pairs to add to each underlying Oracle Cloud Infrastructure instance in the node pool on launch.
* `node_pool_cycling_details` - (Optional) (Updatable) Node Pool Cycling Details
	* `cycle_modes` - (Optional) (Updatable) An ordered list of cycle modes that should be performed on the OKE nodes.  
	* `is_node_cycling_enabled` - (Optional) (Updatable) If cycling operation should be performed on the nodes in the node pool.
	* `maximum_surge` - (Optional) (Updatable) Maximum additional new compute instances that would be temporarily created and added to nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 1, Ranges from 0 to Nodepool size or 0% to 100% 
	* `maximum_unavailable` - (Optional) (Updatable) Maximum active nodes that would be terminated from nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 0, Ranges from 0 to Nodepool size or 0% to 100% 
* `node_shape` - (Required) (Updatable) The name of the node shape of the nodes in the node pool.
* `node_shape_config` - (Optional) (Updatable) Specify the configuration of the shape to launch nodes in the node pool. 
	* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to each node, in gigabytes. 
	* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to each node in the node pool. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `node_source_details` - (Optional) (Updatable) Specify the source to use to launch nodes in the node pool. Currently, image is the only supported source. 
	* `boot_volume_size_in_gbs` - (Optional) (Updatable) The size of the boot volume in GBs. Minimum value is 50 GB. See [here](https://docs.cloud.oracle.com/en-us/iaas/Content/Block/Concepts/bootvolumes.htm) for max custom boot volume sizing and OS-specific requirements.
	* `image_id` - (Required) (Updatable) The OCID of the image used to boot the node.
	* `source_type` - (Required) (Updatable) The source type for the node. Use `IMAGE` when specifying an OCID of an image. 
* `quantity_per_subnet` - (Optional) (Updatable) Optional, default to 1. The number of nodes to create in each subnet specified in subnetIds property. When used, subnetIds is required. This property is deprecated, use nodeConfigDetails instead. 
* `secondary_vnics` - (Optional) (Updatable) A list of secondary vnics to attach to nodes
	* `create_vnic_details` - (Required) (Updatable) The properties of the secondary vnics
		* `application_resources` - (Optional) (Updatable) The application resource that corresponds to this secondary vnic. Used to map pods to this specific vnic for scheduling 
		* `assign_ipv6ip` - (Optional) (Updatable) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet 
		* `assign_public_ip` - (Optional) (Updatable) Whether the VNIC should be assigned a public IP address 
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - (Optional) (Updatable) Display name for secondary vnic
		* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
		* `ip_count` - (Optional) (Updatable) The number of ip addresses to attach to secondary vnic 
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) (Updatable) A list of IPv6 prefixes from which the VNIC should be assigned an IPv6 address. You can provide only the prefix  and Oracle Cloud Infrastructure selects an available address from the range. You can optionally choose to leave the prefix range empty  and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6address` - (Optional) (Updatable) An IPv6 address of your choice. Must be an available IPv6 address within the subnet's prefix 
			* `ipv6subnet_cidr` - (Optional) (Updatable) The IPv6 prefix allocated to the subnet 
		* `nsg_ids` - (Optional) (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to 
		* `skip_source_dest_check` - (Optional) (Updatable) Whether the source/destination check is disabled on the VNIC 
		* `subnet_id` - (Required) (Updatable) the ocid of the subnet to create the vnic in
	* `display_name` - (Optional) (Updatable) Display name for vnic attachment
	* `nic_index` - (Optional) (Updatable) Which physical network interface card (NIC) the VNIC will use
* `ssh_public_key` - (Optional) (Updatable) The SSH public key on each node in the node pool on launch.
* `subnet_ids` - (Optional) (Updatable) The OCIDs of the subnets in which to place nodes for this node pool. When used, quantityPerSubnet can be provided. This property is deprecated, use nodeConfigDetails. Exactly one of the subnetIds or nodeConfigDetails properties must be specified. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_id` - The OCID of the cluster to which this node pool is attached.
* `compartment_id` - The OCID of the compartment in which the node pool exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the node pool.
* `initial_node_labels` - A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	* `key` - The key of the pair.
	* `value` - The value of the pair.
* `kubernetes_version` - The version of Kubernetes running on the nodes in the node pool.
* `lifecycle_details` - Details about the state of the nodepool.
* `name` - The name of the node pool.
* `network_launch_type` - Emulation type for the physical network interface card (NIC) for nodes
* `node_config_details` - The configuration of nodes in the node pool.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	* `kms_key_id` - The OCID of the Key Management Service key assigned to the boot volume.
	* `node_pool_pod_network_option_details` - The CNI related configuration of pods in the node pool. 
		* `cni_type` - The CNI plugin used by this node pool
		* `max_pods_per_node` - The max number of pods per node in the node pool. This value will be limited by the number of VNICs attachable to the node pool shape 
		* `pod_nsg_ids` - The OCIDs of the Network Security Group(s) to associate pods for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
		* `pod_subnet_ids` - The OCIDs of the subnets in which to place pods for this node pool. This can be one of the node pool subnet IDs 
	* `nsg_ids` - The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `placement_configs` - The placement configurations for the node pool. Provide one placement configuration for each availability domain in which you intend to launch a node.

		To use the node pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
		* `availability_domain` - The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - The OCID of the compute capacity reservation in which to place the compute instance.
		* `fault_domains` - A list of fault domains in which to place nodes. 
		* `preemptible_node_config` - Configuration options for preemptible nodes.
			* `preemption_action` - The action to run when the preemptible node is interrupted for eviction.
				* `is_preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - The type of action to run when the instance is interrupted for eviction.
		* `subnet_id` - The OCID of the subnet in which to place nodes.
	* `size` - The number of nodes in the node pool. 
* `node_eviction_node_pool_settings` - Node Eviction Details configuration
	* `eviction_grace_duration` - Duration after which OKE will give up eviction of the pods on the node. PT0M will indicate you want to delete the node without cordon and drain. Default PT60M, Min PT0M, Max: PT60M. Format ISO 8601 e.g PT30M 
	* `is_force_action_after_grace_duration` - If the node action should be performed if not all the pods can be evicted in the grace period
	* `is_force_delete_after_grace_duration` - If the underlying compute instance should be deleted if you cannot evict all the pods in grace period
* `node_image_id` - Deprecated. see `nodeSource`. The OCID of the image running on the nodes in the node pool. 
* `node_image_name` - Deprecated. see `nodeSource`. The name of the image running on the nodes in the node pool. 
* `node_metadata` - A list of key/value pairs to add to each underlying Oracle Cloud Infrastructure instance in the node pool on launch.
* `node_pool_cycling_details` - Node Pool Cycling Details
	* `cycle_modes` - An ordered list of cycle modes that should be performed on the OKE nodes.  
	* `is_node_cycling_enabled` - If cycling operation should be performed on the nodes in the node pool.
	* `maximum_surge` - Maximum additional new compute instances that would be temporarily created and added to nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 1, Ranges from 0 to Nodepool size or 0% to 100% 
	* `maximum_unavailable` - Maximum active nodes that would be terminated from nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 0, Ranges from 0 to Nodepool size or 0% to 100% 
* `node_shape` - The name of the node shape of the nodes in the node pool.
* `node_shape_config` - The shape configuration of the nodes.
	* `memory_in_gbs` - The total amount of memory available to each node, in gigabytes. 
	* `ocpus` - The total number of OCPUs available to each node in the node pool. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `node_source` - Deprecated. see `nodeSourceDetails`. Source running on the nodes in the node pool.
	* `image_id` - The OCID of the image.
	* `source_name` - The user-friendly name of the entity corresponding to the OCID. 
	* `source_type` - The source type of this option. `IMAGE` means the OCID is of an image. 
* `node_source_details` - Source running on the nodes in the node pool.
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB. See [here](https://docs.cloud.oracle.com/en-us/iaas/Content/Block/Concepts/bootvolumes.htm) for max custom boot volume sizing and OS-specific requirements.
	* `image_id` - The OCID of the image used to boot the node.
	* `source_type` - The source type for the node. Use `IMAGE` when specifying an OCID of an image. 
* `nodes` - The nodes in the node pool.
	* `availability_domain` - The name of the availability domain in which this node is placed.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `error` - An error that may be associated with the node.
		* `code` - A short error code that defines the upstream error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
		* `message` - A human-readable error string of the upstream error.
		* `status` - The status of the HTTP response encountered in the upstream error.
	* `fault_domain` - The fault domain of this node.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the compute instance backing this node.
	* `kubernetes_version` - The version of Kubernetes this node is running.
	* `lifecycle_details` - Details about the state of the node.
	* `name` - The name of the node.
	* `node_pool_id` - The OCID of the node pool to which this node belongs.
	* `private_ip` - The private IP address of this node.
	* `public_ip` - The public IP address of this node.
	* `state` - The state of the node. For more information, see [Monitoring Clusters](https://docs.cloud.oracle.com/iaas/Content/ContEng/Tasks/contengmonitoringclusters.htm)
	* `subnet_id` - The OCID of the subnet in which this node is placed.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `quantity_per_subnet` - The number of nodes in each subnet.
* `secondary_vnics` - A list of secondary vnics to attach to nodes
	* `create_vnic_details` - The properties of the secondary vnics
		* `application_resources` - The application resource that corresponds to this secondary vnic. Used to map pods to this specific vnic for scheduling 
		* `assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet 
		* `assign_public_ip` - Whether the VNIC should be assigned a public IP address 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - Display name for secondary vnic
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
		* `ip_count` - The number of ip addresses to attach to secondary vnic 
		* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefixes from which the VNIC should be assigned an IPv6 address. You can provide only the prefix  and Oracle Cloud Infrastructure selects an available address from the range. You can optionally choose to leave the prefix range empty  and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6address` - An IPv6 address of your choice. Must be an available IPv6 address within the subnet's prefix 
			* `ipv6subnet_cidr` - The IPv6 prefix allocated to the subnet 
		* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to 
		* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC 
		* `subnet_id` - the ocid of the subnet to create the vnic in
	* `display_name` - Display name for vnic attachment
	* `nic_index` - Which physical network interface card (NIC) the VNIC will use
* `ssh_public_key` - The SSH public key on each node in the node pool on launch.
* `state` - The state of the nodepool. For more information, see [Monitoring Clusters](https://docs.cloud.oracle.com/iaas/Content/ContEng/Tasks/contengmonitoringclusters.htm)
* `subnet_ids` - The OCIDs of the subnets in which to place nodes for this node pool.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 50 minutes), when creating the Node Pool
	* `update` - (Defaults to 50 minutes), when updating the Node Pool
	* `delete` - (Defaults to 50 minutes), when destroying the Node Pool


## Import

NodePools can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_node_pool.test_node_pool "id"
```

