// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/v56/waas"
)

func WaasHttpRedirectDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["http_redirect_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaasHttpRedirectResource(), fieldMap, readSingularWaasHttpRedirect)
}

func readSingularWaasHttpRedirect(d *schema.ResourceData, m interface{}) error {
	sync := &WaasHttpRedirectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RedirectClient()

	return tfresource.ReadResource(sync)
}

type WaasHttpRedirectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.RedirectClient
	Res    *oci_waas.GetHttpRedirectResponse
}

func (s *WaasHttpRedirectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasHttpRedirectDataSourceCrud) Get() error {
	request := oci_waas.GetHttpRedirectRequest{}

	if httpRedirectId, ok := s.D.GetOkExists("http_redirect_id"); ok {
		tmp := httpRedirectId.(string)
		request.HttpRedirectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetHttpRedirect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasHttpRedirectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("response_code", s.Res.ResponseCode)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Target != nil {
		s.D.Set("target", []interface{}{HttpRedirectTargetToMap(s.Res.Target)})
	} else {
		s.D.Set("target", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
