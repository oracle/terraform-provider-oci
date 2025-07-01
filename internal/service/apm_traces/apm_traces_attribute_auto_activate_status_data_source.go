// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesAttributeAutoActivateStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmTracesAttributeAutoActivateStatus,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_key_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"data_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularApmTracesAttributeAutoActivateStatus(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesAttributeAutoActivateStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AttributesClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesAttributeAutoActivateStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.AttributesClient
	Res    *oci_apm_traces.GetStatusAutoActivateResponse
}

func (s *ApmTracesAttributeAutoActivateStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesAttributeAutoActivateStatusDataSourceCrud) Get() error {
	request := oci_apm_traces.GetStatusAutoActivateRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if dataKeyType, ok := s.D.GetOkExists("data_key_type"); ok {
		request.DataKeyType = oci_apm_traces.GetStatusAutoActivateDataKeyTypeEnum(dataKeyType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetStatusAutoActivate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesAttributeAutoActivateStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesAttributeAutoActivateStatusDataSource-", ApmTracesAttributeAutoActivateStatusDataSource(), s.D))

	s.D.Set("data_key", s.Res.DataKey)

	s.D.Set("state", s.Res.State)

	return nil
}
