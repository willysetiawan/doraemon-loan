package db

import (
	"process-loan/db/dbmodels"
	"process-loan/lib/constant"
	"process-loan/lib/log"
)

func GetActiveInstallment(traceId string) ([]dbmodels.Installment, error) {
	ctrlName := "DB - Get Installment"
	trx := DB
	var resInstallment []dbmodels.Installment
	errInsert := trx.Where(`"installmentActive" = ?`, true).Find(&resInstallment).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resInstallment, errInsert
	}
	return resInstallment, nil
}
