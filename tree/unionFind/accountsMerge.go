package main

import (
	"fmt"
	"sort"
)

/**
721. Accounts Merge

Given a list of accounts where each element accounts[i] is a list of strings, where the first element
accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there
is some common email to both accounts. Note that even if two accounts have the same name, they may
belong to different people as people could have the same name. A person can have any number of accounts
initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format: the first element of each account
is the name, and the rest of the elements are emails in sorted order. The accounts themselves can be returned
in any order.

Example 1:

Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"]
,["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
Explanation:
The first and second John's are the same person as they have the common email "johnsmith@mail.com".
The third John and Mary are different people as none of their email addresses are used by other accounts.
We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.
Example 2:

Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],
["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]
Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],
["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]

Constraints:

1 <= accounts.length <= 1000
2 <= accounts[i].length <= 10
1 <= accounts[i][j].length <= 30
accounts[i][0] consists of English letters.
accounts[i][j] (for j > 0) is a valid email.
*/

/**
Brute force:
try out each email with all emails if it is same, store name and all emails at one place and return
TC - O(N^2*M) SC - O(1)
Where N is the number of accounts and M is the maximum number of emails in an account
*/

/**
1. get all unique email ids
2. union same emailIds with id, maker latter parent of prior
3. iterate over unique emails with Find(id) to create map for all email ids
4. iterate over map and merge and return
*/

type UnionFindAM struct {
	parent []int
}

func NewUnionFindAM(size int) *UnionFindAM {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFindAM{parent}
}

func (uf *UnionFindAM) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFindAM) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		uf.parent[rootY] = rootX
	}
}

// TC - O(ALogA), SC - O(N+A)
func accountsMerge(accounts [][]string) [][]string {
	emailToID := make(map[string]int)
	uf := NewUnionFind(len(accounts))

	for i, account := range accounts {
		for _, email := range account[1:] {
			if owner, exists := emailToID[email]; exists {
				uf.Union(i, owner)
			}
			emailToID[email] = i
		}
	}

	rootToEmails := make(map[int]map[string]struct{})
	for email, id := range emailToID {
		root := uf.Find(id)
		if _, exists := rootToEmails[root]; !exists {
			rootToEmails[root] = make(map[string]struct{})
		}
		rootToEmails[root][email] = struct{}{}
	}

	var mergedAccounts [][]string
	for id, emailsSet := range rootToEmails {
		accountName := accounts[id][0]
		emails := make([]string, 0, len(emailsSet))
		for email := range emailsSet {
			emails = append(emails, email)
		}
		sort.Strings(emails)
		mergedAccounts = append(mergedAccounts, append([]string{accountName}, emails...))
	}

	return mergedAccounts
}

func mainAM() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "john_newyork@mail.com", "john_newyorktes@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}
