---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_lookup"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_lookup"
description: |-
  Provides details about a specific Namespace Lookup in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_lookup
This data source provides details about a specific Namespace Lookup resource in Oracle Cloud Infrastructure Log Analytics service.

Gets detailed information about the lookup with the specified name.


## Example Usage

```hcl
data "oci_log_analytics_namespace_lookup" "test_namespace_lookup" {
	#Required
	lookup_name = var.namespace_lookup_lookup_name
	namespace = var.namespace_lookup_namespace
}
```

## Argument Reference

The following arguments are supported:

* `lookup_name` - (Required) The name of the lookup to operate on.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `active_edit_version` - The active edit version.
* `canonical_link` - The canonical link.
* `categories` - An array of categories assigned to this lookup. The isSystem flag denotes if each category assignment is user-created or Oracle-defined. 
    * `description` - The category description.
    * `display_name` - The category display name.
    * `is_system` - The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
    * `name` - The unique name that identifies the category.
    * `type` - The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The lookup description.
* `edit_version` - The edit version.
* `fields` - The lookup fields.
    * `common_field_name` - The common field name.
    * `default_match_value` - The default match value.
    * `display_name` - The field display name.
    * `is_common_field` - A flag indicating whether or not the lookup field is a common field. 
    * `match_operator` - The match operator.
    * `name` - The field name.
    * `position` - THe field position.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `is_built_in` - A flag indicating if the lookup is custom (user-defined) or built in. 
* `is_hidden` - A flag indicating if the lookup is hidden or not.  A hidden lookup will not be returned in list operations by default. 
* `lookup_display_name` - The lookup display name.
* `lookup_id` - The lookup OCID.
* `lookup_name` - The lookup name.
* `lookup_reference` - The lookup reference as an integer.
* `lookup_reference_string` - The lookup reference as a string.
* `referring_sources` - AutoLookups
    * `canonical_link` - The canonical link.
    * `total_count` - The total count.
* `status_summary` - StatusSummary
    * `chunks_processed` - The number of chunks processed.
    * `failure_details` - The failure details, if any.
    * `filename` - The filename.
    * `status` - The status.
    * `total_chunks` - The total number of chunks.
* `time_updated` - The last updated date.
* `type` - The lookup type. Valid values are Lookup, Dictionary or Module.

