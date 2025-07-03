// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAzureVaultAssociationRequiredOnlyResource = DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationRepresentation)

	DbmulticloudOracleDbAzureVaultAssociationResourceConfig = DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultAssociationRepresentation)

	DbmulticloudOracleDbAzureVaultAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_azure_vault_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association.id}`},
	}

	DbmulticloudOracleDbAzureVaultAssociationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `TestDbAzureVaultAssociation`},
		"oracle_db_azure_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"oracle_db_azure_vault_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id}`},
		// "resource_type":                acctest.Representation{RepType: acctest.Required, Create: `VAULTS`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAzureVaultAssociationDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAzureVaultAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association.id}`}},
	}

	DbmulticloudOracleDbAzureVaultAssociationRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `TestDbAzureVaultAssociation`},
		"oracle_db_azure_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"oracle_db_azure_vault_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id}`},
		// "resource_type":                acctest.Representation{RepType: acctest.Required, Create: `VAULTS`},
		// "defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DbmulticloudOracleDbAzureVaultAssociationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureVaultAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureVaultAssociationResource_basic")
	defer httpreplay.SaveScenario()
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Begin TestDbmulticloudOracleDbAzureVaultAssociationResource_basic>>>>>>>>>>")
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association"
	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_vault_associations.test_oracle_db_azure_vault_associations"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAzureVaultAssociationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationRepresentation), "dbmulticloud", "oracleDbAzureVaultAssociation", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAzureVaultAssociationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAzureVaultAssociation"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_association_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAzureVaultAssociation"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttrSet(resourceName, "display_name", "TestDbAzureVaultAssociation"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_association_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),

				func(s *terraform.State) (err error) {
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Begin OF  resourcediscovery>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Print resId>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					fmt.Println(resId)
					if err != nil {
						fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> error did not complete resourcediscovery >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
						fmt.Println(err)
						return err
					}
					/**
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> error did not complete resourcediscovery >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
							return errExport
						}
					}
					**/
					fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END OF  resourcediscovery>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAzureVaultAssociationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				// resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAzureVaultAssociation"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_association_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAzureVaultAssociation"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_association_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_associations", "test_oracle_db_azure_vault_associations", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TestDbAzureVaultAssociation"),
				// resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				// resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_association_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
				// resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_vault_association_summary_collection.#", "1"),
				// resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_vault_association_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_association", "test_oracle_db_azure_vault_association", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_vault_association_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestDbAzureVaultAssociation"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_resource_accessible"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modification"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				// resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_vault_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAzureVaultAssociationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END OF TestDbmulticloudOracleDbAzureVaultAssociationResource_basic >>>>>>>>>>")
}

func testAccCheckDbmulticloudOracleDbAzureVaultAssociationDestroy(s *terraform.State) error {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Begin of testAccCheckDbmulticloudOracleDbAzureVaultAssociationDestroy>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OracleDbAzureVaultAssociationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_azure_vault_association" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAzureVaultAssociationRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAzureVaultAssociationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAzureVaultAssociation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAzureVaultAssociationLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Got correct status, >>>>>>>>>> ", response.LifecycleState)
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
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ENF OF  testAccCheckDbmulticloudOracleDbAzureVaultAssociationDestroy>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAzureVaultAssociation") {
		resource.AddTestSweepers("DbmulticloudOracleDbAzureVaultAssociation", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAzureVaultAssociation",
			Dependencies: acctest.DependencyGraph["oracleDbAzureVaultAssociation"],
			F:            sweepDbmulticloudOracleDbAzureVaultAssociationResource,
		})
	} else {
		fmt.Println("DbmulticloudOracleDbAzureVaultAssociation is excluded from sweeper.")
	}
}

func sweepDbmulticloudOracleDbAzureVaultAssociationResource(compartment string) error {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BEGIN OF sweepDbmulticloudOracleDbAzureVaultAssociationResource>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	oracleDbAzureVaultAssociationClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultAssociationClient()
	oracleDbAzureVaultAssociationIds, err := getDbmulticloudOracleDbAzureVaultAssociationIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAzureVaultAssociationId := range oracleDbAzureVaultAssociationIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAzureVaultAssociationId]; !ok {
			deleteOracleDbAzureVaultAssociationRequest := oci_dbmulticloud.DeleteOracleDbAzureVaultAssociationRequest{}

			deleteOracleDbAzureVaultAssociationRequest.OracleDbAzureVaultAssociationId = &oracleDbAzureVaultAssociationId

			deleteOracleDbAzureVaultAssociationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := oracleDbAzureVaultAssociationClient.DeleteOracleDbAzureVaultAssociation(context.Background(), deleteOracleDbAzureVaultAssociationRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAzureVaultAssociation %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAzureVaultAssociationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAzureVaultAssociationId, DbmulticloudOracleDbAzureVaultAssociationSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAzureVaultAssociationSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END OF sweepDbmulticloudOracleDbAzureVaultAssociationResource>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	return nil
}

func getDbmulticloudOracleDbAzureVaultAssociationIds(compartment string) ([]string, error) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Begin OF  getDbmulticloudOracleDbAzureVaultAssociationIds>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAzureVaultAssociationId")
	if ids != nil {
		return ids, nil
	}

	var resourceIds []string
	compartmentId := `${var.compartment_id}`
	oracleDbAzureVaultAssociationClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultAssociationClient()
	listOracleDbAzureVaultAssociationsRequest := oci_dbmulticloud.ListOracleDbAzureVaultAssociationsRequest{}
	listOracleDbAzureVaultAssociationsRequest.CompartmentId = &compartmentId
	listOracleDbAzureVaultAssociationsRequest.LifecycleState = oci_dbmulticloud.OracleDbAzureVaultAssociationLifecycleStateActive
	listOracleDbAzureVaultAssociationsResponse, err := oracleDbAzureVaultAssociationClient.ListOracleDbAzureVaultAssociations(context.Background(), listOracleDbAzureVaultAssociationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAzureVaultAssociation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAzureVaultAssociation := range listOracleDbAzureVaultAssociationsResponse.Items {
		id := *oracleDbAzureVaultAssociation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAzureVaultAssociationId", id)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> End OF  getDbmulticloudOracleDbAzureVaultAssociationIds>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	return resourceIds, nil

}

func DbmulticloudOracleDbAzureVaultAssociationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAzureVaultAssociationResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAzureVaultAssociationResponse); ok {
		return oracleDbAzureVaultAssociationResponse.LifecycleState != oci_dbmulticloud.OracleDbAzureVaultAssociationLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAzureVaultAssociationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OracleDbAzureVaultAssociationClient().GetOracleDbAzureVaultAssociation(context.Background(), oci_dbmulticloud.GetOracleDbAzureVaultAssociationRequest{
		OracleDbAzureVaultAssociationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
