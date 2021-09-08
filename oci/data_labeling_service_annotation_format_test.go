// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	annotationFormatDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	AnnotationFormatResourceConfig = ""
)

// issue-routing-tag: data_labeling_service/default
func TestDataLabelingServiceAnnotationFormatResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataLabelingServiceAnnotationFormatResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_labeling_service_annotation_formats.test_annotation_formats"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_annotation_formats", "test_annotation_formats", Required, Create, annotationFormatDataSourceRepresentation) +
				compartmentIdVariableStr + AnnotationFormatResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "annotation_format_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "annotation_format_collection.0.items.#", "4"),
			),
		},
	})
}
