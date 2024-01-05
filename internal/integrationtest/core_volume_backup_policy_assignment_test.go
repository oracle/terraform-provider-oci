// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreVolumeBackupPolicyAssignmentRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", acctest.Required, acctest.Create, CoreVolumeBackupPolicyAssignmentRepresentation)

	CoreCoreVolumeBackupPolicyAssignmentDataSourceRepresentation = map[string]interface{}{
		"asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume.test_volume.id}`},
		"filter":   acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVolumeBackupPolicyAssignmentDataSourceFilterRepresentation}}
	CoreVolumeBackupPolicyAssignmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment.id}`}},
	}

	CoreVolumeBackupPolicyAssignmentRepresentation = map[string]interface{}{
		"asset_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume.test_volume.id}`},
		"policy_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
	}

	CoreVolumeBackupPolicyAssignmentResourceDependencies = utils.VolumeBackupPolicyDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Required, acctest.Create, CoreVolumeRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeBackupPolicyAssignmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupPolicyAssignmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment"
	datasourceName := "data.oci_core_volume_backup_policy_assignments.test_volume_backup_policy_assignments"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreVolumeBackupPolicyAssignmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", acctest.Required, acctest.Create, CoreVolumeBackupPolicyAssignmentRepresentation), "core", "volumeBackupPolicyAssignment", t)

	acctest.ResourceTest(t, testAccCheckCoreVolumeBackupPolicyAssignmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreVolumeBackupPolicyAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", acctest.Required, acctest.Create, CoreVolumeBackupPolicyAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "asset_id"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_backup_policy_assignments", "test_volume_backup_policy_assignments", acctest.Optional, acctest.Update, CoreCoreVolumeBackupPolicyAssignmentDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVolumeBackupPolicyAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", acctest.Optional, acctest.Update, CoreVolumeBackupPolicyAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),

				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policy_assignments.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.asset_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policy_assignments.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreCoreVolumeBackupPolicyAssignmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreVolumeBackupPolicyAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup_policy_assignment" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupPolicyAssignmentRequest{}

			tmp := rs.Primary.ID
			request.PolicyAssignmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetVolumeBackupPolicyAssignment(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
