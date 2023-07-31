\set username '${1}'
\set password '${2}'

CREATE USER :username WITH ENCRYPTED PASSWORD :'password';

CREATE DATABASE hololens WITH OWNER :username; 


\c hololens :username;

CREATE TABLE objects ( 
	object_id serial PRIMARY KEY,
	object_name TEXT UNIQUE NOT NULL,
	x FLOAT4, 
	y FLOAT4, 
	z FLOAT4
);

