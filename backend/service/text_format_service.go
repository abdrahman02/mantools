package service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strings"
)

func FormatJSON(input string) (string, error) {
	var buf bytes.Buffer
	var obj interface{}
	dec := json.NewDecoder(strings.NewReader(input))
	dec.UseNumber()
	if err := dec.Decode(&obj); err != nil {
		return input, err
	}
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	if err := enc.Encode(obj); err != nil {
		return input, err
	}
	// Remove trailing newline added by Encode
	return strings.TrimRight(buf.String(), "\n"), nil
}

func FormatXML(input string) (string, error) {
	decoder := xml.NewDecoder(strings.NewReader(input))
	var out bytes.Buffer
	indent := 0
	writeIndent := func() {
		for i := 0; i < indent; i++ {
			out.WriteString("  ")
		}
	}

	for {
		tok, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			// On parse error, return original input so we don't lose content
			return input, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			writeIndent()
			out.WriteString("<" + t.Name.Local)
			for _, attr := range t.Attr {
				out.WriteString(" " + attr.Name.Local + "=\"" + attr.Value + "\"")
			}
			out.WriteString(">\n")
			indent++
		case xml.EndElement:
			indent--
			writeIndent()
			out.WriteString("</" + t.Name.Local + ">\n")
		case xml.CharData:
			text := strings.TrimSpace(string([]byte(t)))
			if len(text) > 0 {
				writeIndent()
				out.WriteString(text + "\n")
			}
		case xml.Comment:
			writeIndent()
			out.WriteString("<!--" + string(t) + "-->\n")
		case xml.ProcInst:
			writeIndent()
			out.WriteString("<?" + t.Target + " " + string(t.Inst) + "?>\n")
		default:
			// ignore
		}
	}

	return strings.TrimRight(out.String(), "\n"), nil
}