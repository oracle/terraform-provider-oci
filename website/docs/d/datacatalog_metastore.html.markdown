---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_metastore"
sidebar_current: "docs-oci-datasource-datacatalog-metastore"
description: |-
  Provides details about a specific Metastore in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_metastore
This data source provides details about a specific Metastore resource in Oracle Cloud Infrastructure Data Catalog service.

Gets a metastore by identifier.

## Example Usage

```hcl
data "oci_datacatalog_metastore" "test_metastore" {
	#Required
	metastore_id = oci_datacatalog_metastore.test_metastore.id
}
```

## Argument Reference

The following arguments are supported:

* `metastore_id` - (Required) The metastore's OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment which holds the metastore.
* `default_external_table_location` - Location under which external tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `default_managed_table_location` - Location under which managed tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the metastore.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The metastore's OCID.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `state` - The current state of the metastore.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time at which the metastore was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - Time at which the metastore was last modified. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

