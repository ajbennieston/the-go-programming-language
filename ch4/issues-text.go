// issues: Print a table of GitHub issues matching the search terms.

package main

import (
    "log"
    "os"
    "text/template"
    "time"
    "gopl.io/ch4/github"
)

const (
    Day = 24 * time.Hour
    Month = 30 * Day
    Year = 365 * Day
)

func main() {
    const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------------
Number: {{.Number}}
User: {{.User}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    var report = template.Must(template.New("issuelist").
        Funcs(template.FuncMap{"daysAgo": daysAgo}).
        Parse(templ))
    if err := report.Execute(os.Stdout, result); err != nil {
        log.Fatal(err)
    }
}

func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}
