---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_blockchain_platform"
sidebar_current: "docs-oci-resource-blockchain-blockchain_platform"
description: |-
  Provides the Blockchain Platform resource in Oracle Cloud Infrastructure Blockchain service
---

# oci_blockchain_blockchain_platform
This resource provides the Blockchain Platform resource in Oracle Cloud Infrastructure Blockchain service.

Creates a new Blockchain Platform.


## Example Usage

```hcl
resource "oci_blockchain_blockchain_platform" "test_blockchain_platform" {
	#Required
	compartment_id = var.compartment_id
	compute_shape = var.blockchain_platform_compute_shape
	display_name = var.blockchain_platform_display_name
	platform_role = var.blockchain_platform_platform_role

	#Optional
	ca_cert_archive_text = var.blockchain_platform_ca_cert_archive_text
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.blockchain_platform_description
	federated_user_id = oci_identity_user.test_user.id
	freeform_tags = {"bar-key"= "value"}
	idcs_access_token = var.blockchain_platform_idcs_access_token
	is_byol = var.blockchain_platform_is_byol
}
```

## Argument Reference

The following arguments are supported:

* `ca_cert_archive_text` - (Optional) Base64 encoded text in ASCII character set of a Thirdparty CA Certificates archive file. The Archive file is a zip file containing third part CA Certificates, the ca key and certificate files used when issuing enrollment certificates (ECerts) and transaction certificates (TCerts). The chainfile (if it exists) contains the certificate chain which should be trusted for this CA, where the 1st in the chain is always the root CA certificate. File list in zip file [ca-cert.pem,ca-key.pem,ca-chain.pem(optional)]. 
* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `compute_shape` - (Required) Type of compute shape - one of Standard, (Enterprise) Small, Medium, Large or Extra Large
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Platform Instance Description
* `display_name` - (Required) Platform Instance Display name, can be renamed
* `federated_user_id` - (Optional) Identifier for a federated user
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_access_token` - (Optional) IDCS access token
* `is_byol` - (Optional) Bring your own license
* `platform_role` - (Required) Role of platform - founder or participant


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `compute_shape` - Type of compute shape - one of Standard, (Enterprise) Small, Medium, Large or Extra Large
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
* `platform_role` - Role of platform - founder or participant
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

## Import

BlockchainPlatforms can be imported using the `id`, e.g.

```
$ terraform import oci_blockchain_blockchain_platform.test_blockchain_platform "id"
```

