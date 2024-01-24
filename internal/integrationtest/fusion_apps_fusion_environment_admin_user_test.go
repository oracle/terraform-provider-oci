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
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FusionAppsFusionEnvironmentAdminUserResourceConfig = FusionAppsFusionEnvironmentAdminUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_user", "test_fusion_environment_admin_user", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentAdminUserRepresentation)

	FusionAppsFusionAppsFusionEnvironmentAdminUserSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentAdminUserDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentAdminUserDataSourceFilterRepresentation}}
	FusionAppsFusionEnvironmentAdminUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fusion_apps_fusion_environment_admin_user.test_fusion_environment_admin_user.id}`}},
	}

	FusionAppsFusionEnvironmentAdminUserRepresentation = map[string]interface{}{
		"email_address":         acctest.Representation{RepType: acctest.Required, Create: `JohnSmith@example.com`},
		"first_name":            acctest.Representation{RepType: acctest.Required, Create: `firstName`},
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"last_name":             acctest.Representation{RepType: acctest.Required, Create: `lastName`},
		"password":              acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"username":              acctest.Representation{RepType: acctest.Required, Create: `terraformTest`},
	}

	FusionAppsFusionEnvironmentAdminUserResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentAdminUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentAdminUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fusion_apps_fusion_environment_admin_user.test_fusion_environment_admin_user"
	datasourceName := "data.oci_fusion_apps_fusion_environment_admin_users.test_fusion_environment_admin_users"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_admin_user.test_fusion_environment_admin_user"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FusionAppsFusionEnvironmentAdminUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_user", "test_fusion_environment_admin_user", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentAdminUserRepresentation), "fusionapps", "fusionEnvironmentAdminUser", t)

	acctest.ResourceTest(t, testAccCheckFusionAppsFusionEnvironmentAdminUserDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentAdminUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_user", "test_fusion_environment_admin_user", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentAdminUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "email_address", "JohnSmith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "first_name", "firstName"),
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "last_name", "lastName"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "username", "terraformTest"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_users", "test_fusion_environment_admin_users", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentAdminUserDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentAdminUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_user", "test_fusion_environment_admin_user", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentAdminUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttr(datasourceName, "admin_user_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "admin_user_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_admin_user", "test_fusion_environment_admin_user", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentAdminUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentAdminUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
			),
		},
	})
}

func testAccCheckFusionAppsFusionEnvironmentAdminUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FusionApplicationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fusion_apps_fusion_environment_admin_user" {
			noResourceFound = false
			request := oci_fusion_apps.ListAdminUsersRequest{}

			if value, ok := rs.Primary.Attributes["fusion_environment_id"]; ok {
				request.FusionEnvironmentId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")

			_, err := client.ListAdminUsers(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("FusionAppsFusionEnvironmentAdminUser") {
		resource.AddTestSweepers("FusionAppsFusionEnvironmentAdminUser", &resource.Sweeper{
			Name:         "FusionAppsFusionEnvironmentAdminUser",
			Dependencies: acctest.DependencyGraph["fusionEnvironmentAdminUser"],
			F:            sweepFusionAppsFusionEnvironmentAdminUserResource,
		})
	}
}

func sweepFusionAppsFusionEnvironmentAdminUserResource(compartment string) error {
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()
	fusionEnvironmentAdminUserIds, err := getFusionAppsFusionEnvironmentAdminUserIds(compartment)
	if err != nil {
		return err
	}
	for _, fusionEnvironmentAdminUserId := range fusionEnvironmentAdminUserIds {
		if ok := acctest.SweeperDefaultResourceId[fusionEnvironmentAdminUserId]; !ok {
			deleteFusionEnvironmentAdminUserRequest := oci_fusion_apps.DeleteFusionEnvironmentAdminUserRequest{}

			deleteFusionEnvironmentAdminUserRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fusion_apps")
			_, error := fusionApplicationsClient.DeleteFusionEnvironmentAdminUser(context.Background(), deleteFusionEnvironmentAdminUserRequest)
			if error != nil {
				fmt.Printf("Error deleting FusionEnvironmentAdminUser %s %s, It is possible that the resource is already deleted. Please verify manually \n", fusionEnvironmentAdminUserId, error)
				continue
			}
		}
	}
	return nil
}

func getFusionAppsFusionEnvironmentAdminUserIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FusionEnvironmentAdminUserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fusionApplicationsClient := acctest.GetTestClients(&schema.ResourceData{}).FusionApplicationsClient()

	listAdminUsersRequest := oci_fusion_apps.ListAdminUsersRequest{}

	fusionEnvironmentIds, error := getFusionAppsFusionEnvironmentIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fusionEnvironmentId required for FusionEnvironmentAdminUser resource requests \n")
	}
	for _, fusionEnvironmentId := range fusionEnvironmentIds {
		listAdminUsersRequest.FusionEnvironmentId = &fusionEnvironmentId

		listAdminUsersResponse, err := fusionApplicationsClient.ListAdminUsers(context.Background(), listAdminUsersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FusionEnvironmentAdminUser list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fusionEnvironmentAdminUser := range listAdminUsersResponse.Items {
			id := *fusionEnvironmentAdminUser.Username
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FusionEnvironmentAdminUserId", id)
		}

	}
	return resourceIds, nil
}
