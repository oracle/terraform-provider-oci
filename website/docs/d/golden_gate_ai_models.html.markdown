---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_ai_models"
sidebar_current: "docs-oci-datasource-golden_gate-ai_models"
description: |-
  Provides the list of Ai Models in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_ai_models
This data source provides the list of Ai Models in Oracle Cloud Infrastructure Golden Gate service.

Returns the list of AI models for the specified provider. If the
provider requires additional context to resolve its supported models,
that context must be supplied in the request.


## Example Usage

```hcl
data "oci_golden_gate_ai_models" "test_ai_models" {
	#Required
	compartment_id = var.compartment_id
	provider_type = var.ai_model_provider_type

	#Optional
	region = var.ai_model_region
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `provider_type` - (Required) The AI provider type for which model information is requested. 
* `region` - (Optional) Oracle Cloud Infrastructure region identifier, for example us-ashburn-1. 
* `tenancy_id` - (Optional) Oracle Cloud Infrastructure tenancy OCID to use when resolving provider models, for example for Oracle Cloud Infrastructure Generative AI model discovery across a specific tenancy. 


## Attributes Reference

The following attributes are exported:

* `ai_model_collection` - The list of ai_model_collection.

### AiModel Reference

The following attributes are exported:

* `items` - An array of AI models. 
	* `description` - Metadata about this specific object. 
	* `display_name` - An object's Display Name. 
	* `key` - The identifier of the AI model offered by a provider. 
	* `provider_type` - AI Provider type used by the AI Model Connection. 

