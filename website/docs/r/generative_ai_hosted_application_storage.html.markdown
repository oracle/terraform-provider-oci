---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_application_storage"
sidebar_current: "docs-oci-resource-generative_ai-hosted_application_storage"
description: |-
  Provides the Hosted Application Storage resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_hosted_application_storage
This resource provides the Hosted Application Storage resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/HostedApplicationStorage

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a hosted application storage.

## Example Usage

```hcl
resource "oci_generative_ai_hosted_application_storage" "test_hosted_application_storage" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.hosted_application_storage_display_name
	storage_type = var.hosted_application_storage_storage_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.hosted_application_storage_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment OCID to create the hosted application in.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) An optional description of the hosted application.
* `display_name` - (Required) A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `storage_type` - (Required) type like Cache, Postgresql and ADB.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_ids` - A list of application OCID.
* `compartment_id` - The compartment OCID to create the hosted application in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - An optional description of the hosted application storage.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application storage.
* `lifecycle_details` - A message describing the current state of the hosted application storage in more detail that can provide actionable information.
* `state` - The current state of the hosted application storage.
* `storage_type` - type like Cache, Postgresql and ADB.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time the hosted application was created, in the format defined by RFC 3339
* `time_updated` - The date and time the hosted application was updated, in the format defined by RFC 3339

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Hosted Application Storage
	* `update` - (Defaults to 20 minutes), when updating the Hosted Application Storage
	* `delete` - (Defaults to 20 minutes), when destroying the Hosted Application Storage


## Import

HostedApplicationStorages can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_hosted_application_storage.test_hosted_application_storage "id"
```
