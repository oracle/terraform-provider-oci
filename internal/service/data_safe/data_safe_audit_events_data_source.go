// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeAuditEventsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAuditEvents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"audit_event_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"action_taken": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_event_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_policies": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_trail_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audit_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_program": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"command_param": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"command_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_unique_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"error_code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"event_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"extended_event_attributes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_alerted": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_terminal": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"peer_target_database_key": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"target_class": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_collected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trail_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeAuditEvents(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditEventsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditEventsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAuditEventsResponse
}

func (s *DataSafeAuditEventsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditEventsDataSourceCrud) Get() error {
	request := oci_data_safe.ListAuditEventsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAuditEventsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAuditEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	/*
		request.Page = s.Res.OpcNextPage

		for request.Page != nil {
			listResponse, err := s.Client.ListAuditEvents(context.Background(), request)
			if err != nil {
				return err
			}

			s.Res.Items = append(s.Res.Items, listResponse.Items...)
			request.Page = listResponse.OpcNextPage
		}
	*/
	return nil
}

func (s *DataSafeAuditEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditEventsDataSource-", DataSafeAuditEventsDataSource(), s.D))
	resources := []map[string]interface{}{}
	auditEvent := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AuditEventsSummaryToMap(item))
	}
	auditEvent["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAuditEventsDataSource().Schema["audit_event_collection"].Elem.(*schema.Resource).Schema)
		auditEvent["items"] = items
	}

	resources = append(resources, auditEvent)
	if err := s.D.Set("audit_event_collection", resources); err != nil {
		return err
	}

	return nil
}

func AuditEventsSummaryToMap(obj oci_data_safe.AuditEventSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActionTaken != nil {
		result["action_taken"] = string(*obj.ActionTaken)
	}

	if obj.AuditEventTime != nil {
		result["audit_event_time"] = obj.AuditEventTime.String()
	}

	result["audit_location"] = string(obj.AuditLocation)

	if obj.AuditPolicies != nil {
		result["audit_policies"] = string(*obj.AuditPolicies)
	}

	if obj.AuditTrailId != nil {
		result["audit_trail_id"] = string(*obj.AuditTrailId)
	}

	result["audit_type"] = string(obj.AuditType)

	if obj.ClientHostname != nil {
		result["client_hostname"] = string(*obj.ClientHostname)
	}

	if obj.ClientId != nil {
		result["client_id"] = string(*obj.ClientId)
	}

	if obj.ClientIp != nil {
		result["client_ip"] = string(*obj.ClientIp)
	}

	if obj.ClientProgram != nil {
		result["client_program"] = string(*obj.ClientProgram)
	}

	if obj.CommandParam != nil {
		result["command_param"] = string(*obj.CommandParam)
	}

	if obj.CommandText != nil {
		result["command_text"] = string(*obj.CommandText)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["database_type"] = string(obj.DatabaseType)

	if obj.DatabaseUniqueName != nil {
		result["database_unique_name"] = string(*obj.DatabaseUniqueName)
	}

	if obj.DbUserName != nil {
		result["db_user_name"] = string(*obj.DbUserName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ErrorCode != nil {
		result["error_code"] = string(*obj.ErrorCode)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	if obj.EventName != nil {
		result["event_name"] = string(*obj.EventName)
	}

	if obj.ExtendedEventAttributes != nil {
		result["extended_event_attributes"] = string(*obj.ExtendedEventAttributes)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAlerted != nil {
		result["is_alerted"] = bool(*obj.IsAlerted)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.ObjectOwner != nil {
		result["object_owner"] = string(*obj.ObjectOwner)
	}

	if obj.ObjectType != nil {
		result["object_type"] = string(*obj.ObjectType)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	result["operation_status"] = string(obj.OperationStatus)

	if obj.OsTerminal != nil {
		result["os_terminal"] = string(*obj.OsTerminal)
	}

	if obj.OsUserName != nil {
		result["os_user_name"] = string(*obj.OsUserName)
	}

	if obj.PeerTargetDatabaseKey != nil {
		result["peer_target_database_key"] = int(*obj.PeerTargetDatabaseKey)
	}

	result["target_class"] = string(obj.TargetClass)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	if obj.TimeCollected != nil {
		result["time_collected"] = obj.TimeCollected.String()
	}

	result["trail_source"] = string(obj.TrailSource)

	return result
}
