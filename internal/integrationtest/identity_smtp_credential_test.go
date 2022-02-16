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
	smtpCredentialDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: smtpCredentialDataSourceFilterRepresentation}}
	smtpCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_smtp_credential.test_smtp_credential.id}`}},
	}

	smtpCredentialRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"user_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	SmtpCredentialResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentitySmtpCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentitySmtpCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_smtp_credential.test_smtp_credential"
	datasourceName := "data.oci_identity_smtp_credentials.test_smtp_credentials"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SmtpCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", acctest.Required, acctest.Create, smtpCredentialRepresentation), "identity", "smtpCredential", t)

	acctest.ResourceTest(t, testAccCheckIdentitySmtpCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", acctest.Required, acctest.Create, smtpCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "user_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					userId, _ := acctest.FromInstanceState(s, resourceName, "user_id")
					compositeId = "users/" + userId + "/smtpCredentials/" + resId
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
			Config: config + compartmentIdVariableStr + SmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", acctest.Optional, acctest.Update, smtpCredentialRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_smtp_credentials", "test_smtp_credentials", acctest.Optional, acctest.Update, smtpCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + SmtpCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", acctest.Optional, acctest.Update, smtpCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.0.description", "description2"),
				resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.user_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.username"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getSmtpCredentialImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"password",
			},
			ResourceName: resourceName,
		},
	})
}

func getSmtpCredentialImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("users/" + rs.Primary.Attributes["user_id"] + "/smtpCredentials/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckIdentitySmtpCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_smtp_credential" {
			noResourceFound = false
			request := oci_identity.ListSmtpCredentialsRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			response, err := client.ListSmtpCredentials(context.Background(), request)

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
