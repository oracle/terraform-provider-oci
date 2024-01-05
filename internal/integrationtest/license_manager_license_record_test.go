// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LicenseManagerLicenseRecordRequiredOnlyResource = LicenseManagerLicenseRecordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Required, acctest.Create, LicenseManagerLicenseRecordRepresentation)

	LicenseManagerLicenseRecordResourceConfig = LicenseManagerLicenseRecordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Optional, acctest.Update, LicenseManagerLicenseRecordRepresentation)

	LicenseManagerLicenseManagerLicenseRecordSingularDataSourceRepresentation = map[string]interface{}{
		"license_record_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_license_manager_license_record.test_license_record.id}`},
	}

	LicenseManagerLicenseManagerLicenseRecordDataSourceRepresentation = map[string]interface{}{
		"product_license_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_license_manager_product_license.test_product_license.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: LicenseManagerLicenseRecordDataSourceFilterRepresentation}}
	LicenseManagerLicenseRecordDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_license_manager_license_record.test_license_record.id}`}},
	}

	LicenseManagerLicenseRecordRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `LicenseRecordCreate`, Update: `LicenseRecordUpdate`},
		"is_perpetual":       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_unlimited":       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"product_license_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_license_manager_product_license.test_product_license.id}`},
		"expiration_date":    acctest.Representation{RepType: acctest.Required, Create: `2199-06-30T23:59:59.000Z`, Update: `2999-06-30T23:59:59.000Z`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"license_count":      acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `20`},
		"product_id":         acctest.Representation{RepType: acctest.Required, Create: `123`, Update: `234`},
	}

	LicenseManagerLicenseRecordResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Required, acctest.Create, LicenseManagerProductLicenseRepresentation)
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerLicenseRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLicenseManagerLicenseRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_license_manager_license_record.test_license_record"
	datasourceName := "data.oci_license_manager_license_records.test_license_records"
	singularDatasourceName := "data.oci_license_manager_license_record.test_license_record"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LicenseManagerLicenseRecordResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Optional, acctest.Create, LicenseManagerLicenseRecordRepresentation), "licensemanager", "licenseRecord", t)

	acctest.ResourceTest(t, testAccCheckLicenseManagerLicenseRecordDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LicenseManagerLicenseRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Required, acctest.Create, LicenseManagerLicenseRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "LicenseRecordCreate"),
				resource.TestCheckResourceAttr(resourceName, "is_perpetual", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "product_license_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LicenseManagerLicenseRecordResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LicenseManagerLicenseRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Optional, acctest.Create, LicenseManagerLicenseRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "LicenseRecordCreate"),
				resource.TestCheckResourceAttr(resourceName, "expiration_date", "2199-06-30T23:59:59Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_perpetual", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "product_id"),
				resource.TestCheckResourceAttrSet(resourceName, "product_license_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + LicenseManagerLicenseRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Optional, acctest.Update, LicenseManagerLicenseRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "LicenseRecordUpdate"),
				resource.TestCheckResourceAttr(resourceName, "expiration_date", "2999-06-30T23:59:59Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_perpetual", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_count", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "product_id"),
				resource.TestCheckResourceAttrSet(resourceName, "product_license_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_license_records", "test_license_records", acctest.Optional, acctest.Update, LicenseManagerLicenseManagerLicenseRecordDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerLicenseRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Optional, acctest.Update, LicenseManagerLicenseRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "product_license_id"),

				resource.TestCheckResourceAttr(datasourceName, "license_record_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "license_record_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_license_record", "test_license_record", acctest.Required, acctest.Create, LicenseManagerLicenseManagerLicenseRecordSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerLicenseRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_record_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "LicenseRecordUpdate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "expiration_date"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_perpetual", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_unlimited", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_count", "20"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_unit"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_license"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + LicenseManagerLicenseRecordRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLicenseManagerLicenseRecordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LicenseManagerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_license_manager_license_record" {
			noResourceFound = false
			request := oci_license_manager.GetLicenseRecordRequest{}

			tmp := rs.Primary.ID
			request.LicenseRecordId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "license_manager")

			response, err := client.GetLicenseRecord(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_license_manager.LifeCycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !acctest.InSweeperExcludeList("LicenseManagerLicenseRecord") {
		resource.AddTestSweepers("LicenseManagerLicenseRecord", &resource.Sweeper{
			Name:         "LicenseManagerLicenseRecord",
			Dependencies: acctest.DependencyGraph["licenseRecord"],
			F:            sweepLicenseManagerLicenseRecordResource,
		})
	}
}

func sweepLicenseManagerLicenseRecordResource(compartment string) error {
	licenseManagerClient := acctest.GetTestClients(&schema.ResourceData{}).LicenseManagerClient()
	licenseRecordIds, err := getLicenseManagerLicenseRecordIds(compartment)
	if err != nil {
		return err
	}
	for _, licenseRecordId := range licenseRecordIds {
		if ok := acctest.SweeperDefaultResourceId[licenseRecordId]; !ok {
			deleteLicenseRecordRequest := oci_license_manager.DeleteLicenseRecordRequest{}

			deleteLicenseRecordRequest.LicenseRecordId = &licenseRecordId

			deleteLicenseRecordRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "license_manager")
			_, error := licenseManagerClient.DeleteLicenseRecord(context.Background(), deleteLicenseRecordRequest)
			if error != nil {
				fmt.Printf("Error deleting LicenseRecord %s %s, It is possible that the resource is already deleted. Please verify manually \n", licenseRecordId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &licenseRecordId, LicenseManagerLicenseRecordSweepWaitCondition, time.Duration(3*time.Minute),
				LicenseManagerLicenseRecordSweepResponseFetchOperation, "license_manager", true)
		}
	}
	return nil
}

func getLicenseManagerLicenseRecordIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LicenseRecordId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	licenseManagerClient := acctest.GetTestClients(&schema.ResourceData{}).LicenseManagerClient()

	listLicenseRecordsRequest := oci_license_manager.ListLicenseRecordsRequest{}

	productLicenseIds, error := getLicenseManagerProductLicenseIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting productLicenseId required for LicenseRecord resource requests \n")
	}
	for _, productLicenseId := range productLicenseIds {
		listLicenseRecordsRequest.ProductLicenseId = &productLicenseId

		listLicenseRecordsResponse, err := licenseManagerClient.ListLicenseRecords(context.Background(), listLicenseRecordsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LicenseRecord list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, licenseRecord := range listLicenseRecordsResponse.Items {
			id := *licenseRecord.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LicenseRecordId", id)
		}

	}
	return resourceIds, nil
}

func LicenseManagerLicenseRecordSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if licenseRecordResponse, ok := response.Response.(oci_license_manager.GetLicenseRecordResponse); ok {
		return licenseRecordResponse.LifecycleState != oci_license_manager.LifeCycleStateDeleted
	}
	return false
}

func LicenseManagerLicenseRecordSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LicenseManagerClient().GetLicenseRecord(context.Background(), oci_license_manager.GetLicenseRecordRequest{
		LicenseRecordId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
