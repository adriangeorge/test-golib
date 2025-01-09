package audit

import (
	"errors"

	auditmodels "github.com/adriangeorge/test-golib/audit_models"
	"github.com/google/uuid"
)

type ProjectMetadata struct {
	Name string
}

type ProjectAudit struct {
	resourceType auditmodels.AccountLevelResource
	metadata     ProjectMetadata
	userId       string
	resourceID   string
	auditService *auditService
	ownerId      *string
}

// SubmitAuditDetail implements ResourceAudit.
func (p *ProjectAudit) SubmitAuditDetail(parentId string, message string) (string, error) {
	parsedParentId, err := uuid.Parse(parentId)
	if err != nil {
		return "", err
	}
	logId, err := GetAuditService().PutAuditDetail(parsedParentId, message)
	return logId.String(), err
}

// SubmitAuditLog implements ResourceAudit.
func (p *ProjectAudit) SubmitAuditLog(action auditmodels.Action) (string, error) {
	var logId uuid.UUID

	// Check that the action is valid.
	switch action {
	case auditmodels.ActionCreate:
		fallthrough
	case auditmodels.ActionDelete:
		var err error
		logId, err = p.auditService.PutAuditLogAccountLevel(p.resourceType, p.resourceID, action, p.userId, p.ownerId)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("invalid action")
	}
	return logId.String(), nil
}

func NewProjectAudit(userId, projectId string, ownerId *string) ResourceAudit {
	return &ProjectAudit{
		resourceType: auditmodels.Projects,
		userId:       userId,
		auditService: GetAuditService(),
		resourceID:   projectId,
		ownerId:      ownerId,
	}
}
