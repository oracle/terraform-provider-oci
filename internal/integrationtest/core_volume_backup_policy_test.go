// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	CoreVolumeBackupPolicyRequiredOnlyResource = CoreVolumeBackupPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Required, acctest.Create, CoreVolumeBackupPolicyRepresentation)

	CoreCoreVolumeBackupPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVolumeBackupPolicyDataSourceFilterRepresentation}}
	CoreVolumeBackupPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_backup_policy.test_volume_backup_policy.id}`}},
	}

	CoreVolumeBackupPolicyRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"destination_region": acctest.Representation{RepType: acctest.Optional, Create: `${var.destination_region}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `BackupPolicy1`, Update: `BackupPolicy2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"schedules":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreVolumeBackupPolicySchedulesRepresentation},
	}
	CoreVolumeBackupPolicySchedulesRepresentation = map[string]interface{}{
		"backup_type":       acctest.Representation{RepType: acctest.Required, Create: `INCREMENTAL`, Update: `FULL`},
		"period":            acctest.Representation{RepType: acctest.Required, Create: `ONE_DAY`, Update: `ONE_YEAR`},
		"retention_seconds": acctest.Representation{RepType: acctest.Required, Create: `604800`, Update: `2592000`},
		"day_of_month":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"day_of_week":       acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"hour_of_day":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"month":             acctest.Representation{RepType: acctest.Optional, Create: `JANUARY`, Update: `FEBRUARY`},
		"offset_seconds":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `46800`},
		"offset_type":       acctest.Representation{RepType: acctest.Optional, Create: `STRUCTURED`, Update: `NUMERIC_SECONDS`},
		"time_zone":         acctest.Representation{RepType: acctest.Optional, Create: `UTC`, Update: `REGIONAL_DATA_CENTER_TIME`},
	}

	CoreVolumeBackupPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeBackupPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	destinationRegion := utils.GetEnvSettingWithBlankDefault("destination_region")
	destinationRegionVariableStr := fmt.Sprintf("variable \"destination_region\" { default = \"%s\" }\n", destinationRegion)

	resourceName := "oci_core_volume_backup_policy.test_volume_backup_policy"
	datasourceName := "data.oci_core_volume_backup_policies.test_volume_backup_policies"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreVolumeBackupPolicyResourceDependencies+destinationRegionVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Optional, acctest.Create, CoreVolumeBackupPolicyRepresentation), "core", "volumeBackupPolicy", t)

	acctest.ResourceTest(t, testAccCheckCoreVolumeBackupPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreVolumeBackupPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Required, acctest.Create, CoreVolumeBackupPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVolumeBackupPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreVolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Optional, acctest.Create, CoreVolumeBackupPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "destination_region"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "BackupPolicy1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules", map[string]string{
					"backup_type":       "INCREMENTAL",
					"day_of_month":      "10",
					"day_of_week":       "MONDAY",
					"hour_of_day":       "10",
					"month":             "JANUARY",
					"offset_seconds":    "0",
					"offset_type":       "STRUCTURED",
					"period":            "ONE_DAY",
					"retention_seconds": "604800",
					"time_zone":         "UTC",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreVolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Optional, acctest.Update, CoreVolumeBackupPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "destination_region"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "BackupPolicy2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules", map[string]string{
					"backup_type":       "FULL",
					"day_of_month":      "11",
					"day_of_week":       "TUESDAY",
					"hour_of_day":       "11",
					"month":             "FEBRUARY",
					"offset_seconds":    "46800",
					"offset_type":       "NUMERIC_SECONDS",
					"period":            "ONE_YEAR",
					"retention_seconds": "2592000",
					"time_zone":         "REGIONAL_DATA_CENTER_TIME",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_backup_policies", "test_volume_backup_policies", acctest.Optional, acctest.Update, CoreCoreVolumeBackupPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", acctest.Optional, acctest.Update, CoreVolumeBackupPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.destination_region"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.display_name", "BackupPolicy2"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.schedules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "volume_backup_policies.0.schedules", map[string]string{
					"backup_type":       "FULL",
					"day_of_month":      "11",
					"day_of_week":       "TUESDAY",
					"hour_of_day":       "11",
					"month":             "FEBRUARY",
					"offset_seconds":    "46800",
					"offset_type":       "NUMERIC_SECONDS",
					"period":            "ONE_YEAR",
					"retention_seconds": "2592000",
					"time_zone":         "REGIONAL_DATA_CENTER_TIME",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreVolumeBackupPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreVolumeBackupPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup_policy" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			_, err := client.GetVolumeBackupPolicy(context.Background(), request)

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreVolumeBackupPolicy") {
		resource.AddTestSweepers("CoreVolumeBackupPolicy", &resource.Sweeper{
			Name:         "CoreVolumeBackupPolicy",
			Dependencies: acctest.DependencyGraph["volumeBackupPolicy"],
			F:            sweepCoreVolumeBackupPolicyResource,
		})
	}
}

func sweepCoreVolumeBackupPolicyResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	volumeBackupPolicyIds, err := getCoreVolumeBackupPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeBackupPolicyId := range volumeBackupPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[volumeBackupPolicyId]; !ok {
			deleteVolumeBackupPolicyRequest := oci_core.DeleteVolumeBackupPolicyRequest{}

			deleteVolumeBackupPolicyRequest.PolicyId = &volumeBackupPolicyId

			deleteVolumeBackupPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeBackupPolicy(context.Background(), deleteVolumeBackupPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeBackupPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeBackupPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getCoreVolumeBackupPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VolumeBackupPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

	listVolumeBackupPoliciesRequest := oci_core.ListVolumeBackupPoliciesRequest{}
	listVolumeBackupPoliciesRequest.CompartmentId = &compartmentId
	listVolumeBackupPoliciesResponse, err := blockstorageClient.ListVolumeBackupPolicies(context.Background(), listVolumeBackupPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeBackupPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeBackupPolicy := range listVolumeBackupPoliciesResponse.Items {
		id := *volumeBackupPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeBackupPolicyId", id)
	}
	return resourceIds, nil
}
