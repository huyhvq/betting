CREATE TABLE transactions
(
    id           int PRIMARY KEY AUTO_INCREMENT,
    wager_id     int            NOT NULL,
    buying_price decimal(10, 2) NOT NULL,
    created_at   timestamp      NOT NULL,
    updated_at   timestamp      NOT NULL,
    deleted_at   timestamp,
    FOREIGN KEY (wager_id) REFERENCES wagers (id)
);