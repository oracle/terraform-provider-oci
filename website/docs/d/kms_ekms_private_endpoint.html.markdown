---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_ekms_private_endpoint"
sidebar_current: "docs-oci-datasource-kms-ekms_private_endpoint"
description: |-
  Provides details about a specific Ekms Private Endpoint in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_ekms_private_endpoint
This data source provides details about a specific Ekms Private Endpoint resource in Oracle Cloud Infrastructure Kms service.

Gets a specific EKMS private by identifier.

## Example Usage

```hcl
data "oci_kms_ekms_private_endpoint" "test_ekms_private_endpoint" {
	#Required
	ekms_private_endpoint_id = oci_kms_ekms_private_endpoint.test_ekms_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `ekms_private_endpoint_id` - (Required) Unique EKMS private endpoint identifier.


## Attributes Reference

The following attributes are exported:

* `ca_bundle` - CABundle to validate TLS certificate of the external key manager system in PEM format 
* `compartment_id` - Identifier of the compartment this EKMS private endpoint belongs to
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the EKMS private endpoint
* `external_key_manager_ip` - Private IP of the external key manager system to connect to from the EKMS private endpoint 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `port` - The port of the external key manager system
* `private_endpoint_ip` - The IP address in the customer's VCN for the EKMS private endpoint. This is taken from subnet
* `state` - The current state of the EKMS private endpoint resource.
* `subnet_id` - Subnet Identifier
* `time_created` - The time the EKMS private endpoint was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the EKMS private endpoint was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

