-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.m_whitelist
(
    id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    "employeeCif" character varying(30) COLLATE pg_catalog."default",
    "employeeName" character varying(100) COLLATE pg_catalog."default",
    "employeeIdentityNo" character varying(16) COLLATE pg_catalog."default",
    "employeeMobilePhoneNo" character varying(13) COLLATE pg_catalog."default",
    "employeeEmail" character varying COLLATE pg_catalog."default",
    "employeeSalary" numeric,
    "employeeStatus" character varying COLLATE pg_catalog."default",
    "employeeParticipatePayrollMonth" character varying COLLATE pg_catalog."default",
    "employeeMaxLoanAmount" numeric,
    "whitelistCreatedAt" timestamp with time zone,
    "whitelistCreatedBy" character varying COLLATE pg_catalog."default",
    "whitelistUpdatedAt" timestamp with time zone,
    "whitelistUpdatedBy" character varying COLLATE pg_catalog."default",
    "companyId" integer,
    "employeeId" character varying COLLATE pg_catalog."default",
    CONSTRAINT m_whitelist_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd