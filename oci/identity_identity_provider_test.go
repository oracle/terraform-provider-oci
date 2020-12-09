// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v30/common"
	oci_identity "github.com/oracle/oci-go-sdk/v30/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const (
	IdentityProviderPropertyVariables = `
variable "identity_provider_metadata" { default = "" }
variable "identity_provider_metadata_file" { default = "{{.metadata_file}}" }
`
)

var (
	IdentityProviderRequiredOnlyResource = IdentityProviderResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Required, Create, identityProviderRepresentation)

	identityProviderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"protocol":       Representation{repType: Required, create: `SAML2`},
		"filter":         RepresentationGroup{Required, identityProviderDataSourceFilterRepresentation}}
	identityProviderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_identity_provider.test_identity_provider.id}`}},
	}

	identityProviderRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":         Representation{repType: Required, create: `description`, update: `description2`},
		"metadata":            Representation{repType: Required, create: `${file("${var.identity_provider_metadata_file}")}`},
		"metadata_url":        Representation{repType: Required, create: `metadataUrl`, update: `metadataUrl2`},
		"name":                Representation{repType: Required, create: `test-idp-saml2-adfs`},
		"product_type":        Representation{repType: Required, create: `ADFS`},
		"protocol":            Representation{repType: Required, create: `SAML2`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_attributes": Representation{repType: Optional, create: map[string]string{"clientId": "app_sf3kdjf3"}},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	IdentityProviderResourceDependencies = IdentityProviderPropertyVariables +
		DefinedTagsDependencies
)

func TestIdentityIdentityProviderResource_basic(t *testing.T) {
	metadataFile := getEnvSettingWithBlankDefault("identity_provider_metadata_file")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	httpreplay.SetScenario("TestIdentityIdentityProviderResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_identity_provider.test_identity_provider"
	datasourceName := "data.oci_identity_identity_providers.test_identity_providers"

	var resId, resId2 string

	metadataContents, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		log.Panic("Unable to read the file ", metadataFile)
	}
	metadata := string(metadataContents)

	_, tokenFn := tokenizeWithHttpReplay("identity_provider")
	IdentityProviderResourceDependencies = tokenFn(IdentityProviderResourceDependencies, map[string]string{"metadata_file": metadataFile})

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityIdentityProviderDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + IdentityProviderResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Required, Create, identityProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
					resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl"),
					resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
					resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + IdentityProviderResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + IdentityProviderResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Optional, Create, identityProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + IdentityProviderResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Optional, Update, identityProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_identity_identity_providers", "test_identity_providers", Optional, Update, identityProviderDataSourceRepresentation) +
					compartmentIdVariableStr + IdentityProviderResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Optional, Update, identityProviderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "protocol", "SAML2"),

					resource.TestCheckResourceAttr(datasourceName, "identity_providers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.defined_tags.%", "1"),
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckIdentityIdentityProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_identity_provider" {
			noResourceFound = false
			request := oci_identity.GetIdentityProviderRequest{}

			tmp := rs.Primary.ID
			request.IdentityProviderId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

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
