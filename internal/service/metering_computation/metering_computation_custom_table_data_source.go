// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v58/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationCustomTableDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["custom_table_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MeteringComputationCustomTableResource(), fieldMap, readSingularMeteringComputationCustomTable)
}

func readSingularMeteringComputationCustomTable(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTableDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationCustomTableDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetCustomTableResponse
}

func (s *MeteringComputationCustomTableDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationCustomTableDataSourceCrud) Get() error {
	request := oci_metering_computation.GetCustomTableRequest{}

	if customTableId, ok := s.D.GetOkExists("custom_table_id"); ok {
		tmp := customTableId.(string)
		request.CustomTableId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetCustomTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationCustomTableDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.SavedCustomTable != nil {
		s.D.Set("saved_custom_table", []interface{}{SavedCustomTableToMap(s.Res.SavedCustomTable)})
	} else {
		s.D.Set("saved_custom_table", nil)
	}

	if s.Res.SavedReportId != nil {
		s.D.Set("saved_report_id", *s.Res.SavedReportId)
	}

	return nil
}
