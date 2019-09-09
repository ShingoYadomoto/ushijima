
CREATE TABLE months (
    id          SERIAL PRIMARY key,
    display     VARCHAR(255),
    create_date TIMESTAMP,
    update_date TIMESTAMP
);

CREATE TABLE payment_statuses (
    id          SERIAL PRIMARY key,
    name        VARCHAR(255),
    display     VARCHAR(255),
    create_date TIMESTAMP,
    update_date TIMESTAMP
);

CREATE TABLE payment_types (
    id          SERIAL PRIMARY key,
    name        VARCHAR(255),
    display     VARCHAR(255),
    create_date TIMESTAMP,
    update_date TIMESTAMP
);

CREATE TABLE payments (
    id                SERIAL PRIMARY key,
    payment_type_id   int references payment_types(id),
    payment_status_id int references payment_statuses(id),
    month_id          int references months(id),
    amount            int,
    create_date       TIMESTAMP,
    update_date       TIMESTAMP
);

INSERT INTO months (display) VALUES ('2019-09'), ('2019-10'), ('2019-11'), ('2019-12'),('2020-01'), ('2020-02'),
('2020-03'), ('2020-04'), ('2020-05'), ('2020-06'), ('2020-07'), ('2020-08'), ('2020-09'), ('2020-10'), ('2020-11')

INSERT INTO payment_statuses (name, display) VALUES ('paid', '支払済'), ('unpaid', '未支払い'), ('fumiya paid', 'fumiya支払済'), ('shingo paid', 'shingo支払済')

INSERT INTO payment_types (name, display) VALUES ('rent bill', '家賃'), ('electric bill', '電気代'), ('gas bill', 'ガス代'), ('water bill', '水道代'), ('wifi bill', 'WiFi代')

INSERT INTO payments (payment_type_id, payment_status_id, month_id, amount) VALUES
(1, 2,  1, 0), (2, 2,  1, 0), (3, 2,  1, 0), (4, 2,  1, 0), (5, 2,  1, 0),
(1, 2,  2, 0), (2, 2,  2, 0), (3, 2,  2, 0), (4, 2,  2, 0), (5, 2,  2, 0),
(1, 2,  3, 0), (2, 2,  3, 0), (3, 2,  3, 0), (4, 2,  3, 0), (5, 2,  3, 0),
(1, 2,  4, 0), (2, 2,  4, 0), (3, 2,  4, 0), (4, 2,  4, 0), (5, 2,  4, 0),
(1, 2,  5, 0), (2, 2,  5, 0), (3, 2,  5, 0), (4, 2,  5, 0), (5, 2,  5, 0),
(1, 2,  6, 0), (2, 2,  6, 0), (3, 2,  6, 0), (4, 2,  6, 0), (5, 2,  6, 0),
(1, 2,  7, 0), (2, 2,  7, 0), (3, 2,  7, 0), (4, 2,  7, 0), (5, 2,  7, 0),
(1, 2,  8, 0), (2, 2,  8, 0), (3, 2,  8, 0), (4, 2,  8, 0), (5, 2,  8, 0),
(1, 2,  9, 0), (2, 2,  9, 0), (3, 2,  9, 0), (4, 2,  9, 0), (5, 2,  9, 0),
(1, 2, 10, 0), (2, 2, 10, 0), (3, 2, 10, 0), (4, 2, 10, 0), (5, 2, 10, 0),
(1, 2, 11, 0), (2, 2, 11, 0), (3, 2, 11, 0), (4, 2, 11, 0), (5, 2, 11, 0),
(1, 2, 12, 0), (2, 2, 12, 0), (3, 2, 12, 0), (4, 2, 12, 0), (5, 2, 12, 0),
(1, 2, 13, 0), (2, 2, 13, 0), (3, 2, 13, 0), (4, 2, 13, 0), (5, 2, 13, 0),
(1, 2, 14, 0), (2, 2, 14, 0), (3, 2, 14, 0), (4, 2, 14, 0), (5, 2, 14, 0),
(1, 2, 15, 0), (2, 2, 15, 0), (3, 2, 15, 0), (4, 2, 15, 0), (5, 2, 15, 0)