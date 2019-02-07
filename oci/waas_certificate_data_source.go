// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

func WaasCertificateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWaasCertificate,
		Schema: map[string]*schema.Schema{
			"certificate_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"issued_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_name": {
				Type:     schema.TypeList,
				Computed: true,
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

func readSingularWaasCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCertificateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient

	return ReadResource(sync)
}

type WaasCertificateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetCertificateResponse
}

func (s *WaasCertificateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasCertificateDataSourceCrud) Get() error {
	request := oci_waas.GetCertificateRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "waas")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasCertificateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.IssuedBy != nil {
		s.D.Set("issued_by", *s.Res.IssuedBy)
	}

	if s.Res.IssuerName != nil {
		s.D.Set("issuer_name", []interface{}{CertificateSubjectNameToMap(s.Res.IssuerName)})
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
