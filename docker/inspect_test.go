package docker

import "testing"

func TestsplitHostname(t *testing.T) {
	var crcthostnames = []struct {
		a, b, c string
	}{
		{"localhost:5000/hello-world", "localhost:5000", "hello-world"},
		{"myregistrydomain:5000/java", "myregistrydomain:5000", "java"},
		{"docker.io/busybox", "docker.io", "busybox"},
	}
	var wrnghostnames = []struct {
		d, e, f string
	}{
		{"localhost:5000,hello-world", "localhost:5000", "hello-world"},
		{"myregistrydomain:5000&java", "myregistrydomain:5000", "java"},
		{"docker.io@busybox", "docker.io", "busybox"},
	}

	for _, i := range crcthostnames {
		res1, res2 := splitHostname(i.a)
		if res1 != i.b || res2 != i.c {
			t.Errorf("%s is an invalid hostname", i.a)
		}

		for _, j := range wrnghostnames {
			res1, res2 := splitHostname(i.a)
			if res1 == j.e || res2 == j.f {
				t.Errorf("%s is an invalid hostname", j.d)
			}

		}
	}
}

func Testvalidatename(t *testing.T) {
        var crctnames = []struct {
                a string
        }{
                {"localhost:5000/hello-world"},
                {"myregistrydomain:5000/java"},
                {"docker.io/busybox"},
        }
        var wrngnames = []struct {
                b string
        }{
                {"localhost:5000,hello-world"},
                {"myregistrydomain:5000&java"},
                {"docker.io@busybox"},
        }

        for _, i := range crctnames {
                res := validateName(i.a)
                if res != nil{
                        t.Errorf("%s is an invalid name", i.a)
                }

                for _, j := range wrngnames {
                        res := validateName(j.b)
                        if res == nil {
                                t.Errorf("%s is an invalid name", j.b)
                        }

                }
        }
}

