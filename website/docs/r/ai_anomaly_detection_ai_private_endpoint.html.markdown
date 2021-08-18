---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_ai_private_endpoint"
sidebar_current: "docs-oci-resource-ai_anomaly_detection-ai_private_endpoint"
description: |-
  Provides the Ai Private Endpoint resource in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# oci_ai_anomaly_detection_ai_private_endpoint
This resource provides the Ai Private Endpoint resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Create a new private reverse connection endpoint.

## Example Usage

```hcl
resource "oci_ai_anomaly_detection_ai_private_endpoint" "test_ai_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	dns_zones = var.ai_private_endpoint_dns_zones
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.ai_private_endpoint_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Display name of the private endpoint resource being created.
* `dns_zones` - (Required) (Updatable) List of DNS zones to be used by the data assets. Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `subnet_id` - (Required) The OCID of subnet to which the reverse connection is to be created. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ai Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Ai Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Ai Private Endpoint


## Import

AiPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_ai_anomaly_detection_ai_private_endpoint.test_ai_private_endpoint "id"
```

