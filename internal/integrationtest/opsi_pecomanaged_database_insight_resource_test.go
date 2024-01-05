// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	pecomanagedDatabaseInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"credential_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: pecomanagedDatabaseInsightCredentialDetailsRepresentation},
		"database_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.dbsystem_database_id}`},
		"database_resource_type":   acctest.Representation{RepType: acctest.Required, Create: `database`},
		"deployment_type":          acctest.Representation{RepType: acctest.Required, Create: `VIRTUAL_MACHINE`},
		"opsi_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.opsi_private_endpoint_id}`},
		"service_name":             acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
		"status":                   acctest.Representation{RepType: acctest.Required, Create: `DISABLED`},
		"entity_source":            acctest.Representation{RepType: acctest.Required, Create: `PE_COMANAGED_DATABASE`, Update: `PE_COMANAGED_DATABASE`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesPDIRepresentation},
	}

	ignoreChangesPDIRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiPecomanagedResourceDatabaseInsight(t *testing.T) {
	httpreplay.SetScenario("TestOpsiPecomanagedResourceDatabaseInsight")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	opsiPrivateEndpointId := utils.GetEnvSettingWithBlankDefault("opsi_private_endpoint_id")
	opsiPrivateEndpointIdVariableStr := fmt.Sprintf("variable \"opsi_private_endpoint_id\" { default = \"%s\" }\n", opsiPrivateEndpointId)

	dbsystemDatabaseId := utils.GetEnvSettingWithBlankDefault("dbsystem_database_id")
	dbsystemDatabaseIdVariableStr := fmt.Sprintf("variable \"dbsystem_database_id\" { default = \"%s\" }\n", dbsystemDatabaseId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)

	userName := utils.GetEnvSettingWithBlankDefault("user_name")
	userNamedVariableStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)

	resourceName := "oci_opsi_database_insight.test_database_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbsystemDatabaseIdVariableStr+opsiPrivateEndpointIdVariableStr+serviceNameVariableStr+secretIdVariableStr+userNamedVariableStr+PecomanagedDatabaseInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, pecomanagedDatabaseInsightRequiredRepresentation), "opsi", "databaseInsightnkedruktrffiiufuegnf"+
		"", t)

	acctest.ResourceTest(t, testAccCheckOpsiDatabaseInsightDestroy, []resource.TestStep{
		// verify Create with Required
		{
			Config: config + compartmentIdVariableStr + dbsystemDatabaseIdVariableStr + opsiPrivateEndpointIdVariableStr + serviceNameVariableStr + secretIdVariableStr + userNamedVariableStr + PecomanagedDatabaseInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", acctest.Required, acctest.Create, pecomanagedDatabaseInsightRequiredRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + PecomanagedDatabaseInsightRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"deployment_type",
				"service_name",
			},
			ResourceName: resourceName,
		},
	})
}
