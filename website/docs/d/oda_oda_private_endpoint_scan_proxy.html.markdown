---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint_scan_proxy"
sidebar_current: "docs-oci-datasource-oda-oda_private_endpoint_scan_proxy"
description: |-
  Provides details about a specific Oda Private Endpoint Scan Proxy in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_private_endpoint_scan_proxy
This data source provides details about a specific Oda Private Endpoint Scan Proxy resource in Oracle Cloud Infrastructure Digital Assistant service.

Gets the specified ODA Private Endpoint Scan Proxy.

## Example Usage

```hcl
data "oci_oda_oda_private_endpoint_scan_proxy" "test_oda_private_endpoint_scan_proxy" {
	#Required
	oda_private_endpoint_id = oci_oda_oda_private_endpoint.test_oda_private_endpoint.id
	oda_private_endpoint_scan_proxy_id = oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy.id
}
```

## Argument Reference

The following arguments are supported:

* `oda_private_endpoint_id` - (Required) Unique ODA Private Endpoint identifier which is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `oda_private_endpoint_scan_proxy_id` - (Required) Unique ODA Private Endpoint Scan Proxy identifier.


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

