// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_core "github.com/oracle/oci-go-sdk/v25/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VolumeBackupPolicyRequiredOnlyResource = VolumeBackupPolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", Required, Create, volumeBackupPolicyRepresentation)

	volumeBackupPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, volumeBackupPolicyDataSourceFilterRepresentation}}
	volumeBackupPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup_policy.test_volume_backup_policy.id}`}},
	}

	volumeBackupPolicyRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"destination_region": Representation{repType: Optional, create: `${var.destination_region}`},
		"display_name":       Representation{repType: Optional, create: `BackupPolicy1`, update: `BackupPolicy2`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"schedules":          RepresentationGroup{Optional, volumeBackupPolicySchedulesRepresentation},
	}
	volumeBackupPolicySchedulesRepresentation = map[string]interface{}{
		"backup_type":       Representation{repType: Required, create: `INCREMENTAL`, update: `FULL`},
		"period":            Representation{repType: Required, create: `ONE_DAY`, update: `ONE_YEAR`},
		"retention_seconds": Representation{repType: Required, create: `604800`, update: `2592000`},
		"day_of_month":      Representation{repType: Optional, create: `10`, update: `11`},
		"day_of_week":       Representation{repType: Optional, create: `MONDAY`, update: `TUESDAY`},
		"hour_of_day":       Representation{repType: Optional, create: `10`, update: `11`},
		"month":             Representation{repType: Optional, create: `JANUARY`, update: `FEBRUARY`},
		"offset_seconds":    Representation{repType: Optional, create: `0`, update: `46800`},
		"offset_type":       Representation{repType: Optional, create: `STRUCTURED`, update: `NUMERIC_SECONDS`},
		"time_zone":         Representation{repType: Optional, create: `UTC`, update: `REGIONAL_DATA_CENTER_TIME`},
	}

	VolumeBackupPolicyResourceDependencies = DefinedTagsDependencies
)

func TestCoreVolumeBackupPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	destinationRegion := getEnvSettingWithBlankDefault("destination_region")
	destinationRegionVariableStr := fmt.Sprintf("variable \"destination_region\" { default = \"%s\" }\n", destinationRegion)

	resourceName := "oci_core_volume_backup_policy.test_volume_backup_policy"
	datasourceName := "data.oci_core_volume_backup_policies.test_volume_backup_policies"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeBackupPolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupPolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", Required, Create, volumeBackupPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupPolicyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", Optional, Create, volumeBackupPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_region"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "BackupPolicy1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules", map[string]string{
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", Optional, Update, volumeBackupPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_region"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "BackupPolicy2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules", map[string]string{
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_volume_backup_policies", "test_volume_backup_policies", Optional, Update, volumeBackupPolicyDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupPolicyResourceDependencies + destinationRegionVariableStr +
					generateResourceFromRepresentationMap("oci_core_volume_backup_policy", "test_volume_backup_policy", Optional, Update, volumeBackupPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.destination_region"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.display_name", "BackupPolicy2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.schedules.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "volume_backup_policies.0.schedules", map[string]string{
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreVolumeBackupPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup_policy" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreVolumeBackupPolicy") {
		resource.AddTestSweepers("CoreVolumeBackupPolicy", &resource.Sweeper{
			Name:         "CoreVolumeBackupPolicy",
			Dependencies: DependencyGraph["volumeBackupPolicy"],
			F:            sweepCoreVolumeBackupPolicyResource,
		})
	}
}

func sweepCoreVolumeBackupPolicyResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()
	volumeBackupPolicyIds, err := getVolumeBackupPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeBackupPolicyId := range volumeBackupPolicyIds {
		if ok := SweeperDefaultResourceId[volumeBackupPolicyId]; !ok {
			deleteVolumeBackupPolicyRequest := oci_core.DeleteVolumeBackupPolicyRequest{}

			deleteVolumeBackupPolicyRequest.PolicyId = &volumeBackupPolicyId

			deleteVolumeBackupPolicyRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeBackupPolicy(context.Background(), deleteVolumeBackupPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeBackupPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeBackupPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getVolumeBackupPolicyIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VolumeBackupPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()

	listVolumeBackupPoliciesRequest := oci_core.ListVolumeBackupPoliciesRequest{}
	listVolumeBackupPoliciesRequest.CompartmentId = &compartmentId
	listVolumeBackupPoliciesResponse, err := blockstorageClient.ListVolumeBackupPolicies(context.Background(), listVolumeBackupPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeBackupPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeBackupPolicy := range listVolumeBackupPoliciesResponse.Items {
		id := *volumeBackupPolicy.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VolumeBackupPolicyId", id)
	}
	return resourceIds, nil
}
