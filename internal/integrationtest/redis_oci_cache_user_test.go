// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisOciCacheUserRequiredOnlyResource = RedisOciCacheUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserRepresentation)

	RedisOciCacheUserResourceConfig = RedisOciCacheUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Update, RedisOciCacheUserRepresentation)

	RedisOciCacheUserSingularDataSourceRepresentation = map[string]interface{}{
		"oci_cache_user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_user.test_oci_cache_user.id}`},
	}

	RedisOciCacheUserDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test-tf-user`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheUserDataSourceFilterRepresentation}}

	RedisOciCacheUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_user.test_oci_cache_user.id}`}},
	}

	RedisOciCacheUserRepresentation = map[string]interface{}{
		"acl_string":          acctest.Representation{RepType: acctest.Required, Create: `~* +get`, Update: `~* +get +set`},
		"authentication_mode": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheUserAuthenticationModeRepresentation},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":         acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `test-tf-user`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":              acctest.Representation{RepType: acctest.Required, Create: `ON`, Update: `OFF`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRedisTagsChangesRepresentation},
	}
	RedisOciCacheUserAuthenticationModeRepresentation = map[string]interface{}{
		"authentication_type": acctest.Representation{RepType: acctest.Required, Create: `PASSWORD`, Update: `PASSWORD`},
		"hashed_passwords":    acctest.Representation{RepType: acctest.Required, Create: []string{`8a90d6c49320ca051506f337e9e8a04566cff908b06822a5acecc0f7aaeb442f`}, Update: []string{`b8e70354ba6d993f69d706c70cf8a20b0bfe2c256228f27df99f8a7f1ea3be29`}},
		//"authentication_type": acctest.Representation{RepType: acctest.Required, Create: `IAM`, Update: `IAM`},
	}

	RedisOciCacheUserResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: redis/default
func TestRedisOciCacheUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_redis_oci_cache_user.test_oci_cache_user"
	datasourceName := "data.oci_redis_oci_cache_users.test_oci_cache_users"
	singularDatasourceName := "data.oci_redis_oci_cache_user.test_oci_cache_user"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Create, RedisOciCacheUserRepresentation), "redis", "ociCacheUser", t)

	acctest.ResourceTest(t, testAccCheckRedisOciCacheUserDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acl_string", "~* +get"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.authentication_type", "PASSWORD"),
				//resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.hashed_passwords.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttr(resourceName, "status", "ON"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Create, RedisOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acl_string", "~* +get"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.authentication_type", "PASSWORD"),
				//resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.hashed_passwords.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "ON"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RedisOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RedisOciCacheUserRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acl_string", "~* +get"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.authentication_type", "PASSWORD"),
				//resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.hashed_passwords.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "ON"),

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
			Config: config + compartmentIdVariableStr + RedisOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Update, RedisOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acl_string", "~* +get +set"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.authentication_type", "PASSWORD"),
				//resource.TestCheckResourceAttr(resourceName, "authentication_mode.0.hashed_passwords.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "OFF"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_users", "test_oci_cache_users", acctest.Optional, acctest.Update, RedisOciCacheUserDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Optional, acctest.Update, RedisOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oci_cache_user_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oci_cache_user_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oci_cache_user_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "acl_string", "~* +get +set"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authentication_mode.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authentication_mode.0.authentication_type", "PASSWORD"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "authentication_mode.0.hashed_passwords.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test-tf-user"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "OFF"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RedisOciCacheUserRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"authentication_mode.0.hashed_passwords"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRedisOciCacheUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OciCacheUserClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_redis_oci_cache_user" {
			noResourceFound = false
			request := oci_redis.GetOciCacheUserRequest{}

			tmp := rs.Primary.ID
			request.OciCacheUserId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")

			response, err := client.GetOciCacheUser(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_redis.OciCacheUserLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RedisOciCacheUser") {
		resource.AddTestSweepers("RedisOciCacheUser", &resource.Sweeper{
			Name:         "RedisOciCacheUser",
			Dependencies: acctest.DependencyGraph["ociCacheUser"],
			F:            sweepRedisOciCacheUserResource,
		})
	}
}

func sweepRedisOciCacheUserResource(compartment string) error {
	ociCacheUserClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheUserClient()
	ociCacheUserIds, err := getRedisOciCacheUserIds(compartment)
	if err != nil {
		return err
	}
	for _, ociCacheUserId := range ociCacheUserIds {
		if ok := acctest.SweeperDefaultResourceId[ociCacheUserId]; !ok {
			deleteOciCacheUserRequest := oci_redis.DeleteOciCacheUserRequest{}

			deleteOciCacheUserRequest.OciCacheUserId = &ociCacheUserId

			deleteOciCacheUserRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")
			_, error := ociCacheUserClient.DeleteOciCacheUser(context.Background(), deleteOciCacheUserRequest)
			if error != nil {
				fmt.Printf("Error deleting OciCacheUser %s %s, It is possible that the resource is already deleted. Please verify manually \n", ociCacheUserId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ociCacheUserId, RedisOciCacheUserSweepWaitCondition, time.Duration(3*time.Minute),
				RedisOciCacheUserSweepResponseFetchOperation, "redis", true)
		}
	}
	return nil
}

func getRedisOciCacheUserIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OciCacheUserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ociCacheUserClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheUserClient()

	listOciCacheUsersRequest := oci_redis.ListOciCacheUsersRequest{}
	listOciCacheUsersRequest.CompartmentId = &compartmentId
	listOciCacheUsersRequest.LifecycleState = oci_redis.OciCacheUserLifecycleStateActive
	listOciCacheUsersResponse, err := ociCacheUserClient.ListOciCacheUsers(context.Background(), listOciCacheUsersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OciCacheUser list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ociCacheUser := range listOciCacheUsersResponse.Items {
		id := *ociCacheUser.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OciCacheUserId", id)
	}
	return resourceIds, nil
}

func RedisOciCacheUserSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ociCacheUserResponse, ok := response.Response.(oci_redis.GetOciCacheUserResponse); ok {
		return ociCacheUserResponse.LifecycleState != oci_redis.OciCacheUserLifecycleStateDeleted
	}
	return false
}

func RedisOciCacheUserSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OciCacheUserClient().GetOciCacheUser(context.Background(), oci_redis.GetOciCacheUserRequest{
		OciCacheUserId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
