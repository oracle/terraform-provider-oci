// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterCompleteCredentialRotationManagementRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterCompleteCredentialRotationManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterCompleteCredentialRotationManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	singularDatasourceName := "data.oci_containerengine_cluster_credential_rotation_status.test_cluster_credential_rotation_status"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_complete_credential_rotation_management", "test_cluster_complete_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterCompleteCredentialRotationManagementRepresentation), "containerengine", "clusterCompleteCredentialRotationManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create cluster
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.time_credential_expiration"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// start rotation
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_start_credential_rotation_management", "test_cluster_start_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterStartCredentialRotationManagementRepresentation),
		},

		// verify rotation status
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_start_credential_rotation_management", "test_cluster_start_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterStartCredentialRotationManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_credential_rotation_status", "test_cluster_credential_rotation_status",
					acctest.Optional, acctest.Create, ContainerengineClusterCredentialRotationStatusSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "status", "WAITING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status_details", "NEW_CREDENTIALS_ISSUED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_auto_completion_scheduled"),
			),
		},
		// complete rotation
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_complete_credential_rotation_management", "test_cluster_complete_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterCompleteCredentialRotationManagementRepresentation),
		},
		// verify complete rotation status
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, ContainerengineClusterRepresentationForCredentialRotation) +
				compartmentIdVariableStr + ContainerengineClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_complete_credential_rotation_management", "test_cluster_complete_credential_rotation_management", acctest.Required, acctest.Create, ContainerengineClusterCompleteCredentialRotationManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_credential_rotation_status", "test_cluster_credential_rotation_status",
					acctest.Optional, acctest.Create, ContainerengineClusterCredentialRotationStatusSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "status", "COMPLETED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status_details", "COMPLETED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_auto_completion_scheduled"),
			),
		},
	})
}
