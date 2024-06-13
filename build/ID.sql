-- -------------------------------------------------------------
-- TablePlus 6.0.0(550)
--
-- https://tableplus.com/
--
-- Database: ID
-- Generation Time: 2024-06-13 14:10:48.4370
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS dbconnectlogs_id_seq;

-- Table Definition
CREATE TABLE "public"."dbconnectlogs" (
    "id" int4 NOT NULL DEFAULT nextval('dbconnectlogs_id_seq'::regclass),
    "time" timestamp NOT NULL,
    "env" varchar(12) NOT NULL,
    "comment" text,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."refreshtokens" (
    "userid" bpchar(16) NOT NULL,
    "token" text NOT NULL,
    "time" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("userid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."templogincode" (
    "userid" varchar(16) NOT NULL,
    "code" text NOT NULL,
    "time" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."twofatemp" (
    "userid" varchar(16) NOT NULL,
    "token" text NOT NULL,
    "time" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("userid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."users" (
    "userid" varchar(16) NOT NULL,
    "displayid" int4 NOT NULL,
    "email" varchar(255),
    "phone" varchar(20),
    "name" varchar(255),
    "surname" varchar(255),
    "joindate" timestamp DEFAULT CURRENT_TIMESTAMP,
    "verify" bool DEFAULT false,
    "twofa" bool DEFAULT false,
    PRIMARY KEY ("userid")
);

INSERT INTO "public"."dbconnectlogs" ("id", "time", "env", "comment") VALUES
(1, '2024-06-06 20:24:19.116269', 'local', 'development'),
(2, '2024-06-06 20:57:48.70675', 'local', 'development'),
(3, '2024-06-07 14:25:27.676754', 'local', 'development'),
(4, '2024-06-07 15:25:05.249354', 'local', 'development'),
(5, '2024-06-07 15:37:26.28333', 'local', 'development'),
(6, '2024-06-07 15:38:59.256423', 'local', 'development'),
(7, '2024-06-07 15:53:29.259598', 'local', 'development'),
(8, '2024-06-07 16:55:47.913761', 'local', 'development'),
(9, '2024-06-07 16:56:17.521043', 'local', 'development'),
(10, '2024-06-07 17:03:22.809637', 'local', 'development'),
(11, '2024-06-07 17:03:33.230638', 'local', 'development'),
(12, '2024-06-07 17:05:23.198429', 'local', 'development'),
(13, '2024-06-08 11:19:03.575913', 'local', 'development'),
(14, '2024-06-08 11:36:42.480669', 'local', 'development'),
(15, '2024-06-08 11:39:22.905842', 'local', 'development'),
(16, '2024-06-08 12:03:46.836739', 'local', 'development'),
(17, '2024-06-08 12:05:31.328067', 'local', 'development'),
(18, '2024-06-08 12:06:04.020701', 'local', 'development'),
(19, '2024-06-08 12:08:09.404814', 'local', 'development'),
(20, '2024-06-08 13:05:57.28108', 'local', 'development'),
(21, '2024-06-08 13:26:14.254005', 'local', 'development'),
(22, '2024-06-08 13:32:31.355988', 'local', 'development'),
(23, '2024-06-08 15:25:20.127577', 'local', 'development'),
(24, '2024-06-08 15:27:36.820364', 'local', 'development'),
(25, '2024-06-08 15:27:59.43777', 'local', 'development'),
(26, '2024-06-08 15:30:41.345498', 'local', 'development'),
(27, '2024-06-08 15:35:11.695111', 'local', 'development'),
(28, '2024-06-08 15:40:28.463218', 'local', 'development'),
(29, '2024-06-08 15:41:11.235461', 'local', 'development'),
(30, '2024-06-08 16:22:37.678953', 'local', 'development'),
(31, '2024-06-08 16:38:05.95253', 'local', 'development'),
(32, '2024-06-09 09:29:40.216782', 'local', 'development'),
(33, '2024-06-09 09:48:44.350964', 'local', 'development'),
(34, '2024-06-09 09:53:56.551059', 'local', 'development'),
(35, '2024-06-11 16:13:44.384369', 'local', 'development'),
(36, '2024-06-11 16:14:46.905646', 'local', 'development'),
(37, '2024-06-12 13:54:45.321732', 'local', 'development'),
(38, '2024-06-12 13:55:18.717618', 'local', 'development'),
(39, '2024-06-12 13:57:10.281503', 'local', 'development'),
(40, '2024-06-12 13:59:40.110271', 'local', 'development'),
(41, '2024-06-12 14:19:06.410052', 'local', 'development'),
(42, '2024-06-12 14:21:22.845587', 'local', 'development'),
(43, '2024-06-12 14:30:24.231211', 'local', 'development'),
(44, '2024-06-12 14:33:44.274552', 'local', 'development'),
(45, '2024-06-12 14:36:06.46937', 'local', 'development'),
(46, '2024-06-12 15:01:01.780313', 'local', 'development'),
(47, '2024-06-12 15:02:01.946781', 'local', 'development'),
(48, '2024-06-12 15:04:12.690696', 'local', 'development'),
(49, '2024-06-12 15:07:20.457915', 'local', 'development'),
(50, '2024-06-12 15:19:59.212269', 'local', 'development'),
(51, '2024-06-12 15:24:58.119698', 'local', 'development'),
(52, '2024-06-12 15:40:37.52009', 'local', 'development'),
(53, '2024-06-12 15:45:51.496843', 'local', 'development'),
(54, '2024-06-12 15:51:56.55792', 'local', 'development'),
(55, '2024-06-12 15:53:47.914103', 'local', 'development'),
(56, '2024-06-12 15:55:05.431681', 'local', 'development'),
(57, '2024-06-12 15:56:12.165568', 'local', 'development'),
(58, '2024-06-12 15:56:48.454338', 'local', 'development'),
(59, '2024-06-12 15:58:18.119951', 'local', 'development'),
(60, '2024-06-12 16:01:51.272254', 'local', 'development'),
(61, '2024-06-12 16:03:39.233975', 'local', 'development'),
(62, '2024-06-12 16:06:15.826342', 'local', 'development'),
(63, '2024-06-12 17:21:14.890474', 'local', 'development'),
(64, '2024-06-12 17:48:08.761443', 'local', 'development'),
(65, '2024-06-12 17:56:42.703363', 'local', 'development'),
(66, '2024-06-12 18:02:39.647626', 'local', 'development'),
(67, '2024-06-12 18:03:29.908906', 'local', 'development'),
(68, '2024-06-12 18:05:39.671652', 'local', 'development'),
(69, '2024-06-12 21:38:07.857295', 'local', 'development'),
(70, '2024-06-12 22:18:27.389833', 'local', 'development'),
(71, '2024-06-12 22:22:16.585743', 'local', 'development'),
(72, '2024-06-12 22:25:18.72171', 'local', 'development');

INSERT INTO "public"."refreshtokens" ("userid", "token", "time") VALUES
('cdb4bc0c-a2db-42', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA4OTE1MzIsInVzZXJJZCI6ImNkYjRiYzBjLWEyZGItNDIifQ.vzrZOzTeyZSkQntA1gm4OTKKdTaCPj--E3lrqrK8nlQ', '2024-06-12 17:25:32.27828');

INSERT INTO "public"."twofatemp" ("userid", "token", "time") VALUES
('cdb4bc0c-a2db-42', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjozMDQ3NTYsImV4cCI6MTcxODIxNDAxOX0.6_pPl1JsvCEw0bf4_55BkNX48NyDLqObqXZ5-VRy6ig', '2024-06-12 17:20:19.40177');

INSERT INTO "public"."users" ("userid", "displayid", "email", "phone", "name", "surname", "joindate", "verify", "twofa") VALUES
('cdb4bc0c-a2db-42', 636966, 'admin@admin.com', '+79523235502', 'Alex', 'Antipev', '2024-06-08 11:19:06.948374', 't', 'f');

ALTER TABLE "public"."refreshtokens" ADD FOREIGN KEY ("userid") REFERENCES "public"."users"("userid");
ALTER TABLE "public"."templogincode" ADD FOREIGN KEY ("userid") REFERENCES "public"."users"("userid");
ALTER TABLE "public"."twofatemp" ADD FOREIGN KEY ("userid") REFERENCES "public"."users"("userid");


-- Indices
CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);
CREATE UNIQUE INDEX users_phone_key ON public.users USING btree (phone);
