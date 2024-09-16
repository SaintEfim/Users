ALTER TABLE Users ADD COLUMN id_uuid UUID DEFAULT uuid_generate_v4();

UPDATE Users SET id_uuid = uuid_generate_v4();

ALTER TABLE Users DROP COLUMN id;

ALTER TABLE Users RENAME COLUMN id_uuid TO id;

ALTER TABLE Users ADD PRIMARY KEY (id);
