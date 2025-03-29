CREATE TABLE `category` (
  `id` bigint NOT NULL,
  `category_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `rope_length` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_czech_ci;

INSERT INTO `category` (`id`, `category_key`, `label`, `rope_length`) VALUES
(1, 'KAT_I', 'Žáci', 4.5),
(2, 'KAT_II', 'Dorostenci', 4.5),
(3, 'KAT_III', 'Muži', 8),
(4, 'KAT_IV', 'Masters', 8),
(5, 'KAT_V', 'Ženy a dorostenky', 4.5),
(6, 'KAT_VI', 'Žákyně', 4.5);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UK_mdbxetaccaq7s3bhia24htew8` (`category_key`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;


CREATE TABLE `competition` (
  `id` bigint NOT NULL,
  `competition_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `date` date DEFAULT NULL,
  `jugde` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `place` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `sensor_installation` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `starter` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_czech_ci;

ALTER TABLE `competition`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UKn4fi92r1c43e465b20rgam8cp` (`name`,`competition_name`,`date`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `competition`
--
ALTER TABLE `competition`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=126;
COMMIT;

CREATE TABLE `rope_climber` (
  `id` bigint NOT NULL,
  `first_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `last_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `year_of_birth` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_czech_ci;

ALTER TABLE `rope_climber`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `rope_climber`
--
ALTER TABLE `rope_climber`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1464;
COMMIT;


CREATE TABLE `participation` (
  `id` bigint NOT NULL,
  `organization` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  `competition_id` bigint DEFAULT NULL,
  `rope_climber_id` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_czech_ci;

ALTER TABLE `participation`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FKrucg5lqvvhhkemka476ga5b6g` (`category_id`),
  ADD KEY `FKpqea0a5s7bvvfi0mpunegcpwy` (`competition_id`),
  ADD KEY `FK83aykwqmbi925yu24w6kfybvl` (`rope_climber_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `participation`
--
ALTER TABLE `participation`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5383;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `participation`
--
ALTER TABLE `participation`
  ADD CONSTRAINT `FK83aykwqmbi925yu24w6kfybvl` FOREIGN KEY (`rope_climber_id`) REFERENCES `rope_climber` (`id`),
  ADD CONSTRAINT `FKpqea0a5s7bvvfi0mpunegcpwy` FOREIGN KEY (`competition_id`) REFERENCES `competition` (`id`),
  ADD CONSTRAINT `FKrucg5lqvvhhkemka476ga5b6g` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);
COMMIT;

CREATE TABLE `time` (
  `id` bigint NOT NULL,
  `round` int DEFAULT NULL,
  `time` double DEFAULT NULL,
  `participation_id` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_czech_ci;


ALTER TABLE `time`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FK3dyq1l47x20w22b3s25b0qy5l` (`participation_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `time`
--
ALTER TABLE `time`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18503;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `time`
--
ALTER TABLE `time`
  ADD CONSTRAINT `FK3dyq1l47x20w22b3s25b0qy5l` FOREIGN KEY (`participation_id`) REFERENCES `participation` (`id`);
COMMIT;

