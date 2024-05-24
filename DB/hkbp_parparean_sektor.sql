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
-- Database: `hkbp_parparean_sektor`
--

-- --------------------------------------------------------

--
-- Table structure for table `sektors`
--

CREATE TABLE `sektors` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nama_sektor` varchar(255) DEFAULT NULL,
  `deskripsi` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `kontak` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `sektors`
--

INSERT INTO `sektors` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_sektor`, `deskripsi`, `alamat`, `kontak`) VALUES
(1, '2024-05-13 14:17:08', '2024-05-15 00:05:09', NULL, 'Sektor 1', 'ini adalah sektor satu', 'Silangkitang', '111111'),
(2, '2024-05-13 14:21:54', '2024-05-13 14:21:54', '2024-05-13 14:23:03', 'Sektor 2', '', 'balige', '082454658164'),
(3, '2024-05-13 14:23:17', '2024-05-13 14:23:17', NULL, 'Sektor 2', 'Hot Coffee Latte Terbaik', 'sipiongot', '874635768568');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `sektors`
--
ALTER TABLE `sektors`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_sektors_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `sektors`
--
ALTER TABLE `sektors`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
