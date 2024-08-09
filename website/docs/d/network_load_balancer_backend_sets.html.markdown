---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_backend_sets"
sidebar_current: "docs-oci-datasource-network_load_balancer-backend_sets"
description: |-
  Provides the list of Backend Sets in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_backend_sets
This data source provides the list of Backend Sets in Oracle Cloud Infrastructure Network Load Balancer service.

Lists all backend sets associated with a given network load balancer.

## Example Usage

```hcl
data "oci_network_load_balancer_backend_sets" "test_backend_sets" {
	#Required
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `backend_set_collection` - The list of backend_set_collection.

### BackendSet Reference

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
    * `protocol` - The protocol the health check must use; either HTTP, HTTPS, UDP, TCP or DNS.  Example: `HTTP` 
    * `request_data` - Base64 encoded pattern to be sent as UDP or TCP health check probe.
    * `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
    * `response_data` - Base64 encoded pattern to be validated as UDP or TCP health check probe response.
    * `retries` - The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3` 
    * `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200` 
    * `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000` 
    * `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `ip_version` - IP version associated with the backend set.
* `is_fail_open` - If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy. The value is false by default. 
* `is_instant_failover_enabled` - If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy. 
* `is_preserve_source` - If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default. 
* `name` - A user-friendly name for the backend set that must be unique and cannot be changed.

    Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

    Example: `example_backend_set` 
* `policy` - The network load balancer policy for the backend set.  Example: `FIVE_TUPLE` 

