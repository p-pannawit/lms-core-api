-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "uuid" uuid NOT NULL,
  "name" text NULL,
  "amount" bigint NULL,
  "metadata" jsonb NULL,
  PRIMARY KEY ("uuid")
);
