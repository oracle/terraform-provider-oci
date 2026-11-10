// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Exercise SDK serialization without credentials or network requests.
type coreWorkRequestTestTransport func(*http.Request) (*http.Response, error)

func (f coreWorkRequestTestTransport) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type coreWorkRequestTestSigner struct{}

func (coreWorkRequestTestSigner) Sign(*http.Request) error { return nil }

func coreWorkRequestTestResponse(r *http.Request, status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    r,
	}
}

func newCoreWorkRequestTestBaseClient(transport coreWorkRequestTestTransport) oci_common.BaseClient {
	return oci_common.BaseClient{
		HTTPClient: transport,
		Signer:     coreWorkRequestTestSigner{},
		Host:       "https://compute.example.invalid",
		BasePath:   "20160918",
		UserAgent:  "terraform-provider-test",
	}
}

func newVolumeAttachmentTestCrud(t *testing.T, transport coreWorkRequestTestTransport) *CoreVolumeAttachmentResourceCrud {
	t.Helper()
	base := newCoreWorkRequestTestBaseClient(transport)
	d := schema.TestResourceDataRaw(t, CoreVolumeAttachmentResource().Schema, map[string]interface{}{
		"attachment_type": "paravirtualized",
		"instance_id":     "instance-id",
		"volume_id":       "volume-id",
	})
	return &CoreVolumeAttachmentResourceCrud{
		BaseCrud:               tfresource.BaseCrud{D: d},
		Client:                 &oci_core.ComputeClient{BaseClient: base},
		WorkRequestClient:      &oci_work_requests.WorkRequestClient{BaseClient: base},
		DisableNotFoundRetries: true,
	}
}

func TestCoreVolumeAttachmentWorkRequestCompletion(t *testing.T) {
	const relatedInstance = `[{"entityType":"instance","actionType":"RELATED","identifier":"instance-id"}]`
	for _, operation := range []string{"create", "delete"} {
		for _, tc := range []struct {
			name      string
			status    string
			resources string
			errors    string
			wantError string
		}{
			{name: "related instance", status: "SUCCEEDED", resources: relatedInstance},
			{name: "no resources", status: "SUCCEEDED", resources: `[]`},
			{name: "failed", status: "FAILED", resources: relatedInstance, errors: `[{"message":"volume operation failed"}]`, wantError: "volume operation failed"},
			{name: "failed with attachment identifiers", status: "FAILED", resources: `[{"entityType":"volumeattachment","actionType":"CREATED","identifier":"attachment-id"},{"entityType":"volumeattachment","actionType":"DELETED","identifier":"attachment-id"}]`, errors: `[]`, wantError: "FAILED"},
			{name: "canceled without errors", status: "CANCELED", resources: relatedInstance, errors: `[]`, wantError: "CANCELED"},
		} {
			t.Run(operation+"/"+tc.name, func(t *testing.T) {
				var calls []string
				s := newVolumeAttachmentTestCrud(t, func(r *http.Request) (*http.Response, error) {
					call := r.Method + " " + r.URL.Path
					calls = append(calls, call)
					switch call {
					case "POST /20160918/volumeAttachments", "DELETE /20160918/volumeAttachments/attachment-id":
						response := coreWorkRequestTestResponse(r, http.StatusOK, `{"attachmentType":"paravirtualized","id":"attachment-id","lifecycleState":"ATTACHING"}`)
						response.Header.Set("opc-work-request-id", "work-id")
						return response, nil
					case "GET /20160918/workRequests/work-id":
						return coreWorkRequestTestResponse(r, http.StatusOK, fmt.Sprintf(`{"status":%q,"resources":%s}`, tc.status, tc.resources)), nil
					case "GET /20160918/workRequests/work-id/errors":
						return coreWorkRequestTestResponse(r, http.StatusOK, tc.errors), nil
					default:
						return nil, fmt.Errorf("unexpected request: %s", call)
					}
				})
				var err error
				wantCalls := []string{"POST /20160918/volumeAttachments", "GET /20160918/workRequests/work-id"}
				if operation == "create" {
					err = s.Create()
				} else {
					s.D.SetId("attachment-id")
					wantCalls[0] = "DELETE /20160918/volumeAttachments/attachment-id"
					err = s.Delete()
				}
				if tc.wantError != "" {
					if err == nil || !strings.Contains(err.Error(), tc.wantError) || !strings.Contains(err.Error(), tc.status) || !strings.Contains(err.Error(), "work-id") {
						t.Fatalf("error = %v, want work ID, status %s, and %q", err, tc.status, tc.wantError)
					}
					wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id/errors")
				} else {
					if err != nil {
						t.Fatal(err)
					}
				}
				if s.D.Id() != "attachment-id" {
					t.Fatalf("attachment ID = %q, want attachment-id", s.D.Id())
				}
				if !reflect.DeepEqual(calls, wantCalls) {
					t.Fatalf("requests = %v, want %v", calls, wantCalls)
				}
			})
		}
	}
}

