// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	FusionAppsFusionAppsFusionEnvironmentServiceAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"service_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
	}

	FusionAppsFusionAppsFusionEnvironmentServiceAttachmentDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"service_instance_type": acctest.Representation{RepType: acctest.Optional, Create: `DIGITAL_ASSISTANT`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	FusionAppsFusionEnvironmentServiceAttachmentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentServiceAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentServiceAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fusion_apps_fusion_environment_service_attachments.test_fusion_environment_service_attachments"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_service_attachment.test_fusion_environment_service_attachment"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachments", "test_fusion_environment_service_attachments", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentServiceAttachmentDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentServiceAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "service_instance_type", "DIGITAL_ASSISTANT"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "service_attachment_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_service_attachment", "test_fusion_environment_service_attachment", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentServiceAttachmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentServiceAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_attachment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_sku_based"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_instance_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
