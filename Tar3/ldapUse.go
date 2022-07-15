package main

/**
author: Wenjie_pan
*/
import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-ldap/ldap"
)

type User struct {
	username    string
	password    string
	telephone   string
	emailSuffix string
	snUsername  string
	uid         string
	gid         string
}

// LoginBind  connection ldap server and binding ldap server
func LoginBind(ldapUser, ldapPassword, urlUse string) (*ldap.Conn, string, error) {
	f, _ := os.Create("first.txt")
	getVal := string(ldapPassword)
	io.WriteString(f, getVal)
	l, err := ldap.DialURL(urlUse)
	// "ldap:0.0.0.0:389"
	if err != nil {
		return nil, " ", err
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: fmt.Sprintf("cn=%s,dc=devopsman,dc=cn", ldapUser),
		Password: ldapPassword,
	})

	if err != nil {
		// fmt.Println("ldap password is error: ", ldap.LDAPResultInvalidCredentials)
		return nil, "ldap admin or password is error!", err
	}
	res := ldapUser + " is existing!"
	return l, res, nil
}

func (user *User) addUser(conn *ldap.Conn) error {
	ldaprow := ldap.NewAddRequest(fmt.Sprintf("cn=%s,dc=devopsman,dc=cn", user.username), nil)
	ldaprow.Attribute("userPassword", []string{user.password})
	ldaprow.Attribute("homeDirectory", []string{fmt.Sprintf("/home/%s", user.username)})
	ldaprow.Attribute("cn", []string{user.username})
	ldaprow.Attribute("uid", []string{user.username})
	ldaprow.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})
	ldaprow.Attribute("uidNumber", []string{"2201"})
	ldaprow.Attribute("gidNumber", []string{"2201"})
	ldaprow.Attribute("loginShell", []string{"/bin/bash"})
	if err := conn.Add(ldaprow); err != nil {
		return err
	}
	return nil
}

func GetEmployees(con *ldap.Conn) ([]string, error) {
	var employees []string
	sql := ldap.NewSearchRequest("dc=devopsman,dc=cn",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		// "(objectClass=*)",
		//First verify the name
		fmt.Sprintf("(&(objectClass=account)(uid=%s))", ldap.EscapeFilter("wenjie_pan")),
		//this content in string can known as the entry what show in next context
		[]string{"dn", "cn", "objectClass"},
		nil)

	cur, err := con.Search(sql)
	if err != nil {
		return nil, err
	}

	//Check the dn
	userdn := cur.Entries[0].DN
	fmt.Println(userdn)

	//Verify the password of an internal domain user
	err = con.Bind(userdn, "admin123")
	if err != nil {
		fmt.Print("login failure because of the wrong user's password")
		// log.Fatal(err)
	}
	//This step is a circular output
	if len(cur.Entries) > 0 {
		for _, item := range cur.Entries {
			// fmt.Print(item)
			cn := item.GetAttributeValues("cn")
			for _, iCn := range cn {
				employees = append(employees, strings.Split(iCn, "[")[0])
			}
		}
		return employees, nil
	}
	return nil, nil
}
