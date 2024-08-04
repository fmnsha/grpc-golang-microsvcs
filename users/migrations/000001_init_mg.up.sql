-- CREATE TABLE IF NOT EXISTS "tenants"(
--     id serial PRIMARY KEY,
--     tenant_id varchar unique not null,
--     client varchar unique not null,
--     start_date timestamp default now() not null,
--     end_date timestamp default now() not null
-- );


CREATE TABLE IF NOT EXISTS "users"(
    id serial PRIMARY KEY,
    firstName varchar(255) not null,
    lastName varchar(255) not null,
    email varchar unique not null,
    emailConfirmation boolean default false,
    password varchar(255) not null,
	roles           integer[],
    capabilites     integer[],
	createdAt      timestamp default now() not null,
    createdBy      integer default 0 not null,
    updatedAt      timestamp default now() not null ,
    updatedBy      integer default 0 not null
);

-- ALTER TABLE users ENABLE ROW LEVEL SECURITY;

-- CREATE POLICY users_isolation_policy ON users
-- USING (tenant_id = current_setting('app.current_tenant'));

-- CREATE OR REPLACE FUNCTION set_tenant(clientName text) RETURNS void AS $$
-- BEGIN
--     PERFORM set_config('app.current_tenant', (SELECT tenant_id FROM tenants WHERE client = clientName), false);
-- END;
-- $$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS "roles"(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    label varchar(255) not null,
    reserved boolean default false,
	capabilities integer[],
    createdBy      integer default 0 not null
);

CREATE TABLE IF NOT EXISTS "capabilities"(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    label varchar(255) not null,
    reserved boolean default false,
    createdBy integer default 0 not null
);


CREATE TABLE IF NOT EXISTS "categories" (
    id serial PRIMARY KEY,
    name varchar(255) not null,
    image varchar(255) not null,
    details varchar(255) not null
);
CREATE TABLE IF NOT EXISTS "watches" (
    id serial PRIMARY KEY,
    categoryId integer not null,
    name varchar(255) not null,
    image varchar(255) not null,
    details varchar(255) not null,
    FOREIGN KEY (categoryId) REFERENCES categories (id)
);
CREATE TABLE IF NOT EXISTS "files" (
    id serial PRIMARY KEY,
    path varchar(255) not null,
    originalName varchar(255) not null,
    trash boolean not null
);
CREATE TABLE IF NOT EXISTS "warranties" (
    id SERIAL PRIMARY KEY,
    image VARCHAR(255) NOT NULL,
    type_of_repair VARCHAR(255) NOT NULL,
    serial_num VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    point_of_sale VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    date TIMESTAMP with time zone default now() NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    mobile VARCHAR(255) NOT NULL
);



-- capabilities
INSERT INTO capabilities (name,label) VALUES 
                                            ('getallusers','getallusers'),
                                            ('getuser','getuser'),
                                            ('adduser','adduser'),
                                            ('updateuser','updateuser'),
                                            ('updateotherusers','updateotherusers'),
                                            ('deleteuser','deleteuser'),
                                            ('getallroles','getallroles'),
                                            ('getrole','getrole'),
                                            ('addrole','addrole'),
                                            ('updaterole','updaterole'),
                                            ('deleterole','deleterole'),
                                            ('getallcapabilities','getallcapabilities'),
                                            ('getcapability','getcapability'),
                                            ('addcapability','addcapability'),
                                            ('editcapability','editcapability'),
                                            ('deletecapability','deletecapability'),
                                            ('addcategory','addcategory'),
                                            ('updatecategory','updatecategory'),
                                            ('deletecategory','deletecategory'),
                                            ('addwatch','addwatch'),
                                            ('updatewatch','updatewatch'),
                                            ('deletewatch','deletewatch'),
                                            ('getAllWarrantiesByUserCountry','getAllWarrantiesByUserCountry'),
                                            ('getAllWarranties','getAllWarranties');

