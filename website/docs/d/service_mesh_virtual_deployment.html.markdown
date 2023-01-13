---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_virtual_deployment"
sidebar_current: "docs-oci-datasource-service_mesh-virtual_deployment"
description: |-
  Provides details about a specific Virtual Deployment in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_virtual_deployment
This data source provides details about a specific Virtual Deployment resource in Oracle Cloud Infrastructure Service Mesh service.

Gets a VirtualDeployment by identifier.

## Example Usage

```hcl
data "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
	#Required
	virtual_deployment_id = oci_service_mesh_virtual_deployment.test_virtual_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `virtual_deployment_id` - (Required) Unique VirtualDeployment identifier.


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
	* `idle_timeout_in_ms` - The maximum duration in milliseconds for which the request's stream may be idle. The value 0 (zero) indicates that the timeout is disabled.
	* `port` - Port in which virtual deployment is running.
	* `protocol` - Type of protocol used in virtual deployment.
	* `request_timeout_in_ms` - The maximum duration in milliseconds for the deployed service to respond to an incoming request through the listener.  If provided, the timeout value overrides the default timeout of 15 seconds for the HTTP/HTTP2 listeners, and disabled (no timeout) for the GRPC listeners. The value 0 (zero) indicates that the timeout is disabled.  The timeout cannot be configured for the TCP and TLS_PASSTHROUGH listeners.  For streaming responses from the deployed service, consider either keeping the timeout disabled or set a sufficiently high value. 
* `name` - A user-friendly name. The name must be unique within the same virtual service and cannot be changed after creation. Avoid entering confidential information.  Example: `My unique resource name` 
* `service_discovery` - Service Discovery configuration for virtual deployments.
	* `hostname` - The hostname of the virtual deployments.
	* `type` - Type of service discovery.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.
* `virtual_service_id` - The OCID of the virtual service in which this virtual deployment is created.

