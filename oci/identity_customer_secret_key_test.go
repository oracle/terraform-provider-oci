// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v44/common"
	oci_identity "github.com/oracle/oci-go-sdk/v44/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	customerSecretKeyDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
		"filter":  RepresentationGroup{Required, customerSecretKeyDataSourceFilterRepresentation}}
	customerSecretKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_customer_secret_key.test_customer_secret_key.id}`}},
	}

	customerSecretKeyRepresentation = map[string]interface{}{
		"display_name": Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"user_id":      Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	CustomerSecretKeyResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestIdentityCustomerSecretKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCustomerSecretKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_customer_secret_key.test_customer_secret_key"
	datasourceName := "data.oci_identity_customer_secret_keys.test_customer_secret_keys"

	var resId, resId2 string
	var compositeId string

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+CustomerSecretKeyResourceDependencies+
		generateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", Required, Create, customerSecretKeyRepresentation), "identity", "customerSecretKey", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityCustomerSecretKeyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CustomerSecretKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", Required, Create, customerSecretKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						userId, _ := fromInstanceState(s, resourceName, "user_id")
						compositeId = "users/" + userId + "/customerSecretKeys/" + resId
						log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", Optional, Update, customerSecretKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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
					generateDataSourceFromRepresentationMap("oci_identity_customer_secret_keys", "test_customer_secret_keys", Optional, Update, customerSecretKeyDataSourceRepresentation) +
					compartmentIdVariableStr + CustomerSecretKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_customer_secret_key", "test_customer_secret_key", Optional, Update, customerSecretKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_customer_secret_key" {
			noResourceFound = false
			request := oci_identity.ListCustomerSecretKeysRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
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
