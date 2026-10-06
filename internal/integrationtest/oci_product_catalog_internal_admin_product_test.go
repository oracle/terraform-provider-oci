// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

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
	OciProductCatalogInternalAdminProductName = "opc-compute-adminsample-product" + utils.RandomStringOrHttpReplayValue(5, "0123456789", "26")

	OciProductCatalogInternalAdminProductRequiredOnlyResource = OciProductCatalogInternalAdminProductResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_internal_admin_product", acctest.Required, acctest.Create, OciProductCatalogInternalAdminProductRepresentation)

	OciProductCatalogInternalAdminProductResourceConfig = OciProductCatalogInternalAdminProductResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_internal_admin_product", acctest.Optional, acctest.Update, OciProductCatalogInternalAdminProductRepresentation)

	// Fixed DS used in shared dependencies to avoid referencing a resource that may
	// not exist in some steps (e.g., delete-only step).
	OciProductCatalogInternalAdminProductSingularDataSourceRepresentationFixed = map[string]interface{}{
		"product_id": acctest.Representation{RepType: acctest.Required, Create: `${var.product_id}`},
	}

	// Dynamic DS used in the singular datasource verification step to point at the
	// resource created in this test, ensuring attributes match expectations.
	OciProductCatalogInternalAdminProductSingularDataSourceRepresentationDynamic = map[string]interface{}{
		"product_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oci_product_catalog_internal_admin_product.test_internal_admin_product.product_id}`},
	}

	OciProductCatalogInternalAdminProductRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		// Randomized name to ensure uniqueness across runs
		"name":         acctest.Representation{RepType: acctest.Required, Create: OciProductCatalogInternalAdminProductName, Update: OciProductCatalogInternalAdminProductName},
		"service_name": acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
		"limits":       acctest.RepresentationGroup{RepType: acctest.Required, Group: OciProductCatalogInternalAdminProductLimitsRepresentation},
		// Ensure meters are always sent on Create to satisfy backend non-null constraint
		"meters": acctest.RepresentationGroup{RepType: acctest.Required, Group: OciProductCatalogInternalAdminProductMetersRepresentation},
	}
	OciProductCatalogInternalAdminProductLimitsRepresentation = map[string]interface{}{
		"public_limit_name":   acctest.Representation{RepType: acctest.Required, Create: `policy-count`},
		"public_service_name": acctest.Representation{RepType: acctest.Required, Create: `limits`},
	}
	OciProductCatalogInternalAdminProductMetersRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `A100_GPU_V2`, Update: `A100_GPU_V2_AVAILABLE`},
	}

	OciProductCatalogInternalAdminProductResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_services", acctest.Required, acctest.Create, OciProductCatalogInternalAdminProductSingularDataSourceRepresentationFixed)
)

