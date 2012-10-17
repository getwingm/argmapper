package argmapper

type Map map[string]string

func New(args []string) (m Map) {
	m = make(Map)

	if len(args) == 0 {
		return
	}

nextarg:
	for i, s := range args {

		// does s look like an option?
		if len(s) > 1 && s[0] == '-' {
			k := ""
			v := ""

			num_minuses := 1
			if s[1] == '-' {
				num_minuses++
			}

			k = s[num_minuses:]
			if len(k) == 0 || k[0] == '-' || k[0] == '=' {
				continue nextarg
			}

			for i := 1; i < len(k); i++ { // equals cannot be first
				if k[i] == '=' {
					v = k[i+1:]
					k = k[0:i]
					break
				}
			}

			// It must have a value, which might be the next argument.
			remaining := args[i+1:]
			if v == "" && len(remaining) > 0 {
				v = remaining[0] // value is the next arg
			}

			m[k] = v
		}
	}

	return m
}
