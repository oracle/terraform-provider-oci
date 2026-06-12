---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_application_storage"
sidebar_current: "docs-oci-datasource-generative_ai-hosted_application_storage"
description: |-
  Provides details about a specific Hosted Application Storage in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_hosted_application_storage
This data source provides details about a specific Hosted Application Storage resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about a hosted application storage.

## Example Usage

```hcl
data "oci_generative_ai_hosted_application_storage" "test_hosted_application_storage" {
	#Required
	hosted_application_storage_id = oci_generative_ai_hosted_application_storage.test_hosted_application_storage.id
}
```

## Argument Reference

The following arguments are supported:

* `hosted_application_storage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application storage.


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
