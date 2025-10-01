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

Attaches an OAC instance to a ResourceAnalyticsInstance.


## Example Usage

```hcl
resource "oci_resource_analytics_resource_analytics_instance_oac_management" "test_resource_analytics_instance_oac_management" {
	#Required
	resource_analytics_instance_id = oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance.id
	enable_oac = var.enable_oac

	#Optional
	attachment_details {

		#Optional
		idcs_domain_id = oci_identity_domain.test_domain.id
		license_model = var.resource_analytics_instance_oac_management_attachment_details_license_model
		network_details {

			#Optional
			subnet_id = oci_core_subnet.test_subnet.id
			nsg_ids = var.resource_analytics_instance_oac_management_attachment_details_network_details_nsg_ids
		}
	}
	attachment_type = var.resource_analytics_instance_oac_management_attachment_type
}
```

## Argument Reference

The following arguments are supported:

* `attachment_details` - (Optional) Additional details needed when attaching the OAC instance.  Example: `{"idcsDomainId":"ocid...","networkDetails":{...}, ...}` 
	* `idcs_domain_id` - (Optional) IDCS domain [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying a stripe and service administrator user.
	* `license_model` - (Optional) The Oracle license model that applies to the OAC instance.
	* `network_details` - (Optional) Details required when provisioning OAC on a private network.  Example: `{"subnetId":"ocid...", ...}` 
		* `nsg_ids` - (Optional) List of Network Security Group [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)'s for the private network details.  Example: `["ocid...", "ocid..."]` 
		* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet for the private network details.
* `attachment_type` - (Optional) The type of attachment the OAC instance is using.
* `resource_analytics_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance.
* `enable_oac` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Resource Analytics Instance Oac Management
	* `update` - (Defaults to 20 minutes), when updating the Resource Analytics Instance Oac Management
	* `delete` - (Defaults to 20 minutes), when destroying the Resource Analytics Instance Oac Management
