// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsInstanceOperationCertificateManagementsManagementRequiredOnlyResource = BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Required, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation)

	BdsBdsInstanceOperationCertificateManagementsManagementRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `QWRtaW5AMTIz`},
		"services":               acctest.Representation{RepType: acctest.Required, Create: []string{`OOZIE`}},
		"enable_operation_certificate_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"renew_operation_certificate_management":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"host_cert_details":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsInstanceOperationCertificateManagementsManagementHostCertDetailsRepresentation},
		"root_certificate":                        acctest.Representation{RepType: acctest.Optional, Create: ``},
		"server_key_password":                     acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	BdsBdsInstanceOperationCertificateManagementsManagementRepresentation_RENEW = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `QWRtaW5AMTIz`},
		"services":               acctest.Representation{RepType: acctest.Required, Create: []string{`OOZIE`}},
		"enable_operation_certificate_management": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"renew_operation_certificate_management":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"host_cert_details":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsInstanceOperationCertificateManagementsManagementHostCertDetailsRepresentation},
		"root_certificate":                        acctest.Representation{RepType: acctest.Optional, Create: ``},
		"server_key_password":                     acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	BdsBdsInstanceOperationCertificateManagementsManagementHostCertDetailsRepresentation = map[string]interface{}{}

	BdsInstanceOperationCertificateManagementsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, BdsbdsInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceOperationCertificateManagementsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceOperationCertificateManagementsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_instance_operation_certificate_managements_management.test_bds_instance_operation_certificate_managements_management"
	parentResourceName := "oci_bds_bds_instance_operation_certificate_managements_management.test_bds_instance_operation_certificate_managements_management"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsInstanceOperationCertificateManagementsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Optional, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation), "bds", "bdsInstanceOperationCertificateManagementsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Required, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "QWRtaW5AMTIz"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Required, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_operation_certificate_management", "true"),
				resource.TestCheckResourceAttr(parentResourceName, "renew_operation_certificate_management", "false"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies,
		},
		// create with renew
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Required, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation_RENEW),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "QWRtaW5AMTIz"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
			),
		},
		// verify renew
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Required, acctest.Create, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation_RENEW),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_operation_certificate_management", "false"),
				resource.TestCheckResourceAttr(parentResourceName, "renew_operation_certificate_management", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies,
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Optional, acctest.Update, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "QWRtaW5AMTIz"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + BdsInstanceOperationCertificateManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_operation_certificate_managements_management", "test_bds_instance_operation_certificate_managements_management", acctest.Optional, acctest.Update, BdsBdsInstanceOperationCertificateManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_operation_certificate_management", "false"),
			),
		},
	})
}
