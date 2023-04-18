package db

import (
	"process-loan/db/dbmodels"
	"process-loan/lib/constant"
	"process-loan/lib/log"
)

func InsertBookingLoan(traceId string, req dbmodels.BookingLoan) (bool, error) {
	ctrlName := "DB - Insert Booking Loan"
	trx := DB

	errInsert := trx.Create(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Create", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Create")
		return false, errInsert
	}
	return true, nil
}
