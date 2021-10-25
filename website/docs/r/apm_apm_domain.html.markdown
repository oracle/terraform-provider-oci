---
subcategory: "Apm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_apm_domain"
sidebar_current: "docs-oci-resource-apm-apm_domain"
description: |-
  Provides the Apm Domain resource in Oracle Cloud Infrastructure Apm service
---

# oci_apm_apm_domain
This resource provides the Apm Domain resource in Oracle Cloud Infrastructure Apm service.

Creates a new APM domain.


## Example Usage

```hcl
resource "oci_apm_apm_domain" "test_apm_domain" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.apm_domain_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.apm_domain_description
	freeform_tags = {"bar-key"= "value"}
	is_free_tier = var.apm_domain_is_free_tier
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment corresponding to the APM domain.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the APM domain.
* `display_name` - (Required) (Updatable) Display name of the APM domain.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_free_tier` - (Optional) Indicates whether this is an "Always Free" resource. The default value is false.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment corresponding to the APM domain.
* `data_upload_endpoint` - The endpoint where the APM agents upload their observations and metrics.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the APM domain.
* `display_name` - Display name of the APM domain, which can be updated.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `is_free_tier` - Indicates if this is an Always Free resource.
* `state` - The current lifecycle state of the APM domain.
* `time_created` - The time the APM domain was created, expressed in RFC 3339 timestamp format.
* `time_updated` - The time the APM domain was updated, expressed in RFC 3339 timestamp format.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Apm Domain
	* `update` - (Defaults to 20 minutes), when updating the Apm Domain
	* `delete` - (Defaults to 20 minutes), when destroying the Apm Domain


## Import

ApmDomains can be imported using the `id`, e.g.

```
$ terraform import oci_apm_apm_domain.test_apm_domain "id"
```

