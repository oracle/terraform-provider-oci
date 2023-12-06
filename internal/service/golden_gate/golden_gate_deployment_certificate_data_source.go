// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentCertificateDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["certificate_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGateDeploymentCertificateResource(), fieldMap, readSingularGoldenGateDeploymentCertificate)
}

func readSingularGoldenGateDeploymentCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentCertificateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentCertificateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetCertificateResponse
}

func (s *GoldenGateDeploymentCertificateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentCertificateDataSourceCrud) Get() error {
	request := oci_golden_gate.GetCertificateRequest{}

	if certificateKey, ok := s.D.GetOkExists("certificate_key"); ok {
		tmp := certificateKey.(string)
		request.CertificateKey = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDeploymentCertificateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentCertificateDataSource-", GoldenGateDeploymentCertificateDataSource(), s.D))

	if s.Res.AuthorityKeyId != nil {
		s.D.Set("authority_key_id", *s.Res.AuthorityKeyId)
	}

	if s.Res.CertificateContent != nil {
		s.D.Set("certificate_content", *s.Res.CertificateContent)
	}

	if s.Res.IsCa != nil {
		s.D.Set("is_ca", *s.Res.IsCa)
	}

	if s.Res.IsSelfSigned != nil {
		s.D.Set("is_self_signed", *s.Res.IsSelfSigned)
	}

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Md5Hash != nil {
		s.D.Set("md5hash", *s.Res.Md5Hash)
	}

	if s.Res.PublicKey != nil {
		s.D.Set("public_key", *s.Res.PublicKey)
	}

	if s.Res.PublicKeyAlgorithm != nil {
		s.D.Set("public_key_algorithm", *s.Res.PublicKeyAlgorithm)
	}

	if s.Res.PublicKeySize != nil {
		s.D.Set("public_key_size", strconv.FormatInt(*s.Res.PublicKeySize, 10))
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	if s.Res.Sha1Hash != nil {
		s.D.Set("sha1hash", *s.Res.Sha1Hash)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Subject != nil {
		s.D.Set("subject", *s.Res.Subject)
	}

	if s.Res.SubjectKeyId != nil {
		s.D.Set("subject_key_id", *s.Res.SubjectKeyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeValidFrom != nil {
		s.D.Set("time_valid_from", s.Res.TimeValidFrom.String())
	}

	if s.Res.TimeValidTo != nil {
		s.D.Set("time_valid_to", s.Res.TimeValidTo.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
