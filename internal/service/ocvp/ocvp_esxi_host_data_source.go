// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v58/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OcvpEsxiHostDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["esxi_host_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OcvpEsxiHostResource(), fieldMap, readSingularOcvpEsxiHost)
}

func readSingularOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EsxiHostClient()

	return tfresource.ReadResource(sync)
}

type OcvpEsxiHostDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.EsxiHostClient
	Res    *oci_ocvp.GetEsxiHostResponse
}

func (s *OcvpEsxiHostDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpEsxiHostDataSourceCrud) Get() error {
	request := oci_ocvp.GetEsxiHostRequest{}

	if esxiHostId, ok := s.D.GetOkExists("esxi_host_id"); ok {
		tmp := esxiHostId.(string)
		request.EsxiHostId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetEsxiHost(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpEsxiHostDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BillingContractEndDate != nil {
		s.D.Set("billing_contract_end_date", s.Res.BillingContractEndDate.String())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	s.D.Set("current_sku", s.Res.CurrentSku)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedEsxiHostId != nil {
		s.D.Set("failed_esxi_host_id", *s.Res.FailedEsxiHostId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracePeriodEndDate != nil {
		s.D.Set("grace_period_end_date", s.Res.GracePeriodEndDate.String())
	}

	s.D.Set("next_sku", s.Res.NextSku)

	if s.Res.ReplacementEsxiHostId != nil {
		s.D.Set("replacement_esxi_host_id", *s.Res.ReplacementEsxiHostId)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
