---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_gateway"
sidebar_current: "docs-oci-resource-apigateway-gateway"
description: |-
  Provides the Gateway resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_gateway
This resource provides the Gateway resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new gateway.


## Example Usage

```hcl
resource "oci_apigateway_gateway" "test_gateway" {
	#Required
	compartment_id = var.compartment_id
	endpoint_type = var.gateway_endpoint_type
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	certificate_id = "${oci_apigateway_certificate.test_certificate.id}"
	ca_bundles {
		#Required
		type = var.gateway_ca_bundles_type

		#Optional
		ca_bundle_id = oci_apigateway_ca_bundle.test_ca_bundle.id
		certificate_authority_id = oci_apigateway_certificate_authority.test_certificate_authority.id
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.gateway_display_name
	freeform_tags = {"Department"= "Finance"}
	network_security_group_ids = var.gateway_network_security_group_ids
	response_cache_details {
		#Required
		type = var.gateway_response_cache_details_type

		#Optional
		authentication_secret_id = oci_vault_secret.test_secret.id
		authentication_secret_version_number = var.gateway_response_cache_details_authentication_secret_version_number
		connect_timeout_in_ms = var.gateway_response_cache_details_connect_timeout_in_ms
		is_ssl_enabled = var.gateway_response_cache_details_is_ssl_enabled
		is_ssl_verify_disabled = var.gateway_response_cache_details_is_ssl_verify_disabled
		read_timeout_in_ms = var.gateway_response_cache_details_read_timeout_in_ms
		send_timeout_in_ms = var.gateway_response_cache_details_send_timeout_in_ms
		servers {

			#Optional
			host = var.gateway_response_cache_details_servers_host
			port = var.gateway_response_cache_details_servers_port
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `ca_bundles` - (Optional) (Updatable) An array of CA bundles that should be used on the Gateway for TLS validation.
	* `ca_bundle_id` - (Applicable when type=CA_BUNDLE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `certificate_authority_id` - (Applicable when type=CERTIFICATE_AUTHORITY) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `type` - (Required) (Updatable) Type of the CA bundle
* `certificate_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint_type` - (Required) Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be accessible on a private IP address on the subnet.  Example: `PUBLIC` or `PRIVATE` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `network_security_group_ids` - (Optional) (Updatable) An array of Network Security Groups OCIDs associated with this API Gateway. 
* `response_cache_details` - (Optional) (Updatable) Base Gateway response cache. 
	* `authentication_secret_id` - (Required when type=EXTERNAL_RESP_CACHE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
	* `authentication_secret_version_number` - (Required when type=EXTERNAL_RESP_CACHE) (Updatable) The version number of the authentication secret to use. 
	* `connect_timeout_in_ms` - (Applicable when type=EXTERNAL_RESP_CACHE) (Updatable) Defines the timeout for establishing a connection with the Response Cache. 
	* `is_ssl_enabled` - (Applicable when type=EXTERNAL_RESP_CACHE) (Updatable) Defines if the connection should be over SSL. 
	* `is_ssl_verify_disabled` - (Applicable when type=EXTERNAL_RESP_CACHE) (Updatable) Defines whether or not to uphold SSL verification. 
	* `read_timeout_in_ms` - (Applicable when type=EXTERNAL_RESP_CACHE) (Updatable) Defines the timeout for reading data from the Response Cache. 
	* `send_timeout_in_ms` - (Applicable when type=EXTERNAL_RESP_CACHE) (Updatable) Defines the timeout for transmitting data to the Response Cache. 
	* `servers` - (Required when type=EXTERNAL_RESP_CACHE) (Updatable) The set of cache store members to connect to. At present only a single server is supported. 
		* `host` - (Required when type=EXTERNAL_RESP_CACHE) (Updatable) Hostname or IP address (IPv4 only) where the cache store is running.
		* `port` - (Required when type=EXTERNAL_RESP_CACHE) (Updatable) The port the cache store is exposed on.
	* `type` - (Required) (Updatable) Type of the Response Cache.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ca_bundles` - An array of CA bundles that should be used on the Gateway for TLS validation.
	* `ca_bundle_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `certificate_authority_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
	* `type` - Type of the CA bundle
* `certificate_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint_type` - Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be accessible on a private IP address on the subnet.  Example: `PUBLIC` or `PRIVATE` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for APIs deployed on the gateway.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `ip_addresses` - An array of IP addresses associated with the gateway.
	* `ip_address` - An IP address.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `network_security_group_ids` - An array of Network Security Groups OCIDs associated with this API Gateway. 
* `response_cache_details` - Base Gateway response cache. 
	* `authentication_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
	* `authentication_secret_version_number` - The version number of the authentication secret to use. 
	* `connect_timeout_in_ms` - Defines the timeout for establishing a connection with the Response Cache. 
	* `is_ssl_enabled` - Defines if the connection should be over SSL. 
	* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
	* `read_timeout_in_ms` - Defines the timeout for reading data from the Response Cache. 
	* `send_timeout_in_ms` - Defines the timeout for transmitting data to the Response Cache. 
	* `servers` - The set of cache store members to connect to. At present only a single server is supported. 
		* `host` - Hostname or IP address (IPv4 only) where the cache store is running.
		* `port` - The port the cache store is exposed on.
	* `type` - Type of the Response Cache.
* `state` - The current state of the gateway.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created. 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Gateway
	* `update` - (Defaults to 20 minutes), when updating the Gateway
	* `delete` - (Defaults to 20 minutes), when destroying the Gateway


## Import

Gateways can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_gateway.test_gateway "id"
```

