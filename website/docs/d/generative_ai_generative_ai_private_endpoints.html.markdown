---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_generative_ai_private_endpoints"
sidebar_current: "docs-oci-datasource-generative_ai-generative_ai_private_endpoints"
description: |-
  Provides the list of Generative Ai Private Endpoints in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_generative_ai_private_endpoints
This data source provides the list of Generative Ai Private Endpoints in Oracle Cloud Infrastructure Generative AI service.

Lists all Generative AI private endpoints in the specified compartment.


## Example Usage

```hcl
data "oci_generative_ai_generative_ai_private_endpoints" "test_generative_ai_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.generative_ai_private_endpoint_display_name
	id = var.generative_ai_private_endpoint_id
	state = var.generative_ai_private_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint. 
* `state` - (Optional) The lifecycle state of Generative AI private endpoints. 


## Attributes Reference

The following attributes are exported:

* `generative_ai_private_endpoint_collection` - The list of generative_ai_private_endpoint_collection.

### GenerativeAiPrivateEndpoint Reference

The following attributes are exported:

* `description` - A description of this private endpoint. 
* `display_name` - A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `fqdn` - Fully qualified domain name the customer will use for access (for eg: xyz.oraclecloud.com) 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `nsg_ids` - A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to. 
* `private_endpoint_ip` - The private IP address (in the customer's VCN) that represents the access point for the associated endpoint service. 
* `state` - The current state of the Generative AI Private Endpoint. 

