// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	IdentityProviderPropertyVariables = `
variable "identity_provider_metadata" { default = "" }
variable "identity_provider_metadata_file" { default = "{{.metadata_file}}" }
`
)

var (
	IdentityIdentityProviderRequiredOnlyResource = IdentityIdentityProviderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Required, acctest.Create, IdentityIdentityProviderRepresentation)

	IdentityIdentityIdentityProviderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"protocol":       acctest.Representation{RepType: acctest.Required, Create: `SAML2`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test-idp-saml2-adfs`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityIdentityProviderDataSourceFilterRepresentation}}
	IdentityIdentityProviderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_identity_identity_provider.test_identity_provider.id}`}},
	}

	IdentityIdentityProviderRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":         acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"metadata":            acctest.Representation{RepType: acctest.Required, Create: `${file("${var.identity_provider_metadata_file}")}`},
		"metadata_url":        acctest.Representation{RepType: acctest.Required, Create: `metadataUrl`, Update: `metadataUrl2`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `test-idp-saml2-adfs`},
		"product_type":        acctest.Representation{RepType: acctest.Required, Create: `ADFS`},
		"protocol":            acctest.Representation{RepType: acctest.Required, Create: `SAML2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_attributes": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"clientId": "app_sf3kdjf3"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	IdentityIdentityProviderResourceDependencies = IdentityProviderPropertyVariables +
		DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityIdentityProviderResource_basic(t *testing.T) {
	metadataFile := utils.GetEnvSettingWithBlankDefault("identity_provider_metadata_file")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	httpreplay.SetScenario("TestIdentityIdentityProviderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_identity_provider.test_identity_provider"
	datasourceName := "data.oci_identity_identity_providers.test_identity_providers"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityIdentityProviderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Optional, acctest.Create, IdentityIdentityProviderRepresentation), "identity", "identityProvider", t)

	metadataContents, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		log.Panic("Unable to read the file ", metadataFile)
	}
	metadata := string(metadataContents)

	_, tokenFn := acctest.TokenizeWithHttpReplay("identity_provider")
	IdentityIdentityProviderResourceDependencies = tokenFn(IdentityIdentityProviderResourceDependencies, map[string]string{"metadata_file": metadataFile})

	acctest.ResourceTest(t, testAccCheckIdentityIdentityProviderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Required, acctest.Create, IdentityIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
				resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityIdentityProviderResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Optional, acctest.Create, IdentityIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
				resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),
				resource.TestCheckResourceAttrSet(resourceName, "redirect_url"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Optional, acctest.Update, IdentityIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_attributes.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
				resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl2"),
				resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),
				resource.TestCheckResourceAttrSet(resourceName, "redirect_url"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_identity_providers", "test_identity_providers", acctest.Optional, acctest.Update, IdentityIdentityIdentityProviderDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityIdentityProviderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", acctest.Optional, acctest.Update, IdentityIdentityProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(datasourceName, "protocol", "SAML2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "identity_providers.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.freeform_attributes.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.metadata"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.metadata_url", "metadataUrl2"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.product_type", "ADFS"),
				resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.protocol", "SAML2"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.redirect_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + IdentityIdentityProviderRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityIdentityProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_identity_provider" {
			noResourceFound = false
			request := oci_identity.GetIdentityProviderRequest{}

			tmp := rs.Primary.ID
			request.IdentityProviderId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity")

			response, err := client.GetIdentityProvider(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.IdentityProviderLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
