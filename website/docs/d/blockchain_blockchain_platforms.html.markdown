---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_blockchain_platforms"
sidebar_current: "docs-oci-datasource-blockchain-blockchain_platforms"
description: |-
  Provides the list of Blockchain Platforms in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_blockchain_platforms
This data source provides the list of Blockchain Platforms in Oracle Cloud Infrastructure Blockchain service.

Returns a list Blockchain Platform Instances in a compartment

## Example Usage

```hcl
data "oci_blockchain_blockchain_platforms" "test_blockchain_platforms" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.blockchain_platform_display_name
	state = var.blockchain_platform_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Example: `My new resource` 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `blockchain_platform_collection` - The list of blockchain_platform_collection.

### BlockchainPlatform Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `component_details` - Blockchain Platform component details.
	* `osns` - List of OSNs
		* `ad` - Availability Domain of OSN
		* `ocpu_allocation_param` - OCPU allocation parameter
			* `ocpu_allocation_number` - Number of OCPU allocation
		* `osn_key` - OSN identifier
		* `state` - The current state of the OSN.
	* `peers` - List of Peers
		* `ad` - Availability Domain of peer
		* `alias` - peer alias
		* `host` - Host on which the Peer exists
		* `ocpu_allocation_param` - OCPU allocation parameter
			* `ocpu_allocation_number` - Number of OCPU allocation
		* `peer_key` - peer identifier
		* `role` - Peer role
		* `state` - The current state of the peer.
* `compute_shape` - Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE or ENTERPRISE_CUSTOM
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Platform Instance Description
* `display_name` - Platform Instance Display name, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_ocpu_utilization_info` - List of OcpuUtilization for all hosts
	* `host` - Host name of VM
	* `ocpu_capacity_number` - Number of total OCPU capacity on the host
	* `ocpu_utilization_number` - Number of OCPU utilized
* `id` - unique identifier that is immutable on creation
* `is_byol` - Bring your own license
* `is_multi_ad` - True for multi-AD blockchain plaforms, false for single-AD
* `lifecycle_details` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `load_balancer_shape` - Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
* `platform_role` - Role of platform - FOUNDER or PARTICIPANT
* `platform_shape_type` - Type of Platform shape - DEFAULT or CUSTOM
* `platform_version` - Platform Version
* `replicas` - Number of replicas of service components like Rest Proxy, CA and Console
	* `ca_count` - Number of CA replicas
	* `console_count` - Number of console replicas
	* `proxy_count` - Number of REST proxy replicas
* `service_endpoint` - Service endpoint URL, valid post-provisioning
* `service_version` - The version of the Platform Instance.
* `state` - The current state of the Platform Instance.
* `storage_size_in_tbs` - Storage size in TBs
* `storage_used_in_tbs` - Storage used in TBs
* `time_created` - The time the the Platform Instance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Platform Instance was updated. An RFC3339 formatted datetime string
* `total_ocpu_capacity` - Number of total OCPUs allocated to the platform cluster

