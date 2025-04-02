---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_application_instances"
sidebar_current: "docs-oci-datasource-datascience-ml_application_instances"
description: |-
  Provides the list of Ml Application Instances in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_ml_application_instances
This data source provides the list of Ml Application Instances in Oracle Cloud Infrastructure Data Science service.

Returns a list of MlApplicationsInstances.


## Example Usage

```hcl
data "oci_datascience_ml_application_instances" "test_ml_application_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.ml_application_instance_display_name
	ml_application_id = oci_datascience_ml_application.test_ml_application.id
	state = var.ml_application_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `ml_application_id` - (Optional) unique MlApplication identifier
* `state` - (Optional) A filter to return only resources matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `ml_application_instance_collection` - The list of ml_application_instance_collection.

### MlApplicationInstance Reference

The following attributes are exported:

* `auth_configuration` - AuthN/Z configuration for online prediction
	* `application_name` - Name of the IDCS application
	* `domain_id` - Identity Domain OCID
	* `type` - Type of AuthN/Z
* `compartment_id` - The OCID of the compartment where the MlApplicationInstance is created.
* `configuration` - Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplicationImplementation.
	* `key` - Key of configuration property
	* `value` - Value of configuration property
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The name of MlApplicationInstance. System will generate displayName when not provided during creation.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the MlApplicationInstance. Unique identifier that is immutable after creation
* `is_enabled` - States whether the MlApplicationInstance is supposed to be in ACTIVE lifecycle state.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `lifecycle_substate` - The current substate of the MlApplicationInstance. The substate has MlApplicationInstance specific values in comparison with lifecycleState which has standard values common for all Oracle Cloud Infrastructure resources. The NEEDS_ATTENTION and FAILED substates are deprecated in favor of (NON_)?RECOVERABLE_(PROVIDER|SERVICE)_ISSUE and will be removed in next release. 
* `ml_application_id` - The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
* `ml_application_implementation_id` - The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)
* `ml_application_implementation_name` - The name of Ml Application Implementation (based on mlApplicationImplementationId)
* `ml_application_name` - The name of ML Application (based on mlApplicationId).
* `prediction_endpoint_details` - Prediction endpoint related information.
	* `base_prediction_uri` - Base URI of prediction router.
	* `prediction_uris` - Array of all prediction URIs per use-case.
		* `uri` - Prediction URI.
		* `use_case` - Prediction use-case.
* `state` - The current state of the MlApplicationInstance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the MlApplication was created. An RFC3339 formatted datetime string
* `time_updated` - Time of last MlApplicationInstance update in the format defined by RFC 3339.

