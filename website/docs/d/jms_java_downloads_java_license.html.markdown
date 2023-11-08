---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_license"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_license"
description: |-
  Provides details about a specific Java License in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_license
This data source provides details about a specific Java License resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Return details of the specified Java license type.


## Example Usage

```hcl
data "oci_jms_java_downloads_java_license" "test_java_license" {
	#Required
	license_type = var.java_license_license_type
}
```

## Argument Reference

The following arguments are supported:

* `license_type` - (Required) Unique Java license type.


## Attributes Reference

The following attributes are exported:

* `display_name` - Commonly used name for the license type.
* `license_type` - License Type
* `license_url` - Publicly accessible license URL containing the detailed terms and conditions.

