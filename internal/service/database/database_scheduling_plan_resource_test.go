// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package database

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func TestUnitDatabaseSchedulingPlanUpdateWithoutChangesReturnsNormally(t *testing.T) {
	resource := DatabaseSchedulingPlanResource()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"compartment_id":       "ocid1.compartment.oc1..test",
		"resource_id":          "ocid1.resource.oc1..test",
		"scheduling_policy_id": "ocid1.schedulingpolicy.oc1..test",
		"service_type":         "EXADATA",
	})
	crud := &DatabaseSchedulingPlanResourceCrud{}
	crud.D = data

	if err := crud.Update(); err != nil {
		t.Fatalf("Update() unexpected error: %v", err)
	}
}

func TestUnitDatabaseSchedulingPlanCompartmentUpdateRefreshesState(t *testing.T) {
	const (
		planID         = "ocid1.schedulingplan.oc1..test"
		newCompartment = "ocid1.compartment.oc1..new"
	)
	baseClient := oci_common.DefaultBaseClientWithSigner(noopSchedulingPlanSigner{})
	baseClient.Host = "https://database.example.test"
	baseClient.BasePath = "20160918"
	baseClient.UserAgent = "terraform-provider-oci-test"
	baseClient.HTTPClient = schedulingPlanDispatcher{planID: planID, compartmentID: newCompartment}
	client := &oci_database.DatabaseClient{BaseClient: baseClient}
	data := schema.TestResourceDataRaw(t, DatabaseSchedulingPlanResource().Schema, map[string]interface{}{
		"compartment_id":       newCompartment,
		"resource_id":          "ocid1.resource.oc1..test",
		"scheduling_policy_id": "ocid1.schedulingpolicy.oc1..test",
		"service_type":         "EXADATA",
	})
	data.SetId(planID)
	crud := &DatabaseSchedulingPlanResourceCrud{Client: client}
	crud.D = data

	if err := crud.updateCompartmentAndRefresh(newCompartment); err != nil {
		t.Fatalf("updateCompartmentAndRefresh() unexpected error: %v", err)
	}
	if crud.Res == nil || crud.Res.CompartmentId == nil || *crud.Res.CompartmentId != newCompartment {
		t.Fatalf("updated scheduling-plan state was not refreshed: %#v", crud.Res)
	}
	if err := crud.SetData(); err != nil {
		t.Fatalf("SetData() after update unexpected error: %v", err)
	}
}

type noopSchedulingPlanSigner struct{}

func (noopSchedulingPlanSigner) Sign(*http.Request) error { return nil }

type schedulingPlanDispatcher struct {
	planID        string
	compartmentID string
}

func (d schedulingPlanDispatcher) Do(request *http.Request) (*http.Response, error) {
	body := ""
	status := http.StatusOK
	switch {
	case request.Method == http.MethodPost && request.URL.Path == "/20160918/schedulingPlans/"+d.planID+"/actions/changeCompartment":
		body = `{}`
	case request.Method == http.MethodGet && request.URL.Path == "/20160918/schedulingPlans/"+d.planID:
		body = `{"id":"` + d.planID + `","compartmentId":"` + d.compartmentID + `"}`
	default:
		status = http.StatusNotFound
		body = `{"message":"unexpected request"}`
	}
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    request,
	}, nil
}
