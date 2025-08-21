// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeAttributeSetRequiredOnlyResource = DataSafeAttributeSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Required, acctest.Create, DataSafeAttributeSetRepresentation)

	DataSafeAttributeSetResourceConfig = DataSafeAttributeSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Update, DataSafeAttributeSetRepresentation)

	DataSafeAttributeSetSingularDataSourceRepresentation = map[string]interface{}{
		"attribute_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_attribute_set.test_attribute_set.id}`},
	}

	DataSafeAttributeSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"attribute_set_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_attribute_set.test_attribute_set.id}`},
		"attribute_set_type":        acctest.Representation{RepType: acctest.Optional, Create: `IP_ADDRESS`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"in_use":                    acctest.Representation{RepType: acctest.Optional, Create: `NO`},
		"is_user_defined":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeAttributeSetDataSourceFilterRepresentation}}
	DataSafeAttributeSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_attribute_set.test_attribute_set.id}`}},
	}

	DataSafeAttributeSetRepresentation = map[string]interface{}{
		"attribute_set_type":   acctest.Representation{RepType: acctest.Required, Create: `IP_ADDRESS`},
		"attribute_set_values": acctest.Representation{RepType: acctest.Required, Create: []string{`192.168.11.0`}, Update: []string{`192.168.11.1`}},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `IP addresses - AttributeSet`, Update: `displayName2`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `Attribute set for IP addresses`, Update: `description2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAttributeSetRep},
	}

	ignoreAttributeSetRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `system_tags`}},
	}

	DataSafeAttributeSetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAttributeSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAttributeSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_attribute_set.test_attribute_set"
	datasourceName := "data.oci_data_safe_attribute_sets.test_attribute_sets"
	singularDatasourceName := "data.oci_data_safe_attribute_set.test_attribute_set"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeAttributeSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Create, DataSafeAttributeSetRepresentation), "datasafe", "attributeSet", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAttributeSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAttributeSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Required, acctest.Create, DataSafeAttributeSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "attribute_set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "IP addresses - AttributeSet"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAttributeSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeAttributeSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Create, DataSafeAttributeSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "attribute_set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Attribute set for IP addresses"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "IP addresses - AttributeSet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeAttributeSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeAttributeSetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "attribute_set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Attribute set for IP addresses"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "IP addresses - AttributeSet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + DataSafeAttributeSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Update, DataSafeAttributeSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "attribute_set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_attribute_sets", "test_attribute_sets", acctest.Optional, acctest.Update, DataSafeAttributeSetDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAttributeSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Optional, acctest.Update, DataSafeAttributeSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "attribute_set_id"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_user_defined", "false"),

				resource.TestCheckResourceAttr(datasourceName, "attribute_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_set_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_attribute_set", "test_attribute_set", acctest.Required, acctest.Create, DataSafeAttributeSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAttributeSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "attribute_set_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_set_type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_set_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "in_use"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_user_defined"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeAttributeSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeAttributeSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_attribute_set" {
			noResourceFound = false
			request := oci_data_safe.GetAttributeSetRequest{}

			tmp := rs.Primary.ID
			request.AttributeSetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetAttributeSet(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataSafeAttributeSet") {
		resource.AddTestSweepers("DataSafeAttributeSet", &resource.Sweeper{
			Name:         "DataSafeAttributeSet",
			Dependencies: acctest.DependencyGraph["attributeSet"],
			F:            sweepDataSafeAttributeSetResource,
		})
	}
}

func sweepDataSafeAttributeSetResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	attributeSetIds, err := getDataSafeAttributeSetIds(compartment)
	if err != nil {
		return err
	}
	for _, attributeSetId := range attributeSetIds {
		if ok := acctest.SweeperDefaultResourceId[attributeSetId]; !ok {
			deleteAttributeSetRequest := oci_data_safe.DeleteAttributeSetRequest{}

			deleteAttributeSetRequest.AttributeSetId = &attributeSetId

			deleteAttributeSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteAttributeSet(context.Background(), deleteAttributeSetRequest)
			if error != nil {
				fmt.Printf("Error deleting AttributeSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", attributeSetId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeAttributeSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AttributeSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listAttributeSetsRequest := oci_data_safe.ListAttributeSetsRequest{}
	listAttributeSetsRequest.CompartmentId = &compartmentId
	listAttributeSetsResponse, err := dataSafeClient.ListAttributeSets(context.Background(), listAttributeSetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AttributeSet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, attributeSet := range listAttributeSetsResponse.Items {
		id := *attributeSet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AttributeSetId", id)
	}
	return resourceIds, nil
}
