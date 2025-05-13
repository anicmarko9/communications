CREATE TABLE
  "clients" (
    "id" uuid NOT NULL,
    "name" CHARACTER VARYING NOT NULL,
    "email" CHARACTER VARYING NOT NULL,
    "phone" CHARACTER VARYING NOT NULL,
    "website" CHARACTER VARYING,
    "verified" BOOLEAN NOT NULL DEFAULT false,
    "token" CHARACTER VARYING,
    "created_at" TIMESTAMP NOT NULL DEFAULT now (),
    "updated_at" TIMESTAMP NOT NULL DEFAULT now (),
    "deleted_at" TIMESTAMP,
    CONSTRAINT "PK_Client" PRIMARY KEY ("id")
  );

CREATE INDEX "IDX_Client_phone" ON "clients" ("phone");

CREATE INDEX "IDX_Client_email" ON "clients" ("email");

CREATE TABLE
  "leads" (
    "id" SERIAL NOT NULL,
    "type" CHARACTER VARYING NOT NULL,
    "datetime" TIMESTAMP NOT NULL DEFAULT now (),
    "patient_name" CHARACTER VARYING,
    "patient_email" CHARACTER VARYING,
    "patient_phone" CHARACTER VARYING,
    "client_id" uuid NOT NULL,
    CONSTRAINT "PK_Lead" PRIMARY KEY ("id")
  );

CREATE INDEX "IDX_Lead_client_id" ON "leads" ("client_id");

ALTER TABLE "leads"
ADD CONSTRAINT "FK_Lead_Client" FOREIGN KEY ("client_id") REFERENCES "clients" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;