package main

func DebugLog(args ...string) []string {
	debug := []string{"[DEBUG]"}
	return append(debug, args...)
}

func InfoLog(args ...string) []string {
	info := []string{"[INFO]"}
	return append(info, args...)
}

func ErrorLog(args ...string) []string {
	err := []string{"[ERROR]"}
	return append(err, args...)
}
