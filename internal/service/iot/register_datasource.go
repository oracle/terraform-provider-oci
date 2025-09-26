// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_iot_digital_twin_adapter", IotDigitalTwinAdapterDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_adapters", IotDigitalTwinAdaptersDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_instance", IotDigitalTwinInstanceDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_instance_content", IotDigitalTwinInstanceContentDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_instances", IotDigitalTwinInstancesDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_model", IotDigitalTwinModelDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_model_spec", IotDigitalTwinModelSpecDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_models", IotDigitalTwinModelsDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_relationship", IotDigitalTwinRelationshipDataSource())
	tfresource.RegisterDatasource("oci_iot_digital_twin_relationships", IotDigitalTwinRelationshipsDataSource())
	tfresource.RegisterDatasource("oci_iot_iot_domain", IotIotDomainDataSource())
	tfresource.RegisterDatasource("oci_iot_iot_domain_group", IotIotDomainGroupDataSource())
	tfresource.RegisterDatasource("oci_iot_iot_domain_groups", IotIotDomainGroupsDataSource())
	tfresource.RegisterDatasource("oci_iot_iot_domains", IotIotDomainsDataSource())
}
