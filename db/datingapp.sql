-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 30 Jul 2023 pada 20.35
-- Versi server: 5.7.33
-- Versi PHP: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `datingapp`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `matchs`
--

CREATE TABLE `matchs` (
  `id` int(11) NOT NULL,
  `profil_id` int(11) NOT NULL,
  `profil_tujuan` int(11) NOT NULL,
  `profil_match` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT 'MATCH',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `matchs`
--

INSERT INTO `matchs` (`id`, `profil_id`, `profil_tujuan`, `profil_match`, `status`, `created_at`, `updated_at`) VALUES
(5, 1, 2, 0, '', '2023-07-30 19:48:58', '2023-07-30 20:10:55');

-- --------------------------------------------------------

--
-- Struktur dari tabel `persentase`
--

CREATE TABLE `persentase` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `like` int(11) NOT NULL,
  `dislike` int(11) NOT NULL,
  `matchs` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `persentase`
--

INSERT INTO `persentase` (`id`, `user_id`, `like`, `dislike`, `matchs`, `created_at`, `updated_at`) VALUES
(1, 7, 1, 2, 0, '2023-07-30 09:01:19', '2023-07-30 19:58:38'),
(2, 8, 0, 0, 0, '2023-07-30 19:14:17', '2023-07-30 19:14:17'),
(3, 9, 0, 0, 0, '2023-07-30 19:14:57', '2023-07-30 19:14:57'),
(4, 10, 0, 0, 0, '2023-07-30 19:15:24', '2023-07-30 19:15:24');

-- --------------------------------------------------------

--
-- Struktur dari tabel `photo`
--

CREATE TABLE `photo` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `image` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `photo`
--

INSERT INTO `photo` (`id`, `user_id`, `image`, `created_at`, `updated_at`) VALUES
(3, 7, 'saya.jpg', '2023-07-30 09:01:19', '2023-07-30 16:10:14'),
(4, 8, '', '2023-07-30 19:14:17', '2023-07-30 19:14:17'),
(5, 9, '', '2023-07-30 19:14:57', '2023-07-30 19:14:57'),
(6, 10, '', '2023-07-30 19:15:24', '2023-07-30 19:15:24');

-- --------------------------------------------------------

--
-- Struktur dari tabel `premium`
--

CREATE TABLE `premium` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `expired` varchar(20) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `premium`
--

INSERT INTO `premium` (`id`, `user_id`, `expired`, `created_at`, `updated_at`) VALUES
(4, 7, '2023-09-30', '2023-07-30 17:34:57', '2023-07-30 17:34:57');

-- --------------------------------------------------------

--
-- Struktur dari tabel `profil`
--

CREATE TABLE `profil` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `gender` int(11) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `birthdate` varchar(50) DEFAULT NULL,
  `birth_info` varchar(50) DEFAULT NULL,
  `bio` text,
  `lokasi` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `profil`
--

INSERT INTO `profil` (`id`, `user_id`, `gender`, `age`, `birthdate`, `birth_info`, `bio`, `lokasi`, `created_at`, `updated_at`) VALUES
(1, 7, 1, 26, '1997-03-21', 'Rembang', 'Ini saya bingung', 'Rembang', '2023-07-30 09:01:19', '2023-07-30 19:15:51'),
(2, 8, 0, 0, '', '', '', 'Rembang', '2023-07-30 19:14:17', '2023-07-30 19:15:55'),
(3, 9, 0, 0, '', '', '', 'Semarang', '2023-07-30 19:14:57', '2023-07-30 19:15:59'),
(4, 10, 0, 0, '', '', 'Y kudus', 'kudus', '2023-07-30 19:15:24', '2023-07-30 19:52:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `fullname` varchar(60) NOT NULL,
  `email` varchar(50) NOT NULL,
  `notelp` varchar(20) NOT NULL,
  `password` varchar(100) NOT NULL,
  `token` text,
  `status` varchar(10) NOT NULL DEFAULT 'AKTIF',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id`, `username`, `fullname`, `email`, `notelp`, `password`, `token`, `status`, `created_at`, `updated_at`) VALUES
(7, 'hadi12', 'M.Nur Hadi', 'saya@mnurhad.com', '0882005424842', '$2a$10$lUSI3S8Pl3XGxVKm18iDCeV9cExWmHbMNtfHydQetPch3w1Eoizeu', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDbGllbnRLZXkiOiJoYWRpMTIiLCJDbGllbnRTZWNyZXQiOiIkMmEkMTAkbFVTSTNTOFBsM1hHeFZLbTE4aURDZVY5Y0V4V21IYk1OdGZIeWRRZXRQY2gzdzFFb2l6ZXUiLCJleHAiOjE2OTA3NDg0Njd9.XSFcuOITyA0aK3Tv71jwoj9F-c1vVmJgfHMaCoTeJvU', 'AKTIF', '2023-07-30 09:01:19', '2023-07-30 20:06:07'),
(8, 'sonia12', 'Sonia Rahayu', 'saya@sonia.com', '088200544842', '$2a$10$LBlbgQhM9Qma10KUCiJBFO3NLvgMj6NVxafbg6gpZKfulaBKLr2m2', '', 'AKTIF', '2023-07-30 19:14:17', '2023-07-30 19:14:17'),
(9, 'bela12', 'Bela Putri', 'saya@bela.com', '0882005448422', '$2a$10$A3OtybmuG4g7OnRyYNxOGOVb7/SbbVglPds3EO7FOvg5QWzQ422UK', '', 'AKTIF', '2023-07-30 19:14:57', '2023-07-30 19:14:57'),
(10, 'udi12', 'Udi Setiawan', 'saya@udi.com', '08820054484221', '$2a$10$fQeCq7PGnGqOCrPvgEVuAOnzXIhZz23323vQDAufee/aejJMl1EH6', '', 'AKTIF', '2023-07-30 19:15:24', '2023-07-30 19:15:24');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `matchs`
--
ALTER TABLE `matchs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `persentase`
--
ALTER TABLE `persentase`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `photo`
--
ALTER TABLE `photo`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `premium`
--
ALTER TABLE `premium`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `profil`
--
ALTER TABLE `profil`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `matchs`
--
ALTER TABLE `matchs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `persentase`
--
ALTER TABLE `persentase`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `photo`
--
ALTER TABLE `photo`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `premium`
--
ALTER TABLE `premium`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `profil`
--
ALTER TABLE `profil`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `persentase`
--
ALTER TABLE `persentase`
  ADD CONSTRAINT `FK1_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Ketidakleluasaan untuk tabel `photo`
--
ALTER TABLE `photo`
  ADD CONSTRAINT `user_key` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Ketidakleluasaan untuk tabel `premium`
--
ALTER TABLE `premium`
  ADD CONSTRAINT `FK__users` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Ketidakleluasaan untuk tabel `profil`
--
ALTER TABLE `profil`
  ADD CONSTRAINT `user_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
