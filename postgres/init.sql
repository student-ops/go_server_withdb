CREATE TABLE visits (
  id SERIAL PRIMARY KEY,
  count INT DEFAULT 0
);

INSERT INTO visits (count) VALUES (0);
