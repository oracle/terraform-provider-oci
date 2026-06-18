// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package events

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_events "github.com/oracle/oci-go-sdk/v65/events"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestEventsRuleResourceSchema(t *testing.T) {
	if err := EventsRuleResource().InternalValidate(nil, true); err != nil {
		t.Fatalf("resource schema is invalid: %v", err)
	}
}

func TestBuildConditionFromDetails(t *testing.T) {
	condition, err := buildConditionFromDetails(map[string]interface{}{
		"event_types": []interface{}{
			"com.oraclecloud.objectstorage.createbucket",
			"com.oraclecloud.objectstorage.deletebucket",
		},
		"data": `{"resourceId":["test-resource-id"],"freeformTags":{"env":"dev"}}`,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal([]byte(condition), &got); err != nil {
		t.Fatalf("condition is not valid JSON: %v", err)
	}

	want := map[string]interface{}{
		"eventType": []interface{}{
			"com.oraclecloud.objectstorage.createbucket",
			"com.oraclecloud.objectstorage.deletebucket",
		},
		"data": map[string]interface{}{
			"resourceId": []interface{}{"test-resource-id"},
			"freeformTags": map[string]interface{}{
				"env": "dev",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected condition:\ngot:  %#v\nwant: %#v", got, want)
	}
}

func TestMapToActionDetailsListPrefersActionAlias(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, EventsRuleResource().Schema, map[string]interface{}{
		"actions": []interface{}{
			map[string]interface{}{
				"actions": []interface{}{
					map[string]interface{}{
						"action_type": "ONS",
						"is_enabled":  true,
						"topic_id":    "test-topic-id",
					},
				},
				"action": []interface{}{
					map[string]interface{}{
						"action_type": "OSS",
						"is_enabled":  true,
						"stream_id":   "test-stream-id",
					},
				},
			},
		},
	})

	sync := &EventsRuleResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
	}

	actions, err := sync.mapToActionDetailsList("actions.0.%s")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(actions.Actions) != 1 {
		t.Fatalf("expected one action, got %d", len(actions.Actions))
	}
	streamingAction, ok := actions.Actions[0].(oci_events.CreateStreamingServiceActionDetails)
	if !ok {
		t.Fatalf("expected action alias to be used, got %T", actions.Actions[0])
	}
	if streamingAction.StreamId == nil || *streamingAction.StreamId != "test-stream-id" {
		t.Fatalf("unexpected stream id: %#v", streamingAction.StreamId)
	}
}

func TestSelectRuleActionField(t *testing.T) {
	tests := []struct {
		name              string
		oldActionsLen     int
		actionLen         int
		oldActionsChanged bool
		actionChanged     bool
		want              string
	}{
		{
			name:          "uses action alias when only alias is present",
			actionLen:     1,
			actionChanged: true,
			want:          "action",
		},
		{
			name:              "uses deprecated actions when only deprecated field is present",
			oldActionsLen:     1,
			oldActionsChanged: true,
			want:              "actions",
		},
		{
			name:          "migration to action alias ignores stale deprecated actions state",
			oldActionsLen: 1,
			actionLen:     1,
			actionChanged: true,
			want:          "action",
		},
		{
			name:              "migration to deprecated actions ignores stale action alias state",
			oldActionsLen:     1,
			actionLen:         1,
			oldActionsChanged: true,
			want:              "actions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectRuleActionField(tt.oldActionsLen, tt.actionLen, tt.oldActionsChanged, tt.actionChanged)
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
