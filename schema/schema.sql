USE `auster`;

CREATE TABLE `user`
(
    `id`           varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `gender`       varchar(10) COLLATE utf8mb4_bin  NOT NULL,
    `age`          int(11)                          NOT NULL,
    `profile_path` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザ';

CREATE TABLE `user_itinerary_diary`
(
    `id`      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `user_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `date`    date                             NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザの旅行記';

CREATE TABLE `itinerary`
(
    `id`                      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `user_itinerary_diary_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `kind`                    enum ('spot','move')             NOT NULL COMMENT 'spot: 観光地, move: 移動',
    `travel_spot_id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL COMMENT 'kindがspotの場合のみ',
    `from`                    varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'kindがmoveの場合のみ',
    `to`                      varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'kindがmoveの場合のみ',
    `take_time`               int(11)                          NOT NULL COMMENT '移動にかかる時間（分）, kindがspotの場合は0',
    `order`                   int(11)                          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅程';

CREATE TABLE `itinerary_result`
(
    `id`           varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `itinerary_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `done`         tinyint(1)                       NOT NULL COMMENT 'この旅程を行ったかどうか',
    `photo_path`   varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `description`  text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅程の結果';

CREATE TABLE `travel_spot`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `vendor_id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `take_time`   int(11)                          NOT NULL COMMENT '体験にかかる時間（分）',
    `description` text COLLATE utf8mb4_bin         NOT NULL COMMENT '説明',
    `address`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポット';

CREATE TABLE `travel_spot_photo`
(
    `id`             varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`           varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `path`           varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットでの写真';

CREATE TABLE `travel_spot_item`
(
    `id`             varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`           varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `number`         int(11)                          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットで必要な持ち物';

CREATE TABLE `vendor`
(
    `id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行会社';