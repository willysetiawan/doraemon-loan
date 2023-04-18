-- +goose Up
-- +goose StatementBegin
CREATE TABLE `t_debit_payment` (
  `debit_payment_id` int(11) NOT NULL AUTO_INCREMENT,
  `debit_payment_partner_reference_no` varchar(64) NOT NULL,
  `debit_payment_amount` varchar(20) DEFAULT NULL,
  `debit_payment_currency` varchar(3) DEFAULT NULL,
  `debit_payment_url_pay_return` text,
  `debit_payment_type_pay_return` varchar(32) DEFAULT NULL,
  `debit_payment_deeplink_pay_return` varchar(1) DEFAULT NULL,
  `debit_payment_url_pay_notify` text,
  `debit_payment_type_pay_notify` varchar(32) DEFAULT NULL,
  `debit_payment_deeplink_pay_notify` varchar(1) DEFAULT NULL,
  `debit_payment_status` varchar(2) DEFAULT NULL,
  `debit_payment_reference_no` varchar(45) DEFAULT NULL,
  `debit_payment_token` varchar(100) DEFAULT NULL,
  `debit_payment_username` varchar(45) DEFAULT NULL,
  `debit_payment_account_no` varchar(45) DEFAULT NULL,
  `debit_payment_merchant_id` varchar(45) DEFAULT NULL,
  `debit_payment_withdrawal_id` varchar(45) DEFAULT NULL,
  `debit_payment_expire` tinyint(4) DEFAULT NULL COMMENT '0 = transaction not used\\n1 = transaction has been used',
  `debit_payment_created_at` datetime DEFAULT NULL,
  `debit_payment_updated_at` datetime DEFAULT NULL COMMENT 'used for payments from users',
  PRIMARY KEY (`debit_payment_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `t_debit_payment`;
-- +goose StatementEnd
