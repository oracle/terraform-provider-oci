// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func TestDevopsDeploymentsDataSourceSetDataArtifactOverrides(t *testing.T) {
	for _, tc := range []struct {
		name      string
		overrides *oci_devops.DeployArtifactOverrideArgumentCollection
		want      []interface{}
	}{
		{
			name: "populated",
			overrides: &oci_devops.DeployArtifactOverrideArgumentCollection{
				Items: []oci_devops.DeployArtifactOverrideArgument{
					{DeployArtifactId: oci_common.String("artifact-1"), Name: oci_common.String("version"), Value: oci_common.String("v1")},
					{DeployArtifactId: oci_common.String("artifact-2"), Name: oci_common.String("image"), Value: oci_common.String("v2")},
				},
			},
			want: []interface{}{map[string]interface{}{
				"items": []interface{}{
					map[string]interface{}{"deploy_artifact_id": "artifact-1", "name": "version", "value": "v1"},
					map[string]interface{}{"deploy_artifact_id": "artifact-2", "name": "image", "value": "v2"},
				},
			}},
		},
		{
			name:      "empty",
			overrides: &oci_devops.DeployArtifactOverrideArgumentCollection{},
			want:      []interface{}{map[string]interface{}{"items": []interface{}{}}},
		},
		{
			name: "absent",
			want: []interface{}{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			summaries := map[string]oci_devops.DeploymentSummary{
				"PIPELINE_DEPLOYMENT": oci_devops.DeployPipelineDeploymentSummary{
					DeployArtifactOverrideArguments: tc.overrides,
				},
				"PIPELINE_REDEPLOYMENT": oci_devops.DeployPipelineRedeploymentSummary{
					DeployArtifactOverrideArguments: tc.overrides,
				},
				"SINGLE_STAGE_DEPLOYMENT": oci_devops.SingleDeployStageDeploymentSummary{
					DeployArtifactOverrideArguments: tc.overrides,
				},
				"SINGLE_STAGE_REDEPLOYMENT": oci_devops.SingleDeployStageRedeploymentSummary{
					DeployArtifactOverrideArguments: tc.overrides,
				},
			}
			for deploymentType, summary := range summaries {
				t.Run(deploymentType, func(t *testing.T) {
					data := schema.TestResourceDataRaw(t, DevopsDeploymentsDataSource().Schema, nil)
					crud := &DevopsDeploymentsDataSourceCrud{
						D: data,
						Res: &oci_devops.ListDeploymentsResponse{
							DeploymentCollection: oci_devops.DeploymentCollection{
								Items: []oci_devops.DeploymentSummary{summary},
							},
						},
					}
					if err := crud.SetData(); err != nil {
						t.Fatalf("SetData() failed: %v", err)
					}
					items := data.Get("deployment_collection.0.items").([]interface{})
					if len(items) != 1 {
						t.Fatalf("got %d deployments, want 1", len(items))
					}
					item := items[0].(map[string]interface{})
					if got := item["deployment_type"]; got != deploymentType {
						t.Errorf("deployment_type = %v, want %s", got, deploymentType)
					}
					if got := item["deploy_artifact_override_arguments"]; !reflect.DeepEqual(got, tc.want) {
						t.Errorf("deploy_artifact_override_arguments = %#v, want %#v", got, tc.want)
					}
				})
			}
		})
	}
}
