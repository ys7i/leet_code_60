package main

import "strings"

func numUniqueEmails(emails []string) int {
    forwardedMailSet := make(map[string]struct{})

    for _, email := range emails {
        forwardedMail := formatToForwardedMail(email)

        forwardedMailSet[forwardedMail] = struct{}{}
    }

    return len(forwardedMailSet)
}

func formatToForwardedMail(email string) string {
    localNameAndDomain := strings.Split(email, "@")
    localName := localNameAndDomain[0]
    domain := localNameAndDomain[1]
    
    localName = strings.ReplaceAll(localName, ".", "")
    localName = strings.Split(localName, "+")[0]

    return localName + "@" + domain
}