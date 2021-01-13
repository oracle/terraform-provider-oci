// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v32/common"
	oci_core "github.com/oracle/oci-go-sdk/v32/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	volumeBackupPolicyAssignmentDataSourceRepresentation = map[string]interface{}{
		"asset_id": Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
		"filter":   RepresentationGroup{Required, volumeBackupPolicyAssignmentDataSourceFilterRepresentation}}
	volumeBackupPolicyAssignmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment.id}`}},
	}

	volumeBackupPolicyAssignmentRepresentation = map[string]interface{}{
		"asset_id":  Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
		"policy_id": Representation{repType: Required, create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
	}

	VolumeBackupPolicyAssignmentResourceDependencies = VolumeBackupPolicyDependency +
		generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeRepresentation) +
		AvailabilityDomainConfig
)

func TestCoreVolumeBackupPolicyAssignmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupPolicyAssignmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment"
	datasourceName := "data.oci_core_volume_backup_policy_assignments.test_volume_backup_policy_assignments"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeBackupPolicyAssignmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupPolicyAssignmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", Required, Create, volumeBackupPolicyAssignmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "asset_id"),
					resource.TestCheckResourceAttrSet(resourceName, "policy_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateDataSourceFromRepresentationMap("oci_core_volume_backup_policy_assignments", "test_volume_backup_policy_assignments", Optional, Update, volumeBackupPolicyAssignmentDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupPolicyAssignmentResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy_assignment", "test_volume_backup_policy_assignment", Optional, Update, volumeBackupPolicyAssignmentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreVolumeBackupPolicyAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup_policy_assignment" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupPolicyAssignmentRequest{}

			tmp := rs.Primary.ID
			request.PolicyAssignmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
