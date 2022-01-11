-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 11, 2022 at 07:18 AM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 8.0.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_foodmarket`
--

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

CREATE TABLE `cart` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `status` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL,
  `category_name` varchar(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `category`
--

INSERT INTO `category` (`id`, `category_name`) VALUES
(1, 'nasi'),
(2, 'kopi'),
(3, 'dessert'),
(4, 'mie'),
(5, 'sayuran'),
(6, 'es');

-- --------------------------------------------------------

--
-- Table structure for table `favourite`
--

CREATE TABLE `favourite` (
  `id` int(11) NOT NULL,
  `market_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `market`
--

CREATE TABLE `market` (
  `id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `market_name` varchar(128) NOT NULL,
  `market_img` varchar(128) NOT NULL,
  `market_location` varchar(256) NOT NULL,
  `market_rating` varchar(10) NOT NULL,
  `market_product_name` varchar(128) NOT NULL,
  `market_estimate` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `market`
--

INSERT INTO `market` (`id`, `category_id`, `market_name`, `market_img`, `market_location`, `market_rating`, `market_product_name`, `market_estimate`) VALUES
(1, 1, 'Nasi Uduk Bang Bejo', 'nasiuduk.png', 'Tangerang', '4.9', 'Nasi uduk, pecel lele, lalapan, soto ayam, soto babat', '25 menit'),
(2, 1, 'Nasi Daging Bu Joko', 'nasidaging.png', 'Tangerang', '4.8', 'Nasi daging kambing, daging sapi, sambal balado, dan lalapan', '20 menit'),
(3, 1, 'Nasi Jagal Poris', 'nasijagal.png', 'Tangerang', '5', 'Nasi jagal sambal kentang, jagal sambal ijo, jagal kambing sapi', '18 menit'),
(4, 2, 'Kopi Harapan Mantan', 'kopihaman.png', 'Tangerang', '4.8.', 'Kopi Harapan Mantan, Kopi Matcha, Kopi Coklat Dll', '26 menit'),
(5, 2, 'Kopi Milo Cincau Hati', 'kopimilo.png', 'Tangerang', '4.9', 'Kopi Cincau, Kopi Milo, Kopi dll', '15 menit'),
(6, 2, 'Kopi Syur Abis', 'kopisyur.png', 'Tangerang', '4.7', 'Kopi Hitam Manis, Kopi Susu, Dll', '20 menit'),
(7, 3, 'Cheese Cake Jojo', 'dessertcake.png', 'Tangerang', '5', 'Cheese Cake, Oreo Cheese Cake, Cheese Ice Cream Oreo', '18 min'),
(8, 3, 'Waffle Milo Ice Cream', 'dessertwaffle.png', 'Tangerang', '4.9', 'Ice Cream Chocolate Waffle, Milo Ice Oreo Cream, Dll', '20 menit'),
(9, 3, 'Crepes Mantap Mantul', 'dessertcrepes.png', 'Tangerang', '4.6', 'Crepes Coklat, Crepes Daging, Crepes Keju Dll', '12 menit'),
(10, 4, 'Mie Ayam Yamin Rumput Laut', 'mieayam.png', 'Tangerang', '4.9', 'Mie Ayam Rumput Laut, Mie Ayam Yamin, Mie Campur Dll', '15 menit'),
(11, 4, 'Mie Jagal', 'miejagal.png', 'Tangerang', '5', 'Mie Daging Ayam, Mie Daging Sapi Asap, Mie Ayam Jagal', '14 menit'),
(12, 4, 'Mie Kocok', 'miekocok.png', 'Tangerang', '4.8', 'Mie Kocok Ayam, Mie Kocok Sapi, Mie Kocok Sambal Dll', '13 menit'),
(13, 5, 'Salad Tumis Saos Tiram', 'saladsaostiram.png', 'Tangerang', '4.9', 'Salad Tumis, Salad Mayonaise, Salad Saos Tiram dll', '17 menit'),
(14, 5, 'Sop Salad Kangkung', 'sopsalad.png', 'Tangerang', '4.9', 'Sop Kangkung, Salad Kangkung dll', '12 menit'),
(15, 5, 'Raja Tumis Sayuran', 'sayurraja.png', 'Tangerang', '5', 'Tumis Kangkung, Tumis Kol, Tumis Brokoli dll', '20 menit'),
(16, 6, 'Es Capuccino Cincau', 'escapcin.png', 'Tangerang', '4.7', 'Capuccino Cincau Es, Capucinno Cincau dll', '19 menit'),
(17, 6, 'Es Buah Bang Jojo', 'esbuah.png', 'Tangerang', '4.9', 'Es Buah Komplit, Es buah Mangga Melon Dll', '20 menit'),
(18, 6, 'Es Coklat Mantap', 'escoklat.png', 'Tangerang', '4.5', 'Es Coklat Jumbo, Es Coklat Komplit dll', '21 menit');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `market_id` int(11) NOT NULL,
  `product_name` varchar(128) NOT NULL,
  `product_price` varchar(128) NOT NULL,
  `product_img` varchar(128) NOT NULL,
  `product_stock` int(11) NOT NULL,
  `product_status` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `market_id`, `product_name`, `product_price`, `product_img`, `product_stock`, `product_status`) VALUES
(1, 1, 'Nasi Uduk Ayam Paha & Dada', '18000', 'nasiuduk/nasiuduk1.png', 5, 'ready'),
(2, 1, 'Nasi Uduk Goreng Sambal Tomat', '17000', 'nasiuduk/nasiuduk2.png', 5, 'ready'),
(3, 1, 'Nasi Uduk Ati Ampela', '14000', 'nasiuduk/nasiuduk3.png', 5, 'ready'),
(4, 2, 'Nasi Daging Sapi', '25000', 'nasidaging/nasidaging1.png', 5, 'ready'),
(5, 2, 'Nasi Daging Ayam Sambal', '20000', 'nasidaging/nasidaging2.png', 5, 'ready'),
(6, 2, 'Nasi Daging Sapi Ayam Sambal Komplit', '30000', 'nasidaging/nasidaging3.png', 5, 'ready'),
(7, 3, 'Nasi Jagal Kambing', '15000', 'nasijagal/nasijagal1.png', 5, 'ready'),
(8, 3, 'Nasi Jagal Sapi', '18000', 'nasijagal/nasijagal2.png', 5, 'ready'),
(9, 3, 'Nasi Jagal Kambing Sapi Komplit', '30000', 'nasijagal/nasijagal3.png', 5, 'ready'),
(10, 4, 'Kopi Susu Mantan', '20000', 'kopihaman/kopihaman1.png', 5, 'ready'),
(11, 4, 'Kopi Coklat Mantan', '22000', 'kopihaman/kopihaman2.png', 5, 'ready'),
(12, 4, 'Kopi Matcha Mantan', '23000', 'kopihaman/kopihaman3.png', 5, 'ready'),
(13, 5, 'Kopi Milo Cincau', '18000', 'kopimilo/kopimilo1.png', 5, 'ready'),
(14, 5, 'Kopi Milo', '15000', 'kopimilo/kopimilo2.png', 5, 'ready'),
(15, 5, 'Kopi Cincau', '15000', 'kopimilo/kopimilo3.png', 5, 'ready'),
(16, 6, 'Kopi Hitam Mantap', '14000', 'kopisyur/kopisyur1.png', 5, 'ready'),
(17, 6, 'Kopi Syur Spesial', '18000', 'kopisyur/kopisyur2.png', 5, 'ready'),
(18, 6, 'Kopi Syur Mantap', '20000', 'kopisyur/kopisyur3.png', 5, 'ready'),
(19, 7, 'Cheese Cake Coklat', '20000', 'dessertcake/cake1.png', 5, 'ready'),
(20, 7, 'Cheese Cake Keju', '22000', 'dessertcake/cake2.png', 5, 'ready'),
(21, 7, 'Cheese Cake Oreo', '25000', 'dessertcake/cake3.png', 5, 'ready'),
(22, 8, 'Waffle Milo', '30000', 'dessertwaffle/waffle1.png', 5, 'ready'),
(23, 8, 'Waffle Oreo', '30000', 'dessertwaffle/waffle2.png', 5, 'ready'),
(24, 8, 'Waffle Milo Oreo', '38000', 'dessertwaffle/waffle3.png', 5, 'ready'),
(25, 9, 'Crepes Coklat Komplit', '18000', 'dessertcrepes/crepes1.png', 5, 'ready'),
(26, 9, 'Crepes Daging Sapi', '28000', 'dessertcrepes/crepes2.png', 5, 'ready'),
(27, 9, 'Crepes Piscok', '20000', 'dessertcrepes/crepes3.png', 5, 'ready'),
(28, 10, 'Mie Ayam Biasa', '15000', 'mieayam/mieayam1.png', 5, 'ready'),
(29, 10, 'Mie Ayam Rumput Laut', '25000', 'mieayam/mieayam2.png', 5, 'ready'),
(30, 10, 'Mie Ayam Rumput Laut Yamin', '28000', 'mieayam/mieayam3.png', 5, 'ready'),
(31, 11, 'Mie Jagal Ayam', '20000', 'miejagal/miejagal1.png', 5, 'ready'),
(32, 11, 'Mie Jagal Kambing', '25000', 'miejagal/miejagal2.png', 5, 'ready'),
(33, 11, 'Mie Jagal sapi', '28000', 'miejagal/miejagal3.png', 5, 'ready'),
(34, 12, 'Mie Kocok Ayam', '22000', 'miekocok/miekocok1.png', 5, 'ready'),
(35, 12, 'Mie Kocok Kambing', '28000', 'miekocok/miekocok2.png', 5, 'ready'),
(36, 12, 'Mie Kocok Jamur & Sapi', '32000', 'miekocok/miekocok3.png', 5, 'ready'),
(37, 13, 'Salad Mayonaise', '15000', 'saladsaostiram/salad1.png', 5, 'ready'),
(38, 13, 'Salad Saos Tiram', '16000', 'saladsaostiram/salad2.png', 5, 'ready'),
(39, 13, 'Salad Buah Sayur', '12000', 'saladsaostiram/salad3.png', 5, 'ready'),
(40, 14, 'Sop Kol Royko', '12000', 'sopsalad/sop1.png', 5, 'ready'),
(41, 14, 'Sop Kol Masako', '12000', 'sopsalad/sop2.png', 5, 'ready'),
(42, 14, 'Sop Kol Kangkung Saos Tiram', '15000', 'sopsalad/sop3.png', 5, 'ready'),
(43, 15, 'Tumis Kangkung Kecap Asin', '20000', 'sayurraja/tumis1.png', 5, 'ready'),
(44, 15, 'Tumis Kol Lalap Saos Tiram', '22000', 'sayurraja/tumis2.png', 5, 'ready'),
(45, 15, 'Tumis Sawi Kangkung Bumbu Ayam', '25000', 'sayurraja/tumis3.png', 5, 'ready'),
(46, 16, 'Capucinno Cincau Biasa', '7000', 'escapcin/capcin1.png', 5, 'ready'),
(47, 16, 'Capucinno Cincau Coklat', '12000', 'escapcin/capcin2.png', 5, 'ready'),
(48, 16, 'Capucinno Cincau Oreo', '15000', 'escapcin/capcin3.png', 5, 'ready'),
(49, 17, 'Es Buah Biasa', '10000', 'esbuah/esbuah1.png', 5, 'ready'),
(50, 17, 'Es Buah Komplit', '15000', 'esbuah/esbuah2.png', 5, 'ready'),
(51, 17, 'Es Buah Melon Mangga', '12000', 'esbuah/esbuah3.png', 5, 'ready'),
(52, 18, 'Es Coklat Milo', '15000', 'escoklat/escoklat1.png', 5, 'ready'),
(53, 18, 'Es Coklat Oreo', '16000', 'escoklat/escoklat2.png', 5, 'ready'),
(54, 18, 'Es Coklat Komplit', '19000', 'escoklat/escoklat3.png', 5, 'ready');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `fullname` varchar(128) NOT NULL,
  `username` varchar(128) NOT NULL,
  `password` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `user_img` varchar(128) NOT NULL,
  `phone_number` varchar(13) NOT NULL,
  `role_id` int(11) NOT NULL,
  `user_status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `fullname`, `username`, `password`, `email`, `user_img`, `phone_number`, `role_id`, `user_status`) VALUES
(1, 'ical aja jadinya', 'ical852', '123456789', 'shalahuddin@gmail.com', 'ical.jpg', '123', 2, 1),
(5, 'Ical852', 'Ical_', '12345678', 'shalahuddin@gmail.com', 'default.png', '089643936991', 2, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `favourite`
--
ALTER TABLE `favourite`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `market`
--
ALTER TABLE `market`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cart`
--
ALTER TABLE `cart`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `favourite`
--
ALTER TABLE `favourite`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `market`
--
ALTER TABLE `market`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=55;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
