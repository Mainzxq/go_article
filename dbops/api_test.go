package dbops

import "testing"

// init(dblogin, truncate tables, ) -> run test -> clear data(truncate tables)

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate article_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("mainzxq", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("mainzxq")
	if pwd!="123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
  	err := DeleteUser("mainzxq", "123")
  	if err != nil {
		t.Errorf("Error of Del User: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("mainzxq")
	if err != nil {
		t.Errorf("Error of Reget User: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failedß")
	}
}