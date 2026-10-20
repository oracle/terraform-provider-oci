// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiApikeySingularDataSourceRepresentation = map[string]interface{}{
		"api_key_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apikey_id}`},
	}

	GenerativeAiApikeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `${var.apikey_display_name}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.apikey_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiApikeyDataSource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiApikeyDataSource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	apiKeyID := utils.GetEnvSettingWithBlankDefault("apikey_id")
	apiKeyDisplayName := utils.GetEnvSettingWithBlankDefault("apikey_display_name")

	variables := fmt.Sprintf(`
variable "compartment_id" { default = %q }
variable "apikey_id" { default = %q }
variable "apikey_display_name" { default = %q }
`, compartmentID, apiKeyID, apiKeyDisplayName)

	datasourceName := "data.oci_generative_ai_apikeys.test_apikeys"
	singularDatasourceName := "data.oci_generative_ai_apikey.test_apikey"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + variables +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_apikeys", "test_apikeys", acctest.Optional, acctest.Create, GenerativeAiApikeyDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentID),
				resource.TestCheckResourceAttr(datasourceName, "display_name", apiKeyDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "api_key_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "api_key_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "api_key_collection.0.items.0.id", apiKeyID),
				resource.TestCheckResourceAttr(datasourceName, "api_key_collection.0.items.0.display_name", apiKeyDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "api_key_collection.0.items.0.state", "ACTIVE"),
			),
		},
		{
			Config: config + variables +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_apikey", "test_apikey", acctest.Required, acctest.Create, GenerativeAiApikeySingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "api_key_id", apiKeyID),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", apiKeyID),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentID),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", apiKeyDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
			),
		},
	})
}
