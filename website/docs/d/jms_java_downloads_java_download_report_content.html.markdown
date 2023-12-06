---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_report_content"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_download_report_content"
description: |-
  Provides details about a specific Java Download Report Content in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_download_report_content
This data source provides details about a specific Java Download Report Content resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Retrieve a Java download report with the specified identifier.

## Example Usage

```hcl
data "oci_jms_java_downloads_java_download_report_content" "test_java_download_report_content" {
	#Required
	java_download_report_id = oci_jms_java_downloads_java_download_report.test_java_download_report.id
}
```

## Argument Reference

The following arguments are supported:

* `java_download_report_id` - (Required) Unique Java download report identifier.


## Attributes Reference

The following attributes are exported:


