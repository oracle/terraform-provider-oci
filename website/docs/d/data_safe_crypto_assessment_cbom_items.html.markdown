---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_cbom_items"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_cbom_items"
description: |-
  Provides the list of Crypto Assessment Cbom Items in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_cbom_items
This data source provides the list of Crypto Assessment Cbom Items in Oracle Cloud Infrastructure Data Safe service.

Lists the CBOM items for the specified crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_cbom_items" "test_crypto_assessment_cbom_items" {
	#Required
	crypto_assessment_id = oci_data_safe_crypto_assessment.test_crypto_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `crypto_assessment_id` - (Required) The OCID of the crypto assessment.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_cbom_item_collection` - The list of crypto_assessment_cbom_item_collection.

### CryptoAssessmentCbomItem Reference

The following attributes are exported:

* `items` - CBOM items for the crypto assessment.
	* `algorithm` - Cryptographic algorithm or integrity/checksum value observed for the feature.
	* `compliance_driver` - Static compliance standards applicable to the feature.
	* `component_type` - Type of component represented in the CBOM item.
	* `configuration_location` - Locations where this item is configured or stored.
	* `feature` - Feature name represented by the CBOM item.
	* `format` - Cryptographic format used by the feature.
	* `key_size` - Observed key size for the feature.
	* `protocol` - Protocol used by the cryptographic feature.

