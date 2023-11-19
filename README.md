# nyx

first try at developing my own cli-app

As a personal remainder, this is what the table looks like :

+-------------+-------------+------+-----+---------+----------------+
| Field       | Type        | Null | Key | Default | Extra          |
+-------------+-------------+------+-----+---------+----------------+
| id          | int         | NO   | PRI | NULL    | auto_increment |
| commune     | varchar(30) | NO   |     | NULL    |                |
| code_postal | varchar(5)  | NO   |     | NULL    |                |
| departement | varchar(30) | NO   |     | NULL    |                |
+-------------+-------------+------+-----+---------+----------------+

## Results so far : SUCCESSFUL

```php
<?php

class CVilleV5 {

	// properties
	private id;
	private commune;
	private code_postal;
	private departement;

	public function __constructor() {}

	// Some methods begining with letter a to f

	// Getters
	public function getId() { return (int) $this->id; }
	public function getCommune() { return (string) $this->commune; }
	public function getCodePostal() { return (string) $this->code_postal; }
	public function getDepartement() { return (string) $this->departement; }

	// Some methods begining with letter h to z

	// TODO setters

}
```