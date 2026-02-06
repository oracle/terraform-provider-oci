---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_data_file"
sidebar_current: "docs-oci-resource-apm_config-data_file"
description: |-
  Provides the Data File resource in Oracle Cloud Infrastructure Apm Config service
---

# oci_apm_config_data_file
This resource provides the Data File resource in Oracle Cloud Infrastructure Apm Config service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/apm-config/latest/DataFile

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/apm/apm_config

Creates a new data file or replaces an existing one with the same name and type.


## Example Usage

```hcl
resource "oci_apm_config_data_file" "test_data_file" {
	#Required
	put_data_file_body = var.data_file_put_data_file_body
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	apm_type = var.data_file_apm_type
	data_file_name = oci_apm_config_data_file.test_data_file.name

	#Optional
	content_disposition = var.data_file_content_disposition
	content_encoding = var.data_file_content_encoding
	content_language = var.data_file_content_language
	content_md5 = var.data_file_content_md5
	content_type = var.data_file_content_type
	metadata = var.data_file_metadata
}
```

## Argument Reference

The following arguments are supported:

* `content_disposition` - (Optional) (Updatable) Optional parameter that provides presentation information for how the content should be displayed or handled by the recipient.

	For example, to prompt a file download with a custom filename: `attachment; filename="example.txt"` 
* `content_encoding` - (Optional) (Updatable) Optional parameter indicating the content encodings applied to the request body (e.g., gzip, deflate). This value can be used by recipients to determine how to decode the content. 
* `content_language` - (Optional) (Updatable) Optional parameter that indicates the natural language of the content. This value can be used by clients or intermediaries to select or display content based on language preferences. 
* `content_md5` - (Optional) (Updatable) Optional base64-encoded MD5 hash of the request body. If provided, the server will perform a data integrity check by computing the MD5 of the received content and comparing it to the supplied value.

	If the values do not match, the request will be rejected with an HTTP 400 error and a message such as:

	"The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)" 
* `content_type` - (Optional) (Updatable) Optional parameter specifying the media type (MIME type) of the request or response body. If not specified, the default is `application/octet-stream`.

	This value can be used by recipients to determine how to interpret or render the content. 
* `content` - (Optional) The object to upload to the object store. Cannot be defined if `source` is defined. Either one should be defined.
* `source` - (Optional) An absolute path to a file on the local system. Cannot be defined if `content` is defined. Either one should be defined.
* `apm_domain_id` - (Required) (Updatable) The APM Domain ID the request is intended for. 
* `apm_type` - (Required) (Updatable) The type of the data file.
* `data_file_name` - (Required) The name of the data file.
* `metadata` - (Optional) (Updatable) Optional user-defined metadata key and value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data File
	* `update` - (Defaults to 20 minutes), when updating the Data File
	* `delete` - (Defaults to 20 minutes), when destroying the Data File


## Import

DataFiles can be imported using the `id`, e.g.

```
$ terraform import oci_apm_config_data_file.test_data_file "dataFiles/{dataFileName}/apmDomainId/{apmDomainId}/apmType/{apmType}" 
```

