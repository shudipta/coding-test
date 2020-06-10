-- -------------------------------------------------------------
-- TablePlus 3.1.2(296)
--
-- https://tableplus.com/
--
-- Database: ecommerce
-- Generation Time: 2020-06-09 20:22:20.7320
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(15) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `company_name` varchar(40) NOT NULL,
  `contact_name` varchar(30) DEFAULT NULL,
  `contact_title` varchar(30) DEFAULT NULL,
  `address` varchar(60) DEFAULT NULL,
  `city` varchar(15) DEFAULT NULL,
  `region` varchar(15) DEFAULT NULL,
  `postal_code` varchar(10) DEFAULT NULL,
  `country` varchar(15) DEFAULT NULL,
  `phone` varchar(24) DEFAULT NULL,
  `fax` varchar(24) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `order_details`;
CREATE TABLE `order_details` (
  `order_id` smallint(6) NOT NULL,
  `product_id` smallint(6) NOT NULL,
  `unit_price` double NOT NULL,
  `quantity` smallint(6) NOT NULL,
  `discount` double NOT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `order_details_fk_product_id` (`product_id`),
  CONSTRAINT `order_details_fk_order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_details_fk_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `customer_id` smallint(6) DEFAULT NULL,
  `order_date` date DEFAULT NULL,
  `required_date` date DEFAULT NULL,
  `shipped_date` date DEFAULT NULL,
  `freight` double DEFAULT NULL,
  `ship_name` varchar(40) DEFAULT NULL,
  `ship_address` varchar(60) DEFAULT NULL,
  `ship_city` varchar(15) DEFAULT NULL,
  `ship_region` varchar(15) DEFAULT NULL,
  `ship_postal_code` varchar(10) DEFAULT NULL,
  `ship_country` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_fk_customer_id` (`customer_id`),
  CONSTRAINT `orders_fk_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(40) NOT NULL,
  `supplier_id` smallint(6) DEFAULT NULL,
  `category_id` smallint(6) DEFAULT NULL,
  `quantity_per_unit` varchar(20) DEFAULT NULL,
  `unit_price` double DEFAULT NULL,
  `units_in_stock` smallint(6) DEFAULT NULL,
  `units_on_order` smallint(6) DEFAULT NULL,
  `reorder_level` smallint(6) DEFAULT NULL,
  `discontinued` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `products_fk_category_id` (`category_id`),
  CONSTRAINT `products_fk_category_id` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=latin1;

INSERT INTO `categories` (`id`, `category_name`, `description`) VALUES
('1', 'Beverages', 'Soft drinks, coffees, teas, beers, and ales'),
('2', 'Condiments', 'Sweet and savory sauces, relishes, spreads, and seasonings'),
('3', 'Confections', 'Desserts, candies, and sweet breads'),
('4', 'Dairy Products', 'Cheeses'),
('5', 'Grains/Cereals', 'Breads, crackers, pasta, and cereal'),
('6', 'Meat/Poultry', 'Prepared meats'),
('7', 'Produce', 'Dried fruit and bean curd'),
('8', 'Seafood', 'Seaweed and fish');

INSERT INTO `customers` (`id`, `company_name`, `contact_name`, `contact_title`, `address`, `city`, `region`, `postal_code`, `country`, `phone`, `fax`) VALUES
('1', 'Alfreds Futterkiste', 'Maria Anders', 'Sales Representative', 'Obere Str. 57', 'Berlin', NULL, '12209', 'Germany', '030-0074321', '030-0076545'),
('2', 'Ana Trujillo Emparedados y helados', 'Ana Trujillo', 'Owner', 'Avda. de la Constitución 2222', 'México D.F.', NULL, '05021', 'Mexico', '(5) 555-4729', '(5) 555-3745'),
('3', 'Antonio Moreno Taquería', 'Antonio Moreno', 'Owner', 'Mataderos  2312', 'México D.F.', NULL, '05023', 'Mexico', '(5) 555-3932', NULL),
('4', 'Around the Horn', 'Thomas Hardy', 'Sales Representative', '120 Hanover Sq.', 'London', NULL, 'WA1 1DP', 'UK', '(171) 555-7788', '(171) 555-6750'),
('5', 'Berglunds snabbköp', 'Christina Berglund', 'Order Administrator', 'Berguvsvägen  8', 'Luleå', NULL, 'S-958 22', 'Sweden', '0921-12 34 65', '0921-12 34 67'),
('6', 'Blauer See Delikatessen', 'Hanna Moos', 'Sales Representative', 'Forsterstr. 57', 'Mannheim', NULL, '68306', 'Germany', '0621-08460', '0621-08924'),
('7', 'Blondesddsl père et fils', 'Frédérique Citeaux', 'Marketing Manager', '24, place Kléber', 'Strasbourg', NULL, '67000', 'France', '88.60.15.31', '88.60.15.32'),
('8', 'Bólido Comidas preparadas', 'Martín Sommer', 'Owner', 'C/ Araquil, 67', 'Madrid', NULL, '28023', 'Spain', '(91) 555 22 82', '(91) 555 91 99'),
('9', 'Bon app\'', 'Laurence Lebihan', 'Owner', '12, rue des Bouchers', 'Marseille', NULL, '13008', 'France', '91.24.45.40', '91.24.45.41'),
('10', 'Bottom-Dollar Markets', 'Elizabeth Lincoln', 'Accounting Manager', '23 Tsawassen Blvd.', 'Tsawassen', 'BC', 'T2F 8M4', 'Canada', '(604) 555-4729', '(604) 555-3745');

INSERT INTO `order_details` (`order_id`, `product_id`, `unit_price`, `quantity`, `discount`) VALUES
('1', '11', '14', '12', '0'),
('1', '42', '9.80000019', '10', '0'),
('1', '72', '34.7999992', '5', '0'),
('2', '14', '18.6000004', '9', '0'),
('2', '51', '42.4000015', '40', '0'),
('3', '41', '7.69999981', '10', '0'),
('3', '51', '42.4000015', '35', '0.150000006'),
('3', '65', '16.7999992', '15', '0.150000006'),
('4', '22', '16.7999992', '6', '0.0500000007'),
('4', '57', '15.6000004', '15', '0.0500000007'),
('4', '65', '16.7999992', '20', '0'),
('5', '20', '64.8000031', '40', '0.0500000007'),
('5', '33', '2', '25', '0.0500000007'),
('5', '60', '27.2000008', '40', '0'),
('6', '31', '10', '20', '0'),
('6', '39', '14.3999996', '42', '0'),
('6', '49', '16', '40', '0'),
('7', '24', '3.5999999', '15', '0.150000006'),
('7', '55', '19.2000008', '21', '0.150000006'),
('7', '74', '8', '21', '0'),
('8', '2', '15.1999998', '20', '0'),
('8', '16', '13.8999996', '35', '0'),
('8', '36', '15.1999998', '25', '0'),
('8', '59', '44', '30', '0'),
('9', '53', '26.2000008', '15', '0'),
('9', '77', '10.3999996', '12', '0'),
('10', '27', '35.0999985', '25', '0'),
('10', '39', '14.3999996', '6', '0'),
('10', '77', '10.3999996', '15', '0'),
('11', '2', '15.1999998', '50', '0.200000003'),
('11', '5', '17', '65', '0.200000003'),
('11', '32', '25.6000004', '6', '0.200000003'),
('12', '21', '8', '10', '0'),
('12', '37', '20.7999992', '1', '0'),
('13', '41', '7.69999981', '16', '0.25'),
('13', '57', '15.6000004', '50', '0'),
('13', '62', '39.4000015', '15', '0.25'),
('13', '70', '12', '21', '0.25'),
('14', '21', '8', '20', '0'),
('14', '35', '14.3999996', '20', '0'),
('15', '5', '17', '12', '0.200000003'),
('15', '7', '24', '15', '0'),
('15', '56', '30.3999996', '2', '0'),
('16', '16', '13.8999996', '60', '0.25'),
('16', '24', '3.5999999', '28', '0'),
('16', '30', '20.7000008', '60', '0.25'),
('16', '74', '8', '36', '0.25'),
('17', '2', '15.1999998', '35', '0'),
('17', '41', '7.69999981', '25', '0.150000006'),
('18', '17', '31.2000008', '30', '0'),
('18', '70', '12', '20', '0'),
('19', '12', '30.3999996', '12', '0.0500000007'),
('20', '40', '14.6999998', '50', '0'),
('20', '59', '44', '70', '0.150000006'),
('20', '76', '14.3999996', '15', '0.150000006'),
('21', '29', '99', '10', '0'),
('21', '72', '27.7999992', '4', '0');

INSERT INTO `orders` (`id`, `customer_id`, `order_date`, `required_date`, `shipped_date`, `freight`, `ship_name`, `ship_address`, `ship_city`, `ship_region`, `ship_postal_code`, `ship_country`) VALUES
('1', '1', '1996-07-04', '1996-08-01', '1996-07-16', '32.3800011', 'Vins et alcools Chevalier', '59 rue de l\'Abbaye', 'Reims', NULL, '51100', 'France'),
('2', '2', '1996-07-05', '1996-08-16', '1996-07-10', '11.6099997', 'Toms Spezialitäten', 'Luisenstr. 48', 'Münster', NULL, '44087', 'Germany'),
('3', '3', '1996-07-08', '1996-08-05', '1996-07-12', '65.8300018', 'Hanari Carnes', 'Rua do Paço, 67', 'Rio de Janeiro', 'RJ', '05454-876', 'Brazil'),
('4', '4', '1996-07-08', '1996-08-05', '1996-07-15', '41.3400002', 'Victuailles en stock', '2, rue du Commerce', 'Lyon', NULL, '69004', 'France'),
('5', '4', '1996-07-09', '1996-08-06', '1996-07-11', '51.2999992', 'Suprêmes délices', 'Boulevard Tirou, 255', 'Charleroi', NULL, 'B-6000', 'Belgium'),
('6', '3', '1996-07-10', '1996-07-24', '1996-07-16', '58.1699982', 'Hanari Carnes', 'Rua do Paço, 67', 'Rio de Janeiro', 'RJ', '05454-876', 'Brazil'),
('7', '5', '1996-07-11', '1996-08-08', '1996-07-23', '22.9799995', 'Chop-suey Chinese', 'Hauptstr. 31', 'Bern', NULL, '3012', 'Switzerland'),
('8', '9', '1996-07-12', '1996-08-09', '1996-07-15', '148.330002', 'Richter Supermarkt', 'Starenweg 5', 'Genève', NULL, '1204', 'Switzerland'),
('9', '3', '1996-07-15', '1996-08-12', '1996-07-17', '13.9700003', 'Wellington Importadora', 'Rua do Mercado, 12', 'Resende', 'SP', '08737-363', 'Brazil'),
('10', '4', '1996-07-16', '1996-08-13', '1996-07-22', '81.9100037', 'HILARION-Abastos', 'Carrera 22 con Ave. Carlos Soublette #8-35', 'San Cristóbal', 'Táchira', '5022', 'Venezuela'),
('11', '1', '1996-07-17', '1996-08-14', '1996-07-23', '140.509995', 'Ernst Handel', 'Kirchgasse 6', 'Graz', NULL, '8010', 'Austria'),
('12', '4', '1996-07-18', '1996-08-15', '1996-07-25', '3.25', 'Centro comercial Moctezuma', 'Sierras de Granada 9993', 'México D.F.', NULL, '05022', 'Mexico'),
('13', '4', '1996-07-19', '1996-08-16', '1996-07-29', '55.0900002', 'Ottilies Käseladen', 'Mehrheimerstr. 369', 'Köln', NULL, '50739', 'Germany'),
('14', '4', '1996-07-19', '1996-08-16', '1996-07-30', '3.04999995', 'Que Delícia', 'Rua da Panificadora, 12', 'Rio de Janeiro', 'RJ', '02389-673', 'Brazil'),
('15', '8', '1996-07-22', '1996-08-19', '1996-07-25', '48.2900009', 'Rattlesnake Canyon Grocery', '2817 Milton Dr.', 'Albuquerque', 'NM', '87110', 'USA'),
('16', '9', '1996-07-23', '1996-08-20', '1996-07-31', '146.059998', 'Ernst Handel', 'Kirchgasse 6', 'Graz', NULL, '8010', 'Austria'),
('17', '6', '1996-07-24', '1996-08-21', '1996-08-23', '3.67000008', 'Folk och fä HB', 'Åkergatan 24', 'Bräcke', NULL, 'S-844 67', 'Sweden'),
('18', '2', '1996-07-25', '1996-08-22', '1996-08-12', '55.2799988', 'Blondel père et fils', '24, place Kléber', 'Strasbourg', NULL, '67000', 'France'),
('19', '3', '1996-07-26', '1996-09-06', '1996-07-31', '25.7299995', 'Wartian Herkku', 'Torikatu 38', 'Oulu', NULL, '90110', 'Finland'),
('20', '4', '1996-07-29', '1996-08-26', '1996-08-06', '208.580002', 'Frankenversand', 'Berliner Platz 43', 'München', NULL, '80805', 'Germany'),
('21', '8', '1996-07-30', '1996-08-27', '1996-08-02', '66.2900009', 'GROSELLA-Restaurante', '5ª Ave. Los Palos Grandes', 'Caracas', 'DF', '1081', 'Venezuela'),
('22', '5', '1996-07-31', '1996-08-14', '1996-08-09', '4.55999994', 'White Clover Markets', '1029 - 12th Ave. S.', 'Seattle', 'WA', '98124', 'USA'),
('23', '1', '1996-08-01', '1996-08-29', '1996-08-02', '136.539993', 'Wartian Herkku', 'Torikatu 38', 'Oulu', NULL, '90110', 'Finland'),
('24', '6', '1996-08-01', '1996-08-29', '1996-08-30', '4.53999996', 'Split Rail Beer & Ale', 'P.O. Box 555', 'Lander', 'WY', '82520', 'USA'),
('25', '6', '1996-08-02', '1996-08-30', '1996-08-06', '98.0299988', 'Rattlesnake Canyon Grocery', '2817 Milton Dr.', 'Albuquerque', 'NM', '87110', 'USA');

INSERT INTO `products` (`id`, `product_name`, `supplier_id`, `category_id`, `quantity_per_unit`, `unit_price`, `units_in_stock`, `units_on_order`, `reorder_level`, `discontinued`) VALUES
('1', 'Chai', '8', '1', '10 boxes x 30 bags', '18', '39', '0', '10', '1'),
('2', 'Chang', '1', '1', '24 - 12 oz bottles', '19', '17', '40', '25', '1'),
('3', 'Aniseed Syrup', '1', '2', '12 - 550 ml bottles', '10', '13', '70', '25', '0'),
('4', 'Chef Anton\'s Cajun Seasoning', '2', '2', '48 - 6 oz jars', '22', '53', '0', '0', '0'),
('5', 'Chef Anton\'s Gumbo Mix', '2', '2', '36 boxes', '21.3500004', '0', '0', '0', '1'),
('6', 'Grandma\'s Boysenberry Spread', '3', '2', '12 - 8 oz jars', '25', '120', '0', '25', '0'),
('7', 'Uncle Bob\'s Organic Dried Pears', '3', '7', '12 - 1 lb pkgs.', '30', '15', '0', '10', '0'),
('8', 'Northwoods Cranberry Sauce', '3', '2', '12 - 12 oz jars', '40', '6', '0', '0', '0'),
('9', 'Mishi Kobe Niku', '4', '6', '18 - 500 g pkgs.', '97', '29', '0', '0', '1'),
('10', 'Ikura', '4', '8', '12 - 200 ml jars', '31', '31', '0', '0', '0'),
('11', 'Queso Cabrales', '5', '4', '1 kg pkg.', '21', '22', '30', '30', '0'),
('12', 'Queso Manchego La Pastora', '5', '4', '10 - 500 g pkgs.', '38', '86', '0', '0', '0'),
('13', 'Konbu', '6', '8', '2 kg box', '6', '24', '0', '5', '0'),
('14', 'Tofu', '6', '7', '40 - 100 g pkgs.', '23.25', '35', '0', '0', '0'),
('15', 'Genen Shouyu', '6', '2', '24 - 250 ml bottles', '13', '39', '0', '5', '0'),
('16', 'Pavlova', '7', '3', '32 - 500 g boxes', '17.4500008', '29', '0', '10', '0'),
('17', 'Alice Mutton', '7', '6', '20 - 1 kg tins', '39', '0', '0', '0', '1'),
('18', 'Carnarvon Tigers', '7', '8', '16 kg pkg.', '62.5', '42', '0', '0', '0'),
('19', 'Teatime Chocolate Biscuits', '8', '3', '10 boxes x 12 pieces', '9.19999981', '25', '0', '5', '0'),
('20', 'Sir Rodney\'s Marmalade', '8', '3', '30 gift boxes', '81', '40', '0', '0', '0'),
('21', 'Sir Rodney\'s Scones', '8', '3', '24 pkgs. x 4 pieces', '10', '3', '40', '5', '0'),
('22', 'Gustaf\'s Knäckebröd', '9', '5', '24 - 500 g pkgs.', '21', '104', '0', '25', '0'),
('23', 'Tunnbröd', '9', '5', '12 - 250 g pkgs.', '9', '61', '0', '25', '0'),
('24', 'Guaraná Fantástica', '10', '1', '12 - 355 ml cans', '4.5', '20', '0', '0', '1'),
('25', 'NuNuCa Nuß-Nougat-Creme', '11', '3', '20 - 450 g glasses', '14', '76', '0', '30', '0'),
('26', 'Gumbär Gummibärchen', '11', '3', '100 - 250 g bags', '31.2299995', '15', '0', '0', '0'),
('27', 'Schoggi Schokolade', '11', '3', '100 - 100 g pieces', '43.9000015', '49', '0', '30', '0'),
('28', 'Rössle Sauerkraut', '12', '7', '25 - 825 g cans', '45.5999985', '26', '0', '0', '1'),
('29', 'Thüringer Rostbratwurst', '12', '6', '50 bags x 30 sausgs.', '123.790001', '0', '0', '0', '1'),
('30', 'Nord-Ost Matjeshering', '13', '8', '10 - 200 g glasses', '25.8899994', '10', '0', '15', '0'),
('31', 'Gorgonzola Telino', '14', '4', '12 - 100 g pkgs', '12.5', '0', '70', '20', '0'),
('32', 'Mascarpone Fabioli', '14', '4', '24 - 200 g pkgs.', '32', '9', '40', '25', '0'),
('33', 'Geitost', '15', '4', '500 g', '2.5', '112', '0', '20', '0'),
('34', 'Sasquatch Ale', '16', '1', '24 - 12 oz bottles', '14', '111', '0', '15', '0'),
('35', 'Steeleye Stout', '16', '1', '24 - 12 oz bottles', '18', '20', '0', '15', '0'),
('36', 'Inlagd Sill', '17', '8', '24 - 250 g  jars', '19', '112', '0', '20', '0'),
('37', 'Gravad lax', '17', '8', '12 - 500 g pkgs.', '26', '11', '50', '25', '0'),
('38', 'Côte de Blaye', '18', '1', '12 - 75 cl bottles', '263.5', '17', '0', '15', '0'),
('39', 'Chartreuse verte', '18', '1', '750 cc per bottle', '18', '69', '0', '5', '0'),
('40', 'Boston Crab Meat', '19', '8', '24 - 4 oz tins', '18.3999996', '123', '0', '30', '0'),
('41', 'Jack\'s New England Clam Chowder', '19', '8', '12 - 12 oz cans', '9.64999962', '85', '0', '10', '0'),
('42', 'Singaporean Hokkien Fried Mee', '20', '5', '32 - 1 kg pkgs.', '14', '26', '0', '0', '1'),
('43', 'Ipoh Coffee', '20', '1', '16 - 500 g tins', '46', '17', '10', '25', '0'),
('44', 'Gula Malacca', '20', '2', '20 - 2 kg bags', '19.4500008', '27', '0', '15', '0'),
('45', 'Rogede sild', '21', '8', '1k pkg.', '9.5', '5', '70', '15', '0'),
('46', 'Spegesild', '21', '8', '4 - 450 g glasses', '12', '95', '0', '0', '0'),
('47', 'Zaanse koeken', '22', '3', '10 - 4 oz boxes', '9.5', '36', '0', '0', '0'),
('48', 'Chocolade', '22', '3', '10 pkgs.', '12.75', '15', '70', '25', '0'),
('49', 'Maxilaku', '23', '3', '24 - 50 g pkgs.', '20', '10', '60', '15', '0'),
('50', 'Valkoinen suklaa', '23', '3', '12 - 100 g bars', '16.25', '65', '0', '30', '0'),
('51', 'Manjimup Dried Apples', '24', '7', '50 - 300 g pkgs.', '53', '20', '0', '10', '0'),
('52', 'Filo Mix', '24', '5', '16 - 2 kg boxes', '7', '38', '0', '25', '0'),
('53', 'Perth Pasties', '24', '6', '48 pieces', '32.7999992', '0', '0', '0', '1'),
('54', 'Tourtière', '25', '6', '16 pies', '7.44999981', '21', '0', '10', '0'),
('55', 'Pâté chinois', '25', '6', '24 boxes x 2 pies', '24', '115', '0', '20', '0'),
('56', 'Gnocchi di nonna Alice', '26', '5', '24 - 250 g pkgs.', '38', '21', '10', '30', '0'),
('57', 'Ravioli Angelo', '26', '5', '24 - 250 g pkgs.', '19.5', '36', '0', '20', '0'),
('58', 'Escargots de Bourgogne', '27', '8', '24 pieces', '13.25', '62', '0', '20', '0'),
('59', 'Raclette Courdavault', '28', '4', '5 kg pkg.', '55', '79', '0', '0', '0'),
('60', 'Camembert Pierrot', '28', '4', '15 - 300 g rounds', '34', '19', '0', '0', '0'),
('61', 'Sirop d\'érable', '29', '2', '24 - 500 ml bottles', '28.5', '113', '0', '25', '0'),
('62', 'Tarte au sucre', '29', '3', '48 pies', '49.2999992', '17', '0', '0', '0'),
('63', 'Vegie-spread', '7', '2', '15 - 625 g jars', '43.9000015', '24', '0', '5', '0'),
('64', 'Wimmers gute Semmelknödel', '12', '5', '20 bags x 4 pieces', '33.25', '22', '80', '30', '0'),
('65', 'Louisiana Fiery Hot Pepper Sauce', '2', '2', '32 - 8 oz bottles', '21.0499992', '76', '0', '0', '0'),
('66', 'Louisiana Hot Spiced Okra', '2', '2', '24 - 8 oz jars', '17', '4', '100', '20', '0'),
('67', 'Laughing Lumberjack Lager', '16', '1', '24 - 12 oz bottles', '14', '52', '0', '10', '0'),
('68', 'Scottish Longbreads', '8', '3', '10 boxes x 8 pieces', '12.5', '6', '10', '15', '0'),
('69', 'Gudbrandsdalsost', '15', '4', '10 kg pkg.', '36', '26', '0', '15', '0'),
('70', 'Outback Lager', '7', '1', '24 - 355 ml bottles', '15', '15', '10', '30', '0'),
('71', 'Flotemysost', '15', '4', '10 - 500 g pkgs.', '21.5', '26', '0', '0', '0'),
('72', 'Mozzarella di Giovanni', '14', '4', '24 - 200 g pkgs.', '34.7999992', '14', '0', '0', '0'),
('73', 'Röd Kaviar', '17', '8', '24 - 150 g jars', '15', '101', '0', '5', '0'),
('74', 'Longlife Tofu', '4', '7', '5 kg pkg.', '10', '4', '20', '5', '0'),
('75', 'Rhönbräu Klosterbier', '12', '1', '24 - 0.5 l bottles', '7.75', '125', '0', '25', '0'),
('76', 'Lakkalikööri', '23', '1', '500 ml', '18', '57', '0', '20', '0'),
('77', 'Original Frankfurter grüne Soße', '12', '2', '12 boxes', '13', '32', '0', '15', '0');




/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;