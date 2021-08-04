// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_identity "github.com/oracle/oci-go-sdk/v45/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	apiKeyDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
		"filter":  RepresentationGroup{Required, apiKeyDataSourceFilterRepresentation}}
	apiKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_api_key.test_api_key.id}`}},
	}

	apiKeyRepresentation = map[string]interface{}{
		"key_value": Representation{repType: Required, create: `${var.api_key_value}`},
		"user_id":   Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	ApiKeyResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation) + publicKeyVariableStr

	publicKey            = getEnvSettingWithBlankDefault("public_key")
	publicKeyVariableStr = fmt.Sprintf("variable \"api_key_value\" { default = \"%s\" }\n", publicKey)

	publicKeyUpdate            = getEnvSettingWithBlankDefault("public_key_update")
	publicKeyUpdateVariableStr = fmt.Sprintf("variable \"api_key_update_value\" { default = \"%s\" }\n", publicKeyUpdate)
)

// issue-routing-tag: identity/default
func TestIdentityApiKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityApiKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_api_key.test_api_key"
	datasourceName := "data.oci_identity_api_keys.test_api_keys"

	var compositeId, fingerprint string

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ApiKeyResourceDependencies+
		generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Required, Create, apiKeyRepresentation), "identity", "apiKey", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityApiKeyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApiKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Required, Create, apiKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestMatchResourceAttr(resourceName, "key_value", regexp.MustCompile("-----BEGIN PUBL.*")),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApiKeyResourceDependencies,
			},
			// verify create with export
			{
				Config: config + compartmentIdVariableStr + ApiKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Required, Create, apiKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestMatchResourceAttr(resourceName, "key_value", regexp.MustCompile("-----BEGIN PUBL.*")),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						fingerprint, _ = fromInstanceState(s, resourceName, "fingerprint")
						userId, _ := fromInstanceState(s, resourceName, "user_id")
						compositeId = "oci_identity_api_key:users/" + userId + "/apiKeys/" + fingerprint
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

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_api_keys", "test_api_keys", Optional, Update, apiKeyDataSourceRepresentation) +
					compartmentIdVariableStr + ApiKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Optional, Update, apiKeyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "api_keys.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.fingerprint"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.id"),
					resource.TestMatchResourceAttr(datasourceName, "api_keys.0.key_value", regexp.MustCompile("-----BEGIN PUBL.*")),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.user_id"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       getApiKeyImportId(resourceName),
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func getApiKeyImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("oci_identity_api_key:users/" + rs.Primary.Attributes["user_id"] + "/apiKeys/" + rs.Primary.Attributes["fingerprint"]), nil
	}
}

func testAccCheckIdentityApiKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_api_key" {
			noResourceFound = false
			request := oci_identity.ListApiKeysRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			response, err := client.ListApiKeys(context.Background(), request)

			if err == nil {
				fingerprint := rs.Primary.Attributes["fingerprint"]
				for _, item := range response.Items {
					if *item.Fingerprint == fingerprint {
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
