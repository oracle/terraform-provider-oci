---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancers_backend_sets_unified"
sidebar_current: "docs-oci-resource-network_load_balancer-network_load_balancers_backend_sets_unified"
description: |-
  Provides the Network Load Balancers Backend Sets Unified resource in Oracle Cloud Infrastructure Network Load Balancer service
---

# oci_network_load_balancer_network_load_balancers_backend_sets_unified
This resource provides the Network Load Balancers Backend Sets Unified resource in Oracle Cloud Infrastructure Network Load Balancer service.

Adds a backend set to a network load balancer.

## Example Usage

```hcl
resource "oci_network_load_balancer_network_load_balancers_backend_sets_unified" "test_network_load_balancers_backend_sets_unified" {
	#Required
	health_checker {
		#Required
		protocol = var.network_load_balancers_backend_sets_unified_health_checker_protocol

		#Optional
		dns {
			#Required
			domain_name = oci_identity_domain.test_domain.name

			#Optional
			query_class = var.network_load_balancers_backend_sets_unified_health_checker_dns_query_class
			query_type = var.network_load_balancers_backend_sets_unified_health_checker_dns_query_type
			rcodes = var.network_load_balancers_backend_sets_unified_health_checker_dns_rcodes
			transport_protocol = var.network_load_balancers_backend_sets_unified_health_checker_dns_transport_protocol
		}
		interval_in_millis = var.network_load_balancers_backend_sets_unified_health_checker_interval_in_millis
		port = var.network_load_balancers_backend_sets_unified_health_checker_port
		request_data = var.network_load_balancers_backend_sets_unified_health_checker_request_data
		response_body_regex = var.network_load_balancers_backend_sets_unified_health_checker_response_body_regex
		response_data = var.network_load_balancers_backend_sets_unified_health_checker_response_data
		retries = var.network_load_balancers_backend_sets_unified_health_checker_retries
		return_code = var.network_load_balancers_backend_sets_unified_health_checker_return_code
		timeout_in_millis = var.network_load_balancers_backend_sets_unified_health_checker_timeout_in_millis
		url_path = var.network_load_balancers_backend_sets_unified_health_checker_url_path
	}
	name = var.network_load_balancers_backend_sets_unified_name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
	policy = var.network_load_balancers_backend_sets_unified_policy

	#Optional
	backends {
		#Required
		port = var.network_load_balancers_backend_sets_unified_backends_port

		#Optional
		ip_address = var.network_load_balancers_backend_sets_unified_backends_ip_address
		is_backup = var.network_load_balancers_backend_sets_unified_backends_is_backup
		is_drain = var.network_load_balancers_backend_sets_unified_backends_is_drain
		is_offline = var.network_load_balancers_backend_sets_unified_backends_is_offline
		name = var.network_load_balancers_backend_sets_unified_backends_name
		target_id = oci_cloud_guard_target.test_target.id
		weight = var.network_load_balancers_backend_sets_unified_backends_weight
	}
	ip_version = var.network_load_balancers_backend_sets_unified_ip_version
	is_fail_open = var.network_load_balancers_backend_sets_unified_is_fail_open
	is_preserve_source = var.network_load_balancers_backend_sets_unified_is_preserve_source
}
```

## Argument Reference

The following arguments are supported:

