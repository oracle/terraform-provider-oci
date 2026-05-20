---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_om_hub_multicloud_resources"
sidebar_current: "docs-oci-datasource-multicloud-om_hub_multicloud_resources"
description: |-
  Provides the list of Om Hub Multicloud Resources in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_om_hub_multicloud_resources
This data source provides the list of Om Hub Multicloud Resources in Oracle Cloud Infrastructure Multicloud service.

Lists Multicloud resources in the specified Multicloud subscription.
Details for each resource include Multicloud base compartment, name, state, resource type, and network anchor.
For more information, see
[Multicloud Resources](https://docs.cloud.oracle.com/iaas/Content/multicloud-hub/list-resources.htm).


## Example Usage

```hcl
data "oci_multicloud_om_hub_multicloud_resources" "test_om_hub_multicloud_resources" {
	#Required
	subscription_id 			= var.subscription_id
	subscription_service_name 	= var.subscription_service_name

	#Optional
	compartment_id 				= var.compartment_id
	external_location 			= var.external_location
	resource_anchor_id 			= var.resource_anchor_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `external_location` - (Optional) The cloud service provider region.
* `resource_anchor_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource anchor.
* `resource_type` - (Optional) Filter alerts by resource type (e.g. ADBD, VMCluster).
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Optional) The cloud service provider.


## Attributes Reference

The following attributes are exported:

* `multicloud_resource_collection` - The list of OmHubMulticloudResource.

### OmHubMulticloudResource Reference

The following attributes are exported:

* `items` - List of MulticloudResourceSummary.
	* `compartment_id` - Id of the compartment associated with the resource.
	* `compartment_name` - Name of the compartment associated with the resource.
	* `csp_additional_properties` - Properties specific to the cloud service provider. For example, AzureSubnetId for Azure.
	* `csp_resource_id` - The resource Id that comes from the Multicloud control plane.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `network_anchor_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network anchor associated with the resource.
	* `network_anchor_name` - Name of the network anchor associated with the resource.
	* `resource_additional_properties` - Additional attributes specific to certain resource types, used to construct a URL for accessing the resource in the Oracle Cloud Infrastructure console.
	* `resource_display_name` - Endpoint used to retrieve the resource's display name and lifecycle state.
	* `resource_id` - The Id of the multicloud resource.
	* `resource_type` - Type of resource, such as `VMCluster` or `ExaInfra`,
	* `lifecycle_state` - The current state of the Multicloud resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the subscription was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The date and time the subscription was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `vcn_id` - Id of the virtual cloud network (VCN) associated with the resource.
	* `vcn_name` - Name of the virtual cloud network (VCN) associated with the resource.

