package harper

import (
	"testing"
	"time"
)

func TestSystemStatus(t *testing.T) {
	guest := randomID()

	// try with standard super user
	_, err := c.SystemInformationAll()
	if err != nil {
		t.Fatal(err)
	}

	// try with non-super user
	p := Permission{}
	p.SetSuperUser(false)
	p.SetClusterUser(false)
	role, err := c.AddRole(guest, p)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropRole(role.ID)

	err = c.AddUser(guest, guest, role.Role, true)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropUser(guest)

	guestClient := NewClient(DEFAULT_ENDPOINT, guest, guest)
	_, err = guestClient.SystemInformationAll()
	if e, ok := err.(*OperationError); ok && e.IsNotAuthorizedError() {
		return
	}

	t.Fatal("expected SystemInfo call with guest user to fail")
}

func TestGetConfiguration(t *testing.T) {
	if _, err := c.GetConfiguration(); err != nil {
		t.Fatal(err)
	}
}

func TestInstallNodeModule(t *testing.T) {
	if _, err := c.AddCustomFunctionProject("my-project"); err != nil {
		t.Fatal(err)
	}

	wait()

	if _, err := c.InstallNodeModules([]string{"my-project"}, false); err != nil {
		t.Fatal(err)
	}

	wait()

	if _, err := c.DropComponent("my-project", ""); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteRecordsBefore(t *testing.T) {
	database := randomID()
	table := randomID()

	if err := c.CreateDatabase(database); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.CreateTable(database, table, "id"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	record := createTestRecord()

	if _, err := c.Insert(database, table, []interface{}{record}); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	if message, err := c.DeleteRecordsBefore(time.Now(), database, table); err != nil {
		t.Log(message)
		t.Fatal(err)
	}

	time.Sleep(time.Second * 5)
	resp, err := c.DescribeTable(database, table)
	if err != nil {

		t.Fatal(err)
	}
	t.Log(resp)
	t.Log(resp.RecordCount)
	if resp.RecordCount != 0 {
		t.Fatal("There should not be any records left in the table")
	}
}

func TestSetConfiguration(t *testing.T) {
	var configuration = map[string]interface{}{}
	configuration["LOGGING_ROTATION_ENABLED"] = true
	if _, err := c.SetConfiguration(configuration); err != nil {
		t.Fatal(err)
	}
}
