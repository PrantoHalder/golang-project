package postgres

import (
	"fmt"
	"log"

	"main.go/usermgm/storage"
)
//admin register
const registerAdminQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	role
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:role
)RETURNING *`
func(s PostGressStorage) RegisterAdmin(u storage.Admin) (*storage.Admin, error){
	stmt, err := s.DB.PrepareNamed(registerAdminQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Println("error is in the query section of registration section")
		return nil, err
	}
	if u.ID == 0 {
		log.Println("error is in the query section of registration section u.ID == 0")
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}
//doctor register by admin
const registerDoctorAminQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	role
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:role
)RETURNING *`
func(s PostGressStorage) RegisterDoctorAdmin(u storage.DoctorU) (*storage.DoctorU, error){
	stmt, err := s.DB.PrepareNamed(registerDoctorAminQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Println("error is in the query section of registration section")
		return nil, err
	}
	if u.ID == 0 {
		log.Println("error is in the query section of registration section u.ID == 0")
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}

//admin login
const getAdminByUsernameQuery=`SELECT *  
FROM users
WHERE
username = $1
AND
role ='admin'
AND
is_active = true
AND
deleted_at IS NULL`
func (s PostGressStorage) GetAdminByUsername(username string) (*storage.Admin, error) {
	var listUser storage.Admin
	if err := s.DB.Get(&listUser,getAdminByUsernameQuery,username); err != nil {
		log.Println("error is in the query section of usermgm login section")
		return nil, err
	}
	if listUser.ID == 0 {
	 log.Println("error is in the query section of usermgm ID==0 admin login section")
     return nil,fmt.Errorf("unable to find username")
	}
	return &listUser, nil
}
//user register by admin
const registerpatientByAdminQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	role
) values(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:role
)RETURNING *`
func(s PostGressStorage) RegisterPatient(u storage.User) (*storage.User, error){
	stmt, err := s.DB.PrepareNamed(registerpatientByAdminQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Println("error is in the query section of registration section")
		return nil, err
	}
	if u.ID == 0 {
		log.Println("error is in the query section of registration section u.ID == 0")
		log.Println("unable to create user")
		return &u, fmt.Errorf("unable to create user")
	}
	return &u, nil
}