func TestCoreVolumeAttachmentCreateResourceRefresh(t *testing.T) {
	for _, withWorkRequest := range []bool{false, true} {
		t.Run(fmt.Sprintf("work_request_header=%t", withWorkRequest), func(t *testing.T) {
			var calls []string
			s := newVolumeAttachmentTestCrud(t, func(r *http.Request) (*http.Response, error) {
				call := r.Method + " " + r.URL.Path
				calls = append(calls, call)
				switch call {
				case "POST /20160918/volumeAttachments":
					response := coreWorkRequestTestResponse(r, http.StatusOK, `{"attachmentType":"paravirtualized","id":"attachment-id","lifecycleState":"ATTACHING","displayName":"initial response"}`)
					if withWorkRequest {
						response.Header.Set("opc-work-request-id", "work-id")
					}
					return response, nil
				case "GET /20160918/workRequests/work-id":
					return coreWorkRequestTestResponse(r, http.StatusOK, `{"status":"SUCCEEDED","resources":[{"entityType":"instance","actionType":"RELATED","identifier":"instance-id"}]}`), nil
				case "GET /20160918/volumeAttachments/attachment-id":
					return coreWorkRequestTestResponse(r, http.StatusOK, `{"attachmentType":"paravirtualized","id":"attachment-id","lifecycleState":"ATTACHED","displayName":"refreshed attachment"}`), nil
				default:
					return nil, fmt.Errorf("unexpected request: %s", call)
				}
			})
			// Exercise the same wrapper as createCoreVolumeAttachment so ID initialization,
			// lifecycle refresh, and final Terraform state population are covered together.
			if err := tfresource.CreateResource(s.D, s); err != nil {
				t.Fatal(err)
			}
			wantCalls := []string{"POST /20160918/volumeAttachments"}
			if withWorkRequest {
				wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id")
			}
			wantCalls = append(wantCalls, "GET /20160918/volumeAttachments/attachment-id")
			if !reflect.DeepEqual(calls, wantCalls) {
				t.Fatalf("requests = %v, want %v", calls, wantCalls)
			}
			if s.D.Id() != "attachment-id" {
				t.Fatalf("attachment ID = %q, want attachment-id", s.D.Id())
			}
			if state := s.D.Get("state"); state != "ATTACHED" {
				t.Fatalf("state = %v, want ATTACHED", state)
			}
			if name := s.D.Get("display_name"); name != "refreshed attachment" {
				t.Fatalf("display_name = %v, want refreshed attachment", name)
			}
		})
	}
}

func TestCoreVolumeAttachmentWorkRequestPendingStates(t *testing.T) {
	for _, pending := range []string{"ACCEPTED", "IN_PROGRESS", "CANCELING"} {
		t.Run(pending, func(t *testing.T) {
			var ids []string
			s := newVolumeAttachmentTestCrud(t, func(r *http.Request) (*http.Response, error) {
				ids = append(ids, strings.TrimPrefix(r.URL.Path, "/20160918/workRequests/"))
				status := "SUCCEEDED"
				if len(ids) == 1 {
					status = pending
				}
				return coreWorkRequestTestResponse(r, http.StatusOK, fmt.Sprintf(`{"status":%q}`, status)), nil
			})
			if err := s.waitForWorkRequestCompletion("work-id", 5*time.Second); err != nil {
				t.Fatal(err)
			}
			if want := []string{"work-id", "work-id"}; !reflect.DeepEqual(ids, want) {
				t.Fatalf("polled IDs = %v, want %v", ids, want)
			}
		})
	}
}

func TestCoreVolumeAttachmentWorkRequestReadError(t *testing.T) {
	s := newVolumeAttachmentTestCrud(t, func(r *http.Request) (*http.Response, error) {
		return coreWorkRequestTestResponse(r, http.StatusBadRequest, `{"code":"InvalidParameter","message":"cannot read work request"}`), nil
	})
	err := s.waitForWorkRequestCompletion("work-id", time.Second)
	if err == nil || !strings.Contains(err.Error(), "cannot read work request") || !strings.Contains(err.Error(), "work-id") {
		t.Fatalf("error = %v, want work ID and service error", err)
	}
}

func TestCoreVolumeAttachmentWorkRequestTimeout(t *testing.T) {
	s := newVolumeAttachmentTestCrud(t, func(r *http.Request) (*http.Response, error) {
		return coreWorkRequestTestResponse(r, http.StatusOK, `{"status":"IN_PROGRESS"}`), nil
	})
	err := s.waitForWorkRequestCompletion("work-id", 20*time.Millisecond)
	if err == nil || !strings.Contains(err.Error(), "work-id") {
		t.Fatalf("error = %v, want timeout with work ID", err)
	}
	if !strings.Contains(err.Error(), "timeout") && !strings.Contains(err.Error(), "deadline exceeded") {
		t.Fatalf("error = %v, want timeout or context deadline exceeded", err)
	}
}

