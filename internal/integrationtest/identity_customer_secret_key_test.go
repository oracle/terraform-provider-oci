// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	customerSecretKeyDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: customerSecretKeyDataSourceFilterRepresentation}}
	customerSecretKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_customer_secret_key.test_customer_secret_key.id}`}},
	}

	customerSecretKeyRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"user_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	CustomerSecretKeyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityCustomerSecretKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCustomerSecretKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_customer_secret_key.test_customer_secret_key"
	datasourceName := "data.oci_identity_customer_secret_keys.test_customer_secret_keys"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CustomerSecretKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", acctest.Required, acctest.Create, customerSecretKeyRepresentation), "identity", "customerSecretKey", t)

	acctest.ResourceTest(t, testAccCheckIdentityCustomerSecretKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", acctest.Required, acctest.Create, customerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					userId, _ := acctest.FromInstanceState(s, resourceName, "user_id")
					compositeId = "users/" + userId + "/customerSecretKeys/" + resId
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", acctest.Optional, acctest.Update, customerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_customer_secret_keys", "test_customer_secret_keys", acctest.Optional, acctest.Update, customerSecretKeyDataSourceRepresentation) +
				compartmentIdVariableStr + CustomerSecretKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", acctest.Optional, acctest.Update, customerSecretKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "customer_secret_keys.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "customer_secret_keys.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "customer_secret_keys.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "customer_secret_keys.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "customer_secret_keys.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "customer_secret_keys.0.user_id"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getCustomerKeyImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"key",
			},
			ResourceName: resourceName,
		},
	})
}

func getCustomerKeyImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("users/" + rs.Primary.Attributes["user_id"] + "/customerSecretKeys/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckIdentityCustomerSecretKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_customer_secret_key" {
			noResourceFound = false
			request := oci_identity.ListCustomerSecretKeysRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			response, err := client.ListCustomerSecretKeys(context.Background(), request)

			if err == nil {
				id := rs.Primary.Attributes["id"]
				for _, item := range response.Items {
					if *item.Id == id {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, that means item is deleted. continue checking next one
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
