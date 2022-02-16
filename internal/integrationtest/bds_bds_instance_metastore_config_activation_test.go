// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v58/bds"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	bdsDcatDynamicGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `DCAT metastore group for terraform test`},
		"matching_rule":  acctest.Representation{RepType: acctest.Required, Create: `Any {resource.id='${var.metastore_id}'}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `bds_terraform_test_dynamic_group`},
	}

	bdsApiKeyUserGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Users of this group can create BDS API key which can be used for creating bds metastore config in terraform test`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `bds_terraform_test_api_key_user_group`},
	}

	bdsMetastoreConfigTestPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Policy to allow DCAT metastore access to object store for terraform testing`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `bds_terraform_test_policy`},
		"statements": acctest.Representation{RepType: acctest.Required, Create: []string{
			`allow dynamic-group ${oci_identity_dynamic_group.test_dynamic_group.name} to manage object-family in tenancy`,
			`allow group ${oci_identity_group.test_group.name} to {CATALOG_METASTORE_READ, CATALOG_METASTORE_EXECUTE} in tenancy where target.metastore.id='${var.metastore_id}'`}},
	}

	bdsMetastoreConfigTestUserGroupMembershipRepresentation = map[string]interface{}{
		"group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.id}`},
		"user_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.id}`},
	}

	bdsInstanceMetastoreConfigActivateRepresentation1 = acctest.RepresentationCopyWithNewProperties(bdsInstanceMetastoreConfigRepresentation,
		map[string]interface{}{
			"display_name":     acctest.Representation{RepType: acctest.Required, Create: `metastoreActivation1`},
			"activate_trigger": acctest.Representation{RepType: acctest.Optional, Create: `1`}, // Added during optional param test step
		})

	bdsInstanceMetastoreConfigActivateRepresentation2 = acctest.RepresentationCopyWithNewProperties(bdsInstanceMetastoreConfigRepresentation,
		map[string]interface{}{
			"display_name": acctest.Representation{RepType: acctest.Required, Create: `metastoreActivation2`},
			"depends_on":   acctest.Representation{RepType: acctest.Required, Create: []string{`oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config_activation1`}},
		})

	bdsInstanceOdhWithNatGatewayRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation, map[string]interface{}{
		"network_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceMetastoreConfigNetworkRepresentation},
	})

	bdsInstanceMetastoreConfigNetworkRepresentation = map[string]interface{}{
		"cidr_block":              acctest.Representation{RepType: acctest.Required, Create: `111.112.0.0/16`},
		"is_nat_gateway_required": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceMetastoreConfigResource_activation(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceMetastoreConfigResource_activation")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	activationResourceName1 := "oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config_activation1"
	activationResourceName2 := "oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config_activation2"

	var bdsId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsInstanceMetastoreConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Create, bdsInstanceMetastoreConfigRepresentation), "bds", "bdsInstanceMetastoreConfig", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceMetastoreConfigDestroy, []resource.TestStep{
		// Create resource1 followed by resource2, order ensured with depends_on definition in resource2.
		{
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config_activation1", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigActivateRepresentation1) +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config_activation2", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigActivateRepresentation2) +
				metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(activationResourceName1, "display_name", "metastoreActivation1"),
				resource.TestCheckResourceAttr(activationResourceName2, "display_name", "metastoreActivation2"),

				func(s *terraform.State) (err error) {
					bdsId, _ = acctest.FromInstanceState(s, activationResourceName1, "bds_instance_id")
					return err
				},
			),
		},

		// Add optional field activate_trigger to resource1 in order to activate it back.
		{
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config_activation1", acctest.Optional, acctest.Create, bdsInstanceMetastoreConfigActivateRepresentation1) +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config_activation2", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigActivateRepresentation2) +
				metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(activationResourceName1, "display_name", "metastoreActivation1"),
				resource.TestCheckResourceAttr(activationResourceName2, "display_name", "metastoreActivation2"),
				resource.TestCheckResourceAttr(activationResourceName1, "activate_trigger", "1"),
			),
		},

		// Metastore config 1 was manually activated in previous step, so we can delete metastore config 2 now.
		{
			PreConfig: func() {
				activateLocalMetastore(bdsId) // So that metastore config 1 is deactivated allowing deletion of all test resources
			},
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config_activation1", acctest.Optional, acctest.Create, bdsInstanceMetastoreConfigActivateRepresentation1) +
				metastoreIdVariableStr,
		},
	})
}

func activateLocalMetastore(bdsInstanceId string) {
	bdsClient := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	listBdsMetastoreConfigurationsResponse, _ := bdsClient.ListBdsMetastoreConfigurations(context.Background(), oci_bds.ListBdsMetastoreConfigurationsRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: tfresource.GetRetryPolicy(true, "bds"),
	},
		MetastoreType: oci_bds.BdsMetastoreConfigurationMetastoreTypeLocal,
		BdsInstanceId: &bdsInstanceId,
	})

	var localMetastoreId string
	for _, bdsInstanceMetastoreConfig := range listBdsMetastoreConfigurationsResponse.Items {
		localMetastoreId = *bdsInstanceMetastoreConfig.Id
	}
	var clusterPassword = "V2VsY29tZTE="
	request := oci_bds.ActivateBdsMetastoreConfigurationRequest{}
	request.ClusterAdminPassword = &clusterPassword
	request.BdsInstanceId = &bdsInstanceId
	request.MetastoreConfigId = &localMetastoreId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

	response, _ := bdsClient.ActivateBdsMetastoreConfiguration(context.Background(), request)

	workId := response.OpcWorkRequestId
	waitFunction := acctest.WaitTillCondition(acctest.TestAccProvider, workId, activateLocalMetastoreConfigWorkRequestWaitCondition, 20*time.Minute,
		activateLocalMetastoreConfigWorkRequestFetchOperation, "bds", true)
	waitFunction()
}

func activateLocalMetastoreConfigWorkRequestWaitCondition(response common.OCIOperationResponse) bool {
	if getWorkRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
		return getWorkRequestResponse.Status != oci_bds.OperationStatusSucceeded
	}
	return false
}

func activateLocalMetastoreConfigWorkRequestFetchOperation(client *tf_client.OracleClients, workRequestId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetWorkRequest(context.Background(), oci_bds.GetWorkRequestRequest{
		WorkRequestId: workRequestId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
