package db_query

import "github.com/jinzhu/gorm"

type Tables struct {
	db *gorm.DB
}

func Newtables(db *gorm.DB) *Tables {
	return &Tables{db: db}
}

func (t *Tables) CreateSequence() error {
	return t.db.Exec("CREATE SEQUENCE order_by_user;").Error
}

func (t *Tables) CreateUserTable() error {
	return t.db.Exec(
		"CREATE TABLE users (" +
			" id integer DEFAULT nextval('order_by_user')," +
			" f_name VARCHAR(50) NOT NULL," +
			" l_name VARCHAR(50) NOT NULL," +
			" bio TEXT NOT NULL," +
			" PRIMARY KEY (id));").Error
}

func (t *Tables) CreateIdentityTable() error {
	return t.db.Exec(
		"CREATE TABLE identities(" +
			" id integer DEFAULT nextval('order_by_user')," +
			" login VARCHAR(50) UNIQUE NOT NULL," +
			" password VARCHAR(50) NOT NULL," +
			" salt VARCHAR(50) NOT NULL," +
			" blocked BOOLEAN NOT NULL," +
			" PRIMARY KEY (id)," +
			" user_id INTEGER REFERENCES users(id));").Error
}

func (t *Tables) CreateAdminTable() error {
	return t.db.Exec(
		"CREATE TYPE role AS ENUM ('admin', 'superUser');" +
			"CREATE TABLE admins(" +
			" id integer DEFAULT nextval('order_by_user')," +
			" role role," +
			" PRIMARY KEY (id)," +
			" identity_id INTEGER REFERENCES identities(id));").Error
}

func (t *Tables) CreateHobbiesTable() error {
	return t.db.Exec(
		"CREATE TYPE hobby AS ENUM ('sport', 'embroidery');" +
			"CREATE TABLE hobbies(" +
			" id serial," +
			" user_id int NOt Null," +
			" hobby hobby NOT NULL," +
			" PRIMARY KEY (id)," +
			" FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);").Error
}

func (t *Tables) CreateFriendsTable() error {
	return t.db.Exec(
		"DROP TABLE IF EXISTS friends;" +
			"CREATE TABLE friends(" +
			" user_from int NOT NULL," +
			" user_to int NOT NULL," +
			" send timestamp NOT NULL," +
			" approved BOOLEAN NOT NULL," +
			" PRIMARY KEY (user_from, user_to)," +
			" FOREIGN KEY (user_from) REFERENCES users(id) ON UPDATE CASCADE," +
			" FOREIGN KEY (user_to) REFERENCES users(id) ON UPDATE CASCADE);").Error
}
