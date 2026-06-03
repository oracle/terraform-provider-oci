---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_private_endpoint_summaries"
sidebar_current: "docs-oci-datasource-objectstorage-private_endpoint_summaries"
description: |-
  Provides the list of Private Endpoints in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_private_endpoint_summaries
This data source provides the list of Private Endpoints in Oracle Cloud Infrastructure Object Storage service.

Gets a list of all PrivateEndpointSummary items in a compartment. A PrivateEndpointSummary contains only summary fields for the private endpoint
and does not contain fields like the user-defined metadata.

ListPrivateEndpoints returns a PrivateEndpointSummary containing at most 1000 private endpoints. To paginate through more private endpoints, use the returned
`opc-next-page` value with the `page` request parameter.

To use this and other API operations, you must be authorized in an IAM policy. If you are not authorized,
talk to an administrator. If you are an administrator who needs to write policies to give users access, see
[Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).


## Example Usage

```hcl
data "oci_objectstorage_private_endpoint_summaries" "test_pes" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list private endpoints.
* `namespace` - (Required) The Object Storage namespace used for the request.


## Attributes Reference

The following attributes are exported:

* `private_endpoint_summaries` - The list of private_endpoint_summaries.

### Private Endpoint Reference

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
