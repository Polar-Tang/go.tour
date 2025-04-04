package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

	container := [3]string{primary, secondary, tertiary}
	costs := [3]int{len(primary), len(secondary) + len(primary), len(secondary) + len(primary) + len(tertiary)}
	return container, costs
}
