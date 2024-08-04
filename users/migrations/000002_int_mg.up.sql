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
    reserved boolean default false
);

CREATE TABLE IF NOT EXISTS "user_roles"(
    id serial PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    role_id INT REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "capabilities"(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    label varchar(255) not null,
    reserved boolean default false
);

CREATE TABLE IF NOT EXISTS "role_capabilities"(
    id serial PRIMARY KEY,
    role_id INT REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE,
    capability_id INT REFERENCES capabilities (id) ON DELETE CASCADE ON UPDATE CASCADE
);




-- capabilities
INSERT INTO users (firstName,lastName, email, password) VALUES 
                                            ('feras','mnsha','feras@gmail.com','123456'),
                                            ('abc','abc','abc@gmail.com','123456'),
                                            ('xyz','xyz','xyz@gmail.com','123456')
                                            ;

INSERT INTO roles (name,label,reserved) VALUES 
                                            ('admin','admin',true),
                                            ('editor','editor',true),
                                            ('subscriber','subscriber',true)
                                            ;

INSERT INTO user_roles (user_id,role_id) VALUES 
                                            (2,1),
                                            (2,2)
                                            ;

DELETE FROM roles
		WHERE id = 3;                                        
DELETE FROM users
		WHERE id = 1;  

        INSERT INTO users (firstName,lastName, email, password) VALUES 
                                            ('feras','mnsha','feras@gmail.com','123456')
                                            ;

INSERT INTO roles (name,label,reserved) VALUES
                                            ('subscriber','subscriber',true)
                                            ;

SELECT users.email,roles.name FROM userrole JOIN users ON users.id = userrole.user_id JOIN roles on roles.id = userrole.role_id;
SELECT user_id, string_agg(roles.name, ',') FROM userrole JOIN users ON users.id = userrole.user_id JOIN roles on roles.id = userrole.role_id WHERE user_id=4 GROUP BY user_id ;
SELECT capabilities.name,capabilities.label FROM role_capabilities JOIN capabilities ON capabilities.id = role_capabilities.capability_id WHERE role_id = ANY(SELECT role_id FROM user_roles WHERE user_id = 1)  ;