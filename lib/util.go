package fsserver

import "path"

func PathParent(p string) string {
	if len(p) > 0 && p[len(p)-1] == '/' {
		p = p[0 : len(p)-1]
	}
	return path.Dir(p)
}
