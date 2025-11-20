---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_network_anchors"
sidebar_current: "docs-oci-datasource-multicloud-network_anchors"
description: |-
  Provides the list of Network Anchors in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_network_anchors
This data source provides the list of Network Anchors in Oracle Cloud Infrastructure Multicloud service.

Gets a list of NetworkAnchors.

## Example Usage

```hcl
data "oci_multicloud_network_anchors" "test_network_anchors" {

	#Optional
	compartment_id 					= var.compartment_id
	subscription_id 				= var.subscription_id
	subscription_service_name 		= var.subscription_service_name
	network_anchor_lifecycle_state 	= var.network_anchor_lifecycle_state
	display_name 					= var.display_name
	external_location 				= var.external_location
	network_anchor_oci_subnet_id 	= var.network_anchor_oci_subnet_id
	compartment_id_in_subtree 		= var.compartment_id_in_subtree
	network_anchor_oci_vcn_id 		= var.network_anchor_oci_vcn_id
	id 								= var.id
	should_fetch_vcn_name 			= var.should_fetch_vcn_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment or sub-compartment in which to list resources.  A Multicloud base compartment is an Oracle Cloud Infrastructure compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).  
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Optional) The subscription service name of the Cloud Service Provider.
* `network_anchor_lifecycle_state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `external_location` - (Optional) The Cloud Service Provider region.
* `network_anchor_oci_subnet_id` - (Optional) A filter to return only NetworkAnchor resources that match the given Oracle Cloud Infrastructure subnet Id.
* `compartment_id_in_subtree` - (Optional) If set to true, a list operation will return NetworkAnchors from all child compartments in the provided compartmentId parameter.
* `network_anchor_oci_vcn_id` - (Optional) A filter to return only NetworkAnchor resources that match the given Oracle Cloud Infrastructure Vcn Id.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
* `should_fetch_vcn_name` - (Optional) Whether to fetch and include the vcn display name, which may introduce additional latency.

Note: one of the arguments `compartment_id` or `id` must be specified.

## Attributes Reference

The following attributes are exported:

* `network_anchor_collection` - The list of NetworkAnchor.

### NetworkAnchor Reference

The following attributes are exported

* `items` - List of NetworkAnchorSummary
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
    * `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
    * `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
    * `resource_anchor_id` - Oracle Cloud Infrastructure resource anchor Id (OCID).
    * `vcn_id` - Oracle Cloud Infrastructure VCN OCID. CSP can not set this property.
	* `vcn_name` - Name of the VCN associated to the Network Anchor.
    * `network_anchor_connection_status` - Defines status of the Network Anchor.
    * `cluster_placement_group_id` - The CPG ID in which Network Anchor will be created.
    * `time_created` - The date and time the NetworkAnchor was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
    * `time_updated` - The date and time the NetworkAnchor was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
    * `csp_additional_properties` - CSP Specific Additional Properties, AzureSubnetId for Azure
    * `csp_network_anchor_id` - Network Anchor Id in the Cloud Service Provider.
    * `network_anchor_uri` - CSP network anchor Uri
    * `network_anchor_lifecycle_state` - The current state of the NetworkAnchor.
    * `lifecycle_details` - A message that describes the current state of the NetworkAnchor in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
    * `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `subscription_type` - Oracle Cloud Infrastructure Subscription Type.
