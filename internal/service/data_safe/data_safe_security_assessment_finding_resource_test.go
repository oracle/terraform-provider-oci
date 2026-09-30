// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestDataSafeSecurityAssessmentFindingSetDataHandlesDetailsRepresentations(t *testing.T) {
	tests := []struct {
		name    string
		details interface{}
		want    []interface{}
	}{
		{
			name:    "string",
			details: "Password profile does not meet the recommendation.",
			want:    []interface{}{"Password profile does not meet the recommendation."},
		},
		{
			name: "object",
			details: map[string]interface{}{
				"recommendation": "Rotate the password",
				"riskCount":      float64(2),
			},
			want: []interface{}{`{"recommendation":"Rotate the password","riskCount":2}`},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resourceData := schema.TestResourceDataRaw(t, DataSafeSecurityAssessmentFindingResource().Schema, map[string]interface{}{})
			resourceData.SetId(GetSecurityAssessmentFindingCompositeId("test-security-assessment"))
			crud := &DataSafeSecurityAssessmentFindingResourceCrud{
				BaseCrud: tfresource.BaseCrud{D: resourceData},
				Res: &oci_data_safe.FindingSummary{
					Details: &test.details,
				},
			}

			if err := crud.SetData(); err != nil {
				t.Fatalf("SetData() returned an unexpected error: %v", err)
			}
			if got := resourceData.Get("details"); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("details = %#v, want %#v", got, test.want)
			}

			if err := crud.SetData(); err != nil {
				t.Fatalf("second SetData() returned an unexpected error: %v", err)
			}
			if got := resourceData.Get("details"); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("details after second refresh = %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestDataSafeSecurityAssessmentFindingSetDataRejectsUnsupportedDetails(t *testing.T) {
	details := interface{}([]interface{}{"unexpected"})
	resourceData := schema.TestResourceDataRaw(t, DataSafeSecurityAssessmentFindingResource().Schema, map[string]interface{}{})
	resourceData.SetId(GetSecurityAssessmentFindingCompositeId("test-security-assessment"))
	crud := &DataSafeSecurityAssessmentFindingResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
		Res: &oci_data_safe.FindingSummary{
			Details: &details,
		},
	}

	err := crud.SetData()
	if err == nil {
		t.Fatal("SetData() returned nil for an unsupported details type")
	}
	if !strings.Contains(err.Error(), "unsupported security assessment finding details type []interface {}") {
		t.Fatalf("SetData() error = %q, want unsupported type diagnostic", err)
	}
}

func TestDataSafeSecurityAssessmentFindingSetDataRejectsEmptyResponse(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, DataSafeSecurityAssessmentFindingResource().Schema, map[string]interface{}{})
	resourceData.SetId(GetSecurityAssessmentFindingCompositeId("test-security-assessment"))
	crud := &DataSafeSecurityAssessmentFindingResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
	}

	err := crud.SetData()
	if err == nil {
		t.Fatal("SetData() returned nil for an empty finding response")
	}
	if !strings.Contains(err.Error(), "security assessment finding response was empty") {
		t.Fatalf("SetData() error = %q, want empty response diagnostic", err)
	}
}

func TestDataSafeSecurityAssessmentFindingSetDataPreservesGeneratedFields(t *testing.T) {
	doclink := "https://docs.oracle.com/example"
	orp := "ORP-1"
	resourceData := schema.TestResourceDataRaw(t, DataSafeSecurityAssessmentFindingResource().Schema, map[string]interface{}{})
	resourceData.SetId(GetSecurityAssessmentFindingCompositeId("test-security-assessment"))
	crud := &DataSafeSecurityAssessmentFindingResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
		Res: &oci_data_safe.FindingSummary{
			Doclink: &doclink,
			References: &oci_data_safe.References{
				Orp: &orp,
			},
		},
	}

	if err := crud.SetData(); err != nil {
		t.Fatalf("SetData() returned an unexpected error: %v", err)
	}
	if got := resourceData.Get("doclink"); got != doclink {
		t.Fatalf("doclink = %#v, want %#v", got, doclink)
	}
	if got := resourceData.Get("references.0.orp"); got != orp {
		t.Fatalf("references.0.orp = %#v, want %#v", got, orp)
	}
}
