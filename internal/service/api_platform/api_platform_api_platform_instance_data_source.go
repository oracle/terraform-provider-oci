// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package api_platform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiPlatformApiPlatformInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["api_platform_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApiPlatformApiPlatformInstanceResource(), fieldMap, readSingularApiPlatformApiPlatformInstance)
}

func readSingularApiPlatformApiPlatformInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()

	return tfresource.ReadResource(sync)
}

type ApiPlatformApiPlatformInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_api_platform.ApiPlatformClient
	Res    *oci_api_platform.GetApiPlatformInstanceResponse
}

func (s *ApiPlatformApiPlatformInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiPlatformApiPlatformInstanceDataSourceCrud) Get() error {
	request := oci_api_platform.GetApiPlatformInstanceRequest{}

	if apiPlatformInstanceId, ok := s.D.GetOkExists("api_platform_instance_id"); ok {
		tmp := apiPlatformInstanceId.(string)
		request.ApiPlatformInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "api_platform")

	response, err := s.Client.GetApiPlatformInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApiPlatformApiPlatformInstanceDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdcsApp != nil {
		s.D.Set("idcs_app", []interface{}{IdcsAppToMap(s.Res.IdcsApp)})
	} else {
		s.D.Set("idcs_app", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Uris != nil {
		s.D.Set("uris", []interface{}{UrisToMap(s.Res.Uris)})
	} else {
		s.D.Set("uris", nil)
	}

	return nil
}
