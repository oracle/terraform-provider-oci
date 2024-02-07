// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	MediaServicesMediaServicesMediaWorkflowJobFactSingularDataSourceRepresentation = map[string]interface{}{
		"key":                   acctest.Representation{RepType: acctest.Required, Create: `0`},
		"media_workflow_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
	}

	MediaServicesMediaServicesMediaWorkflowJobFactDataSourceRepresentation = map[string]interface{}{
		"media_workflow_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
		"key":                   acctest.Representation{RepType: acctest.Optional, Create: `0`},
		"type":                  acctest.Representation{RepType: acctest.Optional, Create: `runnableJob`},
	}

	MediaServicesMediaWorkflowJobFactResourceConfig = MediaServicesMediaWorkflowRequiredOnlyResource + acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Required, acctest.Create, MediaServicesMediaWorkflowJobRepresentation)
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaWorkflowJobFactResource_basic(t *testing.T) {
	//httpreplay.SetScenario("TestMediaServicesMediaWorkflowJobFactResource_basic")
	//defer httpreplay.SaveScenario()
	//
	//config := acctest.ProviderTestConfig()
	//
	//compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	//compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	//
	//datasourceName := "data.oci_media_services_media_workflow_job_facts.test_media_workflow_job_facts"
	//singularDatasourceName := "data.oci_media_services_media_workflow_job_fact.test_media_workflow_job_fact"
	//
	//acctest.SaveConfigContent("", "", "", t)
	//
	//acctest.ResourceTest(t, nil, []resource.TestStep{
	//	// verify datasource
	//	{
	//		Config: config +
	//			acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_job_facts", "test_media_workflow_job_facts", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowJobFactDataSourceRepresentation) +
	//			compartmentIdVariableStr + MediaServicesMediaWorkflowJobFactResourceConfig,
	//		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
	//			resource.TestCheckResourceAttrSet(datasourceName, "media_workflow_job_id"),
	//			resource.TestCheckResourceAttrSet(datasourceName, "media_workflow_job_fact_collection.#"),
	//		),
	//	},
	//	//verify singular datasource
	//	{
	//		Config: config +
	//			acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_job_fact", "test_media_workflow_job_fact", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowJobFactSingularDataSourceRepresentation) +
	//			compartmentIdVariableStr + MediaServicesMediaWorkflowJobFactResourceConfig,
	//		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
	//			resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
	//			resource.TestCheckResourceAttrSet(singularDatasourceName, "media_workflow_job_id"),
	//		),
	//	},
	//})
}
