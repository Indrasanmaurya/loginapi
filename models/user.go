package models

type User struct {
	Email    string
	Password string // hashed password
}

// Dummy user list (as database)
var Users = []User{
	{
		Email:    "alex.j@example.com",
		Password: "$2a$10$7tZLd7xRUie1Ku2XmDpfaevDldPNM0DNZ93.QBhNsf8HDoVhcvPMO", // password123
	},
}
