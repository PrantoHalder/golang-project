package admin

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"main.go/usermgm/storage"
)

type AdminStore interface {
	GetAdminByUsername(string) (*storage.User, error)
	RegisterAdmin(storage.User) (*storage.User, error)
}

type CoreAdmin struct {
	store AdminStore
}

func NewCoreAdmin(us AdminStore) *CoreAdmin {
	return &CoreAdmin{
		store: us,
	}
}


// Admin registration function
func (cu CoreAdmin) RegisterAdmin(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after GenerateFromPassword")
		return nil, err
	}
	u.Password = string(hashPass)
	ru, err := cu.store.RegisterAdmin(u)
	if err != nil {
		fmt.Println("the error is in the core layer in Register after cu.store.Register")
		return nil, err
	}
	if ru == nil {
		fmt.Println("the error is in the core layer in Register after ru == nil")
		return nil, fmt.Errorf("enable to register")
	}
	return ru, nil
}
//admin login
func (cu CoreAdmin) GetAdminbyUsernameCore(login storage.Login) (*storage.User, error){
    user,err := cu.store.GetAdminByUsername(login.Username) 
	if err != nil {
		fmt.Println("the error is in the core layer in GetStatusbyUsernameCore after cu.store.GetUserByUsername(login) ")
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil{
		fmt.Println("the error is in the core layer in GetStatusbyUsernameCore after bcrypt.CompareHashAndPassword ")
		return nil, err
	}
	return user,nil
}