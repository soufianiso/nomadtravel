package service


import(
	"golang.org/x/crypto/bcrypt"
)

func validatePassword(hashedPassword string, password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil{
		return err
	}
	return 	nil

}

func hashPassword(password string) (string, error){
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil{
		return   "" ,err
	}
	
	return string(hashedpassword), nil
}
