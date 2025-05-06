package main 

const mod = int64(1e9 + 7)

func bigmod(a, b int64) int64 {
	a %= mod 
	res := int64(1)
	for b > 0 {
		if b % 2 == 1 {
			res = res * a % mod 
		}
		a = a * a % mod 
	}
	return res 
}

func inverseMod(n int64) int64 {
	return bigmod(n, mod - 2) % mod 
}



