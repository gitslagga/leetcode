package solution

func GenerateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // You can fill in left brackets
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // You can fill in right bracket
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}
