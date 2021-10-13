---
subcategory: "Service Manager Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_manager_proxy_service_environment"
sidebar_current: "docs-oci-datasource-service_manager_proxy-service_environment"
description: |-
  Provides details about a specific Service Environment in Oracle Cloud Infrastructure Service Manager Proxy service
---

# Data Source: oci_service_manager_proxy_service_environment
This data source provides details about a specific Service Environment resource in Oracle Cloud Infrastructure Service Manager Proxy service.

Gets details of the service environment specified by the serviceEnvironmentId.

## Example Usage

```hcl
data "oci_service_manager_proxy_service_environment" "test_service_environment" {
	#Required
	compartment_id = var.compartment_id
	service_environment_id = oci_service_manager_proxy_service_environment.test_service_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `service_environment_id` - (Required) The Id associated with the service environment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Id associated with the service.
* `console_url` - The URL for the console.
* `id` - Unqiue identifier for the entitlement related to the environment. 
* `service_definition` - Model for details associated with service
	* `display_name` - Display name of the service.
	* `short_display_name` - Short display name of the service.
	* `type` - The service definition type.
* `service_environment_endpoints` - Array of service environment end points.
	* `environment_type` - Service Environemnt EndPoint type.
	* `url` - Service Environemnt Instance EndPoint url.
* `status` - Status of the entitlement registration for the service.
* `subscription_id` - The subscription Id corresponding to the service environment Id. 

