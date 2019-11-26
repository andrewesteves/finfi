CREATE TABLE IF NOT EXISTS clients (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	phone VARCHAR(255),
	description TEXT,
	created_at DATETIME,
	updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS users (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(255),
	created_at DATETIME,
	updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS incomes (
	id INT PRIMARY KEY AUTO_INCREMENT,
	client_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	description text,
	status VARCHAR(255),
	installments INT,
	total decimal(10, 2),
	expired_at DATETIME,
	paid_at DATETIME,
	created_at DATETIME,
	updated_at DATETIME,
	FOREIGN KEY (client_id) REFERENCES clients(id)
);