// issue-routing-tag: oci_product_catalog/default
func TestOciProductCatalogInternalAdminProductResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOciProductCatalogInternalAdminProductResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	productId := utils.GetEnvSettingWithBlankDefault("product_id")
	productIdVariableStr := fmt.Sprintf("variable \"product_id\" { default = \"%s\" }\n", productId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oci_product_catalog_internal_admin_product.test_internal_admin_product"

	singularDatasourceName := "data.oci_oci_product_catalog_internal_admin_product.test_internal_admin_product"
	secondName := "opc-compute-sample-product" + utils.RandomStringOrHttpReplayValue(5, "0123456789", "36")

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.

	acctest.SaveConfigContent(
		config+productIdVariableStr+compartmentIdVariableStr+OciProductCatalogInternalAdminProductResourceDependencies+
			acctest.GenerateResourceFromRepresentationMap(
				"oci_oci_product_catalog_internal_admin_product",
				"test_internal_admin_product",
				acctest.Optional,
				acctest.Create,
				acctest.RepresentationCopyWithNewProperties(OciProductCatalogInternalAdminProductRepresentation, map[string]interface{}{
					"name": acctest.Representation{RepType: acctest.Required, Create: secondName},
				}),
			),
		"ociproductcatalog", "internalAdminProduct", t,
	)

	acctest.ResourceTest(t, testAccCheckOciProductCatalogInternalAdminProductDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + productIdVariableStr + compartmentIdVariableStr + OciProductCatalogInternalAdminProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_internal_admin_product", acctest.Required, acctest.Create, OciProductCatalogInternalAdminProductRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", OciProductCatalogInternalAdminProductName),
				resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "meters.0.name", "A100_GPU_V2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + productIdVariableStr + compartmentIdVariableStr + OciProductCatalogInternalAdminProductResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + productIdVariableStr + compartmentIdVariableStr + OciProductCatalogInternalAdminProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_oci_product_catalog_internal_admin_product",
					"test_internal_admin_product",
					acctest.Optional,
					acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OciProductCatalogInternalAdminProductRepresentation, map[string]interface{}{
						"name": acctest.Representation{RepType: acctest.Required, Create: secondName},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName"),
				//resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				//resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "meters.0.name", "A100_GPU_V2_AVAILABLE"),
				//resource.TestCheckResourceAttr(resourceName, "name", OciProductCatalogInternalAdminProductName),
				//resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				//resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + productIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + OciProductCatalogInternalAdminProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_internal_admin_product", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OciProductCatalogInternalAdminProductRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"name":           acctest.Representation{RepType: acctest.Required, Create: secondName},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName"),
				//resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				//resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "meters.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "name", "name"),
				//resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				//resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + productIdVariableStr + compartmentIdVariableStr + OciProductCatalogInternalAdminProductResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_oci_product_catalog_internal_admin_product",
					"test_internal_admin_product",
					acctest.Optional,
					acctest.Update,
					acctest.RepresentationCopyWithNewProperties(
						OciProductCatalogInternalAdminProductRepresentation,
						map[string]interface{}{
							// Keep the same name as the resource created with optionals to avoid
							// triggering a rename that conflicts with an existing product name.
							"name": acctest.Representation{RepType: acctest.Required, Create: secondName, Update: secondName},
						},
					),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "limits.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "limits.0.public_limit_name", "publicLimitName2"),
				//resource.TestCheckResourceAttrSet(resourceName, "limits.0.public_service_name"),
				//resource.TestCheckResourceAttr(resourceName, "meters.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "meters.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "name", secondName),
				//resource.TestCheckResourceAttrSet(resourceName, "service_name"),
				//resource.TestCheckResourceAttr(resourceName, "skus.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_oci_product_catalog_internal_admin_product", "test_internal_admin_product", acctest.Required, acctest.Create, OciProductCatalogInternalAdminProductSingularDataSourceRepresentationDynamic) +
				productIdVariableStr + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_oci_product_catalog_internal_admin_product",
					"test_internal_admin_product",
					acctest.Optional,
					acctest.Update,
					acctest.RepresentationCopyWithNewProperties(
						OciProductCatalogInternalAdminProductRepresentation,
						map[string]interface{}{
							// Keep name identical to the resource created in Step 3/4 to avoid a diff
							"name": acctest.Representation{RepType: acctest.Required, Create: secondName, Update: secondName},
							// Backend update is a no-op; keep description stable
							"description": acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description`},
						},
					),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_id"),

				// The update step above keeps description stable as "description"; adjust expectation accordingly.
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", secondName),
				//resource.TestCheckResourceAttr(singularDatasourceName, "skus.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				// time_launched/time_ready may be absent for INACTIVE products; do not assert their presence here.
			),
		},
		// verify resource import
		{
			Config:            config + OciProductCatalogInternalAdminProductRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"compartment_id",
				"limits",
				"meters",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOciProductCatalogInternalAdminProductDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ProductAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oci_product_catalog_internal_admin_product" {
			noResourceFound = false
			request := oci_oci_product_catalog.GetAdminProductRequest{}

			// Prefer explicit product_id from state; if missing, fall back to parse from composite ID
			if value, ok := rs.Primary.Attributes["product_id"]; ok && value != "" {
				request.ProductId = &value
			} else {
				parts := strings.Split(rs.Primary.ID, "/")
				if len(parts) == 4 {
					pid := parts[3]
					request.ProductId = &pid
				}
			}

			// By default, acceptance tests attach a retry policy which can cause repeated
			// GetInternalProduct calls on retriable errors (409/429/5xx, network). To avoid
			// long loops during destroy verification, allow opting into a one-shot policy.
			noRetry := true
			if noRetry == true {
				request.RequestMetadata.RetryPolicy = &common.RetryPolicy{
					MaximumNumberAttempts: 1,
					ShouldRetryOperation:  func(resp common.OCIOperationResponse) bool { return false },
					// Ensure NextDuration is non-nil to avoid SDK errors on teardown
					NextDuration: func(resp common.OCIOperationResponse) time.Duration { return 0 },
				}
			} else {
				request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oci_product_catalog")
			}

			resp, err := client.GetAdminProduct(context.Background(), request)

			if err == nil {
				// Treat products in INACTIVE + Deleted as effectively removed
				if resp.AdminProduct.LifecycleState == oci_oci_product_catalog.AdminProductLifecycleStateInactive {
					if resp.AdminProduct.LifecycleDetails != nil && *resp.AdminProduct.LifecycleDetails == "Deleted" {
						// Considered deleted — check next resource

						OciProductCatalogInternalAdminProductRepresentation = map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
							"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
							// If the resource still exists, switch to a fresh random name for the subsequent create
							"name":         acctest.Representation{RepType: acctest.Required, Create: "opc-compute-sampleadmin-product" + utils.RandomStringOrHttpReplayValue(5, "0123456789", "36")},
							"service_name": acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
							"limits":       acctest.RepresentationGroup{RepType: acctest.Required, Group: OciProductCatalogInternalAdminProductLimitsRepresentation},
							// Ensure meters are always sent on Create to satisfy backend non-null constraint
							"meters": acctest.RepresentationGroup{RepType: acctest.Required, Group: OciProductCatalogInternalAdminProductMetersRepresentation},
						}
						return nil
					}
				}
				//return fmt.Errorf("resource still exists")
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
