---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_resource_anchor"
sidebar_current: "docs-oci-datasource-multicloud-resource_anchor"
description: |-
  Provides details about a specific Resource Anchor in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_resource_anchor
This data source provides details about a specific Resource Anchor resource in Oracle Cloud Infrastructure Multicloud service.

Gets information about a ResourceAnchor.

## Example Usage

```hcl
data "oci_multicloud_resource_anchor" "test_resource_anchor" {
	#Required
	resource_anchor_id 			= var.resource_anchor_id
	subscription_id 			= var.subscription_id
	subscription_service_name 	= var.subscription_service_name
	
	#Optional
	should_fetch_compartment_name = var.should_fetch_compartment_name
}
```

## Argument Reference

The following arguments are supported:

* `resource_anchor_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Required) The subscription service name of the Cloud Service Provider.
* `should_fetch_compartment_name` - (Optional) Whether to fetch and include the compartment name, setting this field to yes may introduce additional latency.


## Attributes Reference

The following attributes are exported:

* `cloud_service_provider_metadata_item` - Cloud Service Provider metadata item. Warning - In future this object can change to generic object with future Cloud Service Provider based on  CloudServiceProvider field. This can be one of CSP provider type Azure, GCP and AWS. 
	* `account_id` - AWS accountId that was used for creating this resource anchor resource.
	* `csp_additional_properties` - CSP Specific Additional Properties, AzureSubnetId for Azure
	* `csp_resource_anchor_id` - CSP resource anchor ID.
	* `csp_resource_anchor_name` - CSP resource anchor name.
	* `project_number` - GCP project number that was used for creating this resource anchor resource.
	* `region` - The Azure, AWS or GCP region.
	* `resource_anchor_name` - Oracle Cloud Infrastructure resource anchor name.
	* `resource_anchor_uri` - CSP resource anchor Uri.
	* `resource_group` - Azure resource group that was used for creating this resource.
	* `subscription` - Azure subscription that was used for creating this resource.
	* `subscription_type` - Oracle Cloud Infrastructure Subscription Type.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_name` - The name assigned to the compartment during creation.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `lifecycle_details` - A message that describes the current state of the ResourceAnchor in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `linked_compartment_id` - Optional - Oracle Cloud Infrastructure compartment Id (OCID) which was created or linked by customer with resource anchor.  This compartmentId is different from where resource Anchor live. 
* `linked_compartment_name` - The name assigned to the compartment which was created or linked by customer with resource anchor. This compartment is different from where resource Anchor live.
* `region` - Oracle Cloud Infrastructure Region that resource is created.
* `setup_mode` - AUTO_BIND - when passed compartment will be created on-behalf of customer and bind to this resource anchor NO_AUTO_BIND - compartment will not be created and later customer can bind existing compartment.  to this resource anchor. This is for future use only 
* `lifecycle_state` - The current state of the ResourceAnchor.
* `resource_anchor_subscription_id` - Oracle Cloud Infrastructure Subscription Id
* `subscription_type` - subscription type
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the ResourceAnchor was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ResourceAnchor was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`