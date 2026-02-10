---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_assessor_checks"
sidebar_current: "docs-oci-datasource-database_migration-assessment_assessor_checks"
description: |-
  Provides the list of Assessment Assessor Checks in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessment_assessor_checks
This data source provides the list of Assessment Assessor Checks in Oracle Cloud Infrastructure Database Migration service.

List Assessor Check Summaries.


## Example Usage

```hcl
data "oci_database_migration_assessment_assessor_checks" "test_assessment_assessor_checks" {
	#Required
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessor_name = var.assessment_assessor_check_assessor_name
	compartment_id = var.compartment_id

	#Optional
	display_name = var.assessment_assessor_check_display_name
}
```

## Argument Reference

The following arguments are supported:

* `assessment_id` - (Required) The OCID of the Assessment 
* `assessor_name` - (Required) The name of the Assessor
* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 


## Attributes Reference

The following attributes are exported:

* `assessor_check_summary_collection` - The list of assessor_check_summary_collection.

### AssessmentAssessorCheck Reference

The following attributes are exported:

* `action` - Fixing the issue. 
* `assessor_check_group` - Assessor Check Group
	* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `is_expanded` - True if the group is expanded, false otherwise.
	* `name` - Assessor Check Group name, e.g. ADVANCED.
* `assessor_check_state` - The current state of the Assessor Check.
* `check_action` - Assessor Check Action
	* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `name` - The Assessor Check Action Name.
	* `user_defined_properties` - User defined properties
		* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
		* `help_link_text` - The Help link text.
		* `help_link_url` - The Help URL.
		* `properties` - Array of user defined properties.
			* `default_value` - The default value of the property.
			* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
			* `is_required` - True if the property is required, false otherwise
			* `max_length` - Maximum length of the text
			* `min_length` - Minimum length of the text
			* `name` - The property name.
			* `options` - User defined property options.
				* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
				* `value` - The option value.
			* `type` - The type of the user defined property.
			* `value` - The value of the property.
* `columns` - Array of the column of the objects table. 
	* `display_name` - Display name of column. 
	* `key` - Id of column. 
* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `fixup_script_location` - The path to the fixup script for this check. 
* `help_link_text` - The Help link text.
* `help_link_url` - The Help URL.
* `impact` - Impact of the issue on data migration. 
* `is_exclusion_allowed` - If false, objects cannot be excluded from migration. 
* `issue` - Description of the issue. 
* `key` - Pre-Migration сheck id. 
* `log_location` - Details to access log file in the specified Object Storage bucket, if any. 
	* `bucket` - Name of the bucket containing the log file. 
	* `namespace` - Object Storage namespace. 
	* `object` - Log object name. 
* `metadata` - Metadata of object. 
	* `object_name_column` - The field that stores the name of the object. 
	* `object_type_column` - The field that stores the type of the object. 
	* `object_type_fixed` - The field that stores the fixed type of the object. 
	* `schema_owner_column` - The field that stores the owner of the object. 
* `name` - The Name of the Check.
* `object_count` - Number of database objects to migrate. 
* `objects_display_name` - The objects display name. 

