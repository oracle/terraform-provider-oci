// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
)

const (
	certificateService = "certificate"
)

func WaasCertificateResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createWaasCertificate,
		Read:     readWaasCertificate,
		Update:   updateWaasCertificate,
		Delete:   deleteWaasCertificate,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_data": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key_data": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_trust_verification_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"extensions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_critical": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"issued_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_name": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state_province": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"public_key_info": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"exponent": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject_name": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state_province": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_not_valid_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_not_valid_before": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.CreateResource(d, sync)
}

func readWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

func updateWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type WaasCertificateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.Certificate
	DisableNotFoundRetries bool
}

func (s *WaasCertificateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaasCertificateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waas.CertificateLifecycleStateCreating),
	}
}

func (s *WaasCertificateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waas.CertificateLifecycleStateActive),
	}
}

func (s *WaasCertificateResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waas.CertificateLifecycleStateDeleting),
	}
}

func (s *WaasCertificateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waas.CertificateLifecycleStateDeleted),
	}
}

func (s *WaasCertificateResourceCrud) Create() error {
	request := oci_waas.CreateCertificateRequest{}

	if certificateData, ok := s.D.GetOkExists("certificate_data"); ok {
		tmp := certificateData.(string)
		request.CertificateData = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isTrustVerificationDisabled, ok := s.D.GetOkExists("is_trust_verification_disabled"); ok {
		tmp := isTrustVerificationDisabled.(bool)
		request.IsTrustVerificationDisabled = &tmp
	}

	if privateKeyData, ok := s.D.GetOkExists("private_key_data"); ok {
		tmp := privateKeyData.(string)
		request.PrivateKeyData = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.CreateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *WaasCertificateResourceCrud) Get() error {
	request := oci_waas.GetCertificateRequest{}

	tmp := s.D.Id()
	request.CertificateId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *WaasCertificateResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waas.UpdateCertificateRequest{}

	tmp := s.D.Id()
	request.CertificateId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *WaasCertificateResourceCrud) Delete() error {
	request := oci_waas.DeleteCertificateRequest{}

	tmp := s.D.Id()
	request.CertificateId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, tfresource.WaasService, certificateService, globalvar.DeleteResource)

	_, err := s.Client.DeleteCertificate(context.Background(), request)
	return err
}

func (s *WaasCertificateResourceCrud) SetData() error {
	if s.Res.CertificateData != nil {
		s.D.Set("certificate_data", *s.Res.CertificateData)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	extensions := []interface{}{}
	for _, item := range s.Res.Extensions {
		extensions = append(extensions, CertificateExtensionToMap(item))
	}
	s.D.Set("extensions", extensions)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsTrustVerificationDisabled != nil {
		s.D.Set("is_trust_verification_disabled", *s.Res.IsTrustVerificationDisabled)
	}

	if s.Res.IssuedBy != nil {
		s.D.Set("issued_by", *s.Res.IssuedBy)
	}

	if s.Res.IssuerName != nil {
		s.D.Set("issuer_name", []interface{}{CertificateIssuerNameToMap(s.Res.IssuerName)})
	} else {
		s.D.Set("issuer_name", nil)
	}

	if s.Res.PublicKeyInfo != nil {
		s.D.Set("public_key_info", []interface{}{CertificatePublicKeyInfoToMap(s.Res.PublicKeyInfo)})
	} else {
		s.D.Set("public_key_info", nil)
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", *s.Res.SerialNumber)
	}

	if s.Res.SignatureAlgorithm != nil {
		s.D.Set("signature_algorithm", *s.Res.SignatureAlgorithm)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubjectName != nil {
		s.D.Set("subject_name", []interface{}{CertificateSubjectNameToMap(s.Res.SubjectName)})
	} else {
		s.D.Set("subject_name", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNotValidAfter != nil {
		s.D.Set("time_not_valid_after", s.Res.TimeNotValidAfter.String())
	}

	if s.Res.TimeNotValidBefore != nil {
		s.D.Set("time_not_valid_before", s.Res.TimeNotValidBefore.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func CertificateExtensionToMap(obj oci_waas.CertificateExtensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsCritical != nil {
		result["is_critical"] = bool(*obj.IsCritical)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func CertificateIssuerNameToMap(obj *oci_waas.CertificateIssuerName) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommonName != nil {
		result["common_name"] = string(*obj.CommonName)
	}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.EmailAddress != nil {
		result["email_address"] = string(*obj.EmailAddress)
	}

	if obj.Locality != nil {
		result["locality"] = string(*obj.Locality)
	}

	if obj.Organization != nil {
		result["organization"] = string(*obj.Organization)
	}

	if obj.OrganizationalUnit != nil {
		result["organizational_unit"] = string(*obj.OrganizationalUnit)
	}

	if obj.StateProvince != nil {
		result["state_province"] = string(*obj.StateProvince)
	}

	return result
}

func CertificatePublicKeyInfoToMap(obj *oci_waas.CertificatePublicKeyInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Algorithm != nil {
		result["algorithm"] = string(*obj.Algorithm)
	}

	if obj.Exponent != nil {
		result["exponent"] = int(*obj.Exponent)
	}

	if obj.KeySize != nil {
		result["key_size"] = int(*obj.KeySize)
	}

	return result
}

func CertificateSubjectNameToMap(obj *oci_waas.CertificateSubjectName) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommonName != nil {
		result["common_name"] = string(*obj.CommonName)
	}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.EmailAddress != nil {
		result["email_address"] = string(*obj.EmailAddress)
	}

	if obj.Locality != nil {
		result["locality"] = string(*obj.Locality)
	}

	if obj.Organization != nil {
		result["organization"] = string(*obj.Organization)
	}

	if obj.OrganizationalUnit != nil {
		result["organizational_unit"] = string(*obj.OrganizationalUnit)
	}

	if obj.StateProvince != nil {
		result["state_province"] = string(*obj.StateProvince)
	}

	return result
}

func (s *WaasCertificateResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waas.ChangeCertificateCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CertificateId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.ChangeCertificateCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
