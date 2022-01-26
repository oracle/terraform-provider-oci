// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
)

func CertificatesManagementCertificateVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCertificatesManagementCertificateVersion,
		Schema: map[string]*schema.Schema{
			"certificate_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_version_number": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"issuer_ca_version_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revocation_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"revocation_reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_revocation": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subject_alternative_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"type": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"validity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_of_validity_not_after": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_validity_not_before": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"version_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCertificatesManagementCertificateVersion(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetCertificateVersionResponse
}

func (s *CertificatesManagementCertificateVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateVersionDataSourceCrud) Get() error {
	request := oci_certificates_management.GetCertificateVersionRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	if certificateVersionNumber, ok := s.D.GetOkExists("certificate_version_number"); ok {
		tmp := certificateVersionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert certificateVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CertificateVersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetCertificateVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementCertificateVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificateVersionDataSource-", CertificatesManagementCertificateVersionDataSource(), s.D))

	if s.Res.IssuerCaVersionNumber != nil {
		s.D.Set("issuer_ca_version_number", strconv.FormatInt(*s.Res.IssuerCaVersionNumber, 10))
	}

	if s.Res.RevocationStatus != nil {
		s.D.Set("revocation_status", []interface{}{RevocationStatusToMap(s.Res.RevocationStatus)})
	} else {
		s.D.Set("revocation_status", nil)
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", *s.Res.SerialNumber)
	}

	s.D.Set("stages", s.Res.Stages)

	subjectAlternativeNames := []interface{}{}
	for _, item := range s.Res.SubjectAlternativeNames {
		subjectAlternativeNames = append(subjectAlternativeNames, CertificateSubjectAlternativeNameToMap(item))
	}
	s.D.Set("subject_alternative_names", subjectAlternativeNames)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.Validity != nil {
		s.D.Set("validity", []interface{}{ValidityToMap(s.Res.Validity)})
	} else {
		s.D.Set("validity", nil)
	}

	if s.Res.VersionName != nil {
		s.D.Set("version_name", *s.Res.VersionName)
	}

	if s.Res.VersionNumber != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	return nil
}
