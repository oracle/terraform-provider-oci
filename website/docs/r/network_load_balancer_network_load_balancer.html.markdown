---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancer"
sidebar_current: "docs-oci-resource-network_load_balancer-network_load_balancer"
description: |-
  Provides the Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service
---

# oci_network_load_balancer_network_load_balancer
This resource provides the Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service.

Creates a network load balancer.


## Example Usage

```hcl
resource "oci_network_load_balancer_network_load_balancer" "test_network_load_balancer" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.network_load_balancer_display_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	backend_sets {
		#Required
		health_checker {
			#Required
			protocol = var.network_load_balancer_backend_sets_health_checker_protocol

			#Optional
			dns {
				#Required
				domain_name = oci_identity_domain.test_domain.name

				#Optional
				query_class = var.network_load_balancer_backend_sets_health_checker_dns_query_class
				query_type = var.network_load_balancer_backend_sets_health_checker_dns_query_type
				rcodes = var.network_load_balancer_backend_sets_health_checker_dns_rcodes
				transport_protocol = var.network_load_balancer_backend_sets_health_checker_dns_transport_protocol
			}
			interval_in_millis = var.network_load_balancer_backend_sets_health_checker_interval_in_millis
			port = var.network_load_balancer_backend_sets_health_checker_port
			request_data = var.network_load_balancer_backend_sets_health_checker_request_data
			response_body_regex = var.network_load_balancer_backend_sets_health_checker_response_body_regex
			response_data = var.network_load_balancer_backend_sets_health_checker_response_data
			retries = var.network_load_balancer_backend_sets_health_checker_retries
			return_code = var.network_load_balancer_backend_sets_health_checker_return_code
			timeout_in_millis = var.network_load_balancer_backend_sets_health_checker_timeout_in_millis
			url_path = var.network_load_balancer_backend_sets_health_checker_url_path
		}

		#Optional
		backends {
			#Required
			port = var.network_load_balancer_backend_sets_backends_port

			#Optional
			ip_address = var.network_load_balancer_backend_sets_backends_ip_address
			is_backup = var.network_load_balancer_backend_sets_backends_is_backup
			is_drain = var.network_load_balancer_backend_sets_backends_is_drain
			is_offline = var.network_load_balancer_backend_sets_backends_is_offline
			name = var.network_load_balancer_backend_sets_backends_name
			target_id = oci_cloud_guard_target.test_target.id
			weight = var.network_load_balancer_backend_sets_backends_weight
		}
		ip_version = var.network_load_balancer_backend_sets_ip_version
		is_fail_open = var.network_load_balancer_backend_sets_is_fail_open
		is_preserve_source = var.network_load_balancer_backend_sets_is_preserve_source
		policy = var.network_load_balancer_backend_sets_policy
	}
	#Optional
	assigned_ipv6 = var.network_load_balancer_assigned_ipv6
	assigned_private_ipv4 = var.network_load_balancer_assigned_private_ipv4
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_preserve_source_destination = var.network_load_balancer_is_preserve_source_destination
	is_private = var.network_load_balancer_is_private
	is_symmetric_hash_enabled = var.network_load_balancer_is_symmetric_hash_enabled
	network_security_group_ids = var.network_load_balancer_network_security_group_ids
	nlb_ip_version = var.network_load_balancer_nlb_ip_version
	reserved_ips {
		#Optional
		id = var.network_load_balancer_reserved_ips_id
	}
	subnet_ipv6cidr = var.network_load_balancer_subnet_ipv6cidr
}
```

## Argument Reference

The following arguments are supported:

