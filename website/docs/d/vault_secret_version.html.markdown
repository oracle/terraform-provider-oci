---
subcategory: "Vault"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vault_secret_version"
sidebar_current: "docs-oci-datasource-vault-secret_version"
description: |-
  Provides details about a specific Secret Version in Oracle Cloud Infrastructure Vault service
---

# Data Source: oci_vault_secret_version
This data source provides details about a specific Secret Version resource in Oracle Cloud Infrastructure Vault service.

Gets information about the specified version of a secret.


## Example Usage

```hcl
data "oci_vault_secret_version" "test_secret_version" {
	#Required
	secret_id = "${oci_vault_secret.test_secret.id}"
	secret_version_number = "${var.secret_version_secret_version_number}"
}
```

## Argument Reference

The following arguments are supported:

* `secret_id` - (Required) The OCID of the secret.
* `secret_version_number` - (Required) The version number of the secret.


## Attributes Reference

The following attributes are exported:

* `content_type` - The content type of the secret version's secret contents.
* `name` - The name of the secret version. A name is unique across versions of a secret. 
* `secret_id` - The OCID of the secret.
* `stages` - A list of possible rotation states for the secret version. A secret version marked `CURRENT` is currently in use. A secret version marked `PENDING` is staged and available for use, but has not been applied on the target system and, therefore, has not been rotated into current, active use. The secret most recently uploaded to a vault is always marked `LATEST`. (The first version of a secret is always marked as both `CURRENT` and `LATEST`.) A secret version marked `PREVIOUS` is the secret version that was most recently marked `CURRENT`, before the last secret version rotation. A secret version marked `DEPRECATED` is neither current, pending, nor the previous one in use. Only secret versions marked `DEPRECATED` can be scheduled for deletion. 
* `time_created` - A optional property indicating when the secret version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_current_version_expiry` - An optional property indicating when the current secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the secret version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `version_number` - The version number of the secret.

