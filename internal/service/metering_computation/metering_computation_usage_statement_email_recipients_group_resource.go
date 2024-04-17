// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationUsageStatementEmailRecipientsGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationUsageStatementEmailRecipientsGroup,
		Read:     readMeteringComputationUsageStatementEmailRecipientsGroup,
		Update:   updateMeteringComputationUsageStatementEmailRecipientsGroup,
		Delete:   deleteMeteringComputationUsageStatementEmailRecipientsGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"recipients_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"state": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"first_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			"email_recipients_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMeteringComputationUsageStatementEmailRecipientsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationUsageStatementEmailRecipientsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

func updateMeteringComputationUsageStatementEmailRecipientsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMeteringComputationUsageStatementEmailRecipientsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.EmailRecipientsGroup
	DisableNotFoundRetries bool
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) ID() string {
	return GetUsageStatementEmailRecipientsGroupCompositeId(*s.Res.Id, s.D.Get("subscription_id").(string), s.D.Get("compartment_id").(string))
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_metering_computation.EmailRecipientsGroupLifecycleStateActive),
	}
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) Create() error {
	request := oci_metering_computation.CreateEmailRecipientsGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if recipientsList, ok := s.D.GetOkExists("recipients_list"); ok {
		interfaces := recipientsList.([]interface{})
		tmp := make([]oci_metering_computation.EmailRecipient, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "recipients_list", stateDataIndex)
			converted, err := s.mapToEmailRecipient(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("recipients_list") {
			request.RecipientsList = tmp
		}
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.CreateEmailRecipientsGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailRecipientsGroup
	return nil
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) Get() error {
	request := oci_metering_computation.GetEmailRecipientsGroupRequest{}
	emailRecipientsGroupId, subscriptionId, compartmentId, err := parseUsageStatementEmailRecipientsGroupCompositeId(s.D.Id())
	if err == nil {
		request.EmailRecipientsGroupId = &emailRecipientsGroupId
		request.SubscriptionId = &subscriptionId
		request.CompartmentId = &compartmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.GetEmailRecipientsGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailRecipientsGroup
	return nil
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) Update() error {
	request := oci_metering_computation.UpdateEmailRecipientsGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailRecipientsGroupId, ok := s.D.GetOkExists("email_recipients_group_id"); ok {
		tmp := emailRecipientsGroupId.(string)
		request.EmailRecipientsGroupId = &tmp
	}

	if recipientsList, ok := s.D.GetOkExists("recipients_list"); ok {
		interfaces := recipientsList.([]interface{})
		tmp := make([]oci_metering_computation.EmailRecipient, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "recipients_list", stateDataIndex)
			converted, err := s.mapToEmailRecipient(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("recipients_list") {
			request.RecipientsList = tmp
		}
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.UpdateEmailRecipientsGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmailRecipientsGroup
	return nil
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) Delete() error {
	request := oci_metering_computation.DeleteEmailRecipientsGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailRecipientsGroupId, ok := s.D.GetOkExists("email_recipients_group_id"); ok {
		tmp := emailRecipientsGroupId.(string)
		request.EmailRecipientsGroupId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	_, err := s.Client.DeleteEmailRecipientsGroup(context.Background(), request)
	return err
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) SetData() error {

	emailRecipientsGroupId, subscriptionId, _, err := parseUsageStatementEmailRecipientsGroupCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("email_recipients_group_id", &emailRecipientsGroupId)
		s.D.Set("subscription_id", &subscriptionId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	recipientsList := []interface{}{}
	for _, item := range s.Res.RecipientsList {
		recipientsList = append(recipientsList, EmailRecipientToMap(item))
	}
	s.D.Set("recipients_list", recipientsList)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func GetUsageStatementEmailRecipientsGroupCompositeId(emailRecipientsGroupId string, subscriptionId string, compartmentId string) string {
	emailRecipientsGroupId = url.PathEscape(emailRecipientsGroupId)
	subscriptionId = url.PathEscape(subscriptionId)
	compositeId := "usageStatements/" + subscriptionId + "/emailRecipientsGroups/" + emailRecipientsGroupId + "/compartmentId/" + compartmentId
	return compositeId
}

func parseUsageStatementEmailRecipientsGroupCompositeId(compositeId string) (emailRecipientsGroupId string, subscriptionId string, compartmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("usageStatements/.*/emailRecipientsGroups/.*/compartmentId/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subscriptionId, _ = url.PathUnescape(parts[1])
	emailRecipientsGroupId, _ = url.PathUnescape(parts[3])
	compartmentId, _ = url.PathUnescape(parts[5])
	return
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupResourceCrud) mapToEmailRecipient(fieldKeyFormat string) (oci_metering_computation.EmailRecipient, error) {
	result := oci_metering_computation.EmailRecipient{}

	if emailId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_id")); ok {
		tmp := emailId.(string)
		result.EmailId = &tmp
	}

	if firstName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "first_name")); ok {
		tmp := firstName.(string)
		result.FirstName = &tmp
	}

	if lastName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_name")); ok {
		tmp := lastName.(string)
		result.LastName = &tmp
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		result.LifecycleState = oci_metering_computation.EmailRecipientLifecycleStateEnum(state.(string))
	}

	return result, nil
}

func EmailRecipientToMap(obj oci_metering_computation.EmailRecipient) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailId != nil {
		result["email_id"] = string(*obj.EmailId)
	}

	if obj.FirstName != nil {
		result["first_name"] = string(*obj.FirstName)
	}

	if obj.LastName != nil {
		result["last_name"] = string(*obj.LastName)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func EmailRecipientsGroupSummaryToMap(obj oci_metering_computation.EmailRecipientsGroupSummary, subId string) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}
	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	if obj.Id != nil {
		result["id"] = GetUsageStatementEmailRecipientsGroupCompositeId(string(*obj.Id), subId, string(*obj.CompartmentId))
	}

	recipientsList := []interface{}{}
	for _, item := range obj.RecipientsList {
		recipientsList = append(recipientsList, EmailRecipientToMap(item))
	}
	result["recipients_list"] = recipientsList

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
