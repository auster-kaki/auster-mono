USE `auster`;

CREATE TABLE `user`
(
    `id`           varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`         varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `gender`       varchar(10) COLLATE utf8mb4_bin  NOT NULL,
    `age`          int(11)                          NOT NULL,
    `profile_path` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザ';

CREATE TABLE `hobby`
(
    `id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='趣味';

INSERT INTO auster.hobby (id, name) VALUES ('cstkdiat6c3011a83so0', '釣り');
INSERT INTO auster.hobby (id, name) VALUES ('cstkdiat6c3011a83sog', 'キャンプ');

CREATE TABLE `user_hobby`
(
    `user_id`  varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `hobby_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`user_id`, `hobby_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザの趣味';

CREATE TABLE `user_itinerary_history`
(
    `id`      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `user_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `date`    date                             NOT NULL,
    PRIMARY KEY (`id`),
    KEY       `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='ユーザの旅行記';

CREATE TABLE `itinerary`
(
    `id`                        varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `user_itinerary_history_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `kind`                      enum ('spot','move')             NOT NULL COMMENT 'spot: 観光地, move: 移動',
    `travel_spot_id`            varchar(20) COLLATE utf8mb4_bin  NOT NULL COMMENT 'kindがspotの場合のみ',
    `from`                      varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'kindがmoveの場合のみ',
    `to`                        varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'kindがmoveの場合のみ',
    `take_time`                 int(11)                          NOT NULL COMMENT '移動にかかる時間（分）, kindがspotの場合は0',
    `order`                     int(11)                          NOT NULL,
    PRIMARY KEY (`id`),
    KEY                         `user_itinerary_history_id` (`user_itinerary_history_id`)
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

CREATE TABLE `diary`
(
    `id`         varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`      varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `date`       date                             NOT NULL,
    `photo_path` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `body`       text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='日記';

CREATE TABLE `diary_tag`
(
    `id`       varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `diary_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='日記のタグ';

CREATE TABLE `diary_user`
(
    `diary_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `user_id`  varchar(20) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`diary_id`, `user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='日記のユーザ';

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

INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('1', 'gyo1', '釣り体験 ヒラマサ', 180, '船釣りでヒラマサを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('10', 'can2', 'キャンプ場運営体験 イベントのスタッフ', 300, '周辺地域のショップ様と共同開催するイベントのスタッフ業務をお任せします', '千葉県大網白里市四天木2761-40');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('11', 'can1', 'キャンプ中級者体験 ブッシュクラフト火おこし', 60, 'キャンパーが一度は憧れるブッシュクラフト、火おこし体験ができます', '千葉県長生郡一宮町東浪見1611');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('12', 'can3', 'トレインランニング', 300, '銚子電鉄の駅を自分の足で巡りながら大自然を駆け抜けるトレインランニングを体験できます', '千葉県銚子市新生町2丁目297番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('13', 'can2', 'スペシャルオファー<複業> ショップの企画担当', 0, 'キャンプ場周辺のショップで企画担当を募集します', '千葉県大網白里市四天木2761-40');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('2', 'gyo1', '釣り体験 クロカジキ', 240, '船釣りでクロカジキを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('3', 'gyo1', '釣り体験 ヒラメ', 180, '船釣りでヒラメを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('4', 'gyo1', '漁業体験 競りの見学', 120, '競りの市場を見学できます', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('5', 'gyo1', '漁業体験 定置網漁', 180, '定置網漁を体験できます', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.travel_spot (id, vendor_id, name, take_time, description, address) VALUES ('6', 'gyo2', '社会貢献活動 浜辺のゴミ拾い', 120, '銚子の海岸に打ち上げられた漂流物の清掃を実施します', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');

CREATE TABLE `travel_spot_hobby`
(
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `hobby_id`       varchar(20) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`travel_spot_id`, `hobby_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='体験スポットに当てはまる趣味';

INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('1', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('2', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('3', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('4', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('5', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('6', 'cstkdiat6c3011a83so0');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('10', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('11', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('12', 'cstkdiat6c3011a83sog');
INSERT INTO auster.travel_spot_hobby (travel_spot_id, hobby_id) VALUES ('13', 'cstkdiat6c3011a83sog');

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

INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('1', '1', '釣り体験 ヒラマサ画像','/assets/images/travel_spot_photo_1.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('2', '10', 'キャンプ場運営体験 イベントのスタッフ画像','/assets/images/travel_spot_photo_2.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('3', '11', 'キャンプ中級者体験 ブッシュクラフト火おこし画像','/assets/images/travel_spot_photo_3.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('4', '12', 'トレインランニング画像','/assets/images/travel_spot_photo_4.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('5', '13', 'スペシャルオファー<複業画像','/assets/images/travel_spot_photo_5.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('6', '2', '釣り体験 クロカジキ画像','/assets/images/travel_spot_photo_6.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('7', '3', '釣り体験 ヒラメ画像','/assets/images/travel_spot_photo_7.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('8', '4', '漁業体験 競りの見学画像','/assets/images/travel_spot_photo_8.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('9', '5', '漁業体験 定置網漁画像','/assets/images/travel_spot_photo_9.jpg');
INSERT INTO auster.travel_spot_photo (id, travel_spot_id, name, path) VALUES ('10', '6', '社会貢献活動 浜辺のゴミ拾い画像','/assets/images/travel_spot_photo_10.jpg');

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
    `id`      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行会社';

INSERT INTO auster.vendor (id, name, address) VALUES ('can1', 'CampsiteTORAMI', '千葉県長生郡一宮町東浪見1611');
INSERT INTO auster.vendor (id, name, address) VALUES ('can2', 'Beach Camp 九十九里', '千葉県大網白里市四天木2761-40');
INSERT INTO auster.vendor (id, name, address) VALUES ('can3', '銚子電気鉄道', '千葉県銚子市新生町2丁目297番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo1', '銚子市漁業協同組合', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo2', '銚子市生活環境課 清掃美化班', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');


CREATE TABLE `encounter`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `place`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `user_id`     varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `date`        date                             NOT NULL,
    `description` text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`),
    KEY           `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='出会った人';

--

CREATE TABLE `vendor`
(
    `id`      varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行会社';

INSERT INTO auster.vendor (id, name, address) VALUES ('can1', 'CampsiteTORAMI', '千葉県長生郡一宮町東浪見1611');
INSERT INTO auster.vendor (id, name, address) VALUES ('can2', 'Beach Camp 九十九里', '千葉県大網白里市四天木2761-40');
INSERT INTO auster.vendor (id, name, address) VALUES ('can3', '銚子電気鉄道', '千葉県銚子市新生町2丁目297番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo1', '銚子市漁業協同組合', '千葉県銚子市川口町 2丁目6528番地');
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo2', '銚子市生活環境課 清掃美化班', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');

CREATE TABLE `travel_spot_v2`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `vendor_id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `description` text COLLATE utf8mb4_bin         NOT NULL COMMENT '説明',
    `address`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポット';

CREATE TABLE `travel_spot_itinerary`
(
    `id`                varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `kind`              enum ('spot','move')             NOT NULL COMMENT 'spot: 観光地, move: 移動',
    `travel_spot_v2_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL COMMENT 'kindがspotの場合のみ',
    `take_time`         int(11)                          NOT NULL COMMENT '所要時間（分）',
    `price`             int(11)                          NOT NULL,
    `order`             int(11)                          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットと旅程の関連';

CREATE TABLE `travel_spot_itinerary_item`
(
    `id`                       varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `travel_spot_itinerary_id` varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`                     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `number`                   int(11)                          NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットで必要な持ち物';

CREATE TABLE `travel_spot_diary`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `date`        date                             NOT NULL,
    `photo_path`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `description` text COLLATE utf8mb4_bin         NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='日記';

CREATE TABLE `reservation`
(
    `id`                   varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `user_id`              varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `travel_spot_v2_id`    varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `travel_spot_diary_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `date`                 date                            NOT NULL,
    `is_offer`             tinyint(1)                      NOT NULL,
    `number`               int(11)                         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='予約';

CREATE TABLE `encounter`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `name`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `place`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `user_id`     varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `date`        date                             NOT NULL,
    `description` text COLLATE utf8mb4_bin         NOT NULL,
    PRIMARY KEY (`id`),
    KEY           `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='出会った人';