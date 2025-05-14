CREATE TABLE
  "clients" (
    "id" uuid NOT NULL,
    "name" VARCHAR(31) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(15) NOT NULL,
    "website" VARCHAR(127),
    "created_at" TIMESTAMP NOT NULL DEFAULT now (),
    "updated_at" TIMESTAMP NOT NULL DEFAULT now (),
    "deleted_at" TIMESTAMP,
    CONSTRAINT "PK_Client" PRIMARY KEY ("id")
  );

ALTER TABLE "clients"
ADD CONSTRAINT "UQ_Client_email" UNIQUE ("email");

ALTER TABLE "clients"
ADD CONSTRAINT "UQ_Client_phone" UNIQUE ("phone");

CREATE TABLE
  "leads" (
    "id" SERIAL NOT NULL,
    "type" VARCHAR(15) NOT NULL,
    "datetime" TIMESTAMP NOT NULL DEFAULT now (),
    "name" VARCHAR(31),
    "email" VARCHAR(255),
    "phone" VARCHAR(15),
    "client_id" uuid NOT NULL,
    CONSTRAINT "PK_Lead" PRIMARY KEY ("id")
  );

ALTER TABLE "leads"
ADD CONSTRAINT "FK_Lead_Client" FOREIGN KEY ("client_id") REFERENCES "clients" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE INDEX "IDX_Lead_client_id_datetime" ON "leads" ("client_id", "datetime");