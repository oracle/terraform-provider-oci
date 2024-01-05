// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeAuditArchiveRetrievalRequiredOnlyResource = DataSafeAuditArchiveRetrievalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Required, acctest.Create, auditArchiveRetrievalRepresentation)

	DataSafeAuditArchiveRetrievalResourceConfig = DataSafeAuditArchiveRetrievalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Update, auditArchiveRetrievalRepresentation)

	DataSafeauditArchiveRetrievalSingularDataSourceRepresentation = map[string]interface{}{
		"audit_archive_retrieval_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id}`},
	}

	DataSafeauditArchiveRetrievalDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":               acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"audit_archive_retrieval_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id}`},
		"compartment_id_in_subtree":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `Archive retrieval 2021`, Update: `displayName2`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
		"time_of_expiry":             acctest.Representation{RepType: acctest.Optional, Create: `timeOfExpiry`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: auditArchiveRetrievalDataSourceFilterRepresentation}}
	auditArchiveRetrievalDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id}`}},
	}

	auditArchiveRetrievalRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"end_date":       acctest.Representation{RepType: acctest.Required, Create: `2021-05-01T00:00:00.000Z`},
		"start_date":     acctest.Representation{RepType: acctest.Required, Create: `2021-02-01T00:00:00.000Z`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_target.test_target.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Archive retrieval for target prod_dev from month Feb 2021 to May 2021`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Archive retrieval 2021`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DataSafeAuditArchiveRetrievalResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditArchiveRetrievalResource_basic(t *testing.T) {
	t.Skip("Skip this test as this is an infrequent operation depending on internal Audit data archival to have completed.")
	httpreplay.SetScenario("TestDataSafeAuditArchiveRetrievalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval"
	datasourceName := "data.oci_data_safe_audit_archive_retrievals.test_audit_archive_retrievals"
	singularDatasourceName := "data.oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeAuditArchiveRetrievalResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Create, auditArchiveRetrievalRepresentation), "datasafe", "auditArchiveRetrieval", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAuditArchiveRetrievalDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Required, acctest.Create, auditArchiveRetrievalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "end_date", "2021-05-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "start_date", "2021-02-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Create, auditArchiveRetrievalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Archive retrieval for target prod_dev from month Feb 2021 to May 2021"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Archive retrieval 2021"),
				resource.TestCheckResourceAttr(resourceName, "end_date", "2021-05-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "start_date", "2021-02-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(auditArchiveRetrievalRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Archive retrieval for target prod_dev from month Feb 2021 to May 2021"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Archive retrieval 2021"),
				resource.TestCheckResourceAttr(resourceName, "end_date", "2021-05-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "start_date", "2021-02-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Update, auditArchiveRetrievalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "end_date", "2021-05-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "start_date", "2021-02-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_archive_retrievals", "test_audit_archive_retrievals", acctest.Optional, acctest.Update, DataSafeauditArchiveRetrievalDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Optional, acctest.Update, auditArchiveRetrievalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttrSet(datasourceName, "audit_archive_retrieval_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_of_expiry"),

				resource.TestCheckResourceAttr(datasourceName, "audit_archive_retrieval_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_archive_retrieval", "test_audit_archive_retrieval", acctest.Required, acctest.Create, DataSafeauditArchiveRetrievalSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_archive_retrieval_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_event_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "end_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_info"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "start_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_expiry"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_requested"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditArchiveRetrievalResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeAuditArchiveRetrievalDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_audit_archive_retrieval" {
			noResourceFound = false
			request := oci_data_safe.GetAuditArchiveRetrievalRequest{}

			tmp := rs.Primary.ID
			request.AuditArchiveRetrievalId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetAuditArchiveRetrieval(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.AuditArchiveRetrievalLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeAuditArchiveRetrieval") {
		resource.AddTestSweepers("DataSafeAuditArchiveRetrieval", &resource.Sweeper{
			Name:         "DataSafeAuditArchiveRetrieval",
			Dependencies: acctest.DependencyGraph["auditArchiveRetrieval"],
			F:            sweepDataSafeAuditArchiveRetrievalResource,
		})
	}
}

func sweepDataSafeAuditArchiveRetrievalResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	auditArchiveRetrievalIds, err := getDataSafeAuditArchiveRetrievalIds(compartment)
	if err != nil {
		return err
	}
	for _, auditArchiveRetrievalId := range auditArchiveRetrievalIds {
		if ok := acctest.SweeperDefaultResourceId[auditArchiveRetrievalId]; !ok {
			deleteAuditArchiveRetrievalRequest := oci_data_safe.DeleteAuditArchiveRetrievalRequest{}

			deleteAuditArchiveRetrievalRequest.AuditArchiveRetrievalId = &auditArchiveRetrievalId

			deleteAuditArchiveRetrievalRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteAuditArchiveRetrieval(context.Background(), deleteAuditArchiveRetrievalRequest)
			if error != nil {
				fmt.Printf("Error deleting AuditArchiveRetrieval %s %s, It is possible that the resource is already deleted. Please verify manually \n", auditArchiveRetrievalId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &auditArchiveRetrievalId, DataSafeauditArchiveRetrievalsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeauditArchiveRetrievalsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeAuditArchiveRetrievalIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AuditArchiveRetrievalId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listAuditArchiveRetrievalsRequest := oci_data_safe.ListAuditArchiveRetrievalsRequest{}
	listAuditArchiveRetrievalsRequest.CompartmentId = &compartmentId
	listAuditArchiveRetrievalsRequest.LifecycleState = oci_data_safe.ListAuditArchiveRetrievalsLifecycleStateActive
	listAuditArchiveRetrievalsResponse, err := dataSafeClient.ListAuditArchiveRetrievals(context.Background(), listAuditArchiveRetrievalsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AuditArchiveRetrieval list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, auditArchiveRetrieval := range listAuditArchiveRetrievalsResponse.Items {
		id := *auditArchiveRetrieval.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AuditArchiveRetrievalId", id)
	}
	return resourceIds, nil
}

func DataSafeauditArchiveRetrievalsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if auditArchiveRetrievalResponse, ok := response.Response.(oci_data_safe.GetAuditArchiveRetrievalResponse); ok {
		return auditArchiveRetrievalResponse.LifecycleState != oci_data_safe.AuditArchiveRetrievalLifecycleStateDeleted
	}
	return false
}

func DataSafeauditArchiveRetrievalsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetAuditArchiveRetrieval(context.Background(), oci_data_safe.GetAuditArchiveRetrievalRequest{
		AuditArchiveRetrievalId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
