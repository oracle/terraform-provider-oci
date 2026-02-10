---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_generative_ai_private_endpoint"
sidebar_current: "docs-oci-datasource-generative_ai-generative_ai_private_endpoint"
description: |-
  Provides details about a specific Generative Ai Private Endpoint in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_generative_ai_private_endpoint
This data source provides details about a specific Generative Ai Private Endpoint resource in Oracle Cloud Infrastructure Generative AI service.

Retrieves an Generative AI private endpoint using a `privateEndpointId`.


## Example Usage

```hcl
data "oci_generative_ai_generative_ai_private_endpoint" "test_generative_ai_private_endpoint" {
	#Required
	generative_ai_private_endpoint_id = oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `generative_ai_private_endpoint_id` - (Required) The unique id for a Generative AI private endpoint. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the private endpoint. 
* `display_name` - A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `id` - The OCID of a private endpoint.  
* `private_endpoint_ip` - The private IP address (in the customer's VCN) that represents the access point for the associated endpoint service. 
* `subnet_id` - The OCID of the subnet that the private endpoint belongs to. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the Generative AI private endpoint was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time that the Generative AI private endpoint was updated expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
* `description` - A description of this private endpoint.
* `fqdn` - Fully qualified domain name the customer will use for access (for eg: xyz.oraclecloud.com) 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `nsg_ids` - A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to.
* `state` - The current state of the Generative AI Private Endpoint. 

