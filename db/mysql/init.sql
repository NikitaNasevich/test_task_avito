CREATE TABLE customer
(
    customer_id bigint NOT NULL AUTO_INCREMENT,
    balance     numeric(50, 10),

    PRIMARY KEY (customer_id),
    CHECK (balance >= 0)
);

CREATE TABLE service
(
    service_id   bigint      NOT NULL AUTO_INCREMENT,
    service_name varchar(30) NOT NULL,

    PRIMARY KEY (service_id)

);

CREATE TABLE order_service
(
    order_id    bigint NOT NULL AUTO_INCREMENT,
    service_id  bigint NOT NULL,
    customer_id bigint NOT NULL,
    order_date  date   NOT NULL,
    summ        numeric(50, 10),

    PRIMARY KEY (order_id),
    FOREIGN KEY (service_id) REFERENCES service (service_id),
    FOREIGN KEY (customer_id) REFERENCES customer (customer_id),
    CHECK (summ >= 0)
);

CREATE TABLE reserve
(
    reserve_id  bigint NOT NULL AUTO_INCREMENT,
    customer_id bigint NOT NULL,
    service_id  bigint NOT NULL,
    order_id    bigint NOT NULL,
    summ        numeric(50, 10),

    PRIMARY KEY (reserve_id),
    FOREIGN KEY (customer_id) REFERENCES customer (customer_id),
    FOREIGN KEY (service_id) REFERENCES service (service_id),
    FOREIGN KEY (order_id) REFERENCES order_service (order_id),
    CHECK (summ >= 0)
);

CREATE TABLE profit
(
    profit_id      bigint NOT NULL AUTO_INCREMENT,
    receiving_date date   NOT NULL,
    order_id       bigint NOT NULL,
    summ           numeric(50, 10),

    PRIMARY KEY (profit_id),
    FOREIGN KEY (order_id) REFERENCES order_service (order_id),
    CHECK (summ >= 0)
);

CREATE TABLE history
(
    history_id   bigint NOT NULL AUTO_INCREMENT,
    customer_id  bigint NOT NULL,
    summ         numeric(50, 10),
    service_id   bigint NOT NULL,
    execute_date date   NOT NULL,

    PRIMARY KEY (history_id),
    FOREIGN KEY (customer_id) REFERENCES customer (customer_id),
    FOREIGN KEY (service_id) REFERENCES service (service_id),
    CHECK (summ >= 0)
);

INSERT INTO customer (customer_id)
VALUES (1),
       (2),
       (3),
       (4),
       (5),
       (6),
       (7),
       (8),
       (9),
       (10);

INSERT INTO service (service_name)
VALUES ('Транспорт'),
       ('Сад'),
       ('Ремонт'),
       ('Обучение'),
       ('Красота');

INSERT INTO order_service (service_id, customer_id, order_date, summ)
VALUES (1, 1, '2022-10-05', 10),
       (2, 3, '2022-10-15', 100),
       (3, 2, '2022-10-25', 350);