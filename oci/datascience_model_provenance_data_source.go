// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/datascience"
)

func init() {
	RegisterDatasource("oci_datascience_model_provenance", DatascienceModelProvenanceDataSource())
}

func DatascienceModelProvenanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatascienceModelProvenanceResource(), fieldMap, readSingularDatascienceModelProvenance)
}

func readSingularDatascienceModelProvenance(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelProvenanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

type DatascienceModelProvenanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetModelProvenanceResponse
}

func (s *DatascienceModelProvenanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelProvenanceDataSourceCrud) Get() error {
	request := oci_datascience.GetModelProvenanceRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datascience")

	response, err := s.Client.GetModelProvenance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceModelProvenanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.GitBranch != nil {
		s.D.Set("git_branch", *s.Res.GitBranch)
	}

	if s.Res.GitCommit != nil {
		s.D.Set("git_commit", *s.Res.GitCommit)
	}

	if s.Res.RepositoryUrl != nil {
		s.D.Set("repository_url", *s.Res.RepositoryUrl)
	}

	if s.Res.ScriptDir != nil {
		s.D.Set("script_dir", *s.Res.ScriptDir)
	}

	if s.Res.TrainingScript != nil {
		s.D.Set("training_script", *s.Res.TrainingScript)
	}

	return nil
}
