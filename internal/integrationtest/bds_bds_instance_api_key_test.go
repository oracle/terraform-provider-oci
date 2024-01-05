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
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BdsBdsInstanceApiKeyRequiredOnlyResource = BdsBdsInstanceApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Required, acctest.Create, BdsbdsInstanceApiKeyRepresentation)

	BdsBdsInstanceApiKeyResourceConfig = BdsBdsInstanceApiKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Optional, acctest.Update, BdsbdsInstanceApiKeyRepresentation)

	BdsBdsbdsInstanceApiKeySingularDataSourceRepresentation = map[string]interface{}{
		"api_key_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_api_key.test_bds_instance_api_key.id}`},
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsbdsInstanceApiKeyDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		// "display_name":    Representation{RepType: Optional, Create: `keyAlias`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		// "user_id":         Representation{RepType: Optional, Create: `${oci_identity_user.test_user.id}`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsbdsInstanceApiKeyDataSourceFilterRepresentation}}
	BdsbdsInstanceApiKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_api_key.test_bds_instance_api_key.id}`}},
	}

	BdsbdsInstanceApiKeyRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"key_alias":       acctest.Representation{RepType: acctest.Required, Create: `keyAlias`},
		"passphrase":      acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`},
		"user_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
		"default_region":  acctest.Representation{RepType: acctest.Optional, Create: `us-ashburn-1`},
	}

	BdsBdsInstanceApiKeyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceApiKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceApiKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_instance_api_key.test_bds_instance_api_key"
	datasourceName := "data.oci_bds_bds_instance_api_keys.test_bds_instance_api_keys"
	singularDatasourceName := "data.oci_bds_bds_instance_api_key.test_bds_instance_api_key"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsInstanceApiKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Optional, acctest.Create, BdsbdsInstanceApiKeyRepresentation), "bds", "bdsInstanceApiKey", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceApiKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstanceApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Required, acctest.Create, BdsbdsInstanceApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "key_alias", "keyAlias"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstanceApiKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstanceApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Optional, acctest.Create, BdsbdsInstanceApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "default_region", "us-ashburn-1"),
				resource.TestCheckResourceAttrSet(resourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_alias", "keyAlias"),
				resource.TestCheckResourceAttr(resourceName, "passphrase", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "pemfilepath"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_api_keys", "test_bds_instance_api_keys", acctest.Optional, acctest.Update, BdsBdsbdsInstanceApiKeyDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsInstanceApiKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Optional, acctest.Update, BdsbdsInstanceApiKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				// resource.TestCheckResourceAttr(datasourceName, "display_name", "keyAlias"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				// resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "bds_api_keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "bds_api_keys.0.default_region", "us-ashburn-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_api_keys.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "bds_api_keys.0.key_alias", "keyAlias"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_api_keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_api_keys.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Required, acctest.Create, BdsBdsbdsInstanceApiKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsInstanceApiKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "default_region", "us-ashburn-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_alias", "keyAlias"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pemfilepath"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsBdsInstanceApiKeyRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getBdsApiKeyCompositeId(resourceName),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"passphrase",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceApiKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_api_key" {
			noResourceFound = false
			request := oci_bds.GetBdsApiKeyRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.ApiKeyId = &value
			}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetBdsApiKey(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.BdsApiKeyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceApiKey") {
		resource.AddTestSweepers("BdsBdsInstanceApiKey", &resource.Sweeper{
			Name:         "BdsBdsInstanceApiKey",
			Dependencies: acctest.DependencyGraph["bdsInstanceApiKey"],
			F:            sweepBdsBdsInstanceApiKeyResource,
		})
	}
}

func sweepBdsBdsInstanceApiKeyResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceApiKeyIds, err := getBdsBdsInstanceApiKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceApiKeyId := range bdsInstanceApiKeyIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceApiKeyId]; !ok {
			deleteBdsApiKeyRequest := oci_bds.DeleteBdsApiKeyRequest{}

			deleteBdsApiKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsApiKey(context.Background(), deleteBdsApiKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceApiKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceApiKeyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceApiKeyId, BdsbdsInstanceApiKeysSweepWaitCondition, time.Duration(3*time.Minute),
				BdsbdsInstanceApiKeysSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceApiKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceApiKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listBdsApiKeysRequest := oci_bds.ListBdsApiKeysRequest{}

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceApiKey resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listBdsApiKeysRequest.BdsInstanceId = &bdsInstanceId

		listBdsApiKeysRequest.LifecycleState = oci_bds.BdsApiKeyLifecycleStateActive
		listBdsApiKeysResponse, err := bdsClient.ListBdsApiKeys(context.Background(), listBdsApiKeysRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceApiKey list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceApiKey := range listBdsApiKeysResponse.Items {
			id := *bdsInstanceApiKey.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceApiKeyId", id)
			acctest.SweeperDefaultResourceId[*bdsInstanceApiKey.DefaultRegion] = true

		}

	}
	return resourceIds, nil
}

func BdsbdsInstanceApiKeysSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceApiKeyResponse, ok := response.Response.(oci_bds.GetBdsApiKeyResponse); ok {
		return bdsInstanceApiKeyResponse.LifecycleState != oci_bds.BdsApiKeyLifecycleStateDeleted
	}
	return false
}

func BdsbdsInstanceApiKeysSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsApiKey(context.Background(), oci_bds.GetBdsApiKeyRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getBdsApiKeyCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/apiKeys/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
