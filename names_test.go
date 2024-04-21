package names

import (
	"reflect"
	"testing"
)

func TestFindNickName(t *testing.T) {
	tests := []struct {
		input  string
		expect []string // expected nick names
		ex2    string   // expected output string
	}{
		// Check simple nick name
		{
			input:  "Abc Def (Xyz)",
			expect: []string{"Xyz"},
			ex2:    "Abc Def",
		},

		// Check nick name3s with blanks in them.
		{
			input:  "Abc Def (Xyz Uvw)",
			expect: []string{"Xyz Uvw"},
			ex2:    "Abc Def",
		},
		// Check Multiple Nick Names
		{
			input:  "Abc Def (Xyz Uvw) (Rr)",
			expect: []string{"Xyz Uvw", "Rr"},
			ex2:    "Abc Def",
		},
	}
	// func findNickName(fullName string) (partsFound []string, newFullName string) {
	for ii, tt := range tests {
		var got2 string
		var got []string
		if got, got2 = findNickName(tt.input); !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("Test %d : findNickName() got = %v, expected %v", ii, got, tt.expect)
		}
		if got2 != tt.ex2 {
			t.Errorf("Test %d : findNickName() got2 = %s, ex2 = %s", ii, got2, tt.ex2)
		}
	}

}

func TestDelimitGluedName(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{ // delimit glued name
			input:  "ВладимирМихайловичМясоедов",
			expect: "Владимир Михайлович Мясоедов",
		},
		{ // delimit glued name
			input:  "AbcDefG.(Xyz),Jr",
			expect: "Abc Def G. (Xyz), Jr",
		},
		{ // leave intact correctly delimited name
			input:  "de la Fuzzy, Dr. Philip et Glova (Doc Fuzzy) J. Schlump III, Jr., Genius",
			expect: "de la Fuzzy, Dr. Philip et Glova (Doc Fuzzy) J. Schlump III, Jr., Genius",
		},
	}
	// func delimitGluedName(name string) string {
	for ii, tt := range tests {
		if got := delimitGluedName(tt.input); !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("Test %d : delimitGluedName() got = %s, expected %s", ii, got, tt.expect)
		}
	}
}
func TestSplitNameIntoParts(t *testing.T) {
	tests := []struct {
		input      string
		expectPart []string // expected nick names
		expectComm []bool   // expected output string
	}{
		// Check glued names
		{
			input:      "AbcDefG.(Xyz)",
			expectPart: []string{"Abc", "Def", "G.", "(Xyz)"},
			expectComm: []bool{false, false, false, false},
		},
		// Check simple
		{
			input:      "Abc Def (Xyz)",
			expectPart: []string{"Abc", "Def", "(Xyz)"},
			expectComm: []bool{false, false, false},
		},
		// Check wit a comma
		{
			input:      "Abc Def (Xyz), Genius",
			expectPart: []string{"Abc", "Def", "(Xyz)", "Genius"},
			expectComm: []bool{false, false, true, false},
		},
	}
	// func splitNameIntoParts(fullName string) (parts []string, comma []bool) {
	for ii, tt := range tests {
		if part, comm := splitNameIntoParts(tt.input); !reflect.DeepEqual(part, tt.expectPart) {
			t.Errorf("Test %d : splitNameIntoParts() got = %v, expected = %v", ii, part, tt.expectPart)
		} else if !reflect.DeepEqual(comm, tt.expectComm) {
			t.Errorf("Test %d : splitNameIntoParts() got = %v, expected = %v", ii, comm, tt.expectComm)
		}
	}

}

// func findPartsMap(list map[string]bool, nameParts *[]string, nameCommas *[]bool) (partsFound []string) {

func TestParseFullname(t *testing.T) {
	tests := []struct {
		input  string
		expect ParsedName
	}{
		// Check glued name
		{
			input:  "PhilipJ.Schlump(Fuzzy)",
			expect: ParsedName{First: "Philip", Middle: "J.", Last: "Schlump", Nick: "Fuzzy"},
		},
		// Just the name test
		{
			input:  "Philip Schlump",
			expect: ParsedName{First: "Philip", Last: "Schlump"},
		},
		// test with a title
		{
			input:  "Dr. Philip Schlump",
			expect: ParsedName{Title: "Dr.", First: "Philip", Last: "Schlump"},
		},
		// nickname added
		{
			input:  "Dr. Philip Schlump (Doc Fuzzy)",
			expect: ParsedName{Title: "Dr.", First: "Philip", Last: "Schlump", Nick: "Doc Fuzzy"},
		},
		// have a middle name
		{
			input:  "Philip J. Schlump",
			expect: ParsedName{First: "Philip", Middle: "J.", Last: "Schlump"},
		},
		// use a suffix
		{
			input:  "Philip Schlump III (Doc Fuzzy), Jr.",
			expect: ParsedName{First: "Philip", Last: "Schlump", Nick: "Doc Fuzzy", Suffix: "III, Jr."},
		},
		// all that can be in name world.
		{
			input:  "de la Fuzzy, Dr. Philip et Glova (Doc Fuzzy) J. Schlump III, Jr., Genius",
			expect: ParsedName{Title: "Dr.", First: "Philip et Glova", Middle: "J. Schlump", Last: "de la Fuzzy", Nick: "Doc Fuzzy", Suffix: "III, Jr., Genius"},
		},
		// just a stand alone last
		{
			input:  "Brown",
			expect: ParsedName{Last: "Brown"},
		},
	}
	for ii, tt := range tests {
		if got := ParseFullName(tt.input); !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("Test %d : ParseFullName() got = %v, expected %v", ii, got, tt.expect)
		}
	}
}

func Test_findNickName(t *testing.T) {
	type args struct {
		fullName string
	}
	tests := []struct {
		name            string
		args            args
		wantPartsFound  []string
		wantNewFullName string
	}{
		// func findNickName(fullName string) (partsFound []string, newFullName string) {
		{
			name:            "1st test of findNickName",
			args:            args{fullName: "Philip Schlump (Fuzzy)"},
			wantPartsFound:  []string{"Fuzzy"},
			wantNewFullName: "Philip Schlump",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPartsFound, gotNewFullName := findNickName(tt.args.fullName)
			if !reflect.DeepEqual(gotPartsFound, tt.wantPartsFound) {
				t.Errorf("findNickName() gotPartsFound = %v, want %v", gotPartsFound, tt.wantPartsFound)
			}
			if gotNewFullName != tt.wantNewFullName {
				t.Errorf("findNickName() gotNewFullName = %v, want %v", gotNewFullName, tt.wantNewFullName)
			}
		})
	}
}
