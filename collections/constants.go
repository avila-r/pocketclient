package collections

const (
	TypeBase CollectionType = "base"
	TypeAuth CollectionType = "auth"
	TypeView CollectionType = "view"
)

const (
	Text             FieldType = "text"              // e.g., "", "example"
	Editor           FieldType = "editor"            // e.g., "", "<p>example</p>"
	Number           FieldType = "number"            // e.g., 0, -1, 1, 1.5 (+, - operations supported)
	Bool             FieldType = "bool"              // e.g., false, true
	Email            FieldType = "email"             // e.g., "", "test@example.com"
	URL              FieldType = "url"               // e.g., "", "https://example.com"
	Date             FieldType = "date"              // e.g., "", "2022-01-01 00:00:00.000Z"
	SelectSingle     FieldType = "select_single"     // e.g., "", "optionA"
	SelectMultiple   FieldType = "select_multiple"   // e.g., [], ["optionA", "optionB"] (+, - operations supported)
	RelationSingle   FieldType = "relation_single"   // e.g., "", "JJ2YRU30FBG8MqX"
	RelationMultiple FieldType = "relation_multiple" // e.g., [], ["JJ2YRU30FBG8MqX", "eP2jCr1h3NGtsbz"] (+, - operations supported)
	FileSingle       FieldType = "file_single"       // e.g., "", "example123_Ab24ZjL.png"
	FileMultiple     FieldType = "file_multiple"     // e.g., [], ["file1_Ab24ZjL.png", "file2_Frq24ZjL.txt"] (- operation supported)
	JSON             FieldType = "json"              // e.g., any JSON value
)
