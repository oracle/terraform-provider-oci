// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityIdentityDbCredentialDataSourceRepresentation = map[string]interface{}{
		"user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
		"name":    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":  acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDbCredentialDataSourceFilterRepresentation}}
	IdentityDbCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_db_credential.test_db_credential.id}`}},
	}

	IdentityDbCredentialRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `description`},
		"password":    acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#112`},
		"user_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	IdentityDbCredentialResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityDbCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDbCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_db_credential.test_db_credential"
	datasourceName := "data.oci_identity_db_credentials.test_db_credentials"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDbCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_db_credential", "test_db_credential", acctest.Required, acctest.Create, IdentityDbCredentialRepresentation), "identity", "dbCredential", t)

	acctest.ResourceTest(t, testAccCheckIdentityDbCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_db_credential", "test_db_credential", acctest.Required, acctest.Create, IdentityDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#112"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_db_credentials", "test_db_credentials", acctest.Optional, acctest.Update, IdentityIdentityDbCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDbCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_db_credential", "test_db_credential", acctest.Optional, acctest.Update, IdentityDbCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

				resource.TestCheckResourceAttr(datasourceName, "db_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_credentials.0.description", "description"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_credentials.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_credentials.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_credentials.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_credentials.0.user_id"),
			),
		},
	})
}

func testAccCheckIdentityDbCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_db_credential" {
			noResourceFound = false
			request := oci_identity.ListDbCredentialsRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")
			response, err := client.ListDbCredentials(context.Background(), request)
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
