---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_report_subsetted_objects"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_report_subsetted_objects"
description: |-
  Provides the list of Subsetting Report Subsetted Objects in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_report_subsetted_objects
This data source provides the list of Subsetting Report Subsetted Objects in Oracle Cloud Infrastructure Data Safe service.

Gets a list of subsetted tables present in the specified subsetting report and based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_subsetting_report_subsetted_objects" "test_subsetting_report_subsetted_objects" {
	#Required
	subsetting_report_id = oci_data_safe_subsetting_report.test_subsetting_report.id

	#Optional
	object = var.subsetting_report_subsetted_object_object
	schema_name = var.subsetting_report_subsetted_object_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_report_id` - (Required) The OCID of the subsetting report.


## Attributes Reference

The following attributes are exported:

* `subsetted_object_collection` - The list of subsetted_object_collection.

### SubsettingReportSubsettedObject Reference

The following attributes are exported:

* `items` - An array of subsetted summary objects
	* `object` - The name of the object (table or editioning view) subsetted
	* `object_type` - The type of the object (table or editioning view) subsetted
	* `row_count_after_subsetting` - The count of rows in the subsetted table after subsetting
	* `row_count_before_subsetting` - The count of rows in the subsetted table before subsetting
	* `schema_name` - The name of the schema that contains the subsetted object
	* `size_after_subsetting_in_kbs` - The size of the subsetted table after subsetting in KBs
	* `size_before_subsetting_in_kbs` - The size of the subsetted table before subsetting in KBs

