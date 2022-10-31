CREATE TABLE customer
(
    customer_id bigint NOT NULL AUTO_INCREMENT,
    balance numeric(50, 10),

    PRIMARY KEY(customer_id),
    CHECK (balance >= 0)
);

CREATE TABLE service
(
    service_id bigint NOT NULL AUTO_INCREMENT,
    service_name varchar(30) NOT NULL,

    PRIMARY KEY(service_id)

);

CREATE TABLE order_service
(
    order_id bigint NOT NULL AUTO_INCREMENT,
    service_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    order_date date NOT NULL,
    summ numeric(50, 10),

    PRIMARY KEY(order_id),
    FOREIGN KEY (service_id) REFERENCES service(service_id),
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id),
    CHECK (summ >= 0)
);

CREATE TABLE reserve
(
    reserve_id bigint NOT NULL AUTO_INCREMENT,
    customer_id bigint NOT NULL,
    service_id bigint NOT NULL,
    order_id bigint NOT NULL,
    summ numeric(50, 10),

    PRIMARY KEY(reserve_id),
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id),
    FOREIGN KEY (service_id) REFERENCES service(service_id),
    FOREIGN KEY (order_id) REFERENCES order_service(order_id),
    CHECK (summ >= 0)
);

CREATE TABLE profit
(
    profit_id bigint NOT NULL AUTO_INCREMENT,
    receiving_date date NOT NULL,
    order_id bigint NOT NULL,
    summ numeric(50, 10),

    PRIMARY KEY(profit_id),
    FOREIGN KEY (order_id) REFERENCES order_service(order_id),
    CHECK (summ >=0)
);

CREATE TABLE history
(
    history_id bigint NOT NULL AUTO_INCREMENT,
    customer_id bigint NOT NULL,
    summ numeric(50, 10),
    service_id bigint NOT NULL,
    execute_date date NOT NULL,

    PRIMARY KEY(history_id),
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id),
    FOREIGN KEY (service_id) REFERENCES service(service_id),
    CHECK (summ >= 0)
)