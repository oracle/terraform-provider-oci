// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NosqlTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNosqlTables,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(NosqlTableResource()),
			},
		},
	}
}

func readNosqlTables(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

type NosqlTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_nosql.NosqlClient
	Res    *oci_nosql.ListTablesResponse
}

func (s *NosqlTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NosqlTablesDataSourceCrud) Get() error {
	request := oci_nosql.ListTablesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_nosql.ListTablesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "nosql")

	response, err := s.Client.ListTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NosqlTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NosqlTablesDataSource-", NosqlTablesDataSource(), s.D))

	resources := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		resources = append(resources, TableSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, NosqlTablesDataSource().Schema["table_collection"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("table_collection", resources)

	return nil
}
