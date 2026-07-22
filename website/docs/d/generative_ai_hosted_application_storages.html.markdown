---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_hosted_application_storages"
sidebar_current: "docs-oci-datasource-generative_ai-hosted_application_storages"
description: |-
  Provides the list of Hosted Application Storages in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_hosted_application_storages
This data source provides the list of Hosted Application Storages in Oracle Cloud Infrastructure Generative AI service.

Lists the hosted application storage in a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_hosted_application_storages" "test_hosted_application_storages" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.hosted_application_storage_display_name
	hosted_application_storage_type = var.hosted_application_storage_hosted_application_storage_type
	id = var.hosted_application_storage_id
	state = var.hosted_application_storage_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `hosted_application_storage_type` - (Optional) The type of the hosted application storage.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application storage.
* `state` - (Optional) A filter to return only the hosted applications that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `hosted_application_storage_collection` - The list of hosted_application_storage_collection.

### HostedApplicationStorage Reference

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
