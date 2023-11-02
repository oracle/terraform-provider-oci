---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_families"
sidebar_current: "docs-oci-datasource-jms-java_families"
description: |-
  Provides the list of Java Families in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_java_families
This data source provides the list of Java Families in Oracle Cloud Infrastructure Jms service.

Returns a list of the Java release family information.
A Java release family is typically a major version in the Java version identifier.


## Example Usage

```hcl
data "oci_jms_java_families" "test_java_families" {

	#Optional
	display_name = var.java_family_display_name
	family_version = var.java_family_family_version
	is_supported_version = var.java_family_is_supported_version
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the Java family.
* `family_version` - (Optional) The version identifier for the Java family.
* `is_supported_version` - (Optional) Filter the Java Release Family versions by support status.


## Attributes Reference

The following attributes are exported:

* `java_family_collection` - The list of java_family_collection.

### JavaFamily Reference

The following attributes are exported:

* `display_name` - The display name of the release family.
* `doc_url` - Link to access the documentation for the release.
* `end_of_support_life_date` - The End of Support Life (EOSL) date of the Java release family (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
* `family_version` - The Java release family identifier.
* `is_supported_version` - Whether or not this Java release family is under active support. Refer [Java Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html) for more details. 
* `latest_release_artifacts` - List of artifacts for the latest Java release version in this family. The script URLs in the response can be used from a command line, or in scripts and dockerfiles to always get the artifacts corresponding to the latest update release version. 
	* `approximate_file_size_in_bytes` - Approximate compressed file size in bytes.
	* `architecture` - The target Operating System architecture for the artifact.
	* `artifact_content_type` - Product content type of this artifact.
	* `artifact_description` - Description of the binary artifact. Typically includes the OS, architecture, and installer type.
	* `artifact_file_name` - The file name of the artifact.
	* `artifact_id` - Unique identifier for the artifact.
	* `download_url` - The endpoint that returns a short-lived artifact download URL in the response payload. This download url can then be used for downloading the artifact. See this [API](https://docs.oracle.com/en-us/iaas/api/#/en/jms/20230601/JavaArtifact/GenerateArtifactDownloadUrl) for more details. 
	* `os_family` - The target Operating System family for the artifact.
	* `package_type` - The package type(typically the file extension) of the artifact.
	* `package_type_detail` - Additional information about the package type.
	* `script_checksum_url` - The URL for retrieving the checksum for the artifact. Depending on the context, this can point to the checksum of the archive or latest update release version artifact. 
	* `script_download_url` - The endpoint for downloading this artifact from command line, automatically in scripts and dockerfiles. Depending on the context, this can point to the archive or latest update release version artifact in the specified family. 
	* `sha256` - SHA256 checksum of the artifact.
* `latest_release_version` - Latest Java release version in the family.
* `support_type` - This indicates the support category for the Java release family.

