// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestCertificatesManagementCertificateImportedUnknownPemDiff(t *testing.T) {
	certificateConfig := &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"config_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cert_chain_pem": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_pem": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key_pem": {
				Type:     schema.TypeString,
				Optional: true,
			},
		}},
	}
	resource := &schema.Resource{
		Schema:        map[string]*schema.Schema{"certificate_config": certificateConfig},
		CustomizeDiff: certificatesManagementCertificateCustomizeDiff,
	}
	unknown := cty.UnknownVal(cty.String)
	config := terraform.NewResourceConfigShimmed(cty.ObjectVal(map[string]cty.Value{
		"certificate_config": cty.ListVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{
			"config_type":     cty.StringVal("IMPORTED"),
			"cert_chain_pem":  unknown,
			"certificate_pem": unknown,
			"private_key_pem": unknown,
		})}),
	}), resource.CoreConfigSchema())

	if _, err := resource.Diff(context.Background(), nil, config, nil); err != nil {
		t.Fatalf("Diff() rejected unknown imported certificate material: %v", err)
	}
}

type certificateTestTransport func(*http.Request) (*http.Response, error)

func (f certificateTestTransport) Do(request *http.Request) (*http.Response, error) {
	return f(request)
}

type certificateTestSigner struct{}

func (certificateTestSigner) Sign(*http.Request) error { return nil }

func certificateTestResponse(request *http.Request, status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    request,
	}
}

func TestCertificatesManagementCertificateRestoresPendingNameConflict(t *testing.T) {
	const (
		certificateID = "ocid1.certificate.oc1..pending"
		compartmentID = "ocid1.compartment.oc1..example"
		certificate   = `{"id":"ocid1.certificate.oc1..pending","name":"imported-cert","timeCreated":"2026-01-01T00:00:00Z","lifecycleState":"ACTIVE","compartmentId":"ocid1.compartment.oc1..example","configType":"IMPORTED"}`
	)

	resource := CertificatesManagementCertificateResource()
	config := map[string]interface{}{
		"compartment_id": compartmentID,
		"name":           "imported-cert",
		"certificate_config": []interface{}{map[string]interface{}{
			"config_type":     "IMPORTED",
			"cert_chain_pem":  "cert-chain-pem",
			"certificate_pem": "certificate-pem",
			"private_key_pem": "private-key-pem",
			"version_name":    "restored-version",
		}},
	}
	diff, err := resource.Diff(context.Background(), nil, terraform.NewResourceConfigRaw(config), nil)
	if err != nil {
		t.Fatal(err)
	}
	data, err := schema.InternalMap(resource.Schema).Data(&terraform.InstanceState{Attributes: map[string]string{}}, diff)
	if err != nil {
		t.Fatal(err)
	}

	var calls []string
	client := &oci_certificates_management.CertificatesManagementClient{BaseClient: oci_common.BaseClient{
		Host:      "https://certificates.example.invalid",
		BasePath:  "20210224",
		Signer:    certificateTestSigner{},
		UserAgent: "terraform-provider-oci-unit-test",
		HTTPClient: certificateTestTransport(func(request *http.Request) (*http.Response, error) {
			call := request.Method + " " + request.URL.Path
			calls = append(calls, call)
			switch {
			case call == "POST /20210224/certificates" && len(calls) == 1:
				return certificateTestResponse(request, http.StatusBadRequest, `{"code":"InvalidParameter","message":"A certificate with the name 'imported-cert' already exists."}`), nil
			case call == "GET /20210224/certificates":
				if request.URL.Query().Get("name") != "imported-cert" || request.URL.Query().Get("compartmentId") != compartmentID || request.URL.Query().Get("lifecycleState") != "PENDING_DELETION" {
					t.Fatalf("unexpected recovery query: %s", request.URL.RawQuery)
				}
				pending := strings.Replace(certificate, `"lifecycleState":"ACTIVE"`, `"lifecycleState":"PENDING_DELETION"`, 1)
				return certificateTestResponse(request, http.StatusOK, fmt.Sprintf(`{"items":[%s]}`, pending)), nil
			case call == "POST /20210224/certificates/"+certificateID+"/actions/cancelDeletion":
				return certificateTestResponse(request, http.StatusNoContent, ""), nil
			case call == "GET /20210224/certificates/"+certificateID:
				return certificateTestResponse(request, http.StatusOK, certificate), nil
			case call == "PUT /20210224/certificates/"+certificateID:
				body, err := io.ReadAll(request.Body)
				if err != nil {
					t.Fatal(err)
				}
				if !strings.Contains(string(body), `"certificatePem":"certificate-pem"`) {
					t.Fatalf("restored certificate was not rotated with configured PEM: %s", body)
				}
				return certificateTestResponse(request, http.StatusOK, certificate), nil
			case call == "GET /20210224/certificates/"+certificateID+"/versions":
				return certificateTestResponse(request, http.StatusOK, fmt.Sprintf(`{"items":[{"certificateId":%q,"versionNumber":2,"versionName":"restored-version","timeCreated":"2026-01-01T00:00:00Z","stages":["CURRENT"]}]}`, certificateID)), nil
			default:
				t.Fatalf("unexpected API call: %s", call)
				return nil, fmt.Errorf("unexpected API call")
			}
		}),
	}}
	crud := &CertificatesManagementCertificateResourceCrud{
		BaseCrud:               tfresource.BaseCrud{D: data},
		Client:                 client,
		DisableNotFoundRetries: true,
	}

	if err := tfresource.CreateResource(data, crud); err != nil {
		t.Fatal(err)
	}
	if data.Id() != certificateID {
		t.Fatalf("restored certificate ID = %q, want %q", data.Id(), certificateID)
	}
	if len(calls) != 7 {
		t.Fatalf("API calls = %v", calls)
	}
}

