---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_releases"
sidebar_current: "docs-oci-datasource-jms-java_releases"
description: |-
  Provides the list of Java Releases in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_java_releases
This data source provides the list of Java Releases in Oracle Cloud Infrastructure Jms service.

Returns a list of Java releases.


## Example Usage

```hcl
data "oci_jms_java_releases" "test_java_releases" {

	#Optional
	family_version = var.java_release_family_version
	jre_security_status = var.java_release_jre_security_status
	license_type = var.java_release_license_type
	release_type = var.java_release_release_type
	release_version = var.java_release_release_version
}
```

## Argument Reference

The following arguments are supported:

* `family_version` - (Optional) The version identifier for the Java family.
* `jre_security_status` - (Optional) The security status of the Java Runtime.
* `license_type` - (Optional) Java license type.
* `release_type` - (Optional) Java release type.
* `release_version` - (Optional) Unique Java release version identifier


## Attributes Reference

The following attributes are exported:

* `java_release_collection` - The list of java_release_collection.

### JavaRelease Reference

The following attributes are exported:

* `artifact_content_types` - Artifact content types for the Java version.
* `artifacts` - List of Java artifacts.
	* `approximate_file_size_in_bytes` - Approximate compressed file size in bytes.
	* `architecture` - The target Operating System architecture for the artifact.
	* `artifact_content_type` - Product content type of this artifact.
	* `artifact_description` - Description of the binary artifact. Typically includes the OS, architecture, and installer type.
	* `artifact_file_name` - The file name of the artifact.
	* `artifact_id` - Unique identifier for the artifact.
	* `download_url` - The endpoint that returns a short-lived artifact download URL in the response payload. This download url can then be used for downloading the artifact. See this [API](https://docs.oracle.com/en-us/iaas/api/#/en/jms-java-download/20230601/DownloadUrl/GenerateArtifactDownloadUrl) for more details. 
	* `os_family` - The target Operating System family for the artifact.
	* `package_type` - The package type(typically the file extension) of the artifact.
	* `package_type_detail` - Additional information about the package type.
	* `script_checksum_url` - The URL for retrieving the checksum for the artifact. Depending on the context, this can point to the checksum of the archive or latest update release version artifact. 
	* `script_download_url` - The endpoint for downloading this artifact from command line, automatically in scripts and dockerfiles. Depending on the context, this can point to the archive or latest update release version artifact in the specified family. 
	* `sha256` - SHA256 checksum of the artifact.
* `days_under_security_baseline` - The number of days since this release has been under the security baseline.
* `family_details` - Metadata associated with a specific Java release family. A Java release family is typically a major version in the Java version identifier. 
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
		* `download_url` - The endpoint that returns a short-lived artifact download URL in the response payload. This download url can then be used for downloading the artifact. See this [API](https://docs.oracle.com/en-us/iaas/api/#/en/jms-java-download/20230601/DownloadUrl/GenerateArtifactDownloadUrl) for more details. 
		* `os_family` - The target Operating System family for the artifact.
		* `package_type` - The package type(typically the file extension) of the artifact.
		* `package_type_detail` - Additional information about the package type.
		* `script_checksum_url` - The URL for retrieving the checksum for the artifact. Depending on the context, this can point to the checksum of the archive or latest update release version artifact. 
		* `script_download_url` - The endpoint for downloading this artifact from command line, automatically in scripts and dockerfiles. Depending on the context, this can point to the archive or latest update release version artifact in the specified family. 
		* `sha256` - SHA256 checksum of the artifact.
	* `latest_release_version` - Latest Java release version in the family.
	* `release_date` - The date on which the Java release family was first made available (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 
	* `support_type` - This indicates the support category for the Java release family.
* `family_version` - Java release family identifier.
* `license_details` - Information about a license type for Java.
	* `display_name` - Commonly used name for the license type.
	* `license_type` - License Type
	* `license_url` - Publicly accessible license URL containing the detailed terms and conditions.
* `license_type` - License type for the Java version.
* `mos_patches` - List of My Oracle Support(MoS) patches available for this release. This information is only available for `BPR` release type. 
	* `display_name` - Commonly used name for the MoS release.
	* `patch_url` - MoS URL to access the artifacts for the Java release.
* `parent_release_version` - Parent Java release version identifier. This is applicable for BPR releases.
* `release_date` - The release date of the Java version (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `release_notes_url` - Release notes associated with the Java version.
* `release_type` - Release category of the Java version.
* `release_version` - Java release version identifier.
* `security_status` - The security status of the Java version.

