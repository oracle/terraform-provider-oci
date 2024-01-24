// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationQueryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["query_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MeteringComputationQueryResource(), fieldMap, readSingularMeteringComputationQuery)
}

func readSingularMeteringComputationQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationQueryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationQueryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetQueryResponse
}

func (s *MeteringComputationQueryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationQueryDataSourceCrud) Get() error {
	request := oci_metering_computation.GetQueryRequest{}

	if queryId, ok := s.D.GetOkExists("query_id"); ok {
		tmp := queryId.(string)
		request.QueryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationQueryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.QueryDefinition != nil {
		s.D.Set("query_definition", []interface{}{QueryDefinitionToMap(s.Res.QueryDefinition)})
	} else {
		s.D.Set("query_definition", nil)
	}

	return nil
}
