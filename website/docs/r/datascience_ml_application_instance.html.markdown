---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_application_instance"
sidebar_current: "docs-oci-resource-datascience-ml_application_instance"
description: |-
  Provides the Ml Application Instance resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_ml_application_instance
This resource provides the Ml Application Instance resource in Oracle Cloud Infrastructure Data Science service.

Creates a new MlApplicationInstance.


## Example Usage

```hcl
resource "oci_datascience_ml_application_instance" "test_ml_application_instance" {
	#Required
	compartment_id = var.compartment_id
	ml_application_id = oci_datascience_ml_application.test_ml_application.id
	ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id

	#Optional
	auth_configuration {
		#Required
		type = var.ml_application_instance_auth_configuration_type

		#Optional
		access_token = var.ml_application_instance_auth_configuration_access_token
		application_name = oci_dataflow_application.test_application.name
		audience = var.ml_application_instance_auth_configuration_audience
		domain_id = oci_identity_domain.test_domain.id
		role_name = var.ml_application_instance_auth_configuration_role_name
		scope = var.ml_application_instance_auth_configuration_scope
	}
	configuration {
		#Required
		key = var.ml_application_instance_configuration_key

		#Optional
		value = var.ml_application_instance_configuration_value
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.ml_application_instance_display_name
	freeform_tags = {"Department"= "Finance"}
	is_enabled = var.ml_application_instance_is_enabled
}
```

## Argument Reference

The following arguments are supported:

* `auth_configuration` - (Optional) AuthN/Z configuration for online prediction
	* `application_name` - (Required when type=IDCS) Name of the IDCS application
	* `domain_id` - (Required when type=IDCS) Identity Domain OCID
	* `type` - (Required) Type of AuthN/Z
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the MlApplicationInstance is created.
* `configuration` - (Optional) (Updatable) Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplicationImplementation.
	* `key` - (Required) (Updatable) Key of configuration property
	* `value` - (Optional) (Updatable) Value of configuration property
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) The name of MlApplicationInstance. System will generate displayName when not provided.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_enabled` - (Optional) (Updatable) Defines whether the MlApplicationInstance will be created in ACTIVE (true value) or INACTIVE (false value) lifecycle state.
* `ml_application_id` - (Required) The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
* `ml_application_implementation_id` - (Required) (Updatable) The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ml Application Instance
	* `update` - (Defaults to 20 minutes), when updating the Ml Application Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Ml Application Instance


## Import

MlApplicationInstances can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_ml_application_instance.test_ml_application_instance "id"
```

