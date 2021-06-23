---
subcategory: "Apm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_apm_domains"
sidebar_current: "docs-oci-datasource-apm-apm_domains"
description: |-
  Provides the list of Apm Domains in Oracle Cloud Infrastructure Apm service
---

# Data Source: oci_apm_apm_domains
This data source provides the list of Apm Domains in Oracle Cloud Infrastructure Apm service.

Lists all APM Domains for the specified tenant compartment.


## Example Usage

```hcl
data "oci_apm_apm_domains" "test_apm_domains" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.apm_domain_display_name
	state = var.apm_domain_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) A filter to return only resources that match the given life-cycle state.


## Attributes Reference

The following attributes are exported:

* `apm_domains` - The list of apm_domains.

### ApmDomain Reference

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

