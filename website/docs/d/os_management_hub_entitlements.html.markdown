---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_entitlements"
sidebar_current: "docs-oci-datasource-os_management_hub-entitlements"
description: |-
  Provides the list of Entitlements in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_entitlements
This data source provides the list of Entitlements in Oracle Cloud Infrastructure Os Management Hub service.

Lists entitlements in the specified tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Filter the list against a variety of criteria including but 
not limited to its Customer Support Identifier (CSI), and vendor name.


## Example Usage

```hcl
data "oci_os_management_hub_entitlements" "test_entitlements" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	csi = var.entitlement_csi
	vendor_name = var.entitlement_vendor_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `csi` - (Optional) A filter to return entitlements that match the given CSI.
* `vendor_name` - (Optional) A filter to return only resources that match the given vendor name.


## Attributes Reference

The following attributes are exported:

* `entitlement_collection` - The list of entitlement_collection.

### Entitlement Reference

The following attributes are exported:

* `items` - List of entitlements.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy containing the entitlement.
	* `csi` - The Customer Support Identifier (CSI) which unlocks the software sources. The CSI is is a unique key given to a customer and it uniquely identifies the entitlement.
	* `vendor_name` - The vendor for the entitlement.

