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
INSERT INTO auster.vendor (id, name, address) VALUES ('gyo3', '銚子市観光課', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）');

CREATE TABLE `travel_spot`
(
    `id`          varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `vendor_id`   varchar(20) COLLATE utf8mb4_bin  NOT NULL,
    `title`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `take_time`   int(11)                          NOT NULL COMMENT '所要時間（分）',
    `description` text COLLATE utf8mb4_bin        NOT NULL COMMENT '説明',
    `address`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `photo_path`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポット';

-- 体験スポット
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('1', 'gyo1', '釣り体験 ヒラマサ', 180, '船釣りでヒラマサを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/1/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('10', 'can2', 'キャンプ場運営体験 イベントのスタッフ', 300, '周辺地域のショップ様と共同開催するイベントのスタッフ業務をお任せします', '千葉県大網白里市四天木2761-40', '/assets/images/travel_spots/10/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('11', 'can1', 'キャンプ中級者体験 ブッシュクラフト火おこし', 60, 'キャンパーが一度は憧れるブッシュクラフト、火おこし体験ができます', '千葉県長生郡一宮町東浪見1611', '/assets/images/travel_spots/11/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('12', 'can3', 'トレインランニング', 300, '銚子電鉄の駅を自分の足で巡りながら大自然を駆け抜けるトレインランニングを体験できます', '千葉県銚子市新生町2丁目297番地', '/assets/images/travel_spots/12/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('13', 'can2', 'スペシャルオファー<複業> ショップの企画担当', 0, 'キャンプ場周辺のショップで企画担当を募集します', '千葉県大網白里市四天木2761-40', '/assets/images/travel_spots/13/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('2', 'gyo1', '釣り体験 シイラ', 240, '船釣りでクロカジキを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/2/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('3', 'gyo1', '釣り体験 ヒラメ', 180, '船釣りでヒラメを釣る体験ができます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/3/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('4', 'gyo1', '漁業体験 競りの見学', 120, '競りの市場を見学できます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/4/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('5', 'gyo1', '漁業体験 定置網漁', 180, '定置網漁を体験できます', '千葉県銚子市川口町 2丁目6528番地', '/assets/images/travel_spots/5/');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('6', 'gyo2', '社会貢献活動 浜辺のゴミ拾い', 120, '銚子の海岸に打ち上げられた漂流物の清掃を実施します', '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）', '/assets/images/travel_spots/6/');

-- 観光スポット
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('14', 'gyo3', '犬吠埼灯台', 60, '銚子のシンボルとして知られる白亜の灯台です。本州で最も早く朝日が昇る場所としても有名です。', '千葉県銚子市犬吠埼957612', '');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('15', 'gyo3', '地球の丸く見える丘展望館', 90, '標高90mの屋上展望スペースからは360度の大パノラマが楽しめます。水平線の両端が丸みを帯びて見えることから、その名が付けられました5。', '千葉県銚子市天王台1421-13', '');
INSERT INTO auster.travel_spot (id, vendor_id, title, take_time, description, address, photo_path) VALUES ('16', 'gyo3', '圓福寺（飯沼観音）', 45, '真言宗の寺院で、貴重な美術品や古文書が保管されています。特に「釈迦涅槃殿」は千葉県の重要文化財に指定されています3。', '千葉県銚子市馬場町2934', '');

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

CREATE TABLE `travel_spot_itinerary`
(
    `id`             varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `kind`           enum ('experience','spot','move')            NOT NULL COMMENT 'experience: 体験場所, spot: 観光地, move: 移動',
    `travel_spot_id` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT 'kindがspotの場合のみ',
    `take_time`      int(11)                         NOT NULL COMMENT '所要時間（分）',
    `price`          int(11)                         NOT NULL,
    `order`          int(11)                         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='旅行先の体験スポットと旅程の関連';

-- 旅程: 体験場所
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti1', 'experience', '1', 180, 15000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti2', 'experience', '2', 240, 20000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti3', 'experience', '3', 180, 15000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti4', 'experience', '4', 120, 5000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti5', 'experience', '5', 180, 12000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti6', 'experience', '6', 120, 0, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti10', 'experience', '10', 300, 8000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti11', 'experience', '11', 60, 5000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti12', 'experience', '12', 300, 3000, 4);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti13', 'experience', '13', 0, 0, 4);
-- 旅程: 観光スポット
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti14', 'spot', '14', 60, 500, 2);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti15', 'spot', '15', 90, 800, 2);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti16', 'spot', '16', 45, 300, 2);
-- 旅程: 移動時間(使い回し)
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti17', 'move', '', 30, 400, 1);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti18', 'move', '', 20, 300, 3);
INSERT INTO auster.travel_spot_itinerary (id, kind, travel_spot_id, take_time, price, `order`) VALUES ('iti19', 'move', '', 60, 200, 5);

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

-- 釣り体験（ヒラマサ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_1', 'iti1', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_2', 'iti1', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_3', 'iti1', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item1_4', 'iti1', '酔い止め薬', 1);

-- 釣り体験（クロカジキ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_1', 'iti2', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_2', 'iti2', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_3', 'iti2', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item2_4', 'iti2', '酔い止め薬', 1);

-- 釣り体験（ヒラメ）の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_1', 'iti3', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_2', 'iti3', '日焼け止め', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_3', 'iti3', 'タオル', 2);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item3_4', 'iti3', '酔い止め薬', 1);

-- 競りの見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_1', 'iti4', '長靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_2', 'iti4', '防寒着', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item4_3', 'iti4', 'メモ帳', 1);

-- 定置網漁体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_1', 'iti5', '長靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_2', 'iti5', '作業用手袋', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_3', 'iti5', '防寒着', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item5_4', 'iti5', 'タオル', 2);

-- 浜辺のゴミ拾いの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_1', 'iti6', '軍手', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_2', 'iti6', '帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item6_3', 'iti6', '飲み物', 1);

-- キャンプ場運営体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_1', 'iti10', '作業用手袋', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_2', 'iti10', '動きやすい服装', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item10_3', 'iti10', '長靴', 1);

-- キャンプ中級者体験の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_1', 'iti11', '軍手', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_2', 'iti11', '防虫スプレー', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item11_3', 'iti11', '長袖の服', 1);

-- トレインランニングの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_1', 'iti12', 'ランニングシューズ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_2', 'iti12', '水筒', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item12_3', 'iti12', 'タオル', 2);

-- 複業オファーの持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_1', 'iti13', '筆記用具', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_2', 'iti13', '履歴書', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item13_3', 'iti13', 'ポートフォリオ', 1);

-- 犬吠埼灯台見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_1', 'iti14', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_2', 'iti14', '日傘・帽子', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item14_3', 'iti14', '歩きやすい靴', 1);

-- 地球の丸く見える丘展望館の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_1', 'iti15', 'カメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item15_2', 'iti15', '双眼鏡', 1);

-- 圓福寺見学の持ち物
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_1', 'iti16', '歩きやすい靴', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_2', 'iti16', '写真撮影許可が必要なカメラ', 1);
INSERT INTO auster.travel_spot_itinerary_item (id, travel_spot_itinerary_id, name, number) VALUES ('item16_3', 'iti16', 'お賽銭', 1);

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
    `travel_spot_id`       varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `travel_spot_diary_id` varchar(20) COLLATE utf8mb4_bin NOT NULL,
    `from_date`            date                            NOT NULL,
    `to_date`              date                            NOT NULL,
    `is_offer`             tinyint(1)                      NOT NULL,
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