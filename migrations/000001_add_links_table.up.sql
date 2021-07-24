CREATE TABLE "metal_links" (
  "id" SERIAL PRIMARY KEY,
  "video_id" varchar,
  "video_title" varchar,
  "url" varchar,
  "timestamp" varchar
);

CREATE INDEX ON "metal_links" ("video_title");
