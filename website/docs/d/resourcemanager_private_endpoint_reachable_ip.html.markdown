---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_private_endpoint_reachable_ip"
sidebar_current: "docs-oci-datasource-resourcemanager-private_endpoint_reachable_ip"
description: |-
  Provides details about a specific Private Endpoint Reachable Ip in Oracle Cloud Infrastructure Resource Manager service
---

# Data Source: oci_resourcemanager_private_endpoint_reachable_ip
This data source provides details about a specific Private Endpoint Reachable Ip resource in Oracle Cloud Infrastructure Resource Manager service.

Gets the alternative IP address of the private resource. This IP will be used by Resource Manager Service to connect to the private resource.


## Example Usage

```hcl
data "oci_resourcemanager_private_endpoint_reachable_ip" "test_private_endpoint_reachable_ip" {
	#Required
	private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
	private_ip = var.private_endpoint_reachable_ip_private_ip
}
```

## Argument Reference

The following arguments are supported:

* `private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `private_ip` - (Required) The IP address of the resource in the private subnet. 


## Attributes Reference

The following attributes are exported:

* `ip_address` - An IP address for the Resource Manager service to use for connection to the private resource.

