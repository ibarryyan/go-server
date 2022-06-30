CREATE TABLE `num_info`
(
    `id`       int(11) NOT NULL AUTO_INCREMENT,
    `name`     varchar(255) DEFAULT NULL,
    `info_key` varchar(255) NOT NULL,
    `info_num` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`, `info_key`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

