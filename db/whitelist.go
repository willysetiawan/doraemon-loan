package db

import (
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/log"
)

func InsertWhitelist(traceId string, req dbmodels.Whitelist) (bool, error) {
	ctrlName := "DB - Insert Whitelist"
	trx := DB

	errInsert := trx.Create(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Create", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Create")
		return false, errInsert
	}
	return true, nil
}

func InsertBulkWhitelist(traceId string, req []dbmodels.Whitelist) (bool, error) {
	ctrlName := "DB - Insert Bulk Whitelist"
	trx := DB

	errInsert := trx.CreateInBatches(&req, 100).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Create", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Create")
		return false, errInsert
	}
	return true, nil
}

func UpdateWhitelist(traceId string, req dbmodels.Whitelist) (bool, error) {
	ctrlName := "DB - Update Whitelist"
	trx := DB

	errInsert := trx.Updates(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Updates", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Updates")
		return false, errInsert
	}
	return true, nil
}

func GetWhitelist(traceId string) ([]dbmodels.JoinWhitelistPartner, error) {
	ctrlName := "DB - Get Whitelist"
	trx := DB
	var resWhitelist []dbmodels.JoinWhitelistPartner
	errInsert := trx.Raw(`SELECT * FROM m_whitelist INNER JOIN m_partner ON m_partner.id=m_whitelist."partnerId"`).Find(&resWhitelist).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resWhitelist, errInsert
	}
	return resWhitelist, nil
}

func CheckWhitelist(traceId string, req models.ReqCheckWhitelist) (dbmodels.Whitelist, error) {
	ctrlName := "DB - Check Whitelist"
	trx := DB
	var resWhitelist dbmodels.Whitelist
	errInsert := trx.Where(`"employeeIdentityNo" = ? OR "employeeId" = ? OR "employeeMobilePhoneNo" = ?`, req.CustomerIdentityNo, req.EmployeeId, req.PhoneNumber).Find(&resWhitelist).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resWhitelist, errInsert
	}
	return resWhitelist, nil
}
