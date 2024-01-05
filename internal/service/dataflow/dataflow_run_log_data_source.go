// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"
)

func DataflowRunLogDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataflowRunLog,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataflowRunLog(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunLogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowRunLogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetRunLogResponse
}

func (s *DataflowRunLogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowRunLogDataSourceCrud) Get() error {
	request := oci_dataflow.GetRunLogRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetRunLog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowRunLogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowRunLogDataSource-", DataflowRunLogDataSource(), s.D))

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("unable to read 'content' from response. Error: %v", err)
	} else if base64EncodeContent {
		// This use case is for v0.12, where content should be base64 encoded to avoid
		// being normalized before setting in state.
		s.D.Set("content", base64.StdEncoding.EncodeToString(contentArray))
	} else {
		if s.Res.Content != nil {
			s.D.Set("content", string(contentArray))
		}
	}

	if s.Res.ContentType != nil {
		s.D.Set("content_type", *s.Res.ContentType)
	}

	return nil
}
