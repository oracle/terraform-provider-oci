---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report_content"
sidebar_current: "docs-oci-datasource-data_safe-report_content"
description: |-
  Provides details about a specific Report Content in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_report_content
This data source provides details about a specific Report Content resource in Oracle Cloud Infrastructure Data Safe service.

Downloads the specified report in the form of .xls or .pdf.

## Example Usage

```hcl
data "oci_data_safe_report_content" "test_report_content" {
	#Required
	report_id = oci_data_safe_report.test_report.id
}
```

## Argument Reference

The following arguments are supported:

* `report_id` - (Required) Unique report identifier


## Attributes Reference

The following attributes are exported:


