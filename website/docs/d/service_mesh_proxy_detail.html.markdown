---
subcategory: "Service Mesh"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_mesh_proxy_detail"
sidebar_current: "docs-oci-datasource-service_mesh-proxy_detail"
description: |-
  Provides details about a specific Proxy Detail in Oracle Cloud Infrastructure Service Mesh service
---

# Data Source: oci_service_mesh_proxy_detail
This data source provides details about a specific Proxy Detail resource in Oracle Cloud Infrastructure Service Mesh service.

Returns the attributes of the Proxy such as proxy image version.


## Example Usage

```hcl
data "oci_service_mesh_proxy_detail" "test_proxy_detail" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `proxy_image` - Proxy container image version to be deployed.