* `backend_sets` - (Optional) Backend sets associated with the network load balancer.
	* `backends` - (Optional) An array of backends. 
		* `ip_address` - (Optional) The IP address of the backend server. Example: `10.0.0.3` 
		* `is_backup` - (Optional) Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
		* `is_drain` - (Optional) Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
		* `is_offline` - (Optional) Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
		* `name` - (Optional) A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
		* `port` - (Required) The communication port for the backend server.  Example: `8080` 
		* `target_id` - (Optional) The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
		* `weight` - (Optional) The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
	* `health_checker` - (Required) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm). 
		* `dns` - (Optional) DNS healthcheck configurations.
			* `domain_name` - (Required) The absolute fully-qualified domain name to perform periodic DNS queries. If not provided, an extra dot will be added at the end of a domain name during the query. 
			* `query_class` - (Optional) The class the dns health check query to use; either IN or CH.  Example: `IN` 
			* `query_type` - (Optional) The type the dns health check query to use; A, AAAA, TXT.  Example: `A` 
			* `rcodes` - (Optional) An array that represents accepetable RCODE values for DNS query response. Example: ["NOERROR", "NXDOMAIN"] 
			* `transport_protocol` - (Optional) DNS transport protocol; either UDP or TCP.  Example: `UDP` 
		* `interval_in_millis` - (Optional) The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: `10000` 
		* `port` - (Optional) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the `Backend` object. The port must be specified if the backend port is 0.  Example: `8080` 
		* `protocol` - (Required) The protocol the health check must use; either HTTP, HTTPS, UDP, TCP or DNS.  Example: `HTTP` 
		* `request_data` - (Optional) Base64 encoded pattern to be sent as UDP or TCP health check probe.
		* `response_body_regex` - (Optional) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
		* `response_data` - (Optional) Base64 encoded pattern to be validated as UDP or TCP health check probe response.
		* `retries` - (Optional) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: `3` 
		* `return_code` - (Optional) The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: `200` 
		* `timeout_in_millis` - (Optional) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: `3000` 
		* `url_path` - (Optional) The path against which to run the health check.  Example: `/healthcheck` 
	* `ip_version` - (Optional) IP version associated with the backend set.
	* `is_fail_open` - (Optional) If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy. The value is false by default. 
	* `is_preserve_source` - (Optional) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default. 
	* `policy` - (Optional) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE`

* `assigned_ipv6` - (Optional) IPv6 address to be assigned to the network load balancer being created. This IP address has to be part of one of the prefixes supported by the subnet. Example: "2607:9b80:9a0a:9a7e:abcd:ef01:2345:6789" 
* `assigned_private_ipv4` - (Optional) Private IP address to be assigned to the network load balancer being created. This IP address has to be in the CIDR range of the subnet where network load balancer is being created Example: "10.0.0.1"
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Network load balancer identifier, which can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_preserve_source_destination` - (Optional) (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact. 
* `is_private` - (Optional) Whether the network load balancer has a virtual cloud network-local (private) IP address.

    If "true", then the service assigns a private IP address to the network load balancer.

    If "false", then the service assigns a public IP address to the network load balancer.

	A public network load balancer is accessible from the internet, depending on the [security list rules](https://docs.cloud.oracle.com/iaas/Content/network/Concepts/securitylists.htm) for your virtual cloud network. For more information about public and private network load balancers, see [How Network Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/overview.htm). This value is true by default.

	Example: `true` 
* `is_symmetric_hash_enabled` - (Optional) (Updatable) This can only be enabled when NLB is working in transparent mode with source destination header preservation enabled.  This removes the additional dependency from NLB backends(like Firewalls) to perform SNAT. 

	Example: `true`
	Example: `true` 
* `listeners` - (Optional) Listeners associated with the network load balancer.
	* `default_backend_set_name` - (Required) The name of the associated backend set.  Example: `example_backend_set` 
	* `ip_version` - (Optional) IP version associated with the listener.
	* `name` - (Required) A friendly name for the listener. It must be unique and it cannot be changed.  Example: `example_listener` 
	* `port` - (Required) The communication port for the listener.  Example: `80` 
	* `protocol` - (Required) The protocol on which the listener accepts connection requests. For public network load balancers, ANY protocol refers to TCP/UDP with the wildcard port. For private network load balancers, ANY protocol refers to TCP/UDP/ICMP (note that ICMP requires isPreserveSourceDestination to be set to true). "ListNetworkLoadBalancersProtocols" API is deprecated and it will not return the updated values. Use the allowed values for the protocol instead.  Example: `TCP`
* `network_security_group_ids` - (Optional) (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.

    During the creation of the network load balancer, the service adds the new load balancer to the specified network security groups.

    The benefits of associating the network load balancer with network security groups include:
    *  Network security groups define network security rules to govern ingress and egress traffic for the network load balancer.
    *  The network security rules of other resources can reference the network security groups associated with the network load balancer to ensure access.

    Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `nlb_ip_version` - (Optional) (Updatable) IP version associated with the NLB.
* `reserved_ips` - (Optional) An array of reserved Ips. 
    * `id` - (Optional) OCID of the reserved public IP address created with the virtual cloud network.

        Reserved public IP addresses are IP addresses that are registered using the virtual cloud network API.

        Create a reserved public IP address. When you create the network load balancer, enter the OCID of the reserved public IP address in the reservedIp field to attach the IP address to the network load balancer. This task configures the network load balancer to listen to traffic on this IP address.

        Reserved public IP addresses are not deleted when the network load balancer is deleted. The IP addresses become unattached from the network load balancer.

        Example: "ocid1.publicip.oc1.phx.unique_ID" 
* `subnet_id` - (Required) The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `subnet_ipv6cidr` - (Optional) IPv6 subnet prefix selection. If Ipv6 subnet prefix is passed, Nlb Ipv6 Address would be assign within the cidr block. NLB has to be dual or single stack ipv6 to support this.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Load Balancer
	* `update` - (Defaults to 20 minutes), when updating the Network Load Balancer
	* `delete` - (Defaults to 20 minutes), when destroying the Network Load Balancer


## Import

NetworkLoadBalancers can be imported using the `id`, e.g.

```
$ terraform import oci_network_load_balancer_network_load_balancer.test_network_load_balancer "id"
```

