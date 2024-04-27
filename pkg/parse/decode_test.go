package parse

import "testing"

func TestByteSliceToJSON(t *testing.T) {
	// test cases are base on the great work https://learnxinyminutes.com/docs/yaml/
	tests := []struct {
		name    string
		input   []byte
		want    string
		wantErr bool
	}{
		{
			input: []byte(`
---
key: value
yet_another_key: your value goes here.
a_number: 42
scientific: 1e+12
boolean: true
null_value: null
key with spaces: really?
but: 'A string in quotes.'
'keys can be quoted too.': "useful?"
single quotes: 'remember to escape ''them''.'
double quotes: "have many: \", \0, \t, \u263A, \x0d\x0a == \r\n, and more."
Superscript two: \u00B2
`),
			want:    `{"Superscript two":"\\u00B2","a_number":42,"boolean":true,"but":"A string in quotes.","double quotes":"have many: \", \u0000, \t, ☺, \r\n == \r\n, and more.","key":"value","key with spaces":"really?","keys can be quoted too.":"useful?","null_value":null,"scientific":1000000000000,"single quotes":"remember to escape 'them'.","yet_another_key":"your value goes here."}`,
			wantErr: false,
		},
		{
			input: []byte(`
literal_block: |
  This entire block of text will be the value of the 'literal_block' key,
  with line breaks being preserved.

  The literal continues until de-dented, and the leading indentation is
  stripped.

      Any lines that are 'more-indented' keep the rest of their indentation -
      these lines will be indented by 4 spaces.
folded_style: >
  This entire block of text will be the value of 'folded_style', but this
  time, all newlines will be replaced with a single space.

  Blank lines, like above, are converted to a newline character.

      'More-indented' lines keep their newlines, too -
      this text will appear over two lines.
`),
			want:    `{"folded_style":"This entire block of text will be the value of 'folded_style', but this time, all newlines will be replaced with a single space.\nBlank lines, like above, are converted to a newline character.\n\n    'More-indented' lines keep their newlines, too -\n    this text will appear over two lines.\n","literal_block":"This entire block of text will be the value of the 'literal_block' key,\nwith line breaks being preserved.\n\nThe literal continues until de-dented, and the leading indentation is\nstripped.\n\n    Any lines that are 'more-indented' keep the rest of their indentation -\n    these lines will be indented by 4 spaces.\n"}`,
			wantErr: false,
		},
		{
			input: []byte(`
# Nesting uses indentation. 2 space indent is preferred (but not required).
a_nested_map:
  key: value
  another_key: Another Value
  another_nested_map:
    hello: hello
# Maps don't have to have string keys.
0.25: a float key
# Keys can also be complex, like multi-line objects
# We use ? followed by a space to indicate the start of a complex key.
? |
  This is a key
  that has multiple lines
  : and this is its value
# Sequences (equivalent to lists or arrays) look like this
# (note that the '-' counts as indentation):
a_sequence:
  - Item 1
  - Item 2
  - 0.5  # sequences can contain disparate types.
  - Item 4
  - key: value
    another_key: another_value
  -
    - This is a sequence
    - inside another sequence
  - - - Nested sequence indicators
      - can be collapsed
# Since YAML is a superset of JSON, you can also write JSON-style maps and
# sequences:
json_map: {"key": "value"}
json_seq: [3, 2, 1, "takeoff"]
and quotes are optional: {key: [3, 2, 1, takeoff]}
`),
			want:    `{"0.25":"a float key","This is a key\nthat has multiple lines\n: and this is its value\n":null,"a_nested_map":{"another_key":"Another Value","another_nested_map":{"hello":"hello"},"key":"value"},"a_sequence":["Item 1","Item 2",0.5,"Item 4",{"another_key":"another_value","key":"value"},["This is a sequence","inside another sequence"],[["Nested sequence indicators","can be collapsed"]]],"and quotes are optional":{"key":[3,2,1,"takeoff"]},"json_map":{"key":"value"},"json_seq":[3,2,1,"takeoff"]}`,
			wantErr: false,
		},
		{
			input: []byte(`
anchored_content: &anchor_name This string will appear as the value of two keys.
other_anchor: *anchor_name
base: &base
  name: Everyone has same name
foo:
  <<: *base
  age: 10
bar:
  <<: *base
  age: 20
`),
			want:    `{"anchored_content":"This string will appear as the value of two keys.","bar":{"age":20,"name":"Everyone has same name"},"base":{"name":"Everyone has same name"},"foo":{"age":10,"name":"Everyone has same name"},"other_anchor":"This string will appear as the value of two keys."}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ByteSliceToJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteSliceToJSON() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("ByteSliceToJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
