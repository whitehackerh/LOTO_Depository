package Model

type User struct {
	User_id      int
	User_name    string
	Password     string
	Mail_address string
	Admin_flag   int
	Delete_flag  int
}

// func CreateUser(input_data []string) bool {
// 	query := `INSERT INTO users Values(user_id, user_name, password, mail_address, admin_flag, delete_flag) Values ($1, $2, $3, $4, $5, $6);`
// 	_, err := Db.Exec(query, input_data["user_id"], input_data["user_name"], input_data["password"], input_data["mail_address"], input_data["admin_flag"], input_data["delete_flag"])
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

func GetUser(input_data string) []*User {
	user := User{}
	data := []*User{}

	query := `SELECT user_id, user_name, pgp_sym_decrypt(password, get_passwd()) as password, pgp_sym_decrypt(mail_address, get_passwd()) as mail_address, admin_flag, delete_flag FROM users WHERE user_name = $1;`
	rows, err := Db.Query(query, input_data)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.User_id, &user.User_name, &user.Password, &user.Mail_address, &user.Admin_flag, &user.Delete_flag)
		data = append(data, &User{User_id: user.User_id, User_name: user.User_name, Password: user.Password, Mail_address: user.Mail_address, Admin_flag: user.Admin_flag, Delete_flag: user.Delete_flag})
	}
	return data
}
