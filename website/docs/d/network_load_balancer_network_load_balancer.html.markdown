---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancer"
sidebar_current: "docs-oci-datasource-network_load_balancer-network_load_balancer"
description: |-
  Provides details about a specific Network Load Balancer in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_network_load_balancer
This data source provides details about a specific Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves network load balancer configuration information by identifier.

## Example Usage

```hcl
data "oci_network_load_balancer_network_load_balancer" "test_network_load_balancer" {
	#Required
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `backend_sets` - Backend sets associated with the network load balancer.
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
* `backend_sets` - Backend sets associated with the network load balancer.
	* `backends` - Array of backends. 
		* `ip_address` - The IP address of the backend server. Example: `10.0.0.3` 
		* `is_backup` - Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
		* `is_drain` - Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
		* `is_offline` - Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
		* `name` - A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
		* `port` - The communication port for the backend server.  Example: `8080` 
		* `target_id` - The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
		* `weight` - The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/introducton.htm#Policies).  Example: `3` 
	* `health_checker` - The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/HealthCheckPolicies/health-check-policy-management.htm). 
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
	* `is_instant_failover_enabled` - If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy. 
	* `is_preserve_source` - If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default. 
	* `name` - A user-friendly name for the backend set that must be unique and cannot be changed.

		Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

		Example: `example_backend_set` 
	* `policy` - The network load balancer policy for the backend set.  Example: `FIVE_TUPLE` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name, which does not have to be unique, and can be changed.  Example: `example_load_balancer` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer.
* `ip_addresses` - An array of IP addresses. 
    * `ip_address` - An IP address.  Example: `192.168.0.3` 
    * `ip_version` - IP version associated with this IP address.
    * `is_public` - Whether the IP address is public or private.

        If "true", then the IP address is public and accessible from the internet.

        If "false", then the IP address is private and accessible only from within the associated virtual cloud network. 
    * `reserved_ip` - An object representing a reserved IP address to be attached or that is already attached to a network load balancer. 
        * `id` - OCID of the reserved public IP address created with the virtual cloud network.

            Reserved public IP addresses are IP addresses that are registered using the virtual cloud network API.

            Create a reserved public IP address. When you create the network load balancer, enter the OCID of the reserved public IP address in the reservedIp field to attach the IP address to the network load balancer. This task configures the network load balancer to listen to traffic on this IP address.

            Reserved public IP addresses are not deleted when the network load balancer is deleted. The IP addresses become unattached from the network load balancer.

            Example: "ocid1.publicip.oc1.phx.unique_ID" 
* `is_preserve_source_destination` - When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC. Packets are sent to the backend set without any changes to the source and destination IP. 
* `is_private` - Whether the network load balancer has a virtual cloud network-local (private) IP address.

    If "true", then the service assigns a private IP address to the network load balancer.

    If "false", then the service assigns a public IP address to the network load balancer.

	A public network load balancer is accessible from the internet, depending the [security list rules](https://docs.cloud.oracle.com/iaas/Content/network/Concepts/securitylists.htm) for your virtual cloudn network. For more information about public and private network load balancers, see [How Network Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/overview.htm). This value is true by default.

	Example: `true` 
* `is_symmetric_hash_enabled` - This can only be enabled when NLB is working in transparent mode with source destination header preservation enabled.  This removes the additional dependency from NLB backends(like Firewalls) to perform SNAT. 

    Example: `true` 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `listeners` - Listeners associated with the network load balancer.
	* `default_backend_set_name` - The name of the associated backend set.  Example: `example_backend_set` 
	* `ip_version` - IP version associated with the listener.
	* `name` - A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
	* `port` - The communication port for the listener.  Example: `80` 
	* `protocol` - The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP with the wildcard port. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). "ListNetworkLoadBalancersProtocols" API is deprecated and it will not return the updated values. Use the allowed values for the protocol instead.  Example: `TCP`
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `network_security_group_ids` - An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.

    During the creation of the network load balancer, the service adds the new load balancer to the specified network security groups.

    The benefits of associating the network load balancer with network security groups include:
    *  Network security groups define network security rules to govern ingress and egress traffic for the network load balancer.
    *  The network security rules of other resources can reference the network security groups associated with the network load balancer to ensure access.

    Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `nlb_ip_version` - IP version associated with the NLB.
* `state` - The current state of the network load balancer.
* `subnet_id` - The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)."
* `system_tags` - Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}` 
* `time_created` - The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z` 
* `time_updated` - The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z` 

