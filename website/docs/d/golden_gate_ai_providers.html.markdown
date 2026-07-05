---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_ai_providers"
sidebar_current: "docs-oci-datasource-golden_gate-ai_providers"
description: |-
  Provides the list of Ai Providers in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_ai_providers
This data source provides the list of Ai Providers in Oracle Cloud Infrastructure Golden Gate service.

Returns the list of AI providers along with their supported models.


## Example Usage

```hcl
data "oci_golden_gate_ai_providers" "test_ai_providers" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 


## Attributes Reference

The following attributes are exported:

* `ai_provider_collection` - The list of ai_provider_collection.

### AiProvider Reference

The following attributes are exported:

* `items` - An array of AI providers. 
	* `auth_type` - Authentication types supported by the AI provider. 
	* `default_base_url` - Default base URL for the AI provider. 
	* `description` - Metadata about this specific object. 
	* `display_name` - An object's Display Name. 
	* `models` - List of AI models supported by this provider, when available. This field is null when the provider's models can be retrieved only after supplying additional context. For example, OCI_GENERATIVE_AI model availability may vary by region. 
		* `description` - Metadata about this specific object. 
		* `display_name` - An object's Display Name. 
		* `key` - The identifier of the AI model offered by a provider. 
		* `provider_type` - AI Provider type used by the AI Model Connection. 
	* `provider_type` - AI Provider type used by the AI Model Connection. 

