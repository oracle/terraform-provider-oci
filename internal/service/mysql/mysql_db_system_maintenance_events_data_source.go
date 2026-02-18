// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlDbSystemMaintenanceEventsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMysqlDbSystemMaintenanceEventsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"maintenance_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mysql_version_after_maintenance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mysql_version_before_maintenance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_notes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_scope": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mysql_version_after_maintenance": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mysql_version_before_maintenance": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_mysql_switch_over_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_mysql_switch_over_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMysqlDbSystemMaintenanceEventsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlDbSystemMaintenanceEventsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MysqlDbSystemMaintenanceEventsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbSystemClient
	Res    *oci_mysql.ListMaintenanceEventsResponse
}

func (s *MysqlDbSystemMaintenanceEventsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlDbSystemMaintenanceEventsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_mysql.ListMaintenanceEventsRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if maintenanceAction, ok := s.D.GetOkExists("maintenance_action"); ok {
		request.MaintenanceAction = oci_mysql.ListMaintenanceEventsMaintenanceActionEnum(maintenanceAction.(string))
	}

	if maintenanceStatus, ok := s.D.GetOkExists("maintenance_status"); ok {
		request.MaintenanceStatus = oci_mysql.MaintenanceEventMaintenanceStatusEnum(maintenanceStatus.(string))
	}

	if maintenanceType, ok := s.D.GetOkExists("maintenance_type"); ok {
		request.MaintenanceType = oci_mysql.ListMaintenanceEventsMaintenanceTypeEnum(maintenanceType.(string))
	}

	if mysqlVersionAfterMaintenance, ok := s.D.GetOkExists("mysql_version_after_maintenance"); ok {
		tmp := mysqlVersionAfterMaintenance.(string)
		request.MysqlVersionAfterMaintenance = &tmp
	}

	if mysqlVersionBeforeMaintenance, ok := s.D.GetOkExists("mysql_version_before_maintenance"); ok {
		tmp := mysqlVersionBeforeMaintenance.(string)
		request.MysqlVersionBeforeMaintenance = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListMaintenanceEvents(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaintenanceEvents(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlDbSystemMaintenanceEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlDbSystemMaintenanceEventsDataSource-", MysqlDbSystemMaintenanceEventsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemMaintenanceEvent := map[string]interface{}{
			"db_system_id": *r.DbSystemId,
		}

		if r.CompartmentId != nil {
			dbSystemMaintenanceEvent["compartment_id"] = *r.CompartmentId
		}

		dbSystemMaintenanceEvent["maintenance_action"] = r.MaintenanceAction

		if r.MaintenanceNotes != nil {
			dbSystemMaintenanceEvent["maintenance_notes"] = *r.MaintenanceNotes
		}

		dbSystemMaintenanceEvent["maintenance_scope"] = r.MaintenanceScope

		dbSystemMaintenanceEvent["maintenance_status"] = r.MaintenanceStatus

		dbSystemMaintenanceEvent["maintenance_type"] = r.MaintenanceType

		if r.MysqlVersionAfterMaintenance != nil {
			dbSystemMaintenanceEvent["mysql_version_after_maintenance"] = *r.MysqlVersionAfterMaintenance
		}

		if r.MysqlVersionBeforeMaintenance != nil {
			dbSystemMaintenanceEvent["mysql_version_before_maintenance"] = *r.MysqlVersionBeforeMaintenance
		}

		if r.TimeCreated != nil {
			dbSystemMaintenanceEvent["time_created"] = r.TimeCreated.String()
		}

		if r.TimeEnded != nil {
			dbSystemMaintenanceEvent["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeMysqlSwitchOverEnded != nil {
			dbSystemMaintenanceEvent["time_mysql_switch_over_ended"] = r.TimeMysqlSwitchOverEnded.String()
		}

		if r.TimeMysqlSwitchOverStarted != nil {
			dbSystemMaintenanceEvent["time_mysql_switch_over_started"] = r.TimeMysqlSwitchOverStarted.String()
		}

		if r.TimeStarted != nil {
			dbSystemMaintenanceEvent["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, dbSystemMaintenanceEvent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlDbSystemMaintenanceEventsDataSource().Schema["maintenance_events"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("maintenance_events", resources); err != nil {
		return err
	}

	return nil
}
