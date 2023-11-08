---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_records"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_download_records"
description: |-
  Provides the list of Java Download Records in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_download_records
This data source provides the list of Java Download Records in Oracle Cloud Infrastructure Jms Java Downloads service.

Returns a list of Java download records in a tenancy based on specified parameters.
See [JavaReleases API](https://docs.cloud.oracle.com/iaas/api/#/en/jms/20210610/JavaRelease/ListJavaReleases)
for possible values of `javaFamilyVersion` and `javaReleaseVersion` parameters.


## Example Usage

```hcl
data "oci_jms_java_downloads_java_download_records" "test_java_download_records" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	architecture = var.java_download_record_architecture
	family_version = var.java_download_record_family_version
	os_family = var.java_download_record_os_family
	package_type_detail = var.java_download_record_package_type_detail
	release_version = var.java_download_record_release_version
	time_end = var.java_download_record_time_end
	time_start = var.java_download_record_time_start
}
```

## Argument Reference

The following arguments are supported:

* `architecture` - (Optional) Target Operating System architecture of the artifact.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `family_version` - (Optional) Unique Java family version identifier.
* `os_family` - (Optional) Target Operating System family of the artifact.
* `package_type_detail` - (Optional) Packaging type detail of the artifact.
* `release_version` - (Optional) Unique Java release version identifier.
* `time_end` - (Optional) The end of the time period for which reports are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period for which reports are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `java_download_record_collection` - The list of java_download_record_collection.

### JavaDownloadRecord Reference

The following attributes are exported:

* `items` - A list of Java download records in a tenancy.
	* `architecture` - The target Operating System architecture for the artifact.
	* `download_source_id` - Identifier of the source that downloaded the artifact.
	* `download_type` - Type of download.
	* `family_display_name` - The Java family display name.
	* `family_version` - The Java family version identifier.
	* `os_family` - The target Operating System family for the artifact.
	* `package_type` - The package type(typically the file extension) of the artifact.
	* `package_type_detail` - Additional information about the package type.
	* `release_version` - The Java release version identifier.
	* `time_downloaded` - Timestamp of download.

