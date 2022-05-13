---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_access_policies"
sidebar_current: "docs-oci-datasource-service_mesh-access_policies"
description: |-
  Provides the list of Access Policies in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_access_policies
This data source provides the list of Access Policies in Oracle Cloud Infrastructure Service Mesh service.

Returns a list of AccessPolicy objects.


## Example Usage

```hcl
data "oci_service_mesh_access_policies" "test_access_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.access_policy_id
	mesh_id = oci_service_mesh_mesh.test_mesh.id
	name = var.access_policy_name
	state = var.access_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `id` - (Optional) Unique AccessPolicy identifier.
* `mesh_id` - (Optional) Unique Mesh identifier.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources that match the life cycle state given.


## Attributes Reference

The following attributes are exported:

* `access_policy_collection` - The list of access_policy_collection.

### AccessPolicy Reference

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

