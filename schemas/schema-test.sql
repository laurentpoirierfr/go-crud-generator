-- Example schema.sql file

CREATE TABLE employees (
    employee_id NUMBER(6) NOT NULL,
    first_name VARCHAR2(20),
    last_name VARCHAR2(25) NOT NULL,
    email VARCHAR2(25) NOT NULL,
    phone_number VARCHAR2(20),
    hire_date DATE NOT NULL,
    job_id VARCHAR2(10) NOT NULL,
    salary NUMBER(8,2),
    commission_pct NUMBER(2,2),
    manager_id NUMBER(6),
    department_id NUMBER(4),
    PRIMARY KEY (employee_id)
);

CREATE TABLE departments (
    department_id NUMBER(4) NOT NULL,
    department_name VARCHAR2(30) NOT NULL,
    manager_id NUMBER(6),
    location_id NUMBER(4),
    PRIMARY KEY (department_id)
);

CREATE TABLE jobs (
    job_id VARCHAR2(10) NOT NULL,
    job_title VARCHAR2(35) NOT NULL,
    min_salary NUMBER(6),
    max_salary NUMBER(6),
    PRIMARY KEY (job_id)
);

CREATE TABLE job_history (
    employee_id NUMBER(6) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    job_id VARCHAR2(10) NOT NULL,
    department_id NUMBER(4),
    PRIMARY KEY (employee_id, start_date)
);

CREATE TABLE locations (
    location_id NUMBER(4) NOT NULL,
    street_address VARCHAR2(40),
    postal_code VARCHAR2(12),
    city VARCHAR2(30) NOT NULL,
    state_province VARCHAR2(25),
    country_id CHAR(2),
    PRIMARY KEY (location_id)
);

CREATE TABLE countries (
    country_id CHAR(2) NOT NULL,
    country_name VARCHAR2(40),
    region_id NUMBER,
    PRIMARY KEY (country_id)
);

CREATE TABLE regions (
    region_id NUMBER NOT NULL,
    region_name VARCHAR2(25),
    PRIMARY KEY (region_id)
);
