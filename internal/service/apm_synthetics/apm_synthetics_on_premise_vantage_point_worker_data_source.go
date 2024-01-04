// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsOnPremiseVantagePointWorkerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["on_premise_vantage_point_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["worker_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmSyntheticsOnPremiseVantagePointWorkerResource(), fieldMap, readSingularApmSyntheticsOnPremiseVantagePointWorker)
}

func readSingularApmSyntheticsOnPremiseVantagePointWorker(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsOnPremiseVantagePointWorkerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetWorkerResponse
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetWorkerRequest{}

	if workerCompositeId, ok := s.D.GetOkExists("worker_id"); ok {
		tmp := workerCompositeId.(string)
		onPremiseVantagePointCompositeId, workerId, apmDomainId, err := parseOnPremiseVantagePointWorkerCompositeId(tmp)
		if err == nil {
			request.WorkerId = &workerId
			request.ApmDomainId = &apmDomainId
			onPremiseVantagePointId, _, err2 := parseOnPremiseVantagePointCompositeId(onPremiseVantagePointCompositeId)
			if err2 == nil {
				request.OnPremiseVantagePointId = &onPremiseVantagePointId
			} else {
				log.Printf("[WARN] Get() unable to parse onPremiseVantagePointCompositeId: %s", onPremiseVantagePointCompositeId)
			}
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}

	} else {
		log.Printf("Get() worker_id not found ")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetWorker(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if workerId, ok := s.D.GetOkExists("worker_id"); ok {
		tmp := workerId.(string)
		onPremiseVantagePointCompositeId, workerId1, _, err := parseOnPremiseVantagePointWorkerCompositeId(tmp)
		s.D.Set("worker_id", workerId1)
		s.D.SetId(tmp)
		if err == nil {
			onPremiseVantagePointId, _, err2 := parseOnPremiseVantagePointCompositeId(onPremiseVantagePointCompositeId)
			if err2 == nil {
				s.D.Set("on_premise_vantage_point_id", &onPremiseVantagePointId)
			} else {
				log.Printf("[WARN] Get() unable to parse onPremiseVantagePointCompositeId: %s", onPremiseVantagePointCompositeId)
			}
		} else {
			log.Printf("[WARN] SetData() unable to parse current ID: %s", onPremiseVantagePointCompositeId)
		}
	}

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", configurationDetailsJsonObjectToString(s.Res.ConfigurationDetails))
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GeoInfo != nil {
		s.D.Set("geo_info", *s.Res.GeoInfo)
	}

	if s.Res.IdentityInfo != nil {
		s.D.Set("identity_info", []interface{}{IdentityInfoDetailsToMap(s.Res.IdentityInfo)})
	} else {
		s.D.Set("identity_info", nil)
	}

	monitorList := []interface{}{}
	for _, item := range s.Res.MonitorList {
		monitorList = append(monitorList, WorkerMonitorListToMap(item))
	}
	s.D.Set("monitor_list", monitorList)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.OpvpId != nil {
		s.D.Set("opvp_id", *s.Res.OpvpId)
	}

	if s.Res.OpvpName != nil {
		s.D.Set("opvp_name", *s.Res.OpvpName)
	}

	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	if s.Res.RuntimeId != nil {
		s.D.Set("runtime_id", *s.Res.RuntimeId)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSyncUp != nil {
		s.D.Set("time_last_sync_up", s.Res.TimeLastSyncUp.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VersionDetails != nil {
		s.D.Set("version_details", []interface{}{OnPremiseVpWorkerVersionDetailsToMap(s.Res.VersionDetails)})
	} else {
		s.D.Set("version_details", nil)
	}

	s.D.Set("worker_type", s.Res.WorkerType)

	return nil
}
