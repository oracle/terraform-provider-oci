// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_iot_digital_twin_adapter", IotDigitalTwinAdapterResource())
	tfresource.RegisterResource("oci_iot_digital_twin_instance", IotDigitalTwinInstanceResource())
	tfresource.RegisterResource("oci_iot_digital_twin_model", IotDigitalTwinModelResource())
	tfresource.RegisterResource("oci_iot_digital_twin_relationship", IotDigitalTwinRelationshipResource())
	tfresource.RegisterResource("oci_iot_iot_domain", IotIotDomainResource())
	tfresource.RegisterResource("oci_iot_iot_domain_group", IotIotDomainGroupResource())
}
