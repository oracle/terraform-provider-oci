---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_network_anchor"
sidebar_current: "docs-oci-datasource-multicloud-network_anchor"
description: |-
  Provides details about a specific Network Anchor in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_network_anchor
This data source provides details about a specific Network Anchor resource in Oracle Cloud Infrastructure Multicloud service.

Gets information about a NetworkAnchor.

## Example Usage

```hcl
data "oci_multicloud_network_anchor" "test_network_anchor" {
	#Required
	network_anchor_id 			= var.network_anchor_id
	subscription_id 			= var.subscription_id
	subscription_service_name 	= var.subscription_service_name

	#Optional
	external_location 			= var.network_anchor_external_location
}
```

## Argument Reference

The following arguments are supported:

* `network_anchor_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.
* `subscription_service_name` - (Required) The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
* `external_location` - (Optional) OMHub Control Plane must know underlying CSP CP Region External Location Name.


## Attributes Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `resource_anchor_id` - Oracle Cloud Infrastructure resource anchor Id (OCID).
* `time_created` - The date and time the NetworkAnchor was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `time_updated` - The date and time the NetworkAnchor was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `network_anchor_lifecycle_state` - The current state of the NetworkAnchor.
* `lifecycle_details` - A message that describes the current state of the NetworkAnchor in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `setup_mode` - AUTO_BIND - when passed compartment will be created on-behalf of customer and bind to this resource anchor NO_AUTO_BIND - compartment will not be created and later customer can bind existing compartment.  to this resource anchor. This is for future use only
* `cluster_placement_group_id` - The CPG ID in which Network Anchor will be created.
* `oci_metadata_item` - Oracle Cloud Infrastructure network anchor related meta data items
	* `network_anchor_connection_status` - This can be merge to lifecycleState CONNECTED - Partner and CSI information is assigned and MulticloudLink provisioned. DISCONNECTED - Only partner cloud information is assigned. CONNECTING - Oracle Cloud Infrastructure information is assigned and the control plane is provisioning resources. ACTIVE - Network anchor is connected and resources (VNICs) exist within a subnet. ERROR - DRG attach fails during connection. FAILED - Network anchor creation failed NEEDS_ATTENTION - Network anchor is in temporary bad state UPDATING - Network anchor is getting updated. DELETING - Network anchor is getting deleted DELETED - A connected network anchor is deleted.
	* `vcn` - Oracle Cloud Infrastructure VCN basic information object. It is optional and planned to used for future for network anchor
		* `backup_cidr_blocks` - Oracle Cloud Infrastructure backup cidr block. CSP can set this property It's optional only if disconnect anchor is allowed. IPv4 CIDR blocks for the VCN that meet the following criteria Type: [string (length: 1–32), ...] The CIDR blocks must be valid. They must not overlap with each other or with the on-premises network CIDR block.
		* `cidr_blocks` - Oracle Cloud Infrastructure primary cidr block. CSP can set this property It's optional only if disconnect anchor is allowed IPv4 CIDR blocks for the VCN that meet the following criteria Type: [string (length: 1–32), ...] The CIDR blocks must be valid. They must not overlap with each other or with the on-premises network CIDR block.
		* `dns_label` - Oracle Cloud Infrastructure DNS label. This is optional if DNS config is provided.
		* `vcn_id` - Oracle Cloud Infrastructure VCN OCID. CSP can not set this property.
	* `dns` - Oracle Cloud Infrastructure network anchor related meta data items
		* `custom_domain_name` - Full custom domain name. If this field is passed dnsLabel will be ignored
	* `subnets` - Network subnets
		* `label` - Subnet label. CSP can set this property
		* `subnet_id` - OCID for existing the subnet. CSP can not set this property.
		* `type` - Defines if the subnet is the primary or backup for the network
	* `dns_listening_endpoint_ip_address` - The DNS Listener Endpoint Address.
	* `dns_forwarding_endpoint_ip_address` - The DNS Listener Forwarding Address.
	* `dns_forwarding_config` - DNS forward configuration
* `cloud_service_provider_metadata_item` - Cloud Service Provider metadata item. Warning - In future this object can change to generic object with future Cloud Service Provider based on  CloudServiceProvider field. This can be one of CSP provider type Azure, GCP and AWS 
	* `cidr_blocks` - An Azure/GCP/AWS cidrBlocks
	* `dns_forwarding_config` - DNS domain ip mapping forwarding configuration
	* `network_anchor_uri` - CSP network anchor Uri
	* `odb_network_id` - CSP oracle database network anchor unique ID/name
	* `region` - Azure/GCP/AWS region

