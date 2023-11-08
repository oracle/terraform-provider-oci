---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_licenses"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_licenses"
description: |-
  Provides the list of Java Licenses in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_licenses
This data source provides the list of Java Licenses in Oracle Cloud Infrastructure Jms Java Downloads service.

Return a list with details of all Java licenses.


## Example Usage

```hcl
data "oci_jms_java_downloads_java_licenses" "test_java_licenses" {

	#Optional
	display_name = var.java_license_display_name
	license_type = var.java_license_license_type
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the display name.
* `license_type` - (Optional) Unique Java license type.


## Attributes Reference

The following attributes are exported:

* `java_license_collection` - The list of java_license_collection.

### JavaLicense Reference

The following attributes are exported:

* `display_name` - Commonly used name for the license type.
* `license_type` - License Type
* `license_url` - Publicly accessible license URL containing the detailed terms and conditions.

