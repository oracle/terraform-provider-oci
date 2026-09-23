// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package ai_document

import "testing"

func TestUnitUnsupportedLockCallbacksReturnErrors(t *testing.T) {
	tests := map[string]func() error{
		"add model lock":      (&AiDocumentModelResourceCrud{}).AddModelLock,
		"remove model lock":   (&AiDocumentModelResourceCrud{}).RemoveModelLock,
		"add project lock":    (&AiDocumentProjectResourceCrud{}).AddProjectLock,
		"remove project lock": (&AiDocumentProjectResourceCrud{}).RemoveProjectLock,
	}
	for name, callback := range tests {
		t.Run(name, func(t *testing.T) {
			if err := callback(); err == nil {
				t.Fatal("callback returned nil, want a normal error")
			}
		})
	}
}
