---
subcategory: "Secrets"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_secrets_secretbundle"
sidebar_current: "docs-oci-datasource-secrets-secretbundle"
description: |-
  Provides details about a specific Secretbundle in Oracle Cloud Infrastructure Secrets service
---

# Data Source: oci_secrets_secretbundle
This data source provides details about a specific Secretbundle resource in Oracle Cloud Infrastructure Secrets service.

Gets a secret bundle that matches either the specified `stage`, `label`, or `versionNumber` parameter. 
If none of these parameters are provided, the bundle for the secret version marked as `CURRENT` will be returned.


## Example Usage

```hcl
data "oci_secrets_secretbundle" "test_secretbundle" {
	#Required
	secret_id = oci_vault_secret.test_secret.id

	#Optional
	secret_version_name = oci_vault_secret_version.test_secret_version.name
	stage = var.secretbundle_stage
	version_number = var.secretbundle_version_number
}
```

## Argument Reference

The following arguments are supported:

* `secret_id` - (Required) The OCID of the secret.
* `secret_version_name` - (Optional) The name of the secret. (This might be referred to as the name of the secret version. Names are unique across the different versions of a secret.)
* `stage` - (Optional) The rotation state of the secret version.
* `version_number` - (Optional) The version number of the secret.


## Attributes Reference

The following attributes are exported:

* `metadata` - Customer-provided contextual metadata for the secret. 
* `secret_bundle_content` - The contents of the secret.
	* `content` - The base64-encoded content of the secret.
	* `content_type` - The formatting type of the secret contents.
* `secret_id` - The OCID of the secret.
* `stages` - A list of possible rotation states for the secret version.
* `time_created` - The time when the secret bundle was created.
* `time_of_deletion` - An optional property indicating when to delete the secret version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_expiry` - An optional property indicating when the secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `version_name` - The name of the secret version. Labels are unique across the different versions of a particular secret. 
* `version_number` - The version number of the secret.

