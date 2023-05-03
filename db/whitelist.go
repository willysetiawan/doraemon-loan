package db

import (
	"process-loan/db/dbmodels"
	"process-loan/domain/base/models"
	"process-loan/lib/constant"
	"process-loan/lib/log"
)

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
