package audit

import (
	auditmodels "github.com/adriangeorge/test-golib/audit_models"
)

type ResourceAudit interface {
	SubmitAuditLog(action auditmodels.Action) (string, error)
	SubmitAuditDetail(parentId string, message string) (string, error)
}
