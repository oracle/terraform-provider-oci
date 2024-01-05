// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmKnowledgeBaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["knowledge_base_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AdmKnowledgeBaseResource(), fieldMap, readSingularAdmKnowledgeBase)
}

func readSingularAdmKnowledgeBase(d *schema.ResourceData, m interface{}) error {
	sync := &AdmKnowledgeBaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmKnowledgeBaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.GetKnowledgeBaseResponse
}

func (s *AdmKnowledgeBaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmKnowledgeBaseDataSourceCrud) Get() error {
	request := oci_adm.GetKnowledgeBaseRequest{}

	if knowledgeBaseId, ok := s.D.GetOkExists("knowledge_base_id"); ok {
		tmp := knowledgeBaseId.(string)
		request.KnowledgeBaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.GetKnowledgeBase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AdmKnowledgeBaseDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	return nil
}
