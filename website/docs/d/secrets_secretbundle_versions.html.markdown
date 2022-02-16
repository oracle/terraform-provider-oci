---
subcategory: "Secrets"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_secrets_secretbundle_versions"
sidebar_current: "docs-oci-datasource-secrets-secretbundle_versions"
description: |-
  Provides the list of Secretbundle Versions in Oracle Cloud Infrastructure Secrets service
---

# Data Source: oci_secrets_secretbundle_versions
This data source provides the list of Secretbundle Versions in Oracle Cloud Infrastructure Secrets service.

Lists all secret bundle versions for the specified secret.

## Example Usage

```hcl
data "oci_secrets_secretbundle_versions" "test_secretbundle_versions" {
	#Required
	secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `secret_id` - (Required) The OCID of the secret.


## Attributes Reference

The following attributes are exported:

* `secret_bundle_versions` - The list of secret_bundle_versions.

### SecretbundleVersion Reference

The following attributes are exported:

* `secret_id` - The OCID of the secret.
* `stages` - A list of possible rotation states for the secret bundle.
* `time_created` - The time when the secret bundle was created.
* `time_of_deletion` - An optional property indicating when to delete the secret version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_expiry` - An optional property indicating when the secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `version_name` - The version name of the secret bundle, as provided when the secret was created or last rotated. 
* `version_number` - The version number of the secret.

