package resourcediscovery

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitExtractVarsExportResourceLevel(t *testing.T) {
	type args struct {
		exportVars []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "Test empty input",
			args: args{
				exportVars: []string{""},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Test invalid input with no .",
			args: args{
				exportVars: []string{"oci_core_subnet"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Test invalid input for nested attribute with 3 . ",
			args: args{
				exportVars: []string{"oci_core_subnet.a.b"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Test invalid input with resource not supported in RD",
			args: args{
				exportVars: []string{"abc.abc"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "Test positive input ",
			args: args{
				exportVars: []string{"oci_core_subnet.availability_domain"},
			},
			wantErr: false,
			want:    map[string][]string{"oci_core_subnet": []string{"availability_domain"}},
		},
		{
			name: "Test positive 2 attributes of 1 resource input ",
			args: args{
				exportVars: []string{"oci_core_subnet.availability_domain", "oci_core_subnet.display_name"},
			},
			wantErr: false,
			want:    map[string][]string{"oci_core_subnet": []string{"availability_domain", "display_name"}},
		},
		{
			name: "Test positive 2 attributes of 2 resource input ",
			args: args{
				exportVars: []string{"oci_core_subnet.availability_domain", "oci_core_subnet.display_name", "oci_core_vcn.availability_domain", "oci_core_vcn.display_name"},
			},
			wantErr: false,
			want:    map[string][]string{"oci_core_subnet": []string{"availability_domain", "display_name"}, "oci_core_vcn": []string{"availability_domain", "display_name"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVarsExportResourceLevel(tt.args.exportVars)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVarsExportResourceLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractVarsExportResourceLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitExportAttributeForResourceLevel(t *testing.T) {
	tfHclVersion = &TfHclVersion12{Value: TfVersion12}
	vars = map[string]string{"region": "phx"}
	varsExportForResourceLevel = map[string][]string{}
	referenceMap = map[string]string{}
	sourceAttributes := map[string]interface{}{"available_domain": "ad1"}
	resourceType := "oci_core_vcn"
	resourceName := "Test_vcn"
	varsExport := map[string][]string{"oci_core_vcn": []string{"available_domain"}}
	interpolationMap := map[string]string{"region": "phx"}

	err := exportAttributeForResourceLevel(sourceAttributes, resourceType, resourceName, varsExport, interpolationMap)
	assert.NoError(t, err)
	// vars file should contain this map after exportAttributeForResourceLevel()
	v, exist := vars["oci_core_vcn--available_domain--Test_vcn"]
	assert.True(t, exist)
	assert.Contains(t, v, "ad1")

	// interpolationMap should contain this map after exportAttributeForResourceLevel()
	v, exist = interpolationMap["ad1"]
	assert.True(t, exist)
	assert.Contains(t, v, "oci_core_vcn--available_domain--Test_vcn")
}

func TestUnitGetVarNameFromAttributeOfResources(t *testing.T) {
	type args struct {
		attribute    string
		resourceType string
		resourceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty input",
			args: args{},
			want: "----",
		},
		{
			name: "valid input",
			args: args{
				attribute:    "available_domain",
				resourceType: "oci_core_vcn",
				resourceName: "test_VCN",
			},
			want: "oci_core_vcn--available_domain--test_VCN",
		},
		{
			name: "nested attribute input",
			args: args{
				attribute:    "config.available_domain",
				resourceType: "oci_core_vcn",
				resourceName: "test_VCN",
			},
			want: "oci_core_vcn--config-available_domain--test_VCN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVarNameFromAttributeOfResources(tt.args.attribute, tt.args.resourceType, tt.args.resourceName); got != tt.want {
				t.Errorf("getVarNameFromAttributeOfResources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGetVarNameFromAttributeAndValue(t *testing.T) {
	type args struct {
		attribute string
		value     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive case input with special char: availability_domain, NyKp:PHX-AD-1",
			args: args{
				attribute: "availability_domain",
				value:     "NyKp:PHX-AD-1",
			},
			want: "availability_domain--NyKp-PHX-AD-1",
		},
		{
			name: "Positive case input with numeric and word",
			args: args{
				attribute: "a",
				value:     "b123",
			},
			want: "a--b123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVarNameFromAttributeAndValue(tt.args.attribute, tt.args.value); got != tt.want {
				t.Errorf("getVarNameFromAttributeAndValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitExtractVarsExportGlobalLevel(t *testing.T) {
	type args struct {
		attributes []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "positive test case",
			args: args{
				attributes: []string{"a", "b"},
			},
			want:    []string{"a", "b"},
			wantErr: false,
		},
		{
			name: "negative test case",
			args: args{
				attributes: []string{"a.b", "b"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVarsExportGlobalLevel(tt.args.attributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVarsExportGlobalLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractVarsExportGlobalLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitExportAttributeForGlobalLevel(t *testing.T) {
	tfHclVersion = &TfHclVersion12{Value: TfVersion12}
	vars = map[string]string{"region": "phx"}
	referenceMap = map[string]string{}
	sourceAttributes := map[string]interface{}{"available_domain": "ad1"}
	varsExport := []string{"available_domain"}
	interpolationMap := map[string]string{"region": "phx"}
	resourceName := "Test_vcn"

	err := exportAttributeForGlobalLevel(sourceAttributes, resourceName, varsExport, interpolationMap)
	assert.NoError(t, err)
	// vars file should contain this map after exportAttributeForResourceLevel()
	v, exist := vars["available_domain--ad1"]
	assert.True(t, exist)
	assert.Contains(t, v, "ad1")

	// interpolationMap should contain this map after exportAttributeForResourceLevel()
	v, exist = interpolationMap["ad1"]
	assert.True(t, exist)
	assert.Contains(t, v, "available_domain--ad1")
}
