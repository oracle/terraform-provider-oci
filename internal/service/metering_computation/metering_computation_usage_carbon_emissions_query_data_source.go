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

func MeteringComputationUsageCarbonEmissionsQueryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["usage_carbon_emissions_query_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MeteringComputationUsageCarbonEmissionsQueryResource(), fieldMap, readSingularMeteringComputationUsageCarbonEmissionsQuery)
}

func readSingularMeteringComputationUsageCarbonEmissionsQuery(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageCarbonEmissionsQueryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationUsageCarbonEmissionsQueryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetUsageCarbonEmissionsQueryResponse
}

func (s *MeteringComputationUsageCarbonEmissionsQueryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationUsageCarbonEmissionsQueryDataSourceCrud) Get() error {
	request := oci_metering_computation.GetUsageCarbonEmissionsQueryRequest{}

	if usageCarbonEmissionsQueryId, ok := s.D.GetOkExists("usage_carbon_emissions_query_id"); ok {
		tmp := usageCarbonEmissionsQueryId.(string)
		request.UsageCarbonEmissionsQueryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetUsageCarbonEmissionsQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationUsageCarbonEmissionsQueryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.QueryDefinition != nil {
		s.D.Set("query_definition", []interface{}{UsageCarbonEmissionsQueryDefinitionToMap(s.Res.QueryDefinition)})
	} else {
		s.D.Set("query_definition", nil)
	}

	return nil
}
