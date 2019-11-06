---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_regional_wallet_management"
sidebar_current: "docs-oci-datasource-database-autonomous_database_regional_wallet_management"
description: |-
  Provides details about a specific Autonomous Database Regional Wallet Management in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_regional_wallet_management
This data source provides details about a specific Autonomous Database Regional Wallet Management resource in Oracle Cloud Infrastructure Database service.

Gets the Autonomous Database regional wallet details.


## Example Usage

```hcl
data "oci_database_autonomous_database_regional_wallet_management" "test_autonomous_database_regional_wallet_management" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `state` - The current lifecycle state of the Autonomous Database wallet.
* `time_rotated` - The date and time the wallet was last rotated.

