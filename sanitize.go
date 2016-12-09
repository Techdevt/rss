package main

import (
  "strings"
  "bytes"
  
  "html"
  "html/template"
)

func SanitizeUTF8 (s string) string {
  output := ""

  if !strings.ContainsAny(s, "<>") {
    output = s
  } else {

    s = strings.Replace(s, "\n", "", -1)

    s = strings.Replace(s, "</p>", "\n", -1)
    s = strings.Replace(s, "<br>", "\n", -1)
    s = strings.Replace(s, "</br>", "\n", -1)
    s = strings.Replace(s, "<br/>", "\n", -1)
    s = strings.Replace(s, "<br />", "\n", -1)

    b := bytes.NewBufferString("")
    inTag := false
    for _, r := range s {
      switch r {
      case '<':
        inTag = true
      case '>':
        inTag = false
      default:
        if !inTag {
          b.WriteRune(r)
        }
      }
    }
    output = b.String()
  }

  output = strings.Replace(output, "&#8216;", "'", -1)
  output = strings.Replace(output, "&#8217;", "'", -1)
  output = strings.Replace(output, "&#8220;", "\"", -1)
  output = strings.Replace(output, "&#8221;", "\"", -1)
  output = strings.Replace(output, "&nbsp;", " ", -1)
  output = strings.Replace(output, "&quot;", "\"", -1)
  output = strings.Replace(output, "&apos;", "'", -1)

  output = html.UnescapeString(output)

  output = template.HTMLEscapeString(output)

  output = strings.Replace(output, "&#34;", "\"", -1)
  output = strings.Replace(output, "&#39;", "'", -1)
  output = strings.Replace(output, "&amp; ", "& ", -1)
  output = strings.Replace(output, "&amp;amp; ", "& ", -1)

  return output
}