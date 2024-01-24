// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	oci_oce "github.com/oracle/oci-go-sdk/v65/oce"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OceOceInstanceRequiredOnlyResource = OceOceInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Required, acctest.Create, OceOceInstanceRepresentation)

	OceOceInstanceResourceConfig = OceOceInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Update, OceOceInstanceRepresentation)

	OceOceOceInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"oce_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oce_oce_instance.test_oce_instance.id}`},
	}

	OceOceOceInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tenancy_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_tenancy.test_tenancy.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OceOceInstanceDataSourceFilterRepresentation}}
	OceOceInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_oce_oce_instance.test_oce_instance.id}`}},
	}

	instanceName                 = utils.RandomString(15, utils.CharsetWithoutDigits)
	OceOceInstanceRepresentation = map[string]interface{}{
		"admin_email":              acctest.Representation{RepType: acctest.Required, Create: `${var.admin_email}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"idcs_access_token":        acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: instanceName},
		"object_storage_namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"tenancy_id":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.id}`},
		"tenancy_name":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}`},
		"add_on_features":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLE_ADVANCED_HOSTING`, `CROSS_REGION_DR`}, Update: []string{`ENABLE_ADVANCED_HOSTING`, `CROSS_REGION_DR`}},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"dr_region":                acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`, Update: `us-phoenix-1`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_access_type":     acctest.Representation{RepType: acctest.Optional, Create: `PUBLIC`},
		"instance_license_type":    acctest.Representation{RepType: acctest.Optional, Create: `PREMIUM`},
		"instance_usage_type":      acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY`, Update: `NONPRIMARY`},
		"upgrade_schedule":         acctest.Representation{RepType: acctest.Optional, Create: `UPGRADE_IMMEDIATELY`},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	OceOceInstanceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tenancy", "test_tenancy", acctest.Required, acctest.Create, IdentityIdentityTenancySingularDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: oce/default
func TestOceOceInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOceOceInstanceResource_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "TestOceOceInstanceResource_basic") {
		t.Skip("Skipping suppressed TestOceOceInstanceResource_basic")
	}

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	adminEmail := utils.GetEnvSettingWithBlankDefault("admin_email")
	adminEmailVariableStr := fmt.Sprintf("variable \"admin_email\" { default = \"%s\" }\n", adminEmail)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oce_oce_instance.test_oce_instance"
	datasourceName := "data.oci_oce_oce_instances.test_oce_instances"
	singularDatasourceName := "data.oci_oce_oce_instance.test_oce_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OceOceInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Create, OceOceInstanceRepresentation), "oce", "oceInstance", t)

	acctest.ResourceTest(t, testAccCheckOceOceInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Required, acctest.Create, OceOceInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "name", instanceName),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Create, OceOceInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "add_on_features.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "dr_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "guid"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
				resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "instance_license_type", "PREMIUM"),
				resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "name", instanceName),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
				resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OceOceInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "add_on_features.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "dr_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "guid"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
				resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "instance_license_type", "PREMIUM"),
				resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "name", instanceName),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
				resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),

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
			Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Update, OceOceInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "add_on_features.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "dr_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "guid"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
				resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "instance_license_type", "PREMIUM"),
				resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "NONPRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "name", instanceName),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
				resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_oce_oce_instances", "test_oce_instances", acctest.Optional, acctest.Update, OceOceOceInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Optional, acctest.Update, OceOceInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(datasourceName, "oce_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.add_on_features.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.admin_email"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.dr_region", ""),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.guid"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.idcs_tenancy"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_access_type", "PUBLIC"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_license_type", "PREMIUM"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_usage_type", "NONPRIMARY"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.name", instanceName),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.object_storage_namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.state_message"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.tenancy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.tenancy_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.upgrade_schedule", "UPGRADE_IMMEDIATELY"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", acctest.Required, acctest.Create, OceOceOceInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceOceInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oce_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "add_on_features.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "admin_email"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "guid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_tenancy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_access_type", "PUBLIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_license_type", "PREMIUM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_usage_type", "NONPRIMARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", instanceName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),
			),
		},
		// verify resource import
		{
			Config:            config + OceOceInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"idcs_access_token",
				"dr_region",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOceOceInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OceInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oce_oce_instance" {
			noResourceFound = false
			request := oci_oce.GetOceInstanceRequest{}

			tmp := rs.Primary.ID
			request.OceInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oce")

			response, err := client.GetOceInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oce.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OceOceInstance") {
		resource.AddTestSweepers("OceOceInstance", &resource.Sweeper{
			Name:         "OceOceInstance",
			Dependencies: acctest.DependencyGraph["oceInstance"],
			F:            sweepOceOceInstanceResource,
		})
	}
}

func sweepOceOceInstanceResource(compartment string) error {
	OceInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).OceInstanceClient()
	oceInstanceIds, err := getOceOceInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, oceInstanceId := range oceInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[oceInstanceId]; !ok {
			deleteOceInstanceRequest := oci_oce.DeleteOceInstanceRequest{}

			deleteOceInstanceRequest.OceInstanceId = &oceInstanceId

			deleteOceInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oce")
			_, error := OceInstanceClient.DeleteOceInstance(context.Background(), deleteOceInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting OceInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", oceInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oceInstanceId, OceOceInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				OceOceInstanceSweepResponseFetchOperation, "oce", true)
		}
	}
	return nil
}

func getOceOceInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OceInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	OceInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).OceInstanceClient()

	listOceInstancesRequest := oci_oce.ListOceInstancesRequest{}
	listOceInstancesRequest.CompartmentId = &compartmentId
	listOceInstancesRequest.LifecycleState = oci_oce.ListOceInstancesLifecycleStateActive
	listOceInstancesResponse, err := OceInstanceClient.ListOceInstances(context.Background(), listOceInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OceInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oceInstance := range listOceInstancesResponse.Items {
		id := *oceInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OceInstanceId", id)
	}
	return resourceIds, nil
}

func OceOceInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oceInstanceResponse, ok := response.Response.(oci_oce.GetOceInstanceResponse); ok {
		return oceInstanceResponse.LifecycleState != oci_oce.LifecycleStateDeleted
	}
	return false
}

func OceOceInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OceInstanceClient().GetOceInstance(context.Background(), oci_oce.GetOceInstanceRequest{
		OceInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
