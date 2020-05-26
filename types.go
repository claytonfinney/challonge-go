package challonge

import (
	"bytes"
	"strconv"
	"time"
)

type (
	ChallongeString string
	ChallongeInt    int
	ChallongeInt64  int64
	ChallongeTime   time.Time
	ChallongeBool   bool
)

const (
	NULL_INT    int    = -16487502
	NULL_INT64  int64  = -16487502
	NULL_STRING string = "null"
	NULL_TIME   string = "0001-01-01T00:00:00Z"
)

func IsNull(in interface{}) bool {
	switch in.(type) {
	case int:
		if in == NULL_INT {
			return true
		} else {
			return false
		}
	case int64:
		if in == NULL_INT64 {
			return true
		} else {
			return false
		}
	case string:
		if in == NULL_STRING {
			return true
		} else {
			return false
		}
	case time.Time:
		if in.(time.Time).Format(time.RFC3339) == NULL_TIME {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

///////////////////////////////////
// ChallongeString
// A nullable string type
///////////////////////////////////

func (c ChallongeString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if string(c) == "null" {
		buf.WriteString(`null`)
	} else if string(c) == "" {
		buf.WriteString(`""`)
	} else {
		buf.WriteString(`"` + string(c) + `"`)
	}
	return buf.Bytes(), nil
}

func (c *ChallongeString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*c = "null"
		return nil
	}
	res := ChallongeString(str)
	if len(res) >= 2 {
		res = res[1 : len(res)-1]
	}
	*c = res
	return nil
}

///////////////////////////////////
// ChallongeInt
// A nullable int type
///////////////////////////////////

func (c ChallongeInt) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	s := strconv.Itoa(int(c))
	if int(c) == NULL_INT {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(`` + s + ``)
	}
	return buf.Bytes(), nil
}

func (c *ChallongeInt) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*c = ChallongeInt(NULL_INT)
		return nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*c = ChallongeInt(i)
	return nil
}

///////////////////////////////////
// ChallongeInt64
// A nullable int64 type
///////////////////////////////////

func (c ChallongeInt64) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	s := strconv.Itoa(int(c))
	if int64(c) == NULL_INT64 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(`` + s + ``)
	}
	return buf.Bytes(), nil
}

func (c *ChallongeInt64) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*c = ChallongeInt64(NULL_INT64)
		return nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*c = ChallongeInt64(i)
	return nil
}

///////////////////////////////////
// ChallongeBool
// A nullable bool type
///////////////////////////////////

func (c ChallongeBool) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	var s string
	if bool(c) {
		s = "true"
	} else {
		s = "false"
	}
	buf.WriteString(`` + s + ``)
	return buf.Bytes(), nil
}

func (c *ChallongeBool) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*c = ChallongeBool(false)
		return nil
	} else if str == "true" {
		*c = ChallongeBool(true)
		return nil
	} else {
		*c = ChallongeBool(false)
		return nil
	}
}

///////////////////////////////////
// ChallongeTime
// A nullable time type
///////////////////////////////////

func (c ChallongeTime) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	sin := time.Time(c).Format(time.RFC3339)
	if sin == NULL_TIME {
		buf.WriteString(`null`)
		return buf.Bytes(), nil
	}
	buf.WriteString(`"` + sin + `"`)
	return buf.Bytes(), nil
}

func (c *ChallongeTime) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		tme, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z+08:00")
		*c = ChallongeTime(tme)
		return nil
	} else {
		sin := string(in[1 : len(in)-1])
		tme, err := time.Parse(time.RFC3339, sin)
		if err != nil {
			return err
		}
		*c = ChallongeTime(tme)
		return nil
	}
}

/////////////////////////////////
// Slice Conversions
////////////////////////////////

func normalizeChallongeStringSlice(sl []ChallongeString) []string {
	st := make([]string, 0)
	for _, s := range sl {
		st = append(st, string(s))
	}
	return st
}

func denormalizeStringSlice(sl []string) []ChallongeString {
	st := make([]ChallongeString, 0)
	for _, s := range sl {
		st = append(st, ChallongeString(s))
	}
	return st
}
