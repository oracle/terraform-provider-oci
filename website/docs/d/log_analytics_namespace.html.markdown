---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace"
sidebar_current: "docs-oci-datasource-log_analytics-namespace"
description: |-
  Provides details about a specific Namespace in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace
This data source provides details about a specific Namespace resource in Oracle Cloud Infrastructure Log Analytics service.

This API gets the namespace details of a tenancy already onboarded in Logging Analytics Application


## Example Usage

```hcl
data "oci_log_analytics_namespace" "test_namespace" {
	#Required
	namespace = var.namespace_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `items` - This is the array of namespace of the tenancy request.
	* `compartment_id` - The is the tenancy ID
	* `is_onboarded` - This indicates if the tenancy is onboarded to Logging Analytics
	* `namespace` - This is the namespace name of a tenancy