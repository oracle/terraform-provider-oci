---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_ai_private_endpoints"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-ai_private_endpoints"
description: |-
  Provides the list of Ai Private Endpoints in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_ai_private_endpoints
This data source provides the list of Ai Private Endpoints in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Returns a list of all the AI private endpoints in the specified compartment.


## Example Usage

```hcl
data "oci_ai_anomaly_detection_ai_private_endpoints" "test_ai_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.ai_private_endpoint_display_name
	id = var.ai_private_endpoint_id
	state = var.ai_private_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique AiPrivateEndpoint identifier
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `ai_private_endpoint_collection` - The list of ai_private_endpoint_collection.

### AiPrivateEndpoint Reference

The following attributes are exported:

* `attached_data_assets` - The list of dataAssets using the private reverse connection endpoint.
* `compartment_id` - Compartment Identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Private Reverse Connection Endpoint display name.
* `dns_zones` - List of DNS zones to be used by the data assets. Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `state` - The current state of the private endpoint resource.
* `subnet_id` - Subnet Identifier
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the private endpoint was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the private endpoint was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