func TestCertificatesManagementCertificateCreateWaitsForRecoveryUpdate(t *testing.T) {
	for _, state := range (&CertificatesManagementCertificateResourceCrud{}).CreatedPending() {
		if state == string(oci_certificates_management.CertificateLifecycleStateUpdating) {
			return
		}
	}

	t.Fatal("UPDATING must remain pending while a restored certificate rotates")
}

func TestCertificatesManagementCertificateRecognizesRestoredVersion(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(42),
		Subject:      pkix.Name{CommonName: "restored.example"},
		NotBefore:    time.Now().Add(-time.Minute),
		NotAfter:     time.Now().Add(time.Hour),
	}
	certificateDer, err := x509.CreateCertificate(rand.Reader, template, template, publicKey, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	certificatePem := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificateDer}))

	data := schema.TestResourceDataRaw(t, CertificatesManagementCertificateResource().Schema, map[string]interface{}{
		"compartment_id": "ocid1.compartment.oc1..example",
		"name":           "imported-cert",
		"certificate_config": []interface{}{map[string]interface{}{
			"config_type":     "IMPORTED",
			"certificate_pem": certificatePem,
			"version_name":    "imported-cert-content-hash",
		}},
	})
	data.SetId("ocid1.certificate.oc1..pending")
	client := &oci_certificates_management.CertificatesManagementClient{BaseClient: oci_common.BaseClient{
		Host:      "https://certificates.example.invalid",
		BasePath:  "20210224",
		Signer:    certificateTestSigner{},
		UserAgent: "terraform-provider-oci-unit-test",
		HTTPClient: certificateTestTransport(func(request *http.Request) (*http.Response, error) {
			body := `{"items":[{"certificateId":"ocid1.certificate.oc1..pending","versionNumber":2,"versionName":"imported-cert-content-hash","serialNumber":"00 2A","timeCreated":"2026-01-01T00:00:00Z","stages":["CURRENT"]}]}`
			return certificateTestResponse(request, http.StatusOK, body), nil
		}),
	}}
	crud := &CertificatesManagementCertificateResourceCrud{
		BaseCrud:               tfresource.BaseCrud{D: data},
		Client:                 client,
		DisableNotFoundRetries: true,
	}

	alreadyCurrent, err := crud.importedCertificateVersionAlreadyCurrent("certificate_config.0.%s")
	if err != nil {
		t.Fatal(err)
	}
	if !alreadyCurrent {
		t.Fatal("matching restored certificate version was not recognized")
	}
}

func TestCertificatesManagementCertificateDeletionPendingDays(t *testing.T) {
	validate := CertificatesManagementCertificateResource().Schema["deletion_pending_days"].ValidateFunc
	if _, errors := validate(0, "deletion_pending_days"); len(errors) == 0 {
		t.Fatal("zero deletion pending days must be rejected")
	}
	if _, errors := validate(1, "deletion_pending_days"); len(errors) != 0 {
		t.Fatalf("one deletion pending day was rejected: %v", errors)
	}

	data := schema.TestResourceDataRaw(t, CertificatesManagementCertificateResource().Schema, map[string]interface{}{
		"deletion_pending_days": 1,
	})
	data.SetId("ocid1.certificate.oc1..example")
	beforeDelete := time.Now().UTC()
	client := &oci_certificates_management.CertificatesManagementClient{BaseClient: oci_common.BaseClient{
		Host:      "https://certificates.example.invalid",
		BasePath:  "20210224",
		Signer:    certificateTestSigner{},
		UserAgent: "terraform-provider-oci-unit-test",
		HTTPClient: certificateTestTransport(func(request *http.Request) (*http.Response, error) {
			var body struct {
				TimeOfDeletion time.Time `json:"timeOfDeletion"`
			}
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			minimum := beforeDelete.Add(24 * time.Hour)
			maximum := minimum.Add(time.Second)
			if body.TimeOfDeletion.Before(minimum) || body.TimeOfDeletion.After(maximum) {
				t.Fatalf("timeOfDeletion = %s, want between %s and %s", body.TimeOfDeletion, minimum, maximum)
			}
			return certificateTestResponse(request, http.StatusNoContent, ""), nil
		}),
	}}
	crud := &CertificatesManagementCertificateResourceCrud{
		BaseCrud:               tfresource.BaseCrud{D: data},
		Client:                 client,
		DisableNotFoundRetries: true,
	}

	if err := crud.Delete(); err != nil {
		t.Fatal(err)
	}
}
