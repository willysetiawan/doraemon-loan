package db

import (
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/log"
)

func InsertInstallment(traceId string, req dbmodels.Installment) (bool, error) {
	ctrlName := "DB - Insert Installment"
	trx := DB

	errInsert := trx.Create(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Create", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Create")
		return false, errInsert
	}
	return true, nil
}

func UpdateInstallment(traceId string, req dbmodels.Installment) (bool, error) {
	ctrlName := "DB - Update Installment"
	trx := DB

	errInsert := trx.Updates(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Updates", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Updates")
		return false, errInsert
	}
	return true, nil
}

func GetActiveInstallment(traceId string) ([]dbmodels.Installment, error) {
	ctrlName := "DB - Get Installment"
	trx := DB
	var resInstallment []dbmodels.Installment
	errInsert := trx.Where("active = ?", true).Find(&resInstallment).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resInstallment, errInsert
	}
	return resInstallment, nil
}

func FindInstallment(traceId string, installmentId int) (dbmodels.Installment, error) {
	ctrlName := "DB - Find Installment"
	trx := DB
	var resInstallment dbmodels.Installment
	errInsert := trx.Where("id = ?", installmentId).Find(&resInstallment).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resInstallment, errInsert
	}
	return resInstallment, nil
}

func CheckInstallment(traceId string, req models.ReqInstallment) (dbmodels.Installment, error) {
	ctrlName := "DB - Check Installment"
	trx := DB
	var resInstallment dbmodels.Installment
	errInsert := trx.Where("type = ? AND value = ?", req.Type, req.Value).Find(&resInstallment).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resInstallment, errInsert
	}
	return resInstallment, nil
}
