package canonicalizer

import (
	"github.com/djimenez/iconv-go"
	"os/exec"
	"regexp"
	"strings"
)

var transliterationMap = [][]string{
	{
		// cyrillic
		"Щ", "щ", "Ё", "Ж", "Х", "Ц", "Ч", "Ш", "Ю", "я", "ё", "ж", "х", "ц", "ч", "ш", "ю", "я", "А", "Б", "В",
		"Г", "Д", "Е", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Ь", "Ы", "Ъ", "Э",
		"а", "б", "в", "г", "д", "е", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "ь",
		"ы", "ъ", "э",
		// french
		"Ï", "ï", "Ÿ", "ÿ", "Ê", "ê", "À", "à", "È", "è", "Ù", "ù", "Û", "û",
		// spanish
		"Ñ", "ñ",
	},
	{
		// cyrillic
		"Sc", "sc", "Jo", "Z", "Ch", "C", "C", "S", "Ju", "ja", "jo", "z", "ch", "c", "c", "s", "ju", "ja", "A",
		"B", "V", "G", "D", "E", "Z", "I", "Y", "K", "L", "M", "N", "O", "P", "R", "S", "T", "U", "F", "", "Y", "",
		"E", "a", "b", "v", "g", "d", "e", "z", "i", "j", "k", "l", "m", "n", "o", "p", "r", "s", "t", "u", "f",
		"", "y", "", "e",
		// french
		"I", "i", "Y", "y", "E", "e", "A", "a", "E", "e", "U", "u", "U", "u",
		//spanish
		"N", "n",
	},
}

const unicodeSpecialCharBlacklist = "[^\x09\x0A\x0D\x20-\x7E\u00A0-\u02FF\u0370-\U0010FFFF]"

var specialCharBlacklist = "[`'\"^~]"

var glibcW1250CharMap = [][]string{
	{
		"\xa5", "\xa3", "\xbc", "\x8c", "\xa7", "\x8a", "\xaa", "\x8d", "\x8f", "\x8e", "\xaf", "\xb9", "\xb3", "\xbe",
		"\x9c", "\x9a", "\xba", "\x9d", "\x9f", "\x9e", "\xbf", "\xc0", "\xc1", "\xc2", "\xc3", "\xc4", "\xc5", "\xc6",
		"\xc7", "\xc8", "\xc9", "\xca", "\xcb", "\xcc", "\xcd", "\xce", "\xcf", "\xd0", "\xd1", "\xd2", "\xd3", "\xd4",
		"\xd5", "\xd6", "\xd7", "\xd8", "\xd9", "\xda", "\xdb", "\xdc", "\xdd", "\xde", "\xdf", "\xe0", "\xe1", "\xe2",
		"\xe3", "\xe4", "\xe5", "\xe6", "\xe7", "\xe8", "\xe9", "\xea", "\xeb", "\xec", "\xed", "\xee", "\xef", "\xf0",
		"\xf1", "\xf2", "\xf3", "\xf4", "\xf5", "\xf6", "\xf8", "\xf9", "\xfa", "\xfb", "\xfc", "\xfd", "\xfe", "\x96",
	},
	{
		"A", "L", "L", "S", "S", "S", "S", "T", "Z", "Z", "Z", "a", "l", "l", "s", "s", "s", "t", "z", "z", "z", "R",
		"A", "A", "A", "A", "L", "C", "C", "C", "E", "E", "E", "E", "I", "I", "D", "D", "N", "N", "O", "O", "O", "O",
		"x", "R", "U", "U", "U", "U", "Y", "T", "s", "r", "a", "a", "a", "a", "l", "c", "c", "c", "e", "e", "e", "e",
		"i", "i", "d", "d", "n", "n", "o", "o", "o", "o", "r", "u", "u", "u", "u", "y", "t", "-",
	},
}

const defaultSeparator = "-"

const defaultSuffix = ""

const defaultMaxLength = 0

func zip(slice1, slice2 []string) []string {
	zipped := make([]string, 2*len(slice1))
	for index := range slice1 {
		zipped[index*2] = slice1[index]
		zipped[index*2+1] = slice2[index]
	}
	return zipped
}

func isIconvLibiconv() bool {
	output, err := exec.Command("iconv", "--version").Output()
	if nil != err {
		return false
	}
	return strings.Contains(string(output), "libiconv")
}

