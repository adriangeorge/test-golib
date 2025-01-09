package auditmodels

import (
	"fmt"
	"time"
)

type AccountLevelResource string

const (
	Projects      AccountLevelResource = "al_projects"
	Databases     AccountLevelResource = "al_databases"
	Billing       AccountLevelResource = "al_billing"
	Users         AccountLevelResource = "al_users"
	AccessTokens  AccountLevelResource = "al_access_tokens"
	Collaboration AccountLevelResource = "al_collaborations"
)

func (AccountLevelResource) Check(s string) error {
	switch s {
	case "al_projects", "al_databases", "al_billing", "al_users", "al_access_tokens", "al_collaborations":
		return nil
	default:
		return fmt.Errorf("invalid AccountLevelResource: %s", s)
	}
}

type ProjectLevelResource string

const (
	Deployments           ProjectLevelResource = "pl_deployments"
	CodeUpdates           ProjectLevelResource = "pl_code_updates"
	Collaborators         ProjectLevelResource = "pl_collaborators"
	Environments          ProjectLevelResource = "pl_environments"
	DatabaseAssignments   ProjectLevelResource = "pl_databases"
	ClassPauses           ProjectLevelResource = "pl_class_pauses"
	CustomDomains         ProjectLevelResource = "pl_custom_domains"
	Integrations          ProjectLevelResource = "pl_integrations"
	Authentication        ProjectLevelResource = "pl_authentication"
	AuthenticationMethods ProjectLevelResource = "pl_authentication_methods"
	AuthenticationMail    ProjectLevelResource = "pl_authentication_mail"
	EmailService          ProjectLevelResource = "pl_email_service"
	LogDrains             ProjectLevelResource = "pl_log_drains"
)

func (ProjectLevelResource) Check(s string) error {
	switch s {
	case
		"pl_deployments",
		"pl_code_updates",
		"pl_collaborators",
		"pl_environments",
		"pl_databases",
		"pl_class_pauses",
		"pl_custom_domains",
		"pl_integrations",
		"pl_authentication",
		"pl_authentication_methods",
		"pl_authentication_mail",
		"pl_email_service",
		"pl_log_drains":
		return nil
	default:
		return fmt.Errorf("invalid ProjectLevelResource: %s", s)
	}
}

type Action string

const (
	ActionCreate  Action = "create"
	ActionRead    Action = "read"
	ActionUpdate  Action = "update"
	ActionDelete  Action = "delete"
	ActionEnable  Action = "enable"
	ActionDisable Action = "disable"
)

type AuditFilter struct {
	Resource     AccountLevelResource
	Action       Action
	Before       time.Time
	After        time.Time
	AuthorUserID string
}
