// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRecipeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["remediation_recipe_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AdmRemediationRecipeResource(), fieldMap, readSingularAdmRemediationRecipe)
}

func readSingularAdmRemediationRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRecipeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRecipeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.GetRemediationRecipeResponse
}

func (s *AdmRemediationRecipeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRecipeDataSourceCrud) Get() error {
	request := oci_adm.GetRemediationRecipeRequest{}

	if remediationRecipeId, ok := s.D.GetOkExists("remediation_recipe_id"); ok {
		tmp := remediationRecipeId.(string)
		request.RemediationRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.GetRemediationRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AdmRemediationRecipeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DetectConfiguration != nil {
		s.D.Set("detect_configuration", []interface{}{DetectConfigurationToMap(s.Res.DetectConfiguration)})
	} else {
		s.D.Set("detect_configuration", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRunTriggeredOnKbChange != nil {
		s.D.Set("is_run_triggered_on_kb_change", *s.Res.IsRunTriggeredOnKbChange)
	}

	if s.Res.KnowledgeBaseId != nil {
		s.D.Set("knowledge_base_id", *s.Res.KnowledgeBaseId)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{NetworkConfigurationToMap(s.Res.NetworkConfiguration, true)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	if s.Res.ScmConfiguration != nil {
		scmConfigurationArray := []interface{}{}
		if scmConfigurationMap := ScmConfigurationToMap(&s.Res.ScmConfiguration); scmConfigurationMap != nil {
			scmConfigurationArray = append(scmConfigurationArray, scmConfigurationMap)
		}
		s.D.Set("scm_configuration", scmConfigurationArray)
	} else {
		s.D.Set("scm_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VerifyConfiguration != nil {
		verifyConfigurationArray := []interface{}{}
		if verifyConfigurationMap := VerifyConfigurationToMap(&s.Res.VerifyConfiguration); verifyConfigurationMap != nil {
			verifyConfigurationArray = append(verifyConfigurationArray, verifyConfigurationMap)
		}
		s.D.Set("verify_configuration", verifyConfigurationArray)
	} else {
		s.D.Set("verify_configuration", nil)
	}

	return nil
}
