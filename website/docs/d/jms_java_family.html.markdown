---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_family"
sidebar_current: "docs-oci-datasource-jms-java_family"
description: |-
  Provides details about a specific Java Family in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_java_family
This data source provides details about a specific Java Family resource in Oracle Cloud Infrastructure Jms service.

Returns metadata associated with a specific Java release family.


## Example Usage

```hcl
data "oci_jms_java_family" "test_java_family" {
	#Required
	family_version = var.java_family_family_version
}
```

## Argument Reference

The following arguments are supported:

* `family_version` - (Required) Unique Java family version identifier.


## Attributes Reference

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