* `backends` - (Optional) (Updatable) An array of backends to be associated with the backend set.
	* `ip_address` - (Optional) (Updatable) The IP address of the backend server.  Example: `10.0.0.3` 
	* `is_backup` - (Optional) (Updatable) Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
	* `is_drain` - (Optional) (Updatable) Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
	* `is_offline` - (Optional) (Updatable) Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
	* `name` - (Optional) (Updatable) A read-only field showing the IP address/OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
	* `port` - (Required) (Updatable) The communication port for the backend server.  Example: `8080` 
	* `target_id` - (Optional) (Updatable) The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
	* `weight` - (Optional) (Updatable) The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - (Required) (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm). 
	* `dns` - (Optional) (Updatable) DNS healthcheck configurations.
		* `domain_name` - (Required) (Updatable) The absolute fully-qualified domain name to perform periodic DNS queries. If not provided, an extra dot will be added at the end of a domain name during the query. 
		* `query_class` - (Optional) (Updatable) The class the dns health check query to use; either IN or CH.  Example: `IN` 
		* `query_type` - (Optional) (Updatable) The type the dns health check query to use; A, AAAA, TXT.  Example: `A` 
		* `rcodes` - (Optional) (Updatable) An array that represents accepetable RCODE values for DNS query response. Example: ["NOERROR", "NXDOMAIN"] 
		* `transport_protocol` - (Optional) (Updatable) DNS transport protocol; either UDP or TCP.  Example: `UDP` 
	* `interval_in_millis` - (Optional) (Updatable) The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: `10000` 
	* `port` - (Optional) (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080` 
	* `protocol` - (Required) (Updatable) The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.  Example: `HTTP` 
	* `request_data` - (Optional) (Updatable) Base64 encoded pattern to be sent as UDP or TCP health check probe.
	* `response_body_regex` - (Optional) (Updatable) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `response_data` - (Optional) (Updatable) Base64 encoded pattern to be validated as UDP or TCP health check probe response.
	* `retries` - (Optional) (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3` 
	* `return_code` - (Optional) (Updatable) The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - (Optional) (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000` 
	* `url_path` - (Optional) (Updatable) The path against which to run the health check.  Example: `/healthcheck` 
* `ip_version` - (Optional) (Updatable) IP version associated with the backend set.
* `is_fail_open` - (Optional) (Updatable) If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy. The value is false by default. 
* `is_preserve_source` - (Optional) (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default. 
* `name` - (Required) A user-friendly name for the backend set that must be unique and cannot be changed.

	Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

	Example: `example_backend_set` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
* `policy` - (Required) (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE`` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backends` - Array of backends. 
	* `ip_address` - The IP address of the backend server. Example: `10.0.0.3` 
	* `is_backup` - Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
	* `is_drain` - Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
	* `is_offline` - Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
	* `name` - A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
	* `port` - The communication port for the backend server.  Example: `8080` 
	* `target_id` - The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
	* `weight` - The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm). 
	* `dns` - DNS healthcheck configurations.
		* `domain_name` - The absolute fully-qualified domain name to perform periodic DNS queries. If not provided, an extra dot will be added at the end of a domain name during the query. 
		* `query_class` - The class the dns health check query to use; either IN or CH.  Example: `IN` 
		* `query_type` - The type the dns health check query to use; A, AAAA, TXT.  Example: `A` 
		* `rcodes` - An array that represents accepetable RCODE values for DNS query response. Example: ["NOERROR", "NXDOMAIN"] 
		* `transport_protocol` - DNS transport protocol; either UDP or TCP.  Example: `UDP` 
	* `interval_in_millis` - The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: `10000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.  Example: `HTTP` 
	* `request_data` - Base64 encoded pattern to be sent as UDP or TCP health check probe.
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `response_data` - Base64 encoded pattern to be validated as UDP or TCP health check probe response.
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3` 
	* `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `ip_version` - IP version associated with the backend set.
* `is_fail_open` - If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy. The value is false by default. 
* `is_preserve_source` - If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default. 
* `name` - A user-friendly name for the backend set that must be unique and cannot be changed.

	Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

	Example: `example_backend_set` 
* `policy` - The network load balancer policy for the backend set.  Example: `FIVE_TUPLE` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Load Balancers Backend Sets Unified
	* `update` - (Defaults to 20 minutes), when updating the Network Load Balancers Backend Sets Unified
	* `delete` - (Defaults to 20 minutes), when destroying the Network Load Balancers Backend Sets Unified


## Import

NetworkLoadBalancersBackendSetsUnified can be imported using the `id`, e.g.

```
$ terraform import oci_network_load_balancer_network_load_balancers_backend_sets_unified.test_network_load_balancers_backend_sets_unified "networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}" 
```

