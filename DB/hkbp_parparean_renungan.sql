-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 14, 2024 at 08:03 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `hkbp_parparean_renungan`
--

-- --------------------------------------------------------

--
-- Table structure for table `renungans`
--

CREATE TABLE `renungans` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `judul` varchar(255) DEFAULT NULL,
  `isi_renungan` varchar(255) DEFAULT NULL,
  `ayat_renungan` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `tanggal` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `renungans`
--

INSERT INTO `renungans` (`id`, `created_at`, `updated_at`, `deleted_at`, `judul`, `isi_renungan`, `ayat_renungan`, `status`, `tanggal`) VALUES
(1, '2024-05-13 13:08:45', '2024-05-13 13:24:23', NULL, 'Allianz Arena', 'ayat suci', '12 timotius 3:2', 'Tampilkan', '2024-05-24T20:18'),
(2, '2024-05-13 13:12:43', '2024-05-13 13:24:33', NULL, 'contoh baru paling terbaru', 'ayat alkitab', '12 Yakobus 3:2', 'Tampilkan', '2024-05-15T17:16');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `renungans`
--
ALTER TABLE `renungans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_renungans_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `renungans`
--
ALTER TABLE `renungans`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
