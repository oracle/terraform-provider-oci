package tfresource

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DiagnosticsToError(diag diag.Diagnostics) error {
	var sb strings.Builder
	for _, d := range diag {
		sb.WriteString(fmt.Sprintf("Severity: %s, Summary: %s, Detail: %s\n", d.Severity(), d.Summary(), d.Detail()))
	}
	return errors.New(sb.String())
}
