---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_assessor"
sidebar_current: "docs-oci-datasource-database_migration-assessment_assessor"
description: |-
  Provides details about a specific Assessment Assessor in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessment_assessor
This data source provides details about a specific Assessment Assessor resource in Oracle Cloud Infrastructure Database Migration service.

Display Assessor details.

## Example Usage

```hcl
data "oci_database_migration_assessment_assessor" "test_assessment_assessor" {
	#Required
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessor_name = var.assessment_assessor_assessor_name
}
```

## Argument Reference

The following arguments are supported:

* `assessment_id` - (Required) The OCID of the Assessment 
* `assessor_name` - (Required) The name of the Assessor


## Attributes Reference

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

