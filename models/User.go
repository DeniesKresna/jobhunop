Package Models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"size:255;not null;unique" json:"nickname"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:size:255;not null;unique" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword tring, password string) error {
	return bcrypt.CompareHashAndPassword([]yte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() 
	u.ID = 0
	.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) eror {
	swtch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := chckmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Ivalid Email")
		}
		rturn nil
	case "login":
		if u.Password == "" {
			eturn errors.New("Required Password")
		}
		if u.Email == "" {
			eturn errors.New("Required Email")
		}
		if err = checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Ivalid Email")
		}
		rturn nil
	default:
		if u.Nickname == "" {
			eturn errors.New("Required Nickname")
		}
		if u.Password == "" {
			eturn errors.New("Required Password")
		}
		if u.Email == "" {
			eturn errors.New("Required Email")
		}
		f err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
		return nil
	}
}

func (u *User) SaveUsr(db *gorm.DB) (*User, error) {
	vr err error
	err = db.Debu().Create(&u).Error
	f err != nil {
	return &User{}, err
	} 
	return u, nil
}

func (u *User) FndAllUsers(db *gorm.DB) (*[]User, error){
	var err error
	uers := []User{}
	err = db.Debu().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {

		return &[]User{}, err
	}
	return &users
}