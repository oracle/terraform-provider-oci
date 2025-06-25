---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_configuration"
sidebar_current: "docs-oci-datasource-nosql-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure NoSQL Database service
---

# Data Source: oci_nosql_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure NoSQL Database service.

Retrieves the current service-level configuration.  The
service may of the standard MULTI_TENANCY type, or of the
HOSTED environment type.  In the latter case, information about the
current state of the environment's global encryption key is
included in the response.


## Example Usage

```hcl
data "oci_nosql_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of a table's compartment.


## Attributes Reference

The following attributes are exported:

* `environment` - The service environment type.
* `kms_key` - Information about the state of the service's encryption key management. The following properties are read-only and ignored when this object is used in UpdateConfiguration: kmsKeyState, timeCreated, timeUpdated. 
	* `id` - The OCID of the KMS encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the id will be a null string. 
	* `kms_key_state` - The current state of the encryption key assigned to this Hosted Environment. Oracle-managed keys will always report an ACTIVE state. 
	* `kms_vault_id` - The OCID of the vault containing the encryption key assigned to this Hosted Environment. If the Hosted Environment is using an Oracle-managed Key (i.e., not a CMEK), then the kmsVaultId will be a null string. 
	* `time_created` - The timestamp when encryption key was first enabled for this Hosted Environment. RFC3339 formatted. 
	* `time_updated` - The timestamp of the last update to the encryption key status. RFC3339 formatted. 

