---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_discovery_jobs_result"
sidebar_current: "docs-oci-resource-data_safe-discovery_jobs_result"
description: |-
  Provides the Discovery Jobs Result resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_discovery_jobs_result
This resource provides the Discovery Jobs Result resource in Oracle Cloud Infrastructure Data Safe service.



## Example Usage

```hcl
resource "oci_data_safe_discovery_jobs_result" "test_discovery_jobs_result" {
}
```

## Argument Reference

The following arguments are supported:



** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `app_defined_child_column_keys` - Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column.
* `app_name` - The name of the application. An application is an entity that is identified by a schema and stores sensitive information for that schema. Its value will be same as schemaName, if no value is passed.
* `column_name` - The name of the sensitive column.
* `data_type` - The data type of the sensitive column.
* `db_defined_child_column_keys` - Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column.
* `discovery_type` - The type of the discovery result. It can be one of the following three types: NEW: A new sensitive column in the target database that is not in the sensitive data model. DELETED: A column that is present in the sensitive data model but has been deleted from the target database. MODIFIED: A column that is present in the target database as well as the sensitive data model but some of its attributes have been modified. 
* `estimated_data_value_count` - The estimated number of data values the column has in the associated database.
* `is_result_applied` - Indicates if the discovery result has been processed. You can update this attribute using the PatchDiscoveryJobResults operation to track whether the discovery result has already been processed and applied to the sensitive data model. 
* `key` - The unique key that identifies the discovery result.
* `modified_attributes` - The attributes of a sensitive column that have been modified in the target database. It's populated only in the case of MODIFIED discovery results and shows the new values of the modified attributes. 
	* `app_defined_child_column_keys` - Unique keys identifying the columns that are application-level (non-dictionary) children of the sensitive column.
	* `db_defined_child_column_keys` - Unique keys identifying the columns that are database-level (dictionary-defined) children of the sensitive column.
* `object` - The database object that contains the sensitive column.
* `object_type` - The type of the database object that contains the sensitive column.
* `parent_column_keys` - Unique keys identifying the columns that are parents of the sensitive column. At present, it tracks a single parent only.
* `planned_action` - Specifies how to process the discovery result. It's set to NONE by default. Use the PatchDiscoveryJobResults operation to update this attribute. You can choose one of the following options: ACCEPT: To accept the discovery result and update the sensitive data model to reflect the changes. REJECT: To reject the discovery result so that it doesn't change the sensitive data model. INVALIDATE: To invalidate a newly discovered column. It adds the column to the sensitive data model but marks it as invalid. It helps track false positives and ensure that they aren't reported by future discovery jobs. After specifying the planned action, you can use the ApplyDiscoveryJobResults operation to automatically process the discovery results. 
* `relation_type` - The type of referential relationship the sensitive column has with its parent. NONE indicates that the sensitive column does not have a parent. DB_DEFINED indicates that the relationship is defined in the database dictionary. APP_DEFINED indicates that the relationship is defined at the application level and not in the database dictionary. 
* `sample_data_values` - Original data values collected for the sensitive column from the associated database. Sample data helps review the column and ensure that it actually contains sensitive data. Note that sample data is retrieved by a data discovery job only if the isSampleDataCollectionEnabled attribute is set to true. At present, only one data value is collected per sensitive column. 
* `schema_name` - The database schema that contains the sensitive column.
* `sensitive_columnkey` - The unique key that identifies the sensitive column represented by the discovery result.
* `sensitive_type_id` - The OCID of the sensitive type associated with the sensitive column.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Discovery Jobs Result
	* `update` - (Defaults to 20 minutes), when updating the Discovery Jobs Result
	* `delete` - (Defaults to 20 minutes), when destroying the Discovery Jobs Result


## Import

DiscoveryJobsResults can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_discovery_jobs_result.test_discovery_jobs_result "discoveryJobs/{discoveryJobId}/results/{resultKey}" 
```

