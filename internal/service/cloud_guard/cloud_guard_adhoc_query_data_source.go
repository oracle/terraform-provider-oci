// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardAdhocQueryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["adhoc_query_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardAdhocQueryResource(), fieldMap, readSingularCloudGuardAdhocQuery)
}

func readSingularCloudGuardAdhocQuery(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardAdhocQueryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetAdhocQueryResponse
}

func (s *CloudGuardAdhocQueryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardAdhocQueryDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetAdhocQueryRequest{}

	if adhocQueryId, ok := s.D.GetOkExists("adhoc_query_id"); ok {
		tmp := adhocQueryId.(string)
		request.AdhocQueryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetAdhocQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardAdhocQueryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdhocQueryDetails != nil {
		s.D.Set("adhoc_query_details", []interface{}{AdhocQueryDetailsToMap(s.Res.AdhocQueryDetails)})
	} else {
		s.D.Set("adhoc_query_details", nil)
	}

	adhocQueryRegionalDetails := []interface{}{}
	for _, item := range s.Res.AdhocQueryRegionalDetails {
		adhocQueryRegionalDetails = append(adhocQueryRegionalDetails, AdhocQueryRegionalDetailsToMap(item))
	}
	s.D.Set("adhoc_query_regional_details", adhocQueryRegionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.ErrorMessage != nil {
		s.D.Set("error_message", *s.Res.ErrorMessage)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
