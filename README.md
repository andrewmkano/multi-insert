# Multi-Insert Example
An example used to practice multi-inserts with the go sql library.

Make sure to run the following commands in your local postgres server before running this example.
> CREATE ROLE dummy WITH LOGIN PASSWORD 'dummy';

> CREATE DATABASE notebook WITH OWNER dummy;

> CREATE TABLE IF NOT EXISTS contacts(
 id integer PRIMARY KEY,
first_name character varying(70) NOT NULL DEFAULT '',
last_name character varying(70) NOT NULL DEFAULT '',
email character varying(70) NOT NULL DEFAULT '');
