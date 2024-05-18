CREATE TABLE users (
    id NUMBER PRIMARY KEY,
    username VARCHAR2(50),
    password VARCHAR2(50),
    email VARCHAR2(100)
);

CREATE TABLE products (
    product_id NUMBER PRIMARY KEY,
    name VARCHAR2(100),
    description VARCHAR2(255),
    price NUMBER
);
