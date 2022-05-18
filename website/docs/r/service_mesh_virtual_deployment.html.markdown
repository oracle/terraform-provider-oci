---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_deployment"
sidebar_current: "docs-oci-resource-service_mesh-virtual_deployment"
description: |-
  Provides the Virtual Deployment resource in Oracle Cloud Infrastructure Service Mesh service
---

# oci_service_mesh_virtual_deployment
This resource provides the Virtual Deployment resource in Oracle Cloud Infrastructure Service Mesh service.

Creates a new VirtualDeployment.


## Example Usage

```hcl
resource "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
	#Required
	compartment_id = var.compartment_id
	listeners {
		#Required
		port = var.virtual_deployment_listeners_port
		protocol = var.virtual_deployment_listeners_protocol
	}
	name = var.virtual_deployment_name
	service_discovery {
		#Required
		hostname = var.virtual_deployment_service_discovery_hostname
		type = var.virtual_deployment_service_discovery_type
	}
	virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

	#Optional
	access_logging {

		#Optional
		is_enabled = var.virtual_deployment_access_logging_is_enabled
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.virtual_deployment_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `access_logging` - (Optional) (Updatable) This configuration determines if logging is enabled and where the logs will be output.
	* `is_enabled` - (Optional) (Updatable) Determines if the logging configuration is enabled.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `listeners` - (Required) (Updatable) The listeners for the virtual deployment.
	* `port` - (Required) (Updatable) Port in which virtual deployment is running.
	* `protocol` - (Required) (Updatable) Type of protocol used in virtual deployment.
* `name` - (Required) A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `service_discovery` - (Required) (Updatable) Service Discovery configuration for virtual deployments.
	* `hostname` - (Required) (Updatable) The hostname of the virtual deployments.
	* `type` - (Required) (Updatable) Type of service discovery.
* `virtual_service_id` - (Required) The OCID of the service mesh in which this access policy is created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_logging` - This configuration determines if logging is enabled and where the logs will be output.
	* `is_enabled` - Determines if the logging configuration is enabled.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: `This is my new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `listeners` - The listeners for the virtual deployment
	* `port` - Port in which virtual deployment is running.
	* `protocol` - Type of protocol used in virtual deployment.
* `name` - A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `service_discovery` - Service Discovery configuration for virtual deployments.
	* `hostname` - The hostname of the virtual deployments.
	* `type` - Type of service discovery.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.
* `virtual_service_id` - The OCID of the virtual service in which this virtual deployment is created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Virtual Deployment
	* `update` - (Defaults to 20 minutes), when updating the Virtual Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Virtual Deployment


## Import

VirtualDeployments can be imported using the `id`, e.g.

```
$ terraform import oci_service_mesh_virtual_deployment.test_virtual_deployment "id"
```

