-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 31, 2023 at 06:32 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_crud`
--

-- --------------------------------------------------------

--
-- Table structure for table `dokter`
--

CREATE TABLE `dokter` (
  `id` int(11) NOT NULL,
  `nid` int(9) NOT NULL,
  `nama` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `keahlian` varchar(255) NOT NULL,
  `no_hp` varchar(13) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `dokter`
--

INSERT INTO `dokter` (`id`, `nid`, `nama`, `keahlian`, `no_hp`) VALUES
(1, 114220001, 'Ukasyah', 'Gigi', '0895379114998'),
(2, 114220002, 'Fachriza Farhan', 'Ahli Jantung', '0895379114998'),
(3, 114220003, 'Gaizka', 'Umum', '0876987654325');

-- --------------------------------------------------------

--
-- Table structure for table `hari`
--

CREATE TABLE `hari` (
  `id` int(11) NOT NULL,
  `hari` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `hari`
--

INSERT INTO `hari` (`id`, `hari`) VALUES
(1, 'Senin'),
(2, 'Selasa'),
(3, 'Rabu'),
(4, 'Kamis'),
(5, 'Jum\'at');

-- --------------------------------------------------------

--
-- Table structure for table `jadwal_dokter`
--

CREATE TABLE `jadwal_dokter` (
  `id` int(11) NOT NULL,
  `id_dokter` int(11) NOT NULL,
  `id_hari` int(11) NOT NULL,
  `id_jam` int(11) NOT NULL,
  `id_ruangan` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `jadwal_dokter`
--

INSERT INTO `jadwal_dokter` (`id`, `id_dokter`, `id_hari`, `id_jam`, `id_ruangan`) VALUES
(1, 1, 1, 1, 1),
(2, 2, 1, 2, 2),
(3, 2, 2, 2, 2),
(4, 1, 2, 1, 1),
(5, 3, 3, 1, 3),
(6, 3, 3, 3, 3),
(7, 2, 4, 3, 2);

-- --------------------------------------------------------

--
-- Table structure for table `jam`
--

CREATE TABLE `jam` (
  `id` int(11) NOT NULL,
  `jam` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `jam`
--

INSERT INTO `jam` (`id`, `jam`) VALUES
(1, '08:00'),
(2, '08:30'),
(3, '09.00'),
(4, '09.30'),
(5, '10.00'),
(6, '10.30'),
(7, '11.00'),
(8, '13.00'),
(9, '13.30'),
(10, '14.00'),
(11, '14.30'),
(12, '15.00');

-- --------------------------------------------------------

--
-- Table structure for table `pasien`
--

CREATE TABLE `pasien` (
  `id` int(10) NOT NULL,
  `nama_lengkap` varchar(300) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `jenis_kelamin` varchar(10) NOT NULL,
  `tempat_lahir` varchar(300) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `alamat` text NOT NULL,
  `no_hp` varchar(15) NOT NULL,
  `id_jadwal` int(11) DEFAULT NULL,
  `tgl_reservasi` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pasien`
--

