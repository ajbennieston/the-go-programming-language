// issues: Print a table of GitHub issues matching the search terms.

package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "gopl.io/ch4/github"
)

const (
    Day = 24 * time.Hour
    Month = 30 * Day
    Year = 365 * Day
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    issueMap := make(map[string][]*github.Issue)
    for _, item := range result.Items {
        issueAge := time.Since(item.CreatedAt)
        switch {
        case issueAge < Day:
            issueMap["less than a day"] = append(issueMap["less than a day"], item)
        case issueAge < Month:
            issueMap["less than a month"] = append(issueMap["less than a month"], item)
        case issueAge < Year:
            issueMap["less than a year"] = append(issueMap["less than a year"], item)
        default:
            issueMap["more than a year"] = append(issueMap["more than a year"], item)
        }
    }

    ages := [...]string{"less than a day", "less than a month", "less than a year", "more than a year"}
    for _, age := range ages {
        printIssues(issueMap, age)
    }
}

func printIssues(issueMap map[string][]*github.Issue, age string) {
    if len(issueMap[age]) > 0 {
        fmt.Printf("Issues %s old:\n", age)
        for _, item := range issueMap[age] {
            fmt.Printf("  #%-5d %9.9s %.55s\n",
                       item.Number, item.User.Login, item.Title)
        }
        fmt.Println()
    }
}
