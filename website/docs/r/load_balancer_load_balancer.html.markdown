---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancer"
sidebar_current: "docs-oci-resource-load_balancer-load_balancer"
description: |-
  Provides the Load Balancer resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_load_balancer
This resource provides the Load Balancer resource in Oracle Cloud Infrastructure Load Balancer service.

Creates a new load balancer in the specified compartment. For general information about load balancers,
see [Overview of the Load Balancing Service](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the load balancer to reside. Notice that the load balancer doesn't have to be in the same compartment as the VCN
or backend set. If you're not sure which compartment to use, put the load balancer in the same compartment as the VCN.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You must specify a display name for the load balancer. It does not have to be unique, and you can change it.

For information about Availability Domains, see
[Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Cloud Infrastructure resources, including load balancers, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID
in the response. You can also retrieve a resource's OCID by using a List API operation on that resource type,
or by viewing the resource in the Console. Fore more information, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

When you create a load balancer, the system assigns an IP address.
To get the IP address, use the [GetLoadBalancer](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancer/GetLoadBalancer) operation.


## Supported Aliases

* `oci_load_balancer`

## Example Usage

```hcl
resource "oci_load_balancer_load_balancer" "test_load_balancer" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.load_balancer_display_name
	shape = var.load_balancer_shape
	subnet_ids = var.load_balancer_subnet_ids

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	ip_mode = var.load_balancer_ip_mode
	is_private = var.load_balancer_is_private
	network_security_group_ids = var.load_balancer_network_security_group_ids
	reserved_ips {

		#Optional
		id = var.load_balancer_reserved_ips_id
	}
	shape_details {
		#Required
		maximum_bandwidth_in_mbps = var.load_balancer_shape_details_maximum_bandwidth_in_mbps
		minimum_bandwidth_in_mbps = var.load_balancer_shape_details_minimum_bandwidth_in_mbps
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the load balancer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `example_load_balancer` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ip_mode` - (Optional) IPv6 is currently supported only in the Government Cloud. Whether the load balancer has an IPv4 or IPv6 IP address.

	If "IPV4", the service assigns an IPv4 address and the load balancer supports IPv4 traffic.

	If "IPV6", the service assigns an IPv6 address and the load balancer supports IPv6 traffic.

	Example: "ipMode":"IPV6" 
* `is_private` - (Optional) Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `network_security_group_ids` - (Optional) (Updatable) An array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this load balancer.

	During the load balancer's creation, the service adds the new load balancer to the specified NSGs.

	The benefits of using NSGs with the load balancer include:
	*  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	*  The network security rules of other resources can reference the NSGs associated with the load balancer to ensure access.

	Example: `["ocid1.nsg.oc1.phx.unique_ID"]` 
* `reserved_ips` - (Optional) An array of reserved Ips. Pre-created public IP that will be used as the IP of this load balancer. This reserved IP will not be deleted when load balancer is deleted. This ip should not be already mapped to any other resource.
	* `id` - (Optional) Ocid of the pre-created public IP that should be attached to this load balancer. The public IP will be attached to a private IP. **Note** If public IP resource is present in the config, the terraform plan will throw `After applying this step and refreshing, the plan was not empty` error, and `private_ip_id` needs to be added as an input argument to the public IP resource block or ignore from its lifecycle as shown in [examples](https://github.com/terraform-providers/terraform-provider-oci/blob/507acd0ed6517dbca2fbcfb8100874929c8fd8e1/examples/load_balancer/lb_full/lb_full.tf#L133) to resolve this error.
* `shape` - (Required) (Updatable) A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` *Note: When updating shape for a load balancer, all existing connections to the load balancer will be reset during the update process. Also `10Mbps-Micro` shape cannot be updated to any other shape nor can any other shape be updated to `10Mbps-Micro`.
* `shape_details` - (Optional) (Updatable) The configuration details to create load balancer using Flexible shape. This is required only if shapeName is `Flexible`. 
	* `maximum_bandwidth_in_mbps` - (Required) (Updatable) Bandwidth in Mbps that determines the maximum bandwidth (ingress plus egress) that the load balancer can achieve. This bandwidth cannot be always guaranteed. For a guaranteed bandwidth use the minimumBandwidthInMbps parameter.

		The values must be between minimumBandwidthInMbps and 8192 (8Gbps).

		Example: `1500` 
	* `minimum_bandwidth_in_mbps` - (Required) (Updatable) Bandwidth in Mbps that determines the total pre-provisioned bandwidth (ingress plus egress). The values must be between 10 and the maximumBandwidthInMbps.  Example: `150` 
* `ssl_cipher_suites` - (Optional) The configuration details of an SSL cipher suite.

	The algorithms that compose a cipher suite help you secure Transport Layer Security (TLS) or Secure Socket Layer (SSL) network connections. A cipher suite defines the list of security algorithms your load balancer uses to negotiate with peers while sending and receiving information. The cipher suites you use affect the security level, performance, and compatibility of your data traffic.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.

	Oracle created the following predefined cipher suites that you can specify when you define a resource's [SSL configuration](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/datatypes/SSLConfigurationDetails). You can [create custom cipher suites](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/SSLCipherSuite/CreateSSLCipherSuite) if the predefined cipher suites do not meet your requirements.
	*  __oci-default-ssl-cipher-suite-v1__

	"DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA256" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA384"
	*  __oci-modern-ssl-cipher-suite-v1__

	"AES128-GCM-SHA256" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA256" "DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA384"
	*  __oci-compatible-ssl-cipher-suite-v1__

	"AES128-GCM-SHA256" "AES128-SHA" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA" "AES256-SHA256" "DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA" "ECDHE-RSA-AES256-SHA384"
	*  __oci-wider-compatible-ssl-cipher-suite-v1__

	"AES128-GCM-SHA256" "AES128-SHA" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA" "AES256-SHA256" "CAMELLIA128-SHA" "CAMELLIA256-SHA" "DES-CBC3-SHA" "DH-DSS-AES128-GCM-SHA256" "DH-DSS-AES128-SHA" "DH-DSS-AES128-SHA256" "DH-DSS-AES256-GCM-SHA384" "DH-DSS-AES256-SHA" "DH-DSS-AES256-SHA256" "DH-DSS-CAMELLIA128-SHA" "DH-DSS-CAMELLIA256-SHA" "DH-DSS-DES-CBC3-SHAv" "DH-DSS-SEED-SHA" "DH-RSA-AES128-GCM-SHA256" "DH-RSA-AES128-SHA" "DH-RSA-AES128-SHA256" "DH-RSA-AES256-GCM-SHA384" "DH-RSA-AES256-SHA" "DH-RSA-AES256-SHA256" "DH-RSA-CAMELLIA128-SHA" "DH-RSA-CAMELLIA256-SHA" "DH-RSA-DES-CBC3-SHA" "DH-RSA-SEED-SHA" "DHE-DSS-AES128-GCM-SHA256" "DHE-DSS-AES128-SHA" "DHE-DSS-AES128-SHA256" "DHE-DSS-AES256-GCM-SHA384" "DHE-DSS-AES256-SHA" "DHE-DSS-AES256-SHA256" "DHE-DSS-CAMELLIA128-SHA" "DHE-DSS-CAMELLIA256-SHA" "DHE-DSS-DES-CBC3-SHA" "DHE-DSS-SEED-SHA" "DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA" "DHE-RSA-AES256-SHA256" "DHE-RSA-CAMELLIA128-SHA" "DHE-RSA-CAMELLIA256-SHA" "DHE-RSA-DES-CBC3-SHA" "DHE-RSA-SEED-SHA" "ECDH-ECDSA-AES128-GCM-SHA256" "ECDH-ECDSA-AES128-SHA" "ECDH-ECDSA-AES128-SHA256" "ECDH-ECDSA-AES256-GCM-SHA384" "ECDH-ECDSA-AES256-SHA" "ECDH-ECDSA-AES256-SHA384" "ECDH-ECDSA-DES-CBC3-SHA" "ECDH-ECDSA-RC4-SHA" "ECDH-RSA-AES128-GCM-SHA256" "ECDH-RSA-AES128-SHA" "ECDH-RSA-AES128-SHA256" "ECDH-RSA-AES256-GCM-SHA384" "ECDH-RSA-AES256-SHA" "ECDH-RSA-AES256-SHA384" "ECDH-RSA-DES-CBC3-SHA" "ECDH-RSA-RC4-SHA" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-ECDSA-DES-CBC3-SHA" "ECDHE-ECDSA-RC4-SHA" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA" "ECDHE-RSA-AES256-SHA384" "ECDHE-RSA-DES-CBC3-SHA" "ECDHE-RSA-RC4-SHA" "IDEA-CBC-SHA" "KRB5-DES-CBC3-MD5" "KRB5-DES-CBC3-SHA" "KRB5-IDEA-CBC-MD5" "KRB5-IDEA-CBC-SHA" "KRB5-RC4-MD5" "KRB5-RC4-SHA" "PSK-3DES-EDE-CBC-SHA" "PSK-AES128-CBC-SHA" "PSK-AES256-CBC-SHA" "PSK-RC4-SHA" "RC4-MD5" "RC4-SHA" "SEED-SHA" 
	* `ciphers` - (Required) A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.

		The following ciphers are valid values for this property:
		*  __TLSv1.2 ciphers__

		"AES128-GCM-SHA256" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA256" "DH-DSS-AES128-GCM-SHA256" "DH-DSS-AES128-SHA256" "DH-DSS-AES256-GCM-SHA384" "DH-DSS-AES256-SHA256" "DH-RSA-AES128-GCM-SHA256" "DH-RSA-AES128-SHA256" "DH-RSA-AES256-GCM-SHA384" "DH-RSA-AES256-SHA256" "DHE-DSS-AES128-GCM-SHA256" "DHE-DSS-AES128-SHA256" "DHE-DSS-AES256-GCM-SHA384" "DHE-DSS-AES256-SHA256" "DHE-RSA-AES128-GCM-SHA256" "DHE-RSA-AES128-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES256-SHA256" "ECDH-ECDSA-AES128-GCM-SHA256" "ECDH-ECDSA-AES128-SHA256" "ECDH-ECDSA-AES256-GCM-SHA384" "ECDH-ECDSA-AES256-SHA384" "ECDH-RSA-AES128-GCM-SHA256" "ECDH-RSA-AES128-SHA256" "ECDH-RSA-AES256-GCM-SHA384" "ECDH-RSA-AES256-SHA384" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-SHA384"
		*  __TLSv1 ciphers also supported by TLSv1.2__

		"AES128-SHA" "AES256-SHA" "CAMELLIA128-SHA" "CAMELLIA256-SHA" "DES-CBC3-SHA" "DH-DSS-AES128-SHA" "DH-DSS-AES256-SHA" "DH-DSS-CAMELLIA128-SHA" "DH-DSS-CAMELLIA256-SHA" "DH-DSS-DES-CBC3-SHAv" "DH-DSS-SEED-SHA" "DH-RSA-AES128-SHA" "DH-RSA-AES256-SHA" "DH-RSA-CAMELLIA128-SHA" "DH-RSA-CAMELLIA256-SHA" "DH-RSA-DES-CBC3-SHA" "DH-RSA-SEED-SHA" "DHE-DSS-AES128-SHA" "DHE-DSS-AES256-SHA" "DHE-DSS-CAMELLIA128-SHA" "DHE-DSS-CAMELLIA256-SHA" "DHE-DSS-DES-CBC3-SHA" "DHE-DSS-SEED-SHA" "DHE-RSA-AES128-SHA" "DHE-RSA-AES256-SHA" "DHE-RSA-CAMELLIA128-SHA" "DHE-RSA-CAMELLIA256-SHA" "DHE-RSA-DES-CBC3-SHA" "DHE-RSA-SEED-SHA" "ECDH-ECDSA-AES128-SHA" "ECDH-ECDSA-AES256-SHA" "ECDH-ECDSA-DES-CBC3-SHA" "ECDH-ECDSA-RC4-SHA" "ECDH-RSA-AES128-SHA" "ECDH-RSA-AES256-SHA" "ECDH-RSA-DES-CBC3-SHA" "ECDH-RSA-RC4-SHA" "ECDHE-ECDSA-AES128-SHA" "ECDHE-ECDSA-AES256-SHA" "ECDHE-ECDSA-DES-CBC3-SHA" "ECDHE-ECDSA-RC4-SHA" "ECDHE-RSA-AES128-SHA" "ECDHE-RSA-AES256-SHA" "ECDHE-RSA-DES-CBC3-SHA" "ECDHE-RSA-RC4-SHA" "IDEA-CBC-SHA" "KRB5-DES-CBC3-MD5" "KRB5-DES-CBC3-SHA" "KRB5-IDEA-CBC-MD5" "KRB5-IDEA-CBC-SHA" "KRB5-RC4-MD5" "KRB5-RC4-SHA" "PSK-3DES-EDE-CBC-SHA" "PSK-AES128-CBC-SHA" "PSK-AES256-CBC-SHA" "PSK-RC4-SHA" "RC4-MD5" "RC4-SHA" "SEED-SHA"

		example: `["ECDHE-RSA-AES256-GCM-SHA384","ECDHE-ECDSA-AES256-GCM-SHA384","ECDHE-RSA-AES128-GCM-SHA256"]` 
	* `name` - (Required) A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

		**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
		* oci-default-ssl-cipher-suite-v1
		* oci-modern-ssl-cipher-suite-v1
		* oci-compatible-ssl-cipher-suite-v1
		* oci-wider-compatible-ssl-cipher-suite-v1
		* oci-customized-ssl-cipher-suite

		example: `example_cipher_suite` 

* `subnet_ids` - (Required) An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable.  Example: `example_load_balancer` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer.
* `ip_address_details` - An array of IP addresses. 
	* `ip_address` - An IP address.  Example: `192.168.0.3` 
	* `is_public` - Whether the IP address is public or private.

		If "true", the IP address is public and accessible from the internet.

		If "false", the IP address is private and accessible only from within the associated VCN. 
	* `reserved_ip` - Pre-created public IP that will be used as the IP of this load balancer. This reserved IP will not be deleted when load balancer is deleted. This ip should not be already mapped to any other resource.
		* `id` - Ocid of the pre-created public IP. That should be attahed to this load balancer.
* `ip_addresses` - An array of IP addresses. Deprecated: use ip_address_details instead 
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `network_security_group_ids` - An array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the load balancer.

	During the load balancer's creation, the service adds the new load balancer to the specified NSGs.

	The benefits of associating the load balancer with NSGs include:
	*  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	*  The network security rules of other resources can reference the NSGs associated with the load balancer to ensure access.

	Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `routing_policies` - A named ordered list of routing rules that is applied to a listener.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `condition_language_version` - The version of the language in which `condition` of `rules` are composed. 
	* `name` - The unique name for this list of routing rules. Avoid entering confidential information.  Example: `example_routing_policy` 
	* `rules` - The ordered list of routing rules.
		* `actions` - A list of actions to be applied when conditions of the routing rule are met. 
			* `backend_set_name` - Name of the backend set the listener will forward the traffic to.  Example: `backendSetForImages` 
			* `name` - The name can be one of these values: `FORWARD_TO_BACKENDSET`
		* `condition` - A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action. 
		* `name` - A unique name for the routing policy rule. Avoid entering confidential information. 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `shape_details` - The configuration details to update load balancer to a different shape. 
	* `maximum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the maximum bandwidth (ingress plus egress) that the load balancer can achieve. This bandwidth cannot be always guaranteed. For a guaranteed bandwidth use the minimumBandwidthInMbps parameter.

		The values must be between minimumBandwidthInMbps and 8192 (8Gbps).

		Example: `1500` 
	* `minimum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the total pre-provisioned bandwidth (ingress plus egress). The values must be between 0 and the maximumBandwidthInMbps in multiples of 10. The current allowed maximum value is defined in [Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm).  Example: `150` 
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Load Balancer
	* `update` - (Defaults to 20 minutes), when updating the Load Balancer
	* `delete` - (Defaults to 20 minutes), when destroying the Load Balancer


## Import

LoadBalancers can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_load_balancer.test_load_balancer "id"
```

