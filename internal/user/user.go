package user

import "time"

type User struct {
	Id int64
	Username string
	Password string
	Name string
	Email string
	Phone string
	Address string
	Birthday time.Time
	CompanyId int64
	BranchId int64
 }
