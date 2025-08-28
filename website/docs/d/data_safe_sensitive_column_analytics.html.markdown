---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_column_analytics"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_column_analytics"
description: |-
  Provides the list of Sensitive Column Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_column_analytics
This data source provides the list of Sensitive Column Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets consolidated sensitive columns analytics data based on the specified query parameters.

When you perform the ListSensitiveColumnAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.

To use ListSensitiveColumnAnalytics to get a full list of all compartments and subcompartments in the tenancy from the root compartment,
set the parameter compartmentIdInSubtree to true and accessLevel to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_sensitive_column_analytics" "test_sensitive_column_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sensitive_column_analytic_access_level
	column_name = var.sensitive_column_analytic_column_name
	compartment_id_in_subtree = var.sensitive_column_analytic_compartment_id_in_subtree
	group_by = var.sensitive_column_analytic_group_by
	object = var.sensitive_column_analytic_object
	schema_name = var.sensitive_column_analytic_schema_name
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) The group by parameter to summarize the sensitive columns.
* `object` - (Optional) A filter to return only items related to a specific object name.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_data_model_id` - (Optional) A filter to return only the resources that match the specified sensitive data model OCID.
* `sensitive_type_group_id` - (Optional) An optional filter to return only resources that match the specified OCID of the sensitive type group resource.
* `sensitive_type_id` - (Optional) A filter to return only the sensitive columns that are associated with one of the sensitive types identified by the specified OCIDs.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `sensitive_column_analytics_collection` - The list of sensitive_column_analytics_collection.

### SensitiveColumnAnalytic Reference

The following attributes are exported:

* `items` - An array of sensitive column analytics summary objects.
	* `dimensions` - The dimensions available for sensitive column analytics.
		* `column_name` - The name of the sensitive column.
		* `object` - The database object that contains the sensitive column.
		* `schema_name` - The database schema that contains the sensitive column.
		* `sensitive_data_model_id` - The OCID of the sensitive data model which contains the sensitive column.
		* `sensitive_type_id` - The OCID of the sensitive type associated with the sensitive column.
		* `target_id` - The OCID of the target database associated with the sensitive column.
	* `sensitive_column_analytic_count` - The total count for the aggregation metric.

