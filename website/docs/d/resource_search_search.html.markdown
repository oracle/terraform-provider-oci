---
subcategory: "Resource Search"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_search"
sidebar_current: "docs-oci-datasource-resource_search"
description: |-
	Data source which can be used to search for resources across the tenancy using the OCI Search service. The search query is specified using the OCI Search query syntax, and the results include a list of resources that match the query criteria, along with their attributes.
---

# Data Source: oci_resource_search

Returns a list of results that match the specified query, with a maximum limit of 3,000 resources per request.


## Example Usage

```hcl
data "oci_resource_search" "test_resource_search" {
    #Required
    query = var.search_query
}
```

## Argument Reference

The following arguments are supported:

* `query` - (Required) The query statement. See the [OCI Search query syntax](https://docs.oracle.com/en-us/iaas/Content/Search/Concepts/querysyntax.htm) for details on how to construct queries. A maximum of 3,000 search
  results is supported.

## Attributes Reference

The following attributes are exported:

* `results` - A list of resources.
    * `additional_details` - A map containing additional resource-specific attributes. Keys are strings. Values are exported as strings. Plain string values are returned as-is, and non-string values are JSON-encoded strings. Single-element arrays are flattened before JSON encoding. See the OCI Search query overview for details: https://docs.oracle.com/en-us/iaas/Content/Search/Concepts/queryoverview.htm
    * `availability_domain` - The availability domain where this resource exists, if applicable.
    * `compartment_id` - The OCID of the compartment that contains this resource.
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
    * `display_name` - The display name (or name) of this resource, if one exists.
    * `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
    * `identifier` - The unique identifier for this particular resource, usually an OCID.
    * `resource_type` - The resource type name.
    * `state` - The lifecycle state of this resource, if applicable.
    * `system_tags` - System tags associated with this resource, if any. System tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}`
    * `time_created` - The time that this resource was created.
