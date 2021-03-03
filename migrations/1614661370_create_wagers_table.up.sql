CREATE TABLE `wagers`
(
    `id`                    integer PRIMARY KEY AUTO_INCREMENT,
    `total_wager_value`     integer UNSIGNED NOT NULL,
    `odds`                  integer UNSIGNED NOT NULL,
    `selling_percentage`    integer UNSIGNED NOT NULL,
    `selling_price`         decimal(10, 2)   NOT NULL,
    `current_selling_price` decimal(10, 2)   NOT NULL,
    `amount_sold`           decimal(10, 2)   NOT NULL,
    `percentage_sold`       integer UNSIGNED NOT NULL,
    `created_at`            timestamp        NOT NULL,
    `updated_at`            timestamp        NOT NULL,
    `deleted_at`            timestamp
);