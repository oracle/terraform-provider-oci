---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_migrate_to_native_vcn_statu"
sidebar_current: "docs-oci-datasource-containerengine-migrate_to_native_vcn_statu"
description: |-
  Provides details about a specific Migrate To Native Vcn Statu in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_migrate_to_native_vcn_statu
This data source provides details about a specific Migrate To Native Vcn Statu resource in Oracle Cloud Infrastructure Container Engine service.

Get details on a cluster's migration to native VCN.

## Example Usage

```hcl
data "oci_containerengine_migrate_to_native_vcn_statu" "test_migrate_to_native_vcn_statu" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `state` - The current migration status of the cluster.
* `time_decommission_scheduled` - The date and time the non-native VCN is due to be decommissioned.

