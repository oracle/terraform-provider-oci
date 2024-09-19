// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_bds_auto_scaling_configuration", BdsAutoScalingConfigurationResource())
	tfresource.RegisterResource("oci_bds_bds_instance", BdsBdsInstanceResource())
	tfresource.RegisterResource("oci_bds_bds_instance_api_key", BdsBdsInstanceApiKeyResource())
	tfresource.RegisterResource("oci_bds_bds_instance_metastore_config", BdsBdsInstanceMetastoreConfigResource())
	tfresource.RegisterResource("oci_bds_bds_instance_operation_certificate_managements_management", BdsBdsInstanceOperationCertificateManagementsManagementResource())
	tfresource.RegisterResource("oci_bds_bds_instance_patch_action", BdsBdsInstancePatchActionResource())
	tfresource.RegisterResource("oci_bds_bds_instance_resource_principal_configuration", BdsBdsInstanceResourcePrincipalConfigurationResource())
	tfresource.RegisterResource("oci_bds_bds_instance_os_patch_action", BdsBdsInstanceOSPatchActionResource())
}
