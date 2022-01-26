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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	idpGroupMappingDataSourceRepresentation = map[string]interface{}{
		"identity_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_identity_provider.test_identity_provider.id}`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: idpGroupMappingDataSourceFilterRepresentation}}
	idpGroupMappingDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_idp_group_mapping.test_idp_group_mapping.id}`}},
	}

	idpGroupMappingRepresentation = map[string]interface{}{
		"group_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.id}`},
		"identity_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_identity_provider.test_identity_provider.id}`},
		"idp_group_name":       acctest.Representation{RepType: acctest.Required, Create: `idpGroupName`, Update: `idpGroupName2`},
	}

	IdpGroupMappingResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, groupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Required, acctest.Create, identityProviderRepresentation) +
		IdentityProviderPropertyVariables
)

// issue-routing-tag: identity/default
func TestIdentityIdpGroupMappingResource_basic(t *testing.T) {
	metadataFile := utils.GetEnvSettingWithBlankDefault("identity_provider_metadata_file")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	httpreplay.SetScenario("TestIdentityIdpGroupMappingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_idp_group_mapping.test_idp_group_mapping"
	datasourceName := "data.oci_identity_idp_group_mappings.test_idp_group_mappings"

	var resId, resId2 string
	var compositeId string

	_, tokenFn := acctest.TokenizeWithHttpReplay("idp_group_mapping")
	IdpGroupMappingResourceDependencies = tokenFn(IdpGroupMappingResourceDependencies, map[string]string{"metadata_file": metadataFile})

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdpGroupMappingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_idp_group_mapping", "test_idp_group_mapping", acctest.Required, acctest.Create, idpGroupMappingRepresentation), "identity", "idpGroupMapping", t)

	acctest.ResourceTest(t, testAccCheckIdentityIdpGroupMappingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdpGroupMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_idp_group_mapping", "test_idp_group_mapping", acctest.Required, acctest.Create, idpGroupMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "idp_group_name", "idpGroupName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					identityProviderId, _ := acctest.FromInstanceState(s, resourceName, "identity_provider_id")
					compositeId = "identityProviders/" + identityProviderId + "/groupMappings/" + resId
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
			Config: config + compartmentIdVariableStr + IdpGroupMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_idp_group_mapping", "test_idp_group_mapping", acctest.Optional, acctest.Update, idpGroupMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "idp_group_name", "idpGroupName2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_idp_group_mappings", "test_idp_group_mappings", acctest.Optional, acctest.Update, idpGroupMappingDataSourceRepresentation) +
				compartmentIdVariableStr + IdpGroupMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_idp_group_mapping", "test_idp_group_mapping", acctest.Optional, acctest.Update, idpGroupMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_id"),

				resource.TestCheckResourceAttr(datasourceName, "idp_group_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.identity_provider_id"),
				resource.TestCheckResourceAttr(datasourceName, "idp_group_mappings.0.idp_group_name", "idpGroupName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getIdpGroupMappingImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getIdpGroupMappingImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("identityProviders/" + rs.Primary.Attributes["identity_provider_id"] + "/groupMappings/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckIdentityIdpGroupMappingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_idp_group_mapping" {
			noResourceFound = false
			request := oci_identity.GetIdpGroupMappingRequest{}

			if value, ok := rs.Primary.Attributes["identity_provider_id"]; ok {
				request.IdentityProviderId = &value
			}

			tmp := rs.Primary.ID
			request.MappingId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetIdpGroupMapping(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.IdpGroupMappingLifecycleStateDeleted): true,
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