func TestCoreVolumeAttachmentDeleteResourceWorkRequestErrors(t *testing.T) {
	testCoreDeleteResourceWorkRequestErrors(t, "/20160918/volumeAttachments/attachment-id",
		`{"attachmentType":"paravirtualized","id":"attachment-id","lifecycleState":"DETACHED"}`,
		func(t *testing.T, transport coreWorkRequestTestTransport) (*schema.ResourceData, func() error) {
			s := newVolumeAttachmentTestCrud(t, transport)
			s.D.SetId("attachment-id")
			return s.D, func() error { return tfresource.DeleteResource(s.D, s) }
		})
}

// Run through the generic delete wrappers: testing CRUD.Delete alone misses their
// resource-missing error handling, which must not discard work-request failures.
func testCoreDeleteResourceWorkRequestErrors(t *testing.T, resourcePath, deletedBody string,
	newDelete func(*testing.T, coreWorkRequestTestTransport) (*schema.ResourceData, func() error)) {
	t.Helper()
	const missingBody = `{"code":"NotAuthorizedOrNotFound","message":"not found"}`
	for _, tc := range []struct {
		name          string
		deleteMissing bool
		noWorkRequest bool
		workMissing   bool
		workStatus    string
		errorsMissing bool
		getMissing    bool
		wantError     string
	}{
		{name: "work request 404", workMissing: true, wantError: "not found"},
		{name: "failed work request mentions not found", workStatus: "FAILED", wantError: "dependency not found"},
		{name: "canceled work request mentions not found", workStatus: "CANCELED", wantError: "dependency not found"},
		{name: "work request error list 404", workStatus: "FAILED", errorsMissing: true, wantError: "could not retrieve work request errors"},
		{name: "succeeded then resource deleted", workStatus: "SUCCEEDED"},
		{name: "succeeded then resource 404", workStatus: "SUCCEEDED", getMissing: true},
		{name: "delete resource 404", deleteMissing: true},
		{name: "without work request header", noWorkRequest: true},
	} {
		preserveSettings := []string{"false"}
		if tc.wantError != "" {
			// Preservation mode must retain the failure as well as the resource ID.
			preserveSettings = append(preserveSettings, "true")
		}
		for _, preserve := range preserveSettings {
			t.Run(tc.name+"/preserve="+preserve, func(t *testing.T) {
				t.Setenv("TF_PRESERVE_STATE_ON_404", preserve)
				var calls []string
				d, deleteResource := newDelete(t, func(r *http.Request) (*http.Response, error) {
					call := r.Method + " " + r.URL.Path
					calls = append(calls, call)
					switch call {
					case "DELETE " + resourcePath:
						if tc.deleteMissing {
							return coreWorkRequestTestResponse(r, http.StatusNotFound, missingBody), nil
						}
						response := coreWorkRequestTestResponse(r, http.StatusNoContent, "")
						if !tc.noWorkRequest {
							response.Header.Set("opc-work-request-id", "work-id")
						}
						return response, nil
					case "GET /20160918/workRequests/work-id":
						if tc.workMissing {
							return coreWorkRequestTestResponse(r, http.StatusNotFound, missingBody), nil
						}
						return coreWorkRequestTestResponse(r, http.StatusOK, fmt.Sprintf(`{"status":%q,"resources":[]}`, tc.workStatus)), nil
					case "GET /20160918/workRequests/work-id/errors":
						if tc.errorsMissing {
							return coreWorkRequestTestResponse(r, http.StatusNotFound, missingBody), nil
						}
						return coreWorkRequestTestResponse(r, http.StatusOK, `[{"message":"dependency not found"}]`), nil
					case "GET " + resourcePath:
						if tc.getMissing {
							return coreWorkRequestTestResponse(r, http.StatusNotFound, missingBody), nil
						}
						return coreWorkRequestTestResponse(r, http.StatusOK, deletedBody), nil
					default:
						return nil, fmt.Errorf("unexpected request: %s", call)
					}
				})
				originalID := d.Id()
				err := deleteResource()
				if tc.wantError != "" {
					if err == nil || !strings.Contains(err.Error(), tc.wantError) || !strings.Contains(err.Error(), "work-id") {
						t.Fatalf("error = %v, want work-id and %q", err, tc.wantError)
					}
					if d.Id() != originalID {
						t.Fatalf("resource ID = %q, want preserved %q", d.Id(), originalID)
					}
				} else {
					if err != nil {
						t.Fatal(err)
					}
					if d.Id() != "" {
						t.Fatalf("resource ID = %q, want empty after resource deletion", d.Id())
					}
				}
				wantCalls := []string{"DELETE " + resourcePath}
				if !tc.deleteMissing {
					if !tc.noWorkRequest {
						wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id")
						if tc.workStatus == "FAILED" || tc.workStatus == "CANCELED" {
							wantCalls = append(wantCalls, "GET /20160918/workRequests/work-id/errors")
						}
					}
					if tc.wantError == "" {
						wantCalls = append(wantCalls, "GET "+resourcePath)
					}
				}
				if !reflect.DeepEqual(calls, wantCalls) {
					t.Fatalf("requests = %v, want %v", calls, wantCalls)
				}
			})
		}
	}
}
