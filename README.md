# Multi-Insert Example
An example used to practice multi-inserts with the go sql library.

Make sure to run the following commands in your local postgres server before running this example.
> CREATE ROLE dummy WITH LOGIN PASSWORD 'dummy';

> CREATE DATABASE notebook WITH OWNER dummy;

> CREATE TABLE IF NOT EXISTS contacts(
first_name character varying(70) NOT NULL DEFAULT '',
last_name character varying(70) NOT NULL DEFAULT '',
email character varying(70) NOT NULL DEFAULT '');

### Results

We didn't perform a proper benchmark of our queries but we added an statement in there to get an idea of how long it was taking for the insert to finish. 
Here are the results obtained for the different versions used.

|Approach| Records Submitted |Time (ms) | 
|--------|-----------|-------|
| Single Multi-insert statement|100|4.98245ms|
| Single Multi-insert statement|10000|234.610741ms|
| Single Multi-insert statement|20000|491.543135ms|