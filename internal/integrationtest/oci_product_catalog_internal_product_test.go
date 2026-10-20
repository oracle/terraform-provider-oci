// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_oci_product_catalog "github.com/oracle/oci-go-sdk/v65/ociproductcatalog"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OciProductCatalogInternalProductRequiredOnlyResource = OciProductCatalogInternalProductResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Required, acctest.Create, OciProductCatalogInternalProductRepresentation)

	OciProductCatalogInternalProductResourceConfig = OciProductCatalogInternalProductResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Optional, acctest.Update, OciProductCatalogInternalProductRepresentation)

	OciProductCatalogInternalProductSingularDataSourceRepresentation = map[string]interface{}{
		"product_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
	}

	OciProductCatalogInternalProductRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":          acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"service_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_announcements_service_service.test_service.name}`},
		"backfill_eligibility": acctest.Representation{RepType: acctest.Optional, Create: `BACKFILL_ELIGIBLE`},
		"is_excluded":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"limits":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: OciProductCatalogInternalProductLimitsRepresentation},
		"meters":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: OciProductCatalogInternalProductMetersRepresentation},
	}
	OciProductCatalogInternalProductLimitsRepresentation = map[string]interface{}{
		"public_limit_name":   acctest.Representation{RepType: acctest.Optional, Create: `publicLimitName`, Update: `publicLimitName2`},
		"public_service_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_announcements_service_service.test_service.name}`},
	}
	OciProductCatalogInternalProductMetersRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}

	OciProductCatalogInternalProductResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_announcements_service_services", "test_services", acctest.Required, acctest.Create, AnnouncementsServiceServiceDataSourceRepresentation)
)

// issue-routing-tag: oci_product_catalog/default
func TestOciProductCatalogInternalProductResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOciProductCatalogInternalProductResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oci_product_catalog_internal_product.test_internal_product"

	singularDatasourceName := "data.oci_oci_product_catalog_internal_product.test_internal_product"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OciProductCatalogInternalProductResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Optional, acctest.Create, OciProductCatalogInternalProductRepresentation), "ociproductcatalog", "internalProduct", t)

	acctest.ResourceTest(t, testAccCheckOciProductCatalogInternalProductDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OciProductCatalogInternalProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Required, acctest.Create, OciProductCatalogInternalProductRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OciProductCatalogInternalProductResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OciProductCatalogInternalProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Optional, acctest.Create, OciProductCatalogInternalProductRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backfill_eligibility", "BACKFILL_ELIGIBLE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_excluded", "false"),
				resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName"),
				resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OciProductCatalogInternalProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OciProductCatalogInternalProductRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backfill_eligibility", "BACKFILL_ELIGIBLE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_excluded", "false"),
				resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName"),
				resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + OciProductCatalogInternalProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Optional, acctest.Update, OciProductCatalogInternalProductRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backfill_eligibility", "BACKFILL_ELIGIBLE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_excluded", "true"),
				resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName2"),
				resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meters.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oci_product_catalog_internal_product", "test_internal_product", acctest.Required, acctest.Create, OciProductCatalogInternalProductSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OciProductCatalogInternalProductResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_excluded", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "skus.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_launched"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ready"),
			),
		},
		// verify resource import
		{
			Config:            config + OciProductCatalogInternalProductRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"backfill_eligibility",
				"compartment_id",
				"limits",
				"meters",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOciProductCatalogInternalProductDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ProductInternalClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oci_product_catalog_internal_product" {
			noResourceFound = false
			request := oci_oci_product_catalog.GetInternalProductRequest{}

			if value, ok := rs.Primary.Attributes["product_id"]; ok {
				request.ProductId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oci_product_catalog")

			_, err := client.GetInternalProduct(context.Background(), request)

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
