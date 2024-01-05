// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	//"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
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
	DataSafeAuditTrailResourceConfig = DataSafeAuditTrailResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailUpdateRepresentation)

	DataSafeauditTrailSingularDataSourceRepresentation = map[string]interface{}{
		"audit_trail_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_trail.test_audit_trail.id}`},
	}

	auditTrailDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"audit_trail_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.trail_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `updated-name`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: auditTrailDataSourceFilterRepresentation}}
	auditTrailDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_audit_trail.test_audit_trail.id}`}},
	}

	auditTrailStartRepresentation = map[string]interface{}{
		"audit_trail_id": acctest.Representation{RepType: acctest.Required, Create: `${var.trail_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}

	auditTrailStopRepresentation = map[string]interface{}{
		"audit_trail_id": acctest.Representation{RepType: acctest.Required, Create: `${var.trail_id}`},
		"resume_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `INACTIVE`},
	}

	auditTrailResumeRepresentation = map[string]interface{}{
		"audit_trail_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.trail_id}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `updated-name`, Update: `displayName2`},
		"is_auto_purge_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"resume_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}

	auditTrailUpdateRepresentation = map[string]interface{}{
		"audit_trail_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.trail_id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `updated-name`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_purge_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"resume_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `INACTIVE`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreUpdateChangesRep},
	}

	ignoreUpdateChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DataSafeAuditTrailResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditTrailResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Trail resource")
	httpreplay.SetScenario("TestDataSafeAuditTrailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	trailId := utils.GetEnvSettingWithBlankDefault("trail_ocid")
	trailIdVariableStr := fmt.Sprintf("variable \"trail_id\" { default = \"%s\" }\n", trailId)

	resourceName := "oci_data_safe_audit_trail.test_audit_trail"
	datasourceName := "data.oci_data_safe_audit_trails.test_audit_trails"
	singularDatasourceName := "data.oci_data_safe_audit_trail.test_audit_trail"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+trailIdVariableStr+DataSafeAuditTrailResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Create, auditTrailStartRepresentation), "datasafe", "auditTrail", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAuditTrailDestroy, []resource.TestStep{
		// verify Start
		{
			Config: config + compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailStartRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_trail_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Stop
		{
			Config: config + compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailStopRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_trail_id"),
			),
		},

		// verify Resume
		{
			Config: config + compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailResumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(resourceName, "audit_trail_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_purge_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(resourceName, "audit_trail_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_purge_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_trails", "test_audit_trails", acctest.Optional, acctest.Update, auditTrailDataSourceRepresentation) +
				compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Optional, acctest.Update, auditTrailUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_trail_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "audit_trail_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_trail", "test_audit_trail", acctest.Required, acctest.Create, DataSafeauditTrailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_trail_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_collection_start_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_purge_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_collected"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trail_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + trailIdVariableStr + DataSafeAuditTrailResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + trailIdVariableStr + DataSafeAuditTrailResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`audit_trail_id`, `lifecycle_details`, `resume_trigger`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeAuditTrailDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_audit_trail" {
			noResourceFound = false
			request := oci_data_safe.GetAuditTrailRequest{}

			tmp := rs.Primary.ID
			request.AuditTrailId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetAuditTrail(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeAuditTrail") {
		resource.AddTestSweepers("DataSafeAuditTrail", &resource.Sweeper{
			Name:         "DataSafeAuditTrail",
			Dependencies: acctest.DependencyGraph["auditTrail"],
			F:            sweepDataSafeAuditTrailResource,
		})
	}
}

func sweepDataSafeAuditTrailResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	auditTrailIds, err := getDataSafeAuditTrailIds(compartment)
	if err != nil {
		return err
	}
	for _, auditTrailId := range auditTrailIds {
		if ok := acctest.SweeperDefaultResourceId[auditTrailId]; !ok {
			deleteAuditTrailRequest := oci_data_safe.DeleteAuditTrailRequest{}

			deleteAuditTrailRequest.AuditTrailId = &auditTrailId

			deleteAuditTrailRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteAuditTrail(context.Background(), deleteAuditTrailRequest)
			if error != nil {
				fmt.Printf("Error deleting AuditTrail %s %s, It is possible that the resource is already deleted. Please verify manually \n", auditTrailId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeAuditTrailIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AuditTrailId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listAuditTrailsRequest := oci_data_safe.ListAuditTrailsRequest{}
	listAuditTrailsRequest.CompartmentId = &compartmentId
	listAuditTrailsResponse, err := dataSafeClient.ListAuditTrails(context.Background(), listAuditTrailsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AuditTrail list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, auditTrail := range listAuditTrailsResponse.Items {
		id := *auditTrail.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AuditTrailId", id)
	}
	return resourceIds, nil
}
