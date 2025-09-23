---
subcategory: "Resource Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_analytics_resource_analytics_instance"
sidebar_current: "docs-oci-resource-resource_analytics-resource_analytics_instance"
description: |-
  Provides the Resource Analytics Instance resource in Oracle Cloud Infrastructure Resource Analytics service
---

# oci_resource_analytics_resource_analytics_instance
This resource provides the Resource Analytics Instance resource in Oracle Cloud Infrastructure Resource Analytics service.

Creates a ResourceAnalyticsInstance.


## Example Usage

```hcl
resource "oci_resource_analytics_resource_analytics_instance" "test_resource_analytics_instance" {
	#Required
	adw_admin_password {
		#Required
		password_type = var.resource_analytics_instance_adw_admin_password_password_type

		#Optional
		password = var.resource_analytics_instance_adw_admin_password_password
		secret_id = oci_vault_secret.test_secret.id
	}
	compartment_id = var.compartment_id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.resource_analytics_instance_description
	display_name = var.resource_analytics_instance_display_name
	freeform_tags = {"Department"= "Finance"}
	is_mutual_tls_required = var.resource_analytics_instance_is_mutual_tls_required
	license_model = var.resource_analytics_instance_license_model
	nsg_ids = var.resource_analytics_instance_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `adw_admin_password` - (Required) Details for the ADW Admin password. Password can be passed as `VaultSecretPasswordDetails` or `PlainTextPasswordDetails`.  Example: `{"passwordType":"PLAIN_TEXT","password":"..."}` Example: `{"passwordType":"VAULT_SECRET","secretId":"ocid..."}` 
	* `password` - (Required when password_type=PLAIN_TEXT) Password for the ADW to be created in User Tenancy. The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	* `password_type` - (Required) Password type
	* `secret_id` - (Required when password_type=VAULT_SECRET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault secret to use as the ADW admin password.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the ResourceAnalyticsInstance in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A description of the ResourceAnalyticsInstance instance.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_mutual_tls_required` - (Optional) Require mutual TLS (mTLS) when authenticating connections to the ADW database.
* `license_model` - (Optional) The Oracle license model that applies to the ADW instance.
* `nsg_ids` - (Optional) List of Network Security Group [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)'s.  Example: `["ocid...", "ocid..."]` 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.


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
	* `create` - (Defaults to 20 minutes), when creating the Resource Analytics Instance
	* `update` - (Defaults to 20 minutes), when updating the Resource Analytics Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Resource Analytics Instance


## Import

ResourceAnalyticsInstances can be imported using the `id`, e.g.

```
$ terraform import oci_resource_analytics_resource_analytics_instance.test_resource_analytics_instance "id"
```

