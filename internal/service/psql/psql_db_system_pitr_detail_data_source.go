// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlDbSystemPitrDetailDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularPsqlDbSystemPitrDetailWithContext,
		Schema: map[string]*schema.Schema{
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"pitr_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recovery_time_windows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_recovery_window_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_recovery_window_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularPsqlDbSystemPitrDetailWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsqlDbSystemPitrDetailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsqlDbSystemPitrDetailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetPitrDetailsResponse
}

func (s *PsqlDbSystemPitrDetailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDbSystemPitrDetailDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psql.GetPitrDetailsRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetPitrDetails(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlDbSystemPitrDetailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlDbSystemPitrDetailDataSource-", PsqlDbSystemPitrDetailDataSource(), s.D))

	s.D.Set("pitr_state", s.Res.PitrState)

	recoveryTimeWindows := []interface{}{}
	for _, item := range s.Res.RecoveryTimeWindows {
		recoveryTimeWindows = append(recoveryTimeWindows, PitrTimeWindowToMap(item))
	}
	s.D.Set("recovery_time_windows", recoveryTimeWindows)

	return nil
}

func PitrTimeWindowToMap(obj oci_psql.PitrTimeWindow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeRecoveryWindowEnd != nil {
		result["time_recovery_window_end"] = obj.TimeRecoveryWindowEnd.String()
	}

	if obj.TimeRecoveryWindowStart != nil {
		result["time_recovery_window_start"] = obj.TimeRecoveryWindowStart.String()
	}

	return result
}
