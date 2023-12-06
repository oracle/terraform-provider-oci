---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_ai_private_endpoint"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-ai_private_endpoint"
description: |-
  Provides details about a specific Ai Private Endpoint in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_ai_private_endpoint
This data source provides details about a specific Ai Private Endpoint resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Gets a specific private reverse connection by identifier.

## Example Usage

```hcl
data "oci_ai_anomaly_detection_ai_private_endpoint" "test_ai_private_endpoint" {
	#Required
	ai_private_endpoint_id = oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `ai_private_endpoint_id` - (Required) Unique private reverse connection identifier.


## Attributes Reference

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

