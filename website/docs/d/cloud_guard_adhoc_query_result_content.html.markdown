```

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_adhoc_query_result_content"
sidebar_current: "docs-oci-datasource-cloud_guard-adhoc_query_result_content"
description: |-
  Provides details about a specific Adhoc Query Result Content in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_adhoc_query_result_content
This data source provides details about a specific Adhoc Query Result Content resource in Oracle Cloud Infrastructure Cloud Guard service.

Downloads the results for a given adhoc ID (from includes results from all monitoring regions).

## Example Usage

```hcl
data "oci_cloud_guard_adhoc_query_result_content" "test_adhoc_query_result_content" {
	#Required
	adhoc_query_id = oci_cloud_guard_adhoc_query.test_adhoc_query.id
}

## Argument Reference

The following arguments are supported:

* `adhoc_query_id` - (Required) Adhoc query OCID.


## Attributes Reference

The following attributes are exported:```