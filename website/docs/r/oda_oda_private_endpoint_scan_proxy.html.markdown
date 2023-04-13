---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint_scan_proxy"
sidebar_current: "docs-oci-resource-oda-oda_private_endpoint_scan_proxy"
description: |-
  Provides the Oda Private Endpoint Scan Proxy resource in Oracle Cloud Infrastructure Digital Assistant service
---

# oci_oda_oda_private_endpoint_scan_proxy
This resource provides the Oda Private Endpoint Scan Proxy resource in Oracle Cloud Infrastructure Digital Assistant service.

Starts an asynchronous job to create an ODA Private Endpoint Scan Proxy.

To monitor the status of the job, take the `opc-work-request-id` response
header value and use it to call `GET /workRequests/{workRequestID}`.


## Example Usage

```hcl
resource "oci_oda_oda_private_endpoint_scan_proxy" "test_oda_private_endpoint_scan_proxy" {
	#Required
	oda_private_endpoint_id = oci_oda_oda_private_endpoint.test_oda_private_endpoint.id
	protocol = var.oda_private_endpoint_scan_proxy_protocol
	scan_listener_infos {

		#Optional
		scan_listener_fqdn = var.oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_fqdn
		scan_listener_ip = var.oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_ip
		scan_listener_port = var.oda_private_endpoint_scan_proxy_scan_listener_infos_scan_listener_port
	}
	scan_listener_type = var.oda_private_endpoint_scan_proxy_scan_listener_type
}
```

## Argument Reference

The following arguments are supported:

* `oda_private_endpoint_id` - (Required) Unique ODA Private Endpoint identifier which is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `protocol` - (Required) The protocol used for communication between client, scanProxy and RAC's scan listeners 
* `scan_listener_infos` - (Required) The FQDN/IPs and port information of customer's Real Application Cluster (RAC)'s SCAN listeners. 
	* `scan_listener_fqdn` - (Optional) FQDN of the customer's Real Application Cluster (RAC)'s SCAN listeners. 
	* `scan_listener_ip` - (Optional) A SCAN listener's IP of the customer's Real Application Cluster (RAC). 
	* `scan_listener_port` - (Optional) The port that customer's Real Application Cluster (RAC)'s SCAN listeners are listening on. 
* `scan_listener_type` - (Required) Type indicating whether Scan listener is specified by its FQDN or list of IPs 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Scan Proxy. 
* `protocol` - The protocol used for communication between client, scanProxy and RAC's scan listeners 
* `scan_listener_infos` - The FQDN/IPs and port information of customer's Real Application Cluster (RAC)'s SCAN listeners. 
	* `scan_listener_fqdn` - FQDN of the customer's Real Application Cluster (RAC)'s SCAN listeners. 
	* `scan_listener_ip` - A SCAN listener's IP of the customer's Real Application Cluster (RAC). 
	* `scan_listener_port` - The port that customer's Real Application Cluster (RAC)'s SCAN listeners are listening on. 
* `scan_listener_type` - Type indicating whether Scan listener is specified by its FQDN or list of IPs 
* `state` - The current state of the ODA Private Endpoint Scan Proxy.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oda Private Endpoint Scan Proxy
	* `update` - (Defaults to 20 minutes), when updating the Oda Private Endpoint Scan Proxy
	* `delete` - (Defaults to 20 minutes), when destroying the Oda Private Endpoint Scan Proxy


## Import

OdaPrivateEndpointScanProxies can be imported using the `id`, e.g.

```
$ terraform import oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy "odaPrivateEndpoints/{odaPrivateEndpointId}/odaPrivateEndpointScanProxies/{odaPrivateEndpointScanProxyId}" 
```

