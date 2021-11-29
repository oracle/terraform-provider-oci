---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_metastore"
sidebar_current: "docs-oci-resource-datacatalog-metastore"
description: |-
  Provides the Metastore resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_metastore
This resource provides the Metastore resource in Oracle Cloud Infrastructure Data Catalog service.

Creates a new metastore.


## Example Usage

```hcl
resource "oci_datacatalog_metastore" "test_metastore" {
	#Required
	compartment_id = var.compartment_id
	default_external_table_location = var.metastore_default_external_table_location
	default_managed_table_location = var.metastore_default_managed_table_location

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.metastore_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) OCID of the compartment which holds the metastore.
* `default_external_table_location` - (Required) Location under which external tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `default_managed_table_location` - (Required) Location under which managed tables will be created by default. This references Object Storage using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/ 
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Mutable name of the metastore.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `state` - The current state of the metastore.
* `time_created` - Time at which the metastore was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - Time at which the metastore was last modified. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Metastore
	* `update` - (Defaults to 20 minutes), when updating the Metastore
	* `delete` - (Defaults to 20 minutes), when destroying the Metastore


## Import

Metastores can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_metastore.test_metastore "id"
```

