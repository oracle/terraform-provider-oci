// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"testing"
)

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_toString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"toString",
			fields{TfVersion11},
			"0.11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.toString(); got != tt.want {
				t.Errorf("TfHclVersion11.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getVarHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion11},
			args{"variableName"},
			"\"${var.variableName}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.getVarHclString(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion11.getVarHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getDataSourceHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion11},
			args{"datasource_type", "datasource_name"},
			"\"${data.datasource_type.datasource_name}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.getDataSourceHclString(tt.args.datasourceType, tt.args.datasourceName); got != tt.want {
				t.Errorf("TfHclVersion11.getDataSourceHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getSingleExpHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion11},
			args{"exp"},
			"" +
				"\"${exp}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.getSingleExpHclString(tt.args.expString); got != tt.want {
				t.Errorf("TfHclVersion11.getSingleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion11_getDoubleExpHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion11},
			args{"exp1", "exp2"},
			"\"${exp1.exp2}\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion11{
				Value: tt.fields.Value,
			}
			if got := tfversion.getDoubleExpHclString(tt.args.expString1, tt.args.expString2); got != tt.want {
				t.Errorf("TfHclVersion11.getDoubleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_toString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"versionTest",
			fields{TfVersion12},
			"0.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.toString(); got != tt.want {
				t.Errorf("TfHclVersion12.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getVarHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion12},
			args{"variableName"},
			"var.variableName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.getVarHclString(tt.args.varName); got != tt.want {
				t.Errorf("TfHclVersion12.getVarHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getDataSourceHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion12},
			args{"datasource_type", "datasource_name"},
			"data.datasource_type.datasource_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.getDataSourceHclString(tt.args.datasourceType, tt.args.datasourceName); got != tt.want {
				t.Errorf("TfHclVersion12.getDataSourceHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getSingleExpHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion12},
			args{"exp"},
			"exp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.getSingleExpHclString(tt.args.expString); got != tt.want {
				t.Errorf("TfHclVersion12.getSingleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitTfHclVersion12_getDoubleExpHclString(t *testing.T) {
	type fields struct {
		Value TfVersionEnum
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
			fields{TfVersion12},
			args{"exp1", "exp2"},
			"exp1.exp2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tfversion := &TfHclVersion12{
				Value: tt.fields.Value,
			}
			if got := tfversion.getDoubleExpHclString(tt.args.expString1, tt.args.expString2); got != tt.want {
				t.Errorf("TfHclVersion12.getDoubleExpHclString() = %v, want %v", got, tt.want)
			}
		})
	}
}
