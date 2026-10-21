---
subcategory: "Pss"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_pss_instance"
sidebar_current: "docs-oci-resource-pss-instance"
description: |-
  Provides the Instance resource in Oracle Cloud Infrastructure Pss service
---

# oci_pss_instance
This resource provides the Instance resource in Oracle Cloud Infrastructure Pss service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/pss

Creates an Instance.


## Example Usage

```hcl
resource "oci_pss_instance" "test_instance" {
	#Required
	admin_identity_domain_user_id = oci_identity_user.test_user.id
	compartment_id = var.compartment_ocid
	identity_domain_url = var.instance_identity_domain_url

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `admin_identity_domain_user_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the identity-domain user that initially administers the PSS instance.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Instance in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `identity_domain_url` - (Required) The URL of the identity domain associated with the instance.  Example: `https://idcs-1234567890.identity.oraclecloud.com` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `admin_identity_domain_user_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the identity-domain user that initially administers the PSS instance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Instance.
* `identity_domain_url` - The URL of the identity domain associated with the instance.  Example: `https://idcs-1234567890.identity.oraclecloud.com` 
* `identity_domain_web_auth_client_id` - The client ID of the identity-domain confidential application used for browser authentication.
* `identity_domain_web_auth_client_secret_vault_id` - The OCID of the Vault secret containing the identity-domain confidential application's client secret.
* `lifecycle_details` - A message that describes the current state of the Instance in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `public_ip` - The public IP address for the instance's API gateway.
* `service_endpoint` - The HTTPS endpoint for the instance's PSS API.  Example: `https://example.pss.oraclecloud.com/api/` 
* `state` - The current state of the Instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.example-key": "example-value"}` 
* `time_created` - The date and time the Instance was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Instance was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `ui_endpoint` - The HTTPS endpoint for the instance's PSS user interface.  Example: `https://example.pss.oraclecloud.com/ui/` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 60 minutes), when creating the Instance
	* `update` - (Defaults to 60 minutes), when updating the Instance
	* `delete` - (Defaults to 60 minutes), when destroying the Instance


## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import oci_pss_instance.test_instance "id"
```
