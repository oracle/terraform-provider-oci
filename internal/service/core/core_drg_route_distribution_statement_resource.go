// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreDrgRouteDistributionStatementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgRouteDistributionStatement,
		Read:     readCoreDrgRouteDistributionStatement,
		Update:   updateCoreDrgRouteDistributionStatement,
		Delete:   deleteCoreDrgRouteDistributionStatement,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_route_distribution_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"match_criteria": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"match_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DRG_ATTACHMENT_ID",
								"DRG_ATTACHMENT_TYPE",
								"",
							}, true),
						},
						// Optional
						"attachment_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drg_attachment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func createCoreDrgRouteDistributionStatement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDrgRouteDistributionStatement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.ReadResource(sync)
}

func updateCoreDrgRouteDistributionStatement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrgRouteDistributionStatement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.DeleteResource(d, sync)
}

type CoreDrgRouteDistributionStatementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgRouteDistributionStatement
	DisableNotFoundRetries bool
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) ID() string {
	return getDrgRouteDistributionStatementCompositeId(s.D.Get("drg_route_distribution_id").(string), *s.Res.Id)
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) Create() error {
	request := oci_core.AddDrgRouteDistributionStatementsRequest{}

	if drgRouteDistributionId, ok := s.D.GetOkExists("drg_route_distribution_id"); ok {
		tmp := drgRouteDistributionId.(string)
		request.DrgRouteDistributionId = &tmp
	}

	statement := oci_core.AddDrgRouteDistributionStatementDetails{}

	if action, ok := s.D.GetOkExists("action"); ok {
		statement.Action = oci_core.AddDrgRouteDistributionStatementDetailsActionEnum(action.(string))
	}

	if matchCriteria, ok := s.D.GetOkExists("match_criteria"); ok {
		if tmpList := matchCriteria.([]interface{}); len(tmpList) > 0 {

			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "match_criteria", 0)

			converted, err := s.mapToDrgRouteDistributionMatchCriteria(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert match criteria, encountered error: %v", err)
			}

			statement.MatchCriteria = []oci_core.DrgRouteDistributionMatchCriteria{converted}

		}
	}

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		statement.Priority = &tmp
	}

	tmp := []oci_core.AddDrgRouteDistributionStatementDetails{statement}
	request.Statements = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AddDrgRouteDistributionStatements(context.Background(), request)
	if err != nil {
		return err
	}

	if len(response.Items) > 0 {
		for _, responseStatement := range response.Items {
			if responseStatement.Action != oci_core.DrgRouteDistributionStatementActionEnum(statement.Action) {
				continue
			}
			if len(responseStatement.MatchCriteria) == len(statement.MatchCriteria) {
				if !isDrgRouteDistributionMatchCriteriaEqual(responseStatement.MatchCriteria, statement.MatchCriteria) {
					continue
				}
			}
			if *responseStatement.Priority != *statement.Priority {
				continue
			}
			s.Res = &responseStatement
			break
		}
	} else {
		return fmt.Errorf("distribution statement missing in response")
	}
	return nil
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) Get() error {
	request := oci_core.ListDrgRouteDistributionStatementsRequest{}

	drgRouteDistributionId, statementId, err := parseDrgRouteDistributionStatementCompositeId(s.D.Id())
	if err == nil {
		request.DrgRouteDistributionId = &drgRouteDistributionId
		s.D.Set("drg_route_distribution_id", &drgRouteDistributionId)
	} else {
		log.Printf("[WARN] !!! Get() unable to parse current ID: %s", s.D.Id())
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgRouteDistributionStatements(context.Background(), request)
	if err != nil {
		return err
	}

	var rules []oci_core.DrgRouteDistributionStatement
	rules = response.Items
	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteDistributionStatements(context.Background(), request)
		if err != nil {
			return err
		}

		rules = append(rules, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, r := range rules {
		if *r.Id == statementId {
			s.Res = &r
			break
		}
	}

	if s.Res == nil {
		return fmt.Errorf("drg route distribution statements not found in the list response")
	}

	return nil
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) Update() error {
	request := oci_core.UpdateDrgRouteDistributionStatementsRequest{}

	drgRouteDistributionId, statementId, err := parseDrgRouteDistributionStatementCompositeId(s.D.Id())
	if err == nil {
		request.DrgRouteDistributionId = &drgRouteDistributionId
		s.D.Set("drg_route_distribution_id", &drgRouteDistributionId)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	statement := oci_core.UpdateDrgRouteDistributionStatementDetails{}
	statement.Id = &statementId

	if matchCriteria, ok := s.D.GetOkExists("match_criteria"); ok && s.D.HasChange("match_criteria") {
		if tmpList := matchCriteria.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "match_criteria", 0)

			converted, err := s.mapToDrgRouteDistributionMatchCriteria(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert icmp_options, encountered error: %v", err)
			}
			statement.MatchCriteria = []oci_core.DrgRouteDistributionMatchCriteria{&converted}

		}
	}

	if priority, ok := s.D.GetOkExists("priority"); ok && s.D.HasChange("priority") {
		tmp := priority.(int)
		statement.Priority = &tmp
	}

	tmp := []oci_core.UpdateDrgRouteDistributionStatementDetails{statement}
	request.Statements = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")
	response, err := s.Client.UpdateDrgRouteDistributionStatements(context.Background(), request)
	if err != nil {
		return fmt.Errorf("failed to Update distribution statements, error: %v", err)
	}
	if response.Items != nil && len(response.Items) > 0 {
		_, statementId, err := parseDrgRouteDistributionStatementCompositeId(s.D.Id())
		for _, distributionStatement := range response.Items {
			if *distributionStatement.Id == statementId {
				s.Res = &distributionStatement
			}
		}
		if err != nil {
			return fmt.Errorf("failed to Update distribution statements, error: %v", err)
		}
	}

	return nil
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) Delete() error {

	request := oci_core.RemoveDrgRouteDistributionStatementsRequest{}

	drgRouteDistributionId, statementId, err := parseDrgRouteDistributionStatementCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())

	}
	request.DrgRouteDistributionId = &drgRouteDistributionId
	tmp := []string{statementId}
	request.StatementIds = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")
	_, err = s.Client.RemoveDrgRouteDistributionStatements(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.Action != "" {
		s.D.Set("action", s.Res.Action)
	}

	if s.Res.MatchCriteria != nil {
		matchCriteria := []interface{}{}
		for _, item := range s.Res.MatchCriteria {
			matchCriteria = append(matchCriteria, DrgRouteDistributionMatchCriteriaToMap(item))
		}
		s.D.Set("match_criteria", matchCriteria)

	}
	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	drgRouteDistributionId, statementId, err := parseDrgRouteDistributionStatementCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("drg_route_distribution_id", &drgRouteDistributionId)
		s.D.SetId(getDrgRouteDistributionStatementCompositeId(drgRouteDistributionId, statementId))
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	return nil
}

