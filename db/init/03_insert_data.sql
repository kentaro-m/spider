USE `spiderdb`;

INSERT INTO `sites`
(`id`,
 `title`,
 `url`,
 `created_at`,
 `updated_at`)
VALUES
('faf9c3a7-b3ee-441f-baec-a5b668948382',
 'Learn Something New',
 'https://blog.kentarom.com',
 '2019-04-06 16:03:31',
 '2019-04-06 16:03:31');

INSERT INTO `articles`
(`id`,
`title`,
`url`,
`pub_date`,
`site_id`,
`created_at`,
`updated_at`)
VALUES
('faf9c3a7-b3ee-441f-baec-a5b668948381',
'AWS CDKでサーバーレスアプリケーションのデプロイを試す',
'https://blog.kentarom.com/learn-aws-cdk/',
'2019-01-19 14:13:01',
'faf9c3a7-b3ee-441f-baec-a5b668948382',
'2019-04-06 16:03:31',
'2019-04-06 16:03:31');
