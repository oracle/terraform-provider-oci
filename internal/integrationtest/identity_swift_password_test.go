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
	swiftPasswordDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: swiftPasswordDataSourceFilterRepresentation}}
	swiftPasswordDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_swift_password.test_swift_password.id}`}},
	}

	swiftPasswordRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"user_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	SwiftPasswordResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentitySwiftPasswordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentitySwiftPasswordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_swift_password.test_swift_password"
	datasourceName := "data.oci_identity_swift_passwords.test_swift_passwords"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SwiftPasswordResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_swift_password", "test_swift_password", acctest.Required, acctest.Create, swiftPasswordRepresentation), "identity", "swiftPassword", t)

	acctest.ResourceTest(t, testAccCheckIdentitySwiftPasswordDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SwiftPasswordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_swift_password", "test_swift_password", acctest.Required, acctest.Create, swiftPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					userId, _ := acctest.FromInstanceState(s, resourceName, "user_id")
					compositeId = "users/" + userId + "/swiftPasswords/" + resId
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
			Config: config + compartmentIdVariableStr + SwiftPasswordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_swift_password", "test_swift_password", acctest.Optional, acctest.Update, swiftPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_swift_passwords", "test_swift_passwords", acctest.Optional, acctest.Update, swiftPasswordDataSourceRepresentation) +
				compartmentIdVariableStr + SwiftPasswordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_swift_password", "test_swift_password", acctest.Optional, acctest.Update, swiftPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "passwords.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "passwords.0.description", "description2"),
				resource.TestCheckResourceAttrSet(datasourceName, "passwords.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "passwords.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "passwords.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "passwords.0.user_id"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getSwiftPasswordImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"password",
			},
			ResourceName: resourceName,
		},
	})
}

func getSwiftPasswordImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("users/" + rs.Primary.Attributes["user_id"] + "/swiftPasswords/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckIdentitySwiftPasswordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_swift_password" {
			noResourceFound = false
			request := oci_identity.ListSwiftPasswordsRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			response, err := client.ListSwiftPasswords(context.Background(), request)

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
