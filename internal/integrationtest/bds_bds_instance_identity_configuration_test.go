// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NamespaceSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
	BdsBdsInstanceIdentityConfigurationRequiredOnlyResource = BdsBdsInstanceIdentityConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Required, acctest.Create, BdsBdsInstanceIdentityConfigurationRepresentation)

	BdsBdsInstanceIdentityConfigurationResourceConfig = BdsBdsInstanceIdentityConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceIdentityConfigurationRepresentation)

	BdsBdsInstanceIdentityConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"identity_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_identity_configuration.test_bds_instance_identity_configuration.id}`},
	}

	BdsBdsInstanceIdentityConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `identityDomainConfig`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceIdentityConfigurationDataSourceFilterRepresentation}}
	BdsBdsInstanceIdentityConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_identity_configuration.test_bds_instance_identity_configuration.id}`}},
	}

	BdsBdsInstanceIdentityConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":                              acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"confidential_application_id":                  acctest.Representation{RepType: acctest.Required, Create: `674459b3af89486da22ec1d2da61708f`},
		"display_name":                                 acctest.Representation{RepType: acctest.Required, Create: `identityDomainConfig`},
		"identity_domain_id":                           acctest.Representation{RepType: acctest.Required, Create: `${var.identity_domain_id}`},
		"iam_user_sync_configuration_details":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsInstanceIdentityConfigurationIamUserSyncConfigurationDetailsRepresentation},
		"upst_configuration_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsInstanceIdentityConfigurationUpstConfigurationDetailsRepresentation},
		"activate_iam_user_sync_configuration_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"activate_upst_configuration_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"refresh_confidential_application_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"refresh_upst_token_exchange_keytab_trigger":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	BdsBdsInstanceIdentityConfigurationIamUserSyncConfigurationDetailsRepresentation = map[string]interface{}{
		"is_posix_attributes_addition_required": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	BdsBdsInstanceIdentityConfigurationUpstConfigurationDetailsRepresentation = map[string]interface{}{
		"master_encryption_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.master_encryption_key_id}`},
		"vault_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}

	BdsBdsInstanceIdentityConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceIdentityConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceIdentityConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	identityDomainId := utils.GetEnvSettingWithBlankDefault("identity_domain_ocid")
	identityDomainIdVariableStr := fmt.Sprintf("variable \"identity_domain_id\" { default = \"%s\" }\n", identityDomainId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	masterEncryptionKeyId := utils.GetEnvSettingWithBlankDefault("master_encryption_key_ocid")
	masterEncryptionKeyIdVariableStr := fmt.Sprintf("variable \"master_encryption_key_id\" { default = \"%s\" }\n", masterEncryptionKeyId)

	resourceName := "oci_bds_bds_instance_identity_configuration.test_bds_instance_identity_configuration"
	datasourceName := "data.oci_bds_bds_instance_identity_configurations.test_bds_instance_identity_configurations"
	singularDatasourceName := "data.oci_bds_bds_instance_identity_configuration.test_bds_instance_identity_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+BdsBdsInstanceIdentityConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceIdentityConfigurationRepresentation), "bds", "bdsInstanceIdentityConfiguration", t)
	//fmt.Printf(config + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceIdentityConfigurationRepresentation))
	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceIdentityConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + identityDomainIdVariableStr + vaultIdVariableStr + masterEncryptionKeyIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Required, acctest.Create, BdsBdsInstanceIdentityConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "confidential_application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + identityDomainIdVariableStr + vaultIdVariableStr + masterEncryptionKeyIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceIdentityConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "confidential_application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttr(resourceName, "iam_user_sync_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "iam_user_sync_configuration_details.0.is_posix_attributes_addition_required", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "upst_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "upst_configuration_details.0.master_encryption_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "upst_configuration_details.0.vault_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + identityDomainIdVariableStr + vaultIdVariableStr + masterEncryptionKeyIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceIdentityConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "confidential_application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttr(resourceName, "iam_user_sync_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "iam_user_sync_configuration_details.0.is_posix_attributes_addition_required", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "identity_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "upst_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "upst_configuration_details.0.master_encryption_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "upst_configuration_details.0.vault_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_identity_configurations", "test_bds_instance_identity_configurations", acctest.Optional, acctest.Update, BdsBdsInstanceIdentityConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + identityDomainIdVariableStr + vaultIdVariableStr + masterEncryptionKeyIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceIdentityConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "identity_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_configurations.0.display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_configurations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_configurations.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_identity_configuration", "test_bds_instance_identity_configuration", acctest.Required, acctest.Create, BdsBdsInstanceIdentityConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + identityDomainIdVariableStr + vaultIdVariableStr + masterEncryptionKeyIdVariableStr + BdsBdsInstanceIdentityConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "identityDomainConfig"),
				resource.TestCheckResourceAttr(singularDatasourceName, "iam_user_sync_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "upst_configuration.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsBdsInstanceIdentityConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getBdsIdentityConfigurationCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"iam_user_sync_configuration_details",
				"upst_configuration_details",
				"activate_iam_user_sync_configuration_trigger",
				"activate_upst_configuration_trigger",
				"refresh_confidential_application_trigger",
				"refresh_upst_token_exchange_keytab_trigger",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceIdentityConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_identity_configuration" {
			noResourceFound = false
			request := oci_bds.GetIdentityConfigurationRequest{}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.IdentityConfigurationId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetIdentityConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.IdentityConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceIdentityConfiguration") {
		resource.AddTestSweepers("BdsBdsInstanceIdentityConfiguration", &resource.Sweeper{
			Name:         "BdsBdsInstanceIdentityConfiguration",
			Dependencies: acctest.DependencyGraph["bdsInstanceIdentityConfiguration"],
			F:            sweepBdsBdsInstanceIdentityConfigurationResource,
		})
	}
}

func sweepBdsBdsInstanceIdentityConfigurationResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceIdentityConfigurationIds, err := getBdsBdsInstanceIdentityConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceIdentityConfigurationId := range bdsInstanceIdentityConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceIdentityConfigurationId]; !ok {
			deleteIdentityConfigurationRequest := oci_bds.DeleteIdentityConfigurationRequest{}

			deleteIdentityConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteIdentityConfiguration(context.Background(), deleteIdentityConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceIdentityConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceIdentityConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceIdentityConfigurationId, BdsBdsInstanceIdentityConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsInstanceIdentityConfigurationSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceIdentityConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceIdentityConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listIdentityConfigurationsRequest := oci_bds.ListIdentityConfigurationsRequest{}
	listIdentityConfigurationsRequest.CompartmentId = &compartmentId

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceIdentityConfiguration resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listIdentityConfigurationsRequest.BdsInstanceId = &bdsInstanceId

		listIdentityConfigurationsRequest.LifecycleState = oci_bds.IdentityConfigurationLifecycleStateActive
		listIdentityConfigurationsResponse, err := bdsClient.ListIdentityConfigurations(context.Background(), listIdentityConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceIdentityConfiguration list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceIdentityConfiguration := range listIdentityConfigurationsResponse.Items {
			id := *bdsInstanceIdentityConfiguration.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceIdentityConfigurationId", id)
		}

	}
	return resourceIds, nil
}

func BdsBdsInstanceIdentityConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceIdentityConfigurationResponse, ok := response.Response.(oci_bds.GetIdentityConfigurationResponse); ok {
		return bdsInstanceIdentityConfigurationResponse.LifecycleState != oci_bds.IdentityConfigurationLifecycleStateDeleted
	}
	return false
}

func BdsBdsInstanceIdentityConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetIdentityConfiguration(context.Background(), oci_bds.GetIdentityConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getBdsIdentityConfigurationCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/identityConfigurations/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
