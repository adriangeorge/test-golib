package main

import (
	"fmt"

	"github.com/adriangeorge/test-golib/audit"
	auditmodels "github.com/adriangeorge/test-golib/audit_models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/genezio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	audit.NewAuditServiceWithDB(db)

	pa := audit.NewProjectAudit("9b59dd97-6c04-40f4-bdb2-1bba31da4e31", "32daa7b4-0b3b-4b3b-8b3b-0b3b4b3b4b3b", nil)
	logId, err := pa.SubmitAuditLog(auditmodels.ActionCreate)
	if err != nil {
		panic(err)
	}

	logDetailId, err := pa.SubmitAuditDetail(logId, "test message")
	if err != nil {
		panic(err)
	}

	fmt.Println(logId, logDetailId)
}
