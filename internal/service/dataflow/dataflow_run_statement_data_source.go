// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowRunStatementDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["statement_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataflowRunStatementResource(), fieldMap, readSingularDataflowRunStatement)
}

func readSingularDataflowRunStatement(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunStatementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowRunStatementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetStatementResponse
}

func (s *DataflowRunStatementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowRunStatementDataSourceCrud) Get() error {
	request := oci_dataflow.GetStatementRequest{}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	}

	if statementId, ok := s.D.GetOkExists("statement_id"); ok {
		compositeId := statementId.(string)
		_, statementId, _ := parseRunStatementCompositeId(compositeId)
		tmp := strconv.FormatInt(statementId, 10)
		request.StatementId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetStatement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowRunStatementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(strconv.FormatInt(*s.Res.Id, 10))

	if s.Res.Code != nil {
		s.D.Set("code", *s.Res.Code)
	}

	if s.Res.Output != nil {
		s.D.Set("output", []interface{}{StatementOutputToMap(s.Res.Output)})
	} else {
		s.D.Set("output", nil)
	}

	if s.Res.Progress != nil {
		s.D.Set("progress", *s.Res.Progress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
