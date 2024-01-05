// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
)

func setUp(t *testing.T) {

	SweeperResourceCompartmentIdMap = map[string]map[string][]string{"compartment-1_ocid": {
		"vcn": []string{"vcn-1_ocid"},
	},
	}
}

func tearDown(t *testing.T) {
	SweeperResourceCompartmentIdMap = map[string]map[string][]string{}
}

func TestUnitAddResourceIdToSweeperResourceIdMap(t *testing.T) {

	testsweeperResourceCompartmentIdMapWithCompID := map[string]map[string][]string{
		"compartment-1_ocid": {
			"vcn": []string{"vcn-1_ocid", "vcn-2_ocid"},
		},
	}
	testsweeperResourceCompartmentIdMapWithNonExistentCompID := map[string]map[string][]string{
		"nonexistent_compID": {
			"vcn": []string{"vcn-1_ocid"},
		},
	}
	testsweeperResourceCompartmentIdMapWithNonExistentResourceType := map[string]map[string][]string{
		"compartment-1_ocid": {
			"vcn":                       []string{"vcn-1_ocid"},
			"nonexistent_resource_type": []string{"vcn-1_ocid"},
		},
	}

	type args struct {
		compartmentId string
		resourceType  string
		resourceId    string
	}

	tests := []struct {
		name string
		args args
		want map[string]map[string][]string
	}{
		{
			name: "Test with nonexistent value for compartmentId",
			args: args{
				compartmentId: "nonexistent_compID",
				resourceType:  "vcn",
				resourceId:    "vcn-1_ocid",
			},
			want: testsweeperResourceCompartmentIdMapWithNonExistentCompID,
		},
		{
			name: "Test for existing compartmentId but nonexistent resourceType",
			args: args{
				compartmentId: "compartment-1_ocid",
				resourceType:  "nonexistent_resource_type",
				resourceId:    "vcn-1_ocid",
			},
			want: testsweeperResourceCompartmentIdMapWithNonExistentResourceType,
		},
		{
			name: "Test for existing compartmentId and resourceType",
			args: args{
				compartmentId: "compartment-1_ocid",
				resourceType:  "vcn",
				resourceId:    "vcn-2_ocid",
			},
			want: testsweeperResourceCompartmentIdMapWithCompID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUp(t)
			defer tearDown(t)
			AddResourceIdToSweeperResourceIdMap(tt.args.compartmentId, tt.args.resourceType, tt.args.resourceId)
			if !reflect.DeepEqual(SweeperResourceCompartmentIdMap, tt.want) {
				t.Errorf("AddResourceIdToSweeperResourceIdMap() = %v, want %v", SweeperResourceCompartmentIdMap, tt.want)
			}
		})

	}

}

func TestUnitInSweeperExcludeList(t *testing.T) {
	type args struct {
		sweeperName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sweeperName present in sweep_exclude_list",
			args: args{
				sweeperName: "EmailSuppression",
			},
			want: true,
		},
		{
			name: "sweeperName not present in sweep_exclude_list",
			args: args{
				sweeperName: "AbsentSweeper",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("sweep_exclude_list", "EmailSuppression,EmailSender,BudgetBudget,AiAnomalyDetectionModel,AiAnomalyDetectionDataAsset,AiAnomalyDetectionProject")
			defer os.Unsetenv("sweep_exclude_list")
			if got := InSweeperExcludeList(tt.args.sweeperName); got != tt.want {
				t.Errorf("InSweeperExcludeList() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestUnitGetAvalabilityDomains(t *testing.T) {
	type args struct {
		compartmentId string
	}

	identityClientListAvailabilityDomains = func(client *oci_identity.IdentityClient, adRequest oci_identity.ListAvailabilityDomainsRequest) (oci_identity.ListAvailabilityDomainsResponse, error) {
		Id := "1"
		Name := "Domain1"
		CompartmentId := "compID"
		opcRequestId := "test_opc"
		if *adRequest.CompartmentId == "compID" {
			return oci_identity.ListAvailabilityDomainsResponse{
				OpcRequestId: &opcRequestId,
				RawResponse:  &http.Response{},
				Items: []oci_identity.AvailabilityDomain{{
					CompartmentId: &CompartmentId,
					Id:            &Id,
					Name:          &Name,
				},
				},
			}, nil
		}
		return oci_identity.ListAvailabilityDomainsResponse{}, errors.New("invalid compId")
	}

	tfProviderConfigVar = func(d *schema.ResourceData) (interface{}, error) {
		sdkClientMap := map[string]interface{}{
			"oci_identity.IdentityClient": &oci_identity.IdentityClient{},
		}
		return &client.OracleClients{
			SdkClientMap: sdkClientMap,
		}, nil
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test with valid compartmentID",
			args: args{
				compartmentId: "compID",
			},

			want:    map[string]string{"1": "Domain1"},
			wantErr: false,
		},
		{
			name: "Test with invalid compartmentID",
			args: args{
				compartmentId: "invalid_compID",
			},
			want:    map[string]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAvalabilityDomains(tt.args.compartmentId)
			fmt.Println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvalabilityDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvalabilityDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGetResourceIdsToSweep(t *testing.T) {
	type args struct {
		compartmentId string
		resourceName  string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test for nonexisting compartmentId",
			args: args{
				compartmentId: "non_existing_compID",
				resourceName:  "vcn",
			},
			want: nil,
		},
		{
			name: "Test for existing compartmentId but nonexisting resourceName",
			args: args{
				compartmentId: "compartment-1_ocid",
				resourceName:  "non_existing_resource_name",
			},
			want: nil,
		},
		{
			name: "Test with existing compartmentId and resourceName",
			args: args{
				compartmentId: "compartment-1_ocid",
				resourceName:  "vcn",
			},
			want: []string{"vcn-1_ocid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUp(t)
			defer tearDown(t)
			if got := GetResourceIdsToSweep(tt.args.compartmentId, tt.args.resourceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResourceIdsToSweep() = %v, want %v", got, tt.want)
			}
		})
	}
}
