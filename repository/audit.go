package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditRepository struct {
	DB *gorm.DB
}

type AuditLog struct {
	ID         uuid.UUID `gorm:"type:char(36);column:id;primaryKey;default:(gen_random_uuid())"`
	CreatedAt  time.Time
	Resource   string `gorm:"index"`
	Identifier string `gorm:"index"`
	Action     string `gorm:"index"`
	UserID     string `gorm:"index"`
	// If the resource is account level, this will be the owner's ID.
	// If the owner id is nil, then the owner is the user, otherwise the owner is the organization.
	OwnerID string `gorm:"index"`
	// If the resource is project level, this will be the project's ID.
	ProjectID string `gorm:"index"`
}

type AuditLogDetail struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey;default:(UUID())"`
	CreatedAt  time.Time
	AuditLogID uuid.UUID `gorm:"index"`
	AuditLog   AuditLog  `gorm:"constraint:OnDelete:CASCADE"`
	Message    string
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{DB: db}
}

func (r *AuditRepository) CreateAuditLogProjectLevel(
	resource string,
	identifier string,
	action string,
	userId string,
	projectId string,
) (AuditLog, error) {
	auditLog := AuditLog{
		ID:         uuid.New(),
		Resource:   resource,
		Identifier: identifier,
		Action:     action,
		UserID:     userId,
		ProjectID:  projectId,
	}

	res := r.DB.Create(&auditLog).Commit()
	return auditLog, res.Error
}

func (r *AuditRepository) CreateAuditLogAccountLevel(
	resource string,
	identifier string,
	action string,
	userId string,
	ownerId *string,
) (AuditLog, error) {
	auditLog := AuditLog{
		ID:         uuid.New(),
		Resource:   resource,
		Identifier: identifier,
		Action:     action,
		UserID:     userId,
	}

	if ownerId != nil {
		auditLog.OwnerID = *ownerId
	}

	res := r.DB.Create(&auditLog)

	return auditLog, res.Error
}

func (r *AuditRepository) CreateAuditLogDetail(
	parentId uuid.UUID,
	message string,
) (AuditLogDetail, error) {
	auditLogDetail := AuditLogDetail{
		ID:         uuid.New(),
		AuditLogID: parentId,
		Message:    message,
	}

	err := r.DB.Create(&auditLogDetail).Error
	return auditLogDetail, err
}
