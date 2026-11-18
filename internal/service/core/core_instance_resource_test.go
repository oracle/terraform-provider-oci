// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func newInstanceDeleteTestCrud(t *testing.T, transport coreWorkRequestTestTransport) *CoreInstanceResourceCrud {
	t.Helper()
	base := newCoreWorkRequestTestBaseClient(transport)
	d := schema.TestResourceDataRaw(t, CoreInstanceResource().Schema, map[string]interface{}{
		"preserve_boot_volume":                    true,
		"preserve_data_volumes_created_at_launch": false,
	})
	d.SetId("instance-id")
	return &CoreInstanceResourceCrud{
		BaseCrud:               tfresource.BaseCrud{D: d},
		Client:                 &oci_core.ComputeClient{BaseClient: base},
		WorkRequestClient:      &oci_work_requests.WorkRequestClient{BaseClient: base},
		DisableNotFoundRetries: true,
	}
}

func TestCoreInstanceDeleteWorkRequestCompletion(t *testing.T) {
	for _, tc := range []struct {
		name      string
		statuses  []string
		resources string
		errors    string
		wantError string
	}{
		{name: "succeeded without resources", statuses: []string{"SUCCEEDED"}, resources: `[]`},
		{name: "pending then succeeded", statuses: []string{"ACCEPTED", "IN_PROGRESS", "SUCCEEDED"}, resources: `[]`},
		{name: "failed", statuses: []string{"FAILED"}, resources: `[]`, errors: `[{"message":"termination failed"}]`, wantError: "termination failed"},
		{name: "failed with deleted instance", statuses: []string{"FAILED"}, resources: `[{"entityType":"instance","actionType":"DELETED","identifier":"instance-id"}]`, errors: `[]`, wantError: "FAILED"},
		{name: "canceling then canceled", statuses: []string{"CANCELING", "CANCELED"}, resources: `[]`, errors: `[]`, wantError: "CANCELED"},
		{name: "without work request header"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var calls []string
			polls := 0
			s := newInstanceDeleteTestCrud(t, func(r *http.Request) (*http.Response, error) {
				call := r.Method + " " + r.URL.Path
				calls = append(calls, call)
				switch call {
				case "DELETE /20160918/instances/instance-id":
					if r.URL.Query().Get("preserveBootVolume") != "true" || r.URL.Query().Get("preserveDataVolumesCreatedAtLaunch") != "false" {
						return nil, fmt.Errorf("unexpected preservation flags: %s", r.URL.RawQuery)
					}
					response := coreWorkRequestTestResponse(r, http.StatusNoContent, "")
					if len(tc.statuses) != 0 {
						response.Header.Set("opc-work-request-id", "work-id")
					}
					return response, nil
				case "GET /20160918/workRequests/work-id":
					if polls >= len(tc.statuses) {
						return nil, fmt.Errorf("unexpected work request poll %d", polls)
					}
					status := tc.statuses[polls]
					polls++
					return coreWorkRequestTestResponse(r, http.StatusOK, fmt.Sprintf(`{"status":%q,"resources":%s}`, status, tc.resources)), nil
				case "GET /20160918/workRequests/work-id/errors":
					return coreWorkRequestTestResponse(r, http.StatusOK, tc.errors), nil
				default:
					return nil, fmt.Errorf("unexpected request: %s", call)
				}
			})
			err := s.DeleteWithContext(context.Background())
			wantCalls := []string{"DELETE /20160918/instances/instance-id"}
			for range tc.statuses {
				wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id")
			}
			if tc.wantError != "" {
				status := tc.statuses[len(tc.statuses)-1]
				if err == nil || !strings.Contains(err.Error(), tc.wantError) || !strings.Contains(err.Error(), status) || !strings.Contains(err.Error(), "work-id") {
					t.Fatalf("error = %v, want work ID, status %s, and %q", err, status, tc.wantError)
				}
				wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id/errors")
			} else if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(calls, wantCalls) {
				t.Fatalf("requests = %v, want %v", calls, wantCalls)
			}
			if s.D.Id() != "instance-id" {
				t.Fatalf("instance ID = %q, want instance-id", s.D.Id())
			}
		})
	}
}

func TestCoreInstanceDeleteWorkRequestContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	requestCanceled := make(chan bool, 1)
	s := newInstanceDeleteTestCrud(t, func(r *http.Request) (*http.Response, error) {
		if r.Method == http.MethodDelete {
			response := coreWorkRequestTestResponse(r, http.StatusNoContent, "")
			response.Header.Set("opc-work-request-id", "work-id")
			return response, nil
		}
		// Cancel while polling to prove the caller's context reaches the SDK request.
		cancel()
		select {
		case <-r.Context().Done():
			requestCanceled <- true
			return nil, r.Context().Err()
		case <-time.After(time.Second):
			requestCanceled <- false
			return nil, errors.New("caller cancellation did not reach the work request")
		}
	})
	err := s.DeleteWithContext(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
	if !<-requestCanceled {
		t.Fatal("caller cancellation did not reach the SDK request context")
	}
}

func TestCoreInstanceDeleteResourceWorkRequestErrors(t *testing.T) {
	testCoreDeleteResourceWorkRequestErrors(t, "/20160918/instances/instance-id",
		`{"id":"instance-id","lifecycleState":"TERMINATED"}`,
		func(t *testing.T, transport coreWorkRequestTestTransport) (*schema.ResourceData, func() error) {
			s := newInstanceDeleteTestCrud(t, transport)
			return s.D, func() error { return tfresource.DeleteResourceWithContext(context.Background(), s.D, s) }
		})
}
