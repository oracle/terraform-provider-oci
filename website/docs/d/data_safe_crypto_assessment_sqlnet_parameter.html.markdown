---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_sqlnet_parameter"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_sqlnet_parameter"
description: |-
  Provides details about a specific Crypto Assessment Sqlnet Parameter in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_sqlnet_parameter
This data source provides details about a specific Crypto Assessment Sqlnet Parameter resource in Oracle Cloud Infrastructure Data Safe service.

Gets SQLNET.ORA parameter values and quantum-readiness evaluation for the specified crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_sqlnet_parameter" "test_crypto_assessment_sqlnet_parameter" {
	#Required
	crypto_assessment_id = oci_data_safe_crypto_assessment.test_crypto_assessment.id

	#Optional
	parameter = var.crypto_assessment_sqlnet_parameter_parameter
	quantum_readiness = var.crypto_assessment_sqlnet_parameter_quantum_readiness
}
```

## Argument Reference

The following arguments are supported:

* `crypto_assessment_id` - (Required) The OCID of the crypto assessment.
* `parameter` - (Optional) A filter to return only the SQLNET parameter with the specified name.
* `quantum_readiness` - (Optional) Filters SQLNET parameters by quantum-readiness category.


## Attributes Reference

The following attributes are exported:

* `parameters` - SQLNET parameters observed for the assessment.
	* `name` - SQLNET parameter name.
	* `quantum_readiness` - Quantum-readiness classification of this parameter.
	* `value` - SQLNET parameter value payload.
		* `type` - Value type of the SQLNET parameter.
		* `value` - Parsed SQLNET parameter value. TEXT returns a string, BOOLEAN returns a boolean, and LIST returns an array of strings.
* `source` - Source details for SQLNET parameter values.
	* `type` - Source type of SQLNET parameter data.