var transliterationMapZipped = zip(transliterationMap[0], transliterationMap[1])
var glibcW1250CharMapZipped = zip(glibcW1250CharMap[0], glibcW1250CharMap[1])

type Canonicalizer struct {
	afterCallback  func(string) string
	beforeCallback func(string) string
	maxLength      int
}

func create(maxLength int, beforeCallback func(string) string, afterCallback func(string) string) *Canonicalizer {
	return &Canonicalizer{
		maxLength:      maxLength,
		beforeCallback: beforeCallback,
		afterCallback:  afterCallback,
	}
}

func New() *Canonicalizer {
	return create(defaultMaxLength, nil, nil)
}

func NewWithMaxLength(maxLength int) *Canonicalizer {
	return create(maxLength, nil, nil)
}

func NewWithCallbacks(beforeCallback func(string) string, afterCallback func(string) string) *Canonicalizer {
	return create(defaultMaxLength, beforeCallback, afterCallback)
}

func NewWithMaxLengthAndCallbacks(maxLength int, beforeCallback func(string) string, afterCallback func(string) string) *Canonicalizer {
	return create(maxLength, beforeCallback, afterCallback)
}

func (c *Canonicalizer) SetBeforeCallback(beforeCallback func(string) string) {
	c.beforeCallback = beforeCallback
}

func (c *Canonicalizer) SetAfterCallback(afterCallback func(string) string) {
	c.afterCallback = afterCallback
}

func (c *Canonicalizer) toAscii(str string) string {
	// transliterate cyrillic and other special chars
	str = strings.NewReplacer(transliterationMapZipped...).Replace(str)
	// get rid of some unicode special chars like tabs etc.
	str = regexp.MustCompile(unicodeSpecialCharBlacklist).ReplaceAllString(str, "")
	// get rid of some special chars
	str = regexp.MustCompile(specialCharBlacklist).ReplaceAllString(str, "")
	// transliterate to ASCII/Win-1250
	if !isIconvLibiconv() {
		str, _ = iconv.ConvertString(str, "UTF-8", "WINDOWS-1250//TRANSLIT//IGNORE")
		str = strings.NewReplacer(glibcW1250CharMapZipped...).Replace(str)
		return str
	}
	str, _ = iconv.ConvertString(str, "UTF-8", "ASCII//TRANSLIT//IGNORE")
	// get rid of some special chars (again, since iconv implementations other than glibc might put some back in)
	str = regexp.MustCompile(specialCharBlacklist).ReplaceAllString(str, "")
	return str
}

func (c *Canonicalizer) createSuffix(str string, suffix string, separator string) string {
	suffix = separator + strings.Trim(suffix, separator)
	if c.maxLength > 0 {
		maxLength := c.maxLength - len(suffix)
		if len(str) > maxLength {
			substring := str[0:maxLength]
			str = strings.Trim(substring, separator)
		}
	}
	return str + suffix
}

func (c *Canonicalizer) canonicalize(str string, suffix string, separator string) string {
	if nil != c.beforeCallback {
		beforeCallback := c.beforeCallback
		str = beforeCallback(str)
	}
	str = c.toAscii(str)
	str = strings.ToLower(str)
	str = regexp.MustCompile("(?i)[^a-z0-9]+").ReplaceAllString(str, separator)
	str = strings.Trim(str, separator)
	if defaultSuffix != suffix {
		str = c.createSuffix(str, suffix, separator)
	}
	if nil != c.afterCallback {
		afterCallback := c.afterCallback
		str = afterCallback(str)
	}
	return str
}

func (c *Canonicalizer) Canonicalize(str string) string {
	return c.canonicalize(str, defaultSuffix, defaultSeparator)
}

func (c *Canonicalizer) CanonicalizeWithSuffix(str string, suffix string) string {
	return c.canonicalize(str, suffix, defaultSeparator)
}

func (c *Canonicalizer) CanonicalizeWithSeparator(str string, separator string) string {
	return c.canonicalize(str, defaultSuffix, separator)
}

func (c *Canonicalizer) CanonicalizeWithSeparatorAndSuffix(str string, separator string, suffix string) string {
	return c.canonicalize(str, suffix, separator)
}
