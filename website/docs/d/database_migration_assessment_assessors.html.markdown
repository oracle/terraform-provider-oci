---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_assessors"
sidebar_current: "docs-oci-datasource-database_migration-assessment_assessors"
description: |-
  Provides the list of Assessment Assessors in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessment_assessors
This data source provides the list of Assessment Assessors in Oracle Cloud Infrastructure Database Migration service.

List all Assessors.


## Example Usage

```hcl
data "oci_database_migration_assessment_assessors" "test_assessment_assessors" {
	#Required
	assessment_id = oci_database_migration_assessment.test_assessment.id

	#Optional
	display_name = var.assessment_assessor_display_name
	state = var.assessment_assessor_state
}
```

## Argument Reference

The following arguments are supported:

* `assessment_id` - (Required) The OCID of the Assessment 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `state` - (Optional) The lifecycle state of the Assessor.


## Attributes Reference

The following attributes are exported:

* `assessor_summary_collection` - The list of assessor_summary_collection.

### AssessmentAssessor Reference

The following attributes are exported:

* `actions` - Assessor actions.
	* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `is_disabled` - Defines if the action is enabled or disabled.
	* `name` - The Assessor Action Name.
	* `resource_id` - The OCID of the resource being referenced.
	* `title` - The Assessor Action Title.
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
* `assessment_id` - The OCID of the resource being referenced.
* `assessor_group` - Assessor Group
	* `actions` - Assessor group actions.
		* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
		* `is_disabled` - Defines if the action is enabled or disabled.
		* `name` - The Assessor Action Name.
		* `resource_id` - The OCID of the resource being referenced.
		* `title` - The Assessor Action Title.
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
	* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
	* `name` - Assessor Group name, e.g. ADVANCED.
* `assessor_result` - The Assessor Result text.
* `checks_summary` - The Summary of all Checks.
* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information. 
* `does_script_require_restart` - True if DB restart required after running the script, false otherwise.
* `has_script` - True if script is available either from 'script' property of through download, false otherwise.
* `help_link_text` - The Help link text.
* `help_link_url` - The Help URL.
* `name` - The Assessor Name.
* `script` - The generated SQL script. Can be empty if the script exceeds maxLength. In this case the property 'hasScript' indicates that the script is available for download. 
* `state` - The current state of the Assessor.

