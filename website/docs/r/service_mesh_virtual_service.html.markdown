---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_service"
sidebar_current: "docs-oci-resource-service_mesh-virtual_service"
description: |-
  Provides the Virtual Service resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_virtual_service
This resource provides the Virtual Service resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new VirtualService.


## Example Usage

```hcl
resource "oci_service_mesh_virtual_service" "test_virtual_service" {
	#Required
	compartment_id = var.compartment_id
	mesh_id = oci_service_mesh_mesh.test_mesh.id
	name = var.virtual_service_name

	#Optional
	default_routing_policy {
		#Required
		type = var.virtual_service_default_routing_policy_type
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.virtual_service_description
	freeform_tags = {"bar-key"= "value"}
	hosts = var.virtual_service_hosts
	mtls {
		#Required
		mode = var.virtual_service_mtls_mode

		#Optional
		maximum_validity = var.virtual_service_mtls_maximum_validity
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `default_routing_policy` - (Optional) (Updatable) Routing policy for the virtual service.
	* `type` - (Required) (Updatable) Type of the virtual service routing policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hosts` - (Optional) (Updatable) The DNS hostnames of the virtual service that is used by its callers. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com". Can be omitted if the virtual service will only have TCP virtual deployments. 
* `mesh_id` - (Required) The OCID of the service mesh in which this virtual service is created.
* `mtls` - (Optional) (Updatable) The mTLS authentication mode to use when receiving requests from other virtual services or ingress gateways within the mesh. 
	* `maximum_validity` - (Optional) (Updatable) The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days. 
	* `mode` - (Required) (Updatable) DISABLED: Connection is not tunneled. PERMISSIVE: Connection can be either plaintext or an mTLS tunnel. STRICT: Connection is an mTLS tunnel.  Clients without a valid certificate will be rejected. 
* `name` - (Required) A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `default_routing_policy` - Routing policy for the virtual service.
	* `type` - Type of the virtual service routing policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hosts` - The DNS hostnames of the virtual service that is used by its callers. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com". Can be omitted if the virtual service will only have TCP virtual deployments. 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `mesh_id` - The OCID of the service mesh in which this virtual service is created.
* `mtls` - Mutual TLS settings used when communicating with other virtual services or ingress gateways within the mesh. 
	* `certificate_id` - The OCID of the certificate resource that will be used for mTLS authentication with other virtual services in the mesh. 
	* `maximum_validity` - The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days. 
	* `mode` - DISABLED: Connection is not tunneled. PERMISSIVE: Connection can be either plaintext or an mTLS tunnel. STRICT: Connection is an mTLS tunnel.  Clients without a valid certificate will be rejected. 
* `name` - A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Virtual Service
	* `update` - (Defaults to 20 minutes), when updating the Virtual Service
	* `delete` - (Defaults to 20 minutes), when destroying the Virtual Service


## Import

VirtualServices can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_virtual_service.test_virtual_service "id"
```

