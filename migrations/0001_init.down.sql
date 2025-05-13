ALTER TABLE "leads" DROP CONSTRAINT "FK_Lead_Client";

DROP INDEX "public"."IDX_Lead_client_id";

DROP TABLE "leads";

DROP INDEX "public"."IDX_Client_email";

DROP INDEX "public"."IDX_Client_phone";

DROP TABLE "clients";
