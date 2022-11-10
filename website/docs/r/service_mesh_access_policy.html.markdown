---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_access_policy"
sidebar_current: "docs-oci-resource-service_mesh-access_policy"
description: |-
  Provides the Access Policy resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_access_policy
This resource provides the Access Policy resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new AccessPolicy.


## Example Usage

```hcl
resource "oci_service_mesh_access_policy" "test_access_policy" {
	#Required
	compartment_id = var.compartment_id
	mesh_id = oci_service_mesh_mesh.test_mesh.id
	name = var.access_policy_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.access_policy_description
	freeform_tags = {"bar-key"= "value"}
	rules {
		#Required
		action = var.access_policy_rules_action
		destination {
			#Required
			type = var.access_policy_rules_destination_type

			#Optional
			hostnames = var.access_policy_rules_destination_hostnames
			ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
			ip_addresses = var.access_policy_rules_destination_ip_addresses
			ports = var.access_policy_rules_destination_ports
			protocol = var.access_policy_rules_destination_protocol
			virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
		}
		source {
			#Required
			type = var.access_policy_rules_source_type

			#Optional
			hostnames = var.access_policy_rules_source_hostnames
			ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
			ip_addresses = var.access_policy_rules_source_ip_addresses
			ports = var.access_policy_rules_source_ports
			protocol = var.access_policy_rules_source_protocol
			virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `mesh_id` - (Required) The OCID of the service mesh in which this access policy is created.
* `name` - (Required) A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `rules` - (Optional) (Updatable) List of applicable rules
	* `action` - (Required) (Updatable) Action for the traffic between the source and the destination.
	* `destination` - (Required) (Updatable) Target of the access policy. This can either be the source or the destination of the traffic.
		* `hostnames` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com", "*". Hostname "*" can be used to allow all hosts. 
		* `ingress_gateway_id` - (Required when type=INGRESS_GATEWAY) (Updatable) The OCID of the ingress gateway resource.
		* `ip_addresses` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication. 
		* `ports` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) Ports exposed by an external service. If left empty all ports will be allowed.
		* `protocol` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) Protocol of the external service
		* `type` - (Required) (Updatable) Traffic type of the target.
		* `virtual_service_id` - (Required when type=VIRTUAL_SERVICE) (Updatable) The OCID of the virtual service resource.
	* `source` - (Required) (Updatable) Target of the access policy. This can either be the source or the destination of the traffic.
		* `hostnames` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com", "*". Hostname "*" can be used to allow all hosts. 
		* `ingress_gateway_id` - (Required when type=INGRESS_GATEWAY) (Updatable) The OCID of the ingress gateway resource.
		* `ip_addresses` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication. 
		* `ports` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) Ports exposed by an external service. If left empty all ports will be allowed.
		* `protocol` - (Applicable when type=EXTERNAL_SERVICE) (Updatable) Protocol of the external service
		* `type` - (Required) (Updatable) Traffic type of the target.
		* `virtual_service_id` - (Required when type=VIRTUAL_SERVICE) (Updatable) The OCID of the virtual service resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `mesh_id` - The OCID of the service mesh in which this access policy is created.
* `name` - A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `rules` - List of applicable rules.
	* `action` - Action for the traffic between the source and the destination.
	* `destination` - Target of the access policy. This can either be the source or the destination of the traffic.
		* `hostnames` - The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com", "*". Hostname "*" can be used to allow all hosts. 
		* `ingress_gateway_id` - The OCID of the ingress gateway resource.
		* `ip_addresses` - The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication. 
		* `ports` - Ports exposed by an external service. If left empty all ports will be allowed.
		* `protocol` - Protocol of the external service
		* `type` - Traffic type of the target.
		* `virtual_service_id` - The OCID of the virtual service resource.
	* `source` - Target of the access policy. This can either be the source or the destination of the traffic.
		* `hostnames` - The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", "*.example.com", "*.com", "*". Hostname "*" can be used to allow all hosts. 
		* `ingress_gateway_id` - The OCID of the ingress gateway resource.
		* `ip_addresses` - The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication. 
		* `ports` - Ports exposed by an external service. If left empty all ports will be allowed.
		* `protocol` - Protocol of the external service
		* `type` - Traffic type of the target.
		* `virtual_service_id` - The OCID of the virtual service resource.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Access Policy
	* `update` - (Defaults to 20 minutes), when updating the Access Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Access Policy


## Import

AccessPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_access_policy.test_access_policy "id"
```

