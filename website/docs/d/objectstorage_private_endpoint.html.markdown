---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_private_endpoint"
sidebar_current: "docs-oci-datasource-objectstorage-private-endpoint"
description: |-
  Provides details about a specific Private Endpoint in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_private_endpoint
This data source provides details about a specific Private Endpoint resource in Oracle Cloud Infrastructure Object Storage service.

Gets the current representation of the given private endpoint in the given Object Storage namespace.


## Example Usage

```hcl
data "oci_objectstorage_private_endpoint" "test_pe" {
	#Required
	name = var.pe_name
	namespace = var.namespace
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the private endpoint. Avoid entering confidential information. Example: `my-pe1`
* `namespace` - (Required) The Object Storage namespace used for the request.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment ID in which the private endpoint resource exists in.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the private endpoint.
* `etag` - The entity tag for the Private Endpoint.
* `fqdns` - The object representing FQDN details formed using prefix and additionalPrefixes.
* `name` - The name of the private endpoint. Avoid entering confidential information. Example: my-pe1
* `namespace` - The Object Storage namespace in which the private endpoint resides.
* `prefix` - The DNS prefix value chosen which is the first part of the URL used to access Object Storage.
* `state` - The lifecycle state of the private endpoint resource.
* `time_created` - The date and time the private endpoint was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
* `time_modified` - The date and time the private endpoint was updated, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).

