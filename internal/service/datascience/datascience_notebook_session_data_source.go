// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"
)

func DatascienceNotebookSessionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["notebook_session_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceNotebookSessionResource(), fieldMap, readSingularDatascienceNotebookSession)
}

func readSingularDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceNotebookSessionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetNotebookSessionResponse
}

func (s *DatascienceNotebookSessionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceNotebookSessionDataSourceCrud) Get() error {
	request := oci_datascience.GetNotebookSessionRequest{}

	if notebookSessionId, ok := s.D.GetOkExists("notebook_session_id"); ok {
		tmp := notebookSessionId.(string)
		request.NotebookSessionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceNotebookSessionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotebookSessionConfigurationDetails != nil {
		s.D.Set("notebook_session_configuration_details", []interface{}{NotebookSessionConfigurationDetailsToMap(s.Res.NotebookSessionConfigurationDetails)})
	} else {
		s.D.Set("notebook_session_configuration_details", nil)
	}

	if s.Res.NotebookSessionUrl != nil {
		s.D.Set("notebook_session_url", *s.Res.NotebookSessionUrl)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
