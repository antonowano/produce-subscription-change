-- table copied from production database
create table subscription_change
(
    id                  int auto_increment
        primary key,
    cdate               int                  null,
    mdate               int                  null,
    email               varchar(255)         null,
    email_confirmed     tinyint(1) default 0 not null comment 'Email подтверждён (double opt-in)',
    phone               varchar(255)         null,
    user_id             int                  null,
    source              int                  not null comment 'Источник записи 0-ЛК, 1-из письма, 2-подтверждение почты, 3-регистрация, 4-1С',
    subscription_type   int                  not null comment '0-email, 1-sms',
    subscription_action int                  not null comment '0-отказаться, 1-согласиться',
    reason_unsubscribe  text                 null,
    is_processed        tinyint(1) default 0 not null comment 'Обработана ли строка'
)
    comment 'Лог действий по подпискам на рассылки смс и email' charset = utf8;

-- subscribe phone from source 1
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (1,  1618820074, null, null, 0, '+7(123)123-12-31', 4105022, 1, 1, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (2,  1624542732, null, null, 0, '+7(930)288-08-09', 4105028, 1, 1, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (3,  1616488964, null, null, 0, '+7(930)288-08-07', 4105014, 1, 1, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (4,  1616488964, null, null, 0, '+7(645)123-12-12', 4105021, 1, 1, 1, null, 1);

-- subscribe phone from source 2
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (5,  1618820074, null, null, 0, '+7(123)123-12-01', 4105022, 2, 1, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (6,  1624542732, null, null, 0, '+7(930)288-08-02', 4105028, 2, 1, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (7,  1616488964, null, null, 0, '+7(930)288-08-03', 4105014, 2, 1, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (8,  1616488964, null, null, 0, '+7(645)123-12-04', 4105021, 2, 1, 1, null, 1);

-- subscribe phone from source 3
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (9,  1618820074, null, null, 0, '+7(123)123-12-05', 4105022, 3, 1, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (10, 1624542732, null, null, 0, '+7(930)288-08-06', 4105028, 3, 1, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (11, 1616488964, null, null, 0, '+7(930)288-08-07', 4105014, 3, 1, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (12, 1616488964, null, null, 0, '+7(645)123-12-08', 4105021, 3, 1, 1, null, 1);

-- subscribe phone from source 4
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (13, 1624544305, null, null, 0, '+7(930)288-08-12', 4105035, 4, 1, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (14, 1624542554, null, null, 0, '+7(930)288-08-08', 4105027, 4, 1, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (15, 1624543593, null, null, 0, '+7(930)288-08-10', 4105029, 4, 1, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (16, 1629963582, null, null, 0, '+7(654)123-45-67', 4105036, 4, 1, 1, null, 1);

-- subscribe email from source 1
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (17, 1618493588, null, 'ivan2@antonov.site',       0, null, 4105021, 1, 0, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (18, 1619436542, null, 'qwerty@rambler.ru',        1, null, 4105023, 1, 0, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (19, 1629963582, null, 'ivan3@antonov.site',       1, null, 4105036, 1, 0, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (20, 1639996526, null, 'ivan@antonov.site',        0, null, 4105037, 1, 0, 1, null, 1);

-- subscribe email from source 2
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (21, 1618493588, null, 'ivan2@antonov.site',       1, null, 4105021, 2, 0, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (22, 1619436542, null, 'qwerty@rambler.ru',        0, null, 4105023, 2, 0, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (23, 1629963582, null, 'ivan3@antonov.site',       0, null, 4105036, 2, 0, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (24, 1639996526, null, 'ivan@antonov.site',        1, null, 4105037, 2, 0, 1, null, 1);

-- subscribe email from source 3
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (25, 1618493588, null, 'ivan2@antonov.site',       1, null, 4105021, 3, 0, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (26, 1619436542, null, 'qwerty@rambler.ru',        0, null, 4105023, 3, 0, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (27, 1629963582, null, 'ivan3@antonov.site',       1, null, 4105036, 3, 0, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (28, 1639996526, null, 'ivan@antonov.site',        0, null, 4105037, 3, 0, 1, null, 1);

-- subscribe email from source 4
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (29, 1618493588, null, 'ivan2@antonov.site',       1, null, 4105021, 4, 0, 0, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (30, 1619436542, null, 'qwerty@rambler.ru',        1, null, 4105023, 4, 0, 0, null, 1);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (31, 1629963582, null, 'ivan3@antonov.site',       1, null, 4105036, 4, 0, 1, null, 0);
INSERT INTO subscription_change (id, cdate, mdate, email, email_confirmed, phone, user_id, source, subscription_type, subscription_action, reason_unsubscribe, is_processed) VALUES (32, 1639996526, null, 'ivan@antonov.site',        1, null, 4105037, 4, 0, 1, null, 1);
