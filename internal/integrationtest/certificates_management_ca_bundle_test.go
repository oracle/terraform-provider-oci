// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	caBundleRequiredResourceName = "test-ca-bundle-required-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	caBundleOptionalResourceName = "test-ca-bundle-optional-" + utils.RandomString(10, utils.CharsetWithoutDigits)

	CaBundleRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Required, acctest.Create, caBundleRepresentationRequired)

	CaBundleResourceConfig = CaBundleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Update, caBundleRepresentation)

	caBundleSingularDataSourceRepresentation = map[string]interface{}{
		"ca_bundle_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_ca_bundle.test_ca_bundle.id}`},
	}

	caBundleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: caBundleOptionalResourceName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: caBundleDataSourceFilterRepresentation}}
	caBundleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_certificates_management_ca_bundle.test_ca_bundle.id}`}},
	}

	caBundleRepresentation = map[string]interface{}{
		"ca_bundle_pem":  acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN CERTIFICATE-----\nMIICwjCCAaqgAwIBAgICEAAwDQYJKoZIhvcNAQELBQAwFzEVMBMGA1UEAwwMdGVz\ndC1yb290LWNhMB4XDTIwMDkxMTIyMjczMFoXDTIxMDkxMTIyMjczMFowHzEdMBsG\nA1UEAwwUdGVzdC1pbnRlcm1lZGlhdGUtY2EwggEiMA0GCSqGSIb3DQEBAQUAA4IB\nDwAwggEKAoIBAQDoGijYmLa6P+2C2hcXV7DNBuZ0K+1x41lC8emin9d6wsxsWYJn\nUzthLgJ28eol9CsDQlgRt+rGfYzKgfG7y7Wg7WYQXFHYZt1ANRB0OEVJYm7xurMA\nw6T68gBBctbLCcsk2WmWKp50Js7bwMKGGpCuDQ3YGWb1Bn1cu1qC9eWxl+VXV7RU\nMug1CtloLLrUlZt3FjbKBZEhhq4L6ukKTfNUzAofu4Q0ACxWht7ueuJMWaAaOUR2\nHu5KuGjkiJLvW/wRFeN55+rtuh9Bw/fKnhEq5/q9VCmw4g9sfNO0wIUqBhYx8etF\nYsGww1FWOEtGdUDv8NazYowPdpDwZzWA9UaPAgMBAAGjEDAOMAwGA1UdEwQFMAMB\nAf8wDQYJKoZIhvcNAQELBQADggEBAIrr04h+LmTM/JXEE/7kD8sB77YQhkbaWZPV\nuOwgdwVldAaGJkp4nyFPgcm+W7AXuoMMCq1ONNQuic2voopGU3EFng/tGODdYxiR\nedHyRPX23tJF19MSIJebJ7PNu2LtUU7tQspf4yqt4sBxW0nwjjSNyxPIZdnjNgFG\nHqpGgfIS9rSm89A6XRhwLSiudsQV/bfdHOgYUNuaPZn5LiyxHvHAorsZ1rpymNa6\nL78Mn4YLEBIT6XftA67z9g9N963FhGM0Il2iUVXm4HU/XqM26s94ZHlCIVJnV5xQ\nmjUvxEKNZMc7CRgz1Mr0Am898gQQZZrySHpDNvhpP50OeYH9/mU=\n-----END CERTIFICATE-----`, Update: `-----BEGIN CERTIFICATE-----\n\nMIICwTCCAamgAwIBAgIJAMz8Puud7CUcMA0GCSqGSIb3DQEBCwUAMBcxFTATBgNV\nBAMMDHRlc3Qtcm9vdC1jYTAeFw0yMDA5MTEyMjI3MzBaFw0zMDA5MDkyMjI3MzBa\nMBcxFTATBgNVBAMMDHRlc3Qtcm9vdC1jYTCCASIwDQYJKoZIhvcNAQEBBQADggEP\nADCCAQoCggEBALQql3DFigzMMlTXGJlz4ulZJx3kPmCySwl2rzZ+jZhMF8Oy/5xn\n99ToHgbj+Pp/7XWuogtwkfoZYmU1BTHyjZlaNccTLwkP/4SkjZ7cKsv4TTCN0q/4\nZF6dVoQOom/owREX4YEnISUgN4U5vpZkzkr3NOpiULAmimrROq52TcC2y52ijui9\nSf7QDmtVW7M4kUFjT+If/2yxX/g1dsbKxh+cuZimLWypD3TmhWehglgEZRiDnXJL\n7davAG0DOeO2scnglOY0JAOyM6po4WALmBf6OBxl2jiYpwWgkXMhgGf5u3ID6IDP\nJ2ZeYb9PZ2Z5jYGxUDLRVaaXr0a4a4UbvuECAwEAAaMQMA4wDAYDVR0TBAUwAwEB\n/zANBgkqhkiG9w0BAQsFAAOCAQEATqNEMvk7NUi6M9tv743CVGZa9JLdpdD9ZCbY\nvH6qEl5nVVAiUiKCRSnntDNtUkFd9tVQGk5Hgp2aMm1epq6kC3N0K4Dbsg1dWU4O\nYRG47CSv+sBI4VpDqRjtOS3+FlmMlm237ahs7YcwHNUqqGKbl9QZACxXZeCArYhZ\nt6xDqLb0gKON5gKlhnUqhYTT4+CcBan2nnr7sbdZhwtOmFgrowvs4OVh6WwdbI9m\nSdNKzy7AXAlcodkXAfjIoODeOQdEK6/Ine1K+W4sOk9e+oymoD0KnzhDLI3MDB3V\nGKIiVcct3DMHGI+y04/kcuasqUo4l6lrjt2wJF1pHbaxKIcyGQ==\n-----END CERTIFICATE-----`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: caBundleOptionalResourceName},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}
	caBundleRepresentationRequired = map[string]interface{}{
		"ca_bundle_pem":  acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN CERTIFICATE-----\nMIICwjCCAaqgAwIBAgICEAAwDQYJKoZIhvcNAQELBQAwFzEVMBMGA1UEAwwMdGVz\ndC1yb290LWNhMB4XDTIwMDkxMTIyMjczMFoXDTIxMDkxMTIyMjczMFowHzEdMBsG\nA1UEAwwUdGVzdC1pbnRlcm1lZGlhdGUtY2EwggEiMA0GCSqGSIb3DQEBAQUAA4IB\nDwAwggEKAoIBAQDoGijYmLa6P+2C2hcXV7DNBuZ0K+1x41lC8emin9d6wsxsWYJn\nUzthLgJ28eol9CsDQlgRt+rGfYzKgfG7y7Wg7WYQXFHYZt1ANRB0OEVJYm7xurMA\nw6T68gBBctbLCcsk2WmWKp50Js7bwMKGGpCuDQ3YGWb1Bn1cu1qC9eWxl+VXV7RU\nMug1CtloLLrUlZt3FjbKBZEhhq4L6ukKTfNUzAofu4Q0ACxWht7ueuJMWaAaOUR2\nHu5KuGjkiJLvW/wRFeN55+rtuh9Bw/fKnhEq5/q9VCmw4g9sfNO0wIUqBhYx8etF\nYsGww1FWOEtGdUDv8NazYowPdpDwZzWA9UaPAgMBAAGjEDAOMAwGA1UdEwQFMAMB\nAf8wDQYJKoZIhvcNAQELBQADggEBAIrr04h+LmTM/JXEE/7kD8sB77YQhkbaWZPV\nuOwgdwVldAaGJkp4nyFPgcm+W7AXuoMMCq1ONNQuic2voopGU3EFng/tGODdYxiR\nedHyRPX23tJF19MSIJebJ7PNu2LtUU7tQspf4yqt4sBxW0nwjjSNyxPIZdnjNgFG\nHqpGgfIS9rSm89A6XRhwLSiudsQV/bfdHOgYUNuaPZn5LiyxHvHAorsZ1rpymNa6\nL78Mn4YLEBIT6XftA67z9g9N963FhGM0Il2iUVXm4HU/XqM26s94ZHlCIVJnV5xQ\nmjUvxEKNZMc7CRgz1Mr0Am898gQQZZrySHpDNvhpP50OeYH9/mU=\n-----END CERTIFICATE-----`, Update: `ca-bundle-2`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: caBundleRequiredResourceName},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	CaBundleResourceDependencies = DefinedTagsDependencies
)

func TestCertificatesManagementCaBundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCaBundleResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_certificates_management_ca_bundle.test_ca_bundle"
	datasourceName := "data.oci_certificates_management_ca_bundles.test_ca_bundles"
	singularDatasourceName := "data.oci_certificates_management_ca_bundle.test_ca_bundle"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CaBundleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Create, caBundleRepresentation), "certificatesmanagement", "caBundle", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCertificatesManagementCaBundleDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Required, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(caBundleRepresentation, map[string]interface{}{
							"name": acctest.Representation{RepType: acctest.Required, Create: caBundleRequiredResourceName},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ca_bundle_pem"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", caBundleRequiredResourceName),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CaBundleResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CaBundleResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Create, caBundleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ca_bundle_pem"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", caBundleOptionalResourceName),
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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CaBundleResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(caBundleRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ca_bundle_pem"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", caBundleOptionalResourceName),
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
				Config: config + compartmentIdVariableStr + CaBundleResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Update, caBundleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ca_bundle_pem"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", caBundleOptionalResourceName),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_ca_bundles", "test_ca_bundles", acctest.Optional, acctest.Create, caBundleDataSourceRepresentation) +
					compartmentIdVariableStr + CaBundleResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Optional, acctest.Update, caBundleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", caBundleOptionalResourceName),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "ca_bundle_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ca_bundle_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Required, acctest.Create, caBundleSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CaBundleResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ca_bundle_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", caBundleOptionalResourceName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CaBundleResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ca_bundle_pem",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCertificatesManagementCaBundleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CertificatesManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_certificates_management_ca_bundle" {
			noResourceFound = false
			request := oci_certificates_management.GetCaBundleRequest{}

			tmp := rs.Primary.ID
			request.CaBundleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "certificates_management")

			response, err := client.GetCaBundle(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_certificates_management.CaBundleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CertificatesManagementCaBundle") {
		resource.AddTestSweepers("CertificatesManagementCaBundle", &resource.Sweeper{
			Name:         "CertificatesManagementCaBundle",
			Dependencies: acctest.DependencyGraph["caBundle"],
			F:            sweepCertificatesManagementCaBundleResource,
		})
	}
}

func sweepCertificatesManagementCaBundleResource(compartment string) error {
	certificatesManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CertificatesManagementClient()
	caBundleIds, err := getCaBundleIds(compartment)
	if err != nil {
		return err
	}
	for _, caBundleId := range caBundleIds {
		if ok := acctest.SweeperDefaultResourceId[caBundleId]; !ok {
			deleteCaBundleRequest := oci_certificates_management.DeleteCaBundleRequest{}

			deleteCaBundleRequest.CaBundleId = &caBundleId

			deleteCaBundleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "certificates_management")
			_, error := certificatesManagementClient.DeleteCaBundle(context.Background(), deleteCaBundleRequest)
			if error != nil {
				fmt.Printf("Error deleting CaBundle %s %s, It is possible that the resource is already deleted. Please verify manually \n", caBundleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &caBundleId, caBundleSweepWaitCondition, time.Duration(3*time.Minute),
				caBundleSweepResponseFetchOperation, "certificates_management", true)
		}
	}
	return nil
}

func getCaBundleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CaBundleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	certificatesManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CertificatesManagementClient()

	listCaBundlesRequest := oci_certificates_management.ListCaBundlesRequest{}
	listCaBundlesRequest.CompartmentId = &compartmentId
	listCaBundlesRequest.LifecycleState = oci_certificates_management.ListCaBundlesLifecycleStateActive
	listCaBundlesResponse, err := certificatesManagementClient.ListCaBundles(context.Background(), listCaBundlesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CaBundle list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, caBundle := range listCaBundlesResponse.Items {
		id := *caBundle.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CaBundleId", id)
	}
	return resourceIds, nil
}

func caBundleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if caBundleResponse, ok := response.Response.(oci_certificates_management.GetCaBundleResponse); ok {
		return caBundleResponse.LifecycleState != oci_certificates_management.CaBundleLifecycleStateDeleted
	}
	return false
}

func caBundleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CertificatesManagementClient().GetCaBundle(context.Background(), oci_certificates_management.GetCaBundleRequest{
		CaBundleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
