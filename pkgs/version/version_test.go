package version

import (
	"fmt"
	"testing"
	"github.com/coreos/go-semver/semver"
)

func TestCheckRc(t *testing.T) {
	got := CheckRc("v1.0.0-rc.2")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got_two := CheckRc("v1.0.0")
	want_two := false

	if got_two != want_two {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestGetLatestTag(t *testing.T) {
	got, err := GetLatestTag()
	fmt.Printf("got: %v\n, err: %v", got, err)
}

func TestTrimTag(t *testing.T) {
	got := TrimTag("v1.0.0-rc.2")
	want := "1.0.0-rc.2"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	got_two := TrimTag("v1.0.0")
	want_two := "1.0.0"

	if got_two != want_two {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGetVersionType(t *testing.T) {
	var commitTypes = [...]string{"breaking", "feature", "bugfix"}
	// breaking test
	commitMessage := "breaking - this is a new feature"

	versionType := GetVersionType(commitMessage, commitTypes, "")
	versionType_want := "breaking"

	if versionType != versionType_want {
		t.Errorf("got %s, wanted %s", versionType, versionType_want)
	}

	// feature test

	commitMessage_two := "feature this is a bugfix"

	versionType_two := GetVersionType(commitMessage_two, commitTypes, "")
	versionType_want_two := "feature"

	if versionType_two != versionType_want_two {
		t.Errorf("got %s, wanted %s", versionType_two, versionType_want_two)
	}

	// bugfix test
	commitMessage_three := "this is a bugfix"

	versionType_three := GetVersionType(commitMessage_three, commitTypes, "")
	versionType_want_three := "bugfix"

	if versionType_three != versionType_want_three {
		t.Errorf("got %s, wanted %s", versionType_three, versionType_want_three)
	}


}

func TestSemVerToString(t *testing.T) {
	semver := semver.New("1.2.0")
	got := SemVerToString(semver)
	want := "1.2.0"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestRemovePrefix(t *testing.T) {
	got := RemovePrefix("v5.2.1", "v")
	want := "5.2.1"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestAddV(t *testing.T) {
	got := AddV("5.2.1")
	want := "v5.2.1"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestMakeSemVer(t *testing.T) {
	got := MakeSemVer("5.2.1")
	var s interface{} = got
	if _, ok := s.(*semver.Version); ok {
		fmt.Println("This is an int")
	} else {
		t.Errorf("got %s, wanted *semver.Version", s)
	}
}
func TestRemoveSuffix(t *testing.T) {
	got := RemoveSuffix("v5.2.1-rc.3", "-rc.")
	want := "v5.2.1"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
	got_two := RemoveSuffix("5.2.1-rc.3", "-rc.")
	want_two := "5.2.1"

	if got_two != want_two {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestIncreaseRc(t *testing.T) {
	got, err := IncreaseRc("v5.2.1-rc.3")
	if err != nil {
		t.Fatal("Error converting rc version to int:", err)
	}
	want := "v5.2.1-rc.4"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestBump(t *testing.T) {
	var bumps = map[string]string{"breaking": "major", "feature": "minor", "bugfix": "patch"}

	// major
	bumped := Bump(bumps, "breaking", MakeSemVer("2.5.0"))
	got := SemVerToString(bumped)

	want := "3.0.0"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	// minor
	bumped_two := Bump(bumps, "feature", MakeSemVer("2.5.0"))
	got_two := SemVerToString(bumped_two)

	want_two := "2.6.0"

	if got_two != want_two {
		t.Errorf("got %s, wanted %s", got_two, want_two)
	}

	// patch
	bumped_three := Bump(bumps, "bugfix", MakeSemVer("2.5.0"))
	got_three := SemVerToString(bumped_three)

	want_three := "2.5.1"

	if got_three != want_three {
		t.Errorf("got %s, wanted %s", got_three, want_three)
	}

}

func TestAddRc(t *testing.T) {
	got := AddRc("v5.2.1")
	want := "v5.2.1-rc.1"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
