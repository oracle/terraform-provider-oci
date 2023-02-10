---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_protected_database_fetch_configuration"
sidebar_current: "docs-oci-datasource-recovery-protected_database_fetch_configuration"
description: |-
  Provides details about a specific Protected Database Fetch Configuration in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_protected_database_fetch_configuration
This data source provides details about a specific Protected Database Fetch Configuration resource in Oracle Cloud Infrastructure Recovery service.

Downloads the network service configuration file 'tnsnames.ora' for a specified protected database. Applies to user-defined recovery systems only.

## Example Usage

```hcl
data "oci_recovery_protected_database_fetch_configuration" "test_protected_database_fetch_configuration" {
	#Required
	protected_database_id = oci_recovery_protected_database.test_protected_database.id

	#Optional
	base64_encode_content = true
	configuration_type = var.protected_database_fetch_configuration_configuration_type
}
```

## Argument Reference

The following arguments are supported:

* `base64_encode_content` - (Optional) Encodes the downloaded config in base64. It is recommended to set this to `true` to avoid corrupting the file in Terraform state. The default value is `true`.
* `configuration_type` - (Optional) Currently has four config options ALL, TNSNAMES, HOSTS and CABUNDLE. All will return a zipped folder containing the contents of both tnsnames and the certificateChainPem.
* `protected_database_id` - (Required) The protected database OCID.


## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded config file for recovery service. It is base64 encoded by default. To store the config in plaintext set `base_64_encode_content` to false.

