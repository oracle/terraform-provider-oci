// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v36/ocvp"
)

func init() {
	RegisterDatasource("oci_ocvp_esxi_host", OcvpEsxiHostDataSource())
}

func OcvpEsxiHostDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["esxi_host_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OcvpEsxiHostResource(), fieldMap, readSingularOcvpEsxiHost)
}

func readSingularOcvpEsxiHost(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpEsxiHostDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).esxiHostClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ocvp")

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

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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
