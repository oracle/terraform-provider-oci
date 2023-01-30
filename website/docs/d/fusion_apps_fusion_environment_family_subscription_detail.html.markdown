---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_family_subscription_detail"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_family_subscription_detail"
description: |-
  Provides details about a specific Fusion Environment Family Subscription Detail in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_family_subscription_detail
This data source provides details about a specific Fusion Environment Family Subscription Detail resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets the subscription details of an fusion environment family.

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_family_subscription_detail" "test_fusion_environment_family_subscription_detail" {
	#Required
	fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_family_id` - (Required) The unique identifier (OCID) of the FusionEnvironmentFamily.


## Attributes Reference

The following attributes are exported:

* `subscriptions` - List of subscriptions.
	* `classic_subscription_id` - Subscription id.
	* `id` - OCID of the subscription details for particular root compartment or tenancy.
	* `service_name` - The type of subscription, such as 'CLOUDCM'/'SAAS'/'CRM', etc.
	* `skus` - Stock keeping unit.
		* `description` - Description of the stock units.
		* `license_part_description` - Description of the covered product belonging to this Sku.
		* `metric_name` - Base metric for billing the service.
		* `quantity` - Quantity of the stock units.
		* `sku` - Stock keeping unit id.

