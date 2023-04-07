package gunit

func True(t T, actual bool) {
	t.Helper()
	if actual != true {
		t.Errorf("got false, wanted true")
	}
}

func False(t T, actual bool) {
	t.Helper()
	if actual != false {
		t.Errorf("got true, wanted false")
	}
}
