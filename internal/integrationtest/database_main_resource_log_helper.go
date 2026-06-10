// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func exaccMainResourceLog(t *testing.T, stepName string, resourceName string, previousID *string, currentID *string, attrs ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if s == nil || s.RootModule() == nil {
			t.Logf("[MAIN_RESOURCE_STATE] action=%s | status=missing-state", stepName)
			return nil
		}

		resourceState, ok := s.RootModule().Resources[resourceName]
		if !ok || resourceState == nil || resourceState.Primary == nil || resourceState.Primary.ID == "" {
			t.Logf("[MAIN_RESOURCE_STATE] action=%s | status=absent", stepName)
			return nil
		}

		resourceID := resourceState.Primary.ID
		details := []string{
			fmt.Sprintf("action=%s", stepName),
		}

		if previousID != nil && *previousID != "" {
			details = append(details,
				fmt.Sprintf("updated_in_place=%t", resourceID == *previousID),
				fmt.Sprintf("recreated=%t", resourceID != *previousID),
			)
		}

		if currentID != nil {
			*currentID = resourceID
		}

		for _, attr := range attrs {
			if value := resourceState.Primary.Attributes[attr]; value != "" {
				details = append(details, fmt.Sprintf("%s=%s", attr, value))
			}
		}

		t.Logf("[MAIN_RESOURCE_STATE] %s", strings.Join(details, " | "))
		return nil
	}
}
