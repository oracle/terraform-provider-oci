package datascience

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

type noOpWorkRequestSigner struct{}

func (noOpWorkRequestSigner) Sign(*http.Request) error { return nil }

func TestModelDeploymentWaitForWorkRequestStatuses(t *testing.T) {
	workRequestID := "ocid1.datascienceworkrequest.oc1..example"
	modelDeploymentID := "ocid1.datasciencemodeldeployment.oc1..example"
	action := oci_datascience.WorkRequestResourceActionTypeUpdated

	tests := []struct {
		name          string
		status        oci_datascience.WorkRequestStatusEnum
		wantID        bool
		wantErrorText string
	}{
		{name: "succeeded", status: oci_datascience.WorkRequestStatusSucceeded, wantID: true},
		{name: "failed", status: oci_datascience.WorkRequestStatusFailed, wantErrorText: "child update failed"},
		{name: "canceled", status: oci_datascience.WorkRequestStatusCanceled, wantErrorText: "work request did not succeed"},
		{name: "partially succeeded", status: oci_datascience.WorkRequestStatusPartiallySucceeded, wantErrorText: "child update failed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("opc-request-id", "test-request")
				if strings.HasSuffix(r.URL.Path, "/errors") {
					if tt.status == oci_datascience.WorkRequestStatusCanceled {
						_, _ = fmt.Fprint(w, `[]`)
					} else {
						_, _ = fmt.Fprint(w, `[{"message":"child update failed"}]`)
					}
					return
				}

				resourceJSON := ""
				if tt.wantID {
					resourceJSON = fmt.Sprintf(`,"resources":[{"entityType":"model-deployment","actionType":"%s","identifier":"%s"}]`, action, modelDeploymentID)
				}
				_, _ = fmt.Fprintf(w, `{"status":"%s"%s}`, tt.status, resourceJSON)
			}))
			defer server.Close()

			client := oci_datascience.DataScienceClient{BaseClient: oci_common.BaseClient{
				HTTPClient: http.DefaultClient,
				Signer:     noOpWorkRequestSigner{},
				Host:       server.URL,
				BasePath:   "20190101",
				UserAgent:  "terraform-provider-test",
			}}

			identifier, err := modelDeploymentWaitForWorkRequest(
				context.Background(), &workRequestID, "model-deployment", action, time.Second, false, &client)
			if tt.wantID {
				if err != nil || identifier == nil || *identifier != modelDeploymentID {
					t.Fatalf("waiter returned identifier=%v, error=%v; want %q", identifier, err, modelDeploymentID)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErrorText) {
				t.Fatalf("waiter error=%v, want substring %q", err, tt.wantErrorText)
			}
		})
	}
}
