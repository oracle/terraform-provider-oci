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
	oci_certificates_management "github.com/oracle/oci-go-sdk/v58/certificatesmanagement"
)

func CertificatesManagementCertificateAuthorityVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCertificatesManagementCertificateAuthorityVersion,
		Schema: map[string]*schema.Schema{
			"certificate_authority_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_authority_version_number": {
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
				MaxItems: 1,
				MinItems: 1,
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
				MaxItems: 1,
				MinItems: 1,
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
		},
	}
}

func readSingularCertificatesManagementCertificateAuthorityVersion(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthorityVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateAuthorityVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetCertificateAuthorityVersionResponse
}

func (s *CertificatesManagementCertificateAuthorityVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateAuthorityVersionDataSourceCrud) Get() error {
	request := oci_certificates_management.GetCertificateAuthorityVersionRequest{}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if versionNumber, ok := s.D.GetOkExists("certificate_authority_version_number"); ok {
		tmp := versionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert versionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.CertificateAuthorityVersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetCertificateAuthorityVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementCertificateAuthorityVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificateAuthorityVersionDataSource-", CertificatesManagementCertificateAuthorityVersionDataSource(), s.D))

	if s.Res.CertificateAuthorityId != nil {
		s.D.Set("certificate_authority_id", s.Res.CertificateAuthorityId)
	}

	if s.Res.IssuerCaVersionNumber != nil {
		s.D.Set("issuer_ca_version_number", strconv.FormatInt(*s.Res.IssuerCaVersionNumber, 10))
	}

	if s.Res.RevocationStatus != nil {
		s.D.Set("revocation_status", []interface{}{RevocationStatusToMap(s.Res.RevocationStatus)})
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", string(*s.Res.SerialNumber))
	}

	s.D.Set("stages", s.Res.Stages)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.Validity != nil {
		s.D.Set("validity", []interface{}{ValidityToMap(s.Res.Validity)})
	}

	if s.Res.VersionName != nil {
		s.D.Set("version_name", *s.Res.VersionName)
	}

	if s.Res.VersionNumber != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	return nil
}

func CertificateAuthorityVersionSummaryToMap(obj *oci_certificates_management.CertificateAuthorityVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateAuthorityId != nil {
		result["certificate_authority_id"] = string(*obj.CertificateAuthorityId)
	}

	if obj.IssuerCaVersionNumber != nil {
		result["issuer_ca_version_number"] = strconv.FormatInt(*obj.IssuerCaVersionNumber, 10)
	}

	if obj.RevocationStatus != nil {
		result["revocation_status"] = []interface{}{RevocationStatusToMap(obj.RevocationStatus)}
	}

	if obj.SerialNumber != nil {
		result["serial_number"] = string(*obj.SerialNumber)
	}

	result["stages"] = obj.Stages

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfDeletion != nil {
		result["time_of_deletion"] = obj.TimeOfDeletion.String()
	}

	if obj.Validity != nil {
		result["validity"] = []interface{}{ValidityToMap(obj.Validity)}
	}

	if obj.VersionName != nil {
		result["version_name"] = string(*obj.VersionName)
	}

	if obj.VersionNumber != nil {
		result["version_number"] = strconv.FormatInt(*obj.VersionNumber, 10)
	}

	return result
}
