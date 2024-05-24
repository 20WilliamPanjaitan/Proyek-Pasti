-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 14, 2024 at 08:01 PM
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
-- Database: `hkbp_parparean_acara`
--

-- --------------------------------------------------------

--
-- Table structure for table `acaras`
--

CREATE TABLE `acaras` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nama_acara` varchar(255) DEFAULT NULL,
  `lokasi_acara` varchar(255) DEFAULT NULL,
  `jenis_acara` varchar(255) DEFAULT NULL,
  `waktu_pelaksanaan` varchar(255) DEFAULT NULL,
  `keterangan` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `acaras`
--

INSERT INTO `acaras` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_acara`, `lokasi_acara`, `jenis_acara`, `waktu_pelaksanaan`, `keterangan`) VALUES
(1, '2024-05-13 08:14:21', '2024-05-13 08:14:21', NULL, 'Contoh Acara', 'Balige', 'Ibadah', '2024-05-24T04:03', 'ini adalah contoh acara'),
(2, '2024-05-13 09:48:40', '2024-05-13 10:50:21', NULL, 'Parmingguan', 'WEF', 'Ibadah', '2024-05-31T10:41', 'ABC'),
(3, '2024-05-13 09:49:57', '2024-05-13 09:49:57', '2024-05-13 10:23:18', 'Minggu Palmarum', 'Balige', 'Ibadah', '', 'WRkyjrthe'),
(4, '2024-05-13 09:50:28', '2024-05-13 09:50:28', NULL, 'Minggu Palmarum ke 20', 'oisd', 'Ibadah', '2024-05-23T13:53', 'ystzhnty');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `acaras`
--
ALTER TABLE `acaras`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_acaras_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `acaras`
--
ALTER TABLE `acaras`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
