CREATE TABLE IF NOT EXISTS "files"(
    id serial PRIMARY KEY,
    path varchar(255) not null,
    originalName varchar(255) not null,
	createdAt      timestamp default now() not null,
    createdBy      int default 0 not null
);