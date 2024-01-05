// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlDbSystemPrimaryDbInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularPsqlDbSystemPrimaryDbInstance,
		Schema: map[string]*schema.Schema{
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"db_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularPsqlDbSystemPrimaryDbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemPrimaryDbInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlDbSystemPrimaryDbInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetPrimaryDbInstanceResponse
}

func (s *PsqlDbSystemPrimaryDbInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDbSystemPrimaryDbInstanceDataSourceCrud) Get() error {
	request := oci_psql.GetPrimaryDbInstanceRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetPrimaryDbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlDbSystemPrimaryDbInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlDbSystemPrimaryDbInstanceDataSource-", PsqlDbSystemPrimaryDbInstanceDataSource(), s.D))

	if s.Res.DbInstanceId != nil {
		s.D.Set("db_instance_id", *s.Res.DbInstanceId)
	}

	return nil
}