func getDrgRouteDistributionStatementCompositeId(drgRouteDistributionId string, statementId string) string {
	drgRouteDistributionId = url.PathEscape(drgRouteDistributionId)
	statementId = url.PathEscape(statementId)
	compositeId := "drgRouteDistributions/" + drgRouteDistributionId + "/statements/" + statementId
	return compositeId
}

func parseDrgRouteDistributionStatementCompositeId(compositeId string) (drgRouteDistributionId string, statementsId string, err error) {

	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("drgRouteDistributions/.*/statements/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	drgRouteDistributionId, _ = url.PathUnescape(parts[1])
	statementsId, _ = url.PathUnescape(parts[3])
	return
}

func (s *CoreDrgRouteDistributionStatementResourceCrud) mapToDrgRouteDistributionMatchCriteria(fieldKeyFormat string) (oci_core.DrgRouteDistributionMatchCriteria, error) {

	var baseObject oci_core.DrgRouteDistributionMatchCriteria
	//discriminator
	matchTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "match_type"))
	var matchType string
	if ok {
		matchType = matchTypeRaw.(string)
	} else {
		matchType = "" // default value
		return baseObject, nil
	}
	switch strings.ToLower(matchType) {
	case strings.ToLower("DRG_ATTACHMENT_ID"):
		details := oci_core.DrgAttachmentIdDrgRouteDistributionMatchCriteria{}
		if drgAttachmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "drg_attachment_id")); ok {
			tmp := drgAttachmentId.(string)
			details.DrgAttachmentId = &tmp
		}
		baseObject = details
	case strings.ToLower("DRG_ATTACHMENT_TYPE"):
		details := oci_core.DrgAttachmentTypeDrgRouteDistributionMatchCriteria{}
		if attachmentType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attachment_type")); ok {
			details.AttachmentType = oci_core.DrgAttachmentTypeDrgRouteDistributionMatchCriteriaAttachmentTypeEnum(attachmentType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown match_type '%v' was specified", matchType)
	}
	return baseObject, nil
}

func DrgRouteDistributionMatchCriteriaToMap(obj oci_core.DrgRouteDistributionMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_core.DrgAttachmentIdDrgRouteDistributionMatchCriteria:
		result["match_type"] = "DRG_ATTACHMENT_ID"

		if v.DrgAttachmentId != nil {
			result["drg_attachment_id"] = string(*v.DrgAttachmentId)
		}
	case oci_core.DrgAttachmentTypeDrgRouteDistributionMatchCriteria:
		result["match_type"] = "DRG_ATTACHMENT_TYPE"

		result["attachment_type"] = string(v.AttachmentType)
	default:
		log.Printf("[WARN] Received 'match_type' of unknown type %v", obj)
		return nil
	}
	return result
}

func isDrgRouteDistributionMatchCriteriaEqual(criteria1, criteria2 oci_core.DrgRouteDistributionMatchCriteria) bool {
	mapCriteria1 := DrgRouteDistributionMatchCriteriaToMap(criteria1)
	mapCriteria2 := DrgRouteDistributionMatchCriteriaToMap(criteria2)

	for key, value := range mapCriteria1 {
		if val2, ok := mapCriteria2[key]; ok {
			if val2 != value {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
