---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_model_discoveries"
sidebar_current: "docs-oci-datasource-generative_ai-model_discoveries"
description: |-
  Provides the list of Model Discoveries in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_model_discoveries
This data source provides the list of Model Discoveries in Oracle Cloud Infrastructure Generative AI service.

Retrieves a list of models along with their capabilities, supported features, and deployment availability.
Results can be filtered by attributes such as region, realm, model identifier, supported inference APIs, serving modes, and access type.


## Example Usage

```hcl
data "oci_generative_ai_model_discoveries" "test_model_discoveries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	api_capability = var.model_discovery_api_capability
	capability = var.model_discovery_capability
	is_dedicated_retired = var.model_discovery_is_dedicated_retired
	is_deprecated = var.model_discovery_is_deprecated
	is_on_demand_retired = var.model_discovery_is_on_demand_retired
	model_access = var.model_discovery_model_access
	model_id = oci_generative_ai_model.test_model.id
	realm = var.model_discovery_realm
	region = var.model_discovery_region
	serving_mode = var.model_discovery_serving_mode
}
```

## Argument Reference

The following arguments are supported:

* `api_capability` - (Optional) Filter models that support any of the specified API capabilities.
* `capability` - (Optional) A filter to return only resources their capability matches the given capability.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `is_dedicated_retired` - (Optional) Filter models based on dedicated retirement status.
* `is_deprecated` - (Optional) If true, return only deprecated models; if false, exclude deprecated models.
* `is_on_demand_retired` - (Optional) Filter models based on on-demand retirement status.
* `model_access` - (Optional) Filter models by access type.
* `model_id` - (Optional) A filter to return only resources whose model identifier matches the given modelId.
* `realm` - (Optional) A filter to return only resources whose realm matches the given realm.
* `region` - (Optional) A filter to return only resources whose region matches the given region.
* `serving_mode` - (Optional) A filter to return only resources whose serving modes match the given servingModes.


## Attributes Reference

The following attributes are exported:

* `model_discovery_collection` - The list of model_discovery_collection.

### ModelDiscovery Reference

The following attributes are exported:

* `items` - The list of discovered models matching the search criteria.
	* `api_capability` - The specific API endpoints that this model supports. For example, a chat model may support OPENAI_V1_CHAT_COMPLETIONS and/or OPENAI_V1_RESPONSES. If empty, the model has not yet been annotated with API capabilities.
	* `availability` - The list of availability details for the model across different regions and deployment modes.
		* `realm` - The cloud realm in which the model exists.
		* `region` - The specific geographic region where the model is deployed.
		* `serving_modes` - The supported deployment modes for the model in this region (e.g., on-demand or dedicated).
		* `supported_replacements` - A list of model identifiers that are recommended as replacements after this model is retired.
		* `time_dedicated_retired` - The timestamp when dedicated deployments of the model will be fully retired.
		* `time_deprecated` - The timestamp when the model is marked as deprecated and is no longer recommended for use.
		* `time_on_demand_retired` - The timestamp when the model will no longer be available for on-demand (shared) usage.
	* `capabilities` - Describes what this model can be used for.
	* `modality_support` - The supported input-to-output modality transformations for this model. For example, a model can support TEXT to VIDEO or AUDIO to VIDEO.
		* `input` - The source modality accepted by the model.
		* `output` - The target modality produced by the model.
	* `model_access` - The access level required to use the model, which can be either HOSTED (the model is hosted by the provider and can be accessed via API calls) or PROXY (the model is not directly accessible and requires going through a proxy).
	* `model_id` - A unique identifier for the model
	* `parameters` - The list of configurable parameters supported by the model. For example, temperature and max_tokens for a text generation model.
		* `default_value` - The default value used when this parameter is not supplied.
		* `description` - A human-readable description of this parameter.
		* `maximum` - The maximum allowed value for this parameter. Applicable only when type is FLOAT or INTEGER.
		* `minimum` - The minimum allowed value for this parameter. Applicable only when type is FLOAT or INTEGER.
		* `name` - The name of this parameter.
		* `type` - The data type of the parameter (e.g., float, integer, string).
	* `vendor` - The vendor that offers the model.
