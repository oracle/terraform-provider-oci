---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_resource_analytics_instance_oac_management"
sidebar_current: "docs-oci-resource-resource_analytics-resource_analytics_instance_oac_management"
description: |-
  Provides the Resource Analytics Instance Oac Management resource in Oracle Cloud Infrastructure Resource Analytics service
---

# oci_resource_analytics_resource_analytics_instance_oac_management
This resource provides the Resource Analytics Instance Oac Management resource in Oracle Cloud Infrastructure Resource Analytics service.

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/
Attaches an OAC instance to a ResourceAnalyticsInstance.


## Example Usage

```hcl
resource "oci_resource_analytics_resource_analytics_instance_oac_management" "test_resource_analytics_instance_oac_management" {
	#Required
	attachment_type = var.resource_analytics_instance_oac_management_attachment_type
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
	enable_oac = var.enable_oac

	#Optional
	attachment_details {

		#Optional
		idcs_domain_id = oci_identity_domain.test_domain.id
		capacity_type = var.resource_analytics_instance_oac_management_attachment_details_capacity_type
		capacity_value = var.resource_analytics_instance_oac_management_attachment_details_capacity_value
		license_model = var.resource_analytics_instance_oac_management_attachment_details_license_model
		network_details {

			#Optional
			subnet_id = oci_core_subnet.test_subnet.id
			nsg_ids = var.resource_analytics_instance_oac_management_attachment_details_network_details_nsg_ids
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `attachment_type` - (Required) The type of attachment the OAC instance is using. Example: `MANAGED`
* `resource_analytics_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
* `enable_oac` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.
* `attachment_details` - (Optional) Additional details needed when attaching the OAC instance.  Example: `{"idcsDomainId":"ocid...","networkDetails":{...}, ...}` 
	* `idcs_domain_id` - (Required) IDCS domain [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying a stripe and service administrator user.
	* `capacity_type` - (Optional) The capacity model to use for the Analytics Instance.
	* `capacity_value` - (Optional) The capacity value selected, either the number of OCPUs (OLPU_COUNT) or the number of users (USER_COUNT). This parameter affects the number of OCPUs, amount of memory, and other resources allocated to the Analytics Instance.
	* `license_model` - (Optional) The Oracle license model that applies to the OAC instance.
	* `network_details` - (Optional) Details required when provisioning OAC on a private network.  Example: `{"subnetId":"ocid...", ...}` 
		* `nsg_ids` - (Optional) List of Network Security Group [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)'s for the private network details.  Example: `["ocid...", "ocid..."]` 
		* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet for the private network details.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `adw_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the created ADW instance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the ResourceAnalyticsInstance instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
* `lifecycle_details` - A message that describes the current state of the ResourceAnalyticsInstance in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `oac_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OAC enabled for the ResourceAnalyticsInstance.
* `state` - The current state of the ResourceAnalyticsInstance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the ResourceAnalyticsInstance was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ResourceAnalyticsInstance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Resource Analytics Instance Oac Management
* `update` - (Defaults to 20 minutes), when updating the Resource Analytics Instance Oac Management
* `delete` - (Defaults to 20 minutes), when destroying the Resource Analytics Instance Oac Management
