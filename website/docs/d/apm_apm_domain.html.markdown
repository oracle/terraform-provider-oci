---
subcategory: "Apm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_apm_domain"
sidebar_current: "docs-oci-datasource-apm-apm_domain"
description: |-
  Provides details about a specific Apm Domain in Oracle Cloud Infrastructure Apm service
---

# Data Source: oci_apm_apm_domain
This data source provides details about a specific Apm Domain resource in Oracle Cloud Infrastructure Apm service.

Gets the details of the APM domain specified by OCID.

## Example Usage

```hcl
data "oci_apm_apm_domain" "test_apm_domain" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The OCID of the APM domain


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

