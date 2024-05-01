// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RecoveryRecoveryServiceSubnetRequiredOnlyResource = RecoveryRecoveryServiceSubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetRepresentation)

	RecoveryRecoveryServiceSubnetResourceConfig = RecoveryRecoveryServiceSubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetRepresentation)

	RecoveryRecoveryServiceSubnetSingularDataSourceRepresentation = map[string]interface{}{
		"recovery_service_subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id}`},
	}

	RecoveryRecoveryServiceSubnetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryRecoveryServiceSubnetDataSourceFilterRepresentation}}
	RecoveryRecoveryServiceSubnetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id}`}},
	}

	RecoveryRecoveryServiceSubnetRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"subnets":        acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_subnet.id}`}, Update: []string{`${oci_core_subnet.test_subnet.id}`, `${oci_core_subnet.test_subnet1.id}`}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}

	RecoveryRecoveryServiceSubnetRepresentationForSubnetId = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Required, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}

	RecoveryRecoveryServiceSubnetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet1", acctest.Required, acctest.Create, acctest.GetMultipleUpdatedRepresenationCopy([]string{`cidr_block`}, []interface{}{acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`}}, CoreSubnetRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: recovery/default
func TestRecoveryRecoveryServiceSubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryRecoveryServiceSubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_recovery_recovery_service_subnet.test_recovery_service_subnet"
	datasourceName := "data.oci_recovery_recovery_service_subnets.test_recovery_service_subnets"
	singularDatasourceName := "data.oci_recovery_recovery_service_subnet.test_recovery_service_subnet"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryRecoveryServiceSubnetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create, RecoveryRecoveryServiceSubnetRepresentation), "recovery", "recoveryServiceSubnet", t)

	acctest.ResourceTest(t, testAccCheckRecoveryRecoveryServiceSubnetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create, RecoveryRecoveryServiceSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RecoveryRecoveryServiceSubnetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnets.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_recovery_service_subnets", "test_recovery_service_subnets", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subnets.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "recovery_service_subnet_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "recovery_service_subnet_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_service_subnet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RecoveryRecoveryServiceSubnetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: recovery/default
func TestRecoveryRecoveryServiceSubnetResourceWithSubnetId(t *testing.T) {

	httpreplay.SetScenario("TestRecoveryRecoveryServiceSubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_recovery_recovery_service_subnet.test_recovery_service_subnet"
	datasourceName := "data.oci_recovery_recovery_service_subnets.test_recovery_service_subnets"
	singularDatasourceName := "data.oci_recovery_recovery_service_subnet.test_recovery_service_subnet"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryRecoveryServiceSubnetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create, RecoveryRecoveryServiceSubnetRepresentationForSubnetId), "recovery", "recoveryServiceSubnet", t)

	acctest.ResourceTest(t, testAccCheckRecoveryRecoveryServiceSubnetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetRepresentationForSubnetId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create, RecoveryRecoveryServiceSubnetRepresentationForSubnetId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RecoveryRecoveryServiceSubnetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetRepresentationForSubnetId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_recovery_service_subnets", "test_recovery_service_subnets", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Optional, acctest.Update, RecoveryRecoveryServiceSubnetRepresentationForSubnetId),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "recovery_service_subnet_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "recovery_service_subnet_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryRecoveryServiceSubnetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_service_subnet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RecoveryRecoveryServiceSubnetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRecoveryRecoveryServiceSubnetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_recovery_recovery_service_subnet" {
			noResourceFound = false
			request := oci_recovery.GetRecoveryServiceSubnetRequest{}

			tmp := rs.Primary.ID
			request.RecoveryServiceSubnetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")

			response, err := client.GetRecoveryServiceSubnet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_recovery.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RecoveryRecoveryServiceSubnet") {
		resource.AddTestSweepers("RecoveryRecoveryServiceSubnet", &resource.Sweeper{
			Name:         "RecoveryRecoveryServiceSubnet",
			Dependencies: acctest.DependencyGraph["recoveryServiceSubnet"],
			F:            sweepRecoveryRecoveryServiceSubnetResource,
		})
	}
}

func sweepRecoveryRecoveryServiceSubnetResource(compartment string) error {
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()
	recoveryServiceSubnetIds, err := getRecoveryRecoveryServiceSubnetIds(compartment)
	if err != nil {
		return err
	}
	for _, recoveryServiceSubnetId := range recoveryServiceSubnetIds {
		if ok := acctest.SweeperDefaultResourceId[recoveryServiceSubnetId]; !ok {
			deleteRecoveryServiceSubnetRequest := oci_recovery.DeleteRecoveryServiceSubnetRequest{}

			deleteRecoveryServiceSubnetRequest.RecoveryServiceSubnetId = &recoveryServiceSubnetId

			deleteRecoveryServiceSubnetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")
			_, error := databaseRecoveryClient.DeleteRecoveryServiceSubnet(context.Background(), deleteRecoveryServiceSubnetRequest)
			if error != nil {
				fmt.Printf("Error deleting RecoveryServiceSubnet %s %s, It is possible that the resource is already deleted. Please verify manually \n", recoveryServiceSubnetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &recoveryServiceSubnetId, RecoveryRecoveryServiceSubnetSweepWaitCondition, time.Duration(3*time.Minute),
				RecoveryRecoveryServiceSubnetSweepResponseFetchOperation, "recovery", true)
		}
	}
	return nil
}

func getRecoveryRecoveryServiceSubnetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RecoveryServiceSubnetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()

	listRecoveryServiceSubnetsRequest := oci_recovery.ListRecoveryServiceSubnetsRequest{}
	listRecoveryServiceSubnetsRequest.CompartmentId = &compartmentId
	listRecoveryServiceSubnetsRequest.LifecycleState = oci_recovery.ListRecoveryServiceSubnetsLifecycleStateActive
	listRecoveryServiceSubnetsResponse, err := databaseRecoveryClient.ListRecoveryServiceSubnets(context.Background(), listRecoveryServiceSubnetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RecoveryServiceSubnet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, recoveryServiceSubnet := range listRecoveryServiceSubnetsResponse.Items {
		id := *recoveryServiceSubnet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RecoveryServiceSubnetId", id)
	}
	return resourceIds, nil
}

func RecoveryRecoveryServiceSubnetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if recoveryServiceSubnetResponse, ok := response.Response.(oci_recovery.GetRecoveryServiceSubnetResponse); ok {
		return recoveryServiceSubnetResponse.LifecycleState != oci_recovery.LifecycleStateDeleted
	}
	return false
}

func RecoveryRecoveryServiceSubnetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseRecoveryClient().GetRecoveryServiceSubnet(context.Background(), oci_recovery.GetRecoveryServiceSubnetRequest{
		RecoveryServiceSubnetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
