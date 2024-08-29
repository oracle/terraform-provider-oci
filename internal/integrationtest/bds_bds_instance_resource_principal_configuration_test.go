// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	//"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstanceResourcePrincipalConfigurationRequiredOnlyResource = BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Required, acctest.Create, BdsBdsInstanceResourcePrincipalConfigurationRepresentation)

	BdsBdsInstanceResourcePrincipalConfigurationResourceConfig = BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceResourcePrincipalConfigurationRepresentation)

	BdsBdsInstanceResourcePrincipalConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"resource_principal_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration.id}`},
	}

	BdsBdsInstanceResourcePrincipalConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceResourcePrincipalConfigurationDataSourceFilterRepresentation}}
	BdsBdsInstanceResourcePrincipalConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration.id}`}},
	}

	BdsBdsInstanceResourcePrincipalConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password":                    acctest.Representation{RepType: acctest.Required, Create: `clusterAdminPassword`},
		"display_name":                              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"session_token_life_span_duration_in_hours": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `11`},
		"force_refresh_resource_principal_trigger":  acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `1`},
	}

	//bdsinstanceId            = utils.GetEnvSettingWithBlankDefault("bdsinstance_ocid")
	//bdsinstanceIdVariableStr = fmt.Sprintf("variable \"bdsinstance_id\" { default = \"%s\" }\n", bdsinstanceId)

	BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceResourcePrincipalConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceResourcePrincipalConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	//bdsinstanceId := utils.GetEnvSettingWithBlankDefault("bdsinstance_ocid")
	//bdsinstanceIdVariableStr := fmt.Sprintf("variable \"bdsinstance_id\" { default = \"%s\" }\n", bdsinstanceId)

	resourceName := "oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration"
	datasourceName := "data.oci_bds_bds_instance_resource_principal_configurations.test_bds_instance_resource_principal_configurations"
	singularDatasourceName := "data.oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceResourcePrincipalConfigurationRepresentation), "bds", "bdsInstanceResourcePrincipalConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Required, acctest.Create, BdsBdsInstanceResourcePrincipalConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "clusterAdminPassword"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceResourcePrincipalConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "clusterAdminPassword"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "session_token_life_span_duration_in_hours", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceResourcePrincipalConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "clusterAdminPassword"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "session_token_life_span_duration_in_hours", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configurations", "test_bds_instance_resource_principal_configurations", acctest.Optional, acctest.Update, BdsBdsInstanceResourcePrincipalConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceResourcePrincipalConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "resource_principal_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_principal_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.time_token_expiry"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.time_token_refreshed"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_principal_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_resource_principal_configuration", "test_bds_instance_resource_principal_configuration", acctest.Required, acctest.Create, BdsBdsInstanceResourcePrincipalConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsInstanceResourcePrincipalConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_principal_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_token_life_span_duration_in_hours", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_token_expiry"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_token_refreshed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsBdsInstanceResourcePrincipalConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getBdsResourcePrincipalConfigurationCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"force_refresh_resource_principal_trigger",
			},
			ResourceName: resourceName,
		},
	})
}

func getBdsResourcePrincipalConfigurationCompositeId(resourceName string) resource.ImportStateIdFunc {

	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/resourcePrincipalConfigurations/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}

}
