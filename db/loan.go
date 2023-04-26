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

func FindBookingLoan(traceId string, bookingId string) (dbmodels.BookingLoan, error) {
	ctrlName := "DB - Find Booking Loan"
	trx := DB
	var resBookingLoan dbmodels.BookingLoan
	errInsert := trx.Where(`"bookingId" = ?`, bookingId).Find(&resBookingLoan).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resBookingLoan, errInsert
	}
	return resBookingLoan, nil
}

func UpdateBookingLoan(traceId string, req dbmodels.BookingLoan) (bool, error) {
	ctrlName := "DB - Update Booking Loan"
	trx := DB

	errInsert := trx.Where(`"bookingId" = ?`, req.BookingId).Updates(&req).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Updates", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Updates")
		return false, errInsert
	}
	return true, nil
}

func GetListBookingLoan(traceId string) ([]dbmodels.BookingLoan, error) {
	ctrlName := "DB - Get List Booking Loan"
	trx := DB
	var resBookingLoan []dbmodels.BookingLoan
	errInsert := trx.Find(&resBookingLoan).Error
	if errInsert != nil {
		log.LogData(traceId, ctrlName, "DatabaseCon.Find", constant.LEVEL_LOG_ERROR, errInsert.Error())
		log.LogPrintErrorInsertDB(nil, "DatabaseCon.Find")
		return resBookingLoan, errInsert
	}
	return resBookingLoan, nil
}
