---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_assigned_subscription"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-assigned_subscription"
description: |-
  Provides details about a specific Assigned Subscription in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_assigned_subscription
This data source provides details about a specific Assigned Subscription resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Get the assigned subscription details by assigned subscription ID.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_assigned_subscription" "test_assigned_subscription" {
	#Required
	assigned_subscription_id = oci_tenantmanagercontrolplane_assigned_subscription.test_assigned_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `assigned_subscription_id` - (Required) OCID of the assigned Oracle Cloud Subscription.


## Attributes Reference

The following attributes are exported:

* `classic_subscription_id` - Subscription ID.
* `cloud_amount_currency` - The currency code for the customer associated with the subscription.
* `compartment_id` - The Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the owning compartment. Always a tenancy OCID.
* `csi_number` - Customer service identifier for the customer associated with the subscription.
* `currency_code` - Currency code. For example USD, MXN.
* `customer_country_code` - The country code for the customer associated with the subscription.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `end_date` - Subscription end time.
* `entity_version` - The entity version of the subscription, whether V1 (the legacy schema version), or V2 (the latest 20230401 API version).
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the subscription.
* `is_classic_subscription` - Specifies whether or not the subscription is legacy.
* `is_government_subscription` - Specifies whether or not the subscription is a government subscription.
* `managed_by` - Service or component which is used to provision and manage the subscription.
* `order_ids` - List of subscription order OCIDs that contributed to this subscription.
* `program_type` - Specifies any program that is associated with the subscription.
* `promotion` - List of promotions related to the subscription.
	* `amount` - If a subscription is present, indicates the total amount of promotional subscription credits.
	* `currency_unit` - Currency unit associated with the promotion.
	* `duration` - Specifies how long the promotion related to the subscription, if any, is valid in duration units.
	* `duration_unit` - Unit for the duration.
	* `is_intent_to_pay` - Speficies whether or not the customer intends to pay after the promotion has expired.
	* `status` - If a subscription is present, indicates the current status of the subscription promotion.
	* `time_expired` - Date and time when the promotion ends.
	* `time_started` - Date and time when the promotion starts.
* `purchase_entitlement_id` - Purchase entitlement ID associated with the subscription.
* `region_assignment` - Region for the subscription.
* `service_name` - The type of subscription, such as 'UCM', 'SAAS', 'ERP', 'CRM'.
* `skus` - List of SKUs linked to the subscription.
	* `description` - Description of the stock units.
	* `end_date` - Date and time when the SKU ended.
	* `gsi_order_line_id` - Sales order line identifier.
	* `is_additional_instance` - Specifies if an additional test instance can be provisioned by the SaaS application.
	* `is_base_service_component` - Specifies if the SKU is considered as a parent or child.
	* `license_part_description` - Description of the covered product belonging to this SKU.
	* `metric_name` - Base metric for billing the service.
	* `quantity` - Quantity of the stock units.
	* `sku` - Stock Keeping Unit (SKU) ID.
	* `start_date` - Date and time when the SKU was created.
* `start_date` - Subscription start time.
* `state` - Lifecycle state of the subscription.
* `subscription_number` - Unique Oracle Cloud Subscriptions identifier that is immutable on creation.
* `subscription_tier` - Tier for the subscription, whether a free promotion subscription or a paid subscription.
* `time_created` - The date and time of creation, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `time_updated` - The date and time of update, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 

