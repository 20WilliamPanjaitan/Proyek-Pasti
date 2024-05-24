-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 14, 2024 at 08:02 PM
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
-- Database: `hkbp_parparean_parhalado`
--

-- --------------------------------------------------------

--
-- Table structure for table `parhalados`
--

CREATE TABLE `parhalados` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `tanggal_lahir` varchar(255) DEFAULT NULL,
  `jabatan` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `parhalados`
--

INSERT INTO `parhalados` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama`, `tanggal_lahir`, `jabatan`, `alamat`) VALUES
(1, '2024-05-15 00:14:14', '2024-05-15 00:39:51', '2024-05-15 00:40:53', 'Regen Sibagariang', '1970-07-25', 'Pareses', 'Sitampulak'),
(3, '2024-05-15 00:26:06', '2024-05-15 00:41:13', NULL, 'Leonal Sitompul', '2004-10-20', 'Sintua', 'Laguboti');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `parhalados`
--
ALTER TABLE `parhalados`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_parhalados_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `parhalados`
--
ALTER TABLE `parhalados`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