INSERT INTO `pasien` (`id`, `nama_lengkap`, `nik`, `jenis_kelamin`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `no_hp`, `id_jadwal`, `tgl_reservasi`) VALUES
(55, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(56, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(57, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(58, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(59, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(60, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(61, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(62, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(63, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(64, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-23', 'Jl. Indonesia', '085156764522', 1, '2023-06-05 00:00:00'),
(65, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-03', 'Jl. Indonesia', '085156764522', 4, '2023-06-05 00:00:00'),
(66, 'Dzul', '1234567891234567', 'Pria', 'Bandung', '2023-05-11', 'Jl. Indonesia', '085156764522', 3, '2023-06-05 00:00:00'),
(67, 'Dzul', '3019735829777', 'Pria', 'Bandung', '2023-05-05', 'Jl. Indonesia', '085156764522', 3, '2023-06-05 00:00:00'),
(68, 'Dzul', '3019735829777', 'Pria', 'Bandung', '2023-05-05', 'Jl. Indonesia', '085156764522', 3, '2023-06-05 00:00:00'),
(69, 'Dzul', '3019735829777', 'Pria', 'Bandung', '2023-05-05', 'Jl. Indonesia', '085156764522', 3, '2023-06-05 00:00:00'),
(70, 'Dzul', '3019735829777', 'Pria', 'Bandung', '2023-05-05', 'Jl. Indonesia', '085156764522', 3, '2023-06-05 00:00:00'),
(91, 'Dzul', '545665161', 'Pria', 'Bdg', '2001-08-04', 'Jl. Indonesia', '085156764522', 1, '2023-05-31 22:07:53'),
(92, 'jsadj', '515151', 'sad', 'padnag', '2001-08-04', 'lkjaskdj', '49846546', 4, '2023-05-31 22:26:19'),
(93, 'jkasbndkj', '5561', 'kjasd', 'ajsdnkajsdn', '6226-08-04', 'lkajnsdnl', '65465465', 1, '2023-05-31 22:37:28'),
(94, 'Dzul', '1234567891234567', 'sad', 'Bekasi', '4555-05-05', 'Jl. Indonesia', '085156764522', 1, '2023-05-31 22:52:20'),
(95, 'Dzul', '3019735829777', 'sad', 'Bandung', '5455-04-05', 'Jl. Indonesia', '085156764522', 1, '2023-05-31 22:57:40'),
(96, 'Dzul', '45545', 'kjnadfkj', 'lkansl', '5665-04-05', 'Jl. Indonesia', '085156764522', 1, '2023-05-31 23:07:54'),
(97, 'asdbjasd', '65464', '.nasjkf', 'lkansd', '6846-04-04', 'knasldkn', '644684', 1, '2023-05-31 23:09:22');

-- --------------------------------------------------------

--
-- Table structure for table `ruangan`
--

CREATE TABLE `ruangan` (
  `kode_ruangan` int(11) NOT NULL,
  `nama_ruangan` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ruangan`
--

INSERT INTO `ruangan` (`kode_ruangan`, `nama_ruangan`) VALUES
(1, 'Belalang'),
(2, 'Gajah'),
(3, 'Capung');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dokter`
--
ALTER TABLE `dokter`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `hari`
--
ALTER TABLE `hari`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jadwal_dokter`
--
ALTER TABLE `jadwal_dokter`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_dokter` (`id_dokter`),
  ADD KEY `id_hari` (`id_hari`),
  ADD KEY `id_jam` (`id_jam`),
  ADD KEY `id_ruangan` (`id_ruangan`);

--
-- Indexes for table `jam`
--
ALTER TABLE `jam`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pasien`
--
ALTER TABLE `pasien`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_jadwal` (`id_jadwal`);

--
-- Indexes for table `ruangan`
--
ALTER TABLE `ruangan`
  ADD PRIMARY KEY (`kode_ruangan`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `dokter`
--
ALTER TABLE `dokter`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `hari`
--
ALTER TABLE `hari`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `jadwal_dokter`
--
ALTER TABLE `jadwal_dokter`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `jam`
--
ALTER TABLE `jam`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `pasien`
--
ALTER TABLE `pasien`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=98;

--
-- AUTO_INCREMENT for table `ruangan`
--
ALTER TABLE `ruangan`
  MODIFY `kode_ruangan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `jadwal_dokter`
--
ALTER TABLE `jadwal_dokter`
  ADD CONSTRAINT `jadwal_dokter_ibfk_1` FOREIGN KEY (`id_dokter`) REFERENCES `dokter` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `jadwal_dokter_ibfk_2` FOREIGN KEY (`id_hari`) REFERENCES `hari` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `jadwal_dokter_ibfk_3` FOREIGN KEY (`id_jam`) REFERENCES `jam` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `jadwal_dokter_ibfk_4` FOREIGN KEY (`id_ruangan`) REFERENCES `ruangan` (`kode_ruangan`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `pasien`
--
ALTER TABLE `pasien`
  ADD CONSTRAINT `pasien_ibfk_1` FOREIGN KEY (`id_jadwal`) REFERENCES `jadwal_dokter` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
