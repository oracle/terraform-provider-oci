// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

const (
	IdentityProviderRequiredOnlyResource = IdentityProviderResourceDependenciesRequiredOnly + `
resource "oci_identity_identity_provider" "test_identity_provider" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.identity_provider_description}"
	metadata = "${var.identity_provider_metadata != "" ? var.identity_provider_metadata : "${file("${var.identity_provider_metadata_file}")}"}"
	metadata_url = "${var.identity_provider_metadata_url}"
	name = "${var.identity_provider_name}"
	product_type = "${var.identity_provider_product_type}"
	protocol = "${var.identity_provider_protocol}"
}
`

	IdentityProviderResourceConfig = IdentityProviderResourceDependencies + `
resource "oci_identity_identity_provider" "test_identity_provider" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.identity_provider_description}"
	metadata = "${var.identity_provider_metadata != "" ? var.identity_provider_metadata : "${file("${var.identity_provider_metadata_file}")}"}"
	metadata_url = "${var.identity_provider_metadata_url}"
	name = "${var.identity_provider_name}"
	product_type = "${var.identity_provider_product_type}"
	protocol = "${var.identity_provider_protocol}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.identity_provider_defined_tags_value}")}"
	freeform_tags = "${var.identity_provider_freeform_tags}"
}
`
	IdentityProviderPropertyVariables = `
variable "identity_provider_defined_tags_value" { default = "value" }
variable "identity_provider_description" { default = "description" }
variable "identity_provider_freeform_tags" { default = {"Department"= "Finance"} }
variable "identity_provider_metadata" { default = "" }
variable "identity_provider_metadata_file" { default = "sampleFederationMetadata.xml" }
variable "identity_provider_metadata_url" { default = "metadataUrl" }
variable "identity_provider_name" { default = "test-idp-saml2-adfs" }
variable "identity_provider_product_type" { default = "ADFS" }
variable "identity_provider_protocol" { default = "SAML2" }

`
	IdentityProviderResourceDependenciesRequiredOnly = `
`
	IdentityProviderResourceDependencies = DefinedTagsDependencies
)

func TestIdentityIdentityProviderResource_basic(t *testing.T) {
	metadataFile := getEnvSetting("identity_provider_metadata_file", "")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	resourceName := "oci_identity_identity_provider.test_identity_provider"
	datasourceName := "data.oci_identity_identity_providers.test_identity_providers"

	var resId, resId2 string
	//var resId string

	metadataContents, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		log.Panic("Unable to read the file ", metadataFile)
	}
	metadata := string(metadataContents)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityIdentityProviderDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + IdentityProviderPropertyVariables + compartmentIdVariableStr + IdentityProviderRequiredOnlyResource,
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
				Config: config + IdentityProviderPropertyVariables + compartmentIdVariableStr + IdentityProviderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
					resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl"),
					resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
					resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "identity_provider_defined_tags_value" { default = "updatedValue" }
variable "identity_provider_description" { default = "description2" }
variable "identity_provider_freeform_tags" { default = {"Department"= "Accounting"} }
variable "identity_provider_metadata" { default = "" }
variable "identity_provider_metadata_file" { default = "sampleFederationMetadata.xml" }
variable "identity_provider_metadata_url" { default = "metadataUrl2" }
variable "identity_provider_name" { default = "test-idp-saml2-adfs" }
variable "identity_provider_product_type" { default = "ADFS" }
variable "identity_provider_protocol" { default = "SAML2" }

                ` + compartmentIdVariableStr + IdentityProviderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "metadata", metadata),
					resource.TestCheckResourceAttr(resourceName, "metadata_url", "metadataUrl2"),
					resource.TestCheckResourceAttr(resourceName, "name", "test-idp-saml2-adfs"),
					resource.TestCheckResourceAttr(resourceName, "product_type", "ADFS"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "SAML2"),
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
				Config: config + `
variable "identity_provider_defined_tags_value" { default = "updatedValue" }
variable "identity_provider_description" { default = "description2" }
variable "identity_provider_freeform_tags" { default = {"Department"= "Accounting"} }
variable "identity_provider_metadata" { default = "" }
variable "identity_provider_metadata_file" { default = "sampleFederationMetadata.xml" }
variable "identity_provider_metadata_url" { default = "metadataUrl2" }
variable "identity_provider_name" { default = "test-idp-saml2-adfs" }
variable "identity_provider_product_type" { default = "ADFS" }
variable "identity_provider_protocol" { default = "SAML2" }

data "oci_identity_identity_providers" "test_identity_providers" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	protocol = "${var.identity_provider_protocol}"

    filter {
    	name = "id"
    	values = ["${oci_identity_identity_provider.test_identity_provider.id}"]
    }
}
                ` + compartmentIdVariableStr + IdentityProviderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "protocol", "SAML2"),

					resource.TestCheckResourceAttr(datasourceName, "identity_providers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.name", "test-idp-saml2-adfs"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.product_type", "ADFS"),
					resource.TestCheckResourceAttr(datasourceName, "identity_providers.0.protocol", "SAML2"),
					resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "identity_providers.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}

func testAccCheckIdentityIdentityProviderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_identity_provider" {
			noResourceFound = false
			request := oci_identity.GetIdentityProviderRequest{}

			tmp := rs.Primary.ID
			request.IdentityProviderId = &tmp

			_, err := client.GetIdentityProvider(context.Background(), request)

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
