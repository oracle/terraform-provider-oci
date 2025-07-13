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
	DatascienceModelGroupArtifactContentSingularDataSourceRepresentation = map[string]interface{}{
		"model_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_group.test_model_group.id}`},
		"range":          acctest.Representation{RepType: acctest.Optional, Create: `range`},
	}

	DatascienceModelGroupArtifactContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group", acctest.Required, acctest.Create, DatascienceModelGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceModelGroupArtifactContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelGroupArtifactContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_group_artifact_content", "test_model_group_artifact_content", acctest.Required, acctest.Create, DatascienceModelGroupArtifactContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelGroupArtifactContentResourceConfig,
		},
	})
}
