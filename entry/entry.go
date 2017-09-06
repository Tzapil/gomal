package entry

import (
    "golang.org/x/net/html"

    "github.com/tzapil/anime/parse"
    q "github.com/tzapil/anime/query"

    "strings"
)

const site_adress = "http://www.fansubs.ru/"

type Link struct {
    Name string `json:"name"`
    Link string `json:"link"`
}

func NewLink(n string, l string) *Link {
    return &Link{n, l}
}

type Member struct {
    Name string         `json:"name"`
    Info string         `json:"info"`
    Link string         `json:"link"`
}

func NewMember(n string, i string, l string) *Member {
    return &Member{n, i, l}
}

type Translation struct {
    Team []*Member
    Download string     `json:"download"`
    Info string         `json:"info"`
    Date string         `json:"date"`
    Format string       `json:"format"`
}

func NewTranslation(m []*Member, d string, info string, date string, format string) *Translation {
    return &Translation{m, d, info, date, format}
}

type Entry struct {
    Name string         `json:"name"`
    Information string  `json:"information"`
    Links []*Link       `json:"links"`
    Trs []*Translation  `json:"trs"`
}

func NewEntry(n string, i string) *Entry {
    return &Entry{n, i, make([]*Link, 0), make([]*Translation, 0)}
}

func ParseMember(body *html.Node) *Member {
    var result *Member = nil
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Find(body)

    for i := range r {
        node := r[i].FirstChild
        if node != nil {
            // Team member with info
            if node.Type == html.TextNode {
                a := node.NextSibling
                if a != nil && a.Data == "a" {
                    result = NewMember(a.FirstChild.FirstChild.Data, strings.TrimRight(node.Data, ": "), site_adress + parse.GetHref(a))
                }

                break
            }

            // Team member without info
            if node.Data == "a" {
                result = NewMember(node.FirstChild.FirstChild.Data, "", site_adress + parse.GetHref(node))
                break
            }
        }
    }


    return result
}

func get_date(head *html.Node) string {
    result := ""
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Find(head)

    if len(r) >= 5 {
        fb := r[4].FirstChild
        if fb != nil && fb.Type == html.TextNode {
            result = fb.Data
        }
    }

    return result
}

func get_info(head *html.Node) string {
    result := ""
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("b")).
         Find(head)

    if len(r) > 0 {
        fb := r[0].FirstChild
        if fb != nil && fb.Type == html.TextNode {
            result = fb.Data
        }
    }

    return result
}

func get_format(head *html.Node) string {
    result := ""
    r := q.NewQuery().
         Child(q.NewTag("tbody")). 
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("a")).
         Child(q.NewTag("font")).
         Find(head)

    if len(r) > 0 {
        fb := r[0].FirstChild
        if fb != nil && fb.Type == html.TextNode {
            result = fb.Data
        }
    }

    return result
}

func get_download(head *html.Node) string {
    result := ""
    base := ""
    srt := ""

    r := q.NewQuery().Child(q.NewTag("form")).Find(head)

    if len(r) > 0 {
        base = parse.GetAttr(r[0], "action")
    }

    r = q.NewQuery().Child(q.NewTag("input")).Find(head)

    if len(r) > 0 {
        srt = parse.GetAttr(r[0], "value")
    }

    if base != "" && srt != "" {
        result = site_adress + base + "?srt=" + srt
    }

    return result
}

func ParseTranslation(head *html.Node, body []*html.Node) *Translation {
    team := make([]*Member, 0)
    
    for _, b := range body {
        tr := ParseMember(b)
        if tr != nil {
            team = append(team, tr)
        }
    }

    return NewTranslation(team, get_download(head), get_info(head), get_date(head), get_format(head))
}

func parse_links(node *html.Node) []*Link {
    result := make([]*Link, 0)
    r := q.NewQuery().
         Child(q.NewTag("table")).
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("a")).
         Find(node)

    for _, a := range r {
        result = append(result, NewLink(a.FirstChild.Data, parse.GetHref(a)))
    }

    return result
}

func ParseLinks(node *html.Node) []*Link {
    result := make([]*Link, 0)
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("b")).
         Contains("Ссылки:").
         Find(node)

    if len(r) > 0 {
        fb := r[0]
        blq := fb.NextSibling
        if blq != nil && blq.Data == "blockquote" {
            result = parse_links(blq)
        }
    }

    return result
}

func ParseTitle(node *html.Node) string {
    result := ""
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("b")).
         Find(node)

    if len(r) > 0 {
        fb := r[0]
        blq := fb.FirstChild
        if blq != nil && blq.Type == html.TextNode {
            result = blq.Data
        }
    }

    return result
}

func ParseInformation(node *html.Node) string {
    result := ""
    r := q.NewQuery().
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("b")).
         Contains("Общая информация:").
         Find(node)

    if len(r) > 0 {
        fb := r[0]
        blq := fb.NextSibling
        if blq != nil && blq.Data == "blockquote" {
            result = blq.FirstChild.Data
        }
    }

    return result
}

func ParseEntry(node *html.Node) *Entry {
    result := NewEntry(ParseTitle(node), ParseInformation(node))
    result.Links = append(result.Links, ParseLinks(node)...)
    return result
}

func Parse(root *html.Node) *Entry {
    var result *Entry = nil

    r := q.NewQuery().
         Child(q.NewTag("html")).
         Child(q.NewTag("body")).
         Child(q.NewTag("table")).
         Child(q.NewTag("tbody")).
         Child(q.NewTag("tr")).
         Child(q.NewTag("td")).
         Child(q.NewTag("table")).
         Find(root)

    if len(r) > 7 {
        // Tables with content
        information := r[4]
        result = ParseEntry(information)

        // Translators
        for i := 6; i < len(r); {
            counter := 1
            node := r[i]
            if parse.GetAttr(node, "width") == "750" {
                j := i + 2
                for ;j < len(r) && parse.GetAttr(r[j], "width") != "750"; j++ {}
                counter := j - i

                result.Trs = append(result.Trs, ParseTranslation(r[i], r[i + 1: i + counter]))
            }
            i += counter
        }
    }

    return result
}