// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_jms_java_downloads_java_license_acceptance_record",
		"test_java_license_acceptance_record",
		acctest.Optional,
		acctest.Update,
		JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation)

	JmsJavaDownloadsJavaLicenseAcceptanceRecordSingularDataSourceRepresentation = map[string]interface{}{
		"java_license_acceptance_record_id": acctest.Representation{
			RepType: acctest.Required,
			Create:  `${oci_jms_java_downloads_java_license_acceptance_record.test_java_license_acceptance_record.id}`},
	}

	JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_jms_java_downloads_java_license_acceptance_record.test_java_license_acceptance_record.id}`},
		"license_type":   acctest.Representation{RepType: acctest.Optional, Create: `OTN`},
		"search_by_user": acctest.Representation{RepType: acctest.Optional, Create: JmsJdUserOcid},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `REVOKED`},
		"filter": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
				"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_jms_java_downloads_java_license_acceptance_record.test_java_license_acceptance_record.id}`}},
			},
		},
	}

	JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"license_acceptance_status": acctest.Representation{RepType: acctest.Required, Create: `ACCEPTED`, Update: `REVOKED`},
		"license_type":              acctest.Representation{RepType: acctest.Required, Create: `OTN`},
		"lifecycle": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
			},
		},
	}
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaLicenseAcceptanceRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaLicenseAcceptanceRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_jms_java_downloads_java_license_acceptance_record.test_java_license_acceptance_record"
	datasourceName := "data.oci_jms_java_downloads_java_license_acceptance_records.test_java_license_acceptance_records"
	singularDatasourceName := "data.oci_jms_java_downloads_java_license_acceptance_record.test_java_license_acceptance_record"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_jms_java_downloads_java_license_acceptance_record",
			"test_java_license_acceptance_record",
			acctest.Required,
			acctest.Create,
			JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation),
		"jmsjavadownloads",
		"javaLicenseAcceptanceRecord",
		t,
	)

	acctest.ResourceTest(t, testAccCheckJmsJavaDownloadsJavaLicenseAcceptanceRecordDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license_acceptance_record",
					"test_java_license_acceptance_record",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "license_acceptance_status", "ACCEPTED"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "OTN"),

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
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license_acceptance_record",
					"test_java_license_acceptance_record",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_acceptance_status", "REVOKED"),
				resource.TestCheckResourceAttr(resourceName, "license_type", "OTN"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license_acceptance_records",
					"test_java_license_acceptance_records",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license_acceptance_record",
					"test_java_license_acceptance_record",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaLicenseAcceptanceRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "license_type", "OTN"),
				resource.TestCheckResourceAttr(datasourceName, "search_by_user", JmsJdUserOcid),
				resource.TestCheckResourceAttr(datasourceName, "status", "REVOKED"),

				resource.TestCheckResourceAttr(datasourceName, "java_license_acceptance_record_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "java_license_acceptance_record_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_license_acceptance_record",
					"test_java_license_acceptance_record",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaLicenseAcceptanceRecordSingularDataSourceRepresentation) +
				JmsJavaDownloadsJavaLicenseAcceptanceRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "java_license_acceptance_record_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_updated_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_acceptance_status", "REVOKED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_type", "OTN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_updated"),
			),
		},
	})
}

func testAccCheckJmsJavaDownloadsJavaLicenseAcceptanceRecordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).JavaDownloadClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_jms_java_downloads_java_license_acceptance_record" {
			noResourceFound = false
			request := oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordRequest{}

			tmp := rs.Primary.ID
			request.JavaLicenseAcceptanceRecordId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")

			response, err := client.GetJavaLicenseAcceptanceRecord(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_jms_java_downloads.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("JmsJavaDownloadsJavaLicenseAcceptanceRecord") {
		resource.AddTestSweepers("JmsJavaDownloadsJavaLicenseAcceptanceRecord", &resource.Sweeper{
			Name:         "JmsJavaDownloadsJavaLicenseAcceptanceRecord",
			Dependencies: acctest.DependencyGraph["javaLicenseAcceptanceRecord"],
			F:            sweepJmsJavaDownloadsJavaLicenseAcceptanceRecordResource,
		})
	}
}

func sweepJmsJavaDownloadsJavaLicenseAcceptanceRecordResource(compartment string) error {
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()
	// JmsJavaDownloadsJavaLicenseAcceptanceRecordResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	javaLicenseAcceptanceRecordIds, err := getJmsJavaDownloadsJavaLicenseAcceptanceRecordIds(compartment)
	if err != nil {
		return err
	}
	for _, javaLicenseAcceptanceRecordId := range javaLicenseAcceptanceRecordIds {
		if ok := acctest.SweeperDefaultResourceId[javaLicenseAcceptanceRecordId]; !ok {
			deleteJavaLicenseAcceptanceRecordRequest := oci_jms_java_downloads.DeleteJavaLicenseAcceptanceRecordRequest{}

			deleteJavaLicenseAcceptanceRecordRequest.JavaLicenseAcceptanceRecordId = &javaLicenseAcceptanceRecordId

			deleteJavaLicenseAcceptanceRecordRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")
			_, error := javaDownloadClient.DeleteJavaLicenseAcceptanceRecord(context.Background(), deleteJavaLicenseAcceptanceRecordRequest)
			if error != nil {
				fmt.Printf("Error deleting JavaLicenseAcceptanceRecord %s %s, It is possible that the resource is already deleted. Please verify manually \n", javaLicenseAcceptanceRecordId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &javaLicenseAcceptanceRecordId, JmsJavaDownloadsJavaLicenseAcceptanceRecordSweepWaitCondition, time.Duration(3*time.Minute),
				JmsJavaDownloadsJavaLicenseAcceptanceRecordSweepResponseFetchOperation, "jms_java_downloads", true)
		}
	}
	return nil
}

func getJmsJavaDownloadsJavaLicenseAcceptanceRecordIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JavaLicenseAcceptanceRecordId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()

	listJavaLicenseAcceptanceRecordsRequest := oci_jms_java_downloads.ListJavaLicenseAcceptanceRecordsRequest{}
	listJavaLicenseAcceptanceRecordsRequest.CompartmentId = &compartmentId
	listJavaLicenseAcceptanceRecordsResponse, err := javaDownloadClient.ListJavaLicenseAcceptanceRecords(context.Background(), listJavaLicenseAcceptanceRecordsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting JavaLicenseAcceptanceRecord list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, javaLicenseAcceptanceRecord := range listJavaLicenseAcceptanceRecordsResponse.Items {
		id := *javaLicenseAcceptanceRecord.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JavaLicenseAcceptanceRecordId", id)
	}
	return resourceIds, nil
}

func JmsJavaDownloadsJavaLicenseAcceptanceRecordSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if javaLicenseAcceptanceRecordResponse, ok := response.Response.(oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordResponse); ok {
		return javaLicenseAcceptanceRecordResponse.LifecycleState != oci_jms_java_downloads.LifecycleStateDeleted
	}
	return false
}

func JmsJavaDownloadsJavaLicenseAcceptanceRecordSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.JavaDownloadClient().GetJavaLicenseAcceptanceRecord(context.Background(), oci_jms_java_downloads.GetJavaLicenseAcceptanceRecordRequest{
		JavaLicenseAcceptanceRecordId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
