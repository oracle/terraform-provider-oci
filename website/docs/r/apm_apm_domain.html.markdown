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

Creates a new APM Domain.


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

* `compartment_id` - (Required) (Updatable) The OCID of the compartment corresponding to the APM Domain.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the APM Domain
* `display_name` - (Required) (Updatable) Display name of the APM Domain
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_free_tier` - (Optional) Indicates whether this is an "Always Free" resource. The default value is false.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment corresponding to the APM Domain.
* `data_upload_endpoint` - Where APM Agents upload their observations and metrics.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the APM Domain.
* `display_name` - APM Domain display name, can be updated.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `is_free_tier` - Indicates if this is an Always Free resource.
* `state` - The current lifecycle state of the APM Domain.
* `time_created` - The time the the APM Domain was created. An RFC3339 formatted datetime string
* `time_updated` - The time the APM Domain was updated. An RFC3339 formatted datetime string

## Import

ApmDomains can be imported using the `id`, e.g.

```
$ terraform import oci_apm_apm_domain.test_apm_domain "id"
```

