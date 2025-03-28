// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentDataIngestionJobLogContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularGenerativeAiAgentDataIngestionJobLogContent,
		Schema: map[string]*schema.Schema{
			"data_ingestion_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularGenerativeAiAgentDataIngestionJobLogContent(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobLogContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentDataIngestionJobLogContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetDataIngestionJobLogContentResponse
}

func (s *GenerativeAiAgentDataIngestionJobLogContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentDataIngestionJobLogContentDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.GetDataIngestionJobLogContentRequest{}

	if dataIngestionJobId, ok := s.D.GetOkExists("data_ingestion_job_id"); ok {
		tmp := dataIngestionJobId.(string)
		request.DataIngestionJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetDataIngestionJobLogContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentDataIngestionJobLogContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiAgentDataIngestionJobLogContentDataSource-", GenerativeAiAgentDataIngestionJobLogContentDataSource(), s.D))

	return nil
}
