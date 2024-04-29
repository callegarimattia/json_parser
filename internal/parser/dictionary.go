package parser

// STRUCTURAL CHARACTERS
const (
	// BEGIN_ARRAY is a rune that represents the start of an array.
	BEGIN_ARRAY = '['
	// END_ARRAY is a rune that represents the end of an array.
	END_ARRAY = ']'
	// BEGIN_OBJECT is a rune that represents the start of an object.
	BEGIN_OBJECT = '{'
	// END_OBJECT is a rune that represents the end of an object.
	END_OBJECT = '}'
	// NAME_SEPARATOR is a rune that represents the name separator.
	NAME_SEPARATOR = ':'
	// VALUE_SEPARATOR is a rune that represents the value separator.
	VALUE_SEPARATOR = ','
)

// SPACES
const (
	// SPACE is a rune that represents a space.
	SPACE = ' '
	// TAB is a rune that represents a tab.
	TAB = '\t'
	// NEWLINE is a rune that represents a newline.
	NEWLINE = '\n'
	// CARRIAGE_RETURN is a rune that represents a carriage return.
	CARRIAGE_RETURN = '\r'
)

// VALUES
const (
	// NULL is a string that represents a null value.
	NULL = "null"
	// TRUE is a string that represents a true value.
	TRUE = "true"
	// FALSE is a string that represents a false value.
	FALSE = "false"
	// STRING is a string that represents a string value.
	STRING = "string"
	// NUMBER is a string that represents a number value.
	NUMBER = "number"
	// ARRAY is a string that represents an array value.
	ARRAY = "array"
	// OBJECT is a string that represents an object value.
	OBJECT = "object"
)
