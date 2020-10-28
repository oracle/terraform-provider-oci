
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

Get Namespace of a tenancy already onboarded in Log Analytics Application


## Example Usage

```hcl
data "oci_log_analytics_namespace" "test_namespace" {
	#Required
	namespace = var.namespace_namespace
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Log Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Tenancy ID
* `is_onboarded` - if tenancy is onboarded to logging analytics
* `namespace` - namespace name

