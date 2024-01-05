// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package resourcediscovery

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func newTerraformStateWithValue(name, key, value string) *terraform.State {
	instanceState := terraform.NewInstanceStateShimmedFromValue(cty.Value{}, 0)
	instanceState.Attributes = make(map[string]string)
	instanceState.Attributes[key] = value
	state := terraform.NewState()
	state.RootModule().Resources = make(map[string]*terraform.ResourceState)
	state.RootModule().Resources[name] = &terraform.ResourceState{
		Primary: instanceState,
	}
	return state
}

func TestUnitgetResourceHint(t *testing.T) {
	type args struct {
		resourceClass string
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
	}
	tests := []testFormat{
		{
			name:     "Test hints are present",
			args:     args{resourceClass: "abc"},
			gotError: false,
		},
		{
			name:     "Test hints are absent",
			args:     args{resourceClass: "abcd"},
			gotError: true,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		ctx := &tf_export.ResourceDiscoveryContext{ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"abc": {}}}
		if res, err := ctx.GetResourceHint(test.args.resourceClass); (err != nil) != test.gotError {
			t.Logf("%s", fmt.Sprint(res))
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitdiscover(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type testFormat struct {
		name     string
		gotError bool
		mockFunc func()
	}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": exportParentDefinition},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": true},
		TenancyOcid:         tenancyOcid,
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			CompartmentId:   &compartmentId,
			Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			ExcludeServices: []string{},
		},
	}
	// Create a processing function that adds a new attribute to every discovered resource
	exportParentDefinition.ProcessDiscoveredResourcesFn = func(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
		for _, resource := range resources {
			resource.SourceAttributes["added_by_process_function"] = true
		}
		return resources, nil
	}
	defer func() { exportChildDefinition.ProcessDiscoveredResourcesFn = nil }()

	r := &resourceDiscoveryWithTargetIds{
		resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
			ctx:                 ctx,
			name:                "oci_test_parent",
			discoveredResources: []*tf_export.OCIResource{},
			omittedResources:    []*tf_export.OCIResource{},
		},
		exportIds: map[string]string{"s": "s"},
	}
	tests := []testFormat{
		{
			name:     "Test no error is returned",
			gotError: false,
			mockFunc: func() {
				tf_export.GenerateTerraformNameFromResource = func(resourceAttributes map[string]interface{}, resourceSchema map[string]*schema.Schema) (string, error) {
					return "", errors.New("")
				}
			},
		},
		{
			name:     "Test no error is returned",
			gotError: false,
			mockFunc: func() {
				tf_export.GenerateTerraformNameFromResource = func(resourceAttributes map[string]interface{}, resourceSchema map[string]*schema.Schema) (string, error) {
					return "test", nil
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if err := r.discover(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitwriteConfiguration(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type testFormat struct {
		name     string
		gotError bool
		mockFunc func()
	}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, _ := os.Getwd()
	outputDir = fmt.Sprintf("%s%swriteConfigurationTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": {ResourceClass: "oci_test_parent"}},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": true},
		TenancyOcid:         tenancyOcid,
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			CompartmentId:   &compartmentId,
			Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			ExcludeServices: []string{},
			OutputDir:       &outputDir,
		},
	}
	r := resourceDiscoveryBaseStep{
		ctx:  ctx,
		name: "oci_test_parent",
		discoveredResources: []*tf_export.OCIResource{
			{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: tf_export.TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &tf_export.TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
			},
			{
				// resource with import failure
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: tf_export.TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true,
			},
		},
		omittedResources: []*tf_export.OCIResource{},
	}
	tests := []testFormat{
		{
			name:     "Test no error is returned",
			gotError: false,
			mockFunc: func() {
				tf_export.GetHclStringFromGenericMap = func(builder *strings.Builder, ociRes *tf_export.OCIResource, interpolationMap map[string]string) error {
					return nil
				}
			},
		},
		{
			name:     "Test error is returned from getHclStringFromGenericMap()",
			gotError: true,
			mockFunc: func() {
				tf_export.GetHclStringFromGenericMap = func(builder *strings.Builder, ociRes *tf_export.OCIResource, interpolationMap map[string]string) error {
					return errors.New("")
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if err := r.writeConfiguration(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
	os.RemoveAll(outputDir)
}

func TestUnitwriteTmpState(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type testFormat struct {
		name     string
		gotError bool
		r        resourceDiscoveryBaseStep
	}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, _ := os.Getwd()
	outputDir = fmt.Sprintf("%s%swriteTmpStateTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	tf, _, _ := createTerraformStruct(&tf_export.ExportCommandArgs{
		OutputDir: &outputDir,
	})
	terraformInitMockVar = func(r *resourceDiscoveryBaseStep, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
		return nil
	}
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": {ResourceClass: "oci_test_parent"}},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": true},
		TenancyOcid:         tenancyOcid,
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			CompartmentId:   &compartmentId,
			Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			ExcludeServices: []string{},
			OutputDir:       &outputDir,
		},
		Terraform: tf,
	}
	tests := []testFormat{
		{
			name:     "Test write error is returned",
			gotError: true,
			r: resourceDiscoveryBaseStep{
				ctx:  ctx,
				name: "oci_test_parent",
				discoveredResources: []*tf_export.OCIResource{
					{
						CompartmentId: resourceDiscoveryTestCompartmentOcid,
						TerraformResource: tf_export.TerraformResource{
							Id:                "ocid1.a.b.c",
							TerraformClass:    "oci_resource_type1",
							TerraformName:     "type1_res1",
							TerraformTypeInfo: &tf_export.TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
						},
					},
				},
				omittedResources: []*tf_export.OCIResource{},
			},
		},
		{
			name:     "Test no error is returned",
			gotError: false,
			r: resourceDiscoveryBaseStep{
				ctx:                 ctx,
				name:                "oci_test_parent",
				discoveredResources: []*tf_export.OCIResource{},
				omittedResources:    []*tf_export.OCIResource{},
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		//test.mockFunc()
		if err := test.r.writeTmpState(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
	os.RemoveAll(outputDir)
}

func TestUnitwriteTmpConfigurationForImport(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type args struct {
		resourceClass string
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
	}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, _ := os.Getwd()
	outputDir = fmt.Sprintf("%s%swriteTmpConfigurationForImportTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	tf, _, _ := createTerraformStruct(&tf_export.ExportCommandArgs{
		OutputDir: &outputDir,
	})
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": {ResourceClass: "oci_test_parent"}},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": true},
		TenancyOcid:         tenancyOcid,
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			CompartmentId:   &compartmentId,
			Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			ExcludeServices: []string{},
			OutputDir:       &outputDir,
		},
		Terraform: tf,
	}
	r := resourceDiscoveryBaseStep{
		ctx:  ctx,
		name: "oci_test_parent",
		discoveredResources: []*tf_export.OCIResource{
			{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: tf_export.TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &tf_export.TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
			},
			{
				// resource with import failure
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: tf_export.TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true,
			},
		},
		omittedResources: []*tf_export.OCIResource{},
	}
	tests := []testFormat{
		{
			name:     "Test no error is returned",
			args:     args{resourceClass: "abc"},
			gotError: false,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		//test.mockFunc()
		if err := r.writeTmpConfigurationForImport(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
	os.RemoveAll(outputDir)
}

func TestUnitmergeTempStateFiles(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type args struct {
		resourceClass string
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
	}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, _ := os.Getwd()
	outputDir = fmt.Sprintf("%s%smergeTempStateFilesTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": {ResourceClass: "oci_test_parent"}},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": true},
		TenancyOcid:         tenancyOcid,
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			CompartmentId:   &compartmentId,
			Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			ExcludeServices: []string{},
			OutputDir:       &outputDir,
		},
	}
	r := resourceDiscoveryBaseStep{
		ctx:                 ctx,
		name:                "oci_test_parent",
		discoveredResources: []*tf_export.OCIResource{},
		omittedResources:    []*tf_export.OCIResource{},
	}
	tests := []testFormat{
		{
			name:     "Test no error is returned",
			args:     args{resourceClass: "abc"},
			gotError: false,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		//test.mockFunc()
		if err := r.mergeTempStateFiles(outputDir); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
	os.RemoveAll(outputDir)
}

func TestUnitmergeState(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type args struct {
		state1 interface{}
		state2 interface{}
	}
	type testFormat struct {
		name     string
		args     args
		gotError bool
	}
	s1 := newTerraformStateWithValue("name", "key", "value1")
	s2 := newTerraformStateWithValue("name", "key", "value2")
	tests := []testFormat{
		{
			name:     "Test passing valid states",
			args:     args{state1: s1, state2: s2},
			gotError: false,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if _, err := mergeState(test.args.state1, test.args.state2); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitmergeGeneratedStateFile(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	mergeState = func(state1 interface{}, state2 interface{}) (interface{}, error) {
		return newTerraformStateWithValue("name", "key", "value1"), nil
	}
	type testFormat struct {
		name     string
		gotError bool
		ctx      *tf_export.ResourceDiscoveryContext
		r        resourceDiscoveryBaseStep
	}
	tests := []testFormat{
		{
			name:     "Test ctx state is not nil",
			gotError: false,
			r: resourceDiscoveryBaseStep{
				ctx: &tf_export.ResourceDiscoveryContext{
					State: newTerraformStateWithValue("name", "key", "value1"),
				},
				tempState: newTerraformStateWithValue("name", "key", "value1"),
			},
		},
		{
			name:     "Test ctx state is nil",
			gotError: false,
			r: resourceDiscoveryBaseStep{
				ctx:       &tf_export.ResourceDiscoveryContext{},
				tempState: newTerraformStateWithValue("name", "key", "value1"),
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if err := test.r.mergeGeneratedStateFile(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitpostValidate(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	type testFormat struct {
		name        string
		errorLength int
	}
	ctx := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": exportParentDefinition},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": false},
	}
	tests := []testFormat{
		{
			name:        "Test one missing resource",
			errorLength: 1,
		},
		{
			name:        "Missing resource id check",
			errorLength: 2,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		ctx.PostValidate()
		// Validate encapsulated error message: one or more expected resource ids were not found
		if len(ctx.ErrorList.Errors) == 2 {
			for _, resourceDiscoveryError := range ctx.ErrorList.Errors {
				// element is the element from someSlice for where we are
				fmt.Println("Error", resourceDiscoveryError.Error)
			}
		}
		if len(ctx.ErrorList.Errors) != test.errorLength {

			t.Errorf("Output error length - %d which is not equal to expected error - %d", len(ctx.ErrorList.Errors), test.errorLength)
		}
	}
}
