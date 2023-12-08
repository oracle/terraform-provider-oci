---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalogs"
sidebar_current: "docs-oci-datasource-datacatalog-catalogs"
description: |-
  Provides the list of Catalogs in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalogs
This data source provides the list of Catalogs in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of all the data catalogs in the specified compartment.


## Example Usage

```hcl
data "oci_datacatalog_catalogs" "test_catalogs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.catalog_display_name
	state = var.catalog_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment where you want to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.


## Attributes Reference

The following attributes are exported:

* `catalogs` - The list of catalogs.

### Catalog Reference

The following attributes are exported:

* `attached_catalog_private_endpoints` - The list of private reverse connection endpoints attached to the catalog
* `compartment_id` - Compartment identifier.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Data catalog identifier, which can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - An message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in 'Failed' state. 
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `number_of_objects` - The number of data objects added to the data catalog. Please see the data catalog documentation for further information on how this is calculated. 
* `service_api_url` - The REST front endpoint URL to the data catalog instance.
* `service_console_url` - The console front endpoint URL to the data catalog instance.
* `state` - The current state of the data catalog resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the data catalog was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the data catalog was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

