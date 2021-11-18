---
subcategory: "Service Manager Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_manager_proxy_service_environments"
sidebar_current: "docs-oci-datasource-service_manager_proxy-service_environments"
description: |-
  Provides the list of Service Environments in Oracle Cloud Infrastructure Service Manager Proxy service
---

# Data Source: oci_service_manager_proxy_service_environments
This data source provides the list of Service Environments in Oracle Cloud Infrastructure Service Manager Proxy service.

List details of environments which the service is authorized to view.
This includes the service instance endpoints and service definition
details.


## Example Usage

```hcl
data "oci_service_manager_proxy_service_environments" "test_service_environments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.service_environment_display_name
	service_environment_id = oci_service_manager_proxy_service_environment.test_service_environment.id
	service_environment_type = var.service_environment_service_environment_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `display_name` - (Optional) The display name of the resource.
* `service_environment_id` - (Optional) The Id associated with the service environment.
* `service_environment_type` - (Optional) The service definition type of the environment.


## Attributes Reference

The following attributes are exported:

* `service_environment_collection` - The list of service_environment_collection.

### ServiceEnvironment Reference

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

