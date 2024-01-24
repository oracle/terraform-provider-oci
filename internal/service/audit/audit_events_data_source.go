// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package audit

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_audit "github.com/oracle/oci-go-sdk/v65/audit"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AuditAuditEventsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAuditAuditEvents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"audit_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cloud_events_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"content_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"additional_details": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"event_grouping_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"event_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"identity": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"auth_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"caller_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"caller_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"console_session_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"credentials": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ip_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"principal_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"principal_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"tenant_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_agent": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"request": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"action": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"headers": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"parameters": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"response": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"headers": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"message": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"payload": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"response_time": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"state_change": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"current": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"previous": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
											},
										},
									},
								},
							},
						},
						"event_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_type_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readAuditAuditEvents(d *schema.ResourceData, m interface{}) error {
	sync := &AuditAuditEventsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AuditClient()

	return tfresource.ReadResource(sync)
}

type AuditAuditEventsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_audit.AuditClient
	Res    *oci_audit.ListEventsResponse
}

func (s *AuditAuditEventsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AuditAuditEventsDataSourceCrud) Get() error {
	request := oci_audit.ListEventsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if endTime, ok := s.D.GetOkExists("end_time"); ok {
		tmp, err := time.Parse(time.RFC3339, endTime.(string))
		if err != nil {
			return err
		}
		request.EndTime = &oci_common.SDKTime{Time: tmp}
	}

	if startTime, ok := s.D.GetOkExists("start_time"); ok {
		tmp, err := time.Parse(time.RFC3339, startTime.(string))
		if err != nil {
			return err
		}
		request.StartTime = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "audit")

	response, err := s.Client.ListEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEvents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AuditAuditEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AuditAuditEventsDataSource-", AuditAuditEventsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		auditEvent := map[string]interface{}{}

		if r.CloudEventsVersion != nil {
			auditEvent["cloud_events_version"] = *r.CloudEventsVersion
		}

		if r.ContentType != nil {
			auditEvent["content_type"] = *r.ContentType
		}

		if r.Data != nil {
			auditEvent["data"] = []interface{}{dataToMap(r.Data)}
		} else {
			auditEvent["data"] = nil
		}

		if r.EventId != nil {
			auditEvent["event_id"] = *r.EventId
		}

		if r.EventTime != nil {
			auditEvent["event_time"] = r.EventTime.String()
		}

		if r.EventType != nil {
			auditEvent["event_type"] = *r.EventType
		}

		if r.EventTypeVersion != nil {
			auditEvent["event_type_version"] = *r.EventTypeVersion
		}

		if r.Source != nil {
			auditEvent["source"] = *r.Source
		}

		resources = append(resources, auditEvent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, AuditAuditEventsDataSource().Schema["audit_events"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("audit_events", resources); err != nil {
		return err
	}

	return nil
}

func dataToMap(obj *oci_audit.Data) map[string]interface{} {
	result := map[string]interface{}{}

	for keys := range obj.AdditionalDetails {
		obj.AdditionalDetails[keys] = fmt.Sprintf("%v", obj.AdditionalDetails[keys])
	}

	result["additional_details"] = obj.AdditionalDetails

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.EventGroupingId != nil {
		result["event_grouping_id"] = string(*obj.EventGroupingId)
	}

	if obj.EventName != nil {
		result["event_name"] = string(*obj.EventName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Identity != nil {
		result["identity"] = []interface{}{identityToMap(obj.Identity)}
	}

	if obj.Request != nil {
		result["request"] = []interface{}{requestToMap(obj.Request)}
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.Response != nil {
		result["response"] = []interface{}{responseToMap(obj.Response)}
	}

	if obj.StateChange != nil {
		result["state_change"] = []interface{}{stateChangeToMap(obj.StateChange)}
	}

	return result
}

func identityToMap(obj *oci_audit.Identity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthType != nil {
		result["auth_type"] = string(*obj.AuthType)
	}

	if obj.CallerId != nil {
		result["caller_id"] = string(*obj.CallerId)
	}

	if obj.CallerName != nil {
		result["caller_name"] = string(*obj.CallerName)
	}

	if obj.ConsoleSessionId != nil {
		result["console_session_id"] = string(*obj.ConsoleSessionId)
	}

	if obj.Credentials != nil {
		result["credentials"] = string(*obj.Credentials)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.PrincipalId != nil {
		result["principal_id"] = string(*obj.PrincipalId)
	}

	if obj.PrincipalName != nil {
		result["principal_name"] = string(*obj.PrincipalName)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.UserAgent != nil {
		result["user_agent"] = string(*obj.UserAgent)
	}

	return result
}

func requestToMap(obj *oci_audit.Request) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = string(*obj.Action)
	}

	result["headers"], _ = tfresource.ConvertMapOfStringSlicesToMapOfStrings(obj.Headers)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["parameters"], _ = tfresource.ConvertMapOfStringSlicesToMapOfStrings(obj.Parameters)

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func responseToMap(obj *oci_audit.Response) map[string]interface{} {
	result := map[string]interface{}{}

	result["headers"], _ = tfresource.ConvertMapOfStringSlicesToMapOfStrings(obj.Headers)

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	result["payload"] = obj.Payload

	if obj.ResponseTime != nil {
		result["response_time"] = obj.ResponseTime.String()
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	return result
}

func stateChangeToMap(obj *oci_audit.StateChange) map[string]interface{} {
	result := map[string]interface{}{}

	result["current"] = tfresource.GenericMapToJsonMap(obj.Current)

	result["previous"] = tfresource.GenericMapToJsonMap(obj.Previous)

	return result
}
