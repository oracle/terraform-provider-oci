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
	#Required
	subscription_id 				= var.subscription_id
	subscription_service_name 		= var.subscription_service_name

	#Optional
	compartment_id 					= var.compartment_id
	display_name 					= var.resource_anchor_display_name
	id 								= var.resource_anchor_id
	is_compartment_id_in_subtree 	= var.resource_anchor_is_compartment_id_in_subtree
	linked_compartment_id 			= var.linked_compartment_id
	lifecycle_state 				= var.resource_anchor_state
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.
* `subscription_service_name` - (Required) The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
* `is_compartment_id_in_subtree` - (Optional) Check the sub-compartments of a given compartmentId
* `linked_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which linked to Resource.
* `lifecycle_state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `resource_anchor_collection` - The list of ResourceAnchorCollection.

### ResourceAnchorCollection Reference

The following attributes are exported:

* `items` - List of ResourceAnchorSummary
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
    * `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
    * `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
    * `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
    * `lifecycle_details` - A message that describes the current state of the ResourceAnchor in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
    * `lifecycle_state` - The current state of the ResourceAnchor.
    * `subscription_id` - Oracle Cloud Infrastructure Subscription Id
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `time_created` - The date and time the ResourceAnchor was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
    * `time_updated` - The date and time the ResourceAnchor was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`