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

Gets a list of multicloud resources with multicloud base compartment and subscription across Cloud Service Providers.


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
* `external_location` - (Optional) The Cloud Service Provider region.
* `resource_anchor_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Required) The subscription service name of the Cloud Service Provider.


## Attributes Reference

The following attributes are exported:

* `multicloud_resource_collection` - The list of OmHubMulticloudResource.

### OmHubMulticloudResource Reference

The following attributes are exported:

* `items` - List of MulticloudResourceSummary.
	* `compartment_id` - Compartment Id of the resource.
	* `compartment_name` - Compartment name associated the resource.
	* `csp_additional_properties` - CSP Specific Additional Properties, AzureSubnetId for Azure
	* `csp_resource_id` - Resource Id that comes from the Multi Cloud Control Plane
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `network_anchor_id` - OCID of the Network Anchor
	* `network_anchor_name` - Name of the network anchor associated to the resource.
	* `resource_display_name` - Endpoint used to retrieve displayName and lifeCycleState of the resource.
	* `resource_id` - The Id of the multicloud resource.
	* `resource_type` - What resource it refers to. Eg. VMCluster, ExaInfra, etc.
	* `lifecycle_state` - The current state of the multicloud resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the subscription was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The date and time the subscription was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `vcn_id` - Id of the Virtual Cloud Network associated to the resource.
	* `vcn_name` - Resource Anchor name.
