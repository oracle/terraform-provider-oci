// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
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
	IdentityCompartmentRequiredOnlyResource = IdentityCompartmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Required, acctest.Create, IdentityCompartmentRepresentation)

	IdentityCompartmentResourceConfig = IdentityCompartmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Update, IdentityCompartmentRepresentation)

	IdentityIdentityCompartmentSingularDataSourceRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_compartment.test_compartment.id}`},
	}

	IdentityIdentityCompartmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ANY`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `Network`, Update: `name2`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityCompartmentDataSourceFilterRepresentation}}
	IdentityCompartmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_compartment.test_compartment.id}`}},
	}

	IdentityCompartmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `For network components`, Update: `description2`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `Network`, Update: `name2`},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	IdentityCompartmentResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityCompartmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCompartmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_identity_compartment.test_compartment"
	datasourceName := "data.oci_identity_compartments.test_compartments"
	singularDatasourceName := "data.oci_identity_compartment.test_compartment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityCompartmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Create, IdentityCompartmentRepresentation), "identity", "compartment", t)

	acctest.ResourceTest(t, testAccCheckIdentityCompartmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Required, acctest.Create, IdentityCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
				resource.TestCheckResourceAttr(resourceName, "name", "Network"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Create, IdentityCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Network"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IdentityCompartmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "For network components"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Network"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Update, IdentityCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_compartments", "test_compartments", acctest.Optional, acctest.Update, IdentityIdentityCompartmentDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Optional, acctest.Update, IdentityCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ANY"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "compartments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartments.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartments.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "compartments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "compartments.0.name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "compartments.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Required, acctest.Create, IdentityIdentityCompartmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityCompartmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityCompartmentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"enable_delete",
				// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
				"is_accessible",
			},
			ResourceName: resourceName,
		},
		// restore name of compartment
		{
			Config: config + compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Required, acctest.Create, IdentityCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "Network"),
			),
		},
		// verify error on existing compartment name where automatic import fails due to enable_delete
		{
			Config: config + compartmentIdVariableStr + IdentityCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "test_compartment", acctest.Required, acctest.Create, IdentityCompartmentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_compartment", "name2", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IdentityCompartmentRepresentation, map[string]interface{}{
						"enable_delete": acctest.Representation{RepType: acctest.Required, Create: `true`}})),
			ExpectError: regexp.MustCompile("If you intended to manage an existing compartment, use terraform import instead."),
		},
	})
}

func testAccCheckIdentityCompartmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_compartment" {
			noResourceFound = false
			request := oci_identity.GetCompartmentRequest{}

			tmp := rs.Primary.ID
			request.CompartmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetCompartment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.CompartmentLifecycleStateDeleted): true, // target state when delete_enabled = true
					string(oci_identity.CompartmentLifecycleStateActive):  true, // target state when delete_enabled = false or ""
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
