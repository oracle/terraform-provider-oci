---
layout: "oci"
page_title: "OCI: oci_load_balancer_hostnames"
sidebar_current: "docs-oci-datasource-load_balancer-hostnames"
description: |-
  Provides a list of Hostnames
---

# Data Source: oci_load_balancer_hostnames
The `oci_load_balancer_hostnames` data source allows access to the list of OCI hostnames

Lists all hostname resources associated with the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_hostnames" "test_hostnames" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the hostnames to retrieve. 


## Attributes Reference

The following attributes are exported:

* `hostnames` - The list of hostnames.

### Hostname Reference

The following attributes are exported:

* `hostname` - A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 
* `name` - A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_hostname_001` 

