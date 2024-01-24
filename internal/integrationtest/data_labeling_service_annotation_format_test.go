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
	DataLabelingServiceannotationFormatDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DataLabelingServiceAnnotationFormatResourceConfig = ""
)

// issue-routing-tag: data_labeling_service/default
func TestDataLabelingServiceAnnotationFormatResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataLabelingServiceAnnotationFormatResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_labeling_service_annotation_formats.test_annotation_formats"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_annotation_formats", "test_annotation_formats", acctest.Required, acctest.Create, DataLabelingServiceannotationFormatDataSourceRepresentation) +
				compartmentIdVariableStr + DataLabelingServiceAnnotationFormatResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "annotation_format_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "annotation_format_collection.0.items.#", "4"),
			),
		},
	})
}
