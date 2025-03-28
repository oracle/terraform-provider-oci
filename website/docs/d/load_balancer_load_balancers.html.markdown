---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_load_balancers"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancers"
description: |-
  Provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_load_balancers
This data source provides the list of Load Balancers in Oracle Cloud Infrastructure Load Balancer service.

Lists all load balancers in the specified compartment.

## Supported Aliases

* `oci_load_balancers`

## Example Usage

```hcl
data "oci_load_balancer_load_balancers" "test_load_balancers" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	detail = var.load_balancer_detail
	display_name = var.load_balancer_display_name
	state = var.load_balancer_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to list.
* `detail` - (Optional) The level of detail to return for each result. Can be `full` or `simple`.  Example: `full` 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.  Example: `example_load_balancer` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `SUCCEEDED` 


## Attributes Reference

The following attributes are exported:

* `load_balancers` - The list of load_balancers.

### LoadBalancer Reference

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
	* `reserved_ip` - 
        * `id` - Ocid of the Reserved IP/Public Ip created with VCN.

            Reserved IPs are IPs which already registered using VCN API.

            Create a reserved Public IP and then while creating the load balancer pass the ocid of the reserved IP in this field reservedIp to attach the Ip to Load balancer. Load balancer will be configured to listen to traffic on this IP.

            Reserved IPs will not be deleted when the Load balancer is deleted. They will be unattached from the Load balancer.

            Example: "ocid1.publicip.oc1.phx.unique_ID" 
* `ip_addresses` - An array of IP addresses. Deprecated: use ip_address_details instead.
* `is_delete_protection_enabled` - Whether or not the load balancer has delete protection enabled.

	If "true", the loadbalancer will be protected against deletion if configured to accept traffic.

	If "false", the loadbalancer will not be protected against deletion.

	Delete protection is not be enabled unless this field is set to "true". Example: `true` 
* `is_private` - Whether the load balancer has a VCN-local (private) IP address.

	If "true", the service assigns a private IP address to the load balancer.

	If "false", the service assigns a public IP address to the load balancer.

	A public load balancer is accessible from the internet, depending on your VCN's [security list rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm). For more information about public and private load balancers, see [How Load Balancing Works](https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).

	Example: `true` 
* `is_request_id_enabled` - Whether or not the load balancer has the Request Id feature enabled for HTTP listeners.

	If "true", the load balancer will attach a unique request id header to every request passed through from the load balancer to load balancer backends. This same request id header also will be added to the response the lb received from the backend handling the request before the load balancer returns the response to the requestor. The name of the unique request id header is set the by value of requestIdHeader.

	If "false", the loadbalancer not add this unique request id header to either the request passed through to the load balancer backends nor to the reponse returned to the user.

	Example: `true` 
* `network_security_group_ids` - An array of NSG [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the load balancer.

	During the load balancer's creation, the service adds the new load balancer to the specified NSGs.

	The benefits of associating the load balancer with NSGs include:
	*  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	*  The network security rules of other resources can reference the NSGs associated with the load balancer to ensure access.

	Example: ["ocid1.nsg.oc1.phx.unique_ID"] 
* `request_id_header` - If isRequestIdEnabled is true then this field contains the name of the header field that contains the unique request id that is attached to every request from the load balancer to the load balancer backends and to every response from the load balancer.

	If a request to the load balancer already contains a header with same name as specified in requestIdHeader then the load balancer will not change the value of that field.

	If this field is set to "" this field defaults to X-Request-Id. 
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
* `security_attributes` - Extended Defined tags for ZPR for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit", "usagetype" : "zpr"}}}` 
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress). To get a list of available shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerShape/ListShapes) operation.  Example: `100Mbps` 
* `shape_details` - The configuration details to update load balancer to a different shape. 
	* `maximum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the maximum bandwidth (ingress plus egress) that the load balancer can achieve. This bandwidth cannot be always guaranteed. For a guaranteed bandwidth use the minimumBandwidthInMbps parameter.

		The values must be between minimumBandwidthInMbps and 8000 (8Gbps).

      Example: `1500`
    * `minimum_bandwidth_in_mbps` - Bandwidth in Mbps that determines the total pre-provisioned bandwidth (ingress plus egress). The values must be between 0 and the maximumBandwidthInMbps in multiples of 10. The current allowed maximum value is defined in [Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm).  Example: `150`
* `ssl_cipher_suites` - The configuration details of an SSL cipher suite.

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
	*  __oci-default-http2-ssl-cipher-suite-v1__

	"ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES128-GCM-SHA256"
	*  __oci-default-http2-tls-13-ssl-cipher-suite-v1__

	"TLS-AES-128-GCM-SHA256" "TLS-AES-256-GCM-SHA384" "TLS-CHACHA20-POLY1305-SHA256"
	*  __oci-default-http2-tls-12-13-ssl-cipher-suite-v1__

	"ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "DHE-RSA-AES256-GCM-SHA384" "DHE-RSA-AES128-GCM-SHA256" "TLS-AES-128-GCM-SHA256" "TLS-AES-256-GCM-SHA384" "TLS-CHACHA20-POLY1305-SHA256"
	*  __oci-tls-13-recommended-ssl-cipher-suite-v1__

	"TLS-AES-128-GCM-SHA256" "TLS-AES-256-GCM-SHA384" "TLS-CHACHA20-POLY1305-SHA256"
	*  __oci-tls-12-13-wider-ssl-cipher-suite-v1__

	"TLS-AES-128-GCM-SHA256" "TLS-AES-256-GCM-SHA384" "TLS-CHACHA20-POLY1305-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES256-SHA384" "AES128-GCM-SHA256" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA256"
	*  __oci-tls-11-12-13-wider-ssl-cipher-suite-v1__ "TLS-AES-128-GCM-SHA256" "TLS-AES-256-GCM-SHA384" "TLS-CHACHA20-POLY1305-SHA256" "ECDHE-ECDSA-AES128-GCM-SHA256" "ECDHE-RSA-AES128-GCM-SHA256" "ECDHE-ECDSA-AES128-SHA256" "ECDHE-RSA-AES128-SHA256" "ECDHE-ECDSA-AES256-GCM-SHA384" "ECDHE-RSA-AES256-GCM-SHA384" "ECDHE-ECDSA-AES256-SHA384" "ECDHE-RSA-AES256-SHA384" "AES128-GCM-SHA256" "AES128-SHA256" "AES256-GCM-SHA384" "AES256-SHA256" "ECDHE-ECDSA-AES128-SHA" "ECDHE-RSA-AES128-SHA" "ECDHE-RSA-AES256-SHA" "ECDHE-ECDSA-AES256-SHA" "AES128-SHA" "AES256-SHA" 
	* `name` - A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.

		**Note:** The name of your user-defined cipher suite must not be the same as any of Oracle's predefined or reserved SSL cipher suite names:
		* oci-default-ssl-cipher-suite-v1
		* oci-modern-ssl-cipher-suite-v1
		* oci-compatible-ssl-cipher-suite-v1
		* oci-wider-compatible-ssl-cipher-suite-v1
		* oci-customized-ssl-cipher-suite
		* oci-default-http2-ssl-cipher-suite-v1
		* oci-default-http2-tls-13-ssl-cipher-suite-v1
		* oci-default-http2-tls-12-13-ssl-cipher-suite-v1
		* oci-tls-13-recommended-ssl-cipher-suite-v1
		* oci-tls-12-13-wider-ssl-cipher-suite-v1
		* oci-tls-11-12-13-wider-ssl-cipher-suite-v1

		example: `example_cipher_suite`
* `state` - The current state of the load balancer. 
* `subnet_ids` - An array of subnet [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

