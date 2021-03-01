CREATE TABLE `wagers`
(
    `id`                 integer PRIMARY KEY AUTO_INCREMENT,
    `total_wager_value`  integer UNSIGNED NOT NULL,
    `odds`               integer UNSIGNED NOT NULL,
    `selling_percentage` integer UNSIGNED NOT NULL,
    `selling_price`      decimal(10, 2)   NOT NULL,
    `placed_at`          timestamp        NOT NULL
);