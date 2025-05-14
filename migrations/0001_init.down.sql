DROP INDEX "public"."IDX_Lead_client_id_datetime";

ALTER TABLE "leads" DROP CONSTRAINT "FK_Lead_Client";

DROP TABLE "leads";

ALTER TABLE "clients" DROP CONSTRAINT "UQ_Client_email";

ALTER TABLE "clients" DROP CONSTRAINT "UQ_Client_phone";

DROP TABLE "clients";