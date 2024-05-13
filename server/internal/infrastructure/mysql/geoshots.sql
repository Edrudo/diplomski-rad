create database diplomski;
use diplomski;
CREATE TABLE `geoshots` (
                            `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
                            `eventID` int(11) DEFAULT NULL,
                            `deviceID` int(11) DEFAULT NULL,
                            `imgpath` varchar(100) DEFAULT NULL,
                            `lat` double DEFAULT NULL,
                            `lon` double DEFAULT NULL,
                            `timestamp` datetime DEFAULT NULL,
                            `age` int(11) DEFAULT NULL,
                            `buffered` int(1) DEFAULT NULL,
                            `onstage` int(1) DEFAULT NULL,
                            `eventhos` int(1) DEFAULT NULL,
                            `eventhosStat` int(1) DEFAULT NULL,
                            `jsonpath` varchar(100) DEFAULT NULL,
                            `synced` int(1) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
create user 'edrudo'@'172.17.0.1';
GRANT ALL PRIVILEGES ON diplomski.* TO 'edrudo'@'172.17.0.1';