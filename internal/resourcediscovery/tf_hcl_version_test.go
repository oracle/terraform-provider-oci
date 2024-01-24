// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package resourcediscovery

import (
	"testing"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_toString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"toString",
			fields{tf_export.TfVersion11},
			"0.11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.ToString(); got != tt.want {
				t.Errorf("TfHclVersion11.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getVarHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		varName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"VarHclString",
			fields{tf_export.TfVersion11},
			args{"variableName"},
			"\"${var.variableName}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetVarHclString(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion11.getVarHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitTfHclVersion11_getReference(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		varName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"reference",
			fields{tf_export.TfVersion11},
			args{"reference"},
			"\"reference\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetReference(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion11.getReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getDataSourceHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		datasourceType string
		datasourceName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"DataSourceHclString",
			fields{tf_export.TfVersion11},
			args{"datasource_type", "datasource_name"},
			"\"${data.datasource_type.datasource_name}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetDataSourceHclString(tt.args.datasourceType, tt.args.datasourceName); got != tt.want {
				t.Errorf("TfHclVersion11.getDataSourceHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getSingleExpHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		expString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"SingleExpHclString",
			fields{tf_export.TfVersion11},
			args{"exp"},
			"" +
				"\"${exp}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetSingleExpHclString(tt.args.expString); got != tt.want {
				t.Errorf("TfHclVersion11.getSingleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getDoubleExpHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		expString1 string
		expString2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"DoubleExpHclString",
			fields{tf_export.TfVersion11},
			args{"exp1", "exp2"},
			"\"${exp1.exp2}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetDoubleExpHclString(tt.args.expString1, tt.args.expString2); got != tt.want {
				t.Errorf("TfHclVersion11.getDoubleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_toString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"versionTest",
			fields{tf_export.TfVersion12},
			"0.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.ToString(); got != tt.want {
				t.Errorf("TfHclVersion12.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getVarHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		varName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"VarHclString",
			fields{tf_export.TfVersion12},
			args{"variableName"},
			"var.variableName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetVarHclString(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion12.getVarHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitTfHclVersion12_getReference(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		varName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"reference",
			fields{tf_export.TfVersion12},
			args{"reference"},
			"reference",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetReference(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion12.getReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getDataSourceHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		datasourceType string
		datasourceName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"DataSourceHclString",
			fields{tf_export.TfVersion12},
			args{"datasource_type", "datasource_name"},
			"data.datasource_type.datasource_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetDataSourceHclString(tt.args.datasourceType, tt.args.datasourceName); got != tt.want {
				t.Errorf("TfHclVersion12.getDataSourceHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getSingleExpHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		expString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"SingleExpHclString",
			fields{tf_export.TfVersion12},
			args{"exp"},
			"exp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetSingleExpHclString(tt.args.expString); got != tt.want {
				t.Errorf("TfHclVersion12.getSingleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getDoubleExpHclString(t *testing.T) {
	type fields struct {
		Value tf_export.TfVersionEnum
	}
	type args struct {
		expString1 string
		expString2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"DoubleExpHclString",
			fields{tf_export.TfVersion12},
			args{"exp1", "exp2"},
			"exp1.exp2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &tf_export.TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.GetDoubleExpHclString(tt.args.expString1, tt.args.expString2); got != tt.want {
				t.Errorf("TfHclVersion12.getDoubleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}
