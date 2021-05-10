---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_metastores"
sidebar_current: "docs-oci-datasource-datacatalog-metastores"
description: |-
  Provides the list of Metastores in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_metastores
This data source provides the list of Metastores in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of all metastores in the specified compartment.


## Example Usage

```hcl
data "oci_datacatalog_metastores" "test_metastores" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.metastore_display_name
	state = var.metastore_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment where you want to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.


## Attributes Reference

The following attributes are exported:

* `metastores` - The list of metastores.

### Metastore Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment which holds the metastore.
* `default_external_table_location` - Location under which external tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `default_managed_table_location` - Location under which managed tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the metastore.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The metastore's OCID.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the metastore.
* `time_created` - Time at which the metastore was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - Time at which the metastore was last modified. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

