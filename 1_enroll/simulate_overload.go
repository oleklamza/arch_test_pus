// ---------------------------------------------------------------------------
// Usage (on top of the handler):
// if !simulateOverload(w) {
// 	return
// }
// ---------------------------------------------------------------------------
func simulateOverload(w http.ResponseWriter) bool {
	n := atomic.AddInt64(&load, 1)
	defer atomic.AddInt64(&load, -1)

	d := *delay
	if n > *slowAt {
		d += time.Duration(n - *slowAt) * time.Duration(10 + rand.Intn(10)) * time.Millisecond
	}

	time.Sleep(d)

	if n >= *errorAt && rand.Float64() < 0.15 {
		http.Error(w, `{"error":"service overloaded"}`, http.StatusServiceUnavailable)
		return false
	}

	return true
}