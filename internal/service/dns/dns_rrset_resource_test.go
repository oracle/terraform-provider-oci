// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Exercise Create through the SDK and inspect the serialized request without network calls.
func TestDnsRrsetCreatePreservesTTL(t *testing.T) {
	for _, tc := range []struct {
		name string
		ttls []int
	}{
		{name: "zero", ttls: []int{0}},
		{name: "one", ttls: []int{1}},
		{name: "nonzero", ttls: []int{3600}},
		{name: "multiple zero records", ttls: []int{0, 0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			items := make([]interface{}, len(tc.ttls))
			want := make(map[string]int, len(tc.ttls))
			for i, ttl := range tc.ttls {
				rdata := fmt.Sprintf("192.0.2.%d", i+10)
				items[i] = map[string]interface{}{"domain": "www.example.com", "rtype": "A", "rdata": rdata, "ttl": ttl}
				want[rdata] = ttl
			}
			data := schema.TestResourceDataRaw(t, DnsRrsetResource().Schema, map[string]interface{}{
				"zone_name_or_id": "zone-test-id", "domain": "www.example.com", "rtype": "A", "items": items,
			})
			requests := 0
			client := &oci_dns.DnsClient{BaseClient: oci_common.DefaultBaseClientWithSigner(rrsetTestSigner{})}
			client.Host, client.BasePath = "https://dns.invalid", "/20180115"
			// Capture the request after SDK serialization, before it would reach DNS.
			client.HTTPClient = rrsetTestTransport(func(request *http.Request) (*http.Response, error) {
				requests++
				if request.Method != http.MethodPatch || request.URL.Path != "/20180115/zones/zone-test-id/records/www.example.com/A" {
					return nil, fmt.Errorf("unexpected request: %s %s", request.Method, request.URL.Path)
				}
				var payload struct {
					Items []map[string]interface{} `json:"items"`
				}
				if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
					return nil, err
				}
				if len(payload.Items) != len(want)+1 {
					t.Fatalf("operations = %d; want %d", len(payload.Items), len(want)+1)
				}
				// The empty-RRSet precondition must stay ahead of the ADD operations.
				precondition := payload.Items[0]
				if precondition["operation"] != "PROHIBIT" || precondition["domain"] != "www.example.com" || precondition["rtype"] != "A" {
					t.Fatalf("RRSet creation precondition changed: %#v", precondition)
				}
				for _, item := range payload.Items[1:] {
					if item["operation"] != "ADD" || item["domain"] != "www.example.com" || item["rtype"] != "A" {
						t.Errorf("unexpected ADD operation: %#v", item)
					}
					rdata, _ := item["rdata"].(string)
					expected, exists := want[rdata]
					if !exists {
						t.Fatalf("unexpected or duplicate record: %#v", item)
					}
					// JSON numbers decode as float64; require a numeric TTL, including zero.
					ttl, present := item["ttl"].(float64)
					if !present || ttl != float64(expected) {
						t.Errorf("ADD %s serialized TTL = %#v; want numeric %d", rdata, item["ttl"], expected)
					}
					delete(want, rdata)
				}
				return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"items":[]}`)), Request: request}, nil
			})
			crud := DnsRrsetResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: client, DisableNotFoundRetries: true}
			if err := crud.Create(); err != nil {
				t.Fatal(err)
			}
			if requests != 1 || len(want) != 0 {
				t.Fatalf("requests = %d, unobserved records = %#v", requests, want)
			}
		})
	}
}

// Missing TTLs fail schema validation; an explicit zero must remain valid.
func TestDnsRrsetTTLRequired(t *testing.T) {
	for _, tc := range []struct {
		name      string
		ttl       interface{}
		wantError bool
	}{
		{name: "zero", ttl: 0},
		{name: "one", ttl: 1},
		{name: "nonzero", ttl: 3600},
		{name: "missing", wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			item := map[string]interface{}{"domain": "www.example.com", "rtype": "A", "rdata": "192.0.2.10"}
			if tc.ttl != nil {
				item["ttl"] = tc.ttl
			}
			config := terraform.NewResourceConfigRaw(map[string]interface{}{
				"zone_name_or_id": "zone-test-id", "domain": "www.example.com", "rtype": "A",
				"items": []interface{}{item},
			})
			diagnostics := DnsRrsetResource().Validate(config)
			if diagnostics.HasError() != tc.wantError {
				t.Fatalf("validation errors = %v; want error = %t", diagnostics, tc.wantError)
			}
			if tc.wantError && !strings.Contains(fmt.Sprint(diagnostics), "ttl") {
				t.Fatalf("expected a missing TTL error, got %v", diagnostics)
			}
		})
	}
}

// Cover Get, Update, and both Delete request paths for public and private RRSets.
// Deprecated compartment values must stay out of requests, including values left in state.
// Delete also covers all-protected records (no request) and an empty record set.
func TestDnsRrsetOmitsCompartmentID(t *testing.T) {
	for _, operation := range []string{"read", "update", "delete", "delete mixed records", "delete protected records", "delete empty records"} {
		for _, compartment := range []string{"omitted", "empty", "populated"} {
			for _, private := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/private=%t", operation, compartment, private), func(t *testing.T) {
					items := []interface{}{
						map[string]interface{}{"domain": "www.example.com", "rtype": "A", "rdata": "192.0.2.10", "ttl": 300, "is_protected": operation == "delete mixed records" || operation == "delete protected records"},
						map[string]interface{}{"domain": "www.example.com", "rtype": "A", "rdata": "192.0.2.11", "ttl": 300, "is_protected": operation == "delete protected records"},
					}
					if operation == "delete empty records" {
						items = []interface{}{}
					}
					config := map[string]interface{}{"zone_name_or_id": "zone-test-id", "domain": "www.example.com", "rtype": "A", "items": items}
					if compartment != "omitted" {
						config["compartment_id"] = ""
						if compartment == "populated" {
							config["compartment_id"] = "ocid1.compartment.oc1..legacy"
						}
					}
					if private {
						config["view_id"] = "view-test-id"
					}
					data := schema.TestResourceDataRaw(t, DnsRrsetResource().Schema, config)
					// Model an explicit empty value retained by Terraform state.
					if compartment == "empty" {
						if err := data.Set("compartment_id", ""); err != nil {
							t.Fatal(err)
						}
						if _, exists := data.GetOkExists("compartment_id"); !exists {
							t.Fatal("empty state fixture is not set")
						}
					}
					data.SetId(getRrsetCompositeId("www.example.com", "A", "zone-test-id"))
					requests := 0
					client := &oci_dns.DnsClient{BaseClient: oci_common.DefaultBaseClientWithSigner(rrsetTestSigner{})}
					client.Host, client.BasePath = "https://dns.invalid", "/20180115"
					client.HTTPClient = rrsetTestTransport(func(request *http.Request) (*http.Response, error) {
						requests++
						if operation == "delete protected records" {
							t.Fatal("an entirely protected RRSet must not be modified")
						}
						if _, present := request.URL.Query()["compartmentId"]; present {
							t.Errorf("deprecated compartmentId sent: %s", request.URL.RawQuery)
						}
						wantMethod := map[string]string{"read": http.MethodGet, "update": http.MethodPut, "delete": http.MethodDelete, "delete mixed records": http.MethodPut, "delete empty records": http.MethodDelete}[operation]
						if request.Method != wantMethod || request.URL.Path != "/20180115/zones/zone-test-id/records/www.example.com/A" {
							t.Errorf("unexpected request: %s %s", request.Method, request.URL.Path)
						}
						wantView := ""
						if private {
							wantView = "view-test-id"
						}
						if got := request.URL.Query().Get("viewId"); got != wantView {
							t.Errorf("viewId = %q; want %q", got, wantView)
						}
						if operation == "update" || operation == "delete mixed records" {
							var payload struct {
								Items []struct {
									Rdata string `json:"rdata"`
									Ttl   int    `json:"ttl"`
								} `json:"items"`
							}
							if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
								return nil, err
							}
							wantCount := 2
							if operation == "delete mixed records" {
								wantCount = 1
							}
							if len(payload.Items) != wantCount {
								t.Fatalf("records = %d; want %d", len(payload.Items), wantCount)
							}
							for _, item := range payload.Items {
								if item.Ttl != 300 {
									t.Errorf("TTL changed: %d", item.Ttl)
								}
								if operation == "delete mixed records" && item.Rdata != "192.0.2.10" {
									t.Errorf("expected only the protected record, got %s", item.Rdata)
								}
							}
						}
						return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"items":[]}`)), Request: request}, nil
					})
					crud := DnsRrsetResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: client, DisableNotFoundRetries: true}
					var err error
					switch operation {
					case "read":
						err = crud.Get()
					case "update":
						err = crud.Update()
					case "delete", "delete mixed records", "delete protected records", "delete empty records":
						// Delete uses DELETE for unprotected records and PUT to retain protected records.
						err = crud.Delete()
					}
					if err != nil {
						t.Fatal(err)
					}
					wantRequests := 1
					if operation == "delete protected records" {
						wantRequests = 0
					}
					if requests != wantRequests {
						t.Fatalf("requests = %d; want %d", requests, wantRequests)
					}
				})
			}
		}
	}
}

type rrsetTestTransport func(*http.Request) (*http.Response, error)

func (transport rrsetTestTransport) Do(request *http.Request) (*http.Response, error) {
	return transport(request)
}

// Signing is unnecessary because the in-memory transport never sends requests.
type rrsetTestSigner struct{}

func (rrsetTestSigner) Sign(*http.Request) error { return nil }
