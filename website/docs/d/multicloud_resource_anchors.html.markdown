---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_resource_anchors"
sidebar_current: "docs-oci-datasource-multicloud-resource_anchors"
description: |-
  Provides the list of Resource Anchors in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_resource_anchors
This data source provides the list of Resource Anchors in Oracle Cloud Infrastructure Multicloud service.

Gets a list of ResourceAnchors.

## Example Usage

```hcl
data "oci_multicloud_resource_anchors" "test_resource_anchors" {
	#Optional
	compartment_id 					= var.compartment_id
	linked_compartment_id 			= var.linked_compartment_id
	lifecycle_state 				= var.lifecycle_state
	display_name 					= var.resource_anchor_display_name
	id 								= var.resource_anchor_id
	is_compartment_id_in_subtree 	= var.is_compartment_id_in_subtree
	should_fetch_compartment_name	= var.should_fetch_compartment_name
	subscription_service_name 		= var.subscription_service_name
	subscription_id 				= var.subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment or sub-compartment in which to list resources.  A Multicloud base compartment is an Oracle Cloud Infrastructure compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).  
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
* `is_compartment_id_in_subtree` - (Optional) Check the sub-compartments of a given compartmentId
* `linked_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment linked to the resource.
* `should_fetch_compartment_name` - (Optional) Whether to fetch and include the compartment name, setting this field to yes may introduce additional latency.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Optional) The subscription service name of the Cloud Service Provider.


## Attributes Reference

The following attributes are exported:

* `resource_anchor_collection` - The list of ResourceAnchorCollection.

### ResourceAnchorCollection Reference

The following attributes are exported:

* `items` - List of ResourceAnchorSummary
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
    * `compartment_name` - The name assigned to the compartment during creation.
    * `partner_cloud_account_identifier` - Partner Cloud Account Identifier of the Cloud Service Provider.
    * `csp_resource_anchor_id` - CSP resource anchor ID.
	* `csp_resource_anchor_name` - CSP resource anchor name.
    * `csp_additional_properties` - CSP Specific Additional Properties, AzureSubnetId for Azure
	* `time_created` - The date and time the ResourceAnchor was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	* `time_updated` - The date and time the ResourceAnchor was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	* `lifecycle_state` - The current state of the ResourceAnchor.
	* `lifecycle_details` - A message that describes the current state of the ResourceAnchor in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
    * `subscription_id` - Oracle Cloud Infrastructure Subscription Id
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `linked_compartment_id` - Optional - Oracle Cloud Infrastructure compartment Id (OCID) which was created or linked by customer with resource anchor.  This compartmentId is different from where resource Anchor live.
    * `linked_compartment_name` - The name assigned to the compartment which was created or linked by customer with resource anchor. This compartment is different from where resource Anchor live.
