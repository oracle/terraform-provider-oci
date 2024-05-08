---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_saved_queries"
sidebar_current: "docs-oci-datasource-cloud_guard-saved_queries"
description: |-
  Provides the list of Saved Queries in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_saved_queries
This data source provides the list of Saved Queries in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of saved queries run in a tenancy.


## Example Usage

```hcl
data "oci_cloud_guard_saved_queries" "test_saved_queries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.saved_query_access_level
	compartment_id_in_subtree = var.saved_query_compartment_id_in_subtree
	display_name = var.saved_query_display_name
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `saved_query_collection` - The list of saved_query_collection.

### SavedQuery Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID of the saved query
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the saved query
* `display_name` - Display name of the saved query
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - OCID for the saved query
* `query` - The saved query expression
* `state` - The current lifecycle state of the resource
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the saved query was created. Format defined by RFC3339.
* `time_updated` - The date and time the saved query was updated. Format defined by RFC3339.

