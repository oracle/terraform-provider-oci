// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentDataIngestionJobLogContentSingularDataSourceRepresentation = map[string]interface{}{
		"data_ingestion_job_id": acctest.Representation{RepType: acctest.Required, Create: `${var.data_ingestion_job_id_env}`},
	}

	GenerativeAiAgentDataIngestionJobLogContentResourceConfig = ``
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentDataIngestionJobLogContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentDataIngestionJobLogContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dataIngestionJobId := utils.GetEnvSettingWithBlankDefault("data_ingestion_job_id")
	dataIngestionJobIdVariableStr := fmt.Sprintf("variable \"data_ingestion_job_id_env\" { default = \"%s\" }\n", dataIngestionJobId)

	singularDatasourceName := "data.oci_generative_ai_agent_data_ingestion_job_log_content.test_data_ingestion_job_log_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job_log_content", "test_data_ingestion_job_log_content", acctest.Required, acctest.Create, GenerativeAiAgentDataIngestionJobLogContentSingularDataSourceRepresentation) +
				dataIngestionJobIdVariableStr + compartmentIdVariableStr + GenerativeAiAgentDataIngestionJobLogContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_ingestion_job_id"),
			),
		},
	})
}
