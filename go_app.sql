CREATE TABLE `num_info`
(
    `id`       int(11) NOT NULL AUTO_INCREMENT,
    `name`     varchar(255) DEFAULT NULL,
    `info_key` varchar(255) NOT NULL,
    `info_num` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`, `info_key`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

CREATE TABLE `casbin_rule`
(
    `p_type` varchar(100) NOT NULL DEFAULT '',
    `v0`     varchar(100) NOT NULL DEFAULT '',
    `v1`     varchar(100) NOT NULL DEFAULT '',
    `v2`     varchar(100) NOT NULL DEFAULT '',
    `v3`     varchar(100) NOT NULL DEFAULT '',
    `v4`     varchar(100) NOT NULL DEFAULT '',
    `v5`     varchar(100) NOT NULL DEFAULT '',
    KEY      `IDX_casbin_rule_v4` (`v4`),
    KEY      `IDX_casbin_rule_v5` (`v5`),
    KEY      `IDX_casbin_rule_p_type` (`p_type`),
    KEY      `IDX_casbin_rule_v0` (`v0`),
    KEY      `IDX_casbin_rule_v1` (`v1`),
    KEY      `IDX_casbin_rule_v2` (`v2`),
    KEY      `IDX_casbin_rule_v3` (`v3`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `role_info`
(
    `id`   int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

CREATE TABLE `user_info`
(
    `id`          int(255) NOT NULL AUTO_INCREMENT,
    `name`        varchar(255) NOT NULL,
    `login_name`  varchar(255) NOT NULL,
    `role`        int(255) NOT NULL,
    `pwd`         varchar(255) NOT NULL,
    `create_time` datetime     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

