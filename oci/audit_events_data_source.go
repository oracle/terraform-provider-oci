// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/*
 * A note to maintainers: request_headers, request_parameters and response_headers should be a _map of keys to lists of
 * strings_ ex: `request_headers[accept-encoding]: ["gzip", "deflate"]`, Terraform fails silently if we try to structure
 * the schema as such. The compromise is to flatten the list of strings into a single string of comma delimited items,
 * ex: `request_headers[accept-encoding]: "gzip, deflate"`.
 */

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_audit "github.com/oracle/oci-go-sdk/audit"
	oci_common "github.com/oracle/oci-go-sdk/common"
)

func AuditAuditEventsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAuditAuditEvents,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"audit_events": {
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
						"credential_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_source": {
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
						"principal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_agent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_headers": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_origin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_parameters": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"request_resource": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"response_headers": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"response_payload": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"response_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"response_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": {
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
	sync.Client = m.(*OracleClients).auditClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "audit")

	response, err := s.Client.ListEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	limit := s.D.Get("limit").(int)
	limit--
	for request.Page != nil && limit > 0 {
		listResponse, err := s.Client.ListEvents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
		limit--
	}

	return nil
}

func (s *AuditAuditEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		auditEvent := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CredentialId != nil {
			auditEvent["credential_id"] = *r.CredentialId
		}

		if r.EventId != nil {
			auditEvent["event_id"] = *r.EventId
		}

		if r.EventName != nil {
			auditEvent["event_name"] = *r.EventName
		}

		if r.EventSource != nil {
			auditEvent["event_source"] = *r.EventSource
		}

		if r.EventTime != nil {
			auditEvent["event_time"] = r.EventTime.String()
		}

		if r.EventType != nil {
			auditEvent["event_type"] = *r.EventType
		}

		if r.PrincipalId != nil {
			auditEvent["principal_id"] = *r.PrincipalId
		}

		if r.RequestAction != nil {
			auditEvent["request_action"] = *r.RequestAction
		}

		if r.RequestAgent != nil {
			auditEvent["request_agent"] = *r.RequestAgent
		}

		auditEvent["request_headers"], _ = convertMapOfStringSlicesToMapOfStrings(r.RequestHeaders)

		if r.RequestId != nil {
			auditEvent["request_id"] = *r.RequestId
		}

		if r.RequestOrigin != nil {
			auditEvent["request_origin"] = *r.RequestOrigin
		}

		auditEvent["request_parameters"], _ = convertMapOfStringSlicesToMapOfStrings(r.RequestParameters)

		if r.RequestResource != nil {
			auditEvent["request_resource"] = *r.RequestResource
		}

		auditEvent["response_headers"], _ = convertMapOfStringSlicesToMapOfStrings(r.ResponseHeaders)

		auditEvent["response_payload"] = objectMapToStringMap(r.ResponsePayload)

		if r.ResponseStatus != nil {
			auditEvent["response_status"] = *r.ResponseStatus
		}

		if r.ResponseTime != nil {
			auditEvent["response_time"] = r.ResponseTime.String()
		}

		if r.TenantId != nil {
			auditEvent["tenant_id"] = *r.TenantId
		}

		resources = append(resources, auditEvent)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AuditAuditEventsDataSource().Schema["audit_events"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("audit_events", resources); err != nil {
		return err
	}

	return nil
}
